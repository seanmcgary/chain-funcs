// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TaskAVSRegistrar

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

// TaskAVSRegistrarMetaData contains all meta data concerning the TaskAVSRegistrar contract.
var TaskAVSRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allocationManager\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"},{\"name\":\"keyRegistrar\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allocationManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIAllocationManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deregisterOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperatorSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyRegistrar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIKeyRegistrar\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsAVS\",\"inputs\":[{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateSocket\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeregistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operatorSetIds\",\"type\":\"uint32[]\",\"indexed\":false,\"internalType\":\"uint32[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorSocketSet\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"socket\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerNotOperator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DataLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"KeyNotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAllocationManager\",\"inputs\":[]}]",
	Bin: "0x60e060405234801561000f575f5ffd5b50604051610ce5380380610ce583398101604081905261002e91610142565b6001600160a01b0380841660805280831660a052811660c05282828282828282828282828261005b61006f565b50505050505050505050505050505061018c565b5f54610100900460ff16156100da5760405162461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b606482015260840160405180910390fd5b5f5460ff90811614610129575f805460ff191660ff9081179091556040519081527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b565b6001600160a01b038116811461013f575f5ffd5b50565b5f5f5f60608486031215610154575f5ffd5b835161015f8161012b565b60208501519093506101708161012b565b60408501519092506101818161012b565b809150509250925092565b60805160a05160c051610b0d6101d85f395f818160a3015261051301525f818161017d015281816101d1015261034301525f8181610125015281816101a401526104750152610b0d5ff3fe608060405234801561000f575f5ffd5b5060043610610085575f3560e01c8063b526578711610058578063b526578714610115578063c63fd50214610165578063ca8aa7c714610178578063de1164bb1461019f575f5ffd5b8063303ca956146100895780633ec45c7e1461009e5780636591666a146100e25780638481931d146100f5575b5f5ffd5b61009c610097366004610636565b6101c6565b005b6100c57f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b61009c6100f0366004610732565b610258565b61010861010336600461077d565b61028f565b6040516100d9919061079d565b61015561012336600461077d565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0390811691161490565b60405190151581526020016100d9565b61009c6101733660046107d2565b610338565b6100c57f000000000000000000000000000000000000000000000000000000000000000081565b6100c57f000000000000000000000000000000000000000000000000000000000000000081565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461020f576040516335e5ec5560e21b815260040160405180910390fd5b836001600160a01b03167ff8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1838360405161024a9291906108a3565b60405180910390a250505050565b336001600160a01b038316146102815760405163a5523ee560e01b815260040160405180910390fd5b61028b82826103e9565b5050565b6001600160a01b0381165f9081526033602052604090208054606091906102b5906108ea565b80601f01602080910402602001604051908101604052809291908181526020018280546102e1906108ea565b801561032c5780601f106103035761010080835404028352916020019161032c565b820191905f5260205f20905b81548152906001019060200180831161030f57829003601f168201915b50505050509050919050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610381576040516335e5ec5560e21b815260040160405180910390fd5b61038c868585610457565b61039986858585856105b2565b856001600160a01b03167f9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d85856040516103d49291906108a3565b60405180910390a2505050505050565b505050565b6001600160a01b0382165f90815260336020526040902061040a8282610966565b50816001600160a01b03167f0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b82604051610444919061079d565b60405180910390a25050565b5050505050565b5f5b63ffffffff81168211156105ac575f60405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316815260200185858563ffffffff168181106104ba576104ba610a21565b90506020020160208101906104cf9190610a35565b63ffffffff90811690915260405163029ab09960e21b815282516001600160a01b0390811660048301526020840151909216602482015287821660448201529192507f00000000000000000000000000000000000000000000000000000000000000001690630a6ac26490606401602060405180830381865afa158015610558573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061057c9190610a4e565b6105995760405163815589fb60e01b815260040160405180910390fd5b50806105a481610a6d565b915050610459565b50505050565b5f6105bf82840184610a9d565b90506105cb86826103e9565b505050505050565b80356001600160a01b03811681146105e9575f5ffd5b919050565b5f5f83601f8401126105fe575f5ffd5b50813567ffffffffffffffff811115610615575f5ffd5b6020830191508360208260051b850101111561062f575f5ffd5b9250929050565b5f5f5f5f60608587031215610649575f5ffd5b610652856105d3565b9350610660602086016105d3565b9250604085013567ffffffffffffffff81111561067b575f5ffd5b610687878288016105ee565b95989497509550505050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f8301126106b6575f5ffd5b813567ffffffffffffffff8111156106d0576106d0610693565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156106ff576106ff610693565b604052818152838201602001851015610716575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f60408385031215610743575f5ffd5b61074c836105d3565b9150602083013567ffffffffffffffff811115610767575f5ffd5b610773858286016106a7565b9150509250929050565b5f6020828403121561078d575f5ffd5b610796826105d3565b9392505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b5f5f5f5f5f5f608087890312156107e7575f5ffd5b6107f0876105d3565b95506107fe602088016105d3565b9450604087013567ffffffffffffffff811115610819575f5ffd5b61082589828a016105ee565b909550935050606087013567ffffffffffffffff811115610844575f5ffd5b8701601f81018913610854575f5ffd5b803567ffffffffffffffff81111561086a575f5ffd5b89602082840101111561087b575f5ffd5b60208201935080925050509295509295509295565b803563ffffffff811681146105e9575f5ffd5b602080825281018290525f8360408301825b858110156108e05763ffffffff6108cb84610890565b168252602092830192909101906001016108b5565b5095945050505050565b600181811c908216806108fe57607f821691505b60208210810361091c57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156103e457805f5260205f20601f840160051c810160208510156109475750805b601f840160051c820191505b81811015610450575f8155600101610953565b815167ffffffffffffffff81111561098057610980610693565b6109948161098e84546108ea565b84610922565b6020601f8211600181146109c6575f83156109af5750848201515b5f19600385901b1c1916600184901b178455610450565b5f84815260208120601f198516915b828110156109f557878501518255602094850194600190920191016109d5565b5084821015610a1257868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215610a45575f5ffd5b61079682610890565b5f60208284031215610a5e575f5ffd5b81518015158114610796575f5ffd5b5f63ffffffff821663ffffffff8103610a9457634e487b7160e01b5f52601160045260245ffd5b60010192915050565b5f60208284031215610aad575f5ffd5b813567ffffffffffffffff811115610ac3575f5ffd5b610acf848285016106a7565b94935050505056fea264697066735822122012b6b7c6e00839a14cfca183199f1f4cbd216032f18050368ef58dc89f30b2c564736f6c634300081b0033",
}

// TaskAVSRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use TaskAVSRegistrarMetaData.ABI instead.
var TaskAVSRegistrarABI = TaskAVSRegistrarMetaData.ABI

// TaskAVSRegistrarBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TaskAVSRegistrarMetaData.Bin instead.
var TaskAVSRegistrarBin = TaskAVSRegistrarMetaData.Bin

// DeployTaskAVSRegistrar deploys a new Ethereum contract, binding an instance of TaskAVSRegistrar to it.
func DeployTaskAVSRegistrar(auth *bind.TransactOpts, backend bind.ContractBackend, avs common.Address, allocationManager common.Address, keyRegistrar common.Address) (common.Address, *types.Transaction, *TaskAVSRegistrar, error) {
	parsed, err := TaskAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TaskAVSRegistrarBin), backend, avs, allocationManager, keyRegistrar)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TaskAVSRegistrar{TaskAVSRegistrarCaller: TaskAVSRegistrarCaller{contract: contract}, TaskAVSRegistrarTransactor: TaskAVSRegistrarTransactor{contract: contract}, TaskAVSRegistrarFilterer: TaskAVSRegistrarFilterer{contract: contract}}, nil
}

// TaskAVSRegistrar is an auto generated Go binding around an Ethereum contract.
type TaskAVSRegistrar struct {
	TaskAVSRegistrarCaller     // Read-only binding to the contract
	TaskAVSRegistrarTransactor // Write-only binding to the contract
	TaskAVSRegistrarFilterer   // Log filterer for contract events
}

// TaskAVSRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaskAVSRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaskAVSRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaskAVSRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaskAVSRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaskAVSRegistrarSession struct {
	Contract     *TaskAVSRegistrar // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaskAVSRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaskAVSRegistrarCallerSession struct {
	Contract *TaskAVSRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TaskAVSRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaskAVSRegistrarTransactorSession struct {
	Contract     *TaskAVSRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TaskAVSRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaskAVSRegistrarRaw struct {
	Contract *TaskAVSRegistrar // Generic contract binding to access the raw methods on
}

// TaskAVSRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaskAVSRegistrarCallerRaw struct {
	Contract *TaskAVSRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// TaskAVSRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaskAVSRegistrarTransactorRaw struct {
	Contract *TaskAVSRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaskAVSRegistrar creates a new instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrar(address common.Address, backend bind.ContractBackend) (*TaskAVSRegistrar, error) {
	contract, err := bindTaskAVSRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrar{TaskAVSRegistrarCaller: TaskAVSRegistrarCaller{contract: contract}, TaskAVSRegistrarTransactor: TaskAVSRegistrarTransactor{contract: contract}, TaskAVSRegistrarFilterer: TaskAVSRegistrarFilterer{contract: contract}}, nil
}

// NewTaskAVSRegistrarCaller creates a new read-only instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarCaller(address common.Address, caller bind.ContractCaller) (*TaskAVSRegistrarCaller, error) {
	contract, err := bindTaskAVSRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarCaller{contract: contract}, nil
}

