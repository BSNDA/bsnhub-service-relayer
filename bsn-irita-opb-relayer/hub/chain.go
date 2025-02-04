package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/iritamod-sdk-go/service"
	"github.com/irisnet/core-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	storetypes "github.com/irisnet/core-sdk-go/types/store"
	log "github.com/sirupsen/logrus"
	"relayer/common"
	"relayer/core"
	"relayer/logging"
	"time"
)

type ServiceInfo struct {
	ServiceName string
	Schemas     string
	Provider    string
	ServiceFee  string
	QoS         uint64
}

// IritaHubChain defines the Irita-Hub chain
type IritaHubChain struct {
	ChainID     string
	NodeRPCAddr string

	KeyName    string
	Passphrase string

	ServiceInfo ServiceInfo
	IritaClient *ServiceClient
}

// NewIritaHubChain constructs a new Irita-Hub chain
func NewIritaHubChain(
	chainID string,
	nodeRPCAddr string,
	nodeGRPCAddr string,
	keyMode string,
	keyPath string,
	keyName string,
	passphrase string,
	keyArmor string,
	txFee string,
	serviceName string,
	schemas string,
	provider string,
	serviceFee string,
	timeout uint,
	qos uint64,
) IritaHubChain {
	if len(chainID) == 0 {
		chainID = defaultChainID
	}

	if len(nodeRPCAddr) == 0 {
		nodeRPCAddr = defaultNodeRPCAddr
	}

	if len(nodeGRPCAddr) == 0 {
		nodeGRPCAddr = defaultNodeGRPCAddr
	}

	if len(serviceName) == 0 {
		serviceName = defaultServiceName
	}

	if len(schemas) == 0 {
		schemas = defaultSchemas
	}

	if len(provider) == 0 {
		provider = defaultProvider
	}

	if len(serviceFee) == 0 {
		serviceFee = defaultServiceFee
	}

	if timeout == 0 {
		timeout = defaultTimeout
	}

	if qos == 0 {
		qos = defaultQoS
	}

	if len(txFee) == 0 {
		txFee = defaultFee
	}

	fee, err := types.ParseDecCoins(txFee)
	if err != nil {
		panic(err)
	}
	var keyDAO storetypes.KeyDAO
	if keyMode == "mem" {
		keyDAO = storetypes.NewMemory(nil)
	} else {
		keyDAO = storetypes.NewFileDAO(keyPath)
	}

	config, err := sdk.NewClientConfig(
		nodeRPCAddr,
		nodeGRPCAddr,
		chainID,
		sdk.FeeOption(fee),
		sdk.GasOption(defaultGas),
		sdk.ModeOption(defaultBroadcastMode),
		sdk.AlgoOption(defaultKeyAlgorithm),
		sdk.KeyDAOOption(keyDAO),
		sdk.TimeoutOption(5),
	)
	hub := IritaHubChain{
		ChainID:     chainID,
		NodeRPCAddr: nodeRPCAddr,
		KeyName:     keyName,
		Passphrase:  passphrase,
		ServiceInfo: ServiceInfo{
			ServiceName: serviceName,
			Schemas:     schemas,
			Provider:    provider,
			ServiceFee:  serviceFee,
			QoS:         qos,
		},
		IritaClient: NewServiceClient(config),
	}

	// import key
	if keyMode == "mem" {
		log.WithField("keyName", keyName).Info("use memory key dao, importing key...")
		addr, err := hub.ImportKey(keyName, passphrase, keyArmor)
		if err != nil {
			panic(err)
		}
		log.WithFields(log.Fields{"addr": addr}).Info("successfully import key in hub")
	}
	return hub
}

// BuildIritaHubChain builds an Irita-Hub instance from the given config
func BuildIritaHubChain(config Config) IritaHubChain {
	return NewIritaHubChain(
		config.ChainID,
		config.NodeRPCAddr,
		config.NodeGRPCAddr,
		config.KeyMode,
		config.KeyPath,
		config.KeyName,
		config.Passphrase,
		config.KeyArmor,
		config.Fee,
		config.ServiceName,
		config.Schemas,
		config.Provider,
		config.ServiceFee,
		config.Timeout,
		config.QoS,
	)
}

// GetChainID implements IritaHubChainI
func (ic IritaHubChain) GetChainID() string {
	return ic.ChainID
}

