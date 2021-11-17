package opb

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bianjieai/iritamod-sdk-go/wasm"
	"github.com/irisnet/core-sdk-go/types"
	sdktypes "github.com/irisnet/core-sdk-go/types"
	sdkstore "github.com/irisnet/core-sdk-go/types/store"
	log "github.com/sirupsen/logrus"
	abci "github.com/tendermint/tendermint/abci/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	txstore "relayer/appchains/opb/store"
	"relayer/hub"
	"strings"
	"time"

	"relayer/core"
	"relayer/logging"
	"relayer/store"
)

// opbChain defines the opb chain
type OpbChain struct {
	Config    Config
	OpbClient *hub.ServiceClient
	ChainID   string // unique chain ID

	store      *store.Store // store backend instance
	lastHeight int64        // last height

	done    bool                          // indicates if the chain monitor is done
	handler core.InterchainRequestHandler // handler for the interchain request
}

// NewFISCOChain constructs a new FISCOChain instance
func NewOpbChain(
	config Config,
	store *store.Store,
) (*OpbChain, error) {
	//将接口传递的节点名称通过配置转换为 节点地址，如果不在配置中，不转换
	//随机取一个传入的node
	nodeName := randURL(config.ChainParams.NodeURLs)
	var rpcAddr string
	var grpcAddr string
	//获取配置的nodeURL
	rpcAddrStr, ok := config.RpcAddrsMap[nodeName]
	if ok {
		rpcAddr = rpcAddrStr
	}
	grpcAddrStr, ok := config.GrpcAddrsMap[nodeName]
	if ok {
		grpcAddr = grpcAddrStr
	}
	fees, _ := sdktypes.ParseDecCoins(config.DefaultFee)

	options := []sdktypes.Option{
		sdktypes.CachedOption(true),
		sdktypes.KeyDAOOption(sdkstore.NewMemory(nil)),
		sdktypes.FeeOption(fees),
		sdktypes.GasOption(config.DefaultGas),
		sdktypes.TimeoutOption(config.Timeout),
		sdktypes.AlgoOption(defaultAlgo),
	}

	clientConfig, err := sdktypes.NewClientConfig(
		rpcAddr,
		grpcAddr,
		config.BaseConfig.ChainId,
		options...,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to init clientConfig: %s", err)
	}
	opbClient := hub.NewServiceClient(clientConfig)
	chainID := GetChainID(config.ChainParams)

	opb := &OpbChain{
		Config:    config,
		OpbClient: opbClient,
		ChainID:   chainID,
		store:     store,
		done:      true,
	}

	// import opb key
	addr, err := opb.OpbClient.Import(config.Account.KeyName, config.Account.Passphrase, config.Account.KeyArmor)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"addr": addr}).Info("import opb account success")

	err = opb.storeChainParams()
	if err != nil {
		return nil, err
	}

	err = opb.storeChainID()
	if err != nil {
		return nil, err
	}

	//chainHeight, err := store.GetInt64(HeightKey(chainID))
	//if err != nil {
	//	log.WithFields(log.Fields{"err_info": err.Error(), "chain_id": chainID}).Error("get chain height err when opb chain client is initializing")
	//	return nil, err
	//}
	//log.WithFields(log.Fields{"chain_height": chainHeight, "chain_id": chainID}).Info("finish initializing opb chain client")

	return opb, nil
}

// BuildFISCOChain builds a FISCOChain instance from the given chain params and store
func BuildOpbChain(
	chainParams []byte,
	store *store.Store,
) (*OpbChain, error) {
	var params ChainParams
	err := json.Unmarshal(chainParams, &params)
	if err != nil {
		return nil, err
	}

	baseCfgBz, err := store.Get(BaseConfigKey())
	if err != nil {
		return nil, err
	}

	var baseConfig BaseConfig
	err = json.Unmarshal(baseCfgBz, &baseConfig)
	if err != nil {
		return nil, err
	}

	config := Config{
		BaseConfig:  baseConfig,
		ChainParams: params,
	}

	return NewOpbChain(config, store)
}

// GetChainID implements AppChainI
func (opb *OpbChain) GetChainID() string {
	return opb.ChainID
}