// NewTaskAVSRegistrarTransactor creates a new write-only instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*TaskAVSRegistrarTransactor, error) {
	contract, err := bindTaskAVSRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarTransactor{contract: contract}, nil
}

// NewTaskAVSRegistrarFilterer creates a new log filterer instance of TaskAVSRegistrar, bound to a specific deployed contract.
func NewTaskAVSRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*TaskAVSRegistrarFilterer, error) {
	contract, err := bindTaskAVSRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarFilterer{contract: contract}, nil
}

// bindTaskAVSRegistrar binds a generic wrapper to an already deployed contract.
func bindTaskAVSRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaskAVSRegistrarMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskAVSRegistrar *TaskAVSRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.TaskAVSRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaskAVSRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.contract.Transact(opts, method, params...)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) AllocationManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "allocationManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) AllocationManager() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AllocationManager(&_TaskAVSRegistrar.CallOpts)
}

// AllocationManager is a free data retrieval call binding the contract method 0xca8aa7c7.
//
// Solidity: function allocationManager() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) AllocationManager() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.AllocationManager(&_TaskAVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) Avs(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "avs")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) Avs() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Avs(&_TaskAVSRegistrar.CallOpts)
}

// Avs is a free data retrieval call binding the contract method 0xde1164bb.
//
// Solidity: function avs() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) Avs() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.Avs(&_TaskAVSRegistrar.CallOpts)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) GetOperatorSocket(opts *bind.CallOpts, operator common.Address) (string, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "getOperatorSocket", operator)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _TaskAVSRegistrar.Contract.GetOperatorSocket(&_TaskAVSRegistrar.CallOpts, operator)
}

// GetOperatorSocket is a free data retrieval call binding the contract method 0x8481931d.
//
// Solidity: function getOperatorSocket(address operator) view returns(string)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) GetOperatorSocket(operator common.Address) (string, error) {
	return _TaskAVSRegistrar.Contract.GetOperatorSocket(&_TaskAVSRegistrar.CallOpts, operator)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) KeyRegistrar(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "keyRegistrar")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) KeyRegistrar() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.KeyRegistrar(&_TaskAVSRegistrar.CallOpts)
}

// KeyRegistrar is a free data retrieval call binding the contract method 0x3ec45c7e.
//
// Solidity: function keyRegistrar() view returns(address)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) KeyRegistrar() (common.Address, error) {
	return _TaskAVSRegistrar.Contract.KeyRegistrar(&_TaskAVSRegistrar.CallOpts)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCaller) SupportsAVS(opts *bind.CallOpts, _avs common.Address) (bool, error) {
	var out []interface{}
	err := _TaskAVSRegistrar.contract.Call(opts, &out, "supportsAVS", _avs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.SupportsAVS(&_TaskAVSRegistrar.CallOpts, _avs)
}

// SupportsAVS is a free data retrieval call binding the contract method 0xb5265787.
//
// Solidity: function supportsAVS(address _avs) view returns(bool)
func (_TaskAVSRegistrar *TaskAVSRegistrarCallerSession) SupportsAVS(_avs common.Address) (bool, error) {
	return _TaskAVSRegistrar.Contract.SupportsAVS(&_TaskAVSRegistrar.CallOpts, _avs)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) DeregisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "deregisterOperator", operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.DeregisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// DeregisterOperator is a paid mutator transaction binding the contract method 0x303ca956.
//
// Solidity: function deregisterOperator(address operator, address avs, uint32[] operatorSetIds) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) DeregisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.DeregisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) RegisterOperator(opts *bind.TransactOpts, operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "registerOperator", operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RegisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds, data)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0xc63fd502.
//
// Solidity: function registerOperator(address operator, address avs, uint32[] operatorSetIds, bytes data) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) RegisterOperator(operator common.Address, avs common.Address, operatorSetIds []uint32, data []byte) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.RegisterOperator(&_TaskAVSRegistrar.TransactOpts, operator, avs, operatorSetIds, data)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactor) UpdateSocket(opts *bind.TransactOpts, operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.contract.Transact(opts, "updateSocket", operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarSession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.UpdateSocket(&_TaskAVSRegistrar.TransactOpts, operator, socket)
}

