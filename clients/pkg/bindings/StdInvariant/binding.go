// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StdInvariant

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

// StdInvariantFuzzArtifactSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzArtifactSelector struct {
	Artifact  string
	Selectors [][4]byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// StdInvariantMetaData contains all meta data concerning the StdInvariant contract.
var StdInvariantMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzArtifactSelector[]\",\"components\":[{\"name\":\"artifact\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"}]",
}

// StdInvariantABI is the input ABI used to generate the binding from.
// Deprecated: Use StdInvariantMetaData.ABI instead.
var StdInvariantABI = StdInvariantMetaData.ABI

// StdInvariant is an auto generated Go binding around an Ethereum contract.
type StdInvariant struct {
	StdInvariantCaller     // Read-only binding to the contract
	StdInvariantTransactor // Write-only binding to the contract
	StdInvariantFilterer   // Log filterer for contract events
}

// StdInvariantCaller is an auto generated read-only Go binding around an Ethereum contract.
type StdInvariantCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdInvariantTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StdInvariantTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdInvariantFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StdInvariantFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StdInvariantSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StdInvariantSession struct {
	Contract     *StdInvariant     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StdInvariantCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StdInvariantCallerSession struct {
	Contract *StdInvariantCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StdInvariantTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StdInvariantTransactorSession struct {
	Contract     *StdInvariantTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StdInvariantRaw is an auto generated low-level Go binding around an Ethereum contract.
type StdInvariantRaw struct {
	Contract *StdInvariant // Generic contract binding to access the raw methods on
}

// StdInvariantCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StdInvariantCallerRaw struct {
	Contract *StdInvariantCaller // Generic read-only contract binding to access the raw methods on
}

// StdInvariantTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StdInvariantTransactorRaw struct {
	Contract *StdInvariantTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStdInvariant creates a new instance of StdInvariant, bound to a specific deployed contract.
func NewStdInvariant(address common.Address, backend bind.ContractBackend) (*StdInvariant, error) {
	contract, err := bindStdInvariant(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StdInvariant{StdInvariantCaller: StdInvariantCaller{contract: contract}, StdInvariantTransactor: StdInvariantTransactor{contract: contract}, StdInvariantFilterer: StdInvariantFilterer{contract: contract}}, nil
}

// NewStdInvariantCaller creates a new read-only instance of StdInvariant, bound to a specific deployed contract.
func NewStdInvariantCaller(address common.Address, caller bind.ContractCaller) (*StdInvariantCaller, error) {
	contract, err := bindStdInvariant(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StdInvariantCaller{contract: contract}, nil
}

// NewStdInvariantTransactor creates a new write-only instance of StdInvariant, bound to a specific deployed contract.
func NewStdInvariantTransactor(address common.Address, transactor bind.ContractTransactor) (*StdInvariantTransactor, error) {
	contract, err := bindStdInvariant(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StdInvariantTransactor{contract: contract}, nil
}

// NewStdInvariantFilterer creates a new log filterer instance of StdInvariant, bound to a specific deployed contract.
func NewStdInvariantFilterer(address common.Address, filterer bind.ContractFilterer) (*StdInvariantFilterer, error) {
	contract, err := bindStdInvariant(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StdInvariantFilterer{contract: contract}, nil
}

// bindStdInvariant binds a generic wrapper to an already deployed contract.
func bindStdInvariant(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StdInvariantMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdInvariant *StdInvariantRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdInvariant.Contract.StdInvariantCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdInvariant *StdInvariantRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdInvariant.Contract.StdInvariantTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdInvariant *StdInvariantRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdInvariant.Contract.StdInvariantTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StdInvariant *StdInvariantCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StdInvariant.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StdInvariant *StdInvariantTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StdInvariant.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StdInvariant *StdInvariantTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StdInvariant.Contract.contract.Transact(opts, method, params...)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_StdInvariant *StdInvariantCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_StdInvariant *StdInvariantSession) ExcludeArtifacts() ([]string, error) {
	return _StdInvariant.Contract.ExcludeArtifacts(&_StdInvariant.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_StdInvariant *StdInvariantCallerSession) ExcludeArtifacts() ([]string, error) {
	return _StdInvariant.Contract.ExcludeArtifacts(&_StdInvariant.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_StdInvariant *StdInvariantCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_StdInvariant *StdInvariantSession) ExcludeContracts() ([]common.Address, error) {
	return _StdInvariant.Contract.ExcludeContracts(&_StdInvariant.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_StdInvariant *StdInvariantCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _StdInvariant.Contract.ExcludeContracts(&_StdInvariant.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_StdInvariant *StdInvariantCaller) ExcludeSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "excludeSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_StdInvariant *StdInvariantSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _StdInvariant.Contract.ExcludeSelectors(&_StdInvariant.CallOpts)
}

// ExcludeSelectors is a free data retrieval call binding the contract method 0xb0464fdc.
//
// Solidity: function excludeSelectors() view returns((address,bytes4[])[] excludedSelectors_)
func (_StdInvariant *StdInvariantCallerSession) ExcludeSelectors() ([]StdInvariantFuzzSelector, error) {
	return _StdInvariant.Contract.ExcludeSelectors(&_StdInvariant.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_StdInvariant *StdInvariantCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_StdInvariant *StdInvariantSession) ExcludeSenders() ([]common.Address, error) {
	return _StdInvariant.Contract.ExcludeSenders(&_StdInvariant.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_StdInvariant *StdInvariantCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _StdInvariant.Contract.ExcludeSenders(&_StdInvariant.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_StdInvariant *StdInvariantCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzArtifactSelector, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzArtifactSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzArtifactSelector)).(*[]StdInvariantFuzzArtifactSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_StdInvariant *StdInvariantSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _StdInvariant.Contract.TargetArtifactSelectors(&_StdInvariant.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((string,bytes4[])[] targetedArtifactSelectors_)
func (_StdInvariant *StdInvariantCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzArtifactSelector, error) {
	return _StdInvariant.Contract.TargetArtifactSelectors(&_StdInvariant.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_StdInvariant *StdInvariantCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_StdInvariant *StdInvariantSession) TargetArtifacts() ([]string, error) {
	return _StdInvariant.Contract.TargetArtifacts(&_StdInvariant.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_StdInvariant *StdInvariantCallerSession) TargetArtifacts() ([]string, error) {
	return _StdInvariant.Contract.TargetArtifacts(&_StdInvariant.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_StdInvariant *StdInvariantCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_StdInvariant *StdInvariantSession) TargetContracts() ([]common.Address, error) {
	return _StdInvariant.Contract.TargetContracts(&_StdInvariant.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_StdInvariant *StdInvariantCallerSession) TargetContracts() ([]common.Address, error) {
	return _StdInvariant.Contract.TargetContracts(&_StdInvariant.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_StdInvariant *StdInvariantCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_StdInvariant *StdInvariantSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _StdInvariant.Contract.TargetInterfaces(&_StdInvariant.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_StdInvariant *StdInvariantCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _StdInvariant.Contract.TargetInterfaces(&_StdInvariant.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_StdInvariant *StdInvariantCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_StdInvariant *StdInvariantSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _StdInvariant.Contract.TargetSelectors(&_StdInvariant.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_StdInvariant *StdInvariantCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _StdInvariant.Contract.TargetSelectors(&_StdInvariant.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_StdInvariant *StdInvariantCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StdInvariant.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_StdInvariant *StdInvariantSession) TargetSenders() ([]common.Address, error) {
	return _StdInvariant.Contract.TargetSenders(&_StdInvariant.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_StdInvariant *StdInvariantCallerSession) TargetSenders() ([]common.Address, error) {
	return _StdInvariant.Contract.TargetSenders(&_StdInvariant.CallOpts)
}
