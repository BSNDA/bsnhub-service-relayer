// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iservice

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IServiceCoreExMetaData contains all meta data concerning the IServiceCoreEx contract.
var IServiceCoreExMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_sourceChainID\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_endpointInfo\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_method\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"CrossChainRequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_endpointInfo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_method\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunction\",\"type\":\"bytes4\"}],\"name\":\"sendRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_errMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"setResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001da738038062001da783398181016040528101906200003791906200034e565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600190805190602001906200010a92919062000209565b50600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614620001875781600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001d8565b62000197620001e060201b60201c565b600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b505062000586565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b82805462000217906200047d565b90600052602060002090601f0160209004810192826200023b576000855562000287565b82601f106200025657805160ff191683800117855562000287565b8280016001018555821562000287579182015b828111156200028657825182559160200191906001019062000269565b5b5090506200029691906200029a565b5090565b5b80821115620002b55760008160009055506001016200029b565b5090565b6000620002d0620002ca84620003dd565b620003b4565b905082815260208101848484011115620002ef57620002ee6200054c565b5b620002fc84828562000447565b509392505050565b60008151905062000315816200056c565b92915050565b600082601f83011262000333576200033262000547565b5b815162000345848260208601620002b9565b91505092915050565b6000806040838503121562000368576200036762000556565b5b6000620003788582860162000304565b925050602083015167ffffffffffffffff8111156200039c576200039b62000551565b5b620003aa858286016200031b565b9150509250929050565b6000620003c0620003d3565b9050620003ce8282620004b3565b919050565b6000604051905090565b600067ffffffffffffffff821115620003fb57620003fa62000518565b5b62000406826200055b565b9050602081019050919050565b6000620004208262000427565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60005b83811015620004675780820151818401526020810190506200044a565b8381111562000477576000848401525b50505050565b600060028204905060018216806200049657607f821691505b60208210811415620004ad57620004ac620004e9565b5b50919050565b620004be826200055b565b810181811067ffffffffffffffff82111715620004e057620004df62000518565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b620005778162000413565b81146200058357600080fd5b50565b61181180620005966000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638892bb6a1161005b5780638892bb6a146101155780638da5cb5b146101455780638f32d59b14610163578063f2fde38b1461018157610088565b80631bd752841461008d5780635badbe4c146100bd5780636548e9bc146100db5780638406c079146100f7575b600080fd5b6100a760048036038101906100a29190610c93565b61019d565b6040516100b49190611064565b60405180910390f35b6100c5610336565b6040516100d29190611217565b60405180910390f35b6100f560048036038101906100f09190610bdb565b61033c565b005b6100ff610437565b60405161010c919061102e565b60405180910390f35b61012f600480360381019061012a9190610c08565b61045d565b60405161013c9190611049565b60405180910390f35b61014d61075b565b60405161015a919061102e565b60405180910390f35b61016b610784565b6040516101789190611049565b60405180910390f35b61019b60048036038101906101969190610bdb565b6107db565b005b600085858560008351116101e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101dd906111d7565b60405180910390fd5b600082511161022a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610221906111f7565b60405180910390fd5b600081511161026e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026590611197565b60405180910390fd5b6001600454604051602001610284929190611006565b604051602081830303815290604052805190602001209350600460008154809291906102af9061143f565b91905055507f0faa824f1be7109f16e32fb016edbc264ea83711298f4b3219e1b8fa5aaa8cd7848a8a8a326040516102eb9594939291906110af565b60405180910390a16102fe84878761082e565b60006002600086815260200190815260200160002060006101000a81548160ff02191690831515021790555050505095945050505050565b60045481565b610344610784565b610383576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037a906111b7565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156103f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ea90611137565b60405180910390fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104e690611177565b60405180910390fd5b83600015156002600083815260200190815260200160002060009054906101000a900460ff16151514610557576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054e90611157565b60405180910390fd5b6000600360008781526020019081526020016000206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050905060606000865111156106385785905061063c565b8490505b60016002600089815260200190815260200160002060006101000a81548160ff0219169083151502179055506000826000015173ffffffffffffffffffffffffffffffffffffffff168360200151898460405160240161069d92919061107f565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516107079190610fef565b6000604051808303816000865af19150503d8060008114610744576040519150601f19603f3d011682016040523d82523d6000602084013e610749565b606091505b50509050809450505050509392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6107e3610784565b610822576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610819906111b7565b60405180910390fd5b61082b81610940565b50565b610836610a6d565b82816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508181602001907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050806003600086815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff021916908360e01c021790555090505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156109b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109a790611117565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525090565b6000610acf610aca84611257565b611232565b905082815260208101848484011115610aeb57610aea611524565b5b610af684828561139a565b509392505050565b6000610b11610b0c84611288565b611232565b905082815260208101848484011115610b2d57610b2c611524565b5b610b3884828561139a565b509392505050565b600081359050610b4f81611796565b92915050565b600081359050610b64816117ad565b92915050565b600081359050610b79816117c4565b92915050565b600082601f830112610b9457610b9361151f565b5b8135610ba4848260208601610abc565b91505092915050565b600082601f830112610bc257610bc161151f565b5b8135610bd2848260208601610afe565b91505092915050565b600060208284031215610bf157610bf061152e565b5b6000610bff84828501610b40565b91505092915050565b600080600060608486031215610c2157610c2061152e565b5b6000610c2f86828701610b55565b935050602084013567ffffffffffffffff811115610c5057610c4f611529565b5b610c5c86828701610bad565b925050604084013567ffffffffffffffff811115610c7d57610c7c611529565b5b610c8986828701610bad565b9150509250925092565b600080600080600060a08688031215610caf57610cae61152e565b5b600086013567ffffffffffffffff811115610ccd57610ccc611529565b5b610cd988828901610bad565b955050602086013567ffffffffffffffff811115610cfa57610cf9611529565b5b610d0688828901610bad565b945050604086013567ffffffffffffffff811115610d2757610d26611529565b5b610d3388828901610b7f565b9350506060610d4488828901610b40565b9250506080610d5588828901610b6a565b9150509295509295909350565b610d6b8161131c565b82525050565b610d7a8161132e565b82525050565b610d898161133a565b82525050565b6000610d9a826112ce565b610da481856112e4565b9350610db48185602086016113a9565b610dbd81611533565b840191505092915050565b6000610dd3826112ce565b610ddd81856112f5565b9350610ded8185602086016113a9565b80840191505092915050565b6000610e04826112d9565b610e0e8185611300565b9350610e1e8185602086016113a9565b610e2781611533565b840191505092915050565b60008154610e3f816113dc565b610e498186611311565b94506001821660008114610e645760018114610e7557610ea8565b60ff19831686528186019350610ea8565b610e7e856112b9565b60005b83811015610ea057815481890152600182019150602081019050610e81565b838801955050505b50505092915050565b6000610ebe602683611300565b9150610ec982611544565b604082019050919050565b6000610ee1602f83611300565b9150610eec82611593565b604082019050919050565b6000610f04602483611300565b9150610f0f826115e2565b604082019050919050565b6000610f27602983611300565b9150610f3282611631565b604082019050919050565b6000610f4a602983611300565b9150610f5582611680565b604082019050919050565b6000610f6d602083611300565b9150610f78826116cf565b602082019050919050565b6000610f90602c83611300565b9150610f9b826116f8565b604082019050919050565b6000610fb3602783611300565b9150610fbe82611747565b604082019050919050565b610fd281611390565b82525050565b610fe9610fe482611390565b611488565b82525050565b6000610ffb8284610dc8565b915081905092915050565b60006110128285610e32565b915061101e8284610fd8565b6020820191508190509392505050565b60006020820190506110436000830184610d62565b92915050565b600060208201905061105e6000830184610d71565b92915050565b60006020820190506110796000830184610d80565b92915050565b60006040820190506110946000830185610d80565b81810360208301526110a68184610df9565b90509392505050565b600060a0820190506110c46000830188610d80565b81810360208301526110d68187610df9565b905081810360408301526110ea8186610df9565b905081810360608301526110fe8185610d8f565b905061110d6080830184610d62565b9695505050505050565b6000602082019050818103600083015261113081610eb1565b9050919050565b6000602082019050818103600083015261115081610ed4565b9050919050565b6000602082019050818103600083015261117081610ef7565b9050919050565b6000602082019050818103600083015261119081610f1a565b9050919050565b600060208201905081810360008301526111b081610f3d565b9050919050565b600060208201905081810360008301526111d081610f60565b9050919050565b600060208201905081810360008301526111f081610f83565b9050919050565b6000602082019050818103600083015261121081610fa6565b9050919050565b600060208201905061122c6000830184610fc9565b92915050565b600061123c61124d565b9050611248828261140e565b919050565b6000604051905090565b600067ffffffffffffffff821115611272576112716114f0565b5b61127b82611533565b9050602081019050919050565b600067ffffffffffffffff8211156112a3576112a26114f0565b5b6112ac82611533565b9050602081019050919050565b60008190508160005260206000209050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b600061132782611370565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156113c75780820151818401526020810190506113ac565b838111156113d6576000848401525b50505050565b600060028204905060018216806113f457607f821691505b60208210811415611408576114076114c1565b5b50919050565b61141782611533565b810181811067ffffffffffffffff82111715611436576114356114f0565b5b80604052505050565b600061144a82611390565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561147d5761147c611492565b5b600182019050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f6953657276696365436f726545783a2072656c6179657220616464726573732060008201527f63616e206e6f74206265207a65726f0000000000000000000000000000000000602082015250565b7f6953657276696365436f726545783a206475706c69636174656420726573706f60008201527f6e73652100000000000000000000000000000000000000000000000000000000602082015250565b7f6953657276696365436f726545783a2073656e646572206973206e6f7420746860008201527f652072656c617965720000000000000000000000000000000000000000000000602082015250565b7f6953657276696365436f726545783a2063616c6c446174612063616e206e6f7460008201527f20626520656d7074790000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f6953657276696365436f726545783a2064657374436861696e49442063616e2060008201527f6e6f7420626520656d7074790000000000000000000000000000000000000000602082015250565b7f6953657276696365436f726545783a206d6574686f642063616e206e6f74206260008201527f6520656d70747900000000000000000000000000000000000000000000000000602082015250565b61179f8161131c565b81146117aa57600080fd5b50565b6117b68161133a565b81146117c157600080fd5b50565b6117cd81611344565b81146117d857600080fd5b5056fea26469706673582212200329b153300e253f1eca1a82a7261dd819feedb2ae5376e269446c5b50a0eda264736f6c63430008070033",
}