// UpdateSocket is a paid mutator transaction binding the contract method 0x6591666a.
//
// Solidity: function updateSocket(address operator, string socket) returns()
func (_TaskAVSRegistrar *TaskAVSRegistrarTransactorSession) UpdateSocket(operator common.Address, socket string) (*types.Transaction, error) {
	return _TaskAVSRegistrar.Contract.UpdateSocket(&_TaskAVSRegistrar.TransactOpts, operator, socket)
}

// TaskAVSRegistrarInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarInitializedIterator struct {
	Event *TaskAVSRegistrarInitialized // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarInitialized)
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
		it.Event = new(TaskAVSRegistrarInitialized)
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
func (it *TaskAVSRegistrarInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarInitialized represents a Initialized event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaskAVSRegistrarInitializedIterator, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarInitializedIterator{contract: _TaskAVSRegistrar.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarInitialized) (event.Subscription, error) {

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarInitialized)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseInitialized(log types.Log) (*TaskAVSRegistrarInitialized, error) {
	event := new(TaskAVSRegistrarInitialized)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorDeregisteredIterator is returned from FilterOperatorDeregistered and is used to iterate over the raw logs and unpacked data for OperatorDeregistered events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorDeregisteredIterator struct {
	Event *TaskAVSRegistrarOperatorDeregistered // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorDeregistered)
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
		it.Event = new(TaskAVSRegistrarOperatorDeregistered)
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
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorDeregistered represents a OperatorDeregistered event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorDeregistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeregistered is a free log retrieval operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorDeregistered(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorDeregisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorDeregisteredIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeregistered is a free log subscription operation binding the contract event 0xf8aaad08ee23b49c9bb44e3bca6c7efa43442fc4281245a7f2475aa2632718d1.
//
// Solidity: event OperatorDeregistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorDeregistered(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorDeregistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorDeregistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorDeregistered)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
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
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorDeregistered(log types.Log) (*TaskAVSRegistrarOperatorDeregistered, error) {
	event := new(TaskAVSRegistrarOperatorDeregistered)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRegisteredIterator struct {
	Event *TaskAVSRegistrarOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorRegistered)
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
		it.Event = new(TaskAVSRegistrarOperatorRegistered)
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
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorRegistered represents a OperatorRegistered event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorRegistered struct {
	Operator       common.Address
	OperatorSetIds []uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorRegisteredIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x9efdc3d07eb312e06bf36ea85db02aec96817d7c7421f919027b240eaf34035d.
//
// Solidity: event OperatorRegistered(address indexed operator, uint32[] operatorSetIds)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorRegistered)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorRegistered(log types.Log) (*TaskAVSRegistrarOperatorRegistered, error) {
	event := new(TaskAVSRegistrarOperatorRegistered)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaskAVSRegistrarOperatorSocketSetIterator is returned from FilterOperatorSocketSet and is used to iterate over the raw logs and unpacked data for OperatorSocketSet events raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorSocketSetIterator struct {
	Event *TaskAVSRegistrarOperatorSocketSet // Event containing the contract specifics and raw log

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
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaskAVSRegistrarOperatorSocketSet)
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
		it.Event = new(TaskAVSRegistrarOperatorSocketSet)
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
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaskAVSRegistrarOperatorSocketSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaskAVSRegistrarOperatorSocketSet represents a OperatorSocketSet event raised by the TaskAVSRegistrar contract.
type TaskAVSRegistrarOperatorSocketSet struct {
	Operator common.Address
	Socket   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSocketSet is a free log retrieval operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) FilterOperatorSocketSet(opts *bind.FilterOpts, operator []common.Address) (*TaskAVSRegistrarOperatorSocketSetIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.FilterLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return &TaskAVSRegistrarOperatorSocketSetIterator{contract: _TaskAVSRegistrar.contract, event: "OperatorSocketSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSocketSet is a free log subscription operation binding the contract event 0x0728b43b8c8244bf835bc60bb800c6834d28d6b696427683617f8d4b0878054b.
//
// Solidity: event OperatorSocketSet(address indexed operator, string socket)
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) WatchOperatorSocketSet(opts *bind.WatchOpts, sink chan<- *TaskAVSRegistrarOperatorSocketSet, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TaskAVSRegistrar.contract.WatchLogs(opts, "OperatorSocketSet", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaskAVSRegistrarOperatorSocketSet)
				if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
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
func (_TaskAVSRegistrar *TaskAVSRegistrarFilterer) ParseOperatorSocketSet(log types.Log) (*TaskAVSRegistrarOperatorSocketSet, error) {
	event := new(TaskAVSRegistrarOperatorSocketSet)
	if err := _TaskAVSRegistrar.contract.UnpackLog(event, "OperatorSocketSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
