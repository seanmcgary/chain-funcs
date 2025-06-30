// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StdChains

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
	_ = abi.ConvertType
)

// StdChainsMetaData contains all meta data concerning the StdChains contract.
var StdChainsMetaData = &bind.MetaData{
	ABI: "[]",
}

// StdChainsABI is the input ABI used to generate the binding from.
// Deprecated: Use StdChainsMetaData.ABI instead.
var StdChainsABI = StdChainsMetaData.ABI

// StdChains is an auto generated Go binding around an Ethereum contract.
type StdChains struct {
	StdChainsCaller     // Read-only binding to the contract
	StdChainsTransactor // Write-only binding to the contract
	StdChainsFilterer   // Log filterer for contract events
}

// StdChainsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdChainsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdChainsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdChainsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdChainsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdChainsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdChainsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdChainsSession struct {
	Contract     *StdChains        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdChainsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdChainsCallerSession struct {
	Contract *StdChainsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StdChainsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdChainsTransactorSession struct {
	Contract     *StdChainsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StdChainsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdChainsRaw struct {
	Contract *StdChains // Generic contract binding to access the raw methods on
}

// StdChainsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdChainsCallerRaw struct {
	Contract *StdChainsCaller // Generic read-only contract binding to access the raw methods on
}

// StdChainsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdChainsTransactorRaw struct {
	Contract *StdChainsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdChains creates a new instance of StdChains, bound to a specific deployed contract.
func NewStdChains(address common.Address, backend bind.ContractBackend) (*StdChains, error) {
	contract, err := bindStdChains(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdChains{StdChainsCaller: StdChainsCaller{contract: contract}, StdChainsTransactor: StdChainsTransactor{contract: contract}, StdChainsFilterer: StdChainsFilterer{contract: contract}}, nil
}

// NewStdChainsCaller creates a new read-only instance of StdChains, bound to a specific deployed contract.
func NewStdChainsCaller(address common.Address, caller bind.ContractCaller) (*StdChainsCaller, error) {
	contract, err := bindStdChains(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdChainsCaller{contract: contract}, nil
}

// NewStdChainsTransactor creates a new write-only instance of StdChains, bound to a specific deployed contract.
func NewStdChainsTransactor(address common.Address, transactor bind.ContractTransactor) (*StdChainsTransactor, error) {
	contract, err := bindStdChains(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdChainsTransactor{contract: contract}, nil
}

// NewStdChainsFilterer creates a new log filterer instance of StdChains, bound to a specific deployed contract.
func NewStdChainsFilterer(address common.Address, filterer bind.ContractFilterer) (*StdChainsFilterer, error) {
	contract, err := bindStdChains(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdChainsFilterer{contract: contract}, nil
}

// bindStdChains binds a generic wrapper to an already deployed contract.
func bindStdChains(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdChainsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdChains *StdChainsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdChains.Contract.StdChainsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdChains *StdChainsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdChains.Contract.StdChainsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdChains *StdChainsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdChains.Contract.StdChainsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdChains *StdChainsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdChains.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdChains *StdChainsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdChains.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdChains *StdChainsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdChains.Contract.contract.Transact(opts, method, params...)
}
