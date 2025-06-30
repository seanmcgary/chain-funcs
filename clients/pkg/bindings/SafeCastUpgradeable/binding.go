// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SafeCastUpgradeable

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

// SafeCastUpgradeableMetaData contains all meta data concerning the SafeCastUpgradeable contract.
var SafeCastUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea26469706673582212205b7b0058e12366b82033212bcddc95e36c795495f562d105246d1b169bc1f01564736f6c634300081b0033",
}

// SafeCastUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeCastUpgradeableMetaData.ABI instead.
var SafeCastUpgradeableABI = SafeCastUpgradeableMetaData.ABI

// SafeCastUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeCastUpgradeableMetaData.Bin instead.
var SafeCastUpgradeableBin = SafeCastUpgradeableMetaData.Bin

// DeploySafeCastUpgradeable deploys a new Ethereum contract, binding an instance of SafeCastUpgradeable to it.
func DeploySafeCastUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeCastUpgradeable, error) {
	parsed, err := SafeCastUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeCastUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeCastUpgradeable{SafeCastUpgradeableCaller: SafeCastUpgradeableCaller{contract: contract}, SafeCastUpgradeableTransactor: SafeCastUpgradeableTransactor{contract: contract}, SafeCastUpgradeableFilterer: SafeCastUpgradeableFilterer{contract: contract}}, nil
}

// SafeCastUpgradeable is an auto generated Go binding around an Ethereum contract.
type SafeCastUpgradeable struct {
	SafeCastUpgradeableCaller     // Read-only binding to the contract
	SafeCastUpgradeableTransactor // Write-only binding to the contract
	SafeCastUpgradeableFilterer   // Log filterer for contract events
}

// SafeCastUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeCastUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeCastUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeCastUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeCastUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeCastUpgradeableSession struct {
	Contract     *SafeCastUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeCastUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeCastUpgradeableCallerSession struct {
	Contract *SafeCastUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SafeCastUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeCastUpgradeableTransactorSession struct {
	Contract     *SafeCastUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SafeCastUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeCastUpgradeableRaw struct {
	Contract *SafeCastUpgradeable // Generic contract binding to access the raw methods on
}

// SafeCastUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeCastUpgradeableCallerRaw struct {
	Contract *SafeCastUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// SafeCastUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeCastUpgradeableTransactorRaw struct {
	Contract *SafeCastUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeCastUpgradeable creates a new instance of SafeCastUpgradeable, bound to a specific deployed contract.
func NewSafeCastUpgradeable(address common.Address, backend bind.ContractBackend) (*SafeCastUpgradeable, error) {
	contract, err := bindSafeCastUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeCastUpgradeable{SafeCastUpgradeableCaller: SafeCastUpgradeableCaller{contract: contract}, SafeCastUpgradeableTransactor: SafeCastUpgradeableTransactor{contract: contract}, SafeCastUpgradeableFilterer: SafeCastUpgradeableFilterer{contract: contract}}, nil
}

// NewSafeCastUpgradeableCaller creates a new read-only instance of SafeCastUpgradeable, bound to a specific deployed contract.
func NewSafeCastUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*SafeCastUpgradeableCaller, error) {
	contract, err := bindSafeCastUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastUpgradeableCaller{contract: contract}, nil
}

// NewSafeCastUpgradeableTransactor creates a new write-only instance of SafeCastUpgradeable, bound to a specific deployed contract.
func NewSafeCastUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeCastUpgradeableTransactor, error) {
	contract, err := bindSafeCastUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeCastUpgradeableTransactor{contract: contract}, nil
}

// NewSafeCastUpgradeableFilterer creates a new log filterer instance of SafeCastUpgradeable, bound to a specific deployed contract.
func NewSafeCastUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeCastUpgradeableFilterer, error) {
	contract, err := bindSafeCastUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeCastUpgradeableFilterer{contract: contract}, nil
}

// bindSafeCastUpgradeable binds a generic wrapper to an already deployed contract.
func bindSafeCastUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeCastUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCastUpgradeable *SafeCastUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCastUpgradeable.Contract.SafeCastUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCastUpgradeable *SafeCastUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCastUpgradeable.Contract.SafeCastUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCastUpgradeable *SafeCastUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCastUpgradeable.Contract.SafeCastUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeCastUpgradeable *SafeCastUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeCastUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeCastUpgradeable *SafeCastUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeCastUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeCastUpgradeable *SafeCastUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeCastUpgradeable.Contract.contract.Transact(opts, method, params...)
}
