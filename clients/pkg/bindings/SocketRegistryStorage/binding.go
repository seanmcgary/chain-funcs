// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SocketRegistryStorage

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

// SocketRegistryStorageMetaData contains all meta data concerning the SocketRegistryStorage contract.
var SocketRegistryStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getOperatorSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OperatorSocketSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerNotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DataLengthMismatch\",\"inputs\":[]}]",
}

// SocketRegistryStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SocketRegistryStorageMetaData.ABI instead.
var SocketRegistryStorageABI = SocketRegistryStorageMetaData.ABI

// SocketRegistryStorage is an auto generated Go binding around an Ethereum contract.
type SocketRegistryStorage struct {
	SocketRegistryStorageCaller     // Read-only binding to the contract
	SocketRegistryStorageTransactor // Write-only binding to the contract
	SocketRegistryStorageFilterer   // Log filterer for contract events
}

// SocketRegistryStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SocketRegistryStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketRegistryStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SocketRegistryStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketRegistryStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SocketRegistryStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SocketRegistryStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SocketRegistryStorageSession struct {
	Contract     *SocketRegistryStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SocketRegistryStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SocketRegistryStorageCallerSession struct {
	Contract *SocketRegistryStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SocketRegistryStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SocketRegistryStorageTransactorSession struct {
	Contract     *SocketRegistryStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SocketRegistryStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SocketRegistryStorageRaw struct {
	Contract *SocketRegistryStorage // Generic contract binding to access the raw methods on
}

// SocketRegistryStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SocketRegistryStorageCallerRaw struct {
	Contract *SocketRegistryStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SocketRegistryStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SocketRegistryStorageTransactorRaw struct {
	Contract *SocketRegistryStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSocketRegistryStorage creates a new instance of SocketRegistryStorage, bound to a specific deployed contract.
func NewSocketRegistryStorage(address common.Address, backend bind.ContractBackend) (*SocketRegistryStorage, error) {
	contract, err := bindSocketRegistryStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryStorage{SocketRegistryStorageCaller: SocketRegistryStorageCaller{contract: contract}, SocketRegistryStorageTransactor: SocketRegistryStorageTransactor{contract: contract}, SocketRegistryStorageFilterer: SocketRegistryStorageFilterer{contract: contract}}, nil
}

// NewSocketRegistryStorageCaller creates a new read-only instance of SocketRegistryStorage, bound to a specific deployed contract.
func NewSocketRegistryStorageCaller(address common.Address, caller bind.ContractCaller) (*SocketRegistryStorageCaller, error) {
	contract, err := bindSocketRegistryStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryStorageCaller{contract: contract}, nil
}

// NewSocketRegistryStorageTransactor creates a new write-only instance of SocketRegistryStorage, bound to a specific deployed contract.
func NewSocketRegistryStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SocketRegistryStorageTransactor, error) {
	contract, err := bindSocketRegistryStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryStorageTransactor{contract: contract}, nil
}

// NewSocketRegistryStorageFilterer creates a new log filterer instance of SocketRegistryStorage, bound to a specific deployed contract.
func NewSocketRegistryStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SocketRegistryStorageFilterer, error) {
	contract, err := bindSocketRegistryStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryStorageFilterer{contract: contract}, nil
}

// bindSocketRegistryStorage binds a generic wrapper to an already deployed contract.
func bindSocketRegistryStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SocketRegistryStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocketRegistryStorage *SocketRegistryStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocketRegistryStorage.Contract.SocketRegistryStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocketRegistryStorage *SocketRegistryStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocketRegistryStorage.Contract.SocketRegistryStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocketRegistryStorage *SocketRegistryStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocketRegistryStorage.Contract.SocketRegistryStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SocketRegistryStorage *SocketRegistryStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SocketRegistryStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SocketRegistryStorage *SocketRegistryStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SocketRegistryStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SocketRegistryStorage *SocketRegistryStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SocketRegistryStorage.Contract.contract.Transact(opts, method, params...)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_SocketRegistryStorage *SocketRegistryStorageCaller) GetOperatorSocket(opts *bind.CallOpts, operator common.Address) (string, error) {
	var out []interface{}
	err := _SocketRegistryStorage.contract.Call(opts, &out, "getOperatorSocket", operator)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_SocketRegistryStorage *SocketRegistryStorageSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _SocketRegistryStorage.Contract.GetOperatorSocket(&_SocketRegistryStorage.CallOpts, operator)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_SocketRegistryStorage *SocketRegistryStorageCallerSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _SocketRegistryStorage.Contract.GetOperatorSocket(&_SocketRegistryStorage.CallOpts, operator)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_SocketRegistryStorage *SocketRegistryStorageTransactor) UpdateSocket(opts *bind.TransactOpts, operator common.Address, socket string) (*types.Transaction, error) {
	return _SocketRegistryStorage.contract.Transact(opts, "updateSocket", operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_SocketRegistryStorage *SocketRegistryStorageSession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _SocketRegistryStorage.Contract.UpdateSocket(&_SocketRegistryStorage.TransactOpts, operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_SocketRegistryStorage *SocketRegistryStorageTransactorSession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _SocketRegistryStorage.Contract.UpdateSocket(&_SocketRegistryStorage.TransactOpts, operator, socket)
}

// SocketRegistryStorageOperatorSocketSetIterator is returned from FilterOperatorSocketSet and is used to iterate over the raw logs and unpacked data for OperatorSocketSet events raised by the SocketRegistryStorage contract.
type SocketRegistryStorageOperatorSocketSetIterator struct {
	Event *SocketRegistryStorageOperatorSocketSet // Event containing the contract specifics and raw log

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
func (it *SocketRegistryStorageOperatorSocketSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SocketRegistryStorageOperatorSocketSet)
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
		it.Event = new(SocketRegistryStorageOperatorSocketSet)
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
func (it *SocketRegistryStorageOperatorSocketSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SocketRegistryStorageOperatorSocketSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SocketRegistryStorageOperatorSocketSet represents a OperatorSocketSet event raised by the SocketRegistryStorage contract.
type SocketRegistryStorageOperatorSocketSet struct {
	Operator common.Address
	Socket   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketSet is a free log retrieval operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_SocketRegistryStorage *SocketRegistryStorageFilterer) FilterOperatorSocketSet(opts *bind.FilterOpts, operator []common.Address) (*SocketRegistryStorageOperatorSocketSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SocketRegistryStorage.contract.FilterLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &SocketRegistryStorageOperatorSocketSetIterator{contract: _SocketRegistryStorage.contract, event: "OperatorSocketSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketSet is a free log subscription operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_SocketRegistryStorage *SocketRegistryStorageFilterer) WatchOperatorSocketSet(opts *bind.WatchOpts, sink chan<- *SocketRegistryStorageOperatorSocketSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SocketRegistryStorage.contract.WatchLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SocketRegistryStorageOperatorSocketSet)
				if err := _SocketRegistryStorage.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
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
func (_SocketRegistryStorage *SocketRegistryStorageFilterer) ParseOperatorSocketSet(log types.Log) (*SocketRegistryStorageOperatorSocketSet, error) {
	event := new(SocketRegistryStorageOperatorSocketSet)
	if err := _SocketRegistryStorage.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
