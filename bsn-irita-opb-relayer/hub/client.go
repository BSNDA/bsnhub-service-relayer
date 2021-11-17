package hub

import (
	"github.com/bianjieai/iritamod-sdk-go/wasm"
	"github.com/irisnet/core-sdk-go/client"
	"github.com/irisnet/core-sdk-go/codec"
	cdctypes "github.com/irisnet/core-sdk-go/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/crypto/codec"
	"github.com/irisnet/core-sdk-go/modules/bank"
	sdk "github.com/irisnet/core-sdk-go/types"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"

	"github.com/bianjieai/iritamod-sdk-go/service"
)

type ServiceClient struct {
	encodingConfig sdk.EncodingConfig
	sdk.BaseClient
	Bank    bank.Client
	Service service.Client
	WASM    wasm.Client
}

func NewServiceClient(config sdk.ClientConfig) *ServiceClient {
	encodingConfig := makeEncodingConfig()
	baseClient := client.NewBaseClient(config, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Codec)
	serviceClient := service.NewClient(baseClient, encodingConfig.Codec)
	wasmClient := wasm.NewClient(baseClient)
	sc := &ServiceClient{
		encodingConfig: encodingConfig,
		BaseClient:     baseClient,
		Bank:           bankClient,
		Service:        serviceClient,
		WASM:           wasmClient,
	}

	sc.RegisterModule(
		bankClient,
		serviceClient,
		wasmClient,
	)

	return sc
}

func (sc *ServiceClient) RegisterModule(ms ...sdk.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(sc.encodingConfig.InterfaceRegistry)
	}
}

//client init
func makeEncodingConfig() sdk.EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := sdk.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		TxConfig:          txCfg,
		Amino:             amino,
		Codec:             marshaler,
	}
	registerLegacyAminoCodec(encodingConfig.Amino)
	registerInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
func registerLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*sdk.Msg)(nil), nil)
	cdc.RegisterInterface((*sdk.Tx)(nil), nil)
	cryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers the sdk message type.
func registerInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*sdk.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
}
