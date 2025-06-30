// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StdUtils

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

// StdUtilsMetaData contains all meta data concerning the StdUtils contract.
var StdUtilsMetaData = &bind.MetaData{
	ABI: "[]",
}

// StdUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use StdUtilsMetaData.ABI instead.
var StdUtilsABI = StdUtilsMetaData.ABI

// StdUtils is an auto generated Go binding around an Ethereum contract.
type StdUtils struct {
	StdUtilsCaller     // Read-only binding to the contract
	StdUtilsTransactor // Write-only binding to the contract
	StdUtilsFilterer   // Log filterer for contract events
}

// StdUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdUtilsSession struct {
	Contract     *StdUtils         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdUtilsCallerSession struct {
	Contract *StdUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StdUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdUtilsTransactorSession struct {
	Contract     *StdUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StdUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdUtilsRaw struct {
	Contract *StdUtils // Generic contract binding to access the raw methods on
}

// StdUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdUtilsCallerRaw struct {
	Contract *StdUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// StdUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdUtilsTransactorRaw struct {
	Contract *StdUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdUtils creates a new instance of StdUtils, bound to a specific deployed contract.
func NewStdUtils(address common.Address, backend bind.ContractBackend) (*StdUtils, error) {
	contract, err := bindStdUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdUtils{StdUtilsCaller: StdUtilsCaller{contract: contract}, StdUtilsTransactor: StdUtilsTransactor{contract: contract}, StdUtilsFilterer: StdUtilsFilterer{contract: contract}}, nil
}

// NewStdUtilsCaller creates a new read-only instance of StdUtils, bound to a specific deployed contract.
func NewStdUtilsCaller(address common.Address, caller bind.ContractCaller) (*StdUtilsCaller, error) {
	contract, err := bindStdUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdUtilsCaller{contract: contract}, nil
}

// NewStdUtilsTransactor creates a new write-only instance of StdUtils, bound to a specific deployed contract.
func NewStdUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*StdUtilsTransactor, error) {
	contract, err := bindStdUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdUtilsTransactor{contract: contract}, nil
}

// NewStdUtilsFilterer creates a new log filterer instance of StdUtils, bound to a specific deployed contract.
func NewStdUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*StdUtilsFilterer, error) {
	contract, err := bindStdUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdUtilsFilterer{contract: contract}, nil
}

// bindStdUtils binds a generic wrapper to an already deployed contract.
func bindStdUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdUtilsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdUtils *StdUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdUtils.Contract.StdUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdUtils *StdUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdUtils.Contract.StdUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdUtils *StdUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdUtils.Contract.StdUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdUtils *StdUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdUtils *StdUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdUtils *StdUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdUtils.Contract.contract.Transact(opts, method, params...)
}
