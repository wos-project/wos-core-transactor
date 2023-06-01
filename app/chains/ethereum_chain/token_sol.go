// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereumChain

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractTokenABI is the input ABI used to generate the binding from.
const ContractTokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"Greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Hello\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractTokenBin is the compiled bytecode used for deploying new contracts.
var ContractTokenBin = "0x608060405234801561001057600080fd5b506101d7806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063bcdfe0d51461003b578063efdeaaf514610074575b600080fd5b60408051808201909152600b81526a12195b1b1bc815dbdc9b1960aa1b60208201525b60405161006b9190610085565b60405180910390f35b61005e6100823660046100f0565b90565b600060208083528351808285015260005b818110156100b257858101830151858201604001528201610096565b818111156100c4576000604083870101525b50601f01601f1916929092016040019392505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561010257600080fd5b813567ffffffffffffffff8082111561011a57600080fd5b818401915084601f83011261012e57600080fd5b813581811115610140576101406100da565b604051601f8201601f19908116603f01168101908382118183101715610168576101686100da565b8160405282815287602084870101111561018157600080fd5b82602086016020830137600092810160200192909252509594505050505056fea264697066735822122036a694f35da97c94a51477e90c5cdeb6feea488db2ee1ee66a6b0005e20eba2d64736f6c634300080b0033"

// DeployContractToken deploys a new Ethereum contract, binding an instance of ContractToken to it.
func DeployContractToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractToken{ContractTokenCaller: ContractTokenCaller{contract: contract}, ContractTokenTransactor: ContractTokenTransactor{contract: contract}, ContractTokenFilterer: ContractTokenFilterer{contract: contract}}, nil
}

// ContractToken is an auto generated Go binding around an Ethereum contract.
type ContractToken struct {
	ContractTokenCaller     // Read-only binding to the contract
	ContractTokenTransactor // Write-only binding to the contract
	ContractTokenFilterer   // Log filterer for contract events
}

// ContractTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractTokenSession struct {
	Contract     *ContractToken              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractTokenCallerSession struct {
	Contract *ContractTokenCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ContractTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTokenTransactorSession struct {
	Contract     *ContractTokenTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractTokenRaw struct {
	Contract *ContractToken // Generic contract binding to access the raw methods on
}

// ContractTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractTokenCallerRaw struct {
	Contract *ContractTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTokenTransactorRaw struct {
	Contract *ContractTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractToken creates a new instance of ContractToken, bound to a specific deployed contract.
func NewContractToken(address common.Address, backend bind.ContractBackend) (*ContractToken, error) {
	contract, err := bindContractToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractToken{ContractTokenCaller: ContractTokenCaller{contract: contract}, ContractTokenTransactor: ContractTokenTransactor{contract: contract}, ContractTokenFilterer: ContractTokenFilterer{contract: contract}}, nil
}

// NewContractTokenCaller creates a new read-only instance of ContractToken, bound to a specific deployed contract.
func NewContractTokenCaller(address common.Address, caller bind.ContractCaller) (*ContractTokenCaller, error) {
	contract, err := bindContractToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTokenCaller{contract: contract}, nil
}

// NewContractTokenTransactor creates a new write-only instance of ContractToken, bound to a specific deployed contract.
func NewContractTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTokenTransactor, error) {
	contract, err := bindContractToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTokenTransactor{contract: contract}, nil
}

// NewContractTokenFilterer creates a new log filterer instance of ContractToken, bound to a specific deployed contract.
func NewContractTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractTokenFilterer, error) {
	contract, err := bindContractToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractTokenFilterer{contract: contract}, nil
}

// bindContractToken binds a generic wrapper to an already deployed contract.
func bindContractToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractToken *ContractTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractToken.Contract.ContractTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractToken *ContractTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractToken.Contract.ContractTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractToken *ContractTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractToken.Contract.ContractTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractToken *ContractTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractToken *ContractTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractToken *ContractTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractToken.Contract.contract.Transact(opts, method, params...)
}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) view returns(string)
func (_ContractToken *ContractTokenCaller) Greet(opts *bind.CallOpts, str string) (string, error) {
	var out []interface{}
	err := _ContractToken.contract.Call(opts, &out, "Greet", str)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) view returns(string)
func (_ContractToken *ContractTokenSession) Greet(str string) (string, error) {
	return _ContractToken.Contract.Greet(&_ContractToken.CallOpts, str)
}

// Greet is a free data retrieval call binding the contract method 0xefdeaaf5.
//
// Solidity: function Greet(string str) view returns(string)
func (_ContractToken *ContractTokenCallerSession) Greet(str string) (string, error) {
	return _ContractToken.Contract.Greet(&_ContractToken.CallOpts, str)
}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() view returns(string)
func (_ContractToken *ContractTokenCaller) Hello(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractToken.contract.Call(opts, &out, "Hello")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() view returns(string)
func (_ContractToken *ContractTokenSession) Hello() (string, error) {
	return _ContractToken.Contract.Hello(&_ContractToken.CallOpts)
}

// Hello is a free data retrieval call binding the contract method 0xbcdfe0d5.
//
// Solidity: function Hello() view returns(string)
func (_ContractToken *ContractTokenCallerSession) Hello() (string, error) {
	return _ContractToken.Contract.Hello(&_ContractToken.CallOpts)
}
