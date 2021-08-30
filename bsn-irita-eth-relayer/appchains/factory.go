package appchains

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/viper"

	"relayer/appchains/eth"
	"relayer/config"
	// "relayer/appchains/ethereum"
	"relayer/core"
	"relayer/store"
)

// AppChainFactory defines an application chain factory
type AppChainFactory struct {
	Store *store.Store // store
}

// AppChainFactory defines an application chain factory
type BaseConfigFactory struct {
	config *viper.Viper// store
}

// NewAppChainFactory constructs a new application chain factory
func NewAppChainFactory(store *store.Store) *AppChainFactory {
	return &AppChainFactory{
		Store: store,
	}
}

// NewAppChainFactory constructs a new application chain factory
func NewBaseConfigFactory(v *viper.Viper) *BaseConfigFactory {
	return &BaseConfigFactory{
		config: v,
	}
}

// BuildAppChain implements AppChainFactoryI
func (f *AppChainFactory) BuildAppChain(chainType string, chainParams []byte) (core.AppChainI, error) {
	switch strings.ToLower(chainType) {
	case "eth":
		return eth.BuildEthChain(chainParams, f.Store)

	case "fabric":
		return nil, nil

	case "fisco":
		return nil, nil

	default:
		return nil, fmt.Errorf("application chain %s not supported", chainType)
	}
}

// GetChainID implements AppChainFactoryI
func (f *AppChainFactory) GetChainID(chainType string, chainParams []byte) (chainID string, err error) {
	switch strings.ToLower(chainType) {
	case "eth":
		return eth.GetChainIDFromBytes(chainParams)

	case "fabric":
		return "", nil

	case "fisco":
		return "", nil

	default:
		return "", fmt.Errorf("application chain %s not supported", chainType)
	}
}

// StoreBaseConfig implements AppChainFactoryI
func (f *AppChainFactory) StoreBaseConfig(chainType string, baseConfig []byte) error {
	switch strings.ToLower(chainType) {
	case "eth":
		return eth.StoreBaseConfig(f.Store, baseConfig)

	case "fabric":
		return nil

	case "fisco":
		return nil

	default:
		return fmt.Errorf("application chain %s not supported", chainType)
	}
}

// StoreBaseConfig implements AppChainFactoryI
func (f *AppChainFactory) DeleteChainConfig(chainType string, chainID string) error {
	switch strings.ToLower(chainType) {
	case "eth":
		err := f.Store.Delete(eth.ChainParamsKey(chainID))
		if err != nil{
			return err
		}
		chainIDsbz, _ := f.Store.Get([]byte("chainIDs"))
		chainIDs:= map[string]string{}
		json.Unmarshal(chainIDsbz, &chainIDs)
		delete(chainIDs, chainID)
		bz, err := json.Marshal(chainIDs)
		err = f.Store.Set([]byte("chainIDs"), bz)
		if err != nil{
			return err
		}
		return f.Store.Delete(eth.ChainParamsKey(chainID))
		return nil

	case "fabric":
		return nil

	case "fisco":
		return nil
	default:
		return fmt.Errorf("application chain %s not supported", chainType)
	}
}

// StoreBaseConfig implements AppChainFactoryI
func (bc *BaseConfigFactory) NewBaseConfig(chainType string) (config.BaseConfigI, error) {
	switch strings.ToLower(chainType) {
	case "eth":
		return eth.NewBaseConfig(bc.config), nil

	case "fabric":
		return nil, nil

	case "fisco":
		return nil, nil

	default:
		return nil, fmt.Errorf("application chain %s not supported", chainType)
	}
}
