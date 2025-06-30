// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StdConstants

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

// StdConstantsMetaData contains all meta data concerning the StdConstants contract.
var StdConstantsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220d45c60b11746ce9061a9fc1913dae162a90146eb2cdbf2cc1943fe25f825a3b864736f6c634300081b0033",
}

// StdConstantsABI is the input ABI used to generate the binding from.
// Deprecated: Use StdConstantsMetaData.ABI instead.
var StdConstantsABI = StdConstantsMetaData.ABI

// StdConstantsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdConstantsMetaData.Bin instead.
var StdConstantsBin = StdConstantsMetaData.Bin

// DeployStdConstants deploys a new Ethereum contract, binding an instance of StdConstants to it.
func DeployStdConstants(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdConstants, error) {
	parsed, err := StdConstantsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdConstantsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdConstants{StdConstantsCaller: StdConstantsCaller{contract: contract}, StdConstantsTransactor: StdConstantsTransactor{contract: contract}, StdConstantsFilterer: StdConstantsFilterer{contract: contract}}, nil
}

// StdConstants is an auto generated Go binding around an Ethereum contract.
type StdConstants struct {
	StdConstantsCaller     // Read-only binding to the contract
	StdConstantsTransactor // Write-only binding to the contract
	StdConstantsFilterer   // Log filterer for contract events
}

// StdConstantsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdConstantsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdConstantsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdConstantsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdConstantsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdConstantsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdConstantsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdConstantsSession struct {
	Contract     *StdConstants     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdConstantsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdConstantsCallerSession struct {
	Contract *StdConstantsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StdConstantsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdConstantsTransactorSession struct {
	Contract     *StdConstantsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StdConstantsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdConstantsRaw struct {
	Contract *StdConstants // Generic contract binding to access the raw methods on
}

// StdConstantsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdConstantsCallerRaw struct {
	Contract *StdConstantsCaller // Generic read-only contract binding to access the raw methods on
}

// StdConstantsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdConstantsTransactorRaw struct {
	Contract *StdConstantsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdConstants creates a new instance of StdConstants, bound to a specific deployed contract.
func NewStdConstants(address common.Address, backend bind.ContractBackend) (*StdConstants, error) {
	contract, err := bindStdConstants(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdConstants{StdConstantsCaller: StdConstantsCaller{contract: contract}, StdConstantsTransactor: StdConstantsTransactor{contract: contract}, StdConstantsFilterer: StdConstantsFilterer{contract: contract}}, nil
}

// NewStdConstantsCaller creates a new read-only instance of StdConstants, bound to a specific deployed contract.
func NewStdConstantsCaller(address common.Address, caller bind.ContractCaller) (*StdConstantsCaller, error) {
	contract, err := bindStdConstants(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdConstantsCaller{contract: contract}, nil
}

// NewStdConstantsTransactor creates a new write-only instance of StdConstants, bound to a specific deployed contract.
func NewStdConstantsTransactor(address common.Address, transactor bind.ContractTransactor) (*StdConstantsTransactor, error) {
	contract, err := bindStdConstants(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdConstantsTransactor{contract: contract}, nil
}

// NewStdConstantsFilterer creates a new log filterer instance of StdConstants, bound to a specific deployed contract.
func NewStdConstantsFilterer(address common.Address, filterer bind.ContractFilterer) (*StdConstantsFilterer, error) {
	contract, err := bindStdConstants(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdConstantsFilterer{contract: contract}, nil
}

// bindStdConstants binds a generic wrapper to an already deployed contract.
func bindStdConstants(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdConstantsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdConstants *StdConstantsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdConstants.Contract.StdConstantsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdConstants *StdConstantsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdConstants.Contract.StdConstantsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdConstants *StdConstantsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdConstants.Contract.StdConstantsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdConstants *StdConstantsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdConstants.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdConstants *StdConstantsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdConstants.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdConstants *StdConstantsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdConstants.Contract.contract.Transact(opts, method, params...)
}