// Start implements AppChainI
func (opb *OpbChain) Start(handler core.InterchainRequestHandler) error {
	if !opb.done {
		return fmt.Errorf("chain %s has been started", opb.ChainID)
	}

	opb.done = false
	opb.handler = handler

	go opb.monitor()

	logging.Logger.Infof("chain %s started", opb.ChainID)

	return nil
}

// Stop implements AppChainI
func (opb *OpbChain) Stop() error {
	logging.Logger.Infof("stopping chain %s", opb.ChainID)
	opb.done = true

	return nil
}

// GetHeight implements AppChainI
func (opb *OpbChain) GetHeight() int64 {
	return opb.lastHeight
}

// SendResponse implements AppChainI
func (opb *OpbChain) SendResponse(requestID string, response core.ResponseI) error {
	execAbi := wasm.NewContractABI().
		WithMethod("set_response").
		WithArgs("request_id", requestID).
		WithArgs("err_msg", response.GetErrMsg()).
		WithArgs("output", response.GetOutput())

	data := &txstore.RelayerResInfo{
		RequestId: requestID,
		TxStatus:  txstore.TxStatus_Success,
		ErrMsg:    "",
	}

	defer func(d *txstore.RelayerResInfo) {
		txstore.RelayerResponeRecord(d)
	}(data)

	resultTx, err := opb.OpbClient.WASM.Execute(opb.Config.ChainParams.IServiceCoreAddr, execAbi, nil, opb.BuildBaseTx())
	if err != nil {
		data.TxStatus = txstore.TxStatus_Error
		data.ErrMsg = fmt.Sprintf("call opb setResponse failed :%s", err)
		//mysql.TxErrCollection(requestID, err.Error())
		return err
	}
	data.FromResTxId = resultTx.Hash.String()
	// TODO
	//mysql.OnInterchainRequestResponseSent(requestID, resultTx.Hash)

	err = opb.waitForSuccess(resultTx.Hash.String(), "SetResponse")
	if err != nil {
		//mysql.TxErrCollection(requestID, err.Error())
		data.TxStatus = txstore.TxStatus_Error
		data.ErrMsg = fmt.Sprintf("call opb setResponse failed :%s", err)
		return err
	}

	// TODO
	//mysql.OnInterchainRequestSucceeded(requestID)

	return nil
}

// waitForSuccess waits for the receipt of the given tx
func (opb *OpbChain) waitForSuccess(txHash string, name string) error {
	logging.Logger.Infof("%s: transaction sent to %s, hash: %s", name, opb.GetChainID(), txHash)

	tx, _ := opb.OpbClient.QueryTx(txHash)
	if tx.TxResult.Code != 0 {
		return fmt.Errorf("transaction %s execution failed: %s", txHash, tx.TxResult.Log)
	}

	logging.Logger.Infof("%s: transaction %s execution succeeded", name, txHash)

	return nil
}

// BuildBaseTx builds a base tx
func (opb *OpbChain) BuildBaseTx() types.BaseTx {
	return types.BaseTx{
		From:     opb.Config.Account.KeyName,
		Password: opb.Config.Account.Passphrase,
		Mode:     sdktypes.Commit,
	}
}

// buildInterchainRequest builds an interchain request from the interchain event
func (opb *OpbChain) buildInterchainRequest(e abci.Event) core.InterchainRequest {
	requestID, err := opb.getAttributeValue(e, "request_id")
	if err != nil {
		logging.Logger.Errorf("failed to get requestID: %s", err)
	}
	endpointInfoStr, err := opb.getAttributeValue(e, "endpoint_info")
	if err != nil {
		logging.Logger.Errorf("failed to get endpointInfo: %s", err)
	}
	method, err := opb.getAttributeValue(e, "method")
	if err != nil {
		logging.Logger.Errorf("failed to get method: %s", err)
	}
	callData, err := opb.getAttributeValue(e, "callData")
	if err != nil {
		logging.Logger.Errorf("failed to get callData: %s", err)
	}
	var endpointInfo EndpointInfo
	err = json.Unmarshal([]byte(endpointInfoStr), &endpointInfo)
	if err != nil {
		logging.Logger.Errorf("failed to decode endpointInfo: %s", err)
	}

	callDataBytes := convCallData(callData)
	return core.InterchainRequest{
		ID:              requestID,
		SourceChainID:   opb.ChainID,
		DestChainID:     endpointInfo.DestChainID,
		DestSubChainID:  endpointInfo.DestSubChainID,
		DestChainType:   endpointInfo.DestChainType,
		EndpointAddress: endpointInfo.EndpointAddress,
		EndpointType:    endpointInfo.EndpointType,
		Method:          method,
		CallData:        callDataBytes,
	}
}

