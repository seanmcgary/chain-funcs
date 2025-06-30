// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AVSRegistrar

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

// AVSRegistrarMetaData contains all meta data concerning the AVSRegistrar contract.
var AVSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"_keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"KeyNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllocationManager\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b5060405161086338038061086383398101604081905261002e9161012a565b6001600160a01b0380841660805280831660a052811660c05261004f610057565b505050610174565b5f54610100900460ff16156100c25760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610111575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b0381168114610127575f5ffd5b50565b5f5f5f6060848603121561013c575f5ffd5b835161014781610113565b602085015190935061015881610113565b604085015190925061016981610113565b809150509250925092565b60805160a05160c0516106a46101bf5f395f8181607e015261035b01525f818161012501528181610179015261020b01525f818160cd0152818161014c01526102bd01526106a45ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c8063303ca956146100645780633ec45c7e14610079578063b5265787146100bd578063c63fd5021461010d578063ca8aa7c714610120578063de1164bb14610147575b5f5ffd5b61007761007236600461045d565b61016e565b005b6100a07f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b6100fd6100cb3660046104ba565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b60405190151581526020016100b4565b61007761011b3660046104da565b610200565b6100a07f000000000000000000000000000000000000000000000000000000000000000081565b6100a07f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101b7576040516335e5ec5560e21b815260040160405180910390fd5b836001600160a01b03167ff8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d183836040516101f29291906105ab565b60405180910390a250505050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610249576040516335e5ec5560e21b815260040160405180910390fd5b61025486858561029f565b856001600160a01b03167f9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d858560405161028f9291906105ab565b60405180910390a2505050505050565b5f5b63ffffffff81168211156103f4575f60405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200185858563ffffffff16818110610302576103026105f2565b90506020020160208101906103179190610606565b63ffffffff90811690915260405163029ab09960e21b815282516001600160a01b0390811660048301526020840151909216602482015287821660448201529192507f00000000000000000000000000000000000000000000000000000000000000001690630a6ac26490606401602060405180830381865afa1580156103a0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103c4919061061f565b6103e15760405163815589fb60e01b815260040160405180910390fd5b50806103ec8161063e565b9150506102a1565b50505050565b80356001600160a01b0381168114610410575f5ffd5b919050565b5f5f83601f840112610425575f5ffd5b50813567ffffffffffffffff81111561043c575f5ffd5b6020830191508360208260051b8501011115610456575f5ffd5b9250929050565b5f5f5f5f60608587031215610470575f5ffd5b610479856103fa565b9350610487602086016103fa565b9250604085013567ffffffffffffffff8111156104a2575f5ffd5b6104ae87828801610415565b95989497509550505050565b5f602082840312156104ca575f5ffd5b6104d3826103fa565b9392505050565b5f5f5f5f5f5f608087890312156104ef575f5ffd5b6104f8876103fa565b9550610506602088016103fa565b9450604087013567ffffffffffffffff811115610521575f5ffd5b61052d89828a01610415565b909550935050606087013567ffffffffffffffff81111561054c575f5ffd5b8701601f8101891361055c575f5ffd5b803567ffffffffffffffff811115610572575f5ffd5b896020828401011115610583575f5ffd5b60208201935080925050509295509295509295565b803563ffffffff81168114610410575f5ffd5b602080825281018290525f8360408301825b858110156105e85763ffffffff6105d384610598565b168252602092830192909101906001016105bd565b5095945050505050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215610616575f5ffd5b6104d382610598565b5f6020828403121561062f575f5ffd5b815180151581146104d3575f5ffd5b5f63ffffffff821663ffffffff810361066557634e487b7160e01b5f52601160045260245ffd5b6001019291505056fea2646970667358221220c956b3c3f38e9d508df80dfe1ac3b0765174b4e5a8835d15cb951157f9d4f9c664736f6c634300081b0033",
}

// AVSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use AVSRegistrarMetaData.ABI instead.
var AVSRegistrarABI = AVSRegistrarMetaData.ABI

// AVSRegistrarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AVSRegistrarMetaData.Bin instead.
var AVSRegistrarBin = AVSRegistrarMetaData.Bin

