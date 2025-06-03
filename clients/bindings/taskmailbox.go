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

// TaskMailboxMetaData contains all meta data concerning the TaskMailbox contract.
var TaskMailboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"createTask\",\"inputs\":[{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getTask\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"payload\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enum TaskStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskResult\",\"inputs\":[{\"name\":\"taskId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"}]",
}

// TaskMailboxABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskMailboxMetaData.ABI instead.
var TaskMailboxABI = TaskMailboxMetaData.ABI

// TaskMailbox is an auto generated Go binding around an Ethereum contract.
type TaskMailbox struct {
	TaskMailboxCaller     // Read-only binding to the contract
	TaskMailboxTransactor // Write-only binding to the contract
	TaskMailboxFilterer   // Log filterer for contract events
}

// TaskMailboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskMailboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskMailboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskMailboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskMailboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskMailboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskMailboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskMailboxSession struct {
	Contract     *TaskMailbox      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskMailboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskMailboxCallerSession struct {
	Contract *TaskMailboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TaskMailboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskMailboxTransactorSession struct {
	Contract     *TaskMailboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TaskMailboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskMailboxRaw struct {
	Contract *TaskMailbox // Generic contract binding to access the raw methods on
}

// TaskMailboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskMailboxCallerRaw struct {
	Contract *TaskMailboxCaller // Generic read-only contract binding to access the raw methods on
}

// TaskMailboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskMailboxTransactorRaw struct {
	Contract *TaskMailboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskMailbox creates a new instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailbox(address common.Address, backend bind.ContractBackend) (*TaskMailbox, error) {
	contract, err := bindTaskMailbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskMailbox{TaskMailboxCaller: TaskMailboxCaller{contract: contract}, TaskMailboxTransactor: TaskMailboxTransactor{contract: contract}, TaskMailboxFilterer: TaskMailboxFilterer{contract: contract}}, nil
}

// NewTaskMailboxCaller creates a new read-only instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailboxCaller(address common.Address, caller bind.ContractCaller) (*TaskMailboxCaller, error) {
	contract, err := bindTaskMailbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxCaller{contract: contract}, nil
}

// NewTaskMailboxTransactor creates a new write-only instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailboxTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskMailboxTransactor, error) {
	contract, err := bindTaskMailbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxTransactor{contract: contract}, nil
}

// NewTaskMailboxFilterer creates a new log filterer instance of TaskMailbox, bound to a specific deployed contract.
func NewTaskMailboxFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskMailboxFilterer, error) {
	contract, err := bindTaskMailbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskMailboxFilterer{contract: contract}, nil
}

// bindTaskMailbox binds a generic wrapper to an already deployed contract.
func bindTaskMailbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TaskMailboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskMailbox *TaskMailboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskMailbox.Contract.TaskMailboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// the fallback function (if any).
func (_TaskMailbox *TaskMailboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskMailbox.Contract.TaskMailboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskMailbox *TaskMailboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskMailbox.Contract.TaskMailboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskMailbox *TaskMailboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskMailbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// the fallback function (if any).
func (_TaskMailbox *TaskMailboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskMailbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskMailbox *TaskMailboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskMailbox.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function getTask(bytes32 taskId) view returns(bytes payload, bytes result, uint8 status)
func (_TaskMailbox *TaskMailboxCaller) GetTask(opts *bind.CallOpts, taskId [32]byte) (struct {
	Payload []byte
	Result  []byte
	Status  uint8
}, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getTask", taskId)

	outstruct := new(struct {
		Payload []byte
		Result  []byte
		Status  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Payload = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Result = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetTask is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function getTask(bytes32 taskId) view returns(bytes payload, bytes result, uint8 status)
func (_TaskMailbox *TaskMailboxSession) GetTask(taskId [32]byte) (struct {
	Payload []byte
	Result  []byte
	Status  uint8
}, error) {
	return _TaskMailbox.Contract.GetTask(&_TaskMailbox.CallOpts, taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function getTask(bytes32 taskId) view returns(bytes payload, bytes result, uint8 status)
func (_TaskMailbox *TaskMailboxCallerSession) GetTask(taskId [32]byte) (struct {
	Payload []byte
	Result  []byte
	Status  uint8
}, error) {
	return _TaskMailbox.Contract.GetTask(&_TaskMailbox.CallOpts, taskId)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x865ca1ca.
//
// Solidity: function getTaskResult(bytes32 taskId) view returns(bytes)
func (_TaskMailbox *TaskMailboxCaller) GetTaskResult(opts *bind.CallOpts, taskId [32]byte) ([]byte, error) {
	var out []interface{}
	err := _TaskMailbox.contract.Call(opts, &out, "getTaskResult", taskId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetTaskResult is a free data retrieval call binding the contract method 0x865ca1ca.
//
// Solidity: function getTaskResult(bytes32 taskId) view returns(bytes)
func (_TaskMailbox *TaskMailboxSession) GetTaskResult(taskId [32]byte) ([]byte, error) {
	return _TaskMailbox.Contract.GetTaskResult(&_TaskMailbox.CallOpts, taskId)
}

// GetTaskResult is a free data retrieval call binding the contract method 0x865ca1ca.
//
// Solidity: function getTaskResult(bytes32 taskId) view returns(bytes)
func (_TaskMailbox *TaskMailboxCallerSession) GetTaskResult(taskId [32]byte) ([]byte, error) {
	return _TaskMailbox.Contract.GetTaskResult(&_TaskMailbox.CallOpts, taskId)
}

// CreateTask is a paid mutator transaction binding the contract method 0x9cec7b5f.
//
// Solidity: function createTask(bytes payload) returns(bytes32)
func (_TaskMailbox *TaskMailboxTransactor) CreateTask(opts *bind.TransactOpts, payload []byte) (*types.Transaction, error) {
	return _TaskMailbox.contract.Transact(opts, "createTask", payload)
}

// CreateTask is a paid mutator transaction binding the contract method 0x9cec7b5f.
//
// Solidity: function createTask(bytes payload) returns(bytes32)
func (_TaskMailbox *TaskMailboxSession) CreateTask(payload []byte) (*types.Transaction, error) {
	return _TaskMailbox.Contract.CreateTask(&_TaskMailbox.TransactOpts, payload)
}

// CreateTask is a paid mutator transaction binding the contract method 0x9cec7b5f.
//
// Solidity: function createTask(bytes payload) returns(bytes32)
func (_TaskMailbox *TaskMailboxTransactorSession) CreateTask(payload []byte) (*types.Transaction, error) {
	return _TaskMailbox.Contract.CreateTask(&_TaskMailbox.TransactOpts, payload)
}