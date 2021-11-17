package test

import (
	"github.com/bianjieai/iritamod-sdk-go/wasm"
	sdk "github.com/irisnet/core-sdk-go/types"
	"relayer/hub"
	"testing"
)

var svcClient hub.IritaHubChain

func init() {
	svcClient = hub.NewIritaHubChain(
		"wenchangchain",
		"http://10.1.4.149:36657",
		"10.1.4.149:39090",
		"node0",
		"12345678",
		"-----BEGIN TENDERMINT PRIVATE KEY-----\nkdf: bcrypt\nsalt: E82064503E284EE753B13E9424B08B4C\ntype: sm2\n\nqLgix+DPFfNY+TpWWlNmquy3jUDR314/dJmIxw8JCWGiSn4deFtp8IWGH/mnVe6S\nNdGt6OJ2SbwO098fk16Gw6RO+MgVjShVMXbkggc=\n=h7AT\n-----END TENDERMINT PRIVATE KEY-----",
		"400uirita",
		"cc-contract-call",
		`{"input":{"type":"object"},"output":{"type:"object"}}`,
		"iaa13fnhnwmjmkdf9wdy3f3ee0umfve8hyarwaerqm",
		"1000000upoint",
		20,
		100,
	)
}

func TestExecuteAppContract(t *testing.T) {
	resultTx, err := svcClient.IritaClient.WASM.Execute(
		"iaa1ghd753shjuwexxywmgs4xz7x2q732vcnednxe6",
		wasm.NewContractABI().WithMethod("hello").WithArgs("words", "it's a test string"),
		sdk.NewCoins(sdk.NewCoin("uirita", sdk.NewInt(1000000))),
		sdk.BaseTx{
			From:     "node0",
			Password: "12345678",
		},
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resultTx)
}

func TestQueryAppContractResult(t *testing.T) {
	result, err := svcClient.IritaClient.WASM.QueryContract("iaa1ghd753shjuwexxywmgs4xz7x2q732vcnednxe6", nil)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(result))
}