// SendInterchainRequest implements IritaHubChainI
func (ic IritaHubChain) SendInterchainRequest(
	request core.InterchainRequest,
	cb core.ResponseCallback,
) (core.InterchainRequestInfo, error) {

	info := core.InterchainRequestInfo{}

	invokeServiceReq, err := ic.BuildServiceInvocationRequest(request)
	if err != nil {
		return info, err
	}

	reqCtxID, resTx, err := ic.IritaClient.Service.InvokeService(invokeServiceReq, ic.BuildBaseTx())
	if err != nil {
		//mysql.TxErrCollection(request.ID, err.Error())
		return info, err
	}
	info.HubReqTxId = resTx.Hash.String()
	logging.Logger.Infof("request context created on %s: %s", ic.ChainID, reqCtxID)

	requests, err := ic.IritaClient.Service.QueryRequestsByReqCtx(reqCtxID, 1, nil)
	if err != nil {
		return info, err
	}

	if len(requests) == 0 {
		return info, fmt.Errorf("no service request initiated on %s", ic.ChainID)
	}

	info.IcRequestId = requests[0].ID
	// TODO
	//mysql.OnInterchainRequestSent(request.ID, requests[0].ID, resTx.Hash)

	logging.Logger.Infof("service request initiated on %s: %s", ic.ChainID, requests[0].ID)

	return info, ic.ResponseListener(reqCtxID, requests[0].ID, cb)
}

// BuildServiceInvocationRequest builds the service invocation request from the given interchain request
func (ic IritaHubChain) BuildServiceInvocationRequest(
	request core.InterchainRequest,
) (service.InvokeServiceRequest, error) {
	serviceFeeCap, err := types.ParseDecCoins(ic.ServiceInfo.ServiceFee)
	destID := common.GetDestID(request.DestChainType, request.DestSubChainID, request.DestChainID)

	input := ServiceInput{
		Header: Header{
			ReqSequence: request.ID,
			ChainID:     request.SourceChainID,
		},
		Body: Body{
			Source: Source{
				ID:     request.SourceChainID,
				Sender: request.Sender,
				TxHash: request.TxHash,
			},
			Dest: Dest{
				ID:              destID,
				ChainID:         request.DestChainID,
				SubChainID:      request.DestSubChainID,
				EndpointType:    request.EndpointType,
				EndpointAddress: request.EndpointAddress,
			},
			Method:   request.Method,
			CallData: request.CallData,
		},
	}

	serviceInput, err := json.Marshal(input)
	if err != nil {
		return service.InvokeServiceRequest{}, err
	}

	return service.InvokeServiceRequest{
		ServiceName:   ic.ServiceInfo.ServiceName,
		Providers:     []string{ic.ServiceInfo.Provider},
		Input:         string(serviceInput),
		Timeout:       100,
		ServiceFeeCap: serviceFeeCap,
	}, nil
}

// ResponseListener gets and handles the response of the given request context ID by event subscription
func (ic IritaHubChain) ResponseListener(reqCtxID string, requestID string, cb core.ResponseCallback) error {
	response, err := ic.IritaClient.Service.QueryServiceResponse(requestID)
	if response.RequestContextID == reqCtxID {
		resp := core.ResponseAdaptor{
			StatusCode: 200,
			Result:     response.Result,
			Output:     response.Output,
		}

		cb(requestID, resp)

		return nil
	}

	callbackWrapper := func(reqCtxID, requestID, response string) {
		resp := core.ResponseAdaptor{
			StatusCode: 200,
			//Result:     result,
			Output: response,
		}

		cb(requestID, resp)
	}

	logging.Logger.Infof("waiting for the service response on %s", ic.ChainID)

	subscription, err := ic.IritaClient.Service.SubscribeServiceResponse(reqCtxID, callbackWrapper)
	if err != nil {
		return err
	}

	go func() {
		for {
			reqCtx, err := ic.IritaClient.Service.QueryRequestContext(reqCtxID)
			status, err2 := ic.IritaClient.Status(context.Background())
			req, err3 := ic.IritaClient.Service.QueryServiceRequest(requestID)
			if err != nil || err2 != nil || err3 != nil || reqCtx.BatchState == "BATCH_COMPLETED" || status.SyncInfo.LatestBlockHeight > req.ExpirationHeight {
				logging.Logger.Infof("HUB Unsubscribe RequestID is %s", requestID)
				_ = ic.IritaClient.Unsubscribe(subscription)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	return nil
}

// BuildBaseTx builds a base tx
func (ic IritaHubChain) BuildBaseTx() types.BaseTx {
	return types.BaseTx{
		From:     ic.KeyName,
		Password: ic.Passphrase,
	}
}