// IServiceCoreExABI is the input ABI used to generate the binding from.
// Deprecated: Use IServiceCoreExMetaData.ABI instead.
var IServiceCoreExABI = IServiceCoreExMetaData.ABI

// IServiceCoreExBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use IServiceCoreExMetaData.Bin instead.
var IServiceCoreExBin = IServiceCoreExMetaData.Bin

// DeployIServiceCoreEx deploys a new Ethereum contract, binding an instance of IServiceCoreEx to it.
func DeployIServiceCoreEx(auth *bind.TransactOpts, backend bind.ContractBackend, _relayer common.Address, _sourceChainID string) (common.Address, *types.Transaction, *IServiceCoreEx, error) {
	parsed, err := IServiceCoreExMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(IServiceCoreExBin), backend, _relayer, _sourceChainID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IServiceCoreEx{IServiceCoreExCaller: IServiceCoreExCaller{contract: contract}, IServiceCoreExTransactor: IServiceCoreExTransactor{contract: contract}, IServiceCoreExFilterer: IServiceCoreExFilterer{contract: contract}}, nil
}

// IServiceCoreEx is an auto generated Go binding around an Ethereum contract.
type IServiceCoreEx struct {
	IServiceCoreExCaller     // Read-only binding to the contract
	IServiceCoreExTransactor // Write-only binding to the contract
	IServiceCoreExFilterer   // Log filterer for contract events
}

