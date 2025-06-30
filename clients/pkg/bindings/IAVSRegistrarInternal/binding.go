// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAVSRegistrarInternal

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

// IAVSRegistrarInternalMetaData contains all meta data concerning the IAVSRegistrarInternal contract.
var IAVSRegistrarInternalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"KeyNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllocationManager\",\"inputs\":[]}]",
}

// IAVSRegistrarInternalABI is the input ABI used to generate the binding from.
// Deprecated: Use IAVSRegistrarInternalMetaData.ABI instead.
var IAVSRegistrarInternalABI = IAVSRegistrarInternalMetaData.ABI

// IAVSRegistrarInternal is an auto generated Go binding around an Ethereum contract.
type IAVSRegistrarInternal struct {
	IAVSRegistrarInternalCaller     // Read-only binding to the contract
	IAVSRegistrarInternalTransactor // Write-only binding to the contract
	IAVSRegistrarInternalFilterer   // Log filterer for contract events
}

// IAVSRegistrarInternalCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAVSRegistrarInternalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSRegistrarInternalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAVSRegistrarInternalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSRegistrarInternalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAVSRegistrarInternalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAVSRegistrarInternalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAVSRegistrarInternalSession struct {
	Contract     *IAVSRegistrarInternal // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IAVSRegistrarInternalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAVSRegistrarInternalCallerSession struct {
	Contract *IAVSRegistrarInternalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// IAVSRegistrarInternalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAVSRegistrarInternalTransactorSession struct {
	Contract     *IAVSRegistrarInternalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// IAVSRegistrarInternalRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAVSRegistrarInternalRaw struct {
	Contract *IAVSRegistrarInternal // Generic contract binding to access the raw methods on
}

// IAVSRegistrarInternalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAVSRegistrarInternalCallerRaw struct {
	Contract *IAVSRegistrarInternalCaller // Generic read-only contract binding to access the raw methods on
}

// IAVSRegistrarInternalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAVSRegistrarInternalTransactorRaw struct {
	Contract *IAVSRegistrarInternalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAVSRegistrarInternal creates a new instance of IAVSRegistrarInternal, bound to a specific deployed contract.
func NewIAVSRegistrarInternal(address common.Address, backend bind.ContractBackend) (*IAVSRegistrarInternal, error) {
	contract, err := bindIAVSRegistrarInternal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarInternal{IAVSRegistrarInternalCaller: IAVSRegistrarInternalCaller{contract: contract}, IAVSRegistrarInternalTransactor: IAVSRegistrarInternalTransactor{contract: contract}, IAVSRegistrarInternalFilterer: IAVSRegistrarInternalFilterer{contract: contract}}, nil
}

// NewIAVSRegistrarInternalCaller creates a new read-only instance of IAVSRegistrarInternal, bound to a specific deployed contract.
func NewIAVSRegistrarInternalCaller(address common.Address, caller bind.ContractCaller) (*IAVSRegistrarInternalCaller, error) {
	contract, err := bindIAVSRegistrarInternal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarInternalCaller{contract: contract}, nil
}

// NewIAVSRegistrarInternalTransactor creates a new write-only instance of IAVSRegistrarInternal, bound to a specific deployed contract.
func NewIAVSRegistrarInternalTransactor(address common.Address, transactor bind.ContractTransactor) (*IAVSRegistrarInternalTransactor, error) {
	contract, err := bindIAVSRegistrarInternal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarInternalTransactor{contract: contract}, nil
}

// NewIAVSRegistrarInternalFilterer creates a new log filterer instance of IAVSRegistrarInternal, bound to a specific deployed contract.
func NewIAVSRegistrarInternalFilterer(address common.Address, filterer bind.ContractFilterer) (*IAVSRegistrarInternalFilterer, error) {
	contract, err := bindIAVSRegistrarInternal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarInternalFilterer{contract: contract}, nil
}

// bindIAVSRegistrarInternal binds a generic wrapper to an already deployed contract.
func bindIAVSRegistrarInternal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAVSRegistrarInternalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVSRegistrarInternal *IAVSRegistrarInternalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVSRegistrarInternal.Contract.IAVSRegistrarInternalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVSRegistrarInternal *IAVSRegistrarInternalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVSRegistrarInternal.Contract.IAVSRegistrarInternalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVSRegistrarInternal *IAVSRegistrarInternalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVSRegistrarInternal.Contract.IAVSRegistrarInternalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAVSRegistrarInternal *IAVSRegistrarInternalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAVSRegistrarInternal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAVSRegistrarInternal *IAVSRegistrarInternalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAVSRegistrarInternal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAVSRegistrarInternal *IAVSRegistrarInternalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAVSRegistrarInternal.Contract.contract.Transact(opts, method, params...)
}

// IAVSRegistrarInternalOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the IAVSRegistrarInternal contract.
type IAVSRegistrarInternalOperatorDeregisteredIterator struct {
	Event *IAVSRegistrarInternalOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *IAVSRegistrarInternalOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAVSRegistrarInternalOperatorDeregistered)
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
		it.Event = new(IAVSRegistrarInternalOperatorDeregistered)
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
func (it *IAVSRegistrarInternalOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAVSRegistrarInternalOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAVSRegistrarInternalOperatorDeregistered represents a OperatorDeregistered event raised by the IAVSRegistrarInternal contract.
type IAVSRegistrarInternalOperatorDeregistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_IAVSRegistrarInternal *IAVSRegistrarInternalFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*IAVSRegistrarInternalOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVSRegistrarInternal.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarInternalOperatorDeregisteredIterator{contract: _IAVSRegistrarInternal.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_IAVSRegistrarInternal *IAVSRegistrarInternalFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *IAVSRegistrarInternalOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVSRegistrarInternal.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAVSRegistrarInternalOperatorDeregistered)
				if err := _IAVSRegistrarInternal.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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

// ParseOperatorDeregistered is a log parse operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_IAVSRegistrarInternal *IAVSRegistrarInternalFilterer) ParseOperatorDeregistered(log types.Log) (*IAVSRegistrarInternalOperatorDeregistered, error) {
	event := new(IAVSRegistrarInternalOperatorDeregistered)
	if err := _IAVSRegistrarInternal.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAVSRegistrarInternalOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the IAVSRegistrarInternal contract.
type IAVSRegistrarInternalOperatorRegisteredIterator struct {
	Event *IAVSRegistrarInternalOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *IAVSRegistrarInternalOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAVSRegistrarInternalOperatorRegistered)
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
		it.Event = new(IAVSRegistrarInternalOperatorRegistered)
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
func (it *IAVSRegistrarInternalOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAVSRegistrarInternalOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAVSRegistrarInternalOperatorRegistered represents a OperatorRegistered event raised by the IAVSRegistrarInternal contract.
type IAVSRegistrarInternalOperatorRegistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_IAVSRegistrarInternal *IAVSRegistrarInternalFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*IAVSRegistrarInternalOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVSRegistrarInternal.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &IAVSRegistrarInternalOperatorRegisteredIterator{contract: _IAVSRegistrarInternal.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_IAVSRegistrarInternal *IAVSRegistrarInternalFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *IAVSRegistrarInternalOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _IAVSRegistrarInternal.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAVSRegistrarInternalOperatorRegistered)
				if err := _IAVSRegistrarInternal.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_IAVSRegistrarInternal *IAVSRegistrarInternalFilterer) ParseOperatorRegistered(log types.Log) (*IAVSRegistrarInternalOperatorRegistered, error) {
	event := new(IAVSRegistrarInternalOperatorRegistered)
	if err := _IAVSRegistrarInternal.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
