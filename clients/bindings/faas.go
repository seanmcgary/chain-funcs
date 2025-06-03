// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// FaaSFunction is an auto generated low-level Go binding around an user-defined struct.
type FaaSFunction struct {
	Content []byte
	Url     string
}

// FaaSMetaData contains all meta data concerning the FaaS contract.
var FaaSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskMailbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"callFunction\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployFunction\",\"inputs\":[{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executorOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"functions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"content\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunction\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFaaS.Function\",\"components\":[{\"name\":\"content\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunctionContent\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunctionUrl\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerFunction\",\"inputs\":[{\"name\":\"content\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskMailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FunctionCalled\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"taskId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FunctionDeployed\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"url\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FunctionRegistered\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// FaaSABI is the input ABI used to generate the binding from.
// Deprecated: Use FaaSMetaData.ABI instead.
var FaaSABI = FaaSMetaData.ABI

// FaaS is an auto generated Go binding around an Ethereum contract.
type FaaS struct {
	FaaSCaller     // Read-only binding to the contract
	FaaSTransactor // Write-only binding to the contract
	FaaSFilterer   // Log filterer for contract events
}

// FaaSCaller is an auto generated read-only Go binding around an Ethereum contract.
type FaaSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FaaSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FaaSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FaaSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FaaSSession struct {
	Contract     *FaaS             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaaSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FaaSCallerSession struct {
	Contract *FaaSCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FaaSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FaaSTransactorSession struct {
	Contract     *FaaSTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FaaSRaw is an auto generated low-level Go binding around an Ethereum contract.
type FaaSRaw struct {
	Contract *FaaS // Generic contract binding to access the raw methods on
}

// FaaSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FaaSCallerRaw struct {
	Contract *FaaSCaller // Generic read-only contract binding to access the raw methods on
}

// FaaSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FaaSTransactorRaw struct {
	Contract *FaaSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFaaS creates a new instance of FaaS, bound to a specific deployed contract.
func NewFaaS(address common.Address, backend bind.ContractBackend) (*FaaS, error) {
	contract, err := bindFaaS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FaaS{FaaSCaller: FaaSCaller{contract: contract}, FaaSTransactor: FaaSTransactor{contract: contract}, FaaSFilterer: FaaSFilterer{contract: contract}}, nil
}

// NewFaaSCaller creates a new read-only instance of FaaS, bound to a specific deployed contract.
func NewFaaSCaller(address common.Address, caller bind.ContractCaller) (*FaaSCaller, error) {
	contract, err := bindFaaS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSCaller{contract: contract}, nil
}

// NewFaaSTransactor creates a new write-only instance of FaaS, bound to a specific deployed contract.
func NewFaaSTransactor(address common.Address, transactor bind.ContractTransactor) (*FaaSTransactor, error) {
	contract, err := bindFaaS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FaaSTransactor{contract: contract}, nil
}

// NewFaaSFilterer creates a new log filterer instance of FaaS, bound to a specific deployed contract.
func NewFaaSFilterer(address common.Address, filterer bind.ContractFilterer) (*FaaSFilterer, error) {
	contract, err := bindFaaS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FaaSFilterer{contract: contract}, nil
}

// bindFaaS binds a generic wrapper to an already deployed contract.
func bindFaaS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FaaSMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaS *FaaSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaaS.Contract.FaaSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaS *FaaSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaaS.Contract.FaaSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaS *FaaSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaaS.Contract.FaaSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaS *FaaSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaaS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FaaS *FaaSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaaS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaS *FaaSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaaS.Contract.contract.Transact(opts, method, params...)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_FaaS *FaaSCaller) Avs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "avs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_FaaS *FaaSSession) Avs() (common.Address, error) {
	return _FaaS.Contract.Avs(&_FaaS.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_FaaS *FaaSCallerSession) Avs() (common.Address, error) {
	return _FaaS.Contract.Avs(&_FaaS.CallOpts)
}

// ExecutorOperatorSetId is a free data retrieval call binding the contract method 0xc428b075.
//
// Solidity: function executorOperatorSetId() view returns(uint32)
func (_FaaS *FaaSCaller) ExecutorOperatorSetId(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "executorOperatorSetId")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ExecutorOperatorSetId is a free data retrieval call binding the contract method 0xc428b075.
//
// Solidity: function executorOperatorSetId() view returns(uint32)
func (_FaaS *FaaSSession) ExecutorOperatorSetId() (uint32, error) {
	return _FaaS.Contract.ExecutorOperatorSetId(&_FaaS.CallOpts)
}

// ExecutorOperatorSetId is a free data retrieval call binding the contract method 0xc428b075.
//
// Solidity: function executorOperatorSetId() view returns(uint32)
func (_FaaS *FaaSCallerSession) ExecutorOperatorSetId() (uint32, error) {
	return _FaaS.Contract.ExecutorOperatorSetId(&_FaaS.CallOpts)
}

// Functions is a free data retrieval call binding the contract method 0xb4e02f94.
//
// Solidity: function functions(bytes32 ) view returns(bytes content, string url)
func (_FaaS *FaaSCaller) Functions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Content []byte
	Url     string
}, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "functions", arg0)

	outstruct := new(struct {
		Content []byte
		Url     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Content = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Url = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Functions is a free data retrieval call binding the contract method 0xb4e02f94.
//
// Solidity: function functions(bytes32 ) view returns(bytes content, string url)
func (_FaaS *FaaSSession) Functions(arg0 [32]byte) (struct {
	Content []byte
	Url     string
}, error) {
	return _FaaS.Contract.Functions(&_FaaS.CallOpts, arg0)
}

// Functions is a free data retrieval call binding the contract method 0xb4e02f94.
//
// Solidity: function functions(bytes32 ) view returns(bytes content, string url)
func (_FaaS *FaaSCallerSession) Functions(arg0 [32]byte) (struct {
	Content []byte
	Url     string
}, error) {
	return _FaaS.Contract.Functions(&_FaaS.CallOpts, arg0)
}

// GetFunction is a free data retrieval call binding the contract method 0xe67950d3.
//
// Solidity: function getFunction(bytes32 functionId) view returns((bytes,string))
func (_FaaS *FaaSCaller) GetFunction(opts *bind.CallOpts, functionId [32]byte) (FaaSFunction, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "getFunction", functionId)

	if err != nil {
		return *new(FaaSFunction), err
	}

	out0 := *abi.ConvertType(out[0], new(FaaSFunction)).(*FaaSFunction)

	return out0, err

}

// GetFunction is a free data retrieval call binding the contract method 0xe67950d3.
//
// Solidity: function getFunction(bytes32 functionId) view returns((bytes,string))
func (_FaaS *FaaSSession) GetFunction(functionId [32]byte) (FaaSFunction, error) {
	return _FaaS.Contract.GetFunction(&_FaaS.CallOpts, functionId)
}

// GetFunction is a free data retrieval call binding the contract method 0xe67950d3.
//
// Solidity: function getFunction(bytes32 functionId) view returns((bytes,string))
func (_FaaS *FaaSCallerSession) GetFunction(functionId [32]byte) (FaaSFunction, error) {
	return _FaaS.Contract.GetFunction(&_FaaS.CallOpts, functionId)
}

// GetFunctionContent is a free data retrieval call binding the contract method 0x24281fef.
//
// Solidity: function getFunctionContent(bytes32 functionId) view returns(bytes)
func (_FaaS *FaaSCaller) GetFunctionContent(opts *bind.CallOpts, functionId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "getFunctionContent", functionId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetFunctionContent is a free data retrieval call binding the contract method 0x24281fef.
//
// Solidity: function getFunctionContent(bytes32 functionId) view returns(bytes)
func (_FaaS *FaaSSession) GetFunctionContent(functionId [32]byte) ([]byte, error) {
	return _FaaS.Contract.GetFunctionContent(&_FaaS.CallOpts, functionId)
}

// GetFunctionContent is a free data retrieval call binding the contract method 0x24281fef.
//
// Solidity: function getFunctionContent(bytes32 functionId) view returns(bytes)
func (_FaaS *FaaSCallerSession) GetFunctionContent(functionId [32]byte) ([]byte, error) {
	return _FaaS.Contract.GetFunctionContent(&_FaaS.CallOpts, functionId)
}

// GetFunctionUrl is a free data retrieval call binding the contract method 0x3292c49c.
//
// Solidity: function getFunctionUrl(bytes32 functionId) view returns(string)
func (_FaaS *FaaSCaller) GetFunctionUrl(opts *bind.CallOpts, functionId [32]byte) (string, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "getFunctionUrl", functionId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFunctionUrl is a free data retrieval call binding the contract method 0x3292c49c.
//
// Solidity: function getFunctionUrl(bytes32 functionId) view returns(string)
func (_FaaS *FaaSSession) GetFunctionUrl(functionId [32]byte) (string, error) {
	return _FaaS.Contract.GetFunctionUrl(&_FaaS.CallOpts, functionId)
}

// GetFunctionUrl is a free data retrieval call binding the contract method 0x3292c49c.
//
// Solidity: function getFunctionUrl(bytes32 functionId) view returns(string)
func (_FaaS *FaaSCallerSession) GetFunctionUrl(functionId [32]byte) (string, error) {
	return _FaaS.Contract.GetFunctionUrl(&_FaaS.CallOpts, functionId)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_FaaS *FaaSCaller) TaskMailbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "taskMailbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_FaaS *FaaSSession) TaskMailbox() (common.Address, error) {
	return _FaaS.Contract.TaskMailbox(&_FaaS.CallOpts)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf42a9e13.
//
// Solidity: function taskMailbox() view returns(address)
func (_FaaS *FaaSCallerSession) TaskMailbox() (common.Address, error) {
	return _FaaS.Contract.TaskMailbox(&_FaaS.CallOpts)
}

// CallFunction is a paid mutator transaction binding the contract method 0x49262a73.
//
// Solidity: function callFunction(bytes32 functionId, bytes arguments) returns(bytes32)
func (_FaaS *FaaSTransactor) CallFunction(opts *bind.TransactOpts, functionId [32]byte, arguments []byte) (*types.Transaction, error) {
	return _FaaS.contract.Transact(opts, "callFunction", functionId, arguments)
}

// CallFunction is a paid mutator transaction binding the contract method 0x49262a73.
//
// Solidity: function callFunction(bytes32 functionId, bytes arguments) returns(bytes32)
func (_FaaS *FaaSSession) CallFunction(functionId [32]byte, arguments []byte) (*types.Transaction, error) {
	return _FaaS.Contract.CallFunction(&_FaaS.TransactOpts, functionId, arguments)
}

// CallFunction is a paid mutator transaction binding the contract method 0x49262a73.
//
// Solidity: function callFunction(bytes32 functionId, bytes arguments) returns(bytes32)
func (_FaaS *FaaSTransactorSession) CallFunction(functionId [32]byte, arguments []byte) (*types.Transaction, error) {
	return _FaaS.Contract.CallFunction(&_FaaS.TransactOpts, functionId, arguments)
}

// DeployFunction is a paid mutator transaction binding the contract method 0x5b6116f0.
//
// Solidity: function deployFunction(string url, bytes32 digest) returns(bytes32)
func (_FaaS *FaaSTransactor) DeployFunction(opts *bind.TransactOpts, url string, digest [32]byte) (*types.Transaction, error) {
	return _FaaS.contract.Transact(opts, "deployFunction", url, digest)
}

// DeployFunction is a paid mutator transaction binding the contract method 0x5b6116f0.
//
// Solidity: function deployFunction(string url, bytes32 digest) returns(bytes32)
func (_FaaS *FaaSSession) DeployFunction(url string, digest [32]byte) (*types.Transaction, error) {
	return _FaaS.Contract.DeployFunction(&_FaaS.TransactOpts, url, digest)
}

// DeployFunction is a paid mutator transaction binding the contract method 0x5b6116f0.
//
// Solidity: function deployFunction(string url, bytes32 digest) returns(bytes32)
func (_FaaS *FaaSTransactorSession) DeployFunction(url string, digest [32]byte) (*types.Transaction, error) {
	return _FaaS.Contract.DeployFunction(&_FaaS.TransactOpts, url, digest)
}

// RegisterFunction is a paid mutator transaction binding the contract method 0xe77ebc8f.
//
// Solidity: function registerFunction(bytes content) returns(bytes32)
func (_FaaS *FaaSTransactor) RegisterFunction(opts *bind.TransactOpts, content []byte) (*types.Transaction, error) {
	return _FaaS.contract.Transact(opts, "registerFunction", content)
}

// RegisterFunction is a paid mutator transaction binding the contract method 0xe77ebc8f.
//
// Solidity: function registerFunction(bytes content) returns(bytes32)
func (_FaaS *FaaSSession) RegisterFunction(content []byte) (*types.Transaction, error) {
	return _FaaS.Contract.RegisterFunction(&_FaaS.TransactOpts, content)
}

// RegisterFunction is a paid mutator transaction binding the contract method 0xe77ebc8f.
//
// Solidity: function registerFunction(bytes content) returns(bytes32)
func (_FaaS *FaaSTransactorSession) RegisterFunction(content []byte) (*types.Transaction, error) {
	return _FaaS.Contract.RegisterFunction(&_FaaS.TransactOpts, content)
}

// FaaSFunctionCalledIterator is returned from FilterFunctionCalled and is used to iterate over the raw logs and unpacked data for FunctionCalled events raised by the FaaS contract.
type FaaSFunctionCalledIterator struct {
	Event *FaaSFunctionCalled // Event containing the contract specifics and raw log

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
func (it *FaaSFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaaSFunctionCalled)
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
		it.Event = new(FaaSFunctionCalled)
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
func (it *FaaSFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaaSFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaaSFunctionCalled represents a FunctionCalled event raised by the FaaS contract.
type FaaSFunctionCalled struct {
	FunctionId [32]byte
	TaskId     [32]byte
	Caller     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFunctionCalled is a free log retrieval operation binding the contract event 0x602ae46fd037d4c5652ff0c0e3ac0f0ff6744765ff3f2f8d9fe77f112d3e0985.
//
// Solidity: event FunctionCalled(bytes32 indexed functionId, bytes32 indexed taskId, address indexed caller)
func (_FaaS *FaaSFilterer) FilterFunctionCalled(opts *bind.FilterOpts, functionId [][32]byte, taskId [][32]byte, caller []common.Address) (*FaaSFunctionCalledIterator, error) {

	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _FaaS.contract.FilterLogs(opts, "FunctionCalled", functionIdRule, taskIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &FaaSFunctionCalledIterator{contract: _FaaS.contract, event: "FunctionCalled", logs: logs, sub: sub}, nil
}

// WatchFunctionCalled is a free log subscription operation binding the contract event 0x602ae46fd037d4c5652ff0c0e3ac0f0ff6744765ff3f2f8d9fe77f112d3e0985.
//
// Solidity: event FunctionCalled(bytes32 indexed functionId, bytes32 indexed taskId, address indexed caller)
func (_FaaS *FaaSFilterer) WatchFunctionCalled(opts *bind.WatchOpts, sink chan<- *FaaSFunctionCalled, functionId [][32]byte, taskId [][32]byte, caller []common.Address) (event.Subscription, error) {

	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}
	var taskIdRule []interface{}
	for _, taskIdItem := range taskId {
		taskIdRule = append(taskIdRule, taskIdItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _FaaS.contract.WatchLogs(opts, "FunctionCalled", functionIdRule, taskIdRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaaSFunctionCalled)
				if err := _FaaS.contract.UnpackLog(event, "FunctionCalled", log); err != nil {
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

// ParseFunctionCalled is a log parse operation binding the contract event 0x602ae46fd037d4c5652ff0c0e3ac0f0ff6744765ff3f2f8d9fe77f112d3e0985.
//
// Solidity: event FunctionCalled(bytes32 indexed functionId, bytes32 indexed taskId, address indexed caller)
func (_FaaS *FaaSFilterer) ParseFunctionCalled(log types.Log) (*FaaSFunctionCalled, error) {
	event := new(FaaSFunctionCalled)
	if err := _FaaS.contract.UnpackLog(event, "FunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FaaSFunctionDeployedIterator is returned from FilterFunctionDeployed and is used to iterate over the raw logs and unpacked data for FunctionDeployed events raised by the FaaS contract.
type FaaSFunctionDeployedIterator struct {
	Event *FaaSFunctionDeployed // Event containing the contract specifics and raw log

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
func (it *FaaSFunctionDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaaSFunctionDeployed)
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
		it.Event = new(FaaSFunctionDeployed)
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
func (it *FaaSFunctionDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaaSFunctionDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaaSFunctionDeployed represents a FunctionDeployed event raised by the FaaS contract.
type FaaSFunctionDeployed struct {
	FunctionId [32]byte
	Url        common.Hash
	Digest     [32]byte
	Registrar  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFunctionDeployed is a free log retrieval operation binding the contract event 0x088ac90e03ead29bcc2c251b1567dc223578786d0600fe99db7bc1d8f1497961.
//
// Solidity: event FunctionDeployed(bytes32 indexed functionId, string indexed url, bytes32 indexed digest, address registrar)
func (_FaaS *FaaSFilterer) FilterFunctionDeployed(opts *bind.FilterOpts, functionId [][32]byte, url []string, digest [][32]byte) (*FaaSFunctionDeployedIterator, error) {

	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}
	var urlRule []interface{}
	for _, urlItem := range url {
		urlRule = append(urlRule, urlItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _FaaS.contract.FilterLogs(opts, "FunctionDeployed", functionIdRule, urlRule, digestRule)
	if err != nil {
		return nil, err
	}
	return &FaaSFunctionDeployedIterator{contract: _FaaS.contract, event: "FunctionDeployed", logs: logs, sub: sub}, nil
}

// WatchFunctionDeployed is a free log subscription operation binding the contract event 0x088ac90e03ead29bcc2c251b1567dc223578786d0600fe99db7bc1d8f1497961.
//
// Solidity: event FunctionDeployed(bytes32 indexed functionId, string indexed url, bytes32 indexed digest, address registrar)
func (_FaaS *FaaSFilterer) WatchFunctionDeployed(opts *bind.WatchOpts, sink chan<- *FaaSFunctionDeployed, functionId [][32]byte, url []string, digest [][32]byte) (event.Subscription, error) {

	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}
	var urlRule []interface{}
	for _, urlItem := range url {
		urlRule = append(urlRule, urlItem)
	}
	var digestRule []interface{}
	for _, digestItem := range digest {
		digestRule = append(digestRule, digestItem)
	}

	logs, sub, err := _FaaS.contract.WatchLogs(opts, "FunctionDeployed", functionIdRule, urlRule, digestRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaaSFunctionDeployed)
				if err := _FaaS.contract.UnpackLog(event, "FunctionDeployed", log); err != nil {
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

// ParseFunctionDeployed is a log parse operation binding the contract event 0x088ac90e03ead29bcc2c251b1567dc223578786d0600fe99db7bc1d8f1497961.
//
// Solidity: event FunctionDeployed(bytes32 indexed functionId, string indexed url, bytes32 indexed digest, address registrar)
func (_FaaS *FaaSFilterer) ParseFunctionDeployed(log types.Log) (*FaaSFunctionDeployed, error) {
	event := new(FaaSFunctionDeployed)
	if err := _FaaS.contract.UnpackLog(event, "FunctionDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FaaSFunctionRegisteredIterator is returned from FilterFunctionRegistered and is used to iterate over the raw logs and unpacked data for FunctionRegistered events raised by the FaaS contract.
type FaaSFunctionRegisteredIterator struct {
	Event *FaaSFunctionRegistered // Event containing the contract specifics and raw log

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
func (it *FaaSFunctionRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FaaSFunctionRegistered)
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
		it.Event = new(FaaSFunctionRegistered)
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
func (it *FaaSFunctionRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FaaSFunctionRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FaaSFunctionRegistered represents a FunctionRegistered event raised by the FaaS contract.
type FaaSFunctionRegistered struct {
	FunctionId [32]byte
	Registrar  common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFunctionRegistered is a free log retrieval operation binding the contract event 0xbd508acd226898e9e1b1b8b1f8e4197b63a90c6f2266a22233a966d6f6970bb1.
//
// Solidity: event FunctionRegistered(bytes32 indexed functionId, address indexed registrar)
func (_FaaS *FaaSFilterer) FilterFunctionRegistered(opts *bind.FilterOpts, functionId [][32]byte, registrar []common.Address) (*FaaSFunctionRegisteredIterator, error) {

	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _FaaS.contract.FilterLogs(opts, "FunctionRegistered", functionIdRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return &FaaSFunctionRegisteredIterator{contract: _FaaS.contract, event: "FunctionRegistered", logs: logs, sub: sub}, nil
}

// WatchFunctionRegistered is a free log subscription operation binding the contract event 0xbd508acd226898e9e1b1b8b1f8e4197b63a90c6f2266a22233a966d6f6970bb1.
//
// Solidity: event FunctionRegistered(bytes32 indexed functionId, address indexed registrar)
func (_FaaS *FaaSFilterer) WatchFunctionRegistered(opts *bind.WatchOpts, sink chan<- *FaaSFunctionRegistered, functionId [][32]byte, registrar []common.Address) (event.Subscription, error) {

	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}
	var registrarRule []interface{}
	for _, registrarItem := range registrar {
		registrarRule = append(registrarRule, registrarItem)
	}

	logs, sub, err := _FaaS.contract.WatchLogs(opts, "FunctionRegistered", functionIdRule, registrarRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FaaSFunctionRegistered)
				if err := _FaaS.contract.UnpackLog(event, "FunctionRegistered", log); err != nil {
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

// ParseFunctionRegistered is a log parse operation binding the contract event 0xbd508acd226898e9e1b1b8b1f8e4197b63a90c6f2266a22233a966d6f6970bb1.
//
// Solidity: event FunctionRegistered(bytes32 indexed functionId, address indexed registrar)
func (_FaaS *FaaSFilterer) ParseFunctionRegistered(log types.Log) (*FaaSFunctionRegistered, error) {
	event := new(FaaSFunctionRegistered)
	if err := _FaaS.contract.UnpackLog(event, "FunctionRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
