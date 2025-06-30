// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AVSRegistrarStorage

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

// AVSRegistrarStorageMetaData contains all meta data concerning the AVSRegistrarStorage contract.
var AVSRegistrarStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"KeyNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllocationManager\",\"inputs\":[]}]",
}

// AVSRegistrarStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSRegistrarStorageMetaData.ABI instead.
var AVSRegistrarStorageABI = AVSRegistrarStorageMetaData.ABI

// AVSRegistrarStorage is an auto generated Go binding around an Ethereum contract.
type AVSRegistrarStorage struct {
	AVSRegistrarStorageCaller     // Read-only binding to the contract
	AVSRegistrarStorageTransactor // Write-only binding to the contract
	AVSRegistrarStorageFilterer   // Log filterer for contract events
}

// AVSRegistrarStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSRegistrarStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSRegistrarStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSRegistrarStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSRegistrarStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSRegistrarStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSRegistrarStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSRegistrarStorageSession struct {
	Contract     *AVSRegistrarStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AVSRegistrarStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSRegistrarStorageCallerSession struct {
	Contract *AVSRegistrarStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AVSRegistrarStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSRegistrarStorageTransactorSession struct {
	Contract     *AVSRegistrarStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AVSRegistrarStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSRegistrarStorageRaw struct {
	Contract *AVSRegistrarStorage // Generic contract binding to access the raw methods on
}

// AVSRegistrarStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSRegistrarStorageCallerRaw struct {
	Contract *AVSRegistrarStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AVSRegistrarStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSRegistrarStorageTransactorRaw struct {
	Contract *AVSRegistrarStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSRegistrarStorage creates a new instance of AVSRegistrarStorage, bound to a specific deployed contract.
func NewAVSRegistrarStorage(address common.Address, backend bind.ContractBackend) (*AVSRegistrarStorage, error) {
	contract, err := bindAVSRegistrarStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarStorage{AVSRegistrarStorageCaller: AVSRegistrarStorageCaller{contract: contract}, AVSRegistrarStorageTransactor: AVSRegistrarStorageTransactor{contract: contract}, AVSRegistrarStorageFilterer: AVSRegistrarStorageFilterer{contract: contract}}, nil
}

// NewAVSRegistrarStorageCaller creates a new read-only instance of AVSRegistrarStorage, bound to a specific deployed contract.
func NewAVSRegistrarStorageCaller(address common.Address, caller bind.ContractCaller) (*AVSRegistrarStorageCaller, error) {
	contract, err := bindAVSRegistrarStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarStorageCaller{contract: contract}, nil
}

// NewAVSRegistrarStorageTransactor creates a new write-only instance of AVSRegistrarStorage, bound to a specific deployed contract.
func NewAVSRegistrarStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSRegistrarStorageTransactor, error) {
	contract, err := bindAVSRegistrarStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarStorageTransactor{contract: contract}, nil
}

// NewAVSRegistrarStorageFilterer creates a new log filterer instance of AVSRegistrarStorage, bound to a specific deployed contract.
func NewAVSRegistrarStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSRegistrarStorageFilterer, error) {
	contract, err := bindAVSRegistrarStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarStorageFilterer{contract: contract}, nil
}

// bindAVSRegistrarStorage binds a generic wrapper to an already deployed contract.
func bindAVSRegistrarStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSRegistrarStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSRegistrarStorage *AVSRegistrarStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSRegistrarStorage.Contract.AVSRegistrarStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSRegistrarStorage *AVSRegistrarStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.AVSRegistrarStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSRegistrarStorage *AVSRegistrarStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.AVSRegistrarStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSRegistrarStorage *AVSRegistrarStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSRegistrarStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSRegistrarStorage *AVSRegistrarStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSRegistrarStorage *AVSRegistrarStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSRegistrarStorage.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageSession) AllocationManager() (common.Address, error) {
	return _AVSRegistrarStorage.Contract.AllocationManager(&_AVSRegistrarStorage.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageCallerSession) AllocationManager() (common.Address, error) {
	return _AVSRegistrarStorage.Contract.AllocationManager(&_AVSRegistrarStorage.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageCaller) Avs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSRegistrarStorage.contract.Call(opts, &out, "avs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageSession) Avs() (common.Address, error) {
	return _AVSRegistrarStorage.Contract.Avs(&_AVSRegistrarStorage.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageCallerSession) Avs() (common.Address, error) {
	return _AVSRegistrarStorage.Contract.Avs(&_AVSRegistrarStorage.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSRegistrarStorage.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageSession) KeyRegistrar() (common.Address, error) {
	return _AVSRegistrarStorage.Contract.KeyRegistrar(&_AVSRegistrarStorage.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_AVSRegistrarStorage *AVSRegistrarStorageCallerSession) KeyRegistrar() (common.Address, error) {
	return _AVSRegistrarStorage.Contract.KeyRegistrar(&_AVSRegistrarStorage.CallOpts)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_AVSRegistrarStorage *AVSRegistrarStorageCaller) SupportsAVS(opts *bind.CallOpts, avs common.Address) (bool, error) {
	var out []interface{}
	err := _AVSRegistrarStorage.contract.Call(opts, &out, "supportsAVS", avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_AVSRegistrarStorage *AVSRegistrarStorageSession) SupportsAVS(avs common.Address) (bool, error) {
	return _AVSRegistrarStorage.Contract.SupportsAVS(&_AVSRegistrarStorage.CallOpts, avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address avs) view returns(bool)
func (_AVSRegistrarStorage *AVSRegistrarStorageCallerSession) SupportsAVS(avs common.Address) (bool, error) {
	return _AVSRegistrarStorage.Contract.SupportsAVS(&_AVSRegistrarStorage.CallOpts, avs)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_AVSRegistrarStorage *AVSRegistrarStorageTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSRegistrarStorage.contract.Transact(opts, "deregisterOperator", operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_AVSRegistrarStorage *AVSRegistrarStorageSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.DeregisterOperator(&_AVSRegistrarStorage.TransactOpts, operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_AVSRegistrarStorage *AVSRegistrarStorageTransactorSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.DeregisterOperator(&_AVSRegistrarStorage.TransactOpts, operator, avs, operatorSetIds)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_AVSRegistrarStorage *AVSRegistrarStorageTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _AVSRegistrarStorage.contract.Transact(opts, "registerOperator", operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_AVSRegistrarStorage *AVSRegistrarStorageSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.RegisterOperator(&_AVSRegistrarStorage.TransactOpts, operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_AVSRegistrarStorage *AVSRegistrarStorageTransactorSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _AVSRegistrarStorage.Contract.RegisterOperator(&_AVSRegistrarStorage.TransactOpts, operator, avs, operatorSetIds, data)
}

// AVSRegistrarStorageOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the AVSRegistrarStorage contract.
type AVSRegistrarStorageOperatorDeregisteredIterator struct {
	Event *AVSRegistrarStorageOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *AVSRegistrarStorageOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSRegistrarStorageOperatorDeregistered)
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
		it.Event = new(AVSRegistrarStorageOperatorDeregistered)
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
func (it *AVSRegistrarStorageOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSRegistrarStorageOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSRegistrarStorageOperatorDeregistered represents a OperatorDeregistered event raised by the AVSRegistrarStorage contract.
type AVSRegistrarStorageOperatorDeregistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrarStorage *AVSRegistrarStorageFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*AVSRegistrarStorageOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrarStorage.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarStorageOperatorDeregisteredIterator{contract: _AVSRegistrarStorage.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrarStorage *AVSRegistrarStorageFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *AVSRegistrarStorageOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrarStorage.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSRegistrarStorageOperatorDeregistered)
				if err := _AVSRegistrarStorage.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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
func (_AVSRegistrarStorage *AVSRegistrarStorageFilterer) ParseOperatorDeregistered(log types.Log) (*AVSRegistrarStorageOperatorDeregistered, error) {
	event := new(AVSRegistrarStorageOperatorDeregistered)
	if err := _AVSRegistrarStorage.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSRegistrarStorageOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the AVSRegistrarStorage contract.
type AVSRegistrarStorageOperatorRegisteredIterator struct {
	Event *AVSRegistrarStorageOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *AVSRegistrarStorageOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSRegistrarStorageOperatorRegistered)
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
		it.Event = new(AVSRegistrarStorageOperatorRegistered)
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
func (it *AVSRegistrarStorageOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSRegistrarStorageOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSRegistrarStorageOperatorRegistered represents a OperatorRegistered event raised by the AVSRegistrarStorage contract.
type AVSRegistrarStorageOperatorRegistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrarStorage *AVSRegistrarStorageFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*AVSRegistrarStorageOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrarStorage.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarStorageOperatorRegisteredIterator{contract: _AVSRegistrarStorage.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrarStorage *AVSRegistrarStorageFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *AVSRegistrarStorageOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrarStorage.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSRegistrarStorageOperatorRegistered)
				if err := _AVSRegistrarStorage.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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
func (_AVSRegistrarStorage *AVSRegistrarStorageFilterer) ParseOperatorRegistered(log types.Log) (*AVSRegistrarStorageOperatorRegistered, error) {
	event := new(AVSRegistrarStorageOperatorRegistered)
	if err := _AVSRegistrarStorage.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