// IServiceCoreExCaller is an auto generated read-only Go binding around an Ethereum contract.
type IServiceCoreExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IServiceCoreExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IServiceCoreExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IServiceCoreExSession struct {
	Contract     *IServiceCoreEx   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IServiceCoreExCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IServiceCoreExCallerSession struct {
	Contract *IServiceCoreExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IServiceCoreExTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IServiceCoreExTransactorSession struct {
	Contract     *IServiceCoreExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IServiceCoreExRaw is an auto generated low-level Go binding around an Ethereum contract.
type IServiceCoreExRaw struct {
	Contract *IServiceCoreEx // Generic contract binding to access the raw methods on
}

// IServiceCoreExCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IServiceCoreExCallerRaw struct {
	Contract *IServiceCoreExCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceCoreExTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IServiceCoreExTransactorRaw struct {
	Contract *IServiceCoreExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIServiceCoreEx creates a new instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreEx(address common.Address, backend bind.ContractBackend) (*IServiceCoreEx, error) {
	contract, err := bindIServiceCoreEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreEx{IServiceCoreExCaller: IServiceCoreExCaller{contract: contract}, IServiceCoreExTransactor: IServiceCoreExTransactor{contract: contract}, IServiceCoreExFilterer: IServiceCoreExFilterer{contract: contract}}, nil
}

// NewIServiceCoreExCaller creates a new read-only instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExCaller(address common.Address, caller bind.ContractCaller) (*IServiceCoreExCaller, error) {
	contract, err := bindIServiceCoreEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCaller{contract: contract}, nil
}

// NewIServiceCoreExTransactor creates a new write-only instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExTransactor(address common.Address, transactor bind.ContractTransactor) (*IServiceCoreExTransactor, error) {
	contract, err := bindIServiceCoreEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExTransactor{contract: contract}, nil
}

// NewIServiceCoreExFilterer creates a new log filterer instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExFilterer(address common.Address, filterer bind.ContractFilterer) (*IServiceCoreExFilterer, error) {
	contract, err := bindIServiceCoreEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExFilterer{contract: contract}, nil
}

// bindIServiceCoreEx binds a generic wrapper to an already deployed contract.
func bindIServiceCoreEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceCoreEx *IServiceCoreExRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IServiceCoreEx.Contract.IServiceCoreExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceCoreEx *IServiceCoreExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.IServiceCoreExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.IServiceCoreExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceCoreEx *IServiceCoreExCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IServiceCoreEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IServiceCoreEx *IServiceCoreExCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IServiceCoreEx.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IServiceCoreEx *IServiceCoreExCallerSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IServiceCoreEx.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IServiceCoreEx.contract.Call(opts, &out, "relayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() view returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCaller) RequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IServiceCoreEx.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_IServiceCoreEx *IServiceCoreExSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCallerSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactor) SendRequest(opts *bind.TransactOpts, _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "sendRequest", _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExSession) SendRequest(_endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SendRequest(&_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SendRequest(_endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SendRequest(&_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) SetRelayer(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setRelayer", _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactor) SetResponse(opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setResponse", _requestID, _errMsg, _output)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) SetResponse(_requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetResponse(_requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

// IServiceCoreExCrossChainRequestSentIterator is returned from FilterCrossChainRequestSent and is used to iterate over the raw logs and unpacked data for CrossChainRequestSent events raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainRequestSentIterator struct {
	Event *IServiceCoreExCrossChainRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IServiceCoreExCrossChainRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExCrossChainRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IServiceCoreExCrossChainRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IServiceCoreExCrossChainRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExCrossChainRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExCrossChainRequestSent represents a CrossChainRequestSent event raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainRequestSent struct {
	RequestID    [32]byte
	EndpointInfo string
	Method       string
	CallData     []byte
	Sender       common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCrossChainRequestSent is a free log retrieval operation binding the contract event 0x0faa824f1be7109f16e32fb016edbc264ea83711298f4b3219e1b8fa5aaa8cd7.
//
// Solidity: event CrossChainRequestSent(bytes32 _requestID, string _endpointInfo, string _method, bytes _callData, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterCrossChainRequestSent(opts *bind.FilterOpts) (*IServiceCoreExCrossChainRequestSentIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "CrossChainRequestSent")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCrossChainRequestSentIterator{contract: _IServiceCoreEx.contract, event: "CrossChainRequestSent", logs: logs, sub: sub}, nil
}

// WatchCrossChainRequestSent is a free log subscription operation binding the contract event 0x0faa824f1be7109f16e32fb016edbc264ea83711298f4b3219e1b8fa5aaa8cd7.
//
// Solidity: event CrossChainRequestSent(bytes32 _requestID, string _endpointInfo, string _method, bytes _callData, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchCrossChainRequestSent(opts *bind.WatchOpts, sink chan<- *IServiceCoreExCrossChainRequestSent) (event.Subscription, error) {

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "CrossChainRequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExCrossChainRequestSent)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainRequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainRequestSent is a log parse operation binding the contract event 0x0faa824f1be7109f16e32fb016edbc264ea83711298f4b3219e1b8fa5aaa8cd7.
//
// Solidity: event CrossChainRequestSent(bytes32 _requestID, string _endpointInfo, string _method, bytes _callData, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseCrossChainRequestSent(log types.Log) (*IServiceCoreExCrossChainRequestSent, error) {
	event := new(IServiceCoreExCrossChainRequestSent)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainRequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IServiceCoreExOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IServiceCoreEx contract.
type IServiceCoreExOwnershipTransferredIterator struct {
	Event *IServiceCoreExOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IServiceCoreExOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IServiceCoreExOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IServiceCoreExOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExOwnershipTransferred represents a OwnershipTransferred event raised by the IServiceCoreEx contract.
type IServiceCoreExOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IServiceCoreExOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExOwnershipTransferredIterator{contract: _IServiceCoreEx.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IServiceCoreExOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExOwnershipTransferred)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseOwnershipTransferred(log types.Log) (*IServiceCoreExOwnershipTransferred, error) {
	event := new(IServiceCoreExOwnershipTransferred)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
