package eth

import (
	"encoding/json"
	"github.com/spf13/viper"
	"math/rand"
	cfg "relayer/config"
)

const (
	Prefix = "eth"

	ChainID         = "chain_id"
	GasLimit        = "gas_limit"
	GasPrice        = "gas_price"
	Key             = "key"
	Passphrase      = "passphrase"
	MonitorInterval = "monitor_interval"
	Nodes           = "nodes"

	IServiceEventName  = "iservice_event_name"
	IServiceEventSig   = "iservice_event_sig"
)

// BaseConfig defines the base config
type BaseConfig struct {
	ChainID         string            `yaml:"chain_id"`
	GasLimit        uint64            `yaml:"gas_limit"`
	GasPrice        uint64            `yaml:"gas_price"`
	Key             string            `yaml:"key"`
	Passphrase      string            `yaml:"passphrase"`
	NodesMap        map[string]string `yaml:"nodes"`
	MonitorInterval uint64
	IServiceEventName  string `yaml:"iservice_event_name"`
	IServiceEventSig   string `yaml:"iservice_event_sig"`
}

func (bc *BaseConfig) PrintConfig() {
}

// Config defines the specific chain config
type Config struct {
	BaseConfig
	ChainParams
}

// NewBaseConfig constructs a new BaseConfig instance from viper
func NewBaseConfig(v *viper.Viper) *BaseConfig {
	return &BaseConfig{
		ChainID:         v.GetString(cfg.GetConfigKey(Prefix, ChainID)),
		GasLimit:        v.GetUint64(cfg.GetConfigKey(Prefix, GasLimit)),
		GasPrice:        v.GetUint64(cfg.GetConfigKey(Prefix, GasPrice)),
		Key:             v.GetString(cfg.GetConfigKey(Prefix, Key)),
		Passphrase:      v.GetString(cfg.GetConfigKey(Prefix, Passphrase)),
		MonitorInterval: v.GetUint64(cfg.GetConfigKey(Prefix, MonitorInterval)),
		NodesMap:        v.GetStringMapString(cfg.GetConfigKey(Prefix, Nodes)),
		IServiceEventName:  v.GetString(cfg.GetConfigKey(Prefix, IServiceEventName)),
		IServiceEventSig:   v.GetString(cfg.GetConfigKey(Prefix, IServiceEventSig)),
	}
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
