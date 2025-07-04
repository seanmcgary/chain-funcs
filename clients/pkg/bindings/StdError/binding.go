// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StdError

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

// StdErrorMetaData contains all meta data concerning the StdError contract.
var StdErrorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"arithmeticError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"assertionError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"divisionError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"encodeStorageError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enumConversionError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"indexOOBError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"memOverflowError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"popError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"zeroVarError\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x610233610034600b8282823980515f1a607314602857634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061009b575f3560e01c8063986c5f681161006e578063986c5f68146100d5578063b22dc54d146100dd578063b67689da146100e5578063d160e4de146100ed578063fa784a44146100f5575f5ffd5b806305ee86121461009f57806310332977146100bd5780631de45560146100c55780638995290f146100cd575b5f5ffd5b6100a76100fd565b6040516100b491906101c8565b60405180910390f35b6100a7610138565b6100a761014a565b6100a761015c565b6100a761016e565b6100a7610180565b6100a7610192565b6100a76101a4565b6100a76101b6565b604051603260248201526044015b60408051601f198184030181529190526020810180516001600160e01b0316634e487b7160e01b17905281565b6040516001602482015260440161010b565b6040516021602482015260440161010b565b6040516011602482015260440161010b565b6040516041602482015260440161010b565b6040516031602482015260440161010b565b6040516051602482015260440161010b565b6040516022602482015260440161010b565b6040516012602482015260440161010b565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f8301168401019150509291505056fea2646970667358221220c915f5b778dfac096c4154908c328a261ffba87eb4eaab1d0d6d4bcd830ca7fe64736f6c634300081b0033",
}

// StdErrorABI is the input ABI used to generate the binding from.
// Deprecated: Use StdErrorMetaData.ABI instead.
var StdErrorABI = StdErrorMetaData.ABI

// StdErrorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StdErrorMetaData.Bin instead.
var StdErrorBin = StdErrorMetaData.Bin

// DeployStdError deploys a new Ethereum contract, binding an instance of StdError to it.
func DeployStdError(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StdError, error) {
	parsed, err := StdErrorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StdErrorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StdError{StdErrorCaller: StdErrorCaller{contract: contract}, StdErrorTransactor: StdErrorTransactor{contract: contract}, StdErrorFilterer: StdErrorFilterer{contract: contract}}, nil
}

// StdError is an auto generated Go binding around an Ethereum contract.
type StdError struct {
	StdErrorCaller     // Read-only binding to the contract
	StdErrorTransactor // Write-only binding to the contract
	StdErrorFilterer   // Log filterer for contract events
}

// StdErrorCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdErrorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdErrorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdErrorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdErrorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdErrorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdErrorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdErrorSession struct {
	Contract     *StdError         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdErrorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdErrorCallerSession struct {
	Contract *StdErrorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StdErrorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdErrorTransactorSession struct {
	Contract     *StdErrorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StdErrorRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdErrorRaw struct {
	Contract *StdError // Generic contract binding to access the raw methods on
}

// StdErrorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdErrorCallerRaw struct {
	Contract *StdErrorCaller // Generic read-only contract binding to access the raw methods on
}

// StdErrorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdErrorTransactorRaw struct {
	Contract *StdErrorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdError creates a new instance of StdError, bound to a specific deployed contract.
func NewStdError(address common.Address, backend bind.ContractBackend) (*StdError, error) {
	contract, err := bindStdError(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdError{StdErrorCaller: StdErrorCaller{contract: contract}, StdErrorTransactor: StdErrorTransactor{contract: contract}, StdErrorFilterer: StdErrorFilterer{contract: contract}}, nil
}

// NewStdErrorCaller creates a new read-only instance of StdError, bound to a specific deployed contract.
func NewStdErrorCaller(address common.Address, caller bind.ContractCaller) (*StdErrorCaller, error) {
	contract, err := bindStdError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdErrorCaller{contract: contract}, nil
}

// NewStdErrorTransactor creates a new write-only instance of StdError, bound to a specific deployed contract.
func NewStdErrorTransactor(address common.Address, transactor bind.ContractTransactor) (*StdErrorTransactor, error) {
	contract, err := bindStdError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdErrorTransactor{contract: contract}, nil
}

// NewStdErrorFilterer creates a new log filterer instance of StdError, bound to a specific deployed contract.
func NewStdErrorFilterer(address common.Address, filterer bind.ContractFilterer) (*StdErrorFilterer, error) {
	contract, err := bindStdError(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdErrorFilterer{contract: contract}, nil
}

// bindStdError binds a generic wrapper to an already deployed contract.
func bindStdError(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdErrorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdError *StdErrorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdError.Contract.StdErrorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdError *StdErrorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdError.Contract.StdErrorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdError *StdErrorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdError.Contract.StdErrorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdError *StdErrorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdError.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdError *StdErrorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdError.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdError *StdErrorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdError.Contract.contract.Transact(opts, method, params...)
}

// ArithmeticError is a free data retrieval call binding the contract method 0x8995290f.
//
// Solidity: function arithmeticError() view returns(bytes)
func (_StdError *StdErrorCaller) ArithmeticError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "arithmeticError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ArithmeticError is a free data retrieval call binding the contract method 0x8995290f.
//
// Solidity: function arithmeticError() view returns(bytes)
func (_StdError *StdErrorSession) ArithmeticError() ([]byte, error) {
	return _StdError.Contract.ArithmeticError(&_StdError.CallOpts)
}

// ArithmeticError is a free data retrieval call binding the contract method 0x8995290f.
//
// Solidity: function arithmeticError() view returns(bytes)
func (_StdError *StdErrorCallerSession) ArithmeticError() ([]byte, error) {
	return _StdError.Contract.ArithmeticError(&_StdError.CallOpts)
}

// AssertionError is a free data retrieval call binding the contract method 0x10332977.
//
// Solidity: function assertionError() view returns(bytes)
func (_StdError *StdErrorCaller) AssertionError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "assertionError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// AssertionError is a free data retrieval call binding the contract method 0x10332977.
//
// Solidity: function assertionError() view returns(bytes)
func (_StdError *StdErrorSession) AssertionError() ([]byte, error) {
	return _StdError.Contract.AssertionError(&_StdError.CallOpts)
}

// AssertionError is a free data retrieval call binding the contract method 0x10332977.
//
// Solidity: function assertionError() view returns(bytes)
func (_StdError *StdErrorCallerSession) AssertionError() ([]byte, error) {
	return _StdError.Contract.AssertionError(&_StdError.CallOpts)
}

// DivisionError is a free data retrieval call binding the contract method 0xfa784a44.
//
// Solidity: function divisionError() view returns(bytes)
func (_StdError *StdErrorCaller) DivisionError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "divisionError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DivisionError is a free data retrieval call binding the contract method 0xfa784a44.
//
// Solidity: function divisionError() view returns(bytes)
func (_StdError *StdErrorSession) DivisionError() ([]byte, error) {
	return _StdError.Contract.DivisionError(&_StdError.CallOpts)
}

// DivisionError is a free data retrieval call binding the contract method 0xfa784a44.
//
// Solidity: function divisionError() view returns(bytes)
func (_StdError *StdErrorCallerSession) DivisionError() ([]byte, error) {
	return _StdError.Contract.DivisionError(&_StdError.CallOpts)
}

// EncodeStorageError is a free data retrieval call binding the contract method 0xd160e4de.
//
// Solidity: function encodeStorageError() view returns(bytes)
func (_StdError *StdErrorCaller) EncodeStorageError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "encodeStorageError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodeStorageError is a free data retrieval call binding the contract method 0xd160e4de.
//
// Solidity: function encodeStorageError() view returns(bytes)
func (_StdError *StdErrorSession) EncodeStorageError() ([]byte, error) {
	return _StdError.Contract.EncodeStorageError(&_StdError.CallOpts)
}

// EncodeStorageError is a free data retrieval call binding the contract method 0xd160e4de.
//
// Solidity: function encodeStorageError() view returns(bytes)
func (_StdError *StdErrorCallerSession) EncodeStorageError() ([]byte, error) {
	return _StdError.Contract.EncodeStorageError(&_StdError.CallOpts)
}

// EnumConversionError is a free data retrieval call binding the contract method 0x1de45560.
//
// Solidity: function enumConversionError() view returns(bytes)
func (_StdError *StdErrorCaller) EnumConversionError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "enumConversionError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnumConversionError is a free data retrieval call binding the contract method 0x1de45560.
//
// Solidity: function enumConversionError() view returns(bytes)
func (_StdError *StdErrorSession) EnumConversionError() ([]byte, error) {
	return _StdError.Contract.EnumConversionError(&_StdError.CallOpts)
}

// EnumConversionError is a free data retrieval call binding the contract method 0x1de45560.
//
// Solidity: function enumConversionError() view returns(bytes)
func (_StdError *StdErrorCallerSession) EnumConversionError() ([]byte, error) {
	return _StdError.Contract.EnumConversionError(&_StdError.CallOpts)
}

// IndexOOBError is a free data retrieval call binding the contract method 0x05ee8612.
//
// Solidity: function indexOOBError() view returns(bytes)
func (_StdError *StdErrorCaller) IndexOOBError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "indexOOBError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// IndexOOBError is a free data retrieval call binding the contract method 0x05ee8612.
//
// Solidity: function indexOOBError() view returns(bytes)
func (_StdError *StdErrorSession) IndexOOBError() ([]byte, error) {
	return _StdError.Contract.IndexOOBError(&_StdError.CallOpts)
}

// IndexOOBError is a free data retrieval call binding the contract method 0x05ee8612.
//
// Solidity: function indexOOBError() view returns(bytes)
func (_StdError *StdErrorCallerSession) IndexOOBError() ([]byte, error) {
	return _StdError.Contract.IndexOOBError(&_StdError.CallOpts)
}

// MemOverflowError is a free data retrieval call binding the contract method 0x986c5f68.
//
// Solidity: function memOverflowError() view returns(bytes)
func (_StdError *StdErrorCaller) MemOverflowError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "memOverflowError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MemOverflowError is a free data retrieval call binding the contract method 0x986c5f68.
//
// Solidity: function memOverflowError() view returns(bytes)
func (_StdError *StdErrorSession) MemOverflowError() ([]byte, error) {
	return _StdError.Contract.MemOverflowError(&_StdError.CallOpts)
}

// MemOverflowError is a free data retrieval call binding the contract method 0x986c5f68.
//
// Solidity: function memOverflowError() view returns(bytes)
func (_StdError *StdErrorCallerSession) MemOverflowError() ([]byte, error) {
	return _StdError.Contract.MemOverflowError(&_StdError.CallOpts)
}

// PopError is a free data retrieval call binding the contract method 0xb22dc54d.
//
// Solidity: function popError() view returns(bytes)
func (_StdError *StdErrorCaller) PopError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "popError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// PopError is a free data retrieval call binding the contract method 0xb22dc54d.
//
// Solidity: function popError() view returns(bytes)
func (_StdError *StdErrorSession) PopError() ([]byte, error) {
	return _StdError.Contract.PopError(&_StdError.CallOpts)
}

// PopError is a free data retrieval call binding the contract method 0xb22dc54d.
//
// Solidity: function popError() view returns(bytes)
func (_StdError *StdErrorCallerSession) PopError() ([]byte, error) {
	return _StdError.Contract.PopError(&_StdError.CallOpts)
}

// ZeroVarError is a free data retrieval call binding the contract method 0xb67689da.
//
// Solidity: function zeroVarError() view returns(bytes)
func (_StdError *StdErrorCaller) ZeroVarError(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _StdError.contract.Call(opts, &out, "zeroVarError")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ZeroVarError is a free data retrieval call binding the contract method 0xb67689da.
//
// Solidity: function zeroVarError() view returns(bytes)
func (_StdError *StdErrorSession) ZeroVarError() ([]byte, error) {
	return _StdError.Contract.ZeroVarError(&_StdError.CallOpts)
}

// ZeroVarError is a free data retrieval call binding the contract method 0xb67689da.
//
// Solidity: function zeroVarError() view returns(bytes)
func (_StdError *StdErrorCallerSession) ZeroVarError() ([]byte, error) {
	return _StdError.Contract.ZeroVarError(&_StdError.CallOpts)
}
