// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Script

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

// ScriptMetaData contains all meta data concerning the Script contract.
var ScriptMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_SCRIPT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
}

// ScriptABI is the input ABI used to generate the binding from.
// Deprecated: Use ScriptMetaData.ABI instead.
var ScriptABI = ScriptMetaData.ABI

// Script is an auto generated Go binding around an Ethereum contract.
type Script struct {
	ScriptCaller     // Read-only binding to the contract
	ScriptTransactor // Write-only binding to the contract
	ScriptFilterer   // Log filterer for contract events
}

// ScriptCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScriptCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScriptTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScriptFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScriptSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScriptSession struct {
	Contract     *Script           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScriptCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScriptCallerSession struct {
	Contract *ScriptCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ScriptTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScriptTransactorSession struct {
	Contract     *ScriptTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScriptRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScriptRaw struct {
	Contract *Script // Generic contract binding to access the raw methods on
}

// ScriptCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScriptCallerRaw struct {
	Contract *ScriptCaller // Generic read-only contract binding to access the raw methods on
}

// ScriptTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScriptTransactorRaw struct {
	Contract *ScriptTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScript creates a new instance of Script, bound to a specific deployed contract.
func NewScript(address common.Address, backend bind.ContractBackend) (*Script, error) {
	contract, err := bindScript(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Script{ScriptCaller: ScriptCaller{contract: contract}, ScriptTransactor: ScriptTransactor{contract: contract}, ScriptFilterer: ScriptFilterer{contract: contract}}, nil
}

// NewScriptCaller creates a new read-only instance of Script, bound to a specific deployed contract.
func NewScriptCaller(address common.Address, caller bind.ContractCaller) (*ScriptCaller, error) {
	contract, err := bindScript(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScriptCaller{contract: contract}, nil
}

// NewScriptTransactor creates a new write-only instance of Script, bound to a specific deployed contract.
func NewScriptTransactor(address common.Address, transactor bind.ContractTransactor) (*ScriptTransactor, error) {
	contract, err := bindScript(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScriptTransactor{contract: contract}, nil
}

// NewScriptFilterer creates a new log filterer instance of Script, bound to a specific deployed contract.
func NewScriptFilterer(address common.Address, filterer bind.ContractFilterer) (*ScriptFilterer, error) {
	contract, err := bindScript(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScriptFilterer{contract: contract}, nil
}

// bindScript binds a generic wrapper to an already deployed contract.
func bindScript(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScriptMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Script *ScriptRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Script.Contract.ScriptCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Script *ScriptRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Script.Contract.ScriptTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Script *ScriptRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Script.Contract.ScriptTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Script *ScriptCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Script.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Script *ScriptTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Script.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Script *ScriptTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Script.Contract.contract.Transact(opts, method, params...)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_Script *ScriptCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Script.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_Script *ScriptSession) ISSCRIPT() (bool, error) {
	return _Script.Contract.ISSCRIPT(&_Script.CallOpts)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_Script *ScriptCallerSession) ISSCRIPT() (bool, error) {
	return _Script.Contract.ISSCRIPT(&_Script.CallOpts)
}
