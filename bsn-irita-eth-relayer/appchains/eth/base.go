package eth

import (
	"encoding/json"
)

const (
	ChainType              = "eth"
	DefaultMonitorInterval = 1 // 1 second by default
)

// CompactBlock represents the compact block with tx hashes
type CompactBlock struct {
	Txs []string `json:"transactions"`
}

// ChainParams defines the params for the specific chain
type ChainParams struct {
	NodeURLs         []string `json:"nodes"`
	ChainID          string    `json:"chainId"`
	IServiceCoreAddr string   `json:"iserviceCoreAddr"`
}

type EndpointInfo struct {
	DestSubChainID  string `json:"dest_sub_chain_id"`
	DestChainID     string `json:"dest_chain_id"`
	DestChainType   string `json:"dest_chain_type"`
	EndpointAddress string `json:"endpoint_address"`
	EndpointType    string `json:"endpoint_type"`
}

// GetChainID returns the unique chain id from the specified chain params
func GetChainID(params ChainParams) string {
	return  params.ChainID
}

// GetChainIDFromBytes returns the unique chain id from the given chain params bytes
func GetChainIDFromBytes(params []byte) (string, error) {
	var chainParams ChainParams
	err := json.Unmarshal(params, &chainParams)
	if err != nil {
		return "", err
	}

	return GetChainID(chainParams), nil
}
