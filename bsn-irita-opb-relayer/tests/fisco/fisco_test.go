package fisco_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	fiscoclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"

	"relayer/appchains/opb/iservice"
)

type FISCOTestSuite struct {
	suite.Suite
	client *fiscoclient.Client
}

func TestFISCOTestSuite(t *testing.T) {
	suite.Run(t, new(FISCOTestSuite))
}

func (suite *FISCOTestSuite) SetupTest() {
	configs, err := conf.ParseConfigFile("config.toml")
	suite.NoError(err)

	client, err := fiscoclient.Dial(&configs[0])
	suite.NoError(err)

	suite.client = client
}

func (suite *FISCOTestSuite) TestDeployIServiceContracts() {
	transactOpts := suite.client.GetTransactOpts()
	// deploy iservice core extension contract
	iserviceCoreAddr, tx, _, err := iservice.DeployIServiceCoreEx(transactOpts, suite.client, transactOpts.From, "")
	suite.NoError(err)

	fmt.Printf("iservice core extension deployed, contract address: %s, tx hash: %s", iserviceCoreAddr.String(), tx.Hash().String())

	// deploy iservice proxy contract
	iserviceDelegatorAddr, tx, _, err := iservice.DeployIServiceDelegator(transactOpts, suite.client, iserviceCoreAddr)
	suite.NoError(err)

	fmt.Printf("iservice delegator deployed, contract address: %s, tx hash: %s", iserviceDelegatorAddr.String(), tx.Hash().String())
}
