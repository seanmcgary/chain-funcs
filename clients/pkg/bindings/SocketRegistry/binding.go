// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SocketRegistry

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

// SocketRegistryMetaData contains all meta data concerning the SocketRegistry contract.
var SocketRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getOperatorSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OperatorSocketSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerNotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DataLengthMismatch\",\"inputs\":[]}]",
}

// SocketRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use SocketRegistryMetaData.ABI instead.
var SocketRegistryABI = SocketRegistryMetaData.ABI

// SocketRegistry is an auto generated Go binding around an Ethereum contract.
type SocketRegistry struct {
	SocketRegistryCaller     // Read-only binding to the contract
	SocketRegistryTransactor // Write-only binding to the contract
	SocketRegistryFilterer   // Log filterer for contract events
}

// SocketRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SocketRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SocketRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SocketRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SocketRegistrySession struct {
	Contract     *SocketRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SocketRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SocketRegistryCallerSession struct {
	Contract *SocketRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SocketRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SocketRegistryTransactorSession struct {
	Contract     *SocketRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SocketRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SocketRegistryRaw struct {
	Contract *SocketRegistry // Generic contract binding to access the raw methods on
}

// SocketRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SocketRegistryCallerRaw struct {
	Contract *SocketRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// SocketRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SocketRegistryTransactorRaw struct {
	Contract *SocketRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSocketRegistry creates a new instance of SocketRegistry, bound to a specific deployed contract.
func NewSocketRegistry(address common.Address, backend bind.ContractBackend) (*SocketRegistry, error) {
	contract, err := bindSocketRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SocketRegistry{SocketRegistryCaller: SocketRegistryCaller{contract: contract}, SocketRegistryTransactor: SocketRegistryTransactor{contract: contract}, SocketRegistryFilterer: SocketRegistryFilterer{contract: contract}}, nil
}

// NewSocketRegistryCaller creates a new read-only instance of SocketRegistry, bound to a specific deployed contract.
func NewSocketRegistryCaller(address common.Address, caller bind.ContractCaller) (*SocketRegistryCaller, error) {
	contract, err := bindSocketRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryCaller{contract: contract}, nil
}

// NewSocketRegistryTransactor creates a new write-only instance of SocketRegistry, bound to a specific deployed contract.
func NewSocketRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*SocketRegistryTransactor, error) {
	contract, err := bindSocketRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryTransactor{contract: contract}, nil
}

// NewSocketRegistryFilterer creates a new log filterer instance of SocketRegistry, bound to a specific deployed contract.
func NewSocketRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*SocketRegistryFilterer, error) {
	contract, err := bindSocketRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryFilterer{contract: contract}, nil
}

// bindSocketRegistry binds a generic wrapper to an already deployed contract.
func bindSocketRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SocketRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocketRegistry *SocketRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocketRegistry.Contract.SocketRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocketRegistry *SocketRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocketRegistry.Contract.SocketRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocketRegistry *SocketRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocketRegistry.Contract.SocketRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocketRegistry *SocketRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocketRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocketRegistry *SocketRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocketRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocketRegistry *SocketRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocketRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_SocketRegistry *SocketRegistryCaller) GetOperatorSocket(opts *bind.CallOpts, operator common.Address) (string, error) {
	var out []interface{}
	err := _SocketRegistry.contract.Call(opts, &out, "getOperatorSocket", operator)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_SocketRegistry *SocketRegistrySession) GetOperatorSocket(operator common.Address) (string, error) {
	return _SocketRegistry.Contract.GetOperatorSocket(&_SocketRegistry.CallOpts, operator)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_SocketRegistry *SocketRegistryCallerSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _SocketRegistry.Contract.GetOperatorSocket(&_SocketRegistry.CallOpts, operator)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_SocketRegistry *SocketRegistryTransactor) UpdateSocket(opts *bind.TransactOpts, operator common.Address, socket string) (*types.Transaction, error) {
	return _SocketRegistry.contract.Transact(opts, "updateSocket", operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_SocketRegistry *SocketRegistrySession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _SocketRegistry.Contract.UpdateSocket(&_SocketRegistry.TransactOpts, operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_SocketRegistry *SocketRegistryTransactorSession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _SocketRegistry.Contract.UpdateSocket(&_SocketRegistry.TransactOpts, operator, socket)
}

// SocketRegistryOperatorSocketSetIterator is returned from FilterOperatorSocketSet and is used to iterate over the raw logs and unpacked data for OperatorSocketSet events raised by the SocketRegistry contract.
type SocketRegistryOperatorSocketSetIterator struct {
	Event *SocketRegistryOperatorSocketSet // Event containing the contract specifics and raw log

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
func (it *SocketRegistryOperatorSocketSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocketRegistryOperatorSocketSet)
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
		it.Event = new(SocketRegistryOperatorSocketSet)
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
func (it *SocketRegistryOperatorSocketSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocketRegistryOperatorSocketSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocketRegistryOperatorSocketSet represents a OperatorSocketSet event raised by the SocketRegistry contract.
type SocketRegistryOperatorSocketSet struct {
	Operator common.Address
	Socket   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketSet is a free log retrieval operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_SocketRegistry *SocketRegistryFilterer) FilterOperatorSocketSet(opts *bind.FilterOpts, operator []common.Address) (*SocketRegistryOperatorSocketSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SocketRegistry.contract.FilterLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryOperatorSocketSetIterator{contract: _SocketRegistry.contract, event: "OperatorSocketSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketSet is a free log subscription operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_SocketRegistry *SocketRegistryFilterer) WatchOperatorSocketSet(opts *bind.WatchOpts, sink chan<- *SocketRegistryOperatorSocketSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SocketRegistry.contract.WatchLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocketRegistryOperatorSocketSet)
				if err := _SocketRegistry.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
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

// ParseOperatorSocketSet is a log parse operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_SocketRegistry *SocketRegistryFilterer) ParseOperatorSocketSet(log types.Log) (*SocketRegistryOperatorSocketSet, error) {
	event := new(SocketRegistryOperatorSocketSet)
	if err := _SocketRegistry.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
