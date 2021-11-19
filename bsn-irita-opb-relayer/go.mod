module relayer

go 1.16

require (
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/bianjieai/iritamod-sdk-go v0.0.0-20211119065750-d3edd49ebbe1
	github.com/cockroachdb/pebble v0.0.0-20201118202804-75ede898b66c
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.0
	github.com/irisnet/core-sdk-go v0.0.0-20211118114422-2efa1178f1e2
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v1.1.1
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/tendermint/tendermint v0.34.14
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
	github.com/ugorji/go => github.com/ugorji/go v1.1.2
)
