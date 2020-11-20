package bcos

import (
	"github.com/spf13/viper"

	cfg "relayer/config"
)

const (
	Prefix = "bcos"

	ConfigFile = "config_file"

	IServiceCoreAddr   = "iservice_core_addr"
	IServiceMarketAddr = "iservice_market_addr"

	MonitorInterval = "monitor_interval"
)

// Config represents the BCOS chain config
type Config struct {
	ConfigFile         string `yaml:"config_file"`
	IServiceCoreAddr   string `yaml:"iservice_core_addr"`
	IServiceMarketAddr string `yaml:"iservice_market_addr"`
	MonitorInterval    uint64 `yaml:"monitor_interval"`
}

// NewConfig constructs a new Config from viper
func NewConfig(v *viper.Viper) Config {
	return Config{
		ConfigFile:         v.GetString(cfg.GetConfigKey(Prefix, ConfigFile)),
		IServiceCoreAddr:   v.GetString(cfg.GetConfigKey(Prefix, IServiceCoreAddr)),
		IServiceMarketAddr: v.GetString(cfg.GetConfigKey(Prefix, IServiceMarketAddr)),
		MonitorInterval:    v.GetUint64(cfg.GetConfigKey(Prefix, MonitorInterval)),
	}
}
