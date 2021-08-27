package eth

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"relayer/appchains/eth/iservice"
	txstore "relayer/appchains/eth/store"
	"relayer/core"
	"relayer/logging"
	"relayer/store"
)

// EthChain defines the Eth chain
type EthChain struct {
	Config  Config
	Client  *ethclient.Client
	ChainID string // unique chain ID

	IServiceCoreContract *iservice.IServiceCoreEx // iService Core Extension contract
	IServiceCoreABI      abi.ABI                  // parsed iService Core Extension ABI

	store              *store.Store // store backend instance
	done    bool
	ClientSubscription ethereum.Subscription
}

// NewEthChain constructs a new EthChain instance
func NewEthChain(
	config Config,
	store *store.Store,
) (*EthChain, error) {

	nodeName := randURL(config.NodeURLs)
	//获取配置的nodeURL
	nodeUrl, ok := config.NodesMap[nodeName]
	if ok {
		nodeName = nodeUrl
	}

	client, err := ethclient.Dial(nodeUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to eth node: %s", err)
	}

	iServiceCoreABI, err := abi.JSON(strings.NewReader(iservice.IServiceCoreExABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse iService Core Extension ABI: %s", err)
	}

	iServiceCore, err := iservice.NewIServiceCoreEx(ethcmn.HexToAddress(config.IServiceCoreAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the iService Core Extension contract: %s", err)
	}

	chainID := GetChainID(config.ChainParams)

	eth := &EthChain{
		Config:               config,
		Client:               client,
		ChainID:              chainID,
		IServiceCoreContract: iServiceCore,
		IServiceCoreABI:      iServiceCoreABI,
		store:                store,
	}

	err = eth.storeChainParams()
	if err != nil {
		return nil, err
	}

	err = eth.storeChainID()
	if err != nil {
		return nil, err
	}

	return eth, nil
}

// BuildEthChain builds a EthChain instance from the given chain params and store
func BuildEthChain(
	chainParams []byte,
	store *store.Store,
) (*EthChain, error) {
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

	return NewEthChain(config, store)
}

// GetChainID implements AppChainI
func (ec *EthChain) GetChainID() string {
	return ec.ChainID
}

// Start implements AppChainI
func (ec *EthChain) Start(handler core.InterchainRequestHandler) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filterQuery := ethereum.FilterQuery{
		Addresses: []ethcmn.Address{ethcmn.HexToAddress(ec.Config.IServiceCoreAddr)},
		Topics:    [][]ethcmn.Hash{{crypto.Keccak256Hash([]byte(ec.Config.IServiceEventSig))}},
	}

	ch := make(chan ethtypes.Log)

	sub, err := ec.Client.SubscribeFilterLogs(ctx, filterQuery, ch)
	if err != nil {
		return err
	}

	ec.ClientSubscription = sub
	ec.done = false

	logHandler := func(log ethtypes.Log) {
		iServiceRequestEvent, err := ec.parseLog(log)
		if err != nil {
			logging.Logger.Errorf("failed to parse log %+v: %s", log, err)
		} else {
			request := ec.buildInterchainRequest(&iServiceRequestEvent)
			handler(ec.ChainID, request, log.TxHash.String())
		}
	}

	go ec.logListener(sub, ch, logHandler)

	return nil
}

// Stop implements AppChainI
func (ec *EthChain) Stop() error {
	logging.Logger.Infof("stopping chain %s", ec.ChainID)
	if ec.ClientSubscription != nil{
		ec.ClientSubscription.Unsubscribe()
	}
	ec.done = true

	return nil
}

func (ec *EthChain) Close() {
	ec.Client.Close()
}

// GetHeight implements AppChainI
func (ec *EthChain) GetHeight() int64 {
	return int64(1)
}

// SendResponse implements AppChainI
func (ec *EthChain) SendResponse(requestID string, response core.ResponseI) error {
	auth, err := ec.buildAuthTransactor()
	if err != nil {
		return err
	}
	requestIDBytes, err := hex.DecodeString(requestID)
	if err != nil {
		return err
	}

	data := &txstore.RelayerResInfo{
		RequestId: requestID,
		TxStatus:  txstore.TxStatus_Success,
		ErrMsg:    "",
	}

	defer func(d *txstore.RelayerResInfo) {
		txstore.RelayerResponeRecord(d)
	}(data)

	var requestID32Bytes [32]byte
	copy(requestID32Bytes[:], requestIDBytes)

	tx, err := ec.IServiceCoreContract.SetResponse(auth, requestID32Bytes, response.GetErrMsg(), response.GetOutput())
	if err != nil {
		data.TxStatus = txstore.TxStatus_Error
		data.ErrMsg = fmt.Sprintf("call eth setResponse failed :%s", err)

		return err
	}

	data.FromResTxId = tx.Hash().Hex()

	// TODO
	//mysql.OnInterchainRequestResponseSent(requestID, tx.Hash().Hex())

	err = ec.waitForReceipt(tx, "SetResponse")
	if err != nil {

		data.TxStatus = txstore.TxStatus_Error
		data.ErrMsg = fmt.Sprintf("call eth setResponse failed :%s", err)
		return err
	}

	// TODO
	//mysql.OnInterchainRequestSucceeded(requestID)

	return nil
}

// buildInterchainRequest builds an interchain request from the interchain event
func (ec *EthChain) buildInterchainRequest(e *iservice.IServiceCoreExCrossChainRequestSent) core.InterchainRequest {
	var endpointInfo EndpointInfo
	err := json.Unmarshal([]byte(e.EndpointInfo), &endpointInfo)
	if err != nil {
		logging.Logger.Errorf("failed to decode endpointInfo: %s", err)
	}
	return core.InterchainRequest{
		ID:              hex.EncodeToString(e.RequestID[:]),
		SourceChainID:   ec.ChainID,
		DestChainID:     endpointInfo.DestChainID,
		DestSubChainID:  endpointInfo.DestSubChainID,
		DestChainType:   endpointInfo.DestChainType,
		EndpointAddress: endpointInfo.EndpointAddress,
		EndpointType:    endpointInfo.EndpointType,
		Method:          e.Method,
		CallData:        e.CallData,
		Sender:          e.Sender.String(),
	}
}

// waitForReceipt waits for the receipt of the given tx
func (ec *EthChain) waitForReceipt(tx *ethtypes.Transaction, name string) error {
	logging.Logger.Infof("%s: transaction sent to %s, hash: %s", name, ec.GetChainID(), tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		return fmt.Errorf("failed to mint the transaction %s: %s", tx.Hash().Hex(), err)
	}

	if receipt.Status == ethtypes.ReceiptStatusFailed {
		return fmt.Errorf("%s: transaction %s execution failed", name, tx.Hash().Hex())
	}

	logging.Logger.Infof("%s: transaction %s execution succeeded", name, tx.Hash().Hex())

	return nil
}

// getBlock gets the block in the given height
func (ec *EthChain) getBlock(height int64) (block *ethtypes.Block, err error) {
	return ec.Client.BlockByNumber(context.Background(), big.NewInt(height))
}

// logListener listens to the log sent by the given channel and handles it with the specified handler
func (ec EthChain) logListener(sub ethereum.Subscription, logChan chan ethtypes.Log, handler func(log ethtypes.Log)) {
	for {
		select {
		case log := <-logChan:
			handler(log)
		case err := <-sub.Err():
			logging.Logger.Errorf("Error on log subscription: %s", err)
		}
		if ec.done {
			return
		}
		time.Sleep(time.Duration(ec.Config.MonitorInterval) * time.Second)
	}
}

// parseServiceInvokedEvents parses the ServiceInvoked events from the receipt
func (ec *EthChain) parseLog(log ethtypes.Log) (iservice.IServiceCoreExCrossChainRequestSent, error) {
	var event iservice.IServiceCoreExCrossChainRequestSent
	err := ec.IServiceCoreABI.Unpack(&event, ec.Config.IServiceEventName, log.Data)
	if err != nil {
		return event, err
	}
	return event, nil
}

// storeChainParams stores the chain params
func (ec *EthChain) storeChainParams() error {
	bz, err := json.Marshal(ec.Config.ChainParams)
	if err != nil {
		return err
	}

	return ec.store.Set(ChainParamsKey(ec.ChainID), bz)
}

func (ec *EthChain) storeChainID() error {
	chainIDsbz, err := ec.store.Get([]byte("chainIDs"))
	if err != nil {
		return err
	}
	chainIDs := map[string]string{}
	err = json.Unmarshal(chainIDsbz, &chainIDs)
	if err != nil {
		return err
	}
	chainIDs[ec.ChainID] = "eth"
	bz, _ := json.Marshal(chainIDs)
	return ec.store.Set([]byte("chainIDs"), bz)
}

// buildAuthTransactor builds an authenticated transactor
func (ec *EthChain) buildAuthTransactor() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(ec.Config.Key)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)

	nextNonce, err := ec.Client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}

	auth.GasLimit = ec.Config.GasLimit
	auth.GasPrice = big.NewInt(int64(ec.Config.GasPrice))
	auth.Nonce = big.NewInt(int64(nextNonce))

	return auth, nil
}
