package appchains

import (
	"fmt"
	"strings"

	"relayer/appchains/fisco"
	"relayer/core"
	"relayer/store"
)

// AppChainFactory defines an application chain factory
type AppChainFactory struct {
	Store *store.Store // store
}

// NewAppChainFactory constructs a new application chain factory
func NewAppChainFactory(store *store.Store) *AppChainFactory {
	return &AppChainFactory{
		Store: store,
	}
}

// BuildAppChain implements AppChainFactoryI
func (f *AppChainFactory) BuildAppChain(chainType string, chainParams []byte) (core.AppChainI, error) {
	switch strings.ToLower(chainType) {
	case "eth":
		// return ethereum.MakeEthChain(ethereum.NewConfig(f.Config)), nil
		return nil, nil

	case "fabric":
		// return fabric.MakeFabricChain(fabric.NewConfig(f.Config)), nil
		return nil, nil

	case "fisco":
		return fisco.BuildFISCOChain(chainParams, f.Store)

	default:
		return nil, fmt.Errorf("application chain %s not supported", chainType)
	}
}

// GetChainID implements AppChainFactoryI
func (f *AppChainFactory) GetChainID(chainType string, chainParams []byte) (chainID string, err error) {
	switch strings.ToLower(chainType) {
	case "eth":
		return "", nil

	case "fabric":
		return "", nil

	case "fisco":
		return fisco.GetChainIDFromBytes(chainParams)

	default:
		return "", fmt.Errorf("application chain %s not supported", chainType)
	}
}

// StoreBaseConfig implements AppChainFactoryI
func (f *AppChainFactory) StoreBaseConfig(chainType string, baseConfig []byte) error {
	switch strings.ToLower(chainType) {
	case "eth":
		return nil

	case "fabric":
		return nil

	case "fisco":
		return fisco.StoreBaseConfig(f.Store, baseConfig)

	default:
		return fmt.Errorf("application chain %s not supported", chainType)
	}
}