//convCallData 按照hex base64 string的顺序解析字符串
func convCallData(data string) []byte {
	logging.Logger.Infof("Convert CallData : %s", data)
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err == nil {
		return bytes
	}

	return []byte(data)

}

// monitor is responsible for monitoring the chain
func (opb *OpbChain) monitor() {
	for {
		opb.scan()

		if opb.done {
			return
		}

		time.Sleep(time.Duration(opb.Config.MonitorInterval) * time.Second)
	}
}

// scan performs chain scanning
func (opb *OpbChain) scan() {
	currentHeight, err := opb.getBlockNumber()
	if err != nil {
		logging.Logger.Errorf("failed to get the current block height: %s", err)
		return
	}

	if opb.lastHeight == 0 {
		opb.lastHeight = currentHeight - 1
	}

	if currentHeight <= opb.lastHeight {
		return
	}

	opb.scanBlocks(opb.lastHeight+1, currentHeight)
}

// scanBlocks scans the blocks of the specified range
func (opb *OpbChain) scanBlocks(startHeight int64, endHeight int64) {
	for h := startHeight; h <= endHeight; {
		blockResult, err := opb.OpbClient.BlockResults(context.Background(), &h)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			time.Sleep(time.Duration(10) * time.Second)
			continue
		}
		block, err := opb.OpbClient.Block(context.Background(), &h)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			time.Sleep(time.Duration(10) * time.Second)
			continue
		}
		opb.parseCrossChainRequest(blockResult.TxsResults, block)

		err = opb.updateHeight(h)
		if err != nil {
			logging.Logger.Errorf("failed to update height: %s", err)
		}

		h++
	}
}

// getBlockNumber retrieves the current block number
func (opb *OpbChain) getBlockNumber() (int64, error) {
	resultState, err := opb.OpbClient.Status(context.Background())
	if err != nil {
		return -1, err
	}

	return resultState.SyncInfo.LatestBlockHeight, nil
}

// parseServiceInvokedEvents parses the ServiceInvoked events from the receipt
func (opb *OpbChain) parseCrossChainRequest(txResults []*abci.ResponseDeliverTx, block *ctypes.ResultBlock) {
	for i, txResult := range txResults {
		for _, e := range txResult.Events {
			if e.Type == "wasm" && len(e.Attributes) > 1 {
				contractAddr, _ := opb.getAttributeValue(e, "_contract_address")
				if contractAddr == opb.Config.ChainParams.IServiceCoreAddr {
					request := opb.buildInterchainRequest(e)
					err := opb.handler(opb.ChainID, request, strings.ToUpper(hex.EncodeToString(block.Block.Txs[i].Hash())))
					if err != nil {

					}
				}
			}
		}
	}
}

func (opb *OpbChain) getAttributeValue(event abci.Event, attributeKey string) (string, error) {
	for _, attr := range event.Attributes {
		if string(attr.Key) == attributeKey {
			return string(attr.Value), nil
		}
	}

	return "", fmt.Errorf("attribute key %s does not exist", attributeKey)
}

// storeChainParams stores the chain params
func (opb *OpbChain) storeChainParams() error {
	bz, err := json.Marshal(opb.Config.ChainParams)
	if err != nil {
		return err
	}

	return opb.store.Set(ChainParamsKey(opb.ChainID), bz)
}

func (opb *OpbChain) storeChainID() error {
	chainIDsbz, err := opb.store.Get([]byte("chainIDs"))
	if err != nil {
		return err
	}
	chainIDs := map[string]string{}
	err = json.Unmarshal(chainIDsbz, &chainIDs)
	if err != nil {
		return err
	}
	chainIDs[opb.ChainID] = "opb"
	bz, err := json.Marshal(chainIDs)
	return opb.store.Set([]byte("chainIDs"), bz)
}

// updateHeight updates the height
func (opb *OpbChain) updateHeight(height int64) error {
	opb.lastHeight = height
	return opb.store.SetInt64(HeightKey(opb.ChainID), height)
}