// DeployAVSRegistrar deploys a new Ethereum contract, binding an instance of AVSRegistrar to it.
func DeployAVSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, _avs common.Address, _allocationManager common.Address, _keyRegistrar common.Address) (common.Address, *types.Transaction, *AVSRegistrar, error) {
	parsed, err := AVSRegistrarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AVSRegistrarBin), backend, _avs, _allocationManager, _keyRegistrar)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AVSRegistrar{AVSRegistrarCaller: AVSRegistrarCaller{contract: contract}, AVSRegistrarTransactor: AVSRegistrarTransactor{contract: contract}, AVSRegistrarFilterer: AVSRegistrarFilterer{contract: contract}}, nil
}

// AVSRegistrar is an auto generated Go binding around an Ethereum contract.
type AVSRegistrar struct {
	AVSRegistrarCaller     // Read-only binding to the contract
	AVSRegistrarTransactor // Write-only binding to the contract
	AVSRegistrarFilterer   // Log filterer for contract events
}

// AVSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type AVSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AVSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AVSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AVSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AVSRegistrarSession struct {
	Contract     *AVSRegistrar     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AVSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AVSRegistrarCallerSession struct {
	Contract *AVSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AVSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AVSRegistrarTransactorSession struct {
	Contract     *AVSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AVSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type AVSRegistrarRaw struct {
	Contract *AVSRegistrar // Generic contract binding to access the raw methods on
}

// AVSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AVSRegistrarCallerRaw struct {
	Contract *AVSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// AVSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AVSRegistrarTransactorRaw struct {
	Contract *AVSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAVSRegistrar creates a new instance of AVSRegistrar, bound to a specific deployed contract.
func NewAVSRegistrar(address common.Address, backend bind.ContractBackend) (*AVSRegistrar, error) {
	contract, err := bindAVSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrar{AVSRegistrarCaller: AVSRegistrarCaller{contract: contract}, AVSRegistrarTransactor: AVSRegistrarTransactor{contract: contract}, AVSRegistrarFilterer: AVSRegistrarFilterer{contract: contract}}, nil
}

// NewAVSRegistrarCaller creates a new read-only instance of AVSRegistrar, bound to a specific deployed contract.
func NewAVSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*AVSRegistrarCaller, error) {
	contract, err := bindAVSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarCaller{contract: contract}, nil
}

// NewAVSRegistrarTransactor creates a new write-only instance of AVSRegistrar, bound to a specific deployed contract.
func NewAVSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*AVSRegistrarTransactor, error) {
	contract, err := bindAVSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarTransactor{contract: contract}, nil
}

// NewAVSRegistrarFilterer creates a new log filterer instance of AVSRegistrar, bound to a specific deployed contract.
func NewAVSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*AVSRegistrarFilterer, error) {
	contract, err := bindAVSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarFilterer{contract: contract}, nil
}

