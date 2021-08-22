package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	_ "net/http/pprof"
	"relayer/appchains"
	cfg "relayer/config"
	"relayer/core"
	"relayer/hub"
	"relayer/logging"
	"relayer/mysql"
	"relayer/server"
	"relayer/store"

	txstore "relayer/appchains/eth/store"
)

const (
	_HttpPort = "base.http_port"
)

// StartCmd implements the start command
func StartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start [config-file]",
		Short: "Start the relayer demon",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 0 {
				configFileName = cfg.DefaultConfigFileName
			} else {
				configFileName = args[0]
			}

			config, err := cfg.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			appChainType := config.GetString(cfg.ConfigKeyAppChainType)

			store, err := store.NewStore(config.GetString(cfg.ConfigKeyStorePath))
			if err != nil {
				return err
			}

			mysqlConfig := mysql.NewConfig(config)

			txstore.InitMysql(mysqlConfig.DSN())
			//mysql.NewDB(mysqlConfig)
			//defer mysql.Close()

			appChainFactory := appchains.NewAppChainFactory(store)
			hubChain := hub.BuildIritaHubChain(hub.NewConfig(config))
			relayerInstance := core.NewRelayer(appChainType, hubChain, appChainFactory, logging.Logger)

			baseConfigFactory := appchains.NewBaseConfigFactory(config)
			BaseConfig, err := baseConfigFactory.NewBaseConfig(appChainType)
			if err != nil {
				return err
			}
			baseConfigByte, _ := json.Marshal(BaseConfig)
			appChainFactory.StoreBaseConfig(appChainType, baseConfigByte)
			chainIDsbz, _ := store.Get([]byte("chainIDs"))
			if chainIDsbz == nil {
				chainIDsbz, err = json.Marshal(map[string]string{})
				if err != nil {
					return err
				}
				store.Set([]byte("chainIDs"), chainIDsbz)
			} else {
				chainIDs := map[string]string{}
				json.Unmarshal(chainIDsbz, &chainIDs)
				for chainID, chainType := range chainIDs {
					if chainType == appChainType {
						chainParams, err := store.Get([]byte(fmt.Sprintf("%s:params:%s", appChainType, chainID)))
						if err != nil {
							return err
						}
						chain, err := relayerInstance.AppChainFactory.BuildAppChain(chainType, chainParams)
						if err != nil {
							return err
						}

						if err := chain.Start(relayerInstance.HandleInterchainRequest); err != nil {
							return err
						}

						relayerInstance.AppChains[chainID] = chain
						relayerInstance.AppChainStates[chainID] = true
					}
				}
			}

			chainManager := server.NewChainManager(relayerInstance)

			httpPort := config.GetInt(_HttpPort)
			if httpPort == 0 {
				httpPort = 8082
			}

			http.ListenAndServe("0.0.0.0:6060", nil)
			server.StartWebServer(chainManager, httpPort)

			return nil
		},
	}

	return cmd
}
