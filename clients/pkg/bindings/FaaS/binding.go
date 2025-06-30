// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FaaS

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

// FaaSFunctionMetadata is an auto generated low-level Go binding around an user-defined struct.
type FaaSFunctionMetadata struct {
	HasContent    bool
	HasUrl        bool
	ContentLength uint32
}

// FaaSMetaData contains all meta data concerning the FaaS contract.
var FaaSMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskMailbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_avs\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executorOperatorSetId\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"avs\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"callFunction\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"arguments\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployFunction\",\"inputs\":[{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executorOperatorSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"functionContent\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"functionMetadata\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"hasContent\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"hasUrl\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"contentLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"functionUrls\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunction\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFaaS.Function\",\"components\":[{\"name\":\"content\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunctionContent\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunctionMetadata\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFaaS.FunctionMetadata\",\"components\":[{\"name\":\"hasContent\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"hasUrl\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"contentLength\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFunctionUrl\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerFunction\",\"inputs\":[{\"name\":\"content\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"taskMailbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractITaskMailbox\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"FunctionCalled\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"taskId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FunctionDeployed\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"url\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FunctionRegistered\",\"inputs\":[{\"name\":\"functionId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"registrar\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x60c060405234801561000f575f5ffd5b506040516110d63803806110d683398101604081905261002e91610079565b6001600160a01b03928316608052911660a0525f805463ffffffff191663ffffffff9092169190911790556100c5565b80516001600160a01b0381168114610074575f5ffd5b919050565b5f5f5f6060848603121561008b575f5ffd5b6100948461005e565b92506100a26020850161005e565b9150604084015163ffffffff811681146100ba575f5ffd5b809150509250925092565b60805160a051610fe26100f45f395f8181610256015261064001525f81816102c801526106a40152610fe25ff3fe608060405234801561000f575f5ffd5b50600436106100cb575f3560e01c8063c428b07511610088578063e77ebc8f11610063578063e77ebc8f146102b0578063f42a9e13146102c3578063f8910880146102ea578063faddefeb146102fd575f5ffd5b8063c428b0751461022d578063de1164bb14610251578063e67950d314610290575f5ffd5b8063064a3a62146100cf578063228764b21461013157806324281fef146101c65780633292c49c146101e657806349262a73146101f95780635b6116f01461021a575b5f5ffd5b6101096100dd366004610b5c565b60016020525f908152604090205460ff8082169161010081049091169062010000900463ffffffff1683565b604080519315158452911515602084015263ffffffff16908201526060015b60405180910390f35b61019a61013f366004610b5c565b60408051606080820183525f80835260208084018290529284018190529384526001825292829020825193840183525460ff8082161515855261010082041615159184019190915262010000900463ffffffff169082015290565b604080518251151581526020808401511515908201529181015163ffffffff1690820152606001610128565b6101d96101d4366004610b5c565b610310565b6040516101289190610ba1565b6101d96101f4366004610b5c565b6103af565b61020c610207366004610c63565b6103cb565b604051908152602001610128565b61020c610228366004610ca7565b610754565b5f5461023c9063ffffffff1681565b60405163ffffffff9091168152602001610128565b6102787f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b039091168152602001610128565b6102a361029e366004610b5c565b610823565b6040516101289190610cfc565b61020c6102be366004610d3d565b6109f0565b6102787f000000000000000000000000000000000000000000000000000000000000000081565b6101d96102f8366004610b5c565b610aad565b6101d961030b366004610b5c565b610b44565b5f81815260026020526040902080546060919061032c90610d77565b80601f016020809104026020016040519081016040528092919081815260200182805461035890610d77565b80156103a35780601f1061037a576101008083540402835291602001916103a3565b820191905f5260205f20905b81548152906001019060200180831161038657829003601f168201915b50505050509050919050565b5f81815260036020526040902080546060919061032c90610d77565b5f828152600160205260408120805460ff16806103ee57508054610100900460ff165b6104335760405162461bcd60e51b8152602060048201526012602482015271119d5b98dd1a5bdb881b9bdd08199bdd5b9960721b604482015260640160405180910390fd5b805460609060ff16156105355761046a6040518060800160405280606081526020015f815260200160608152602001606081525090565b5f868152600260205260409020805461048290610d77565b80601f01602080910402602001604051908101604052809291908181526020018280546104ae90610d77565b80156104f95780601f106104d0576101008083540402835291602001916104f9565b820191905f5260205f20905b8154815290600101906020018083116104dc57829003601f168201915b5050509183525050602080820187905260408083018790525161051e91839101610daf565b60405160208183030381529060405291505061062d565b61055f6040518060800160405280606081526020015f815260200160608152602001606081525090565b602080820187905260408083018790525f88815260039092529020805461058590610d77565b80601f01602080910402602001604051908101604052809291908181526020018280546105b190610d77565b80156105fc5780601f106105d3576101008083540402835291602001916105fc565b820191905f5260205f20905b8154815290600101906020018083116105df57829003601f168201915b505050505081606001819052508060405160200161061a9190610daf565b6040516020818303038152906040529150505b6040805180820182526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811682525f805463ffffffff16602080850191909152845160808101865233815290810182905280850184905260608101869052935162221dbd60e51b8152929390927f000000000000000000000000000000000000000000000000000000000000000090921691630443b7a0916106d991600401610e0f565b6020604051808303815f875af11580156106f5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107199190610e78565b6040519091503390829089907f602ae46fd037d4c5652ff0c0e3ac0f0ff6744765ff3f2f8d9fe77f112d3e0985905f90a49695505050505050565b604080516060810182525f80825260016020808401828152848601848152878552928252858420945185549151935163ffffffff16620100000265ffffffff0000199415156101000261ff00199215159290921661ffff199093169290921717929092169190911790925560039091529081206107d18482610edb565b5081836040516107e19190610f96565b6040519081900381203382529084907f088ac90e03ead29bcc2c251b1567dc223578786d0600fe99db7bc1d8f14979619060200160405180910390a450919050565b604080518082018252606080825260208083018290525f85815260018252849020845192830185525460ff8082161515845261010082041615159183019190915262010000900463ffffffff16818401528251808401909352805191929091819061089b57604080515f815260208101909152610930565b5f85815260026020526040902080546108b390610d77565b80601f01602080910402602001604051908101604052809291908181526020018280546108df90610d77565b801561092a5780601f106109015761010080835404028352916020019161092a565b820191905f5260205f20905b81548152906001019060200180831161090d57829003601f168201915b50505050505b815260200182602001516109525760405180602001604052805f8152506109e7565b5f858152600360205260409020805461096a90610d77565b80601f016020809104026020016040519081016040528092919081815260200182805461099690610d77565b80156109e15780601f106109b8576101008083540402835291602001916109e1565b820191905f5260205f20905b8154815290600101906020018083116109c457829003601f168201915b50505050505b90529392505050565b80516020808301919091206040805160608101825260018082525f828601818152875163ffffffff908116858701908152878452938852858320945185549251945161ffff1990931690151561ff00191617610100941515949094029390931765ffffffff000019166201000091909316029190911790915560029093528220610a7a8482610edb565b50604051339082907fbd508acd226898e9e1b1b8b1f8e4197b63a90c6f2266a22233a966d6f6970bb1905f90a392915050565b60036020525f908152604090208054610ac590610d77565b80601f0160208091040260200160405190810160405280929190818152602001828054610af190610d77565b8015610b3c5780601f10610b1357610100808354040283529160200191610b3c565b820191905f5260205f20905b815481529060010190602001808311610b1f57829003601f168201915b505050505081565b60026020525f908152604090208054610ac590610d77565b5f60208284031215610b6c575f5ffd5b5035919050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f610bb36020830184610b73565b9392505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f67ffffffffffffffff841115610be857610be8610bba565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715610c1757610c17610bba565b604052838152905080828401851015610c2e575f5ffd5b838360208301375f60208583010152509392505050565b5f82601f830112610c54575f5ffd5b610bb383833560208501610bce565b5f5f60408385031215610c74575f5ffd5b82359150602083013567ffffffffffffffff811115610c91575f5ffd5b610c9d85828601610c45565b9150509250929050565b5f5f60408385031215610cb8575f5ffd5b823567ffffffffffffffff811115610cce575f5ffd5b8301601f81018513610cde575f5ffd5b610ced85823560208401610bce565b95602094909401359450505050565b602081525f825160406020840152610d176060840182610b73565b90506020840151601f19848303016040850152610d348282610b73565b95945050505050565b5f60208284031215610d4d575f5ffd5b813567ffffffffffffffff811115610d63575f5ffd5b610d6f84828501610c45565b949350505050565b600181811c90821680610d8b57607f821691505b602082108103610da957634e487b7160e01b5f52602260045260245ffd5b50919050565b602081525f825160806020840152610dca60a0840182610b73565b9050602084015160408401526040840151601f19848303016060850152610df18282610b73565b9150506060840151601f19848303016080850152610d348282610b73565b6020815260018060a01b0382511660208201526bffffffffffffffffffffffff60208301511660408201525f604083015160018060a01b03815116606084015263ffffffff602082015116608084015250606083015160a080840152610d6f60c0840182610b73565b5f60208284031215610e88575f5ffd5b5051919050565b601f821115610ed657805f5260205f20601f840160051c81016020851015610eb45750805b601f840160051c820191505b81811015610ed3575f8155600101610ec0565b50505b505050565b815167ffffffffffffffff811115610ef557610ef5610bba565b610f0981610f038454610d77565b84610e8f565b6020601f821160018114610f3b575f8315610f245750848201515b5f19600385901b1c1916600184901b178455610ed3565b5f84815260208120601f198516915b82811015610f6a5787850151825560209485019460019092019101610f4a565b5084821015610f8757868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82518060208501845e5f92019182525091905056fea26469706673582212205174dc4c964ccbcd095010833791c502b601d37565564f92d33c337a35f5d53664736f6c634300081b0033",
}

