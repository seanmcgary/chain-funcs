// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StdCheats

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

// StdCheatsMetaData contains all meta data concerning the StdCheats contract.
var StdCheatsMetaData = &bind.MetaData{
	ABI: "[]",
}

// StdCheatsABI is the input ABI used to generate the binding from.
// Deprecated: Use StdCheatsMetaData.ABI instead.
var StdCheatsABI = StdCheatsMetaData.ABI

// StdCheats is an auto generated Go binding around an Ethereum contract.
type StdCheats struct {
	StdCheatsCaller     // Read-only binding to the contract
	StdCheatsTransactor // Write-only binding to the contract
	StdCheatsFilterer   // Log filterer for contract events
}

// StdCheatsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdCheatsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdCheatsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdCheatsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdCheatsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdCheatsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdCheatsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdCheatsSession struct {
	Contract     *StdCheats        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdCheatsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdCheatsCallerSession struct {
	Contract *StdCheatsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StdCheatsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdCheatsTransactorSession struct {
	Contract     *StdCheatsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StdCheatsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdCheatsRaw struct {
	Contract *StdCheats // Generic contract binding to access the raw methods on
}

// StdCheatsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdCheatsCallerRaw struct {
	Contract *StdCheatsCaller // Generic read-only contract binding to access the raw methods on
}

// StdCheatsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdCheatsTransactorRaw struct {
	Contract *StdCheatsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdCheats creates a new instance of StdCheats, bound to a specific deployed contract.
func NewStdCheats(address common.Address, backend bind.ContractBackend) (*StdCheats, error) {
	contract, err := bindStdCheats(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdCheats{StdCheatsCaller: StdCheatsCaller{contract: contract}, StdCheatsTransactor: StdCheatsTransactor{contract: contract}, StdCheatsFilterer: StdCheatsFilterer{contract: contract}}, nil
}

// NewStdCheatsCaller creates a new read-only instance of StdCheats, bound to a specific deployed contract.
func NewStdCheatsCaller(address common.Address, caller bind.ContractCaller) (*StdCheatsCaller, error) {
	contract, err := bindStdCheats(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdCheatsCaller{contract: contract}, nil
}

// NewStdCheatsTransactor creates a new write-only instance of StdCheats, bound to a specific deployed contract.
func NewStdCheatsTransactor(address common.Address, transactor bind.ContractTransactor) (*StdCheatsTransactor, error) {
	contract, err := bindStdCheats(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdCheatsTransactor{contract: contract}, nil
}

// NewStdCheatsFilterer creates a new log filterer instance of StdCheats, bound to a specific deployed contract.
func NewStdCheatsFilterer(address common.Address, filterer bind.ContractFilterer) (*StdCheatsFilterer, error) {
	contract, err := bindStdCheats(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdCheatsFilterer{contract: contract}, nil
}

// bindStdCheats binds a generic wrapper to an already deployed contract.
func bindStdCheats(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdCheatsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdCheats *StdCheatsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdCheats.Contract.StdCheatsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdCheats *StdCheatsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdCheats.Contract.StdCheatsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdCheats *StdCheatsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdCheats.Contract.StdCheatsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdCheats *StdCheatsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdCheats.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdCheats *StdCheatsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdCheats.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdCheats *StdCheatsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdCheats.Contract.contract.Transact(opts, method, params...)
}
