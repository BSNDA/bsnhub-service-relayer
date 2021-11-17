package hub

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/irisnet/core-sdk-go/types"

	cfg "relayer/config"
)

// default config variables
var (
	defaultChainID       = "irita-hub"
	defaultNodeRPCAddr   = "http://127.0.0.1:26657"
	defaultNodeGRPCAddr  = "127.0.0.1:9090"
	defaultGas           = uint64(200000)
	defaultFee           = "4point"
	defaultBroadcastMode = types.Commit
	defaultKeyAlgorithm  = "sm2"
	defaultServiceName   = "cc-contract-call"
	defaultSchemas       = ""
	defaultProvider      = "iaa1fe6gm5kyam6xfs0wngw3d23l9djlyw82xxcjm2"
	defaultServiceFee    = "1000000upoint"
	defaultTimeout       = uint(20)
	defaultQoS           = uint64(100)
)

const (
	Prefix        = "hub"
	ServicePrefix = "service"
	ChainID       = "chain_id"
	NodeRPCAddr   = "node_rpc_addr"
	NodeGRPCAddr  = "node_grpc_addr"
	AccountKey    = "account"
	Fee           = "fee"
	ServiceName   = "service_name"
	Schemas       = "schemas"
	Provider      = "provider"
	ServiceFee    = "service_fee"
	Timeout       = "timeout"
	QoS           = "qos"
)

// Config is a config struct for IRITA-HUB
type Config struct {
	// chain cfg -> "hub.*"
	ChainID      string  `yaml:"chain_id"`
	NodeRPCAddr  string  `yaml:"node_rpc_addr"`
	NodeGRPCAddr string  `yaml:"node_grpc_addr"`
	Account      Account `yaml:"account"`
	Fee          string  `yaml:"fee"`

	// service cfg -> "service.*"
	ServiceName string `yaml:"service_name"` // service name
	Schemas     string `yaml:"schemas"`      // input and output schemas
	Provider    string `yaml:"provider"`     // service provider
	ServiceFee  string `yaml:"service_fee"`  // service fee
	Timeout     uint   `yaml:"timeout"`
	QoS         uint64 `yaml:"qos"` // quality of service, in terms of the minimum response time
}

type Account struct {
	KeyName    string `yaml:"key_name" mapstructure:"key_name"`
	Passphrase string `yaml:"passphrase"`
	KeyArmor   string `yaml:"key_armor" mapstructure:"key_armor"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	account := Account{}
	err := v.UnmarshalKey(cfg.GetConfigKey(Prefix, AccountKey), &account)
	if err != nil {
		fmt.Println(err)
	}
	return Config{
		ChainID:      v.GetString(cfg.GetConfigKey(Prefix, ChainID)),
		NodeRPCAddr:  v.GetString(cfg.GetConfigKey(Prefix, NodeRPCAddr)),
		NodeGRPCAddr: v.GetString(cfg.GetConfigKey(Prefix, NodeGRPCAddr)),
		Account:      account,
		Fee:          v.GetString(cfg.GetConfigKey(Prefix, Fee)),
		ServiceName:  v.GetString(cfg.GetConfigKey(ServicePrefix, ServiceName)),
		Schemas:      v.GetString(cfg.GetConfigKey(ServicePrefix, Schemas)),
		Provider:     v.GetString(cfg.GetConfigKey(ServicePrefix, Provider)),
		ServiceFee:   v.GetString(cfg.GetConfigKey(ServicePrefix, ServiceFee)),
		Timeout:      v.GetUint(cfg.GetConfigKey(ServicePrefix, Timeout)),
		QoS:          v.GetUint64(cfg.GetConfigKey(ServicePrefix, QoS)),
	}
}
