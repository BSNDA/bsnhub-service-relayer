package hub

import (
	"context"
	"encoding/json"
	"fmt"
	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
	"relayer/core"
	"relayer/logging"
	"relayer/mysql"
	"strconv"
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

	KeyPath    string
	KeyName    string
	Passphrase string

	ServiceInfo   ServiceInfo
	ServiceClient servicesdk.ServiceClient
}

// NewIritaHubChain constructs a new Irita-Hub chain
func NewIritaHubChain(
	chainID string,
	nodeRPCAddr string,
	nodeGRPCAddr string,
	keyPath string,
	keyName string,
	passphrase string,
	serviceName string,
	schemas string,
	provider string,
	serviceFee string,
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

	if len(keyPath) == 0 {
		keyPath = defaultKeyPath
	}

	if len(serviceName) == 0 {
		keyPath = defaultServiceName
	}

	if len(schemas) == 0 {
		keyPath = defaultSchemas
	}

	if len(provider) == 0 {
		keyPath = defaultProvider
	}

	if len(serviceFee) == 0 {
		keyPath = defaultServiceFee
	}

	if qos == 0 {
		qos = defaultQoS
	}

	fee, err := types.ParseDecCoins(defaultFee)
	if err != nil {
		panic(err)
	}

	config := types.ClientConfig{
		NodeURI:  nodeRPCAddr,
		GRPCAddr: nodeGRPCAddr,
		ChainID:  chainID,
		Gas:      defaultGas,
		Fee:      fee,
		Mode:     defaultBroadcastMode,
		Algo:     defaultKeyAlgorithm,
		KeyDAO:   store.NewFileDAO(keyPath),
		Level:    "debug",
	}

	hub := IritaHubChain{
		ChainID:     chainID,
		NodeRPCAddr: nodeRPCAddr,
		KeyPath:     keyPath,
		KeyName:     keyName,
		Passphrase:  passphrase,
		ServiceInfo: ServiceInfo{
			ServiceName: serviceName,
			Schemas:     schemas,
			Provider:    provider,
			ServiceFee:  serviceFee,
			QoS:         qos,
		},
		ServiceClient: servicesdk.NewServiceClient(config),
	}

	return hub
}

// BuildIritaHubChain builds an Irita-Hub instance from the given config
func BuildIritaHubChain(config Config) IritaHubChain {
	return NewIritaHubChain(
		config.ChainID,
		config.NodeRPCAddr,
		config.NodeGRPCAddr,
		config.KeyPath,
		config.KeyName,
		config.Passphrase,
		config.ServiceName,
		config.Schemas,
		config.Provider,
		config.ServiceFee,
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
) error {
	invokeServiceReq, err := ic.BuildServiceInvocationRequest(request)
	if err != nil {
		return err
	}

	reqCtxID, resTx, err := ic.ServiceClient.InvokeService(invokeServiceReq, ic.BuildBaseTx())
	if err != nil {
		mysql.TxErrCollection(request.ID, err.Error())
		return err
	}

	logging.Logger.Infof("request context created on %s: %s", ic.ChainID, reqCtxID)

	requests, err := ic.ServiceClient.QueryRequestsByReqCtx(reqCtxID, 1)
	if err != nil {
		return err
	}

	if len(requests) == 0 {
		return fmt.Errorf("no service request initiated on %s", ic.ChainID)
	}

	// TODO
	mysql.OnInterchainRequestSent(request.ID, requests[0].ID, resTx.Hash)

	logging.Logger.Infof("service request initiated on %s: %s", ic.ChainID, requests[0].ID)

	return ic.ResponseListener(reqCtxID, requests[0].ID, cb)
}

// BuildServiceInvocationRequest builds the service invocation request from the given interchain request
func (ic IritaHubChain) BuildServiceInvocationRequest(
	request core.InterchainRequest,
) (service.InvokeServiceRequest, error) {
	serviceFeeCap, err := types.ParseDecCoins(ic.ServiceInfo.ServiceFee)
	if err != nil {
		return service.InvokeServiceRequest{}, err
	}
	var methodAndArgs MethodAndArgs
	inputStr,err:= strconv.Unquote("\""+request.MethodAndArgs+"\"")
	err = json.Unmarshal([]byte(inputStr), &methodAndArgs)
	if err != nil {
		return service.InvokeServiceRequest{}, err
	}
	ArgsByte, err := json.Marshal(methodAndArgs.Args)
	if err != nil {
		return service.InvokeServiceRequest{}, err
	}

	input := ServiceInput{
		Header: Header{
			ReqSequence: request.ID,
			ChainID:     request.ChainID,
		},
		Body: Body{
			Source: Source{
				ID:     ic.ChainID,
				Sender: request.Sender,
				TxHash: request.TxHash,
			},
			Dest: Dest{
				ID:              request.ChainID,
				EndpointType:    request.EndpointType,
				EndpointAddress: request.EndpointAddress,
			},
			method: methodAndArgs.Method,
			args: string(ArgsByte),
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
	response, err := ic.ServiceClient.QueryServiceResponse(requestID)
	if response.RequestContextID == reqCtxID {
		resp := core.ResponseAdaptor{
			StatusCode: 200,
			Result:     response.Result,
			Output:     response.Output,
		}

		cb(requestID, resp)

		return nil
	}

	callbackWrapper := func(reqCtxID, requestID, result string, response string) {
		resp := core.ResponseAdaptor{
			StatusCode: 200,
			Result:     result,
			Output:     response,
		}

		cb(requestID, resp)
	}

	logging.Logger.Infof("waiting for the service response on %s", ic.ChainID)

	subscription, err := ic.ServiceClient.SubscribeServiceResponse(reqCtxID, callbackWrapper)
	if err != nil {
		return err
	}

	go func() {
		for {
			reqCtx, err := ic.ServiceClient.QueryRequestContext(reqCtxID)
			status, err2 := ic.ServiceClient.Status(context.Background())
			req, err3 := ic.ServiceClient.QueryServiceRequest(requestID)
			if err != nil || err2 != nil || err3 != nil || reqCtx.BatchState == "BATCH_COMPLETED" || status.SyncInfo.LatestBlockHeight > req.ExpirationHeight {
				_ = ic.ServiceClient.Unsubscribe(subscription)
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
