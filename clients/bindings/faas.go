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
)

// FaaSMetaData contains all meta data concerning the FaaS contract.
var FaaSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskMailbox\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callFunction\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"functions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunction\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerFunction\",\"inputs\":[{\"name\":\"content\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskMailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contract ITaskMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FunctionCalled\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"taskId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FunctionRegistered\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
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
	parsed, err := abi.JSON(strings.NewReader(FaaSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FaaS *FaaSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FaaS.Contract.FaaSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// the fallback function (if any).
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
// the fallback function (if any).
func (_FaaS *FaaSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FaaS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FaaS *FaaSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FaaS.Contract.contract.Transact(opts, method, params...)
}

// Functions is a free data retrieval call binding the contract method 0x0e8fbb5a.
//
// Solidity: function functions(bytes32 ) view returns(bytes)
func (_FaaS *FaaSCaller) Functions(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "functions", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Functions is a free data retrieval call binding the contract method 0x0e8fbb5a.
//
// Solidity: function functions(bytes32 ) view returns(bytes)
func (_FaaS *FaaSSession) Functions(arg0 [32]byte) ([]byte, error) {
	return _FaaS.Contract.Functions(&_FaaS.CallOpts, arg0)
}

// Functions is a free data retrieval call binding the contract method 0x0e8fbb5a.
//
// Solidity: function functions(bytes32 ) view returns(bytes)
func (_FaaS *FaaSCallerSession) Functions(arg0 [32]byte) ([]byte, error) {
	return _FaaS.Contract.Functions(&_FaaS.CallOpts, arg0)
}

// GetFunction is a free data retrieval call binding the contract method 0x15834a73.
//
// Solidity: function getFunction(bytes32 functionId) view returns(bytes)
func (_FaaS *FaaSCaller) GetFunction(opts *bind.CallOpts, functionId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "getFunction", functionId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetFunction is a free data retrieval call binding the contract method 0x15834a73.
//
// Solidity: function getFunction(bytes32 functionId) view returns(bytes)
func (_FaaS *FaaSSession) GetFunction(functionId [32]byte) ([]byte, error) {
	return _FaaS.Contract.GetFunction(&_FaaS.CallOpts, functionId)
}

// GetFunction is a free data retrieval call binding the contract method 0x15834a73.
//
// Solidity: function getFunction(bytes32 functionId) view returns(bytes)
func (_FaaS *FaaSCallerSession) GetFunction(functionId [32]byte) ([]byte, error) {
	return _FaaS.Contract.GetFunction(&_FaaS.CallOpts, functionId)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf83d08ba.
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

// TaskMailbox is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function taskMailbox() view returns(address)
func (_FaaS *FaaSSession) TaskMailbox() (common.Address, error) {
	return _FaaS.Contract.TaskMailbox(&_FaaS.CallOpts)
}

// TaskMailbox is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function taskMailbox() view returns(address)
func (_FaaS *FaaSCallerSession) TaskMailbox() (common.Address, error) {
	return _FaaS.Contract.TaskMailbox(&_FaaS.CallOpts)
}

// CallFunction is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function callFunction(bytes32 functionId, bytes arguments) returns(bytes32)
func (_FaaS *FaaSTransactor) CallFunction(opts *bind.TransactOpts, functionId [32]byte, arguments []byte) (*types.Transaction, error) {
	return _FaaS.contract.Transact(opts, "callFunction", functionId, arguments)
}

// CallFunction is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function callFunction(bytes32 functionId, bytes arguments) returns(bytes32)
func (_FaaS *FaaSSession) CallFunction(functionId [32]byte, arguments []byte) (*types.Transaction, error) {
	return _FaaS.Contract.CallFunction(&_FaaS.TransactOpts, functionId, arguments)
}

// CallFunction is a paid mutator transaction binding the contract method 0xa5843f08.
//
// Solidity: function callFunction(bytes32 functionId, bytes arguments) returns(bytes32)
func (_FaaS *FaaSTransactorSession) CallFunction(functionId [32]byte, arguments []byte) (*types.Transaction, error) {
	return _FaaS.Contract.CallFunction(&_FaaS.TransactOpts, functionId, arguments)
}

// RegisterFunction is a paid mutator transaction binding the contract method 0x30ec90b3.
//
// Solidity: function registerFunction(bytes content) returns(bytes32)
func (_FaaS *FaaSTransactor) RegisterFunction(opts *bind.TransactOpts, content []byte) (*types.Transaction, error) {
	return _FaaS.contract.Transact(opts, "registerFunction", content)
}

// RegisterFunction is a paid mutator transaction binding the contract method 0x30ec90b3.
//
// Solidity: function registerFunction(bytes content) returns(bytes32)
func (_FaaS *FaaSSession) RegisterFunction(content []byte) (*types.Transaction, error) {
	return _FaaS.Contract.RegisterFunction(&_FaaS.TransactOpts, content)
}

// RegisterFunction is a paid mutator transaction binding the contract method 0x30ec90b3.
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

// FilterFunctionCalled is a free log retrieval operation binding the contract event 0x7dc0d1e5f47b0d6e7d2e55c5b6ea29e59e3cf50e9d4e5ed1b4bb3c5d1e6f7a3b.
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

// WatchFunctionCalled is a free log subscription operation binding the contract event 0x7dc0d1e5f47b0d6e7d2e55c5b6ea29e59e3cf50e9d4e5ed1b4bb3c5d1e6f7a3b.
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

// ParseFunctionCalled is a log parse operation binding the contract event 0x7dc0d1e5f47b0d6e7d2e55c5b6ea29e59e3cf50e9d4e5ed1b4bb3c5d1e6f7a3b.
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

// FilterFunctionRegistered is a free log retrieval operation binding the contract event 0x6cc5fd28f4b9e8a1e0dcb5a05f1b8b9f3c3e5f4f2e6e3c8b7d4f8e9f6b3c2d1a.
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

// WatchFunctionRegistered is a free log subscription operation binding the contract event 0x6cc5fd28f4b9e8a1e0dcb5a05f1b8b9f3c3e5f4f2e6e3c8b7d4f8e9f6b3c2d1a.
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

// ParseFunctionRegistered is a log parse operation binding the contract event 0x6cc5fd28f4b9e8a1e0dcb5a05f1b8b9f3c3e5f4f2e6e3c8b7d4f8e9f6b3c2d1a.
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