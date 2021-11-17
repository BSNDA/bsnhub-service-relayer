package opb

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"math/rand"

	cfg "relayer/config"
)

const (
	Prefix = "opb"

	// base config
	ChainId         = "chain_id"
	AccountKey      = "account"
	RpcAddrsMap     = "rpc_addrs"
	GrpcAddrsMap    = "grpc_addrs"
	MonitorInterval = "monitor_interval"
	DefaultFee      = "default_fee"
	DefaultGas      = "default_gas"
	Timeout         = "timeout"
)

const (
	defaultAlgo = "sm2"
)

// BaseConfig defines the base config
type BaseConfig struct {
	Account         Account `yaml:"account"`
	RpcAddrsMap     map[string]string
	GrpcAddrsMap    map[string]string
	ChainId         string
	DefaultFee      string
	Timeout         uint
	DefaultGas      uint64
	MonitorInterval uint64
}

type Account struct {
	KeyName    string `yaml:"key_name" mapstructure:"key_name"`
	Passphrase string `yaml:"passphrase"`
	KeyArmor   string `yaml:"key_armor" mapstructure:"key_armor"`
}

func (bc *BaseConfig) PrintConfig() {
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) (*BaseConfig, error) {
	account := Account{}
	err := v.UnmarshalKey(cfg.GetConfigKey(Prefix, AccountKey), &account)
	if err != nil {
		fmt.Println(err)
	}
	config := new(BaseConfig)
	config.ChainId = v.GetString(cfg.GetConfigKey(Prefix, ChainId))
	config.RpcAddrsMap = v.GetStringMapString(cfg.GetConfigKey(Prefix, RpcAddrsMap))
	config.GrpcAddrsMap = v.GetStringMapString(cfg.GetConfigKey(Prefix, GrpcAddrsMap))
	config.Timeout = v.GetUint(cfg.GetConfigKey(Prefix, Timeout))
	config.DefaultFee = v.GetString(cfg.GetConfigKey(Prefix, DefaultFee))
	config.DefaultGas = v.GetUint64(cfg.GetConfigKey(Prefix, DefaultGas))
	config.MonitorInterval = v.GetUint64(cfg.GetConfigKey(Prefix, MonitorInterval))
	config.Account = account

	return config, nil
}
func randURL(m []string) string {
	if len(m) == 0 {
		return ""
	}
	for _, index := range rand.Perm(len(m)) {
		return m[index]
	}
	return ""
}

// ValidBaseConfig validates if the given bytes is valid BaseConfig
func ValidateBaseConfig(baseCfg []byte) error {
	var baseConfig BaseConfig
	return json.Unmarshal(baseCfg, &baseConfig)
}