// FaaSABI is the input ABI used to generate the binding from.
// Deprecated: Use FaaSMetaData.ABI instead.
var FaaSABI = FaaSMetaData.ABI

// FaaSBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FaaSMetaData.Bin instead.
var FaaSBin = FaaSMetaData.Bin

// DeployFaaS deploys a new Ethereum contract, binding an instance of FaaS to it.
func DeployFaaS(auth *bind.TransactOpts, backend bind.ContractBackend, _taskMailbox common.Address, _avs common.Address, _executorOperatorSetId uint32) (common.Address, *types.Transaction, *FaaS, error) {
	parsed, err := FaaSMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FaaSBin), backend, _taskMailbox, _avs, _executorOperatorSetId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FaaS{FaaSCaller: FaaSCaller{contract: contract}, FaaSTransactor: FaaSTransactor{contract: contract}, FaaSFilterer: FaaSFilterer{contract: contract}}, nil
}

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

// FunctionContent is a free data retrieval call binding the contract method 0xfaddefeb.
//
// Solidity: function functionContent(bytes32 ) view returns(bytes)
func (_FaaS *FaaSCaller) FunctionContent(opts *bind.CallOpts, arg0 [32]byte) ([]byte, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "functionContent", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FunctionContent is a free data retrieval call binding the contract method 0xfaddefeb.
//
// Solidity: function functionContent(bytes32 ) view returns(bytes)
func (_FaaS *FaaSSession) FunctionContent(arg0 [32]byte) ([]byte, error) {
	return _FaaS.Contract.FunctionContent(&_FaaS.CallOpts, arg0)
}

// FunctionContent is a free data retrieval call binding the contract method 0xfaddefeb.
//
// Solidity: function functionContent(bytes32 ) view returns(bytes)
func (_FaaS *FaaSCallerSession) FunctionContent(arg0 [32]byte) ([]byte, error) {
	return _FaaS.Contract.FunctionContent(&_FaaS.CallOpts, arg0)
}

// FunctionMetadata is a free data retrieval call binding the contract method 0x064a3a62.
//
// Solidity: function functionMetadata(bytes32 ) view returns(bool hasContent, bool hasUrl, uint32 contentLength)
func (_FaaS *FaaSCaller) FunctionMetadata(opts *bind.CallOpts, arg0 [32]byte) (struct {
	HasContent    bool
	HasUrl        bool
	ContentLength uint32
}, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "functionMetadata", arg0)

	outstruct := new(struct {
		HasContent    bool
		HasUrl        bool
		ContentLength uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HasContent = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.HasUrl = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.ContentLength = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// FunctionMetadata is a free data retrieval call binding the contract method 0x064a3a62.
//
// Solidity: function functionMetadata(bytes32 ) view returns(bool hasContent, bool hasUrl, uint32 contentLength)
func (_FaaS *FaaSSession) FunctionMetadata(arg0 [32]byte) (struct {
	HasContent    bool
	HasUrl        bool
	ContentLength uint32
}, error) {
	return _FaaS.Contract.FunctionMetadata(&_FaaS.CallOpts, arg0)
}

// FunctionMetadata is a free data retrieval call binding the contract method 0x064a3a62.
//
// Solidity: function functionMetadata(bytes32 ) view returns(bool hasContent, bool hasUrl, uint32 contentLength)
func (_FaaS *FaaSCallerSession) FunctionMetadata(arg0 [32]byte) (struct {
	HasContent    bool
	HasUrl        bool
	ContentLength uint32
}, error) {
	return _FaaS.Contract.FunctionMetadata(&_FaaS.CallOpts, arg0)
}

// FunctionUrls is a free data retrieval call binding the contract method 0xf8910880.
//
// Solidity: function functionUrls(bytes32 ) view returns(string)
func (_FaaS *FaaSCaller) FunctionUrls(opts *bind.CallOpts, arg0 [32]byte) (string, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "functionUrls", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// FunctionUrls is a free data retrieval call binding the contract method 0xf8910880.
//
// Solidity: function functionUrls(bytes32 ) view returns(string)
func (_FaaS *FaaSSession) FunctionUrls(arg0 [32]byte) (string, error) {
	return _FaaS.Contract.FunctionUrls(&_FaaS.CallOpts, arg0)
}

// FunctionUrls is a free data retrieval call binding the contract method 0xf8910880.
//
// Solidity: function functionUrls(bytes32 ) view returns(string)
func (_FaaS *FaaSCallerSession) FunctionUrls(arg0 [32]byte) (string, error) {
	return _FaaS.Contract.FunctionUrls(&_FaaS.CallOpts, arg0)
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

// GetFunctionMetadata is a free data retrieval call binding the contract method 0x228764b2.
//
// Solidity: function getFunctionMetadata(bytes32 functionId) view returns((bool,bool,uint32))
func (_FaaS *FaaSCaller) GetFunctionMetadata(opts *bind.CallOpts, functionId [32]byte) (FaaSFunctionMetadata, error) {
	var out []interface{}
	err := _FaaS.contract.Call(opts, &out, "getFunctionMetadata", functionId)

	if err != nil {
		return *new(FaaSFunctionMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(FaaSFunctionMetadata)).(*FaaSFunctionMetadata)

	return out0, err

}

// GetFunctionMetadata is a free data retrieval call binding the contract method 0x228764b2.
//
// Solidity: function getFunctionMetadata(bytes32 functionId) view returns((bool,bool,uint32))
func (_FaaS *FaaSSession) GetFunctionMetadata(functionId [32]byte) (FaaSFunctionMetadata, error) {
	return _FaaS.Contract.GetFunctionMetadata(&_FaaS.CallOpts, functionId)
}

// GetFunctionMetadata is a free data retrieval call binding the contract method 0x228764b2.
//
// Solidity: function getFunctionMetadata(bytes32 functionId) view returns((bool,bool,uint32))
func (_FaaS *FaaSCallerSession) GetFunctionMetadata(functionId [32]byte) (FaaSFunctionMetadata, error) {
	return _FaaS.Contract.GetFunctionMetadata(&_FaaS.CallOpts, functionId)
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