// bindAVSRegistrar binds a generic wrapper to an already deployed contract.
func bindAVSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AVSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSRegistrar *AVSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSRegistrar.Contract.AVSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSRegistrar *AVSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.AVSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSRegistrar *AVSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.AVSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AVSRegistrar *AVSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AVSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AVSRegistrar *AVSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AVSRegistrar *AVSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_AVSRegistrar *AVSRegistrarCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSRegistrar.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_AVSRegistrar *AVSRegistrarSession) AllocationManager() (common.Address, error) {
	return _AVSRegistrar.Contract.AllocationManager(&_AVSRegistrar.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_AVSRegistrar *AVSRegistrarCallerSession) AllocationManager() (common.Address, error) {
	return _AVSRegistrar.Contract.AllocationManager(&_AVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_AVSRegistrar *AVSRegistrarCaller) Avs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSRegistrar.contract.Call(opts, &out, "avs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_AVSRegistrar *AVSRegistrarSession) Avs() (common.Address, error) {
	return _AVSRegistrar.Contract.Avs(&_AVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_AVSRegistrar *AVSRegistrarCallerSession) Avs() (common.Address, error) {
	return _AVSRegistrar.Contract.Avs(&_AVSRegistrar.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_AVSRegistrar *AVSRegistrarCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AVSRegistrar.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_AVSRegistrar *AVSRegistrarSession) KeyRegistrar() (common.Address, error) {
	return _AVSRegistrar.Contract.KeyRegistrar(&_AVSRegistrar.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_AVSRegistrar *AVSRegistrarCallerSession) KeyRegistrar() (common.Address, error) {
	return _AVSRegistrar.Contract.KeyRegistrar(&_AVSRegistrar.CallOpts)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_AVSRegistrar *AVSRegistrarCaller) SupportsAVS(opts *bind.CallOpts, _avs common.Address) (bool, error) {
	var out []interface{}
	err := _AVSRegistrar.contract.Call(opts, &out, "supportsAVS", _avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_AVSRegistrar *AVSRegistrarSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _AVSRegistrar.Contract.SupportsAVS(&_AVSRegistrar.CallOpts, _avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_AVSRegistrar *AVSRegistrarCallerSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _AVSRegistrar.Contract.SupportsAVS(&_AVSRegistrar.CallOpts, _avs)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_AVSRegistrar *AVSRegistrarTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSRegistrar.contract.Transact(opts, "deregisterOperator", operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_AVSRegistrar *AVSRegistrarSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.DeregisterOperator(&_AVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_AVSRegistrar *AVSRegistrarTransactorSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.DeregisterOperator(&_AVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_AVSRegistrar *AVSRegistrarTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _AVSRegistrar.contract.Transact(opts, "registerOperator", operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_AVSRegistrar *AVSRegistrarSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.RegisterOperator(&_AVSRegistrar.TransactOpts, operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_AVSRegistrar *AVSRegistrarTransactorSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _AVSRegistrar.Contract.RegisterOperator(&_AVSRegistrar.TransactOpts, operator, avs, operatorSetIds, data)
}

// AVSRegistrarInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AVSRegistrar contract.
type AVSRegistrarInitializedIterator struct {
	Event *AVSRegistrarInitialized // Event containing the contract specifics and raw log

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
func (it *AVSRegistrarInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSRegistrarInitialized)
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
		it.Event = new(AVSRegistrarInitialized)
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
func (it *AVSRegistrarInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSRegistrarInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSRegistrarInitialized represents a Initialized event raised by the AVSRegistrar contract.
type AVSRegistrarInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSRegistrar *AVSRegistrarFilterer) FilterInitialized(opts *bind.FilterOpts) (*AVSRegistrarInitializedIterator, error) {

	logs, sub, err := _AVSRegistrar.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarInitializedIterator{contract: _AVSRegistrar.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSRegistrar *AVSRegistrarFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AVSRegistrarInitialized) (event.Subscription, error) {

	logs, sub, err := _AVSRegistrar.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSRegistrarInitialized)
				if err := _AVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AVSRegistrar *AVSRegistrarFilterer) ParseInitialized(log types.Log) (*AVSRegistrarInitialized, error) {
	event := new(AVSRegistrarInitialized)
	if err := _AVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSRegistrarOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the AVSRegistrar contract.
type AVSRegistrarOperatorDeregisteredIterator struct {
	Event *AVSRegistrarOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *AVSRegistrarOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSRegistrarOperatorDeregistered)
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
		it.Event = new(AVSRegistrarOperatorDeregistered)
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
func (it *AVSRegistrarOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSRegistrarOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSRegistrarOperatorDeregistered represents a OperatorDeregistered event raised by the AVSRegistrar contract.
type AVSRegistrarOperatorDeregistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrar *AVSRegistrarFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*AVSRegistrarOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrar.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarOperatorDeregisteredIterator{contract: _AVSRegistrar.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrar *AVSRegistrarFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *AVSRegistrarOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrar.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSRegistrarOperatorDeregistered)
				if err := _AVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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
func (_AVSRegistrar *AVSRegistrarFilterer) ParseOperatorDeregistered(log types.Log) (*AVSRegistrarOperatorDeregistered, error) {
	event := new(AVSRegistrarOperatorDeregistered)
	if err := _AVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AVSRegistrarOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the AVSRegistrar contract.
type AVSRegistrarOperatorRegisteredIterator struct {
	Event *AVSRegistrarOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *AVSRegistrarOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AVSRegistrarOperatorRegistered)
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
		it.Event = new(AVSRegistrarOperatorRegistered)
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
func (it *AVSRegistrarOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AVSRegistrarOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AVSRegistrarOperatorRegistered represents a OperatorRegistered event raised by the AVSRegistrar contract.
type AVSRegistrarOperatorRegistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrar *AVSRegistrarFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*AVSRegistrarOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrar.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &AVSRegistrarOperatorRegisteredIterator{contract: _AVSRegistrar.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_AVSRegistrar *AVSRegistrarFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *AVSRegistrarOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _AVSRegistrar.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AVSRegistrarOperatorRegistered)
				if err := _AVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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
func (_AVSRegistrar *AVSRegistrarFilterer) ParseOperatorRegistered(log types.Log) (*AVSRegistrarOperatorRegistered, error) {
	event := new(AVSRegistrarOperatorRegistered)
	if err := _AVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
