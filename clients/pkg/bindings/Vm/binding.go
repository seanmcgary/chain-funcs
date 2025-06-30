// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Vm

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

// VmSafeAccessListItem is an auto generated low-level Go binding around an user-defined struct.
type VmSafeAccessListItem struct {
	Target      common.Address
	StorageKeys [][32]byte
}

// VmSafeAccountAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeAccountAccess struct {
	ChainInfo       VmSafeChainInfo
	Kind            uint8
	Account         common.Address
	Accessor        common.Address
	Initialized     bool
	OldBalance      *big.Int
	NewBalance      *big.Int
	DeployedCode    []byte
	Value           *big.Int
	Data            []byte
	Reverted        bool
	StorageAccesses []VmSafeStorageAccess
	Depth           uint64
}

// VmSafeBroadcastTxSummary is an auto generated low-level Go binding around an user-defined struct.
type VmSafeBroadcastTxSummary struct {
	TxHash          [32]byte
	TxType          uint8
	ContractAddress common.Address
	BlockNumber     uint64
	Success         bool
}

// VmSafeChain is an auto generated low-level Go binding around an user-defined struct.
type VmSafeChain struct {
	Name       string
	ChainId    *big.Int
	ChainAlias string
	RpcUrl     string
}

// VmSafeChainInfo is an auto generated low-level Go binding around an user-defined struct.
type VmSafeChainInfo struct {
	ForkId  *big.Int
	ChainId *big.Int
}

// VmSafeDebugStep is an auto generated low-level Go binding around an user-defined struct.
type VmSafeDebugStep struct {
	Stack        []*big.Int
	MemoryInput  []byte
	Opcode       uint8
	Depth        uint64
	IsOutOfGas   bool
	ContractAddr common.Address
}

// VmSafeDirEntry is an auto generated low-level Go binding around an user-defined struct.
type VmSafeDirEntry struct {
	ErrorMessage string
	Path         string
	Depth        uint64
	IsDir        bool
	IsSymlink    bool
}

// VmSafeEthGetLogs is an auto generated low-level Go binding around an user-defined struct.
type VmSafeEthGetLogs struct {
	Emitter          common.Address
	Topics           [][32]byte
	Data             []byte
	BlockHash        [32]byte
	BlockNumber      uint64
	TransactionHash  [32]byte
	TransactionIndex uint64
	LogIndex         *big.Int
	Removed          bool
}

// VmSafeFfiResult is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFfiResult struct {
	ExitCode int32
	Stdout   []byte
	Stderr   []byte
}

// VmSafeFsMetadata is an auto generated low-level Go binding around an user-defined struct.
type VmSafeFsMetadata struct {
	IsDir     bool
	IsSymlink bool
	Length    *big.Int
	ReadOnly  bool
	Modified  *big.Int
	Accessed  *big.Int
	Created   *big.Int
}

// VmSafeGas is an auto generated low-level Go binding around an user-defined struct.
type VmSafeGas struct {
	GasLimit      uint64
	GasTotalUsed  uint64
	GasMemoryUsed uint64
	GasRefunded   int64
	GasRemaining  uint64
}

// VmSafeLog is an auto generated low-level Go binding around an user-defined struct.
type VmSafeLog struct {
	Topics  [][32]byte
	Data    []byte
	Emitter common.Address
}

// VmSafePotentialRevert is an auto generated low-level Go binding around an user-defined struct.
type VmSafePotentialRevert struct {
	Reverter     common.Address
	PartialMatch bool
	RevertData   []byte
}

// VmSafeRpc is an auto generated low-level Go binding around an user-defined struct.
type VmSafeRpc struct {
	Key string
	Url string
}

// VmSafeSignedDelegation is an auto generated low-level Go binding around an user-defined struct.
type VmSafeSignedDelegation struct {
	V              uint8
	R              [32]byte
	S              [32]byte
	Nonce          uint64
	Implementation common.Address
}

// VmSafeStorageAccess is an auto generated low-level Go binding around an user-defined struct.
type VmSafeStorageAccess struct {
	Account       common.Address
	Slot          [32]byte
	IsWrite       bool
	PreviousValue [32]byte
	NewValue      [32]byte
	Reverted      bool
}

// VmSafeWallet is an auto generated low-level Go binding around an user-defined struct.
type VmSafeWallet struct {
	Addr       common.Address
	PublicKeyX *big.Int
	PublicKeyY *big.Int
	PrivateKey *big.Int
}

// VmMetaData contains all meta data concerning the Vm contract.
var VmMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"accessList\",\"inputs\":[{\"name\":\"access\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.AccessListItem[]\",\"components\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"storageKeys\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accesses\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"readSlots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"writeSlots\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"activeFork\",\"inputs\":[],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addr\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"keyAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"allowCheatcodes\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbs\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqAbsDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRel\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertApproxEqRelDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"maxPercentDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertFalse\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertFalse\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertGtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLe\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLeDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLt\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertLtDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"right\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"right\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"right\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"right\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"right\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"right\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"right\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"right\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"right\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"right\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"right\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"right\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEq\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"right\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertNotEqDecimal\",\"inputs\":[{\"name\":\"left\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"right\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertTrue\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assertTrue\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assume\",\"inputs\":[{\"name\":\"condition\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assumeNoRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assumeNoRevert\",\"inputs\":[{\"name\":\"potentialReverts\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.PotentialRevert[]\",\"components\":[{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"partialMatch\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"assumeNoRevert\",\"inputs\":[{\"name\":\"potentialRevert\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.PotentialRevert\",\"components\":[{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"partialMatch\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"attachBlob\",\"inputs\":[{\"name\":\"blob\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"attachDelegation\",\"inputs\":[{\"name\":\"signedDelegation\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.SignedDelegation\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blobBaseFee\",\"inputs\":[{\"name\":\"newBlobBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blobhashes\",\"inputs\":[{\"name\":\"hashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"breakpoint\",\"inputs\":[{\"name\":\"char\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"breakpoint\",\"inputs\":[{\"name\":\"char\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"broadcast\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcast\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcast\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"broadcastRawTransaction\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[{\"name\":\"newChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearMockedCalls\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cloneAccount\",\"inputs\":[{\"name\":\"source\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"closeFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"coinbase\",\"inputs\":[{\"name\":\"newCoinbase\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"computeCreate2Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"computeCreate2Address\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initCodeHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"computeCreateAddress\",\"inputs\":[{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"contains\",\"inputs\":[{\"name\":\"subject\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"search\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cool\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"coolSlot\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"copyFile\",\"inputs\":[{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"copied\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"copyStorage\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"recursive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createFork\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createFork\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createFork\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSelectFork\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSelectFork\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createSelectFork\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createWallet\",\"inputs\":[{\"name\":\"walletLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createWallet\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createWallet\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"walletLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deal\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteSnapshot\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteSnapshots\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteStateSnapshot\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteStateSnapshots\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"constructorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"constructorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"constructorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deployCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"constructorArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"derivationPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"language\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"language\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"deriveKey\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"derivationPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"index\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"difficulty\",\"inputs\":[{\"name\":\"newDifficulty\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dumpState\",\"inputs\":[{\"name\":\"pathToStateJson\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ensNamehash\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"envAddress\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envAddress\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBool\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBool\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes32\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envBytes32\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envExists\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envInt\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envInt\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envOr\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"defaultValue\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envString\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envString\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envUint\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"envUint\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delim\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"etch\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newRuntimeBytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eth_getLogs\",\"inputs\":[{\"name\":\"fromBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"logs\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.EthGetLogs[]\",\"components\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"transactionHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"transactionIndex\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"logIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"removed\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"exists\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"expectCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCallMinGas\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCallMinGas\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCreate\",\"inputs\":[{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectCreate2\",\"inputs\":[{\"name\":\"bytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"deployer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[{\"name\":\"checkTopic1\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic2\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic3\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkData\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[{\"name\":\"checkTopic1\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic2\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic3\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[{\"name\":\"checkTopic1\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic2\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic3\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmit\",\"inputs\":[{\"name\":\"checkTopic1\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic2\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic3\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmitAnonymous\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmitAnonymous\",\"inputs\":[{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmitAnonymous\",\"inputs\":[{\"name\":\"checkTopic0\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic1\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic2\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic3\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkData\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectEmitAnonymous\",\"inputs\":[{\"name\":\"checkTopic0\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic1\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic2\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkTopic3\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"checkData\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectPartialRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectPartialRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"reverter\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"count\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectRevert\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectSafeMemory\",\"inputs\":[{\"name\":\"min\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"max\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"expectSafeMemoryCall\",\"inputs\":[{\"name\":\"min\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"max\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fee\",\"inputs\":[{\"name\":\"newBasefee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ffi\",\"inputs\":[{\"name\":\"commandInput\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"foundryVersionAtLeast\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"foundryVersionCmp\",\"inputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fsMetadata\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"metadata\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.FsMetadata\",\"components\":[{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"readOnly\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"modified\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"accessed\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"created\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getArtifactPathByCode\",\"inputs\":[{\"name\":\"code\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getArtifactPathByDeployedCode\",\"inputs\":[{\"name\":\"deployedCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlobBaseFee\",\"inputs\":[],\"outputs\":[{\"name\":\"blobBaseFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlobhashes\",\"inputs\":[],\"outputs\":[{\"name\":\"hashes\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"height\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBlockTimestamp\",\"inputs\":[],\"outputs\":[{\"name\":\"timestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBroadcast\",\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"txType\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.BroadcastTxType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.BroadcastTxSummary\",\"components\":[{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"txType\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.BroadcastTxType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBroadcasts\",\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.BroadcastTxSummary[]\",\"components\":[{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"txType\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.BroadcastTxType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBroadcasts\",\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"txType\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.BroadcastTxType\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.BroadcastTxSummary[]\",\"components\":[{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"txType\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.BroadcastTxType\"},{\"name\":\"contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChain\",\"inputs\":[{\"name\":\"chainAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"chain\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Chain\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChain\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chain\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Chain\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rpcUrl\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"creationBytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployedCode\",\"inputs\":[{\"name\":\"artifactPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"runtimeBytecode\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployment\",\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployment\",\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"deployedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDeployments\",\"inputs\":[{\"name\":\"contractName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"deployedAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFoundryVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLabel\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"currentLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMappingKeyAndParentOf\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"elementSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"found\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"parent\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMappingLength\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mappingSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMappingSlotAt\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"mappingSlot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"idx\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRecordedLogs\",\"inputs\":[],\"outputs\":[{\"name\":\"logs\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.Log[]\",\"components\":[{\"name\":\"topics\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"emitter\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getStateDiff\",\"inputs\":[],\"outputs\":[{\"name\":\"diff\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateDiffJson\",\"inputs\":[],\"outputs\":[{\"name\":\"diff\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWallets\",\"inputs\":[],\"outputs\":[{\"name\":\"wallets\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"indexOf\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"interceptInitcode\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isContext\",\"inputs\":[{\"name\":\"context\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.ForgeContext\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPersistent\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"persistent\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyExists\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyExistsJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"keyExistsToml\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"label\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newLabel\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastCallGas\",\"inputs\":[],\"outputs\":[{\"name\":\"gas\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Gas\",\"components\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasTotalUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasMemoryUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasRefunded\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"gasRemaining\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"load\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"loadAllocs\",\"inputs\":[{\"name\":\"pathToAllocsJson\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"makePersistent\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"makePersistent\",\"inputs\":[{\"name\":\"account0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account1\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"makePersistent\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"makePersistent\",\"inputs\":[{\"name\":\"account0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"account2\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCall\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCallRevert\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCallRevert\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCallRevert\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCallRevert\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"revertData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCalls\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockCalls\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"returnData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mockFunction\",\"inputs\":[{\"name\":\"callee\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"noAccessList\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"parseAddress\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseBool\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseBytes\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseBytes32\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseInt\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonAddress\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonAddressArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBool\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBoolArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytes\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytes32\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytes32Array\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonBytesArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonInt\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonIntArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonKeys\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"keys\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonString\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonStringArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonType\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonType\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonTypeArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonUint\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseJsonUintArray\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseToml\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseToml\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"abiEncodedData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlAddress\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlAddressArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBool\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBoolArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytes\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytes32\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytes32Array\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlBytesArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlInt\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlIntArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlKeys\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"keys\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlString\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlStringArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlType\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlType\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlTypeArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlUint\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseTomlUintArray\",\"inputs\":[{\"name\":\"toml\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"parseUint\",\"inputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"parsedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"pauseGasMetering\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseTracing\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"txOrigin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"txOrigin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prevrandao\",\"inputs\":[{\"name\":\"newPrevrandao\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"prevrandao\",\"inputs\":[{\"name\":\"newPrevrandao\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"projectRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prompt\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptAddress\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptSecret\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptSecretUint\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"promptUint\",\"inputs\":[{\"name\":\"promptText\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"publicKeyP256\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"randomAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"randomBool\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randomBytes\",\"inputs\":[{\"name\":\"len\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randomBytes4\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randomBytes8\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes8\",\"internalType\":\"bytes8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randomInt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randomInt\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randomUint\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"randomUint\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"randomUint\",\"inputs\":[{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"readCallers\",\"inputs\":[],\"outputs\":[{\"name\":\"callerMode\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.CallerMode\"},{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"txOrigin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"readDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxDepth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.DirEntry[]\",\"components\":[{\"name\":\"errorMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"maxDepth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"followLinks\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.DirEntry[]\",\"components\":[{\"name\":\"errorMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"entries\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.DirEntry[]\",\"components\":[{\"name\":\"errorMessage\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isDir\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isSymlink\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readFileBinary\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readLine\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"line\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"readLink\",\"inputs\":[{\"name\":\"linkPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"targetPath\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"record\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recordLogs\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rememberKey\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"keyAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rememberKeys\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"derivationPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"count\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"keyAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rememberKeys\",\"inputs\":[{\"name\":\"mnemonic\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"derivationPath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"language\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"count\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"keyAddrs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeDir\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"recursive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"replace\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"from\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"to\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"resetGasMetering\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resumeGasMetering\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resumeTracing\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"revertTo\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertToAndDelete\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertToState\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revertToStateAndDelete\",\"inputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokePersistent\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokePersistent\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roll\",\"inputs\":[{\"name\":\"newHeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollFork\",\"inputs\":[{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollFork\",\"inputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollFork\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollFork\",\"inputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rpc\",\"inputs\":[{\"name\":\"urlOrAlias\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"method\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rpc\",\"inputs\":[{\"name\":\"method\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rpcUrl\",\"inputs\":[{\"name\":\"rpcAlias\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rpcUrlStructs\",\"inputs\":[],\"outputs\":[{\"name\":\"urls\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.Rpc[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"url\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rpcUrls\",\"inputs\":[],\"outputs\":[{\"name\":\"urls\",\"type\":\"string[2][]\",\"internalType\":\"string[2][]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"selectFork\",\"inputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeAddress\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeAddress\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBool\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBool\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes32\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeBytes32\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeInt\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeInt\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeJson\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeJsonType\",\"inputs\":[{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"serializeJsonType\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"typeDescription\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeString\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeString\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeUint\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeUint\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"serializeUintToHex\",\"inputs\":[{\"name\":\"objectKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setArbitraryStorage\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"overwrite\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setArbitraryStorage\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBlockhash\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEnv\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNonceUnsafe\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shuffle\",\"inputs\":[{\"name\":\"array\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sign\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signAndAttachDelegation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"signedDelegation\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.SignedDelegation\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signAndAttachDelegation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"signedDelegation\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.SignedDelegation\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.Wallet\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"publicKeyX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"publicKeyY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signCompact\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"vs\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"signDelegation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"signedDelegation\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.SignedDelegation\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signDelegation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"signedDelegation\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.SignedDelegation\",\"components\":[{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signP256\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"digest\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"skip\",\"inputs\":[{\"name\":\"skipTest\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"skip\",\"inputs\":[{\"name\":\"skipTest\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sleep\",\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"snapshot\",\"inputs\":[],\"outputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"snapshotGasLastCall\",\"inputs\":[{\"name\":\"group\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"snapshotGasLastCall\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"snapshotState\",\"inputs\":[],\"outputs\":[{\"name\":\"snapshotId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"snapshotValue\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"snapshotValue\",\"inputs\":[{\"name\":\"group\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sort\",\"inputs\":[{\"name\":\"array\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"split\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"delimiter\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"outputs\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"startBroadcast\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBroadcast\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBroadcast\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startDebugTraceRecording\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startMappingRecording\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startPrank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startPrank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startPrank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"txOrigin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startPrank\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"txOrigin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegateCall\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startSnapshotGas\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startSnapshotGas\",\"inputs\":[{\"name\":\"group\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startStateDiffRecording\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopAndReturnDebugTraceRecording\",\"inputs\":[],\"outputs\":[{\"name\":\"step\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.DebugStep[]\",\"components\":[{\"name\":\"stack\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"memoryInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"opcode\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isOutOfGas\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"contractAddr\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopAndReturnStateDiff\",\"inputs\":[],\"outputs\":[{\"name\":\"accountAccesses\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.AccountAccess[]\",\"components\":[{\"name\":\"chainInfo\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.ChainInfo\",\"components\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"kind\",\"type\":\"uint8\",\"internalType\":\"enumVmSafe.AccountAccessKind\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"accessor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialized\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"oldBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deployedCode\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"reverted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"storageAccesses\",\"type\":\"tuple[]\",\"internalType\":\"structVmSafe.StorageAccess[]\",\"components\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"isWrite\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"previousValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"newValue\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reverted\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"depth\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopBroadcast\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopExpectSafeMemory\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopMappingRecording\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopPrank\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopSnapshotGas\",\"inputs\":[{\"name\":\"group\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopSnapshotGas\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stopSnapshotGas\",\"inputs\":[],\"outputs\":[{\"name\":\"gasUsed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"store\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toBase64\",\"inputs\":[{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toBase64\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toBase64URL\",\"inputs\":[{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toBase64URL\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toLowercase\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toString\",\"inputs\":[{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"stringifiedValue\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"toUppercase\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transact\",\"inputs\":[{\"name\":\"forkId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transact\",\"inputs\":[{\"name\":\"txHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"trim\",\"inputs\":[{\"name\":\"input\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"tryFfi\",\"inputs\":[{\"name\":\"commandInput\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"tuple\",\"internalType\":\"structVmSafe.FfiResult\",\"components\":[{\"name\":\"exitCode\",\"type\":\"int32\",\"internalType\":\"int32\"},{\"name\":\"stdout\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"stderr\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"txGasPrice\",\"inputs\":[{\"name\":\"newGasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unixTime\",\"inputs\":[],\"outputs\":[{\"name\":\"milliseconds\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"warmSlot\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"warp\",\"inputs\":[{\"name\":\"newTimestamp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeFile\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeFileBinary\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeJson\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeLine\",\"inputs\":[{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"data\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeToml\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"valueKey\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"writeToml\",\"inputs\":[{\"name\":\"json\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"path\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// VmABI is the input ABI used to generate the binding from.
// Deprecated: Use VmMetaData.ABI instead.
var VmABI = VmMetaData.ABI

// Vm is an auto generated Go binding around an Ethereum contract.
type Vm struct {
	VmCaller     // Read-only binding to the contract
	VmTransactor // Write-only binding to the contract
	VmFilterer   // Log filterer for contract events
}

// VmCaller is an auto generated read-only Go binding around an Ethereum contract.
type VmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VmSession struct {
	Contract     *Vm               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VmCallerSession struct {
	Contract *VmCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VmTransactorSession struct {
	Contract     *VmTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VmRaw is an auto generated low-level Go binding around an Ethereum contract.
type VmRaw struct {
	Contract *Vm // Generic contract binding to access the raw methods on
}

// VmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VmCallerRaw struct {
	Contract *VmCaller // Generic read-only contract binding to access the raw methods on
}

// VmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VmTransactorRaw struct {
	Contract *VmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVm creates a new instance of Vm, bound to a specific deployed contract.
func NewVm(address common.Address, backend bind.ContractBackend) (*Vm, error) {
	contract, err := bindVm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vm{VmCaller: VmCaller{contract: contract}, VmTransactor: VmTransactor{contract: contract}, VmFilterer: VmFilterer{contract: contract}}, nil
}

// NewVmCaller creates a new read-only instance of Vm, bound to a specific deployed contract.
func NewVmCaller(address common.Address, caller bind.ContractCaller) (*VmCaller, error) {
	contract, err := bindVm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VmCaller{contract: contract}, nil
}

// NewVmTransactor creates a new write-only instance of Vm, bound to a specific deployed contract.
func NewVmTransactor(address common.Address, transactor bind.ContractTransactor) (*VmTransactor, error) {
	contract, err := bindVm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VmTransactor{contract: contract}, nil
}

// NewVmFilterer creates a new log filterer instance of Vm, bound to a specific deployed contract.
func NewVmFilterer(address common.Address, filterer bind.ContractFilterer) (*VmFilterer, error) {
	contract, err := bindVm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VmFilterer{contract: contract}, nil
}

// bindVm binds a generic wrapper to an already deployed contract.
func bindVm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vm *VmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vm.Contract.VmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vm *VmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.Contract.VmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vm *VmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vm.Contract.VmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vm *VmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vm *VmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vm *VmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vm.Contract.contract.Transact(opts, method, params...)
}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmCaller) ActiveFork(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "activeFork")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmSession) ActiveFork() (*big.Int, error) {
	return _Vm.Contract.ActiveFork(&_Vm.CallOpts)
}

// ActiveFork is a free data retrieval call binding the contract method 0x2f103f22.
//
// Solidity: function activeFork() view returns(uint256 forkId)
func (_Vm *VmCallerSession) ActiveFork() (*big.Int, error) {
	return _Vm.Contract.ActiveFork(&_Vm.CallOpts)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmCaller) Addr(opts *bind.CallOpts, privateKey *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "addr", privateKey)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _Vm.Contract.Addr(&_Vm.CallOpts, privateKey)
}

// Addr is a free data retrieval call binding the contract method 0xffa18649.
//
// Solidity: function addr(uint256 privateKey) pure returns(address keyAddr)
func (_Vm *VmCallerSession) Addr(privateKey *big.Int) (common.Address, error) {
	return _Vm.Contract.Addr(&_Vm.CallOpts, privateKey)
}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs", left, right, maxDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs is a free data retrieval call binding the contract method 0x16d207c6.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs0", left, right, maxDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs0(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs0(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs0 is a free data retrieval call binding the contract method 0x240f839d.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs0(left *big.Int, right *big.Int, maxDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbs0(&_Vm.CallOpts, left, right, maxDelta)
}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs1", left, right, maxDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs1(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs1(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs1 is a free data retrieval call binding the contract method 0x8289e621.
//
// Solidity: function assertApproxEqAbs(int256 left, int256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs1(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs1(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbs2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbs2", left, right, maxDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbs2(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs2(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbs2 is a free data retrieval call binding the contract method 0xf710b062.
//
// Solidity: function assertApproxEqAbs(uint256 left, uint256 right, uint256 maxDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbs2(left *big.Int, right *big.Int, maxDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbs2(&_Vm.CallOpts, left, right, maxDelta, error)
}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal", left, right, maxDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal is a free data retrieval call binding the contract method 0x045c55ce.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal0", left, right, maxDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal0(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal0(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal0 is a free data retrieval call binding the contract method 0x3d5bc8bc.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal0(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal0(&_Vm.CallOpts, left, right, maxDelta, decimals)
}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal1", left, right, maxDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal1(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal1(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal1 is a free data retrieval call binding the contract method 0x60429eb2.
//
// Solidity: function assertApproxEqAbsDecimal(uint256 left, uint256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal1(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal1(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqAbsDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqAbsDecimal2", left, right, maxDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqAbsDecimal2(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal2(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqAbsDecimal2 is a free data retrieval call binding the contract method 0x6a5066d4.
//
// Solidity: function assertApproxEqAbsDecimal(int256 left, int256 right, uint256 maxDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqAbsDecimal2(left *big.Int, right *big.Int, maxDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqAbsDecimal2(&_Vm.CallOpts, left, right, maxDelta, decimals, error)
}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel", left, right, maxPercentDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRel(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel is a free data retrieval call binding the contract method 0x1ecb7d33.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel0", left, right, maxPercentDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqRel0(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel0(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel0 is a free data retrieval call binding the contract method 0x8cf25ef4.
//
// Solidity: function assertApproxEqRel(uint256 left, uint256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel0(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel0(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel1", left, right, maxPercentDelta, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRel1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel1(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel1 is a free data retrieval call binding the contract method 0xef277d72.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRel1(&_Vm.CallOpts, left, right, maxPercentDelta, error)
}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCaller) AssertApproxEqRel2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRel2", left, right, maxPercentDelta)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmSession) AssertApproxEqRel2(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel2(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRel2 is a free data retrieval call binding the contract method 0xfea2d14f.
//
// Solidity: function assertApproxEqRel(int256 left, int256 right, uint256 maxPercentDelta) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRel2(left *big.Int, right *big.Int, maxPercentDelta *big.Int) error {
	return _Vm.Contract.AssertApproxEqRel2(&_Vm.CallOpts, left, right, maxPercentDelta)
}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal", left, right, maxPercentDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal is a free data retrieval call binding the contract method 0x21ed2977.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal0", left, right, maxPercentDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal0(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal0(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal0 is a free data retrieval call binding the contract method 0x82d6c8fd.
//
// Solidity: function assertApproxEqRelDecimal(uint256 left, uint256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal0(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal0(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal1", left, right, maxPercentDelta, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal1(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal1 is a free data retrieval call binding the contract method 0xabbf21cc.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal1(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertApproxEqRelDecimal1(&_Vm.CallOpts, left, right, maxPercentDelta, decimals)
}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertApproxEqRelDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertApproxEqRelDecimal2", left, right, maxPercentDelta, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertApproxEqRelDecimal2(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal2(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertApproxEqRelDecimal2 is a free data retrieval call binding the contract method 0xfccc11c4.
//
// Solidity: function assertApproxEqRelDecimal(int256 left, int256 right, uint256 maxPercentDelta, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertApproxEqRelDecimal2(left *big.Int, right *big.Int, maxPercentDelta *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertApproxEqRelDecimal2(&_Vm.CallOpts, left, right, maxPercentDelta, decimals, error)
}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCaller) AssertEq(opts *bind.CallOpts, left [][32]byte, right [][32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmSession) AssertEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertEq(&_Vm.CallOpts, left, right)
}

// AssertEq is a free data retrieval call binding the contract method 0x0cc9ee84.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertEq(&_Vm.CallOpts, left, right)
}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq0(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq0(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq0(&_Vm.CallOpts, left, right, error)
}

// AssertEq0 is a free data retrieval call binding the contract method 0x191f1b30.
//
// Solidity: function assertEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq0(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq0(&_Vm.CallOpts, left, right, error)
}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_Vm *VmCaller) AssertEq1(opts *bind.CallOpts, left common.Address, right common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_Vm *VmSession) AssertEq1(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertEq1(&_Vm.CallOpts, left, right, error)
}

// AssertEq1 is a free data retrieval call binding the contract method 0x2f2769d1.
//
// Solidity: function assertEq(address left, address right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq1(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertEq1(&_Vm.CallOpts, left, right, error)
}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertEq10(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq10", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertEq10(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq10(&_Vm.CallOpts, left, right, error)
}

// AssertEq10 is a free data retrieval call binding the contract method 0x714a2f13.
//
// Solidity: function assertEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq10(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq10(&_Vm.CallOpts, left, right, error)
}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCaller) AssertEq11(opts *bind.CallOpts, left [32]byte, right [32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq11", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmSession) AssertEq11(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertEq11(&_Vm.CallOpts, left, right)
}

// AssertEq11 is a free data retrieval call binding the contract method 0x7c84c69b.
//
// Solidity: function assertEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCallerSession) AssertEq11(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertEq11(&_Vm.CallOpts, left, right)
}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertEq12(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq12", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertEq12(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq12(&_Vm.CallOpts, left, right, error)
}

// AssertEq12 is a free data retrieval call binding the contract method 0x88b44c85.
//
// Solidity: function assertEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq12(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertEq12(&_Vm.CallOpts, left, right, error)
}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCaller) AssertEq13(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq13", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmSession) AssertEq13(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq13(&_Vm.CallOpts, left, right)
}

// AssertEq13 is a free data retrieval call binding the contract method 0x975d5a12.
//
// Solidity: function assertEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq13(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq13(&_Vm.CallOpts, left, right)
}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_Vm *VmCaller) AssertEq14(opts *bind.CallOpts, left []byte, right []byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq14", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_Vm *VmSession) AssertEq14(left []byte, right []byte) error {
	return _Vm.Contract.AssertEq14(&_Vm.CallOpts, left, right)
}

// AssertEq14 is a free data retrieval call binding the contract method 0x97624631.
//
// Solidity: function assertEq(bytes left, bytes right) pure returns()
func (_Vm *VmCallerSession) AssertEq14(left []byte, right []byte) error {
	return _Vm.Contract.AssertEq14(&_Vm.CallOpts, left, right)
}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertEq15(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq15", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertEq15(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq15(&_Vm.CallOpts, left, right)
}

// AssertEq15 is a free data retrieval call binding the contract method 0x98296c54.
//
// Solidity: function assertEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertEq15(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq15(&_Vm.CallOpts, left, right)
}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCaller) AssertEq16(opts *bind.CallOpts, left [32]byte, right [32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq16", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmSession) AssertEq16(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertEq16(&_Vm.CallOpts, left, right, error)
}

// AssertEq16 is a free data retrieval call binding the contract method 0xc1fa1ed0.
//
// Solidity: function assertEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq16(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertEq16(&_Vm.CallOpts, left, right, error)
}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_Vm *VmCaller) AssertEq17(opts *bind.CallOpts, left []string, right []string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq17", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_Vm *VmSession) AssertEq17(left []string, right []string) error {
	return _Vm.Contract.AssertEq17(&_Vm.CallOpts, left, right)
}

// AssertEq17 is a free data retrieval call binding the contract method 0xcf1c049c.
//
// Solidity: function assertEq(string[] left, string[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq17(left []string, right []string) error {
	return _Vm.Contract.AssertEq17(&_Vm.CallOpts, left, right)
}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq18(opts *bind.CallOpts, left [][32]byte, right [][32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq18", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq18(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertEq18(&_Vm.CallOpts, left, right, error)
}

// AssertEq18 is a free data retrieval call binding the contract method 0xe03e9177.
//
// Solidity: function assertEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq18(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertEq18(&_Vm.CallOpts, left, right, error)
}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCaller) AssertEq19(opts *bind.CallOpts, left []byte, right []byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq19", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmSession) AssertEq19(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertEq19(&_Vm.CallOpts, left, right, error)
}

// AssertEq19 is a free data retrieval call binding the contract method 0xe24fed00.
//
// Solidity: function assertEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq19(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertEq19(&_Vm.CallOpts, left, right, error)
}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_Vm *VmCaller) AssertEq2(opts *bind.CallOpts, left string, right string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_Vm *VmSession) AssertEq2(left string, right string, error string) error {
	return _Vm.Contract.AssertEq2(&_Vm.CallOpts, left, right, error)
}

// AssertEq2 is a free data retrieval call binding the contract method 0x36f656d8.
//
// Solidity: function assertEq(string left, string right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq2(left string, right string, error string) error {
	return _Vm.Contract.AssertEq2(&_Vm.CallOpts, left, right, error)
}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq20(opts *bind.CallOpts, left []bool, right []bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq20", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq20(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertEq20(&_Vm.CallOpts, left, right, error)
}

// AssertEq20 is a free data retrieval call binding the contract method 0xe48a8f8d.
//
// Solidity: function assertEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq20(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertEq20(&_Vm.CallOpts, left, right, error)
}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCaller) AssertEq21(opts *bind.CallOpts, left [][]byte, right [][]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq21", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmSession) AssertEq21(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertEq21(&_Vm.CallOpts, left, right)
}

// AssertEq21 is a free data retrieval call binding the contract method 0xe5fb9b4a.
//
// Solidity: function assertEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq21(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertEq21(&_Vm.CallOpts, left, right)
}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq22(opts *bind.CallOpts, left []string, right []string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq22", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq22(left []string, right []string, error string) error {
	return _Vm.Contract.AssertEq22(&_Vm.CallOpts, left, right, error)
}

// AssertEq22 is a free data retrieval call binding the contract method 0xeff6b27d.
//
// Solidity: function assertEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq22(left []string, right []string, error string) error {
	return _Vm.Contract.AssertEq22(&_Vm.CallOpts, left, right, error)
}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_Vm *VmCaller) AssertEq23(opts *bind.CallOpts, left string, right string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq23", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_Vm *VmSession) AssertEq23(left string, right string) error {
	return _Vm.Contract.AssertEq23(&_Vm.CallOpts, left, right)
}

// AssertEq23 is a free data retrieval call binding the contract method 0xf320d963.
//
// Solidity: function assertEq(string left, string right) pure returns()
func (_Vm *VmCallerSession) AssertEq23(left string, right string) error {
	return _Vm.Contract.AssertEq23(&_Vm.CallOpts, left, right)
}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq24(opts *bind.CallOpts, left [][]byte, right [][]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq24", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq24(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertEq24(&_Vm.CallOpts, left, right, error)
}

// AssertEq24 is a free data retrieval call binding the contract method 0xf413f0b6.
//
// Solidity: function assertEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq24(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertEq24(&_Vm.CallOpts, left, right, error)
}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_Vm *VmCaller) AssertEq25(opts *bind.CallOpts, left bool, right bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq25", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_Vm *VmSession) AssertEq25(left bool, right bool) error {
	return _Vm.Contract.AssertEq25(&_Vm.CallOpts, left, right)
}

// AssertEq25 is a free data retrieval call binding the contract method 0xf7fe3477.
//
// Solidity: function assertEq(bool left, bool right) pure returns()
func (_Vm *VmCallerSession) AssertEq25(left bool, right bool) error {
	return _Vm.Contract.AssertEq25(&_Vm.CallOpts, left, right)
}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertEq26(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq26", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq26(&_Vm.CallOpts, left, right)
}

// AssertEq26 is a free data retrieval call binding the contract method 0xfe74f05b.
//
// Solidity: function assertEq(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertEq26(&_Vm.CallOpts, left, right)
}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_Vm *VmCaller) AssertEq3(opts *bind.CallOpts, left []common.Address, right []common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq3", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_Vm *VmSession) AssertEq3(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertEq3(&_Vm.CallOpts, left, right)
}

// AssertEq3 is a free data retrieval call binding the contract method 0x3868ac34.
//
// Solidity: function assertEq(address[] left, address[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq3(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertEq3(&_Vm.CallOpts, left, right)
}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq4(opts *bind.CallOpts, left []common.Address, right []common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq4", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq4(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertEq4(&_Vm.CallOpts, left, right, error)
}

// AssertEq4 is a free data retrieval call binding the contract method 0x3e9173c5.
//
// Solidity: function assertEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq4(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertEq4(&_Vm.CallOpts, left, right, error)
}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCaller) AssertEq5(opts *bind.CallOpts, left bool, right bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq5", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_Vm *VmSession) AssertEq5(left bool, right bool, error string) error {
	return _Vm.Contract.AssertEq5(&_Vm.CallOpts, left, right, error)
}

// AssertEq5 is a free data retrieval call binding the contract method 0x4db19e7e.
//
// Solidity: function assertEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq5(left bool, right bool, error string) error {
	return _Vm.Contract.AssertEq5(&_Vm.CallOpts, left, right, error)
}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_Vm *VmCaller) AssertEq6(opts *bind.CallOpts, left common.Address, right common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq6", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_Vm *VmSession) AssertEq6(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertEq6(&_Vm.CallOpts, left, right)
}

// AssertEq6 is a free data retrieval call binding the contract method 0x515361f6.
//
// Solidity: function assertEq(address left, address right) pure returns()
func (_Vm *VmCallerSession) AssertEq6(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertEq6(&_Vm.CallOpts, left, right)
}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertEq7(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq7", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmSession) AssertEq7(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq7(&_Vm.CallOpts, left, right, error)
}

// AssertEq7 is a free data retrieval call binding the contract method 0x5d18c73a.
//
// Solidity: function assertEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertEq7(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertEq7(&_Vm.CallOpts, left, right, error)
}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCaller) AssertEq8(opts *bind.CallOpts, left []bool, right []bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq8", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmSession) AssertEq8(left []bool, right []bool) error {
	return _Vm.Contract.AssertEq8(&_Vm.CallOpts, left, right)
}

// AssertEq8 is a free data retrieval call binding the contract method 0x707df785.
//
// Solidity: function assertEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq8(left []bool, right []bool) error {
	return _Vm.Contract.AssertEq8(&_Vm.CallOpts, left, right)
}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCaller) AssertEq9(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEq9", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmSession) AssertEq9(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq9(&_Vm.CallOpts, left, right)
}

// AssertEq9 is a free data retrieval call binding the contract method 0x711043ac.
//
// Solidity: function assertEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCallerSession) AssertEq9(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertEq9(&_Vm.CallOpts, left, right)
}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertEqDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal is a free data retrieval call binding the contract method 0x27af7d9c.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertEqDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal0", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal0(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal0 is a free data retrieval call binding the contract method 0x48016c04.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertEqDecimal0(&_Vm.CallOpts, left, right, decimals)
}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertEqDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal1 is a free data retrieval call binding the contract method 0x7e77b0c5.
//
// Solidity: function assertEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertEqDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertEqDecimal2", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertEqDecimal2 is a free data retrieval call binding the contract method 0xd0cbbdef.
//
// Solidity: function assertEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_Vm *VmCaller) AssertFalse(opts *bind.CallOpts, condition bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertFalse", condition, error)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_Vm *VmSession) AssertFalse(condition bool, error string) error {
	return _Vm.Contract.AssertFalse(&_Vm.CallOpts, condition, error)
}

// AssertFalse is a free data retrieval call binding the contract method 0x7ba04809.
//
// Solidity: function assertFalse(bool condition, string error) pure returns()
func (_Vm *VmCallerSession) AssertFalse(condition bool, error string) error {
	return _Vm.Contract.AssertFalse(&_Vm.CallOpts, condition, error)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_Vm *VmCaller) AssertFalse0(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertFalse0", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_Vm *VmSession) AssertFalse0(condition bool) error {
	return _Vm.Contract.AssertFalse0(&_Vm.CallOpts, condition)
}

// AssertFalse0 is a free data retrieval call binding the contract method 0xa5982885.
//
// Solidity: function assertFalse(bool condition) pure returns()
func (_Vm *VmCallerSession) AssertFalse0(condition bool) error {
	return _Vm.Contract.AssertFalse0(&_Vm.CallOpts, condition)
}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertGe(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertGe(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe(&_Vm.CallOpts, left, right)
}

// AssertGe is a free data retrieval call binding the contract method 0x0a30b771.
//
// Solidity: function assertGe(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertGe(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe(&_Vm.CallOpts, left, right)
}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGe0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertGe0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe0(&_Vm.CallOpts, left, right, error)
}

// AssertGe0 is a free data retrieval call binding the contract method 0xa84328dd.
//
// Solidity: function assertGe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGe0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe0(&_Vm.CallOpts, left, right, error)
}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertGe1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertGe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe1(&_Vm.CallOpts, left, right)
}

// AssertGe1 is a free data retrieval call binding the contract method 0xa8d4d1d9.
//
// Solidity: function assertGe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertGe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGe1(&_Vm.CallOpts, left, right)
}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGe2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGe2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertGe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe2(&_Vm.CallOpts, left, right, error)
}

// AssertGe2 is a free data retrieval call binding the contract method 0xe25242c0.
//
// Solidity: function assertGe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGe2(&_Vm.CallOpts, left, right, error)
}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGeDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertGeDecimal is a free data retrieval call binding the contract method 0x3d1fe08a.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGeDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal0 is a free data retrieval call binding the contract method 0x5df93c9b.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGeDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal1 is a free data retrieval call binding the contract method 0x8bff9133.
//
// Solidity: function assertGeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGeDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGeDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertGeDecimal2 is a free data retrieval call binding the contract method 0xdc28c0f1.
//
// Solidity: function assertGeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertGt(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertGt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt(&_Vm.CallOpts, left, right)
}

// AssertGt is a free data retrieval call binding the contract method 0x5a362d45.
//
// Solidity: function assertGt(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertGt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt(&_Vm.CallOpts, left, right)
}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGt0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertGt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt0(&_Vm.CallOpts, left, right, error)
}

// AssertGt0 is a free data retrieval call binding the contract method 0xd9a3c4d2.
//
// Solidity: function assertGt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt0(&_Vm.CallOpts, left, right, error)
}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertGt1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertGt1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt1(&_Vm.CallOpts, left, right)
}

// AssertGt1 is a free data retrieval call binding the contract method 0xdb07fcd2.
//
// Solidity: function assertGt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertGt1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertGt1(&_Vm.CallOpts, left, right)
}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertGt2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGt2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertGt2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt2(&_Vm.CallOpts, left, right, error)
}

// AssertGt2 is a free data retrieval call binding the contract method 0xf8d33b9b.
//
// Solidity: function assertGt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertGt2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertGt2(&_Vm.CallOpts, left, right, error)
}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGtDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGtDecimal(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal is a free data retrieval call binding the contract method 0x04a5c7ab.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertGtDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertGtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal0 is a free data retrieval call binding the contract method 0x64949a8d.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertGtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGtDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal1", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGtDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertGtDecimal1 is a free data retrieval call binding the contract method 0x78611f0e.
//
// Solidity: function assertGtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertGtDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertGtDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertGtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertGtDecimal2 is a free data retrieval call binding the contract method 0xeccd2437.
//
// Solidity: function assertGtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertGtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertGtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLe(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertLe(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe(&_Vm.CallOpts, left, right, error)
}

// AssertLe is a free data retrieval call binding the contract method 0x4dfe692c.
//
// Solidity: function assertLe(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLe(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe(&_Vm.CallOpts, left, right, error)
}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertLe0(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe0", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertLe0(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe0(&_Vm.CallOpts, left, right)
}

// AssertLe0 is a free data retrieval call binding the contract method 0x8466f415.
//
// Solidity: function assertLe(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertLe0(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe0(&_Vm.CallOpts, left, right)
}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertLe1(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe1", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertLe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe1(&_Vm.CallOpts, left, right)
}

// AssertLe1 is a free data retrieval call binding the contract method 0x95fd154e.
//
// Solidity: function assertLe(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertLe1(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLe1(&_Vm.CallOpts, left, right)
}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLe2(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLe2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertLe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe2(&_Vm.CallOpts, left, right, error)
}

// AssertLe2 is a free data retrieval call binding the contract method 0xd17d4b0d.
//
// Solidity: function assertLe(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLe2(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLe2(&_Vm.CallOpts, left, right, error)
}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLeDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLeDecimal is a free data retrieval call binding the contract method 0x11d1364a.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLeDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal0 is a free data retrieval call binding the contract method 0x7fefbbe0.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLeDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal1 is a free data retrieval call binding the contract method 0xaa5cf788.
//
// Solidity: function assertLeDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLeDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLeDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLeDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLeDecimal2 is a free data retrieval call binding the contract method 0xc304aab7.
//
// Solidity: function assertLeDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLeDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLeDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertLt(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertLt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt(&_Vm.CallOpts, left, right)
}

// AssertLt is a free data retrieval call binding the contract method 0x3e914080.
//
// Solidity: function assertLt(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertLt(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt(&_Vm.CallOpts, left, right)
}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLt0(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt0", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertLt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt0(&_Vm.CallOpts, left, right, error)
}

// AssertLt0 is a free data retrieval call binding the contract method 0x65d5c135.
//
// Solidity: function assertLt(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLt0(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt0(&_Vm.CallOpts, left, right, error)
}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertLt1(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertLt1(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt1(&_Vm.CallOpts, left, right, error)
}

// AssertLt1 is a free data retrieval call binding the contract method 0x9ff531e3.
//
// Solidity: function assertLt(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertLt1(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertLt1(&_Vm.CallOpts, left, right, error)
}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertLt2(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLt2", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertLt2(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt2(&_Vm.CallOpts, left, right)
}

// AssertLt2 is a free data retrieval call binding the contract method 0xb12fc005.
//
// Solidity: function assertLt(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertLt2(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertLt2(&_Vm.CallOpts, left, right)
}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLtDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLtDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLtDecimal is a free data retrieval call binding the contract method 0x2077337e.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLtDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal0 is a free data retrieval call binding the contract method 0x40f0b4e0.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertLtDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal1", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertLtDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal1 is a free data retrieval call binding the contract method 0xa972d037.
//
// Solidity: function assertLtDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal1(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertLtDecimal1(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertLtDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertLtDecimal2", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertLtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertLtDecimal2 is a free data retrieval call binding the contract method 0xdbe8d88b.
//
// Solidity: function assertLtDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertLtDecimal2(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertLtDecimal2(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq(opts *bind.CallOpts, left [][32]byte, right [][32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmSession) AssertNotEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertNotEq(&_Vm.CallOpts, left, right)
}

// AssertNotEq is a free data retrieval call binding the contract method 0x0603ea68.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq(left [][32]byte, right [][32]byte) error {
	return _Vm.Contract.AssertNotEq(&_Vm.CallOpts, left, right)
}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq0(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq0", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmSession) AssertNotEq0(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq0(&_Vm.CallOpts, left, right)
}

// AssertNotEq0 is a free data retrieval call binding the contract method 0x0b72f4ef.
//
// Solidity: function assertNotEq(int256[] left, int256[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq0(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq0(&_Vm.CallOpts, left, right)
}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq1(opts *bind.CallOpts, left bool, right bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq1", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq1(left bool, right bool, error string) error {
	return _Vm.Contract.AssertNotEq1(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq1 is a free data retrieval call binding the contract method 0x1091a261.
//
// Solidity: function assertNotEq(bool left, bool right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq1(left bool, right bool, error string) error {
	return _Vm.Contract.AssertNotEq1(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_Vm *VmCaller) AssertNotEq10(opts *bind.CallOpts, left string, right string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq10", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_Vm *VmSession) AssertNotEq10(left string, right string) error {
	return _Vm.Contract.AssertNotEq10(&_Vm.CallOpts, left, right)
}

// AssertNotEq10 is a free data retrieval call binding the contract method 0x6a8237b3.
//
// Solidity: function assertNotEq(string left, string right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq10(left string, right string) error {
	return _Vm.Contract.AssertNotEq10(&_Vm.CallOpts, left, right)
}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq11(opts *bind.CallOpts, left []common.Address, right []common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq11", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq11(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertNotEq11(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq11 is a free data retrieval call binding the contract method 0x72c7e0b5.
//
// Solidity: function assertNotEq(address[] left, address[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq11(left []common.Address, right []common.Address, error string) error {
	return _Vm.Contract.AssertNotEq11(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq12(opts *bind.CallOpts, left string, right string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq12", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq12(left string, right string, error string) error {
	return _Vm.Contract.AssertNotEq12(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq12 is a free data retrieval call binding the contract method 0x78bdcea7.
//
// Solidity: function assertNotEq(string left, string right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq12(left string, right string, error string) error {
	return _Vm.Contract.AssertNotEq12(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq13(opts *bind.CallOpts, left common.Address, right common.Address, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq13", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq13(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertNotEq13(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq13 is a free data retrieval call binding the contract method 0x8775a591.
//
// Solidity: function assertNotEq(address left, address right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq13(left common.Address, right common.Address, error string) error {
	return _Vm.Contract.AssertNotEq13(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCaller) AssertNotEq14(opts *bind.CallOpts, left [32]byte, right [32]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq14", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmSession) AssertNotEq14(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertNotEq14(&_Vm.CallOpts, left, right)
}

// AssertNotEq14 is a free data retrieval call binding the contract method 0x898e83fc.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq14(left [32]byte, right [32]byte) error {
	return _Vm.Contract.AssertNotEq14(&_Vm.CallOpts, left, right)
}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq15(opts *bind.CallOpts, left []byte, right []byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq15", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq15(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertNotEq15(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq15 is a free data retrieval call binding the contract method 0x9507540e.
//
// Solidity: function assertNotEq(bytes left, bytes right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq15(left []byte, right []byte, error string) error {
	return _Vm.Contract.AssertNotEq15(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq16(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq16", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq16(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq16(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq16 is a free data retrieval call binding the contract method 0x98f9bdbd.
//
// Solidity: function assertNotEq(uint256 left, uint256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq16(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq16(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq17(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq17", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq17(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq17(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq17 is a free data retrieval call binding the contract method 0x9a7fbd8f.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq17(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq17(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_Vm *VmCaller) AssertNotEq18(opts *bind.CallOpts, left common.Address, right common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq18", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_Vm *VmSession) AssertNotEq18(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertNotEq18(&_Vm.CallOpts, left, right)
}

// AssertNotEq18 is a free data retrieval call binding the contract method 0xb12e1694.
//
// Solidity: function assertNotEq(address left, address right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq18(left common.Address, right common.Address) error {
	return _Vm.Contract.AssertNotEq18(&_Vm.CallOpts, left, right)
}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq19(opts *bind.CallOpts, left [32]byte, right [32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq19", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq19(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertNotEq19(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq19 is a free data retrieval call binding the contract method 0xb2332f51.
//
// Solidity: function assertNotEq(bytes32 left, bytes32 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq19(left [32]byte, right [32]byte, error string) error {
	return _Vm.Contract.AssertNotEq19(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq2(opts *bind.CallOpts, left [][]byte, right [][]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq2", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq2(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertNotEq2(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq2 is a free data retrieval call binding the contract method 0x1dcd1f68.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq2(left [][]byte, right [][]byte, error string) error {
	return _Vm.Contract.AssertNotEq2(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq20(opts *bind.CallOpts, left []string, right []string, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq20", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq20(left []string, right []string, error string) error {
	return _Vm.Contract.AssertNotEq20(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq20 is a free data retrieval call binding the contract method 0xb67187f3.
//
// Solidity: function assertNotEq(string[] left, string[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq20(left []string, right []string, error string) error {
	return _Vm.Contract.AssertNotEq20(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCaller) AssertNotEq21(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq21", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmSession) AssertNotEq21(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq21(&_Vm.CallOpts, left, right)
}

// AssertNotEq21 is a free data retrieval call binding the contract method 0xb7909320.
//
// Solidity: function assertNotEq(uint256 left, uint256 right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq21(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq21(&_Vm.CallOpts, left, right)
}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq22(opts *bind.CallOpts, left [][32]byte, right [][32]byte, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq22", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq22(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertNotEq22(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq22 is a free data retrieval call binding the contract method 0xb873634c.
//
// Solidity: function assertNotEq(bytes32[] left, bytes32[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq22(left [][32]byte, right [][32]byte, error string) error {
	return _Vm.Contract.AssertNotEq22(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq23(opts *bind.CallOpts, left []string, right []string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq23", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_Vm *VmSession) AssertNotEq23(left []string, right []string) error {
	return _Vm.Contract.AssertNotEq23(&_Vm.CallOpts, left, right)
}

// AssertNotEq23 is a free data retrieval call binding the contract method 0xbdfacbe8.
//
// Solidity: function assertNotEq(string[] left, string[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq23(left []string, right []string) error {
	return _Vm.Contract.AssertNotEq23(&_Vm.CallOpts, left, right)
}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq24(opts *bind.CallOpts, left []*big.Int, right []*big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq24", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq24(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq24(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq24 is a free data retrieval call binding the contract method 0xd3977322.
//
// Solidity: function assertNotEq(int256[] left, int256[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq24(left []*big.Int, right []*big.Int, error string) error {
	return _Vm.Contract.AssertNotEq24(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq25(opts *bind.CallOpts, left [][]byte, right [][]byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq25", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmSession) AssertNotEq25(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertNotEq25(&_Vm.CallOpts, left, right)
}

// AssertNotEq25 is a free data retrieval call binding the contract method 0xedecd035.
//
// Solidity: function assertNotEq(bytes[] left, bytes[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq25(left [][]byte, right [][]byte) error {
	return _Vm.Contract.AssertNotEq25(&_Vm.CallOpts, left, right)
}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_Vm *VmCaller) AssertNotEq26(opts *bind.CallOpts, left *big.Int, right *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq26", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_Vm *VmSession) AssertNotEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq26(&_Vm.CallOpts, left, right)
}

// AssertNotEq26 is a free data retrieval call binding the contract method 0xf4c004e3.
//
// Solidity: function assertNotEq(int256 left, int256 right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq26(left *big.Int, right *big.Int) error {
	return _Vm.Contract.AssertNotEq26(&_Vm.CallOpts, left, right)
}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_Vm *VmCaller) AssertNotEq3(opts *bind.CallOpts, left bool, right bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq3", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_Vm *VmSession) AssertNotEq3(left bool, right bool) error {
	return _Vm.Contract.AssertNotEq3(&_Vm.CallOpts, left, right)
}

// AssertNotEq3 is a free data retrieval call binding the contract method 0x236e4d66.
//
// Solidity: function assertNotEq(bool left, bool right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq3(left bool, right bool) error {
	return _Vm.Contract.AssertNotEq3(&_Vm.CallOpts, left, right)
}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq4(opts *bind.CallOpts, left []bool, right []bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq4", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmSession) AssertNotEq4(left []bool, right []bool) error {
	return _Vm.Contract.AssertNotEq4(&_Vm.CallOpts, left, right)
}

// AssertNotEq4 is a free data retrieval call binding the contract method 0x286fafea.
//
// Solidity: function assertNotEq(bool[] left, bool[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq4(left []bool, right []bool) error {
	return _Vm.Contract.AssertNotEq4(&_Vm.CallOpts, left, right)
}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_Vm *VmCaller) AssertNotEq5(opts *bind.CallOpts, left []byte, right []byte) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq5", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_Vm *VmSession) AssertNotEq5(left []byte, right []byte) error {
	return _Vm.Contract.AssertNotEq5(&_Vm.CallOpts, left, right)
}

// AssertNotEq5 is a free data retrieval call binding the contract method 0x3cf78e28.
//
// Solidity: function assertNotEq(bytes left, bytes right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq5(left []byte, right []byte) error {
	return _Vm.Contract.AssertNotEq5(&_Vm.CallOpts, left, right)
}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq6(opts *bind.CallOpts, left []common.Address, right []common.Address) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq6", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_Vm *VmSession) AssertNotEq6(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertNotEq6(&_Vm.CallOpts, left, right)
}

// AssertNotEq6 is a free data retrieval call binding the contract method 0x46d0b252.
//
// Solidity: function assertNotEq(address[] left, address[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq6(left []common.Address, right []common.Address) error {
	return _Vm.Contract.AssertNotEq6(&_Vm.CallOpts, left, right)
}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq7(opts *bind.CallOpts, left *big.Int, right *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq7", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq7(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq7(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq7 is a free data retrieval call binding the contract method 0x4724c5b9.
//
// Solidity: function assertNotEq(int256 left, int256 right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq7(left *big.Int, right *big.Int, error string) error {
	return _Vm.Contract.AssertNotEq7(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCaller) AssertNotEq8(opts *bind.CallOpts, left []*big.Int, right []*big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq8", left, right)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmSession) AssertNotEq8(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq8(&_Vm.CallOpts, left, right)
}

// AssertNotEq8 is a free data retrieval call binding the contract method 0x56f29cba.
//
// Solidity: function assertNotEq(uint256[] left, uint256[] right) pure returns()
func (_Vm *VmCallerSession) AssertNotEq8(left []*big.Int, right []*big.Int) error {
	return _Vm.Contract.AssertNotEq8(&_Vm.CallOpts, left, right)
}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCaller) AssertNotEq9(opts *bind.CallOpts, left []bool, right []bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEq9", left, right, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmSession) AssertNotEq9(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertNotEq9(&_Vm.CallOpts, left, right, error)
}

// AssertNotEq9 is a free data retrieval call binding the contract method 0x62c6f9fb.
//
// Solidity: function assertNotEq(bool[] left, bool[] right, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEq9(left []bool, right []bool, error string) error {
	return _Vm.Contract.AssertNotEq9(&_Vm.CallOpts, left, right, error)
}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal is a free data retrieval call binding the contract method 0x14e75680.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal0(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal0", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal0 is a free data retrieval call binding the contract method 0x33949f0b.
//
// Solidity: function assertNotEqDecimal(int256 left, int256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal0(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal0(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal1(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal1", left, right, decimals)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal1 is a free data retrieval call binding the contract method 0x669efca7.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal1(left *big.Int, right *big.Int, decimals *big.Int) error {
	return _Vm.Contract.AssertNotEqDecimal1(&_Vm.CallOpts, left, right, decimals)
}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCaller) AssertNotEqDecimal2(opts *bind.CallOpts, left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertNotEqDecimal2", left, right, decimals, error)

	if err != nil {
		return err
	}

	return err

}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmSession) AssertNotEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertNotEqDecimal2 is a free data retrieval call binding the contract method 0xf5a55558.
//
// Solidity: function assertNotEqDecimal(uint256 left, uint256 right, uint256 decimals, string error) pure returns()
func (_Vm *VmCallerSession) AssertNotEqDecimal2(left *big.Int, right *big.Int, decimals *big.Int, error string) error {
	return _Vm.Contract.AssertNotEqDecimal2(&_Vm.CallOpts, left, right, decimals, error)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_Vm *VmCaller) AssertTrue(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertTrue", condition)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_Vm *VmSession) AssertTrue(condition bool) error {
	return _Vm.Contract.AssertTrue(&_Vm.CallOpts, condition)
}

// AssertTrue is a free data retrieval call binding the contract method 0x0c9fd581.
//
// Solidity: function assertTrue(bool condition) pure returns()
func (_Vm *VmCallerSession) AssertTrue(condition bool) error {
	return _Vm.Contract.AssertTrue(&_Vm.CallOpts, condition)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_Vm *VmCaller) AssertTrue0(opts *bind.CallOpts, condition bool, error string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assertTrue0", condition, error)

	if err != nil {
		return err
	}

	return err

}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_Vm *VmSession) AssertTrue0(condition bool, error string) error {
	return _Vm.Contract.AssertTrue0(&_Vm.CallOpts, condition, error)
}

// AssertTrue0 is a free data retrieval call binding the contract method 0xa34edc03.
//
// Solidity: function assertTrue(bool condition, string error) pure returns()
func (_Vm *VmCallerSession) AssertTrue0(condition bool, error string) error {
	return _Vm.Contract.AssertTrue0(&_Vm.CallOpts, condition, error)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmCaller) Assume(opts *bind.CallOpts, condition bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assume", condition)

	if err != nil {
		return err
	}

	return err

}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmSession) Assume(condition bool) error {
	return _Vm.Contract.Assume(&_Vm.CallOpts, condition)
}

// Assume is a free data retrieval call binding the contract method 0x4c63e562.
//
// Solidity: function assume(bool condition) pure returns()
func (_Vm *VmCallerSession) Assume(condition bool) error {
	return _Vm.Contract.Assume(&_Vm.CallOpts, condition)
}

// AssumeNoRevert is a free data retrieval call binding the contract method 0x285b366a.
//
// Solidity: function assumeNoRevert() pure returns()
func (_Vm *VmCaller) AssumeNoRevert(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assumeNoRevert")

	if err != nil {
		return err
	}

	return err

}

// AssumeNoRevert is a free data retrieval call binding the contract method 0x285b366a.
//
// Solidity: function assumeNoRevert() pure returns()
func (_Vm *VmSession) AssumeNoRevert() error {
	return _Vm.Contract.AssumeNoRevert(&_Vm.CallOpts)
}

// AssumeNoRevert is a free data retrieval call binding the contract method 0x285b366a.
//
// Solidity: function assumeNoRevert() pure returns()
func (_Vm *VmCallerSession) AssumeNoRevert() error {
	return _Vm.Contract.AssumeNoRevert(&_Vm.CallOpts)
}

// AssumeNoRevert0 is a free data retrieval call binding the contract method 0x8a4592cc.
//
// Solidity: function assumeNoRevert((address,bool,bytes)[] potentialReverts) pure returns()
func (_Vm *VmCaller) AssumeNoRevert0(opts *bind.CallOpts, potentialReverts []VmSafePotentialRevert) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assumeNoRevert0", potentialReverts)

	if err != nil {
		return err
	}

	return err

}

// AssumeNoRevert0 is a free data retrieval call binding the contract method 0x8a4592cc.
//
// Solidity: function assumeNoRevert((address,bool,bytes)[] potentialReverts) pure returns()
func (_Vm *VmSession) AssumeNoRevert0(potentialReverts []VmSafePotentialRevert) error {
	return _Vm.Contract.AssumeNoRevert0(&_Vm.CallOpts, potentialReverts)
}

// AssumeNoRevert0 is a free data retrieval call binding the contract method 0x8a4592cc.
//
// Solidity: function assumeNoRevert((address,bool,bytes)[] potentialReverts) pure returns()
func (_Vm *VmCallerSession) AssumeNoRevert0(potentialReverts []VmSafePotentialRevert) error {
	return _Vm.Contract.AssumeNoRevert0(&_Vm.CallOpts, potentialReverts)
}

// AssumeNoRevert1 is a free data retrieval call binding the contract method 0xd8591eeb.
//
// Solidity: function assumeNoRevert((address,bool,bytes) potentialRevert) pure returns()
func (_Vm *VmCaller) AssumeNoRevert1(opts *bind.CallOpts, potentialRevert VmSafePotentialRevert) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "assumeNoRevert1", potentialRevert)

	if err != nil {
		return err
	}

	return err

}

// AssumeNoRevert1 is a free data retrieval call binding the contract method 0xd8591eeb.
//
// Solidity: function assumeNoRevert((address,bool,bytes) potentialRevert) pure returns()
func (_Vm *VmSession) AssumeNoRevert1(potentialRevert VmSafePotentialRevert) error {
	return _Vm.Contract.AssumeNoRevert1(&_Vm.CallOpts, potentialRevert)
}

// AssumeNoRevert1 is a free data retrieval call binding the contract method 0xd8591eeb.
//
// Solidity: function assumeNoRevert((address,bool,bytes) potentialRevert) pure returns()
func (_Vm *VmCallerSession) AssumeNoRevert1(potentialRevert VmSafePotentialRevert) error {
	return _Vm.Contract.AssumeNoRevert1(&_Vm.CallOpts, potentialRevert)
}

// Breakpoint is a free data retrieval call binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) pure returns()
func (_Vm *VmCaller) Breakpoint(opts *bind.CallOpts, char string) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "breakpoint", char)

	if err != nil {
		return err
	}

	return err

}

// Breakpoint is a free data retrieval call binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) pure returns()
func (_Vm *VmSession) Breakpoint(char string) error {
	return _Vm.Contract.Breakpoint(&_Vm.CallOpts, char)
}

// Breakpoint is a free data retrieval call binding the contract method 0xf0259e92.
//
// Solidity: function breakpoint(string char) pure returns()
func (_Vm *VmCallerSession) Breakpoint(char string) error {
	return _Vm.Contract.Breakpoint(&_Vm.CallOpts, char)
}

// Breakpoint0 is a free data retrieval call binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) pure returns()
func (_Vm *VmCaller) Breakpoint0(opts *bind.CallOpts, char string, value bool) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "breakpoint0", char, value)

	if err != nil {
		return err
	}

	return err

}

// Breakpoint0 is a free data retrieval call binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) pure returns()
func (_Vm *VmSession) Breakpoint0(char string, value bool) error {
	return _Vm.Contract.Breakpoint0(&_Vm.CallOpts, char, value)
}

// Breakpoint0 is a free data retrieval call binding the contract method 0xf7d39a8d.
//
// Solidity: function breakpoint(string char, bool value) pure returns()
func (_Vm *VmCallerSession) Breakpoint0(char string, value bool) error {
	return _Vm.Contract.Breakpoint0(&_Vm.CallOpts, char, value)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmCaller) ComputeCreate2Address(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreate2Address", salt, initCodeHash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address(&_Vm.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address is a free data retrieval call binding the contract method 0x890c283b.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreate2Address(salt [32]byte, initCodeHash [32]byte) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address(&_Vm.CallOpts, salt, initCodeHash)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmCaller) ComputeCreate2Address0(opts *bind.CallOpts, salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreate2Address0", salt, initCodeHash, deployer)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address0(&_Vm.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreate2Address0 is a free data retrieval call binding the contract method 0xd323826a.
//
// Solidity: function computeCreate2Address(bytes32 salt, bytes32 initCodeHash, address deployer) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreate2Address0(salt [32]byte, initCodeHash [32]byte, deployer common.Address) (common.Address, error) {
	return _Vm.Contract.ComputeCreate2Address0(&_Vm.CallOpts, salt, initCodeHash, deployer)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmCaller) ComputeCreateAddress(opts *bind.CallOpts, deployer common.Address, nonce *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "computeCreateAddress", deployer, nonce)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _Vm.Contract.ComputeCreateAddress(&_Vm.CallOpts, deployer, nonce)
}

// ComputeCreateAddress is a free data retrieval call binding the contract method 0x74637a7a.
//
// Solidity: function computeCreateAddress(address deployer, uint256 nonce) pure returns(address)
func (_Vm *VmCallerSession) ComputeCreateAddress(deployer common.Address, nonce *big.Int) (common.Address, error) {
	return _Vm.Contract.ComputeCreateAddress(&_Vm.CallOpts, deployer, nonce)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey", mnemonic, derivationPath, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey(&_Vm.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey is a free data retrieval call binding the contract method 0x29233b1f.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey(mnemonic string, derivationPath string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey(&_Vm.CallOpts, mnemonic, derivationPath, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey0(opts *bind.CallOpts, mnemonic string, index uint32, language string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey0", mnemonic, index, language)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey0(&_Vm.CallOpts, mnemonic, index, language)
}

// DeriveKey0 is a free data retrieval call binding the contract method 0x32c8176d.
//
// Solidity: function deriveKey(string mnemonic, uint32 index, string language) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey0(mnemonic string, index uint32, language string) (*big.Int, error) {
	return _Vm.Contract.DeriveKey0(&_Vm.CallOpts, mnemonic, index, language)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey1(opts *bind.CallOpts, mnemonic string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey1", mnemonic, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey1(&_Vm.CallOpts, mnemonic, index)
}

// DeriveKey1 is a free data retrieval call binding the contract method 0x6229498b.
//
// Solidity: function deriveKey(string mnemonic, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey1(mnemonic string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey1(&_Vm.CallOpts, mnemonic, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCaller) DeriveKey2(opts *bind.CallOpts, mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "deriveKey2", mnemonic, derivationPath, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey2(&_Vm.CallOpts, mnemonic, derivationPath, index)
}

// DeriveKey2 is a free data retrieval call binding the contract method 0x6bcb2c1b.
//
// Solidity: function deriveKey(string mnemonic, string derivationPath, uint32 index) pure returns(uint256 privateKey)
func (_Vm *VmCallerSession) DeriveKey2(mnemonic string, derivationPath string, index uint32) (*big.Int, error) {
	return _Vm.Contract.DeriveKey2(&_Vm.CallOpts, mnemonic, derivationPath, index)
}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_Vm *VmCaller) EnsNamehash(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "ensNamehash", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_Vm *VmSession) EnsNamehash(name string) ([32]byte, error) {
	return _Vm.Contract.EnsNamehash(&_Vm.CallOpts, name)
}

// EnsNamehash is a free data retrieval call binding the contract method 0x8c374c65.
//
// Solidity: function ensNamehash(string name) pure returns(bytes32)
func (_Vm *VmCallerSession) EnsNamehash(name string) ([32]byte, error) {
	return _Vm.Contract.EnsNamehash(&_Vm.CallOpts, name)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmCaller) EnvAddress(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envAddress", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmSession) EnvAddress(name string) (common.Address, error) {
	return _Vm.Contract.EnvAddress(&_Vm.CallOpts, name)
}

// EnvAddress is a free data retrieval call binding the contract method 0x350d56bf.
//
// Solidity: function envAddress(string name) view returns(address value)
func (_Vm *VmCallerSession) EnvAddress(name string) (common.Address, error) {
	return _Vm.Contract.EnvAddress(&_Vm.CallOpts, name)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmCaller) EnvAddress0(opts *bind.CallOpts, name string, delim string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envAddress0", name, delim)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _Vm.Contract.EnvAddress0(&_Vm.CallOpts, name, delim)
}

// EnvAddress0 is a free data retrieval call binding the contract method 0xad31b9fa.
//
// Solidity: function envAddress(string name, string delim) view returns(address[] value)
func (_Vm *VmCallerSession) EnvAddress0(name string, delim string) ([]common.Address, error) {
	return _Vm.Contract.EnvAddress0(&_Vm.CallOpts, name, delim)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmCaller) EnvBool(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBool", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmSession) EnvBool(name string) (bool, error) {
	return _Vm.Contract.EnvBool(&_Vm.CallOpts, name)
}

// EnvBool is a free data retrieval call binding the contract method 0x7ed1ec7d.
//
// Solidity: function envBool(string name) view returns(bool value)
func (_Vm *VmCallerSession) EnvBool(name string) (bool, error) {
	return _Vm.Contract.EnvBool(&_Vm.CallOpts, name)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmCaller) EnvBool0(opts *bind.CallOpts, name string, delim string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBool0", name, delim)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _Vm.Contract.EnvBool0(&_Vm.CallOpts, name, delim)
}

// EnvBool0 is a free data retrieval call binding the contract method 0xaaaddeaf.
//
// Solidity: function envBool(string name, string delim) view returns(bool[] value)
func (_Vm *VmCallerSession) EnvBool0(name string, delim string) ([]bool, error) {
	return _Vm.Contract.EnvBool0(&_Vm.CallOpts, name, delim)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmCaller) EnvBytes(opts *bind.CallOpts, name string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes", name)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmSession) EnvBytes(name string) ([]byte, error) {
	return _Vm.Contract.EnvBytes(&_Vm.CallOpts, name)
}

// EnvBytes is a free data retrieval call binding the contract method 0x4d7baf06.
//
// Solidity: function envBytes(string name) view returns(bytes value)
func (_Vm *VmCallerSession) EnvBytes(name string) ([]byte, error) {
	return _Vm.Contract.EnvBytes(&_Vm.CallOpts, name)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmCaller) EnvBytes0(opts *bind.CallOpts, name string, delim string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes0", name, delim)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _Vm.Contract.EnvBytes0(&_Vm.CallOpts, name, delim)
}

// EnvBytes0 is a free data retrieval call binding the contract method 0xddc2651b.
//
// Solidity: function envBytes(string name, string delim) view returns(bytes[] value)
func (_Vm *VmCallerSession) EnvBytes0(name string, delim string) ([][]byte, error) {
	return _Vm.Contract.EnvBytes0(&_Vm.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmCaller) EnvBytes32(opts *bind.CallOpts, name string, delim string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes32", name, delim)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _Vm.Contract.EnvBytes32(&_Vm.CallOpts, name, delim)
}

// EnvBytes32 is a free data retrieval call binding the contract method 0x5af231c1.
//
// Solidity: function envBytes32(string name, string delim) view returns(bytes32[] value)
func (_Vm *VmCallerSession) EnvBytes32(name string, delim string) ([][32]byte, error) {
	return _Vm.Contract.EnvBytes32(&_Vm.CallOpts, name, delim)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmCaller) EnvBytes320(opts *bind.CallOpts, name string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envBytes320", name)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmSession) EnvBytes320(name string) ([32]byte, error) {
	return _Vm.Contract.EnvBytes320(&_Vm.CallOpts, name)
}

// EnvBytes320 is a free data retrieval call binding the contract method 0x97949042.
//
// Solidity: function envBytes32(string name) view returns(bytes32 value)
func (_Vm *VmCallerSession) EnvBytes320(name string) ([32]byte, error) {
	return _Vm.Contract.EnvBytes320(&_Vm.CallOpts, name)
}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_Vm *VmCaller) EnvExists(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envExists", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_Vm *VmSession) EnvExists(name string) (bool, error) {
	return _Vm.Contract.EnvExists(&_Vm.CallOpts, name)
}

// EnvExists is a free data retrieval call binding the contract method 0xce8365f9.
//
// Solidity: function envExists(string name) view returns(bool result)
func (_Vm *VmCallerSession) EnvExists(name string) (bool, error) {
	return _Vm.Contract.EnvExists(&_Vm.CallOpts, name)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmCaller) EnvInt(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envInt", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvInt(&_Vm.CallOpts, name, delim)
}

// EnvInt is a free data retrieval call binding the contract method 0x42181150.
//
// Solidity: function envInt(string name, string delim) view returns(int256[] value)
func (_Vm *VmCallerSession) EnvInt(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvInt(&_Vm.CallOpts, name, delim)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmCaller) EnvInt0(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envInt0", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmSession) EnvInt0(name string) (*big.Int, error) {
	return _Vm.Contract.EnvInt0(&_Vm.CallOpts, name)
}

// EnvInt0 is a free data retrieval call binding the contract method 0x892a0c61.
//
// Solidity: function envInt(string name) view returns(int256 value)
func (_Vm *VmCallerSession) EnvInt0(name string) (*big.Int, error) {
	return _Vm.Contract.EnvInt0(&_Vm.CallOpts, name)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmCaller) EnvOr(opts *bind.CallOpts, name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr", name, delim, defaultValue)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _Vm.Contract.EnvOr(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr is a free data retrieval call binding the contract method 0x2281f367.
//
// Solidity: function envOr(string name, string delim, bytes32[] defaultValue) view returns(bytes32[] value)
func (_Vm *VmCallerSession) EnvOr(name string, delim string, defaultValue [][32]byte) ([][32]byte, error) {
	return _Vm.Contract.EnvOr(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmCaller) EnvOr0(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr0", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr0(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr0 is a free data retrieval call binding the contract method 0x4700d74b.
//
// Solidity: function envOr(string name, string delim, int256[] defaultValue) view returns(int256[] value)
func (_Vm *VmCallerSession) EnvOr0(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr0(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmCaller) EnvOr1(opts *bind.CallOpts, name string, defaultValue bool) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr1", name, defaultValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _Vm.Contract.EnvOr1(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr1 is a free data retrieval call binding the contract method 0x4777f3cf.
//
// Solidity: function envOr(string name, bool defaultValue) view returns(bool value)
func (_Vm *VmCallerSession) EnvOr1(name string, defaultValue bool) (bool, error) {
	return _Vm.Contract.EnvOr1(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmCaller) EnvOr10(opts *bind.CallOpts, name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr10", name, delim, defaultValue)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _Vm.Contract.EnvOr10(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr10 is a free data retrieval call binding the contract method 0xc74e9deb.
//
// Solidity: function envOr(string name, string delim, address[] defaultValue) view returns(address[] value)
func (_Vm *VmCallerSession) EnvOr10(name string, delim string, defaultValue []common.Address) ([]common.Address, error) {
	return _Vm.Contract.EnvOr10(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmCaller) EnvOr11(opts *bind.CallOpts, name string, defaultValue string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr11", name, defaultValue)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _Vm.Contract.EnvOr11(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr11 is a free data retrieval call binding the contract method 0xd145736c.
//
// Solidity: function envOr(string name, string defaultValue) view returns(string value)
func (_Vm *VmCallerSession) EnvOr11(name string, defaultValue string) (string, error) {
	return _Vm.Contract.EnvOr11(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmCaller) EnvOr12(opts *bind.CallOpts, name string, delim string, defaultValue []bool) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr12", name, delim, defaultValue)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _Vm.Contract.EnvOr12(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr12 is a free data retrieval call binding the contract method 0xeb85e83b.
//
// Solidity: function envOr(string name, string delim, bool[] defaultValue) view returns(bool[] value)
func (_Vm *VmCallerSession) EnvOr12(name string, delim string, defaultValue []bool) ([]bool, error) {
	return _Vm.Contract.EnvOr12(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmCaller) EnvOr2(opts *bind.CallOpts, name string, defaultValue common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr2", name, defaultValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _Vm.Contract.EnvOr2(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr2 is a free data retrieval call binding the contract method 0x561fe540.
//
// Solidity: function envOr(string name, address defaultValue) view returns(address value)
func (_Vm *VmCallerSession) EnvOr2(name string, defaultValue common.Address) (common.Address, error) {
	return _Vm.Contract.EnvOr2(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmCaller) EnvOr3(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr3", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr3(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr3 is a free data retrieval call binding the contract method 0x5e97348f.
//
// Solidity: function envOr(string name, uint256 defaultValue) view returns(uint256 value)
func (_Vm *VmCallerSession) EnvOr3(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr3(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmCaller) EnvOr4(opts *bind.CallOpts, name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr4", name, delim, defaultValue)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _Vm.Contract.EnvOr4(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr4 is a free data retrieval call binding the contract method 0x64bc3e64.
//
// Solidity: function envOr(string name, string delim, bytes[] defaultValue) view returns(bytes[] value)
func (_Vm *VmCallerSession) EnvOr4(name string, delim string, defaultValue [][]byte) ([][]byte, error) {
	return _Vm.Contract.EnvOr4(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmCaller) EnvOr5(opts *bind.CallOpts, name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr5", name, delim, defaultValue)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr5(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr5 is a free data retrieval call binding the contract method 0x74318528.
//
// Solidity: function envOr(string name, string delim, uint256[] defaultValue) view returns(uint256[] value)
func (_Vm *VmCallerSession) EnvOr5(name string, delim string, defaultValue []*big.Int) ([]*big.Int, error) {
	return _Vm.Contract.EnvOr5(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmCaller) EnvOr6(opts *bind.CallOpts, name string, delim string, defaultValue []string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr6", name, delim, defaultValue)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _Vm.Contract.EnvOr6(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr6 is a free data retrieval call binding the contract method 0x859216bc.
//
// Solidity: function envOr(string name, string delim, string[] defaultValue) view returns(string[] value)
func (_Vm *VmCallerSession) EnvOr6(name string, delim string, defaultValue []string) ([]string, error) {
	return _Vm.Contract.EnvOr6(&_Vm.CallOpts, name, delim, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmCaller) EnvOr7(opts *bind.CallOpts, name string, defaultValue []byte) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr7", name, defaultValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _Vm.Contract.EnvOr7(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr7 is a free data retrieval call binding the contract method 0xb3e47705.
//
// Solidity: function envOr(string name, bytes defaultValue) view returns(bytes value)
func (_Vm *VmCallerSession) EnvOr7(name string, defaultValue []byte) ([]byte, error) {
	return _Vm.Contract.EnvOr7(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmCaller) EnvOr8(opts *bind.CallOpts, name string, defaultValue [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr8", name, defaultValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _Vm.Contract.EnvOr8(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr8 is a free data retrieval call binding the contract method 0xb4a85892.
//
// Solidity: function envOr(string name, bytes32 defaultValue) view returns(bytes32 value)
func (_Vm *VmCallerSession) EnvOr8(name string, defaultValue [32]byte) ([32]byte, error) {
	return _Vm.Contract.EnvOr8(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmCaller) EnvOr9(opts *bind.CallOpts, name string, defaultValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envOr9", name, defaultValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr9(&_Vm.CallOpts, name, defaultValue)
}

// EnvOr9 is a free data retrieval call binding the contract method 0xbbcb713e.
//
// Solidity: function envOr(string name, int256 defaultValue) view returns(int256 value)
func (_Vm *VmCallerSession) EnvOr9(name string, defaultValue *big.Int) (*big.Int, error) {
	return _Vm.Contract.EnvOr9(&_Vm.CallOpts, name, defaultValue)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmCaller) EnvString(opts *bind.CallOpts, name string, delim string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envString", name, delim)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmSession) EnvString(name string, delim string) ([]string, error) {
	return _Vm.Contract.EnvString(&_Vm.CallOpts, name, delim)
}

// EnvString is a free data retrieval call binding the contract method 0x14b02bc9.
//
// Solidity: function envString(string name, string delim) view returns(string[] value)
func (_Vm *VmCallerSession) EnvString(name string, delim string) ([]string, error) {
	return _Vm.Contract.EnvString(&_Vm.CallOpts, name, delim)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmCaller) EnvString0(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envString0", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmSession) EnvString0(name string) (string, error) {
	return _Vm.Contract.EnvString0(&_Vm.CallOpts, name)
}

// EnvString0 is a free data retrieval call binding the contract method 0xf877cb19.
//
// Solidity: function envString(string name) view returns(string value)
func (_Vm *VmCallerSession) EnvString0(name string) (string, error) {
	return _Vm.Contract.EnvString0(&_Vm.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmCaller) EnvUint(opts *bind.CallOpts, name string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envUint", name)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmSession) EnvUint(name string) (*big.Int, error) {
	return _Vm.Contract.EnvUint(&_Vm.CallOpts, name)
}

// EnvUint is a free data retrieval call binding the contract method 0xc1978d1f.
//
// Solidity: function envUint(string name) view returns(uint256 value)
func (_Vm *VmCallerSession) EnvUint(name string) (*big.Int, error) {
	return _Vm.Contract.EnvUint(&_Vm.CallOpts, name)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmCaller) EnvUint0(opts *bind.CallOpts, name string, delim string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "envUint0", name, delim)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvUint0(&_Vm.CallOpts, name, delim)
}

// EnvUint0 is a free data retrieval call binding the contract method 0xf3dec099.
//
// Solidity: function envUint(string name, string delim) view returns(uint256[] value)
func (_Vm *VmCallerSession) EnvUint0(name string, delim string) ([]*big.Int, error) {
	return _Vm.Contract.EnvUint0(&_Vm.CallOpts, name, delim)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) view returns(bool result)
func (_Vm *VmCaller) Exists(opts *bind.CallOpts, path string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "exists", path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) view returns(bool result)
func (_Vm *VmSession) Exists(path string) (bool, error) {
	return _Vm.Contract.Exists(&_Vm.CallOpts, path)
}

// Exists is a free data retrieval call binding the contract method 0x261a323e.
//
// Solidity: function exists(string path) view returns(bool result)
func (_Vm *VmCallerSession) Exists(path string) (bool, error) {
	return _Vm.Contract.Exists(&_Vm.CallOpts, path)
}

// FoundryVersionAtLeast is a free data retrieval call binding the contract method 0x6248be1f.
//
// Solidity: function foundryVersionAtLeast(string version) view returns(bool)
func (_Vm *VmCaller) FoundryVersionAtLeast(opts *bind.CallOpts, version string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "foundryVersionAtLeast", version)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FoundryVersionAtLeast is a free data retrieval call binding the contract method 0x6248be1f.
//
// Solidity: function foundryVersionAtLeast(string version) view returns(bool)
func (_Vm *VmSession) FoundryVersionAtLeast(version string) (bool, error) {
	return _Vm.Contract.FoundryVersionAtLeast(&_Vm.CallOpts, version)
}

// FoundryVersionAtLeast is a free data retrieval call binding the contract method 0x6248be1f.
//
// Solidity: function foundryVersionAtLeast(string version) view returns(bool)
func (_Vm *VmCallerSession) FoundryVersionAtLeast(version string) (bool, error) {
	return _Vm.Contract.FoundryVersionAtLeast(&_Vm.CallOpts, version)
}

// FoundryVersionCmp is a free data retrieval call binding the contract method 0xca7b0a09.
//
// Solidity: function foundryVersionCmp(string version) view returns(int256)
func (_Vm *VmCaller) FoundryVersionCmp(opts *bind.CallOpts, version string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "foundryVersionCmp", version)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FoundryVersionCmp is a free data retrieval call binding the contract method 0xca7b0a09.
//
// Solidity: function foundryVersionCmp(string version) view returns(int256)
func (_Vm *VmSession) FoundryVersionCmp(version string) (*big.Int, error) {
	return _Vm.Contract.FoundryVersionCmp(&_Vm.CallOpts, version)
}

// FoundryVersionCmp is a free data retrieval call binding the contract method 0xca7b0a09.
//
// Solidity: function foundryVersionCmp(string version) view returns(int256)
func (_Vm *VmCallerSession) FoundryVersionCmp(version string) (*big.Int, error) {
	return _Vm.Contract.FoundryVersionCmp(&_Vm.CallOpts, version)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmCaller) FsMetadata(opts *bind.CallOpts, path string) (VmSafeFsMetadata, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "fsMetadata", path)

	if err != nil {
		return *new(VmSafeFsMetadata), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeFsMetadata)).(*VmSafeFsMetadata)

	return out0, err

}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _Vm.Contract.FsMetadata(&_Vm.CallOpts, path)
}

// FsMetadata is a free data retrieval call binding the contract method 0xaf368a08.
//
// Solidity: function fsMetadata(string path) view returns((bool,bool,uint256,bool,uint256,uint256,uint256) metadata)
func (_Vm *VmCallerSession) FsMetadata(path string) (VmSafeFsMetadata, error) {
	return _Vm.Contract.FsMetadata(&_Vm.CallOpts, path)
}

// GetArtifactPathByCode is a free data retrieval call binding the contract method 0xeb74848c.
//
// Solidity: function getArtifactPathByCode(bytes code) view returns(string path)
func (_Vm *VmCaller) GetArtifactPathByCode(opts *bind.CallOpts, code []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getArtifactPathByCode", code)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetArtifactPathByCode is a free data retrieval call binding the contract method 0xeb74848c.
//
// Solidity: function getArtifactPathByCode(bytes code) view returns(string path)
func (_Vm *VmSession) GetArtifactPathByCode(code []byte) (string, error) {
	return _Vm.Contract.GetArtifactPathByCode(&_Vm.CallOpts, code)
}

// GetArtifactPathByCode is a free data retrieval call binding the contract method 0xeb74848c.
//
// Solidity: function getArtifactPathByCode(bytes code) view returns(string path)
func (_Vm *VmCallerSession) GetArtifactPathByCode(code []byte) (string, error) {
	return _Vm.Contract.GetArtifactPathByCode(&_Vm.CallOpts, code)
}

// GetArtifactPathByDeployedCode is a free data retrieval call binding the contract method 0x6d853ba5.
//
// Solidity: function getArtifactPathByDeployedCode(bytes deployedCode) view returns(string path)
func (_Vm *VmCaller) GetArtifactPathByDeployedCode(opts *bind.CallOpts, deployedCode []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getArtifactPathByDeployedCode", deployedCode)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetArtifactPathByDeployedCode is a free data retrieval call binding the contract method 0x6d853ba5.
//
// Solidity: function getArtifactPathByDeployedCode(bytes deployedCode) view returns(string path)
func (_Vm *VmSession) GetArtifactPathByDeployedCode(deployedCode []byte) (string, error) {
	return _Vm.Contract.GetArtifactPathByDeployedCode(&_Vm.CallOpts, deployedCode)
}

// GetArtifactPathByDeployedCode is a free data retrieval call binding the contract method 0x6d853ba5.
//
// Solidity: function getArtifactPathByDeployedCode(bytes deployedCode) view returns(string path)
func (_Vm *VmCallerSession) GetArtifactPathByDeployedCode(deployedCode []byte) (string, error) {
	return _Vm.Contract.GetArtifactPathByDeployedCode(&_Vm.CallOpts, deployedCode)
}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_Vm *VmCaller) GetBlobBaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlobBaseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_Vm *VmSession) GetBlobBaseFee() (*big.Int, error) {
	return _Vm.Contract.GetBlobBaseFee(&_Vm.CallOpts)
}

// GetBlobBaseFee is a free data retrieval call binding the contract method 0x1f6d6ef7.
//
// Solidity: function getBlobBaseFee() view returns(uint256 blobBaseFee)
func (_Vm *VmCallerSession) GetBlobBaseFee() (*big.Int, error) {
	return _Vm.Contract.GetBlobBaseFee(&_Vm.CallOpts)
}

// GetBlobhashes is a free data retrieval call binding the contract method 0xf56ff18b.
//
// Solidity: function getBlobhashes() view returns(bytes32[] hashes)
func (_Vm *VmCaller) GetBlobhashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlobhashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBlobhashes is a free data retrieval call binding the contract method 0xf56ff18b.
//
// Solidity: function getBlobhashes() view returns(bytes32[] hashes)
func (_Vm *VmSession) GetBlobhashes() ([][32]byte, error) {
	return _Vm.Contract.GetBlobhashes(&_Vm.CallOpts)
}

// GetBlobhashes is a free data retrieval call binding the contract method 0xf56ff18b.
//
// Solidity: function getBlobhashes() view returns(bytes32[] hashes)
func (_Vm *VmCallerSession) GetBlobhashes() ([][32]byte, error) {
	return _Vm.Contract.GetBlobhashes(&_Vm.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmSession) GetBlockNumber() (*big.Int, error) {
	return _Vm.Contract.GetBlockNumber(&_Vm.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256 height)
func (_Vm *VmCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Vm.Contract.GetBlockNumber(&_Vm.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmCaller) GetBlockTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBlockTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmSession) GetBlockTimestamp() (*big.Int, error) {
	return _Vm.Contract.GetBlockTimestamp(&_Vm.CallOpts)
}

// GetBlockTimestamp is a free data retrieval call binding the contract method 0x796b89b9.
//
// Solidity: function getBlockTimestamp() view returns(uint256 timestamp)
func (_Vm *VmCallerSession) GetBlockTimestamp() (*big.Int, error) {
	return _Vm.Contract.GetBlockTimestamp(&_Vm.CallOpts)
}

// GetBroadcast is a free data retrieval call binding the contract method 0x3dc90cb3.
//
// Solidity: function getBroadcast(string contractName, uint64 chainId, uint8 txType) view returns((bytes32,uint8,address,uint64,bool))
func (_Vm *VmCaller) GetBroadcast(opts *bind.CallOpts, contractName string, chainId uint64, txType uint8) (VmSafeBroadcastTxSummary, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBroadcast", contractName, chainId, txType)

	if err != nil {
		return *new(VmSafeBroadcastTxSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeBroadcastTxSummary)).(*VmSafeBroadcastTxSummary)

	return out0, err

}

// GetBroadcast is a free data retrieval call binding the contract method 0x3dc90cb3.
//
// Solidity: function getBroadcast(string contractName, uint64 chainId, uint8 txType) view returns((bytes32,uint8,address,uint64,bool))
func (_Vm *VmSession) GetBroadcast(contractName string, chainId uint64, txType uint8) (VmSafeBroadcastTxSummary, error) {
	return _Vm.Contract.GetBroadcast(&_Vm.CallOpts, contractName, chainId, txType)
}

// GetBroadcast is a free data retrieval call binding the contract method 0x3dc90cb3.
//
// Solidity: function getBroadcast(string contractName, uint64 chainId, uint8 txType) view returns((bytes32,uint8,address,uint64,bool))
func (_Vm *VmCallerSession) GetBroadcast(contractName string, chainId uint64, txType uint8) (VmSafeBroadcastTxSummary, error) {
	return _Vm.Contract.GetBroadcast(&_Vm.CallOpts, contractName, chainId, txType)
}

// GetBroadcasts is a free data retrieval call binding the contract method 0xf2fa4a26.
//
// Solidity: function getBroadcasts(string contractName, uint64 chainId) view returns((bytes32,uint8,address,uint64,bool)[])
func (_Vm *VmCaller) GetBroadcasts(opts *bind.CallOpts, contractName string, chainId uint64) ([]VmSafeBroadcastTxSummary, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBroadcasts", contractName, chainId)

	if err != nil {
		return *new([]VmSafeBroadcastTxSummary), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeBroadcastTxSummary)).(*[]VmSafeBroadcastTxSummary)

	return out0, err

}

// GetBroadcasts is a free data retrieval call binding the contract method 0xf2fa4a26.
//
// Solidity: function getBroadcasts(string contractName, uint64 chainId) view returns((bytes32,uint8,address,uint64,bool)[])
func (_Vm *VmSession) GetBroadcasts(contractName string, chainId uint64) ([]VmSafeBroadcastTxSummary, error) {
	return _Vm.Contract.GetBroadcasts(&_Vm.CallOpts, contractName, chainId)
}

// GetBroadcasts is a free data retrieval call binding the contract method 0xf2fa4a26.
//
// Solidity: function getBroadcasts(string contractName, uint64 chainId) view returns((bytes32,uint8,address,uint64,bool)[])
func (_Vm *VmCallerSession) GetBroadcasts(contractName string, chainId uint64) ([]VmSafeBroadcastTxSummary, error) {
	return _Vm.Contract.GetBroadcasts(&_Vm.CallOpts, contractName, chainId)
}

// GetBroadcasts0 is a free data retrieval call binding the contract method 0xf7afe919.
//
// Solidity: function getBroadcasts(string contractName, uint64 chainId, uint8 txType) view returns((bytes32,uint8,address,uint64,bool)[])
func (_Vm *VmCaller) GetBroadcasts0(opts *bind.CallOpts, contractName string, chainId uint64, txType uint8) ([]VmSafeBroadcastTxSummary, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getBroadcasts0", contractName, chainId, txType)

	if err != nil {
		return *new([]VmSafeBroadcastTxSummary), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeBroadcastTxSummary)).(*[]VmSafeBroadcastTxSummary)

	return out0, err

}

// GetBroadcasts0 is a free data retrieval call binding the contract method 0xf7afe919.
//
// Solidity: function getBroadcasts(string contractName, uint64 chainId, uint8 txType) view returns((bytes32,uint8,address,uint64,bool)[])
func (_Vm *VmSession) GetBroadcasts0(contractName string, chainId uint64, txType uint8) ([]VmSafeBroadcastTxSummary, error) {
	return _Vm.Contract.GetBroadcasts0(&_Vm.CallOpts, contractName, chainId, txType)
}

// GetBroadcasts0 is a free data retrieval call binding the contract method 0xf7afe919.
//
// Solidity: function getBroadcasts(string contractName, uint64 chainId, uint8 txType) view returns((bytes32,uint8,address,uint64,bool)[])
func (_Vm *VmCallerSession) GetBroadcasts0(contractName string, chainId uint64, txType uint8) ([]VmSafeBroadcastTxSummary, error) {
	return _Vm.Contract.GetBroadcasts0(&_Vm.CallOpts, contractName, chainId, txType)
}

// GetChain is a free data retrieval call binding the contract method 0x4cc1c2bb.
//
// Solidity: function getChain(string chainAlias) view returns((string,uint256,string,string) chain)
func (_Vm *VmCaller) GetChain(opts *bind.CallOpts, chainAlias string) (VmSafeChain, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getChain", chainAlias)

	if err != nil {
		return *new(VmSafeChain), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeChain)).(*VmSafeChain)

	return out0, err

}

// GetChain is a free data retrieval call binding the contract method 0x4cc1c2bb.
//
// Solidity: function getChain(string chainAlias) view returns((string,uint256,string,string) chain)
func (_Vm *VmSession) GetChain(chainAlias string) (VmSafeChain, error) {
	return _Vm.Contract.GetChain(&_Vm.CallOpts, chainAlias)
}

// GetChain is a free data retrieval call binding the contract method 0x4cc1c2bb.
//
// Solidity: function getChain(string chainAlias) view returns((string,uint256,string,string) chain)
func (_Vm *VmCallerSession) GetChain(chainAlias string) (VmSafeChain, error) {
	return _Vm.Contract.GetChain(&_Vm.CallOpts, chainAlias)
}

// GetChain0 is a free data retrieval call binding the contract method 0xb6791ad4.
//
// Solidity: function getChain(uint256 chainId) view returns((string,uint256,string,string) chain)
func (_Vm *VmCaller) GetChain0(opts *bind.CallOpts, chainId *big.Int) (VmSafeChain, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getChain0", chainId)

	if err != nil {
		return *new(VmSafeChain), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeChain)).(*VmSafeChain)

	return out0, err

}

// GetChain0 is a free data retrieval call binding the contract method 0xb6791ad4.
//
// Solidity: function getChain(uint256 chainId) view returns((string,uint256,string,string) chain)
func (_Vm *VmSession) GetChain0(chainId *big.Int) (VmSafeChain, error) {
	return _Vm.Contract.GetChain0(&_Vm.CallOpts, chainId)
}

// GetChain0 is a free data retrieval call binding the contract method 0xb6791ad4.
//
// Solidity: function getChain(uint256 chainId) view returns((string,uint256,string,string) chain)
func (_Vm *VmCallerSession) GetChain0(chainId *big.Int) (VmSafeChain, error) {
	return _Vm.Contract.GetChain0(&_Vm.CallOpts, chainId)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmCaller) GetCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmSession) GetCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetCode(&_Vm.CallOpts, artifactPath)
}

// GetCode is a free data retrieval call binding the contract method 0x8d1cc925.
//
// Solidity: function getCode(string artifactPath) view returns(bytes creationBytecode)
func (_Vm *VmCallerSession) GetCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetCode(&_Vm.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmCaller) GetDeployedCode(opts *bind.CallOpts, artifactPath string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getDeployedCode", artifactPath)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetDeployedCode(&_Vm.CallOpts, artifactPath)
}

// GetDeployedCode is a free data retrieval call binding the contract method 0x3ebf73b4.
//
// Solidity: function getDeployedCode(string artifactPath) view returns(bytes runtimeBytecode)
func (_Vm *VmCallerSession) GetDeployedCode(artifactPath string) ([]byte, error) {
	return _Vm.Contract.GetDeployedCode(&_Vm.CallOpts, artifactPath)
}

// GetDeployment is a free data retrieval call binding the contract method 0x0debd5d6.
//
// Solidity: function getDeployment(string contractName, uint64 chainId) view returns(address deployedAddress)
func (_Vm *VmCaller) GetDeployment(opts *bind.CallOpts, contractName string, chainId uint64) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getDeployment", contractName, chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDeployment is a free data retrieval call binding the contract method 0x0debd5d6.
//
// Solidity: function getDeployment(string contractName, uint64 chainId) view returns(address deployedAddress)
func (_Vm *VmSession) GetDeployment(contractName string, chainId uint64) (common.Address, error) {
	return _Vm.Contract.GetDeployment(&_Vm.CallOpts, contractName, chainId)
}

// GetDeployment is a free data retrieval call binding the contract method 0x0debd5d6.
//
// Solidity: function getDeployment(string contractName, uint64 chainId) view returns(address deployedAddress)
func (_Vm *VmCallerSession) GetDeployment(contractName string, chainId uint64) (common.Address, error) {
	return _Vm.Contract.GetDeployment(&_Vm.CallOpts, contractName, chainId)
}

// GetDeployment0 is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string contractName) view returns(address deployedAddress)
func (_Vm *VmCaller) GetDeployment0(opts *bind.CallOpts, contractName string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getDeployment0", contractName)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDeployment0 is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string contractName) view returns(address deployedAddress)
func (_Vm *VmSession) GetDeployment0(contractName string) (common.Address, error) {
	return _Vm.Contract.GetDeployment0(&_Vm.CallOpts, contractName)
}

// GetDeployment0 is a free data retrieval call binding the contract method 0xa8091d97.
//
// Solidity: function getDeployment(string contractName) view returns(address deployedAddress)
func (_Vm *VmCallerSession) GetDeployment0(contractName string) (common.Address, error) {
	return _Vm.Contract.GetDeployment0(&_Vm.CallOpts, contractName)
}

// GetDeployments is a free data retrieval call binding the contract method 0x74e133dd.
//
// Solidity: function getDeployments(string contractName, uint64 chainId) view returns(address[] deployedAddresses)
func (_Vm *VmCaller) GetDeployments(opts *bind.CallOpts, contractName string, chainId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getDeployments", contractName, chainId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeployments is a free data retrieval call binding the contract method 0x74e133dd.
//
// Solidity: function getDeployments(string contractName, uint64 chainId) view returns(address[] deployedAddresses)
func (_Vm *VmSession) GetDeployments(contractName string, chainId uint64) ([]common.Address, error) {
	return _Vm.Contract.GetDeployments(&_Vm.CallOpts, contractName, chainId)
}

// GetDeployments is a free data retrieval call binding the contract method 0x74e133dd.
//
// Solidity: function getDeployments(string contractName, uint64 chainId) view returns(address[] deployedAddresses)
func (_Vm *VmCallerSession) GetDeployments(contractName string, chainId uint64) ([]common.Address, error) {
	return _Vm.Contract.GetDeployments(&_Vm.CallOpts, contractName, chainId)
}

// GetFoundryVersion is a free data retrieval call binding the contract method 0xea991bb5.
//
// Solidity: function getFoundryVersion() view returns(string version)
func (_Vm *VmCaller) GetFoundryVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getFoundryVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetFoundryVersion is a free data retrieval call binding the contract method 0xea991bb5.
//
// Solidity: function getFoundryVersion() view returns(string version)
func (_Vm *VmSession) GetFoundryVersion() (string, error) {
	return _Vm.Contract.GetFoundryVersion(&_Vm.CallOpts)
}

// GetFoundryVersion is a free data retrieval call binding the contract method 0xea991bb5.
//
// Solidity: function getFoundryVersion() view returns(string version)
func (_Vm *VmCallerSession) GetFoundryVersion() (string, error) {
	return _Vm.Contract.GetFoundryVersion(&_Vm.CallOpts)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmCaller) GetLabel(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getLabel", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmSession) GetLabel(account common.Address) (string, error) {
	return _Vm.Contract.GetLabel(&_Vm.CallOpts, account)
}

// GetLabel is a free data retrieval call binding the contract method 0x28a249b0.
//
// Solidity: function getLabel(address account) view returns(string currentLabel)
func (_Vm *VmCallerSession) GetLabel(account common.Address) (string, error) {
	return _Vm.Contract.GetLabel(&_Vm.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmCaller) GetNonce(opts *bind.CallOpts, account common.Address) (uint64, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getNonce", account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmSession) GetNonce(account common.Address) (uint64, error) {
	return _Vm.Contract.GetNonce(&_Vm.CallOpts, account)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address account) view returns(uint64 nonce)
func (_Vm *VmCallerSession) GetNonce(account common.Address) (uint64, error) {
	return _Vm.Contract.GetNonce(&_Vm.CallOpts, account)
}

// GetStateDiff is a free data retrieval call binding the contract method 0x80df01cc.
//
// Solidity: function getStateDiff() view returns(string diff)
func (_Vm *VmCaller) GetStateDiff(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getStateDiff")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetStateDiff is a free data retrieval call binding the contract method 0x80df01cc.
//
// Solidity: function getStateDiff() view returns(string diff)
func (_Vm *VmSession) GetStateDiff() (string, error) {
	return _Vm.Contract.GetStateDiff(&_Vm.CallOpts)
}

// GetStateDiff is a free data retrieval call binding the contract method 0x80df01cc.
//
// Solidity: function getStateDiff() view returns(string diff)
func (_Vm *VmCallerSession) GetStateDiff() (string, error) {
	return _Vm.Contract.GetStateDiff(&_Vm.CallOpts)
}

// GetStateDiffJson is a free data retrieval call binding the contract method 0xf54fe009.
//
// Solidity: function getStateDiffJson() view returns(string diff)
func (_Vm *VmCaller) GetStateDiffJson(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "getStateDiffJson")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetStateDiffJson is a free data retrieval call binding the contract method 0xf54fe009.
//
// Solidity: function getStateDiffJson() view returns(string diff)
func (_Vm *VmSession) GetStateDiffJson() (string, error) {
	return _Vm.Contract.GetStateDiffJson(&_Vm.CallOpts)
}

// GetStateDiffJson is a free data retrieval call binding the contract method 0xf54fe009.
//
// Solidity: function getStateDiffJson() view returns(string diff)
func (_Vm *VmCallerSession) GetStateDiffJson() (string, error) {
	return _Vm.Contract.GetStateDiffJson(&_Vm.CallOpts)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_Vm *VmCaller) IndexOf(opts *bind.CallOpts, input string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "indexOf", input, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_Vm *VmSession) IndexOf(input string, key string) (*big.Int, error) {
	return _Vm.Contract.IndexOf(&_Vm.CallOpts, input, key)
}

// IndexOf is a free data retrieval call binding the contract method 0x8a0807b7.
//
// Solidity: function indexOf(string input, string key) pure returns(uint256)
func (_Vm *VmCallerSession) IndexOf(input string, key string) (*big.Int, error) {
	return _Vm.Contract.IndexOf(&_Vm.CallOpts, input, key)
}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_Vm *VmCaller) IsContext(opts *bind.CallOpts, context uint8) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "isContext", context)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_Vm *VmSession) IsContext(context uint8) (bool, error) {
	return _Vm.Contract.IsContext(&_Vm.CallOpts, context)
}

// IsContext is a free data retrieval call binding the contract method 0x64af255d.
//
// Solidity: function isContext(uint8 context) view returns(bool result)
func (_Vm *VmCallerSession) IsContext(context uint8) (bool, error) {
	return _Vm.Contract.IsContext(&_Vm.CallOpts, context)
}

// IsDir is a free data retrieval call binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) view returns(bool result)
func (_Vm *VmCaller) IsDir(opts *bind.CallOpts, path string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "isDir", path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDir is a free data retrieval call binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) view returns(bool result)
func (_Vm *VmSession) IsDir(path string) (bool, error) {
	return _Vm.Contract.IsDir(&_Vm.CallOpts, path)
}

// IsDir is a free data retrieval call binding the contract method 0x7d15d019.
//
// Solidity: function isDir(string path) view returns(bool result)
func (_Vm *VmCallerSession) IsDir(path string) (bool, error) {
	return _Vm.Contract.IsDir(&_Vm.CallOpts, path)
}

// IsFile is a free data retrieval call binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) view returns(bool result)
func (_Vm *VmCaller) IsFile(opts *bind.CallOpts, path string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "isFile", path)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFile is a free data retrieval call binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) view returns(bool result)
func (_Vm *VmSession) IsFile(path string) (bool, error) {
	return _Vm.Contract.IsFile(&_Vm.CallOpts, path)
}

// IsFile is a free data retrieval call binding the contract method 0xe0eb04d4.
//
// Solidity: function isFile(string path) view returns(bool result)
func (_Vm *VmCallerSession) IsFile(path string) (bool, error) {
	return _Vm.Contract.IsFile(&_Vm.CallOpts, path)
}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmCaller) IsPersistent(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "isPersistent", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmSession) IsPersistent(account common.Address) (bool, error) {
	return _Vm.Contract.IsPersistent(&_Vm.CallOpts, account)
}

// IsPersistent is a free data retrieval call binding the contract method 0xd92d8efd.
//
// Solidity: function isPersistent(address account) view returns(bool persistent)
func (_Vm *VmCallerSession) IsPersistent(account common.Address) (bool, error) {
	return _Vm.Contract.IsPersistent(&_Vm.CallOpts, account)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmCaller) KeyExists(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "keyExists", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmSession) KeyExists(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExists(&_Vm.CallOpts, json, key)
}

// KeyExists is a free data retrieval call binding the contract method 0x528a683c.
//
// Solidity: function keyExists(string json, string key) view returns(bool)
func (_Vm *VmCallerSession) KeyExists(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExists(&_Vm.CallOpts, json, key)
}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_Vm *VmCaller) KeyExistsJson(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "keyExistsJson", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_Vm *VmSession) KeyExistsJson(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsJson(&_Vm.CallOpts, json, key)
}

// KeyExistsJson is a free data retrieval call binding the contract method 0xdb4235f6.
//
// Solidity: function keyExistsJson(string json, string key) view returns(bool)
func (_Vm *VmCallerSession) KeyExistsJson(json string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsJson(&_Vm.CallOpts, json, key)
}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_Vm *VmCaller) KeyExistsToml(opts *bind.CallOpts, toml string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "keyExistsToml", toml, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_Vm *VmSession) KeyExistsToml(toml string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsToml(&_Vm.CallOpts, toml, key)
}

// KeyExistsToml is a free data retrieval call binding the contract method 0x600903ad.
//
// Solidity: function keyExistsToml(string toml, string key) view returns(bool)
func (_Vm *VmCallerSession) KeyExistsToml(toml string, key string) (bool, error) {
	return _Vm.Contract.KeyExistsToml(&_Vm.CallOpts, toml, key)
}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_Vm *VmCaller) LastCallGas(opts *bind.CallOpts) (VmSafeGas, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "lastCallGas")

	if err != nil {
		return *new(VmSafeGas), err
	}

	out0 := *abi.ConvertType(out[0], new(VmSafeGas)).(*VmSafeGas)

	return out0, err

}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_Vm *VmSession) LastCallGas() (VmSafeGas, error) {
	return _Vm.Contract.LastCallGas(&_Vm.CallOpts)
}

// LastCallGas is a free data retrieval call binding the contract method 0x2b589b28.
//
// Solidity: function lastCallGas() view returns((uint64,uint64,uint64,int64,uint64) gas)
func (_Vm *VmCallerSession) LastCallGas() (VmSafeGas, error) {
	return _Vm.Contract.LastCallGas(&_Vm.CallOpts)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmCaller) Load(opts *bind.CallOpts, target common.Address, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "load", target, slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _Vm.Contract.Load(&_Vm.CallOpts, target, slot)
}

// Load is a free data retrieval call binding the contract method 0x667f9d70.
//
// Solidity: function load(address target, bytes32 slot) view returns(bytes32 data)
func (_Vm *VmCallerSession) Load(target common.Address, slot [32]byte) ([32]byte, error) {
	return _Vm.Contract.Load(&_Vm.CallOpts, target, slot)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmCaller) ParseAddress(opts *bind.CallOpts, stringifiedValue string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseAddress", stringifiedValue)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _Vm.Contract.ParseAddress(&_Vm.CallOpts, stringifiedValue)
}

// ParseAddress is a free data retrieval call binding the contract method 0xc6ce059d.
//
// Solidity: function parseAddress(string stringifiedValue) pure returns(address parsedValue)
func (_Vm *VmCallerSession) ParseAddress(stringifiedValue string) (common.Address, error) {
	return _Vm.Contract.ParseAddress(&_Vm.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmCaller) ParseBool(opts *bind.CallOpts, stringifiedValue string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBool", stringifiedValue)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmSession) ParseBool(stringifiedValue string) (bool, error) {
	return _Vm.Contract.ParseBool(&_Vm.CallOpts, stringifiedValue)
}

// ParseBool is a free data retrieval call binding the contract method 0x974ef924.
//
// Solidity: function parseBool(string stringifiedValue) pure returns(bool parsedValue)
func (_Vm *VmCallerSession) ParseBool(stringifiedValue string) (bool, error) {
	return _Vm.Contract.ParseBool(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmCaller) ParseBytes(opts *bind.CallOpts, stringifiedValue string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBytes", stringifiedValue)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _Vm.Contract.ParseBytes(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes is a free data retrieval call binding the contract method 0x8f5d232d.
//
// Solidity: function parseBytes(string stringifiedValue) pure returns(bytes parsedValue)
func (_Vm *VmCallerSession) ParseBytes(stringifiedValue string) ([]byte, error) {
	return _Vm.Contract.ParseBytes(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmCaller) ParseBytes32(opts *bind.CallOpts, stringifiedValue string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseBytes32", stringifiedValue)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _Vm.Contract.ParseBytes32(&_Vm.CallOpts, stringifiedValue)
}

// ParseBytes32 is a free data retrieval call binding the contract method 0x087e6e81.
//
// Solidity: function parseBytes32(string stringifiedValue) pure returns(bytes32 parsedValue)
func (_Vm *VmCallerSession) ParseBytes32(stringifiedValue string) ([32]byte, error) {
	return _Vm.Contract.ParseBytes32(&_Vm.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmCaller) ParseInt(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseInt", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseInt(&_Vm.CallOpts, stringifiedValue)
}

// ParseInt is a free data retrieval call binding the contract method 0x42346c5e.
//
// Solidity: function parseInt(string stringifiedValue) pure returns(int256 parsedValue)
func (_Vm *VmCallerSession) ParseInt(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseInt(&_Vm.CallOpts, stringifiedValue)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseJson(opts *bind.CallOpts, json string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJson", json)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseJson(json string) ([]byte, error) {
	return _Vm.Contract.ParseJson(&_Vm.CallOpts, json)
}

// ParseJson is a free data retrieval call binding the contract method 0x6a82600a.
//
// Solidity: function parseJson(string json) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseJson(json string) ([]byte, error) {
	return _Vm.Contract.ParseJson(&_Vm.CallOpts, json)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseJson0(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJson0", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseJson0(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJson0(&_Vm.CallOpts, json, key)
}

// ParseJson0 is a free data retrieval call binding the contract method 0x85940ef1.
//
// Solidity: function parseJson(string json, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseJson0(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJson0(&_Vm.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmCaller) ParseJsonAddress(opts *bind.CallOpts, json string, key string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonAddress", json, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _Vm.Contract.ParseJsonAddress(&_Vm.CallOpts, json, key)
}

// ParseJsonAddress is a free data retrieval call binding the contract method 0x1e19e657.
//
// Solidity: function parseJsonAddress(string json, string key) pure returns(address)
func (_Vm *VmCallerSession) ParseJsonAddress(json string, key string) (common.Address, error) {
	return _Vm.Contract.ParseJsonAddress(&_Vm.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmCaller) ParseJsonAddressArray(opts *bind.CallOpts, json string, key string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonAddressArray", json, key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseJsonAddressArray(&_Vm.CallOpts, json, key)
}

// ParseJsonAddressArray is a free data retrieval call binding the contract method 0x2fce7883.
//
// Solidity: function parseJsonAddressArray(string json, string key) pure returns(address[])
func (_Vm *VmCallerSession) ParseJsonAddressArray(json string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseJsonAddressArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmCaller) ParseJsonBool(opts *bind.CallOpts, json string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBool", json, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmSession) ParseJsonBool(json string, key string) (bool, error) {
	return _Vm.Contract.ParseJsonBool(&_Vm.CallOpts, json, key)
}

// ParseJsonBool is a free data retrieval call binding the contract method 0x9f86dc91.
//
// Solidity: function parseJsonBool(string json, string key) pure returns(bool)
func (_Vm *VmCallerSession) ParseJsonBool(json string, key string) (bool, error) {
	return _Vm.Contract.ParseJsonBool(&_Vm.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmCaller) ParseJsonBoolArray(opts *bind.CallOpts, json string, key string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBoolArray", json, key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _Vm.Contract.ParseJsonBoolArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBoolArray is a free data retrieval call binding the contract method 0x91f3b94f.
//
// Solidity: function parseJsonBoolArray(string json, string key) pure returns(bool[])
func (_Vm *VmCallerSession) ParseJsonBoolArray(json string, key string) ([]bool, error) {
	return _Vm.Contract.ParseJsonBoolArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmCaller) ParseJsonBytes(opts *bind.CallOpts, json string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes", json, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJsonBytes(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes is a free data retrieval call binding the contract method 0xfd921be8.
//
// Solidity: function parseJsonBytes(string json, string key) pure returns(bytes)
func (_Vm *VmCallerSession) ParseJsonBytes(json string, key string) ([]byte, error) {
	return _Vm.Contract.ParseJsonBytes(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmCaller) ParseJsonBytes32(opts *bind.CallOpts, json string, key string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes32", json, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32 is a free data retrieval call binding the contract method 0x1777e59d.
//
// Solidity: function parseJsonBytes32(string json, string key) pure returns(bytes32)
func (_Vm *VmCallerSession) ParseJsonBytes32(json string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmCaller) ParseJsonBytes32Array(opts *bind.CallOpts, json string, key string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytes32Array", json, key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32Array(&_Vm.CallOpts, json, key)
}

// ParseJsonBytes32Array is a free data retrieval call binding the contract method 0x91c75bc3.
//
// Solidity: function parseJsonBytes32Array(string json, string key) pure returns(bytes32[])
func (_Vm *VmCallerSession) ParseJsonBytes32Array(json string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseJsonBytes32Array(&_Vm.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmCaller) ParseJsonBytesArray(opts *bind.CallOpts, json string, key string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonBytesArray", json, key)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseJsonBytesArray(&_Vm.CallOpts, json, key)
}

// ParseJsonBytesArray is a free data retrieval call binding the contract method 0x6631aa99.
//
// Solidity: function parseJsonBytesArray(string json, string key) pure returns(bytes[])
func (_Vm *VmCallerSession) ParseJsonBytesArray(json string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseJsonBytesArray(&_Vm.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmCaller) ParseJsonInt(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonInt", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonInt(&_Vm.CallOpts, json, key)
}

// ParseJsonInt is a free data retrieval call binding the contract method 0x7b048ccd.
//
// Solidity: function parseJsonInt(string json, string key) pure returns(int256)
func (_Vm *VmCallerSession) ParseJsonInt(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonInt(&_Vm.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmCaller) ParseJsonIntArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonIntArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonIntArray(&_Vm.CallOpts, json, key)
}

// ParseJsonIntArray is a free data retrieval call binding the contract method 0x9983c28a.
//
// Solidity: function parseJsonIntArray(string json, string key) pure returns(int256[])
func (_Vm *VmCallerSession) ParseJsonIntArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonIntArray(&_Vm.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmCaller) ParseJsonKeys(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonKeys", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonKeys(&_Vm.CallOpts, json, key)
}

// ParseJsonKeys is a free data retrieval call binding the contract method 0x213e4198.
//
// Solidity: function parseJsonKeys(string json, string key) pure returns(string[] keys)
func (_Vm *VmCallerSession) ParseJsonKeys(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonKeys(&_Vm.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmCaller) ParseJsonString(opts *bind.CallOpts, json string, key string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonString", json, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmSession) ParseJsonString(json string, key string) (string, error) {
	return _Vm.Contract.ParseJsonString(&_Vm.CallOpts, json, key)
}

// ParseJsonString is a free data retrieval call binding the contract method 0x49c4fac8.
//
// Solidity: function parseJsonString(string json, string key) pure returns(string)
func (_Vm *VmCallerSession) ParseJsonString(json string, key string) (string, error) {
	return _Vm.Contract.ParseJsonString(&_Vm.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmCaller) ParseJsonStringArray(opts *bind.CallOpts, json string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonStringArray", json, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonStringArray(&_Vm.CallOpts, json, key)
}

// ParseJsonStringArray is a free data retrieval call binding the contract method 0x498fdcf4.
//
// Solidity: function parseJsonStringArray(string json, string key) pure returns(string[])
func (_Vm *VmCallerSession) ParseJsonStringArray(json string, key string) ([]string, error) {
	return _Vm.Contract.ParseJsonStringArray(&_Vm.CallOpts, json, key)
}

// ParseJsonType is a free data retrieval call binding the contract method 0xa9da313b.
//
// Solidity: function parseJsonType(string json, string typeDescription) pure returns(bytes)
func (_Vm *VmCaller) ParseJsonType(opts *bind.CallOpts, json string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonType", json, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonType is a free data retrieval call binding the contract method 0xa9da313b.
//
// Solidity: function parseJsonType(string json, string typeDescription) pure returns(bytes)
func (_Vm *VmSession) ParseJsonType(json string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseJsonType(&_Vm.CallOpts, json, typeDescription)
}

// ParseJsonType is a free data retrieval call binding the contract method 0xa9da313b.
//
// Solidity: function parseJsonType(string json, string typeDescription) pure returns(bytes)
func (_Vm *VmCallerSession) ParseJsonType(json string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseJsonType(&_Vm.CallOpts, json, typeDescription)
}

// ParseJsonType0 is a free data retrieval call binding the contract method 0xe3f5ae33.
//
// Solidity: function parseJsonType(string json, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCaller) ParseJsonType0(opts *bind.CallOpts, json string, key string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonType0", json, key, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonType0 is a free data retrieval call binding the contract method 0xe3f5ae33.
//
// Solidity: function parseJsonType(string json, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmSession) ParseJsonType0(json string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseJsonType0(&_Vm.CallOpts, json, key, typeDescription)
}

// ParseJsonType0 is a free data retrieval call binding the contract method 0xe3f5ae33.
//
// Solidity: function parseJsonType(string json, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCallerSession) ParseJsonType0(json string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseJsonType0(&_Vm.CallOpts, json, key, typeDescription)
}

// ParseJsonTypeArray is a free data retrieval call binding the contract method 0x0175d535.
//
// Solidity: function parseJsonTypeArray(string json, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCaller) ParseJsonTypeArray(opts *bind.CallOpts, json string, key string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonTypeArray", json, key, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseJsonTypeArray is a free data retrieval call binding the contract method 0x0175d535.
//
// Solidity: function parseJsonTypeArray(string json, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmSession) ParseJsonTypeArray(json string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseJsonTypeArray(&_Vm.CallOpts, json, key, typeDescription)
}

// ParseJsonTypeArray is a free data retrieval call binding the contract method 0x0175d535.
//
// Solidity: function parseJsonTypeArray(string json, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCallerSession) ParseJsonTypeArray(json string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseJsonTypeArray(&_Vm.CallOpts, json, key, typeDescription)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmCaller) ParseJsonUint(opts *bind.CallOpts, json string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonUint", json, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonUint(&_Vm.CallOpts, json, key)
}

// ParseJsonUint is a free data retrieval call binding the contract method 0xaddde2b6.
//
// Solidity: function parseJsonUint(string json, string key) pure returns(uint256)
func (_Vm *VmCallerSession) ParseJsonUint(json string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseJsonUint(&_Vm.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmCaller) ParseJsonUintArray(opts *bind.CallOpts, json string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseJsonUintArray", json, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonUintArray(&_Vm.CallOpts, json, key)
}

// ParseJsonUintArray is a free data retrieval call binding the contract method 0x522074ab.
//
// Solidity: function parseJsonUintArray(string json, string key) pure returns(uint256[])
func (_Vm *VmCallerSession) ParseJsonUintArray(json string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseJsonUintArray(&_Vm.CallOpts, json, key)
}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseToml(opts *bind.CallOpts, toml string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseToml", toml, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseToml(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseToml(&_Vm.CallOpts, toml, key)
}

// ParseToml is a free data retrieval call binding the contract method 0x37736e08.
//
// Solidity: function parseToml(string toml, string key) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseToml(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseToml(&_Vm.CallOpts, toml, key)
}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_Vm *VmCaller) ParseToml0(opts *bind.CallOpts, toml string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseToml0", toml)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_Vm *VmSession) ParseToml0(toml string) ([]byte, error) {
	return _Vm.Contract.ParseToml0(&_Vm.CallOpts, toml)
}

// ParseToml0 is a free data retrieval call binding the contract method 0x592151f0.
//
// Solidity: function parseToml(string toml) pure returns(bytes abiEncodedData)
func (_Vm *VmCallerSession) ParseToml0(toml string) ([]byte, error) {
	return _Vm.Contract.ParseToml0(&_Vm.CallOpts, toml)
}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_Vm *VmCaller) ParseTomlAddress(opts *bind.CallOpts, toml string, key string) (common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlAddress", toml, key)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_Vm *VmSession) ParseTomlAddress(toml string, key string) (common.Address, error) {
	return _Vm.Contract.ParseTomlAddress(&_Vm.CallOpts, toml, key)
}

// ParseTomlAddress is a free data retrieval call binding the contract method 0x65e7c844.
//
// Solidity: function parseTomlAddress(string toml, string key) pure returns(address)
func (_Vm *VmCallerSession) ParseTomlAddress(toml string, key string) (common.Address, error) {
	return _Vm.Contract.ParseTomlAddress(&_Vm.CallOpts, toml, key)
}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_Vm *VmCaller) ParseTomlAddressArray(opts *bind.CallOpts, toml string, key string) ([]common.Address, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlAddressArray", toml, key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_Vm *VmSession) ParseTomlAddressArray(toml string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseTomlAddressArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlAddressArray is a free data retrieval call binding the contract method 0x65c428e7.
//
// Solidity: function parseTomlAddressArray(string toml, string key) pure returns(address[])
func (_Vm *VmCallerSession) ParseTomlAddressArray(toml string, key string) ([]common.Address, error) {
	return _Vm.Contract.ParseTomlAddressArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_Vm *VmCaller) ParseTomlBool(opts *bind.CallOpts, toml string, key string) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBool", toml, key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_Vm *VmSession) ParseTomlBool(toml string, key string) (bool, error) {
	return _Vm.Contract.ParseTomlBool(&_Vm.CallOpts, toml, key)
}

// ParseTomlBool is a free data retrieval call binding the contract method 0xd30dced6.
//
// Solidity: function parseTomlBool(string toml, string key) pure returns(bool)
func (_Vm *VmCallerSession) ParseTomlBool(toml string, key string) (bool, error) {
	return _Vm.Contract.ParseTomlBool(&_Vm.CallOpts, toml, key)
}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_Vm *VmCaller) ParseTomlBoolArray(opts *bind.CallOpts, toml string, key string) ([]bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBoolArray", toml, key)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_Vm *VmSession) ParseTomlBoolArray(toml string, key string) ([]bool, error) {
	return _Vm.Contract.ParseTomlBoolArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBoolArray is a free data retrieval call binding the contract method 0x127cfe9a.
//
// Solidity: function parseTomlBoolArray(string toml, string key) pure returns(bool[])
func (_Vm *VmCallerSession) ParseTomlBoolArray(toml string, key string) ([]bool, error) {
	return _Vm.Contract.ParseTomlBoolArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_Vm *VmCaller) ParseTomlBytes(opts *bind.CallOpts, toml string, key string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytes", toml, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_Vm *VmSession) ParseTomlBytes(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseTomlBytes(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes is a free data retrieval call binding the contract method 0xd77bfdb9.
//
// Solidity: function parseTomlBytes(string toml, string key) pure returns(bytes)
func (_Vm *VmCallerSession) ParseTomlBytes(toml string, key string) ([]byte, error) {
	return _Vm.Contract.ParseTomlBytes(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_Vm *VmCaller) ParseTomlBytes32(opts *bind.CallOpts, toml string, key string) ([32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytes32", toml, key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_Vm *VmSession) ParseTomlBytes32(toml string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32 is a free data retrieval call binding the contract method 0x8e214810.
//
// Solidity: function parseTomlBytes32(string toml, string key) pure returns(bytes32)
func (_Vm *VmCallerSession) ParseTomlBytes32(toml string, key string) ([32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_Vm *VmCaller) ParseTomlBytes32Array(opts *bind.CallOpts, toml string, key string) ([][32]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytes32Array", toml, key)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_Vm *VmSession) ParseTomlBytes32Array(toml string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32Array(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytes32Array is a free data retrieval call binding the contract method 0x3e716f81.
//
// Solidity: function parseTomlBytes32Array(string toml, string key) pure returns(bytes32[])
func (_Vm *VmCallerSession) ParseTomlBytes32Array(toml string, key string) ([][32]byte, error) {
	return _Vm.Contract.ParseTomlBytes32Array(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_Vm *VmCaller) ParseTomlBytesArray(opts *bind.CallOpts, toml string, key string) ([][]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlBytesArray", toml, key)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_Vm *VmSession) ParseTomlBytesArray(toml string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseTomlBytesArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlBytesArray is a free data retrieval call binding the contract method 0xb197c247.
//
// Solidity: function parseTomlBytesArray(string toml, string key) pure returns(bytes[])
func (_Vm *VmCallerSession) ParseTomlBytesArray(toml string, key string) ([][]byte, error) {
	return _Vm.Contract.ParseTomlBytesArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_Vm *VmCaller) ParseTomlInt(opts *bind.CallOpts, toml string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlInt", toml, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_Vm *VmSession) ParseTomlInt(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlInt(&_Vm.CallOpts, toml, key)
}

// ParseTomlInt is a free data retrieval call binding the contract method 0xc1350739.
//
// Solidity: function parseTomlInt(string toml, string key) pure returns(int256)
func (_Vm *VmCallerSession) ParseTomlInt(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlInt(&_Vm.CallOpts, toml, key)
}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_Vm *VmCaller) ParseTomlIntArray(opts *bind.CallOpts, toml string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlIntArray", toml, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_Vm *VmSession) ParseTomlIntArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlIntArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlIntArray is a free data retrieval call binding the contract method 0xd3522ae6.
//
// Solidity: function parseTomlIntArray(string toml, string key) pure returns(int256[])
func (_Vm *VmCallerSession) ParseTomlIntArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlIntArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_Vm *VmCaller) ParseTomlKeys(opts *bind.CallOpts, toml string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlKeys", toml, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_Vm *VmSession) ParseTomlKeys(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlKeys(&_Vm.CallOpts, toml, key)
}

// ParseTomlKeys is a free data retrieval call binding the contract method 0x812a44b2.
//
// Solidity: function parseTomlKeys(string toml, string key) pure returns(string[] keys)
func (_Vm *VmCallerSession) ParseTomlKeys(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlKeys(&_Vm.CallOpts, toml, key)
}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_Vm *VmCaller) ParseTomlString(opts *bind.CallOpts, toml string, key string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlString", toml, key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_Vm *VmSession) ParseTomlString(toml string, key string) (string, error) {
	return _Vm.Contract.ParseTomlString(&_Vm.CallOpts, toml, key)
}

// ParseTomlString is a free data retrieval call binding the contract method 0x8bb8dd43.
//
// Solidity: function parseTomlString(string toml, string key) pure returns(string)
func (_Vm *VmCallerSession) ParseTomlString(toml string, key string) (string, error) {
	return _Vm.Contract.ParseTomlString(&_Vm.CallOpts, toml, key)
}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_Vm *VmCaller) ParseTomlStringArray(opts *bind.CallOpts, toml string, key string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlStringArray", toml, key)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_Vm *VmSession) ParseTomlStringArray(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlStringArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlStringArray is a free data retrieval call binding the contract method 0x9f629281.
//
// Solidity: function parseTomlStringArray(string toml, string key) pure returns(string[])
func (_Vm *VmCallerSession) ParseTomlStringArray(toml string, key string) ([]string, error) {
	return _Vm.Contract.ParseTomlStringArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlType is a free data retrieval call binding the contract method 0x47fa5e11.
//
// Solidity: function parseTomlType(string toml, string typeDescription) pure returns(bytes)
func (_Vm *VmCaller) ParseTomlType(opts *bind.CallOpts, toml string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlType", toml, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseTomlType is a free data retrieval call binding the contract method 0x47fa5e11.
//
// Solidity: function parseTomlType(string toml, string typeDescription) pure returns(bytes)
func (_Vm *VmSession) ParseTomlType(toml string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseTomlType(&_Vm.CallOpts, toml, typeDescription)
}

// ParseTomlType is a free data retrieval call binding the contract method 0x47fa5e11.
//
// Solidity: function parseTomlType(string toml, string typeDescription) pure returns(bytes)
func (_Vm *VmCallerSession) ParseTomlType(toml string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseTomlType(&_Vm.CallOpts, toml, typeDescription)
}

// ParseTomlType0 is a free data retrieval call binding the contract method 0xf9fa5cdb.
//
// Solidity: function parseTomlType(string toml, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCaller) ParseTomlType0(opts *bind.CallOpts, toml string, key string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlType0", toml, key, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseTomlType0 is a free data retrieval call binding the contract method 0xf9fa5cdb.
//
// Solidity: function parseTomlType(string toml, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmSession) ParseTomlType0(toml string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseTomlType0(&_Vm.CallOpts, toml, key, typeDescription)
}

// ParseTomlType0 is a free data retrieval call binding the contract method 0xf9fa5cdb.
//
// Solidity: function parseTomlType(string toml, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCallerSession) ParseTomlType0(toml string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseTomlType0(&_Vm.CallOpts, toml, key, typeDescription)
}

// ParseTomlTypeArray is a free data retrieval call binding the contract method 0x49be3743.
//
// Solidity: function parseTomlTypeArray(string toml, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCaller) ParseTomlTypeArray(opts *bind.CallOpts, toml string, key string, typeDescription string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlTypeArray", toml, key, typeDescription)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ParseTomlTypeArray is a free data retrieval call binding the contract method 0x49be3743.
//
// Solidity: function parseTomlTypeArray(string toml, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmSession) ParseTomlTypeArray(toml string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseTomlTypeArray(&_Vm.CallOpts, toml, key, typeDescription)
}

// ParseTomlTypeArray is a free data retrieval call binding the contract method 0x49be3743.
//
// Solidity: function parseTomlTypeArray(string toml, string key, string typeDescription) pure returns(bytes)
func (_Vm *VmCallerSession) ParseTomlTypeArray(toml string, key string, typeDescription string) ([]byte, error) {
	return _Vm.Contract.ParseTomlTypeArray(&_Vm.CallOpts, toml, key, typeDescription)
}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_Vm *VmCaller) ParseTomlUint(opts *bind.CallOpts, toml string, key string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlUint", toml, key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_Vm *VmSession) ParseTomlUint(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlUint(&_Vm.CallOpts, toml, key)
}

// ParseTomlUint is a free data retrieval call binding the contract method 0xcc7b0487.
//
// Solidity: function parseTomlUint(string toml, string key) pure returns(uint256)
func (_Vm *VmCallerSession) ParseTomlUint(toml string, key string) (*big.Int, error) {
	return _Vm.Contract.ParseTomlUint(&_Vm.CallOpts, toml, key)
}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_Vm *VmCaller) ParseTomlUintArray(opts *bind.CallOpts, toml string, key string) ([]*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseTomlUintArray", toml, key)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_Vm *VmSession) ParseTomlUintArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlUintArray(&_Vm.CallOpts, toml, key)
}

// ParseTomlUintArray is a free data retrieval call binding the contract method 0xb5df27c8.
//
// Solidity: function parseTomlUintArray(string toml, string key) pure returns(uint256[])
func (_Vm *VmCallerSession) ParseTomlUintArray(toml string, key string) ([]*big.Int, error) {
	return _Vm.Contract.ParseTomlUintArray(&_Vm.CallOpts, toml, key)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmCaller) ParseUint(opts *bind.CallOpts, stringifiedValue string) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "parseUint", stringifiedValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseUint(&_Vm.CallOpts, stringifiedValue)
}

// ParseUint is a free data retrieval call binding the contract method 0xfa91454d.
//
// Solidity: function parseUint(string stringifiedValue) pure returns(uint256 parsedValue)
func (_Vm *VmCallerSession) ParseUint(stringifiedValue string) (*big.Int, error) {
	return _Vm.Contract.ParseUint(&_Vm.CallOpts, stringifiedValue)
}

// PauseTracing is a free data retrieval call binding the contract method 0xc94d1f90.
//
// Solidity: function pauseTracing() view returns()
func (_Vm *VmCaller) PauseTracing(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "pauseTracing")

	if err != nil {
		return err
	}

	return err

}

// PauseTracing is a free data retrieval call binding the contract method 0xc94d1f90.
//
// Solidity: function pauseTracing() view returns()
func (_Vm *VmSession) PauseTracing() error {
	return _Vm.Contract.PauseTracing(&_Vm.CallOpts)
}

// PauseTracing is a free data retrieval call binding the contract method 0xc94d1f90.
//
// Solidity: function pauseTracing() view returns()
func (_Vm *VmCallerSession) PauseTracing() error {
	return _Vm.Contract.PauseTracing(&_Vm.CallOpts)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmCaller) ProjectRoot(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "projectRoot")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmSession) ProjectRoot() (string, error) {
	return _Vm.Contract.ProjectRoot(&_Vm.CallOpts)
}

// ProjectRoot is a free data retrieval call binding the contract method 0xd930a0e6.
//
// Solidity: function projectRoot() view returns(string path)
func (_Vm *VmCallerSession) ProjectRoot() (string, error) {
	return _Vm.Contract.ProjectRoot(&_Vm.CallOpts)
}

// PublicKeyP256 is a free data retrieval call binding the contract method 0xc453949e.
//
// Solidity: function publicKeyP256(uint256 privateKey) pure returns(uint256 publicKeyX, uint256 publicKeyY)
func (_Vm *VmCaller) PublicKeyP256(opts *bind.CallOpts, privateKey *big.Int) (struct {
	PublicKeyX *big.Int
	PublicKeyY *big.Int
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "publicKeyP256", privateKey)

	outstruct := new(struct {
		PublicKeyX *big.Int
		PublicKeyY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PublicKeyX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PublicKeyY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PublicKeyP256 is a free data retrieval call binding the contract method 0xc453949e.
//
// Solidity: function publicKeyP256(uint256 privateKey) pure returns(uint256 publicKeyX, uint256 publicKeyY)
func (_Vm *VmSession) PublicKeyP256(privateKey *big.Int) (struct {
	PublicKeyX *big.Int
	PublicKeyY *big.Int
}, error) {
	return _Vm.Contract.PublicKeyP256(&_Vm.CallOpts, privateKey)
}

// PublicKeyP256 is a free data retrieval call binding the contract method 0xc453949e.
//
// Solidity: function publicKeyP256(uint256 privateKey) pure returns(uint256 publicKeyX, uint256 publicKeyY)
func (_Vm *VmCallerSession) PublicKeyP256(privateKey *big.Int) (struct {
	PublicKeyX *big.Int
	PublicKeyY *big.Int
}, error) {
	return _Vm.Contract.PublicKeyP256(&_Vm.CallOpts, privateKey)
}

// RandomBool is a free data retrieval call binding the contract method 0xcdc126bd.
//
// Solidity: function randomBool() view returns(bool)
func (_Vm *VmCaller) RandomBool(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "randomBool")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RandomBool is a free data retrieval call binding the contract method 0xcdc126bd.
//
// Solidity: function randomBool() view returns(bool)
func (_Vm *VmSession) RandomBool() (bool, error) {
	return _Vm.Contract.RandomBool(&_Vm.CallOpts)
}

// RandomBool is a free data retrieval call binding the contract method 0xcdc126bd.
//
// Solidity: function randomBool() view returns(bool)
func (_Vm *VmCallerSession) RandomBool() (bool, error) {
	return _Vm.Contract.RandomBool(&_Vm.CallOpts)
}

// RandomBytes is a free data retrieval call binding the contract method 0x6c5d32a9.
//
// Solidity: function randomBytes(uint256 len) view returns(bytes)
func (_Vm *VmCaller) RandomBytes(opts *bind.CallOpts, len *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "randomBytes", len)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// RandomBytes is a free data retrieval call binding the contract method 0x6c5d32a9.
//
// Solidity: function randomBytes(uint256 len) view returns(bytes)
func (_Vm *VmSession) RandomBytes(len *big.Int) ([]byte, error) {
	return _Vm.Contract.RandomBytes(&_Vm.CallOpts, len)
}

// RandomBytes is a free data retrieval call binding the contract method 0x6c5d32a9.
//
// Solidity: function randomBytes(uint256 len) view returns(bytes)
func (_Vm *VmCallerSession) RandomBytes(len *big.Int) ([]byte, error) {
	return _Vm.Contract.RandomBytes(&_Vm.CallOpts, len)
}

// RandomBytes4 is a free data retrieval call binding the contract method 0x9b7cd579.
//
// Solidity: function randomBytes4() view returns(bytes4)
func (_Vm *VmCaller) RandomBytes4(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "randomBytes4")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// RandomBytes4 is a free data retrieval call binding the contract method 0x9b7cd579.
//
// Solidity: function randomBytes4() view returns(bytes4)
func (_Vm *VmSession) RandomBytes4() ([4]byte, error) {
	return _Vm.Contract.RandomBytes4(&_Vm.CallOpts)
}

// RandomBytes4 is a free data retrieval call binding the contract method 0x9b7cd579.
//
// Solidity: function randomBytes4() view returns(bytes4)
func (_Vm *VmCallerSession) RandomBytes4() ([4]byte, error) {
	return _Vm.Contract.RandomBytes4(&_Vm.CallOpts)
}

// RandomBytes8 is a free data retrieval call binding the contract method 0x0497b0a5.
//
// Solidity: function randomBytes8() view returns(bytes8)
func (_Vm *VmCaller) RandomBytes8(opts *bind.CallOpts) ([8]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "randomBytes8")

	if err != nil {
		return *new([8]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([8]byte)).(*[8]byte)

	return out0, err

}

// RandomBytes8 is a free data retrieval call binding the contract method 0x0497b0a5.
//
// Solidity: function randomBytes8() view returns(bytes8)
func (_Vm *VmSession) RandomBytes8() ([8]byte, error) {
	return _Vm.Contract.RandomBytes8(&_Vm.CallOpts)
}

// RandomBytes8 is a free data retrieval call binding the contract method 0x0497b0a5.
//
// Solidity: function randomBytes8() view returns(bytes8)
func (_Vm *VmCallerSession) RandomBytes8() ([8]byte, error) {
	return _Vm.Contract.RandomBytes8(&_Vm.CallOpts)
}

// RandomInt is a free data retrieval call binding the contract method 0x111f1202.
//
// Solidity: function randomInt() view returns(int256)
func (_Vm *VmCaller) RandomInt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "randomInt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomInt is a free data retrieval call binding the contract method 0x111f1202.
//
// Solidity: function randomInt() view returns(int256)
func (_Vm *VmSession) RandomInt() (*big.Int, error) {
	return _Vm.Contract.RandomInt(&_Vm.CallOpts)
}

// RandomInt is a free data retrieval call binding the contract method 0x111f1202.
//
// Solidity: function randomInt() view returns(int256)
func (_Vm *VmCallerSession) RandomInt() (*big.Int, error) {
	return _Vm.Contract.RandomInt(&_Vm.CallOpts)
}

// RandomInt0 is a free data retrieval call binding the contract method 0x12845966.
//
// Solidity: function randomInt(uint256 bits) view returns(int256)
func (_Vm *VmCaller) RandomInt0(opts *bind.CallOpts, bits *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "randomInt0", bits)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomInt0 is a free data retrieval call binding the contract method 0x12845966.
//
// Solidity: function randomInt(uint256 bits) view returns(int256)
func (_Vm *VmSession) RandomInt0(bits *big.Int) (*big.Int, error) {
	return _Vm.Contract.RandomInt0(&_Vm.CallOpts, bits)
}

// RandomInt0 is a free data retrieval call binding the contract method 0x12845966.
//
// Solidity: function randomInt(uint256 bits) view returns(int256)
func (_Vm *VmCallerSession) RandomInt0(bits *big.Int) (*big.Int, error) {
	return _Vm.Contract.RandomInt0(&_Vm.CallOpts, bits)
}

// RandomUint0 is a free data retrieval call binding the contract method 0xcf81e69c.
//
// Solidity: function randomUint(uint256 bits) view returns(uint256)
func (_Vm *VmCaller) RandomUint0(opts *bind.CallOpts, bits *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "randomUint0", bits)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomUint0 is a free data retrieval call binding the contract method 0xcf81e69c.
//
// Solidity: function randomUint(uint256 bits) view returns(uint256)
func (_Vm *VmSession) RandomUint0(bits *big.Int) (*big.Int, error) {
	return _Vm.Contract.RandomUint0(&_Vm.CallOpts, bits)
}

// RandomUint0 is a free data retrieval call binding the contract method 0xcf81e69c.
//
// Solidity: function randomUint(uint256 bits) view returns(uint256)
func (_Vm *VmCallerSession) RandomUint0(bits *big.Int) (*big.Int, error) {
	return _Vm.Contract.RandomUint0(&_Vm.CallOpts, bits)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir(opts *bind.CallOpts, path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir", path, maxDepth)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir(&_Vm.CallOpts, path, maxDepth)
}

// ReadDir is a free data retrieval call binding the contract method 0x1497876c.
//
// Solidity: function readDir(string path, uint64 maxDepth) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir(path string, maxDepth uint64) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir(&_Vm.CallOpts, path, maxDepth)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir0(opts *bind.CallOpts, path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir0", path, maxDepth, followLinks)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir0(&_Vm.CallOpts, path, maxDepth, followLinks)
}

// ReadDir0 is a free data retrieval call binding the contract method 0x8102d70d.
//
// Solidity: function readDir(string path, uint64 maxDepth, bool followLinks) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir0(path string, maxDepth uint64, followLinks bool) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir0(&_Vm.CallOpts, path, maxDepth, followLinks)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCaller) ReadDir1(opts *bind.CallOpts, path string) ([]VmSafeDirEntry, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readDir1", path)

	if err != nil {
		return *new([]VmSafeDirEntry), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeDirEntry)).(*[]VmSafeDirEntry)

	return out0, err

}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir1(&_Vm.CallOpts, path)
}

// ReadDir1 is a free data retrieval call binding the contract method 0xc4bc59e0.
//
// Solidity: function readDir(string path) view returns((string,string,uint64,bool,bool)[] entries)
func (_Vm *VmCallerSession) ReadDir1(path string) ([]VmSafeDirEntry, error) {
	return _Vm.Contract.ReadDir1(&_Vm.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmCaller) ReadFile(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readFile", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmSession) ReadFile(path string) (string, error) {
	return _Vm.Contract.ReadFile(&_Vm.CallOpts, path)
}

// ReadFile is a free data retrieval call binding the contract method 0x60f9bb11.
//
// Solidity: function readFile(string path) view returns(string data)
func (_Vm *VmCallerSession) ReadFile(path string) (string, error) {
	return _Vm.Contract.ReadFile(&_Vm.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmCaller) ReadFileBinary(opts *bind.CallOpts, path string) ([]byte, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readFileBinary", path)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmSession) ReadFileBinary(path string) ([]byte, error) {
	return _Vm.Contract.ReadFileBinary(&_Vm.CallOpts, path)
}

// ReadFileBinary is a free data retrieval call binding the contract method 0x16ed7bc4.
//
// Solidity: function readFileBinary(string path) view returns(bytes data)
func (_Vm *VmCallerSession) ReadFileBinary(path string) ([]byte, error) {
	return _Vm.Contract.ReadFileBinary(&_Vm.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmCaller) ReadLine(opts *bind.CallOpts, path string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readLine", path)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmSession) ReadLine(path string) (string, error) {
	return _Vm.Contract.ReadLine(&_Vm.CallOpts, path)
}

// ReadLine is a free data retrieval call binding the contract method 0x70f55728.
//
// Solidity: function readLine(string path) view returns(string line)
func (_Vm *VmCallerSession) ReadLine(path string) (string, error) {
	return _Vm.Contract.ReadLine(&_Vm.CallOpts, path)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmCaller) ReadLink(opts *bind.CallOpts, linkPath string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "readLink", linkPath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmSession) ReadLink(linkPath string) (string, error) {
	return _Vm.Contract.ReadLink(&_Vm.CallOpts, linkPath)
}

// ReadLink is a free data retrieval call binding the contract method 0x9f5684a2.
//
// Solidity: function readLink(string linkPath) view returns(string targetPath)
func (_Vm *VmCallerSession) ReadLink(linkPath string) (string, error) {
	return _Vm.Contract.ReadLink(&_Vm.CallOpts, linkPath)
}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_Vm *VmCaller) Replace(opts *bind.CallOpts, input string, from string, to string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "replace", input, from, to)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_Vm *VmSession) Replace(input string, from string, to string) (string, error) {
	return _Vm.Contract.Replace(&_Vm.CallOpts, input, from, to)
}

// Replace is a free data retrieval call binding the contract method 0xe00ad03e.
//
// Solidity: function replace(string input, string from, string to) pure returns(string output)
func (_Vm *VmCallerSession) Replace(input string, from string, to string) (string, error) {
	return _Vm.Contract.Replace(&_Vm.CallOpts, input, from, to)
}

// ResumeTracing is a free data retrieval call binding the contract method 0x72a09ccb.
//
// Solidity: function resumeTracing() view returns()
func (_Vm *VmCaller) ResumeTracing(opts *bind.CallOpts) error {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "resumeTracing")

	if err != nil {
		return err
	}

	return err

}

// ResumeTracing is a free data retrieval call binding the contract method 0x72a09ccb.
//
// Solidity: function resumeTracing() view returns()
func (_Vm *VmSession) ResumeTracing() error {
	return _Vm.Contract.ResumeTracing(&_Vm.CallOpts)
}

// ResumeTracing is a free data retrieval call binding the contract method 0x72a09ccb.
//
// Solidity: function resumeTracing() view returns()
func (_Vm *VmCallerSession) ResumeTracing() error {
	return _Vm.Contract.ResumeTracing(&_Vm.CallOpts)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmCaller) RpcUrl(opts *bind.CallOpts, rpcAlias string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrl", rpcAlias)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmSession) RpcUrl(rpcAlias string) (string, error) {
	return _Vm.Contract.RpcUrl(&_Vm.CallOpts, rpcAlias)
}

// RpcUrl is a free data retrieval call binding the contract method 0x975a6ce9.
//
// Solidity: function rpcUrl(string rpcAlias) view returns(string json)
func (_Vm *VmCallerSession) RpcUrl(rpcAlias string) (string, error) {
	return _Vm.Contract.RpcUrl(&_Vm.CallOpts, rpcAlias)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmCaller) RpcUrlStructs(opts *bind.CallOpts) ([]VmSafeRpc, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrlStructs")

	if err != nil {
		return *new([]VmSafeRpc), err
	}

	out0 := *abi.ConvertType(out[0], new([]VmSafeRpc)).(*[]VmSafeRpc)

	return out0, err

}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _Vm.Contract.RpcUrlStructs(&_Vm.CallOpts)
}

// RpcUrlStructs is a free data retrieval call binding the contract method 0x9d2ad72a.
//
// Solidity: function rpcUrlStructs() view returns((string,string)[] urls)
func (_Vm *VmCallerSession) RpcUrlStructs() ([]VmSafeRpc, error) {
	return _Vm.Contract.RpcUrlStructs(&_Vm.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmCaller) RpcUrls(opts *bind.CallOpts) ([][2]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "rpcUrls")

	if err != nil {
		return *new([][2]string), err
	}

	out0 := *abi.ConvertType(out[0], new([][2]string)).(*[][2]string)

	return out0, err

}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmSession) RpcUrls() ([][2]string, error) {
	return _Vm.Contract.RpcUrls(&_Vm.CallOpts)
}

// RpcUrls is a free data retrieval call binding the contract method 0xa85a8418.
//
// Solidity: function rpcUrls() view returns(string[2][] urls)
func (_Vm *VmCallerSession) RpcUrls() ([][2]string, error) {
	return _Vm.Contract.RpcUrls(&_Vm.CallOpts)
}

// SerializeJsonType is a free data retrieval call binding the contract method 0x6d4f96a6.
//
// Solidity: function serializeJsonType(string typeDescription, bytes value) pure returns(string json)
func (_Vm *VmCaller) SerializeJsonType(opts *bind.CallOpts, typeDescription string, value []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "serializeJsonType", typeDescription, value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SerializeJsonType is a free data retrieval call binding the contract method 0x6d4f96a6.
//
// Solidity: function serializeJsonType(string typeDescription, bytes value) pure returns(string json)
func (_Vm *VmSession) SerializeJsonType(typeDescription string, value []byte) (string, error) {
	return _Vm.Contract.SerializeJsonType(&_Vm.CallOpts, typeDescription, value)
}

// SerializeJsonType is a free data retrieval call binding the contract method 0x6d4f96a6.
//
// Solidity: function serializeJsonType(string typeDescription, bytes value) pure returns(string json)
func (_Vm *VmCallerSession) SerializeJsonType(typeDescription string, value []byte) (string, error) {
	return _Vm.Contract.SerializeJsonType(&_Vm.CallOpts, typeDescription, value)
}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCaller) Sign(opts *bind.CallOpts, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "sign", digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign(digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign(&_Vm.CallOpts, digest)
}

// Sign is a free data retrieval call binding the contract method 0x799cd333.
//
// Solidity: function sign(bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) Sign(digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign(&_Vm.CallOpts, digest)
}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCaller) Sign0(opts *bind.CallOpts, signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "sign0", signer, digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign0(signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign0(&_Vm.CallOpts, signer, digest)
}

// Sign0 is a free data retrieval call binding the contract method 0x8c1aa205.
//
// Solidity: function sign(address signer, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) Sign0(signer common.Address, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign0(&_Vm.CallOpts, signer, digest)
}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCaller) Sign2(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "sign2", privateKey, digest)

	outstruct := new(struct {
		V uint8
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.V = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.R = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign2(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign2(&_Vm.CallOpts, privateKey, digest)
}

// Sign2 is a free data retrieval call binding the contract method 0xe341eaa4.
//
// Solidity: function sign(uint256 privateKey, bytes32 digest) pure returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) Sign2(privateKey *big.Int, digest [32]byte) (struct {
	V uint8
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.Sign2(&_Vm.CallOpts, privateKey, digest)
}

// SignCompact0 is a free data retrieval call binding the contract method 0x8e2f97bf.
//
// Solidity: function signCompact(address signer, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmCaller) SignCompact0(opts *bind.CallOpts, signer common.Address, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "signCompact0", signer, digest)

	outstruct := new(struct {
		R  [32]byte
		Vs [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Vs = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignCompact0 is a free data retrieval call binding the contract method 0x8e2f97bf.
//
// Solidity: function signCompact(address signer, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmSession) SignCompact0(signer common.Address, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _Vm.Contract.SignCompact0(&_Vm.CallOpts, signer, digest)
}

// SignCompact0 is a free data retrieval call binding the contract method 0x8e2f97bf.
//
// Solidity: function signCompact(address signer, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmCallerSession) SignCompact0(signer common.Address, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _Vm.Contract.SignCompact0(&_Vm.CallOpts, signer, digest)
}

// SignCompact1 is a free data retrieval call binding the contract method 0xa282dc4b.
//
// Solidity: function signCompact(bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmCaller) SignCompact1(opts *bind.CallOpts, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "signCompact1", digest)

	outstruct := new(struct {
		R  [32]byte
		Vs [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Vs = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignCompact1 is a free data retrieval call binding the contract method 0xa282dc4b.
//
// Solidity: function signCompact(bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmSession) SignCompact1(digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _Vm.Contract.SignCompact1(&_Vm.CallOpts, digest)
}

// SignCompact1 is a free data retrieval call binding the contract method 0xa282dc4b.
//
// Solidity: function signCompact(bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmCallerSession) SignCompact1(digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _Vm.Contract.SignCompact1(&_Vm.CallOpts, digest)
}

// SignCompact2 is a free data retrieval call binding the contract method 0xcc2a781f.
//
// Solidity: function signCompact(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmCaller) SignCompact2(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "signCompact2", privateKey, digest)

	outstruct := new(struct {
		R  [32]byte
		Vs [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Vs = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignCompact2 is a free data retrieval call binding the contract method 0xcc2a781f.
//
// Solidity: function signCompact(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmSession) SignCompact2(privateKey *big.Int, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _Vm.Contract.SignCompact2(&_Vm.CallOpts, privateKey, digest)
}

// SignCompact2 is a free data retrieval call binding the contract method 0xcc2a781f.
//
// Solidity: function signCompact(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 vs)
func (_Vm *VmCallerSession) SignCompact2(privateKey *big.Int, digest [32]byte) (struct {
	R  [32]byte
	Vs [32]byte
}, error) {
	return _Vm.Contract.SignCompact2(&_Vm.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmCaller) SignP256(opts *bind.CallOpts, privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "signP256", privateKey, digest)

	outstruct := new(struct {
		R [32]byte
		S [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.R = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.S = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.SignP256(&_Vm.CallOpts, privateKey, digest)
}

// SignP256 is a free data retrieval call binding the contract method 0x83211b40.
//
// Solidity: function signP256(uint256 privateKey, bytes32 digest) pure returns(bytes32 r, bytes32 s)
func (_Vm *VmCallerSession) SignP256(privateKey *big.Int, digest [32]byte) (struct {
	R [32]byte
	S [32]byte
}, error) {
	return _Vm.Contract.SignP256(&_Vm.CallOpts, privateKey, digest)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_Vm *VmCaller) Split(opts *bind.CallOpts, input string, delimiter string) ([]string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "split", input, delimiter)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_Vm *VmSession) Split(input string, delimiter string) ([]string, error) {
	return _Vm.Contract.Split(&_Vm.CallOpts, input, delimiter)
}

// Split is a free data retrieval call binding the contract method 0x8bb75533.
//
// Solidity: function split(string input, string delimiter) pure returns(string[] outputs)
func (_Vm *VmCallerSession) Split(input string, delimiter string) ([]string, error) {
	return _Vm.Contract.Split(&_Vm.CallOpts, input, delimiter)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmCaller) ToBase64(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmSession) ToBase64(data string) (string, error) {
	return _Vm.Contract.ToBase64(&_Vm.CallOpts, data)
}

// ToBase64 is a free data retrieval call binding the contract method 0x3f8be2c8.
//
// Solidity: function toBase64(string data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64(data string) (string, error) {
	return _Vm.Contract.ToBase64(&_Vm.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmCaller) ToBase640(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase640", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmSession) ToBase640(data []byte) (string, error) {
	return _Vm.Contract.ToBase640(&_Vm.CallOpts, data)
}

// ToBase640 is a free data retrieval call binding the contract method 0xa5cbfe65.
//
// Solidity: function toBase64(bytes data) pure returns(string)
func (_Vm *VmCallerSession) ToBase640(data []byte) (string, error) {
	return _Vm.Contract.ToBase640(&_Vm.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmCaller) ToBase64URL(opts *bind.CallOpts, data string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64URL", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmSession) ToBase64URL(data string) (string, error) {
	return _Vm.Contract.ToBase64URL(&_Vm.CallOpts, data)
}

// ToBase64URL is a free data retrieval call binding the contract method 0xae3165b3.
//
// Solidity: function toBase64URL(string data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64URL(data string) (string, error) {
	return _Vm.Contract.ToBase64URL(&_Vm.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmCaller) ToBase64URL0(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toBase64URL0", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmSession) ToBase64URL0(data []byte) (string, error) {
	return _Vm.Contract.ToBase64URL0(&_Vm.CallOpts, data)
}

// ToBase64URL0 is a free data retrieval call binding the contract method 0xc8bd0e4a.
//
// Solidity: function toBase64URL(bytes data) pure returns(string)
func (_Vm *VmCallerSession) ToBase64URL0(data []byte) (string, error) {
	return _Vm.Contract.ToBase64URL0(&_Vm.CallOpts, data)
}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_Vm *VmCaller) ToLowercase(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toLowercase", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_Vm *VmSession) ToLowercase(input string) (string, error) {
	return _Vm.Contract.ToLowercase(&_Vm.CallOpts, input)
}

// ToLowercase is a free data retrieval call binding the contract method 0x50bb0884.
//
// Solidity: function toLowercase(string input) pure returns(string output)
func (_Vm *VmCallerSession) ToLowercase(input string) (string, error) {
	return _Vm.Contract.ToLowercase(&_Vm.CallOpts, input)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString(opts *bind.CallOpts, value common.Address) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString(value common.Address) (string, error) {
	return _Vm.Contract.ToString(&_Vm.CallOpts, value)
}

// ToString is a free data retrieval call binding the contract method 0x56ca623e.
//
// Solidity: function toString(address value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString(value common.Address) (string, error) {
	return _Vm.Contract.ToString(&_Vm.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString0(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString0", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString0(value *big.Int) (string, error) {
	return _Vm.Contract.ToString0(&_Vm.CallOpts, value)
}

// ToString0 is a free data retrieval call binding the contract method 0x6900a3ae.
//
// Solidity: function toString(uint256 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString0(value *big.Int) (string, error) {
	return _Vm.Contract.ToString0(&_Vm.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString1(opts *bind.CallOpts, value []byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString1", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString1(value []byte) (string, error) {
	return _Vm.Contract.ToString1(&_Vm.CallOpts, value)
}

// ToString1 is a free data retrieval call binding the contract method 0x71aad10d.
//
// Solidity: function toString(bytes value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString1(value []byte) (string, error) {
	return _Vm.Contract.ToString1(&_Vm.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString2(opts *bind.CallOpts, value bool) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString2", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString2(value bool) (string, error) {
	return _Vm.Contract.ToString2(&_Vm.CallOpts, value)
}

// ToString2 is a free data retrieval call binding the contract method 0x71dce7da.
//
// Solidity: function toString(bool value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString2(value bool) (string, error) {
	return _Vm.Contract.ToString2(&_Vm.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString3(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString3", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString3(value *big.Int) (string, error) {
	return _Vm.Contract.ToString3(&_Vm.CallOpts, value)
}

// ToString3 is a free data retrieval call binding the contract method 0xa322c40e.
//
// Solidity: function toString(int256 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString3(value *big.Int) (string, error) {
	return _Vm.Contract.ToString3(&_Vm.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmCaller) ToString4(opts *bind.CallOpts, value [32]byte) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toString4", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmSession) ToString4(value [32]byte) (string, error) {
	return _Vm.Contract.ToString4(&_Vm.CallOpts, value)
}

// ToString4 is a free data retrieval call binding the contract method 0xb11a19e8.
//
// Solidity: function toString(bytes32 value) pure returns(string stringifiedValue)
func (_Vm *VmCallerSession) ToString4(value [32]byte) (string, error) {
	return _Vm.Contract.ToString4(&_Vm.CallOpts, value)
}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_Vm *VmCaller) ToUppercase(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "toUppercase", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_Vm *VmSession) ToUppercase(input string) (string, error) {
	return _Vm.Contract.ToUppercase(&_Vm.CallOpts, input)
}

// ToUppercase is a free data retrieval call binding the contract method 0x074ae3d7.
//
// Solidity: function toUppercase(string input) pure returns(string output)
func (_Vm *VmCallerSession) ToUppercase(input string) (string, error) {
	return _Vm.Contract.ToUppercase(&_Vm.CallOpts, input)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_Vm *VmCaller) Trim(opts *bind.CallOpts, input string) (string, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "trim", input)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_Vm *VmSession) Trim(input string) (string, error) {
	return _Vm.Contract.Trim(&_Vm.CallOpts, input)
}

// Trim is a free data retrieval call binding the contract method 0xb2dad155.
//
// Solidity: function trim(string input) pure returns(string output)
func (_Vm *VmCallerSession) Trim(input string) (string, error) {
	return _Vm.Contract.Trim(&_Vm.CallOpts, input)
}

// UnixTime is a free data retrieval call binding the contract method 0x625387dc.
//
// Solidity: function unixTime() view returns(uint256 milliseconds)
func (_Vm *VmCaller) UnixTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vm.contract.Call(opts, &out, "unixTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnixTime is a free data retrieval call binding the contract method 0x625387dc.
//
// Solidity: function unixTime() view returns(uint256 milliseconds)
func (_Vm *VmSession) UnixTime() (*big.Int, error) {
	return _Vm.Contract.UnixTime(&_Vm.CallOpts)
}

// UnixTime is a free data retrieval call binding the contract method 0x625387dc.
//
// Solidity: function unixTime() view returns(uint256 milliseconds)
func (_Vm *VmCallerSession) UnixTime() (*big.Int, error) {
	return _Vm.Contract.UnixTime(&_Vm.CallOpts)
}

// AccessList is a paid mutator transaction binding the contract method 0x743e4cb7.
//
// Solidity: function accessList((address,bytes32[])[] access) returns()
func (_Vm *VmTransactor) AccessList(opts *bind.TransactOpts, access []VmSafeAccessListItem) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "accessList", access)
}

// AccessList is a paid mutator transaction binding the contract method 0x743e4cb7.
//
// Solidity: function accessList((address,bytes32[])[] access) returns()
func (_Vm *VmSession) AccessList(access []VmSafeAccessListItem) (*types.Transaction, error) {
	return _Vm.Contract.AccessList(&_Vm.TransactOpts, access)
}

// AccessList is a paid mutator transaction binding the contract method 0x743e4cb7.
//
// Solidity: function accessList((address,bytes32[])[] access) returns()
func (_Vm *VmTransactorSession) AccessList(access []VmSafeAccessListItem) (*types.Transaction, error) {
	return _Vm.Contract.AccessList(&_Vm.TransactOpts, access)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmTransactor) Accesses(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "accesses", target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Accesses(&_Vm.TransactOpts, target)
}

// Accesses is a paid mutator transaction binding the contract method 0x65bc9481.
//
// Solidity: function accesses(address target) returns(bytes32[] readSlots, bytes32[] writeSlots)
func (_Vm *VmTransactorSession) Accesses(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Accesses(&_Vm.TransactOpts, target)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmTransactor) AllowCheatcodes(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "allowCheatcodes", account)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmSession) AllowCheatcodes(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.AllowCheatcodes(&_Vm.TransactOpts, account)
}

// AllowCheatcodes is a paid mutator transaction binding the contract method 0xea060291.
//
// Solidity: function allowCheatcodes(address account) returns()
func (_Vm *VmTransactorSession) AllowCheatcodes(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.AllowCheatcodes(&_Vm.TransactOpts, account)
}

// AttachBlob is a paid mutator transaction binding the contract method 0x10cb385c.
//
// Solidity: function attachBlob(bytes blob) returns()
func (_Vm *VmTransactor) AttachBlob(opts *bind.TransactOpts, blob []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "attachBlob", blob)
}

// AttachBlob is a paid mutator transaction binding the contract method 0x10cb385c.
//
// Solidity: function attachBlob(bytes blob) returns()
func (_Vm *VmSession) AttachBlob(blob []byte) (*types.Transaction, error) {
	return _Vm.Contract.AttachBlob(&_Vm.TransactOpts, blob)
}

// AttachBlob is a paid mutator transaction binding the contract method 0x10cb385c.
//
// Solidity: function attachBlob(bytes blob) returns()
func (_Vm *VmTransactorSession) AttachBlob(blob []byte) (*types.Transaction, error) {
	return _Vm.Contract.AttachBlob(&_Vm.TransactOpts, blob)
}

// AttachDelegation is a paid mutator transaction binding the contract method 0x14ae3519.
//
// Solidity: function attachDelegation((uint8,bytes32,bytes32,uint64,address) signedDelegation) returns()
func (_Vm *VmTransactor) AttachDelegation(opts *bind.TransactOpts, signedDelegation VmSafeSignedDelegation) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "attachDelegation", signedDelegation)
}

// AttachDelegation is a paid mutator transaction binding the contract method 0x14ae3519.
//
// Solidity: function attachDelegation((uint8,bytes32,bytes32,uint64,address) signedDelegation) returns()
func (_Vm *VmSession) AttachDelegation(signedDelegation VmSafeSignedDelegation) (*types.Transaction, error) {
	return _Vm.Contract.AttachDelegation(&_Vm.TransactOpts, signedDelegation)
}

// AttachDelegation is a paid mutator transaction binding the contract method 0x14ae3519.
//
// Solidity: function attachDelegation((uint8,bytes32,bytes32,uint64,address) signedDelegation) returns()
func (_Vm *VmTransactorSession) AttachDelegation(signedDelegation VmSafeSignedDelegation) (*types.Transaction, error) {
	return _Vm.Contract.AttachDelegation(&_Vm.TransactOpts, signedDelegation)
}

// BlobBaseFee is a paid mutator transaction binding the contract method 0x6d315d7e.
//
// Solidity: function blobBaseFee(uint256 newBlobBaseFee) returns()
func (_Vm *VmTransactor) BlobBaseFee(opts *bind.TransactOpts, newBlobBaseFee *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "blobBaseFee", newBlobBaseFee)
}

// BlobBaseFee is a paid mutator transaction binding the contract method 0x6d315d7e.
//
// Solidity: function blobBaseFee(uint256 newBlobBaseFee) returns()
func (_Vm *VmSession) BlobBaseFee(newBlobBaseFee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.BlobBaseFee(&_Vm.TransactOpts, newBlobBaseFee)
}

// BlobBaseFee is a paid mutator transaction binding the contract method 0x6d315d7e.
//
// Solidity: function blobBaseFee(uint256 newBlobBaseFee) returns()
func (_Vm *VmTransactorSession) BlobBaseFee(newBlobBaseFee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.BlobBaseFee(&_Vm.TransactOpts, newBlobBaseFee)
}

// Blobhashes is a paid mutator transaction binding the contract method 0x129de7eb.
//
// Solidity: function blobhashes(bytes32[] hashes) returns()
func (_Vm *VmTransactor) Blobhashes(opts *bind.TransactOpts, hashes [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "blobhashes", hashes)
}

// Blobhashes is a paid mutator transaction binding the contract method 0x129de7eb.
//
// Solidity: function blobhashes(bytes32[] hashes) returns()
func (_Vm *VmSession) Blobhashes(hashes [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Blobhashes(&_Vm.TransactOpts, hashes)
}

// Blobhashes is a paid mutator transaction binding the contract method 0x129de7eb.
//
// Solidity: function blobhashes(bytes32[] hashes) returns()
func (_Vm *VmTransactorSession) Blobhashes(hashes [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Blobhashes(&_Vm.TransactOpts, hashes)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmTransactor) Broadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast")
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmSession) Broadcast() (*types.Transaction, error) {
	return _Vm.Contract.Broadcast(&_Vm.TransactOpts)
}

// Broadcast is a paid mutator transaction binding the contract method 0xafc98040.
//
// Solidity: function broadcast() returns()
func (_Vm *VmTransactorSession) Broadcast() (*types.Transaction, error) {
	return _Vm.Contract.Broadcast(&_Vm.TransactOpts)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmTransactor) Broadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast0", signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast0(&_Vm.TransactOpts, signer)
}

// Broadcast0 is a paid mutator transaction binding the contract method 0xe6962cdb.
//
// Solidity: function broadcast(address signer) returns()
func (_Vm *VmTransactorSession) Broadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast0(&_Vm.TransactOpts, signer)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmTransactor) Broadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcast1", privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast1(&_Vm.TransactOpts, privateKey)
}

// Broadcast1 is a paid mutator transaction binding the contract method 0xf67a965b.
//
// Solidity: function broadcast(uint256 privateKey) returns()
func (_Vm *VmTransactorSession) Broadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Broadcast1(&_Vm.TransactOpts, privateKey)
}

// BroadcastRawTransaction is a paid mutator transaction binding the contract method 0x8c0c72e0.
//
// Solidity: function broadcastRawTransaction(bytes data) returns()
func (_Vm *VmTransactor) BroadcastRawTransaction(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "broadcastRawTransaction", data)
}

// BroadcastRawTransaction is a paid mutator transaction binding the contract method 0x8c0c72e0.
//
// Solidity: function broadcastRawTransaction(bytes data) returns()
func (_Vm *VmSession) BroadcastRawTransaction(data []byte) (*types.Transaction, error) {
	return _Vm.Contract.BroadcastRawTransaction(&_Vm.TransactOpts, data)
}

// BroadcastRawTransaction is a paid mutator transaction binding the contract method 0x8c0c72e0.
//
// Solidity: function broadcastRawTransaction(bytes data) returns()
func (_Vm *VmTransactorSession) BroadcastRawTransaction(data []byte) (*types.Transaction, error) {
	return _Vm.Contract.BroadcastRawTransaction(&_Vm.TransactOpts, data)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmTransactor) ChainId(opts *bind.TransactOpts, newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "chainId", newChainId)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmSession) ChainId(newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.ChainId(&_Vm.TransactOpts, newChainId)
}

// ChainId is a paid mutator transaction binding the contract method 0x4049ddd2.
//
// Solidity: function chainId(uint256 newChainId) returns()
func (_Vm *VmTransactorSession) ChainId(newChainId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.ChainId(&_Vm.TransactOpts, newChainId)
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmTransactor) ClearMockedCalls(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "clearMockedCalls")
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmSession) ClearMockedCalls() (*types.Transaction, error) {
	return _Vm.Contract.ClearMockedCalls(&_Vm.TransactOpts)
}

// ClearMockedCalls is a paid mutator transaction binding the contract method 0x3fdf4e15.
//
// Solidity: function clearMockedCalls() returns()
func (_Vm *VmTransactorSession) ClearMockedCalls() (*types.Transaction, error) {
	return _Vm.Contract.ClearMockedCalls(&_Vm.TransactOpts)
}

// CloneAccount is a paid mutator transaction binding the contract method 0x533d61c9.
//
// Solidity: function cloneAccount(address source, address target) returns()
func (_Vm *VmTransactor) CloneAccount(opts *bind.TransactOpts, source common.Address, target common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "cloneAccount", source, target)
}

// CloneAccount is a paid mutator transaction binding the contract method 0x533d61c9.
//
// Solidity: function cloneAccount(address source, address target) returns()
func (_Vm *VmSession) CloneAccount(source common.Address, target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.CloneAccount(&_Vm.TransactOpts, source, target)
}

// CloneAccount is a paid mutator transaction binding the contract method 0x533d61c9.
//
// Solidity: function cloneAccount(address source, address target) returns()
func (_Vm *VmTransactorSession) CloneAccount(source common.Address, target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.CloneAccount(&_Vm.TransactOpts, source, target)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmTransactor) CloseFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "closeFile", path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmSession) CloseFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.CloseFile(&_Vm.TransactOpts, path)
}

// CloseFile is a paid mutator transaction binding the contract method 0x48c3241f.
//
// Solidity: function closeFile(string path) returns()
func (_Vm *VmTransactorSession) CloseFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.CloseFile(&_Vm.TransactOpts, path)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmTransactor) Coinbase(opts *bind.TransactOpts, newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "coinbase", newCoinbase)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmSession) Coinbase(newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Coinbase(&_Vm.TransactOpts, newCoinbase)
}

// Coinbase is a paid mutator transaction binding the contract method 0xff483c54.
//
// Solidity: function coinbase(address newCoinbase) returns()
func (_Vm *VmTransactorSession) Coinbase(newCoinbase common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Coinbase(&_Vm.TransactOpts, newCoinbase)
}

// Contains is a paid mutator transaction binding the contract method 0x3fb18aec.
//
// Solidity: function contains(string subject, string search) returns(bool result)
func (_Vm *VmTransactor) Contains(opts *bind.TransactOpts, subject string, search string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "contains", subject, search)
}

// Contains is a paid mutator transaction binding the contract method 0x3fb18aec.
//
// Solidity: function contains(string subject, string search) returns(bool result)
func (_Vm *VmSession) Contains(subject string, search string) (*types.Transaction, error) {
	return _Vm.Contract.Contains(&_Vm.TransactOpts, subject, search)
}

// Contains is a paid mutator transaction binding the contract method 0x3fb18aec.
//
// Solidity: function contains(string subject, string search) returns(bool result)
func (_Vm *VmTransactorSession) Contains(subject string, search string) (*types.Transaction, error) {
	return _Vm.Contract.Contains(&_Vm.TransactOpts, subject, search)
}

// Cool is a paid mutator transaction binding the contract method 0x40ff9f21.
//
// Solidity: function cool(address target) returns()
func (_Vm *VmTransactor) Cool(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "cool", target)
}

// Cool is a paid mutator transaction binding the contract method 0x40ff9f21.
//
// Solidity: function cool(address target) returns()
func (_Vm *VmSession) Cool(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Cool(&_Vm.TransactOpts, target)
}

// Cool is a paid mutator transaction binding the contract method 0x40ff9f21.
//
// Solidity: function cool(address target) returns()
func (_Vm *VmTransactorSession) Cool(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Cool(&_Vm.TransactOpts, target)
}

// CoolSlot is a paid mutator transaction binding the contract method 0x8c78e654.
//
// Solidity: function coolSlot(address target, bytes32 slot) returns()
func (_Vm *VmTransactor) CoolSlot(opts *bind.TransactOpts, target common.Address, slot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "coolSlot", target, slot)
}

// CoolSlot is a paid mutator transaction binding the contract method 0x8c78e654.
//
// Solidity: function coolSlot(address target, bytes32 slot) returns()
func (_Vm *VmSession) CoolSlot(target common.Address, slot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CoolSlot(&_Vm.TransactOpts, target, slot)
}

// CoolSlot is a paid mutator transaction binding the contract method 0x8c78e654.
//
// Solidity: function coolSlot(address target, bytes32 slot) returns()
func (_Vm *VmTransactorSession) CoolSlot(target common.Address, slot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CoolSlot(&_Vm.TransactOpts, target, slot)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmTransactor) CopyFile(opts *bind.TransactOpts, from string, to string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "copyFile", from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _Vm.Contract.CopyFile(&_Vm.TransactOpts, from, to)
}

// CopyFile is a paid mutator transaction binding the contract method 0xa54a87d8.
//
// Solidity: function copyFile(string from, string to) returns(uint64 copied)
func (_Vm *VmTransactorSession) CopyFile(from string, to string) (*types.Transaction, error) {
	return _Vm.Contract.CopyFile(&_Vm.TransactOpts, from, to)
}

// CopyStorage is a paid mutator transaction binding the contract method 0x203dac0d.
//
// Solidity: function copyStorage(address from, address to) returns()
func (_Vm *VmTransactor) CopyStorage(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "copyStorage", from, to)
}

// CopyStorage is a paid mutator transaction binding the contract method 0x203dac0d.
//
// Solidity: function copyStorage(address from, address to) returns()
func (_Vm *VmSession) CopyStorage(from common.Address, to common.Address) (*types.Transaction, error) {
	return _Vm.Contract.CopyStorage(&_Vm.TransactOpts, from, to)
}

// CopyStorage is a paid mutator transaction binding the contract method 0x203dac0d.
//
// Solidity: function copyStorage(address from, address to) returns()
func (_Vm *VmTransactorSession) CopyStorage(from common.Address, to common.Address) (*types.Transaction, error) {
	return _Vm.Contract.CopyStorage(&_Vm.TransactOpts, from, to)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmTransactor) CreateDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createDir", path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.CreateDir(&_Vm.TransactOpts, path, recursive)
}

// CreateDir is a paid mutator transaction binding the contract method 0x168b64d3.
//
// Solidity: function createDir(string path, bool recursive) returns()
func (_Vm *VmTransactorSession) CreateDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.CreateDir(&_Vm.TransactOpts, path, recursive)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork(opts *bind.TransactOpts, urlOrAlias string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork", urlOrAlias)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork(&_Vm.TransactOpts, urlOrAlias)
}

// CreateFork is a paid mutator transaction binding the contract method 0x31ba3498.
//
// Solidity: function createFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork(&_Vm.TransactOpts, urlOrAlias)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork0(opts *bind.TransactOpts, urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork0", urlOrAlias, blockNumber)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork0(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork0(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateFork0 is a paid mutator transaction binding the contract method 0x6ba3ba2b.
//
// Solidity: function createFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork0(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork0(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateFork1(opts *bind.TransactOpts, urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createFork1", urlOrAlias, txHash)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmSession) CreateFork1(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork1(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateFork1 is a paid mutator transaction binding the contract method 0x7ca29682.
//
// Solidity: function createFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateFork1(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateFork1(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork(opts *bind.TransactOpts, urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork", urlOrAlias, blockNumber)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateSelectFork is a paid mutator transaction binding the contract method 0x71ee464d.
//
// Solidity: function createSelectFork(string urlOrAlias, uint256 blockNumber) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork(urlOrAlias string, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork(&_Vm.TransactOpts, urlOrAlias, blockNumber)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork0(opts *bind.TransactOpts, urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork0", urlOrAlias, txHash)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork0(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork0(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork0 is a paid mutator transaction binding the contract method 0x84d52b7a.
//
// Solidity: function createSelectFork(string urlOrAlias, bytes32 txHash) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork0(urlOrAlias string, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork0(&_Vm.TransactOpts, urlOrAlias, txHash)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactor) CreateSelectFork1(opts *bind.TransactOpts, urlOrAlias string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createSelectFork1", urlOrAlias)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmSession) CreateSelectFork1(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork1(&_Vm.TransactOpts, urlOrAlias)
}

// CreateSelectFork1 is a paid mutator transaction binding the contract method 0x98680034.
//
// Solidity: function createSelectFork(string urlOrAlias) returns(uint256 forkId)
func (_Vm *VmTransactorSession) CreateSelectFork1(urlOrAlias string) (*types.Transaction, error) {
	return _Vm.Contract.CreateSelectFork1(&_Vm.TransactOpts, urlOrAlias)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet(opts *bind.TransactOpts, walletLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet", walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet(&_Vm.TransactOpts, walletLabel)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x7404f1d2.
//
// Solidity: function createWallet(string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet(walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet(&_Vm.TransactOpts, walletLabel)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet0(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet0", privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet0(&_Vm.TransactOpts, privateKey)
}

// CreateWallet0 is a paid mutator transaction binding the contract method 0x7a675bb6.
//
// Solidity: function createWallet(uint256 privateKey) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet0(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet0(&_Vm.TransactOpts, privateKey)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactor) CreateWallet1(opts *bind.TransactOpts, privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "createWallet1", privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet1(&_Vm.TransactOpts, privateKey, walletLabel)
}

// CreateWallet1 is a paid mutator transaction binding the contract method 0xed7c5462.
//
// Solidity: function createWallet(uint256 privateKey, string walletLabel) returns((address,uint256,uint256,uint256) wallet)
func (_Vm *VmTransactorSession) CreateWallet1(privateKey *big.Int, walletLabel string) (*types.Transaction, error) {
	return _Vm.Contract.CreateWallet1(&_Vm.TransactOpts, privateKey, walletLabel)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmTransactor) Deal(opts *bind.TransactOpts, account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deal", account, newBalance)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmSession) Deal(account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Deal(&_Vm.TransactOpts, account, newBalance)
}

// Deal is a paid mutator transaction binding the contract method 0xc88a5e6d.
//
// Solidity: function deal(address account, uint256 newBalance) returns()
func (_Vm *VmTransactorSession) Deal(account common.Address, newBalance *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Deal(&_Vm.TransactOpts, account, newBalance)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) DeleteSnapshot(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteSnapshot", snapshotId)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) DeleteSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteSnapshot is a paid mutator transaction binding the contract method 0xa6368557.
//
// Solidity: function deleteSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) DeleteSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmTransactor) DeleteSnapshots(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteSnapshots")
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmSession) DeleteSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshots(&_Vm.TransactOpts)
}

// DeleteSnapshots is a paid mutator transaction binding the contract method 0x421ae469.
//
// Solidity: function deleteSnapshots() returns()
func (_Vm *VmTransactorSession) DeleteSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteSnapshots(&_Vm.TransactOpts)
}

// DeleteStateSnapshot is a paid mutator transaction binding the contract method 0x08d6b37a.
//
// Solidity: function deleteStateSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) DeleteStateSnapshot(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteStateSnapshot", snapshotId)
}

// DeleteStateSnapshot is a paid mutator transaction binding the contract method 0x08d6b37a.
//
// Solidity: function deleteStateSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) DeleteStateSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteStateSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteStateSnapshot is a paid mutator transaction binding the contract method 0x08d6b37a.
//
// Solidity: function deleteStateSnapshot(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) DeleteStateSnapshot(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeleteStateSnapshot(&_Vm.TransactOpts, snapshotId)
}

// DeleteStateSnapshots is a paid mutator transaction binding the contract method 0xe0933c74.
//
// Solidity: function deleteStateSnapshots() returns()
func (_Vm *VmTransactor) DeleteStateSnapshots(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deleteStateSnapshots")
}

// DeleteStateSnapshots is a paid mutator transaction binding the contract method 0xe0933c74.
//
// Solidity: function deleteStateSnapshots() returns()
func (_Vm *VmSession) DeleteStateSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteStateSnapshots(&_Vm.TransactOpts)
}

// DeleteStateSnapshots is a paid mutator transaction binding the contract method 0xe0933c74.
//
// Solidity: function deleteStateSnapshots() returns()
func (_Vm *VmTransactorSession) DeleteStateSnapshots() (*types.Transaction, error) {
	return _Vm.Contract.DeleteStateSnapshots(&_Vm.TransactOpts)
}

// DeployCode is a paid mutator transaction binding the contract method 0x002cb687.
//
// Solidity: function deployCode(string artifactPath, uint256 value, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode(opts *bind.TransactOpts, artifactPath string, value *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode", artifactPath, value, salt)
}

// DeployCode is a paid mutator transaction binding the contract method 0x002cb687.
//
// Solidity: function deployCode(string artifactPath, uint256 value, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode(artifactPath string, value *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode(&_Vm.TransactOpts, artifactPath, value, salt)
}

// DeployCode is a paid mutator transaction binding the contract method 0x002cb687.
//
// Solidity: function deployCode(string artifactPath, uint256 value, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode(artifactPath string, value *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode(&_Vm.TransactOpts, artifactPath, value, salt)
}

// DeployCode0 is a paid mutator transaction binding the contract method 0x016155bf.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode0(opts *bind.TransactOpts, artifactPath string, constructorArgs []byte, salt [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode0", artifactPath, constructorArgs, salt)
}

// DeployCode0 is a paid mutator transaction binding the contract method 0x016155bf.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode0(artifactPath string, constructorArgs []byte, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode0(&_Vm.TransactOpts, artifactPath, constructorArgs, salt)
}

// DeployCode0 is a paid mutator transaction binding the contract method 0x016155bf.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode0(artifactPath string, constructorArgs []byte, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode0(&_Vm.TransactOpts, artifactPath, constructorArgs, salt)
}

// DeployCode1 is a paid mutator transaction binding the contract method 0x0af6a701.
//
// Solidity: function deployCode(string artifactPath, uint256 value) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode1(opts *bind.TransactOpts, artifactPath string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode1", artifactPath, value)
}

// DeployCode1 is a paid mutator transaction binding the contract method 0x0af6a701.
//
// Solidity: function deployCode(string artifactPath, uint256 value) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode1(artifactPath string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode1(&_Vm.TransactOpts, artifactPath, value)
}

// DeployCode1 is a paid mutator transaction binding the contract method 0x0af6a701.
//
// Solidity: function deployCode(string artifactPath, uint256 value) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode1(artifactPath string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode1(&_Vm.TransactOpts, artifactPath, value)
}

// DeployCode2 is a paid mutator transaction binding the contract method 0x17ab1d79.
//
// Solidity: function deployCode(string artifactPath, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode2(opts *bind.TransactOpts, artifactPath string, salt [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode2", artifactPath, salt)
}

// DeployCode2 is a paid mutator transaction binding the contract method 0x17ab1d79.
//
// Solidity: function deployCode(string artifactPath, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode2(artifactPath string, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode2(&_Vm.TransactOpts, artifactPath, salt)
}

// DeployCode2 is a paid mutator transaction binding the contract method 0x17ab1d79.
//
// Solidity: function deployCode(string artifactPath, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode2(artifactPath string, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode2(&_Vm.TransactOpts, artifactPath, salt)
}

// DeployCode3 is a paid mutator transaction binding the contract method 0x29ce9dde.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode3(opts *bind.TransactOpts, artifactPath string, constructorArgs []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode3", artifactPath, constructorArgs)
}

// DeployCode3 is a paid mutator transaction binding the contract method 0x29ce9dde.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode3(artifactPath string, constructorArgs []byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode3(&_Vm.TransactOpts, artifactPath, constructorArgs)
}

// DeployCode3 is a paid mutator transaction binding the contract method 0x29ce9dde.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode3(artifactPath string, constructorArgs []byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode3(&_Vm.TransactOpts, artifactPath, constructorArgs)
}

// DeployCode4 is a paid mutator transaction binding the contract method 0x3aa773ea.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, uint256 value, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode4(opts *bind.TransactOpts, artifactPath string, constructorArgs []byte, value *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode4", artifactPath, constructorArgs, value, salt)
}

// DeployCode4 is a paid mutator transaction binding the contract method 0x3aa773ea.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, uint256 value, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode4(artifactPath string, constructorArgs []byte, value *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode4(&_Vm.TransactOpts, artifactPath, constructorArgs, value, salt)
}

// DeployCode4 is a paid mutator transaction binding the contract method 0x3aa773ea.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, uint256 value, bytes32 salt) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode4(artifactPath string, constructorArgs []byte, value *big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode4(&_Vm.TransactOpts, artifactPath, constructorArgs, value, salt)
}

// DeployCode5 is a paid mutator transaction binding the contract method 0x9a8325a0.
//
// Solidity: function deployCode(string artifactPath) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode5(opts *bind.TransactOpts, artifactPath string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode5", artifactPath)
}

// DeployCode5 is a paid mutator transaction binding the contract method 0x9a8325a0.
//
// Solidity: function deployCode(string artifactPath) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode5(artifactPath string) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode5(&_Vm.TransactOpts, artifactPath)
}

// DeployCode5 is a paid mutator transaction binding the contract method 0x9a8325a0.
//
// Solidity: function deployCode(string artifactPath) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode5(artifactPath string) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode5(&_Vm.TransactOpts, artifactPath)
}

// DeployCode6 is a paid mutator transaction binding the contract method 0xff5d64e4.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, uint256 value) returns(address deployedAddress)
func (_Vm *VmTransactor) DeployCode6(opts *bind.TransactOpts, artifactPath string, constructorArgs []byte, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "deployCode6", artifactPath, constructorArgs, value)
}

// DeployCode6 is a paid mutator transaction binding the contract method 0xff5d64e4.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, uint256 value) returns(address deployedAddress)
func (_Vm *VmSession) DeployCode6(artifactPath string, constructorArgs []byte, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode6(&_Vm.TransactOpts, artifactPath, constructorArgs, value)
}

// DeployCode6 is a paid mutator transaction binding the contract method 0xff5d64e4.
//
// Solidity: function deployCode(string artifactPath, bytes constructorArgs, uint256 value) returns(address deployedAddress)
func (_Vm *VmTransactorSession) DeployCode6(artifactPath string, constructorArgs []byte, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.DeployCode6(&_Vm.TransactOpts, artifactPath, constructorArgs, value)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmTransactor) Difficulty(opts *bind.TransactOpts, newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "difficulty", newDifficulty)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmSession) Difficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Difficulty(&_Vm.TransactOpts, newDifficulty)
}

// Difficulty is a paid mutator transaction binding the contract method 0x46cc92d9.
//
// Solidity: function difficulty(uint256 newDifficulty) returns()
func (_Vm *VmTransactorSession) Difficulty(newDifficulty *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Difficulty(&_Vm.TransactOpts, newDifficulty)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmTransactor) DumpState(opts *bind.TransactOpts, pathToStateJson string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "dumpState", pathToStateJson)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmSession) DumpState(pathToStateJson string) (*types.Transaction, error) {
	return _Vm.Contract.DumpState(&_Vm.TransactOpts, pathToStateJson)
}

// DumpState is a paid mutator transaction binding the contract method 0x709ecd3f.
//
// Solidity: function dumpState(string pathToStateJson) returns()
func (_Vm *VmTransactorSession) DumpState(pathToStateJson string) (*types.Transaction, error) {
	return _Vm.Contract.DumpState(&_Vm.TransactOpts, pathToStateJson)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmTransactor) Etch(opts *bind.TransactOpts, target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "etch", target, newRuntimeBytecode)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmSession) Etch(target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.Contract.Etch(&_Vm.TransactOpts, target, newRuntimeBytecode)
}

// Etch is a paid mutator transaction binding the contract method 0xb4d6c782.
//
// Solidity: function etch(address target, bytes newRuntimeBytecode) returns()
func (_Vm *VmTransactorSession) Etch(target common.Address, newRuntimeBytecode []byte) (*types.Transaction, error) {
	return _Vm.Contract.Etch(&_Vm.TransactOpts, target, newRuntimeBytecode)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmTransactor) EthGetLogs(opts *bind.TransactOpts, fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "eth_getLogs", fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.EthGetLogs(&_Vm.TransactOpts, fromBlock, toBlock, target, topics)
}

// EthGetLogs is a paid mutator transaction binding the contract method 0x35e1349b.
//
// Solidity: function eth_getLogs(uint256 fromBlock, uint256 toBlock, address target, bytes32[] topics) returns((address,bytes32[],bytes,bytes32,uint64,bytes32,uint64,uint256,bool)[] logs)
func (_Vm *VmTransactorSession) EthGetLogs(fromBlock *big.Int, toBlock *big.Int, target common.Address, topics [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.EthGetLogs(&_Vm.TransactOpts, fromBlock, toBlock, target, topics)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall", callee, msgValue, gas, data)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmSession) ExpectCall(callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall(&_Vm.TransactOpts, callee, msgValue, gas, data)
}

// ExpectCall is a paid mutator transaction binding the contract method 0x23361207.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall(callee common.Address, msgValue *big.Int, gas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall(&_Vm.TransactOpts, callee, msgValue, gas, data)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall0", callee, msgValue, gas, data, count)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall0(callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall0(&_Vm.TransactOpts, callee, msgValue, gas, data, count)
}

// ExpectCall0 is a paid mutator transaction binding the contract method 0x65b7b7cc.
//
// Solidity: function expectCall(address callee, uint256 msgValue, uint64 gas, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall0(callee common.Address, msgValue *big.Int, gas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall0(&_Vm.TransactOpts, callee, msgValue, gas, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall1(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall1", callee, msgValue, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall1(callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall1(&_Vm.TransactOpts, callee, msgValue, data, count)
}

// ExpectCall1 is a paid mutator transaction binding the contract method 0xa2b1a1ae.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall1(callee common.Address, msgValue *big.Int, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall1(&_Vm.TransactOpts, callee, msgValue, data, count)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall2(opts *bind.TransactOpts, callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall2", callee, data)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmSession) ExpectCall2(callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall2(&_Vm.TransactOpts, callee, data)
}

// ExpectCall2 is a paid mutator transaction binding the contract method 0xbd6af434.
//
// Solidity: function expectCall(address callee, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall2(callee common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall2(&_Vm.TransactOpts, callee, data)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCall3(opts *bind.TransactOpts, callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall3", callee, data, count)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCall3(callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall3(&_Vm.TransactOpts, callee, data, count)
}

// ExpectCall3 is a paid mutator transaction binding the contract method 0xc1adbbff.
//
// Solidity: function expectCall(address callee, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCall3(callee common.Address, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall3(&_Vm.TransactOpts, callee, data, count)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmTransactor) ExpectCall4(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCall4", callee, msgValue, data)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmSession) ExpectCall4(callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall4(&_Vm.TransactOpts, callee, msgValue, data)
}

// ExpectCall4 is a paid mutator transaction binding the contract method 0xf30c7ba3.
//
// Solidity: function expectCall(address callee, uint256 msgValue, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCall4(callee common.Address, msgValue *big.Int, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCall4(&_Vm.TransactOpts, callee, msgValue, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmTransactor) ExpectCallMinGas(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCallMinGas", callee, msgValue, minGas, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmSession) ExpectCallMinGas(callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas(&_Vm.TransactOpts, callee, msgValue, minGas, data)
}

// ExpectCallMinGas is a paid mutator transaction binding the contract method 0x08e4e116.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data) returns()
func (_Vm *VmTransactorSession) ExpectCallMinGas(callee common.Address, msgValue *big.Int, minGas uint64, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas(&_Vm.TransactOpts, callee, msgValue, minGas, data)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmTransactor) ExpectCallMinGas0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCallMinGas0", callee, msgValue, minGas, data, count)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmSession) ExpectCallMinGas0(callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas0(&_Vm.TransactOpts, callee, msgValue, minGas, data, count)
}

// ExpectCallMinGas0 is a paid mutator transaction binding the contract method 0xe13a1834.
//
// Solidity: function expectCallMinGas(address callee, uint256 msgValue, uint64 minGas, bytes data, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectCallMinGas0(callee common.Address, msgValue *big.Int, minGas uint64, data []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCallMinGas0(&_Vm.TransactOpts, callee, msgValue, minGas, data, count)
}

// ExpectCreate is a paid mutator transaction binding the contract method 0x73cdce36.
//
// Solidity: function expectCreate(bytes bytecode, address deployer) returns()
func (_Vm *VmTransactor) ExpectCreate(opts *bind.TransactOpts, bytecode []byte, deployer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCreate", bytecode, deployer)
}

// ExpectCreate is a paid mutator transaction binding the contract method 0x73cdce36.
//
// Solidity: function expectCreate(bytes bytecode, address deployer) returns()
func (_Vm *VmSession) ExpectCreate(bytecode []byte, deployer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCreate(&_Vm.TransactOpts, bytecode, deployer)
}

// ExpectCreate is a paid mutator transaction binding the contract method 0x73cdce36.
//
// Solidity: function expectCreate(bytes bytecode, address deployer) returns()
func (_Vm *VmTransactorSession) ExpectCreate(bytecode []byte, deployer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCreate(&_Vm.TransactOpts, bytecode, deployer)
}

// ExpectCreate2 is a paid mutator transaction binding the contract method 0xea54a472.
//
// Solidity: function expectCreate2(bytes bytecode, address deployer) returns()
func (_Vm *VmTransactor) ExpectCreate2(opts *bind.TransactOpts, bytecode []byte, deployer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectCreate2", bytecode, deployer)
}

// ExpectCreate2 is a paid mutator transaction binding the contract method 0xea54a472.
//
// Solidity: function expectCreate2(bytes bytecode, address deployer) returns()
func (_Vm *VmSession) ExpectCreate2(bytecode []byte, deployer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCreate2(&_Vm.TransactOpts, bytecode, deployer)
}

// ExpectCreate2 is a paid mutator transaction binding the contract method 0xea54a472.
//
// Solidity: function expectCreate2(bytes bytecode, address deployer) returns()
func (_Vm *VmTransactorSession) ExpectCreate2(bytecode []byte, deployer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectCreate2(&_Vm.TransactOpts, bytecode, deployer)
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmTransactor) ExpectEmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit")
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmSession) ExpectEmit() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit(&_Vm.TransactOpts)
}

// ExpectEmit is a paid mutator transaction binding the contract method 0x440ed10d.
//
// Solidity: function expectEmit() returns()
func (_Vm *VmTransactorSession) ExpectEmit() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit(&_Vm.TransactOpts)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactor) ExpectEmit0(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit0", checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmSession) ExpectEmit0(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit0(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit0 is a paid mutator transaction binding the contract method 0x491cc7c2.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactorSession) ExpectEmit0(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit0(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x4c74a335.
//
// Solidity: function expectEmit(uint64 count) returns()
func (_Vm *VmTransactor) ExpectEmit1(opts *bind.TransactOpts, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit1", count)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x4c74a335.
//
// Solidity: function expectEmit(uint64 count) returns()
func (_Vm *VmSession) ExpectEmit1(count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit1(&_Vm.TransactOpts, count)
}

// ExpectEmit1 is a paid mutator transaction binding the contract method 0x4c74a335.
//
// Solidity: function expectEmit(uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectEmit1(count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit1(&_Vm.TransactOpts, count)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x5e1d1c33.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, uint64 count) returns()
func (_Vm *VmTransactor) ExpectEmit2(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit2", checkTopic1, checkTopic2, checkTopic3, checkData, count)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x5e1d1c33.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, uint64 count) returns()
func (_Vm *VmSession) ExpectEmit2(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit2(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, count)
}

// ExpectEmit2 is a paid mutator transaction binding the contract method 0x5e1d1c33.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectEmit2(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit2(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, count)
}

// ExpectEmit3 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactor) ExpectEmit3(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit3", checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit3 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmSession) ExpectEmit3(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit3(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit3 is a paid mutator transaction binding the contract method 0x81bad6f3.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmit3(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit3(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmit4 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmTransactor) ExpectEmit4(opts *bind.TransactOpts, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit4", emitter)
}

// ExpectEmit4 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmSession) ExpectEmit4(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit4(&_Vm.TransactOpts, emitter)
}

// ExpectEmit4 is a paid mutator transaction binding the contract method 0x86b9620d.
//
// Solidity: function expectEmit(address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmit4(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit4(&_Vm.TransactOpts, emitter)
}

// ExpectEmit5 is a paid mutator transaction binding the contract method 0xb43aece3.
//
// Solidity: function expectEmit(address emitter, uint64 count) returns()
func (_Vm *VmTransactor) ExpectEmit5(opts *bind.TransactOpts, emitter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit5", emitter, count)
}

// ExpectEmit5 is a paid mutator transaction binding the contract method 0xb43aece3.
//
// Solidity: function expectEmit(address emitter, uint64 count) returns()
func (_Vm *VmSession) ExpectEmit5(emitter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit5(&_Vm.TransactOpts, emitter, count)
}

// ExpectEmit5 is a paid mutator transaction binding the contract method 0xb43aece3.
//
// Solidity: function expectEmit(address emitter, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectEmit5(emitter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit5(&_Vm.TransactOpts, emitter, count)
}

// ExpectEmit6 is a paid mutator transaction binding the contract method 0xc339d02c.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter, uint64 count) returns()
func (_Vm *VmTransactor) ExpectEmit6(opts *bind.TransactOpts, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmit6", checkTopic1, checkTopic2, checkTopic3, checkData, emitter, count)
}

// ExpectEmit6 is a paid mutator transaction binding the contract method 0xc339d02c.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter, uint64 count) returns()
func (_Vm *VmSession) ExpectEmit6(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit6(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter, count)
}

// ExpectEmit6 is a paid mutator transaction binding the contract method 0xc339d02c.
//
// Solidity: function expectEmit(bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectEmit6(checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmit6(&_Vm.TransactOpts, checkTopic1, checkTopic2, checkTopic3, checkData, emitter, count)
}

// ExpectEmitAnonymous is a paid mutator transaction binding the contract method 0x2e5f270c.
//
// Solidity: function expectEmitAnonymous() returns()
func (_Vm *VmTransactor) ExpectEmitAnonymous(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmitAnonymous")
}

// ExpectEmitAnonymous is a paid mutator transaction binding the contract method 0x2e5f270c.
//
// Solidity: function expectEmitAnonymous() returns()
func (_Vm *VmSession) ExpectEmitAnonymous() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous(&_Vm.TransactOpts)
}

// ExpectEmitAnonymous is a paid mutator transaction binding the contract method 0x2e5f270c.
//
// Solidity: function expectEmitAnonymous() returns()
func (_Vm *VmTransactorSession) ExpectEmitAnonymous() (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous(&_Vm.TransactOpts)
}

// ExpectEmitAnonymous0 is a paid mutator transaction binding the contract method 0x6fc68705.
//
// Solidity: function expectEmitAnonymous(address emitter) returns()
func (_Vm *VmTransactor) ExpectEmitAnonymous0(opts *bind.TransactOpts, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmitAnonymous0", emitter)
}

// ExpectEmitAnonymous0 is a paid mutator transaction binding the contract method 0x6fc68705.
//
// Solidity: function expectEmitAnonymous(address emitter) returns()
func (_Vm *VmSession) ExpectEmitAnonymous0(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous0(&_Vm.TransactOpts, emitter)
}

// ExpectEmitAnonymous0 is a paid mutator transaction binding the contract method 0x6fc68705.
//
// Solidity: function expectEmitAnonymous(address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmitAnonymous0(emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous0(&_Vm.TransactOpts, emitter)
}

// ExpectEmitAnonymous1 is a paid mutator transaction binding the contract method 0x71c95899.
//
// Solidity: function expectEmitAnonymous(bool checkTopic0, bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactor) ExpectEmitAnonymous1(opts *bind.TransactOpts, checkTopic0 bool, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmitAnonymous1", checkTopic0, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmitAnonymous1 is a paid mutator transaction binding the contract method 0x71c95899.
//
// Solidity: function expectEmitAnonymous(bool checkTopic0, bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmSession) ExpectEmitAnonymous1(checkTopic0 bool, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous1(&_Vm.TransactOpts, checkTopic0, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmitAnonymous1 is a paid mutator transaction binding the contract method 0x71c95899.
//
// Solidity: function expectEmitAnonymous(bool checkTopic0, bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData, address emitter) returns()
func (_Vm *VmTransactorSession) ExpectEmitAnonymous1(checkTopic0 bool, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool, emitter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous1(&_Vm.TransactOpts, checkTopic0, checkTopic1, checkTopic2, checkTopic3, checkData, emitter)
}

// ExpectEmitAnonymous2 is a paid mutator transaction binding the contract method 0xc948db5e.
//
// Solidity: function expectEmitAnonymous(bool checkTopic0, bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactor) ExpectEmitAnonymous2(opts *bind.TransactOpts, checkTopic0 bool, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectEmitAnonymous2", checkTopic0, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmitAnonymous2 is a paid mutator transaction binding the contract method 0xc948db5e.
//
// Solidity: function expectEmitAnonymous(bool checkTopic0, bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmSession) ExpectEmitAnonymous2(checkTopic0 bool, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous2(&_Vm.TransactOpts, checkTopic0, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectEmitAnonymous2 is a paid mutator transaction binding the contract method 0xc948db5e.
//
// Solidity: function expectEmitAnonymous(bool checkTopic0, bool checkTopic1, bool checkTopic2, bool checkTopic3, bool checkData) returns()
func (_Vm *VmTransactorSession) ExpectEmitAnonymous2(checkTopic0 bool, checkTopic1 bool, checkTopic2 bool, checkTopic3 bool, checkData bool) (*types.Transaction, error) {
	return _Vm.Contract.ExpectEmitAnonymous2(&_Vm.TransactOpts, checkTopic0, checkTopic1, checkTopic2, checkTopic3, checkData)
}

// ExpectPartialRevert is a paid mutator transaction binding the contract method 0x11fb5b9c.
//
// Solidity: function expectPartialRevert(bytes4 revertData) returns()
func (_Vm *VmTransactor) ExpectPartialRevert(opts *bind.TransactOpts, revertData [4]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectPartialRevert", revertData)
}

// ExpectPartialRevert is a paid mutator transaction binding the contract method 0x11fb5b9c.
//
// Solidity: function expectPartialRevert(bytes4 revertData) returns()
func (_Vm *VmSession) ExpectPartialRevert(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectPartialRevert(&_Vm.TransactOpts, revertData)
}

// ExpectPartialRevert is a paid mutator transaction binding the contract method 0x11fb5b9c.
//
// Solidity: function expectPartialRevert(bytes4 revertData) returns()
func (_Vm *VmTransactorSession) ExpectPartialRevert(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectPartialRevert(&_Vm.TransactOpts, revertData)
}

// ExpectPartialRevert0 is a paid mutator transaction binding the contract method 0x51aa008a.
//
// Solidity: function expectPartialRevert(bytes4 revertData, address reverter) returns()
func (_Vm *VmTransactor) ExpectPartialRevert0(opts *bind.TransactOpts, revertData [4]byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectPartialRevert0", revertData, reverter)
}

// ExpectPartialRevert0 is a paid mutator transaction binding the contract method 0x51aa008a.
//
// Solidity: function expectPartialRevert(bytes4 revertData, address reverter) returns()
func (_Vm *VmSession) ExpectPartialRevert0(revertData [4]byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectPartialRevert0(&_Vm.TransactOpts, revertData, reverter)
}

// ExpectPartialRevert0 is a paid mutator transaction binding the contract method 0x51aa008a.
//
// Solidity: function expectPartialRevert(bytes4 revertData, address reverter) returns()
func (_Vm *VmTransactorSession) ExpectPartialRevert0(revertData [4]byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectPartialRevert0(&_Vm.TransactOpts, revertData, reverter)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0x1ff5f952.
//
// Solidity: function expectRevert(address reverter, uint64 count) returns()
func (_Vm *VmTransactor) ExpectRevert(opts *bind.TransactOpts, reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert", reverter, count)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0x1ff5f952.
//
// Solidity: function expectRevert(address reverter, uint64 count) returns()
func (_Vm *VmSession) ExpectRevert(reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert(&_Vm.TransactOpts, reverter, count)
}

// ExpectRevert is a paid mutator transaction binding the contract method 0x1ff5f952.
//
// Solidity: function expectRevert(address reverter, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectRevert(reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert(&_Vm.TransactOpts, reverter, count)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0x260bc5de.
//
// Solidity: function expectRevert(bytes4 revertData, address reverter) returns()
func (_Vm *VmTransactor) ExpectRevert0(opts *bind.TransactOpts, revertData [4]byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert0", revertData, reverter)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0x260bc5de.
//
// Solidity: function expectRevert(bytes4 revertData, address reverter) returns()
func (_Vm *VmSession) ExpectRevert0(revertData [4]byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert0(&_Vm.TransactOpts, revertData, reverter)
}

// ExpectRevert0 is a paid mutator transaction binding the contract method 0x260bc5de.
//
// Solidity: function expectRevert(bytes4 revertData, address reverter) returns()
func (_Vm *VmTransactorSession) ExpectRevert0(revertData [4]byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert0(&_Vm.TransactOpts, revertData, reverter)
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0x4994c273.
//
// Solidity: function expectRevert(bytes revertData, uint64 count) returns()
func (_Vm *VmTransactor) ExpectRevert1(opts *bind.TransactOpts, revertData []byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert1", revertData, count)
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0x4994c273.
//
// Solidity: function expectRevert(bytes revertData, uint64 count) returns()
func (_Vm *VmSession) ExpectRevert1(revertData []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert1(&_Vm.TransactOpts, revertData, count)
}

// ExpectRevert1 is a paid mutator transaction binding the contract method 0x4994c273.
//
// Solidity: function expectRevert(bytes revertData, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectRevert1(revertData []byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert1(&_Vm.TransactOpts, revertData, count)
}

// ExpectRevert10 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmTransactor) ExpectRevert10(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert10")
}

// ExpectRevert10 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmSession) ExpectRevert10() (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert10(&_Vm.TransactOpts)
}

// ExpectRevert10 is a paid mutator transaction binding the contract method 0xf4844814.
//
// Solidity: function expectRevert() returns()
func (_Vm *VmTransactorSession) ExpectRevert10() (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert10(&_Vm.TransactOpts)
}

// ExpectRevert2 is a paid mutator transaction binding the contract method 0x4ee38244.
//
// Solidity: function expectRevert(uint64 count) returns()
func (_Vm *VmTransactor) ExpectRevert2(opts *bind.TransactOpts, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert2", count)
}

// ExpectRevert2 is a paid mutator transaction binding the contract method 0x4ee38244.
//
// Solidity: function expectRevert(uint64 count) returns()
func (_Vm *VmSession) ExpectRevert2(count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert2(&_Vm.TransactOpts, count)
}

// ExpectRevert2 is a paid mutator transaction binding the contract method 0x4ee38244.
//
// Solidity: function expectRevert(uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectRevert2(count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert2(&_Vm.TransactOpts, count)
}

// ExpectRevert3 is a paid mutator transaction binding the contract method 0x61ebcf12.
//
// Solidity: function expectRevert(bytes revertData, address reverter) returns()
func (_Vm *VmTransactor) ExpectRevert3(opts *bind.TransactOpts, revertData []byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert3", revertData, reverter)
}

// ExpectRevert3 is a paid mutator transaction binding the contract method 0x61ebcf12.
//
// Solidity: function expectRevert(bytes revertData, address reverter) returns()
func (_Vm *VmSession) ExpectRevert3(revertData []byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert3(&_Vm.TransactOpts, revertData, reverter)
}

// ExpectRevert3 is a paid mutator transaction binding the contract method 0x61ebcf12.
//
// Solidity: function expectRevert(bytes revertData, address reverter) returns()
func (_Vm *VmTransactorSession) ExpectRevert3(revertData []byte, reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert3(&_Vm.TransactOpts, revertData, reverter)
}

// ExpectRevert4 is a paid mutator transaction binding the contract method 0xb0762d73.
//
// Solidity: function expectRevert(bytes4 revertData, address reverter, uint64 count) returns()
func (_Vm *VmTransactor) ExpectRevert4(opts *bind.TransactOpts, revertData [4]byte, reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert4", revertData, reverter, count)
}

// ExpectRevert4 is a paid mutator transaction binding the contract method 0xb0762d73.
//
// Solidity: function expectRevert(bytes4 revertData, address reverter, uint64 count) returns()
func (_Vm *VmSession) ExpectRevert4(revertData [4]byte, reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert4(&_Vm.TransactOpts, revertData, reverter, count)
}

// ExpectRevert4 is a paid mutator transaction binding the contract method 0xb0762d73.
//
// Solidity: function expectRevert(bytes4 revertData, address reverter, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectRevert4(revertData [4]byte, reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert4(&_Vm.TransactOpts, revertData, reverter, count)
}

// ExpectRevert5 is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmTransactor) ExpectRevert5(opts *bind.TransactOpts, revertData [4]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert5", revertData)
}

// ExpectRevert5 is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmSession) ExpectRevert5(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert5(&_Vm.TransactOpts, revertData)
}

// ExpectRevert5 is a paid mutator transaction binding the contract method 0xc31eb0e0.
//
// Solidity: function expectRevert(bytes4 revertData) returns()
func (_Vm *VmTransactorSession) ExpectRevert5(revertData [4]byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert5(&_Vm.TransactOpts, revertData)
}

// ExpectRevert6 is a paid mutator transaction binding the contract method 0xd345fb1f.
//
// Solidity: function expectRevert(bytes revertData, address reverter, uint64 count) returns()
func (_Vm *VmTransactor) ExpectRevert6(opts *bind.TransactOpts, revertData []byte, reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert6", revertData, reverter, count)
}

// ExpectRevert6 is a paid mutator transaction binding the contract method 0xd345fb1f.
//
// Solidity: function expectRevert(bytes revertData, address reverter, uint64 count) returns()
func (_Vm *VmSession) ExpectRevert6(revertData []byte, reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert6(&_Vm.TransactOpts, revertData, reverter, count)
}

// ExpectRevert6 is a paid mutator transaction binding the contract method 0xd345fb1f.
//
// Solidity: function expectRevert(bytes revertData, address reverter, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectRevert6(revertData []byte, reverter common.Address, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert6(&_Vm.TransactOpts, revertData, reverter, count)
}

// ExpectRevert7 is a paid mutator transaction binding the contract method 0xd814f38a.
//
// Solidity: function expectRevert(address reverter) returns()
func (_Vm *VmTransactor) ExpectRevert7(opts *bind.TransactOpts, reverter common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert7", reverter)
}

// ExpectRevert7 is a paid mutator transaction binding the contract method 0xd814f38a.
//
// Solidity: function expectRevert(address reverter) returns()
func (_Vm *VmSession) ExpectRevert7(reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert7(&_Vm.TransactOpts, reverter)
}

// ExpectRevert7 is a paid mutator transaction binding the contract method 0xd814f38a.
//
// Solidity: function expectRevert(address reverter) returns()
func (_Vm *VmTransactorSession) ExpectRevert7(reverter common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert7(&_Vm.TransactOpts, reverter)
}

// ExpectRevert8 is a paid mutator transaction binding the contract method 0xe45ca72d.
//
// Solidity: function expectRevert(bytes4 revertData, uint64 count) returns()
func (_Vm *VmTransactor) ExpectRevert8(opts *bind.TransactOpts, revertData [4]byte, count uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert8", revertData, count)
}

// ExpectRevert8 is a paid mutator transaction binding the contract method 0xe45ca72d.
//
// Solidity: function expectRevert(bytes4 revertData, uint64 count) returns()
func (_Vm *VmSession) ExpectRevert8(revertData [4]byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert8(&_Vm.TransactOpts, revertData, count)
}

// ExpectRevert8 is a paid mutator transaction binding the contract method 0xe45ca72d.
//
// Solidity: function expectRevert(bytes4 revertData, uint64 count) returns()
func (_Vm *VmTransactorSession) ExpectRevert8(revertData [4]byte, count uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert8(&_Vm.TransactOpts, revertData, count)
}

// ExpectRevert9 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmTransactor) ExpectRevert9(opts *bind.TransactOpts, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectRevert9", revertData)
}

// ExpectRevert9 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmSession) ExpectRevert9(revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert9(&_Vm.TransactOpts, revertData)
}

// ExpectRevert9 is a paid mutator transaction binding the contract method 0xf28dceb3.
//
// Solidity: function expectRevert(bytes revertData) returns()
func (_Vm *VmTransactorSession) ExpectRevert9(revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.ExpectRevert9(&_Vm.TransactOpts, revertData)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmTransactor) ExpectSafeMemory(opts *bind.TransactOpts, min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectSafeMemory", min, max)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmSession) ExpectSafeMemory(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemory(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemory is a paid mutator transaction binding the contract method 0x6d016688.
//
// Solidity: function expectSafeMemory(uint64 min, uint64 max) returns()
func (_Vm *VmTransactorSession) ExpectSafeMemory(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemory(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmTransactor) ExpectSafeMemoryCall(opts *bind.TransactOpts, min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "expectSafeMemoryCall", min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmSession) ExpectSafeMemoryCall(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemoryCall(&_Vm.TransactOpts, min, max)
}

// ExpectSafeMemoryCall is a paid mutator transaction binding the contract method 0x05838bf4.
//
// Solidity: function expectSafeMemoryCall(uint64 min, uint64 max) returns()
func (_Vm *VmTransactorSession) ExpectSafeMemoryCall(min uint64, max uint64) (*types.Transaction, error) {
	return _Vm.Contract.ExpectSafeMemoryCall(&_Vm.TransactOpts, min, max)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmTransactor) Fee(opts *bind.TransactOpts, newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "fee", newBasefee)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmSession) Fee(newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Fee(&_Vm.TransactOpts, newBasefee)
}

// Fee is a paid mutator transaction binding the contract method 0x39b37ab0.
//
// Solidity: function fee(uint256 newBasefee) returns()
func (_Vm *VmTransactorSession) Fee(newBasefee *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Fee(&_Vm.TransactOpts, newBasefee)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmTransactor) Ffi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "ffi", commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.Ffi(&_Vm.TransactOpts, commandInput)
}

// Ffi is a paid mutator transaction binding the contract method 0x89160467.
//
// Solidity: function ffi(string[] commandInput) returns(bytes result)
func (_Vm *VmTransactorSession) Ffi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.Ffi(&_Vm.TransactOpts, commandInput)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmTransactor) GetMappingKeyAndParentOf(opts *bind.TransactOpts, target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingKeyAndParentOf", target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingKeyAndParentOf(&_Vm.TransactOpts, target, elementSlot)
}

// GetMappingKeyAndParentOf is a paid mutator transaction binding the contract method 0x876e24e6.
//
// Solidity: function getMappingKeyAndParentOf(address target, bytes32 elementSlot) returns(bool found, bytes32 key, bytes32 parent)
func (_Vm *VmTransactorSession) GetMappingKeyAndParentOf(target common.Address, elementSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingKeyAndParentOf(&_Vm.TransactOpts, target, elementSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmTransactor) GetMappingLength(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingLength", target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingLength(&_Vm.TransactOpts, target, mappingSlot)
}

// GetMappingLength is a paid mutator transaction binding the contract method 0x2f2fd63f.
//
// Solidity: function getMappingLength(address target, bytes32 mappingSlot) returns(uint256 length)
func (_Vm *VmTransactorSession) GetMappingLength(target common.Address, mappingSlot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingLength(&_Vm.TransactOpts, target, mappingSlot)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmTransactor) GetMappingSlotAt(opts *bind.TransactOpts, target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getMappingSlotAt", target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingSlotAt(&_Vm.TransactOpts, target, mappingSlot, idx)
}

// GetMappingSlotAt is a paid mutator transaction binding the contract method 0xebc73ab4.
//
// Solidity: function getMappingSlotAt(address target, bytes32 mappingSlot, uint256 idx) returns(bytes32 value)
func (_Vm *VmTransactorSession) GetMappingSlotAt(target common.Address, mappingSlot [32]byte, idx *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.GetMappingSlotAt(&_Vm.TransactOpts, target, mappingSlot, idx)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmTransactor) GetNonce0(opts *bind.TransactOpts, wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getNonce0", wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.Contract.GetNonce0(&_Vm.TransactOpts, wallet)
}

// GetNonce0 is a paid mutator transaction binding the contract method 0xa5748aad.
//
// Solidity: function getNonce((address,uint256,uint256,uint256) wallet) returns(uint64 nonce)
func (_Vm *VmTransactorSession) GetNonce0(wallet VmSafeWallet) (*types.Transaction, error) {
	return _Vm.Contract.GetNonce0(&_Vm.TransactOpts, wallet)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmTransactor) GetRecordedLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getRecordedLogs")
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmSession) GetRecordedLogs() (*types.Transaction, error) {
	return _Vm.Contract.GetRecordedLogs(&_Vm.TransactOpts)
}

// GetRecordedLogs is a paid mutator transaction binding the contract method 0x191553a4.
//
// Solidity: function getRecordedLogs() returns((bytes32[],bytes,address)[] logs)
func (_Vm *VmTransactorSession) GetRecordedLogs() (*types.Transaction, error) {
	return _Vm.Contract.GetRecordedLogs(&_Vm.TransactOpts)
}

// GetWallets is a paid mutator transaction binding the contract method 0xdb7a4605.
//
// Solidity: function getWallets() returns(address[] wallets)
func (_Vm *VmTransactor) GetWallets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "getWallets")
}

// GetWallets is a paid mutator transaction binding the contract method 0xdb7a4605.
//
// Solidity: function getWallets() returns(address[] wallets)
func (_Vm *VmSession) GetWallets() (*types.Transaction, error) {
	return _Vm.Contract.GetWallets(&_Vm.TransactOpts)
}

// GetWallets is a paid mutator transaction binding the contract method 0xdb7a4605.
//
// Solidity: function getWallets() returns(address[] wallets)
func (_Vm *VmTransactorSession) GetWallets() (*types.Transaction, error) {
	return _Vm.Contract.GetWallets(&_Vm.TransactOpts)
}

// InterceptInitcode is a paid mutator transaction binding the contract method 0x838653c7.
//
// Solidity: function interceptInitcode() returns()
func (_Vm *VmTransactor) InterceptInitcode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "interceptInitcode")
}

// InterceptInitcode is a paid mutator transaction binding the contract method 0x838653c7.
//
// Solidity: function interceptInitcode() returns()
func (_Vm *VmSession) InterceptInitcode() (*types.Transaction, error) {
	return _Vm.Contract.InterceptInitcode(&_Vm.TransactOpts)
}

// InterceptInitcode is a paid mutator transaction binding the contract method 0x838653c7.
//
// Solidity: function interceptInitcode() returns()
func (_Vm *VmTransactorSession) InterceptInitcode() (*types.Transaction, error) {
	return _Vm.Contract.InterceptInitcode(&_Vm.TransactOpts)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmTransactor) Label(opts *bind.TransactOpts, account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "label", account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.Contract.Label(&_Vm.TransactOpts, account, newLabel)
}

// Label is a paid mutator transaction binding the contract method 0xc657c718.
//
// Solidity: function label(address account, string newLabel) returns()
func (_Vm *VmTransactorSession) Label(account common.Address, newLabel string) (*types.Transaction, error) {
	return _Vm.Contract.Label(&_Vm.TransactOpts, account, newLabel)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmTransactor) LoadAllocs(opts *bind.TransactOpts, pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "loadAllocs", pathToAllocsJson)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmSession) LoadAllocs(pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.Contract.LoadAllocs(&_Vm.TransactOpts, pathToAllocsJson)
}

// LoadAllocs is a paid mutator transaction binding the contract method 0xb3a056d7.
//
// Solidity: function loadAllocs(string pathToAllocsJson) returns()
func (_Vm *VmTransactorSession) LoadAllocs(pathToAllocsJson string) (*types.Transaction, error) {
	return _Vm.Contract.LoadAllocs(&_Vm.TransactOpts, pathToAllocsJson)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmTransactor) MakePersistent(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent", accounts)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmSession) MakePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent(&_Vm.TransactOpts, accounts)
}

// MakePersistent is a paid mutator transaction binding the contract method 0x1d9e269e.
//
// Solidity: function makePersistent(address[] accounts) returns()
func (_Vm *VmTransactorSession) MakePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent(&_Vm.TransactOpts, accounts)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmTransactor) MakePersistent0(opts *bind.TransactOpts, account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent0", account0, account1)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmSession) MakePersistent0(account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent0(&_Vm.TransactOpts, account0, account1)
}

// MakePersistent0 is a paid mutator transaction binding the contract method 0x4074e0a8.
//
// Solidity: function makePersistent(address account0, address account1) returns()
func (_Vm *VmTransactorSession) MakePersistent0(account0 common.Address, account1 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent0(&_Vm.TransactOpts, account0, account1)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmTransactor) MakePersistent1(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent1", account)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmSession) MakePersistent1(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent1(&_Vm.TransactOpts, account)
}

// MakePersistent1 is a paid mutator transaction binding the contract method 0x57e22dde.
//
// Solidity: function makePersistent(address account) returns()
func (_Vm *VmTransactorSession) MakePersistent1(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent1(&_Vm.TransactOpts, account)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmTransactor) MakePersistent2(opts *bind.TransactOpts, account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "makePersistent2", account0, account1, account2)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmSession) MakePersistent2(account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent2(&_Vm.TransactOpts, account0, account1, account2)
}

// MakePersistent2 is a paid mutator transaction binding the contract method 0xefb77a75.
//
// Solidity: function makePersistent(address account0, address account1, address account2) returns()
func (_Vm *VmTransactorSession) MakePersistent2(account0 common.Address, account1 common.Address, account2 common.Address) (*types.Transaction, error) {
	return _Vm.Contract.MakePersistent2(&_Vm.TransactOpts, account0, account1, account2)
}

// MockCall is a paid mutator transaction binding the contract method 0x08e0c537.
//
// Solidity: function mockCall(address callee, bytes4 data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall(opts *bind.TransactOpts, callee common.Address, data [4]byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall", callee, data, returnData)
}

// MockCall is a paid mutator transaction binding the contract method 0x08e0c537.
//
// Solidity: function mockCall(address callee, bytes4 data, bytes returnData) returns()
func (_Vm *VmSession) MockCall(callee common.Address, data [4]byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCall is a paid mutator transaction binding the contract method 0x08e0c537.
//
// Solidity: function mockCall(address callee, bytes4 data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall(callee common.Address, data [4]byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall0", callee, msgValue, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmSession) MockCall0(callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall0(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCall0 is a paid mutator transaction binding the contract method 0x81409b91.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall0(callee common.Address, msgValue *big.Int, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall0(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCall1 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall1(opts *bind.TransactOpts, callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall1", callee, data, returnData)
}

// MockCall1 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmSession) MockCall1(callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall1(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCall1 is a paid mutator transaction binding the contract method 0xb96213e4.
//
// Solidity: function mockCall(address callee, bytes data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall1(callee common.Address, data []byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall1(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCall2 is a paid mutator transaction binding the contract method 0xe7b36a3d.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes4 data, bytes returnData) returns()
func (_Vm *VmTransactor) MockCall2(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data [4]byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCall2", callee, msgValue, data, returnData)
}

// MockCall2 is a paid mutator transaction binding the contract method 0xe7b36a3d.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes4 data, bytes returnData) returns()
func (_Vm *VmSession) MockCall2(callee common.Address, msgValue *big.Int, data [4]byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall2(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCall2 is a paid mutator transaction binding the contract method 0xe7b36a3d.
//
// Solidity: function mockCall(address callee, uint256 msgValue, bytes4 data, bytes returnData) returns()
func (_Vm *VmTransactorSession) MockCall2(callee common.Address, msgValue *big.Int, data [4]byte, returnData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCall2(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0x2dfba5df.
//
// Solidity: function mockCallRevert(address callee, bytes4 data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert(opts *bind.TransactOpts, callee common.Address, data [4]byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert", callee, data, revertData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0x2dfba5df.
//
// Solidity: function mockCallRevert(address callee, bytes4 data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert(callee common.Address, data [4]byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert(&_Vm.TransactOpts, callee, data, revertData)
}

// MockCallRevert is a paid mutator transaction binding the contract method 0x2dfba5df.
//
// Solidity: function mockCallRevert(address callee, bytes4 data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert(callee common.Address, data [4]byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert(&_Vm.TransactOpts, callee, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0x596c8f04.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes4 data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert0(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data [4]byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert0", callee, msgValue, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0x596c8f04.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes4 data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert0(callee common.Address, msgValue *big.Int, data [4]byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert0(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert0 is a paid mutator transaction binding the contract method 0x596c8f04.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes4 data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert0(callee common.Address, msgValue *big.Int, data [4]byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert0(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert1 is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert1(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert1", callee, msgValue, data, revertData)
}

// MockCallRevert1 is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert1(callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert1(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert1 is a paid mutator transaction binding the contract method 0xd23cd037.
//
// Solidity: function mockCallRevert(address callee, uint256 msgValue, bytes data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert1(callee common.Address, msgValue *big.Int, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert1(&_Vm.TransactOpts, callee, msgValue, data, revertData)
}

// MockCallRevert2 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmTransactor) MockCallRevert2(opts *bind.TransactOpts, callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCallRevert2", callee, data, revertData)
}

// MockCallRevert2 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmSession) MockCallRevert2(callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert2(&_Vm.TransactOpts, callee, data, revertData)
}

// MockCallRevert2 is a paid mutator transaction binding the contract method 0xdbaad147.
//
// Solidity: function mockCallRevert(address callee, bytes data, bytes revertData) returns()
func (_Vm *VmTransactorSession) MockCallRevert2(callee common.Address, data []byte, revertData []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCallRevert2(&_Vm.TransactOpts, callee, data, revertData)
}

// MockCalls is a paid mutator transaction binding the contract method 0x08bcbae1.
//
// Solidity: function mockCalls(address callee, uint256 msgValue, bytes data, bytes[] returnData) returns()
func (_Vm *VmTransactor) MockCalls(opts *bind.TransactOpts, callee common.Address, msgValue *big.Int, data []byte, returnData [][]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCalls", callee, msgValue, data, returnData)
}

// MockCalls is a paid mutator transaction binding the contract method 0x08bcbae1.
//
// Solidity: function mockCalls(address callee, uint256 msgValue, bytes data, bytes[] returnData) returns()
func (_Vm *VmSession) MockCalls(callee common.Address, msgValue *big.Int, data []byte, returnData [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCalls(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCalls is a paid mutator transaction binding the contract method 0x08bcbae1.
//
// Solidity: function mockCalls(address callee, uint256 msgValue, bytes data, bytes[] returnData) returns()
func (_Vm *VmTransactorSession) MockCalls(callee common.Address, msgValue *big.Int, data []byte, returnData [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCalls(&_Vm.TransactOpts, callee, msgValue, data, returnData)
}

// MockCalls0 is a paid mutator transaction binding the contract method 0x5c5c3de9.
//
// Solidity: function mockCalls(address callee, bytes data, bytes[] returnData) returns()
func (_Vm *VmTransactor) MockCalls0(opts *bind.TransactOpts, callee common.Address, data []byte, returnData [][]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockCalls0", callee, data, returnData)
}

// MockCalls0 is a paid mutator transaction binding the contract method 0x5c5c3de9.
//
// Solidity: function mockCalls(address callee, bytes data, bytes[] returnData) returns()
func (_Vm *VmSession) MockCalls0(callee common.Address, data []byte, returnData [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCalls0(&_Vm.TransactOpts, callee, data, returnData)
}

// MockCalls0 is a paid mutator transaction binding the contract method 0x5c5c3de9.
//
// Solidity: function mockCalls(address callee, bytes data, bytes[] returnData) returns()
func (_Vm *VmTransactorSession) MockCalls0(callee common.Address, data []byte, returnData [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.MockCalls0(&_Vm.TransactOpts, callee, data, returnData)
}

// MockFunction is a paid mutator transaction binding the contract method 0xadf84d21.
//
// Solidity: function mockFunction(address callee, address target, bytes data) returns()
func (_Vm *VmTransactor) MockFunction(opts *bind.TransactOpts, callee common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "mockFunction", callee, target, data)
}

// MockFunction is a paid mutator transaction binding the contract method 0xadf84d21.
//
// Solidity: function mockFunction(address callee, address target, bytes data) returns()
func (_Vm *VmSession) MockFunction(callee common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockFunction(&_Vm.TransactOpts, callee, target, data)
}

// MockFunction is a paid mutator transaction binding the contract method 0xadf84d21.
//
// Solidity: function mockFunction(address callee, address target, bytes data) returns()
func (_Vm *VmTransactorSession) MockFunction(callee common.Address, target common.Address, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.MockFunction(&_Vm.TransactOpts, callee, target, data)
}

// NoAccessList is a paid mutator transaction binding the contract method 0x238ad778.
//
// Solidity: function noAccessList() returns()
func (_Vm *VmTransactor) NoAccessList(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "noAccessList")
}

// NoAccessList is a paid mutator transaction binding the contract method 0x238ad778.
//
// Solidity: function noAccessList() returns()
func (_Vm *VmSession) NoAccessList() (*types.Transaction, error) {
	return _Vm.Contract.NoAccessList(&_Vm.TransactOpts)
}

// NoAccessList is a paid mutator transaction binding the contract method 0x238ad778.
//
// Solidity: function noAccessList() returns()
func (_Vm *VmTransactorSession) NoAccessList() (*types.Transaction, error) {
	return _Vm.Contract.NoAccessList(&_Vm.TransactOpts)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmTransactor) PauseGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "pauseGasMetering")
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmSession) PauseGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.PauseGasMetering(&_Vm.TransactOpts)
}

// PauseGasMetering is a paid mutator transaction binding the contract method 0xd1a5b36f.
//
// Solidity: function pauseGasMetering() returns()
func (_Vm *VmTransactorSession) PauseGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.PauseGasMetering(&_Vm.TransactOpts)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactor) Prank(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank", msgSender, txOrigin)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmSession) Prank(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank(&_Vm.TransactOpts, msgSender, txOrigin)
}

// Prank is a paid mutator transaction binding the contract method 0x47e50cce.
//
// Solidity: function prank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactorSession) Prank(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank(&_Vm.TransactOpts, msgSender, txOrigin)
}

// Prank0 is a paid mutator transaction binding the contract method 0x7d73d042.
//
// Solidity: function prank(address msgSender, address txOrigin, bool delegateCall) returns()
func (_Vm *VmTransactor) Prank0(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank0", msgSender, txOrigin, delegateCall)
}

// Prank0 is a paid mutator transaction binding the contract method 0x7d73d042.
//
// Solidity: function prank(address msgSender, address txOrigin, bool delegateCall) returns()
func (_Vm *VmSession) Prank0(msgSender common.Address, txOrigin common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.Prank0(&_Vm.TransactOpts, msgSender, txOrigin, delegateCall)
}

// Prank0 is a paid mutator transaction binding the contract method 0x7d73d042.
//
// Solidity: function prank(address msgSender, address txOrigin, bool delegateCall) returns()
func (_Vm *VmTransactorSession) Prank0(msgSender common.Address, txOrigin common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.Prank0(&_Vm.TransactOpts, msgSender, txOrigin, delegateCall)
}

// Prank1 is a paid mutator transaction binding the contract method 0xa7f8bf5c.
//
// Solidity: function prank(address msgSender, bool delegateCall) returns()
func (_Vm *VmTransactor) Prank1(opts *bind.TransactOpts, msgSender common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank1", msgSender, delegateCall)
}

// Prank1 is a paid mutator transaction binding the contract method 0xa7f8bf5c.
//
// Solidity: function prank(address msgSender, bool delegateCall) returns()
func (_Vm *VmSession) Prank1(msgSender common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.Prank1(&_Vm.TransactOpts, msgSender, delegateCall)
}

// Prank1 is a paid mutator transaction binding the contract method 0xa7f8bf5c.
//
// Solidity: function prank(address msgSender, bool delegateCall) returns()
func (_Vm *VmTransactorSession) Prank1(msgSender common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.Prank1(&_Vm.TransactOpts, msgSender, delegateCall)
}

// Prank2 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmTransactor) Prank2(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prank2", msgSender)
}

// Prank2 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmSession) Prank2(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank2(&_Vm.TransactOpts, msgSender)
}

// Prank2 is a paid mutator transaction binding the contract method 0xca669fa7.
//
// Solidity: function prank(address msgSender) returns()
func (_Vm *VmTransactorSession) Prank2(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.Prank2(&_Vm.TransactOpts, msgSender)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmTransactor) Prevrandao(opts *bind.TransactOpts, newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prevrandao", newPrevrandao)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmSession) Prevrandao(newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao(&_Vm.TransactOpts, newPrevrandao)
}

// Prevrandao is a paid mutator transaction binding the contract method 0x3b925549.
//
// Solidity: function prevrandao(bytes32 newPrevrandao) returns()
func (_Vm *VmTransactorSession) Prevrandao(newPrevrandao [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao(&_Vm.TransactOpts, newPrevrandao)
}

// Prevrandao0 is a paid mutator transaction binding the contract method 0x9cb1c0d4.
//
// Solidity: function prevrandao(uint256 newPrevrandao) returns()
func (_Vm *VmTransactor) Prevrandao0(opts *bind.TransactOpts, newPrevrandao *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prevrandao0", newPrevrandao)
}

// Prevrandao0 is a paid mutator transaction binding the contract method 0x9cb1c0d4.
//
// Solidity: function prevrandao(uint256 newPrevrandao) returns()
func (_Vm *VmSession) Prevrandao0(newPrevrandao *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao0(&_Vm.TransactOpts, newPrevrandao)
}

// Prevrandao0 is a paid mutator transaction binding the contract method 0x9cb1c0d4.
//
// Solidity: function prevrandao(uint256 newPrevrandao) returns()
func (_Vm *VmTransactorSession) Prevrandao0(newPrevrandao *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Prevrandao0(&_Vm.TransactOpts, newPrevrandao)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_Vm *VmTransactor) Prompt(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "prompt", promptText)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_Vm *VmSession) Prompt(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.Prompt(&_Vm.TransactOpts, promptText)
}

// Prompt is a paid mutator transaction binding the contract method 0x47eaf474.
//
// Solidity: function prompt(string promptText) returns(string input)
func (_Vm *VmTransactorSession) Prompt(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.Prompt(&_Vm.TransactOpts, promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_Vm *VmTransactor) PromptAddress(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptAddress", promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_Vm *VmSession) PromptAddress(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptAddress(&_Vm.TransactOpts, promptText)
}

// PromptAddress is a paid mutator transaction binding the contract method 0x62ee05f4.
//
// Solidity: function promptAddress(string promptText) returns(address)
func (_Vm *VmTransactorSession) PromptAddress(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptAddress(&_Vm.TransactOpts, promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_Vm *VmTransactor) PromptSecret(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptSecret", promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_Vm *VmSession) PromptSecret(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecret(&_Vm.TransactOpts, promptText)
}

// PromptSecret is a paid mutator transaction binding the contract method 0x1e279d41.
//
// Solidity: function promptSecret(string promptText) returns(string input)
func (_Vm *VmTransactorSession) PromptSecret(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecret(&_Vm.TransactOpts, promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_Vm *VmTransactor) PromptSecretUint(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptSecretUint", promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_Vm *VmSession) PromptSecretUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecretUint(&_Vm.TransactOpts, promptText)
}

// PromptSecretUint is a paid mutator transaction binding the contract method 0x69ca02b7.
//
// Solidity: function promptSecretUint(string promptText) returns(uint256)
func (_Vm *VmTransactorSession) PromptSecretUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptSecretUint(&_Vm.TransactOpts, promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_Vm *VmTransactor) PromptUint(opts *bind.TransactOpts, promptText string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "promptUint", promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_Vm *VmSession) PromptUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptUint(&_Vm.TransactOpts, promptText)
}

// PromptUint is a paid mutator transaction binding the contract method 0x652fd489.
//
// Solidity: function promptUint(string promptText) returns(uint256)
func (_Vm *VmTransactorSession) PromptUint(promptText string) (*types.Transaction, error) {
	return _Vm.Contract.PromptUint(&_Vm.TransactOpts, promptText)
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_Vm *VmTransactor) RandomAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "randomAddress")
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_Vm *VmSession) RandomAddress() (*types.Transaction, error) {
	return _Vm.Contract.RandomAddress(&_Vm.TransactOpts)
}

// RandomAddress is a paid mutator transaction binding the contract method 0xd5bee9f5.
//
// Solidity: function randomAddress() returns(address)
func (_Vm *VmTransactorSession) RandomAddress() (*types.Transaction, error) {
	return _Vm.Contract.RandomAddress(&_Vm.TransactOpts)
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_Vm *VmTransactor) RandomUint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "randomUint")
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_Vm *VmSession) RandomUint() (*types.Transaction, error) {
	return _Vm.Contract.RandomUint(&_Vm.TransactOpts)
}

// RandomUint is a paid mutator transaction binding the contract method 0x25124730.
//
// Solidity: function randomUint() returns(uint256)
func (_Vm *VmTransactorSession) RandomUint() (*types.Transaction, error) {
	return _Vm.Contract.RandomUint(&_Vm.TransactOpts)
}

// RandomUint1 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_Vm *VmTransactor) RandomUint1(opts *bind.TransactOpts, min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "randomUint1", min, max)
}

// RandomUint1 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_Vm *VmSession) RandomUint1(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RandomUint1(&_Vm.TransactOpts, min, max)
}

// RandomUint1 is a paid mutator transaction binding the contract method 0xd61b051b.
//
// Solidity: function randomUint(uint256 min, uint256 max) returns(uint256)
func (_Vm *VmTransactorSession) RandomUint1(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RandomUint1(&_Vm.TransactOpts, min, max)
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmTransactor) ReadCallers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "readCallers")
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmSession) ReadCallers() (*types.Transaction, error) {
	return _Vm.Contract.ReadCallers(&_Vm.TransactOpts)
}

// ReadCallers is a paid mutator transaction binding the contract method 0x4ad0bac9.
//
// Solidity: function readCallers() returns(uint8 callerMode, address msgSender, address txOrigin)
func (_Vm *VmTransactorSession) ReadCallers() (*types.Transaction, error) {
	return _Vm.Contract.ReadCallers(&_Vm.TransactOpts)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmTransactor) Record(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "record")
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmSession) Record() (*types.Transaction, error) {
	return _Vm.Contract.Record(&_Vm.TransactOpts)
}

// Record is a paid mutator transaction binding the contract method 0x266cf109.
//
// Solidity: function record() returns()
func (_Vm *VmTransactorSession) Record() (*types.Transaction, error) {
	return _Vm.Contract.Record(&_Vm.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmTransactor) RecordLogs(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "recordLogs")
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmSession) RecordLogs() (*types.Transaction, error) {
	return _Vm.Contract.RecordLogs(&_Vm.TransactOpts)
}

// RecordLogs is a paid mutator transaction binding the contract method 0x41af2f52.
//
// Solidity: function recordLogs() returns()
func (_Vm *VmTransactorSession) RecordLogs() (*types.Transaction, error) {
	return _Vm.Contract.RecordLogs(&_Vm.TransactOpts)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmTransactor) RememberKey(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rememberKey", privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RememberKey(&_Vm.TransactOpts, privateKey)
}

// RememberKey is a paid mutator transaction binding the contract method 0x22100064.
//
// Solidity: function rememberKey(uint256 privateKey) returns(address keyAddr)
func (_Vm *VmTransactorSession) RememberKey(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RememberKey(&_Vm.TransactOpts, privateKey)
}

// RememberKeys is a paid mutator transaction binding the contract method 0x97cb9189.
//
// Solidity: function rememberKeys(string mnemonic, string derivationPath, uint32 count) returns(address[] keyAddrs)
func (_Vm *VmTransactor) RememberKeys(opts *bind.TransactOpts, mnemonic string, derivationPath string, count uint32) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rememberKeys", mnemonic, derivationPath, count)
}

// RememberKeys is a paid mutator transaction binding the contract method 0x97cb9189.
//
// Solidity: function rememberKeys(string mnemonic, string derivationPath, uint32 count) returns(address[] keyAddrs)
func (_Vm *VmSession) RememberKeys(mnemonic string, derivationPath string, count uint32) (*types.Transaction, error) {
	return _Vm.Contract.RememberKeys(&_Vm.TransactOpts, mnemonic, derivationPath, count)
}

// RememberKeys is a paid mutator transaction binding the contract method 0x97cb9189.
//
// Solidity: function rememberKeys(string mnemonic, string derivationPath, uint32 count) returns(address[] keyAddrs)
func (_Vm *VmTransactorSession) RememberKeys(mnemonic string, derivationPath string, count uint32) (*types.Transaction, error) {
	return _Vm.Contract.RememberKeys(&_Vm.TransactOpts, mnemonic, derivationPath, count)
}

// RememberKeys0 is a paid mutator transaction binding the contract method 0xf8d58eaf.
//
// Solidity: function rememberKeys(string mnemonic, string derivationPath, string language, uint32 count) returns(address[] keyAddrs)
func (_Vm *VmTransactor) RememberKeys0(opts *bind.TransactOpts, mnemonic string, derivationPath string, language string, count uint32) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rememberKeys0", mnemonic, derivationPath, language, count)
}

// RememberKeys0 is a paid mutator transaction binding the contract method 0xf8d58eaf.
//
// Solidity: function rememberKeys(string mnemonic, string derivationPath, string language, uint32 count) returns(address[] keyAddrs)
func (_Vm *VmSession) RememberKeys0(mnemonic string, derivationPath string, language string, count uint32) (*types.Transaction, error) {
	return _Vm.Contract.RememberKeys0(&_Vm.TransactOpts, mnemonic, derivationPath, language, count)
}

// RememberKeys0 is a paid mutator transaction binding the contract method 0xf8d58eaf.
//
// Solidity: function rememberKeys(string mnemonic, string derivationPath, string language, uint32 count) returns(address[] keyAddrs)
func (_Vm *VmTransactorSession) RememberKeys0(mnemonic string, derivationPath string, language string, count uint32) (*types.Transaction, error) {
	return _Vm.Contract.RememberKeys0(&_Vm.TransactOpts, mnemonic, derivationPath, language, count)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmTransactor) RemoveDir(opts *bind.TransactOpts, path string, recursive bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "removeDir", path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.RemoveDir(&_Vm.TransactOpts, path, recursive)
}

// RemoveDir is a paid mutator transaction binding the contract method 0x45c62011.
//
// Solidity: function removeDir(string path, bool recursive) returns()
func (_Vm *VmTransactorSession) RemoveDir(path string, recursive bool) (*types.Transaction, error) {
	return _Vm.Contract.RemoveDir(&_Vm.TransactOpts, path, recursive)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmTransactor) RemoveFile(opts *bind.TransactOpts, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "removeFile", path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmSession) RemoveFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.RemoveFile(&_Vm.TransactOpts, path)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string path) returns()
func (_Vm *VmTransactorSession) RemoveFile(path string) (*types.Transaction, error) {
	return _Vm.Contract.RemoveFile(&_Vm.TransactOpts, path)
}

// ResetGasMetering is a paid mutator transaction binding the contract method 0xbe367dd3.
//
// Solidity: function resetGasMetering() returns()
func (_Vm *VmTransactor) ResetGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "resetGasMetering")
}

// ResetGasMetering is a paid mutator transaction binding the contract method 0xbe367dd3.
//
// Solidity: function resetGasMetering() returns()
func (_Vm *VmSession) ResetGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResetGasMetering(&_Vm.TransactOpts)
}

// ResetGasMetering is a paid mutator transaction binding the contract method 0xbe367dd3.
//
// Solidity: function resetGasMetering() returns()
func (_Vm *VmTransactorSession) ResetGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResetGasMetering(&_Vm.TransactOpts)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmTransactor) ResetNonce(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "resetNonce", account)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmSession) ResetNonce(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ResetNonce(&_Vm.TransactOpts, account)
}

// ResetNonce is a paid mutator transaction binding the contract method 0x1c72346d.
//
// Solidity: function resetNonce(address account) returns()
func (_Vm *VmTransactorSession) ResetNonce(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.ResetNonce(&_Vm.TransactOpts, account)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmTransactor) ResumeGasMetering(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "resumeGasMetering")
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmSession) ResumeGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResumeGasMetering(&_Vm.TransactOpts)
}

// ResumeGasMetering is a paid mutator transaction binding the contract method 0x2bcd50e0.
//
// Solidity: function resumeGasMetering() returns()
func (_Vm *VmTransactorSession) ResumeGasMetering() (*types.Transaction, error) {
	return _Vm.Contract.ResumeGasMetering(&_Vm.TransactOpts)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertTo(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertTo", snapshotId)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertTo(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertTo(&_Vm.TransactOpts, snapshotId)
}

// RevertTo is a paid mutator transaction binding the contract method 0x44d7f0a4.
//
// Solidity: function revertTo(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertTo(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertTo(&_Vm.TransactOpts, snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertToAndDelete(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertToAndDelete", snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertToAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevertToAndDelete is a paid mutator transaction binding the contract method 0x03e0aca9.
//
// Solidity: function revertToAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertToAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevertToState is a paid mutator transaction binding the contract method 0xc2527405.
//
// Solidity: function revertToState(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertToState(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertToState", snapshotId)
}

// RevertToState is a paid mutator transaction binding the contract method 0xc2527405.
//
// Solidity: function revertToState(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertToState(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToState(&_Vm.TransactOpts, snapshotId)
}

// RevertToState is a paid mutator transaction binding the contract method 0xc2527405.
//
// Solidity: function revertToState(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertToState(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToState(&_Vm.TransactOpts, snapshotId)
}

// RevertToStateAndDelete is a paid mutator transaction binding the contract method 0x3a1985dc.
//
// Solidity: function revertToStateAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactor) RevertToStateAndDelete(opts *bind.TransactOpts, snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revertToStateAndDelete", snapshotId)
}

// RevertToStateAndDelete is a paid mutator transaction binding the contract method 0x3a1985dc.
//
// Solidity: function revertToStateAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmSession) RevertToStateAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToStateAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevertToStateAndDelete is a paid mutator transaction binding the contract method 0x3a1985dc.
//
// Solidity: function revertToStateAndDelete(uint256 snapshotId) returns(bool success)
func (_Vm *VmTransactorSession) RevertToStateAndDelete(snapshotId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RevertToStateAndDelete(&_Vm.TransactOpts, snapshotId)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmTransactor) RevokePersistent(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revokePersistent", accounts)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmSession) RevokePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent(&_Vm.TransactOpts, accounts)
}

// RevokePersistent is a paid mutator transaction binding the contract method 0x3ce969e6.
//
// Solidity: function revokePersistent(address[] accounts) returns()
func (_Vm *VmTransactorSession) RevokePersistent(accounts []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent(&_Vm.TransactOpts, accounts)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmTransactor) RevokePersistent0(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "revokePersistent0", account)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmSession) RevokePersistent0(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent0(&_Vm.TransactOpts, account)
}

// RevokePersistent0 is a paid mutator transaction binding the contract method 0x997a0222.
//
// Solidity: function revokePersistent(address account) returns()
func (_Vm *VmTransactorSession) RevokePersistent0(account common.Address) (*types.Transaction, error) {
	return _Vm.Contract.RevokePersistent0(&_Vm.TransactOpts, account)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmTransactor) Roll(opts *bind.TransactOpts, newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "roll", newHeight)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmSession) Roll(newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Roll(&_Vm.TransactOpts, newHeight)
}

// Roll is a paid mutator transaction binding the contract method 0x1f7b4f30.
//
// Solidity: function roll(uint256 newHeight) returns()
func (_Vm *VmTransactorSession) Roll(newHeight *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Roll(&_Vm.TransactOpts, newHeight)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmTransactor) RollFork(opts *bind.TransactOpts, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork", txHash)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmSession) RollFork(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork(&_Vm.TransactOpts, txHash)
}

// RollFork is a paid mutator transaction binding the contract method 0x0f29772b.
//
// Solidity: function rollFork(bytes32 txHash) returns()
func (_Vm *VmTransactorSession) RollFork(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork(&_Vm.TransactOpts, txHash)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmTransactor) RollFork0(opts *bind.TransactOpts, forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork0", forkId, blockNumber)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmSession) RollFork0(forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork0(&_Vm.TransactOpts, forkId, blockNumber)
}

// RollFork0 is a paid mutator transaction binding the contract method 0xd74c83a4.
//
// Solidity: function rollFork(uint256 forkId, uint256 blockNumber) returns()
func (_Vm *VmTransactorSession) RollFork0(forkId *big.Int, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork0(&_Vm.TransactOpts, forkId, blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmTransactor) RollFork1(opts *bind.TransactOpts, blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork1", blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmSession) RollFork1(blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork1(&_Vm.TransactOpts, blockNumber)
}

// RollFork1 is a paid mutator transaction binding the contract method 0xd9bbf3a1.
//
// Solidity: function rollFork(uint256 blockNumber) returns()
func (_Vm *VmTransactorSession) RollFork1(blockNumber *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.RollFork1(&_Vm.TransactOpts, blockNumber)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactor) RollFork2(opts *bind.TransactOpts, forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rollFork2", forkId, txHash)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmSession) RollFork2(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork2(&_Vm.TransactOpts, forkId, txHash)
}

// RollFork2 is a paid mutator transaction binding the contract method 0xf2830f7b.
//
// Solidity: function rollFork(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactorSession) RollFork2(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.RollFork2(&_Vm.TransactOpts, forkId, txHash)
}

// Rpc is a paid mutator transaction binding the contract method 0x0199a220.
//
// Solidity: function rpc(string urlOrAlias, string method, string params) returns(bytes data)
func (_Vm *VmTransactor) Rpc(opts *bind.TransactOpts, urlOrAlias string, method string, params string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rpc", urlOrAlias, method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x0199a220.
//
// Solidity: function rpc(string urlOrAlias, string method, string params) returns(bytes data)
func (_Vm *VmSession) Rpc(urlOrAlias string, method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc(&_Vm.TransactOpts, urlOrAlias, method, params)
}

// Rpc is a paid mutator transaction binding the contract method 0x0199a220.
//
// Solidity: function rpc(string urlOrAlias, string method, string params) returns(bytes data)
func (_Vm *VmTransactorSession) Rpc(urlOrAlias string, method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc(&_Vm.TransactOpts, urlOrAlias, method, params)
}

// Rpc0 is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmTransactor) Rpc0(opts *bind.TransactOpts, method string, params string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "rpc0", method, params)
}

// Rpc0 is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmSession) Rpc0(method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc0(&_Vm.TransactOpts, method, params)
}

// Rpc0 is a paid mutator transaction binding the contract method 0x1206c8a8.
//
// Solidity: function rpc(string method, string params) returns(bytes data)
func (_Vm *VmTransactorSession) Rpc0(method string, params string) (*types.Transaction, error) {
	return _Vm.Contract.Rpc0(&_Vm.TransactOpts, method, params)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmTransactor) SelectFork(opts *bind.TransactOpts, forkId *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "selectFork", forkId)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmSession) SelectFork(forkId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SelectFork(&_Vm.TransactOpts, forkId)
}

// SelectFork is a paid mutator transaction binding the contract method 0x9ebf6827.
//
// Solidity: function selectFork(uint256 forkId) returns()
func (_Vm *VmTransactorSession) SelectFork(forkId *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SelectFork(&_Vm.TransactOpts, forkId)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmTransactor) SerializeAddress(opts *bind.TransactOpts, objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeAddress", objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress is a paid mutator transaction binding the contract method 0x1e356e1a.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeAddress(objectKey string, valueKey string, values []common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmTransactor) SerializeAddress0(opts *bind.TransactOpts, objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeAddress0", objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeAddress0 is a paid mutator transaction binding the contract method 0x972c6062.
//
// Solidity: function serializeAddress(string objectKey, string valueKey, address value) returns(string json)
func (_Vm *VmTransactorSession) SerializeAddress0(objectKey string, valueKey string, value common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SerializeAddress0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBool(opts *bind.TransactOpts, objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBool", objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool is a paid mutator transaction binding the contract method 0x92925aa1.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBool(objectKey string, valueKey string, values []bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmTransactor) SerializeBool0(opts *bind.TransactOpts, objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBool0", objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBool0 is a paid mutator transaction binding the contract method 0xac22e971.
//
// Solidity: function serializeBool(string objectKey, string valueKey, bool value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBool0(objectKey string, valueKey string, value bool) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBool0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBytes(opts *bind.TransactOpts, objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes", objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes is a paid mutator transaction binding the contract method 0x9884b232.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes(objectKey string, valueKey string, values [][]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmTransactor) SerializeBytes0(opts *bind.TransactOpts, objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes0", objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes0 is a paid mutator transaction binding the contract method 0xf21d52c7.
//
// Solidity: function serializeBytes(string objectKey, string valueKey, bytes value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes0(objectKey string, valueKey string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmTransactor) SerializeBytes32(opts *bind.TransactOpts, objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes32", objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes32(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes32 is a paid mutator transaction binding the contract method 0x201e43e2.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes32(objectKey string, valueKey string, values [][32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes32(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmTransactor) SerializeBytes320(opts *bind.TransactOpts, objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeBytes320", objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes320(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeBytes320 is a paid mutator transaction binding the contract method 0x2d812b44.
//
// Solidity: function serializeBytes32(string objectKey, string valueKey, bytes32 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeBytes320(objectKey string, valueKey string, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeBytes320(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmTransactor) SerializeInt(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeInt", objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt is a paid mutator transaction binding the contract method 0x3f33db60.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeInt(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmTransactor) SerializeInt0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeInt0", objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeInt0 is a paid mutator transaction binding the contract method 0x7676e127.
//
// Solidity: function serializeInt(string objectKey, string valueKey, int256[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeInt0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeInt0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmTransactor) SerializeJson(opts *bind.TransactOpts, objectKey string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeJson", objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJson(&_Vm.TransactOpts, objectKey, value)
}

// SerializeJson is a paid mutator transaction binding the contract method 0x9b3358b0.
//
// Solidity: function serializeJson(string objectKey, string value) returns(string json)
func (_Vm *VmTransactorSession) SerializeJson(objectKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJson(&_Vm.TransactOpts, objectKey, value)
}

// SerializeJsonType0 is a paid mutator transaction binding the contract method 0x6f93bccb.
//
// Solidity: function serializeJsonType(string objectKey, string valueKey, string typeDescription, bytes value) returns(string json)
func (_Vm *VmTransactor) SerializeJsonType0(opts *bind.TransactOpts, objectKey string, valueKey string, typeDescription string, value []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeJsonType0", objectKey, valueKey, typeDescription, value)
}

// SerializeJsonType0 is a paid mutator transaction binding the contract method 0x6f93bccb.
//
// Solidity: function serializeJsonType(string objectKey, string valueKey, string typeDescription, bytes value) returns(string json)
func (_Vm *VmSession) SerializeJsonType0(objectKey string, valueKey string, typeDescription string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJsonType0(&_Vm.TransactOpts, objectKey, valueKey, typeDescription, value)
}

// SerializeJsonType0 is a paid mutator transaction binding the contract method 0x6f93bccb.
//
// Solidity: function serializeJsonType(string objectKey, string valueKey, string typeDescription, bytes value) returns(string json)
func (_Vm *VmTransactorSession) SerializeJsonType0(objectKey string, valueKey string, typeDescription string, value []byte) (*types.Transaction, error) {
	return _Vm.Contract.SerializeJsonType0(&_Vm.TransactOpts, objectKey, valueKey, typeDescription, value)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmTransactor) SerializeString(opts *bind.TransactOpts, objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeString", objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeString is a paid mutator transaction binding the contract method 0x561cd6f3.
//
// Solidity: function serializeString(string objectKey, string valueKey, string[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeString(objectKey string, valueKey string, values []string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmTransactor) SerializeString0(opts *bind.TransactOpts, objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeString0", objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeString0 is a paid mutator transaction binding the contract method 0x88da6d35.
//
// Solidity: function serializeString(string objectKey, string valueKey, string value) returns(string json)
func (_Vm *VmTransactorSession) SerializeString0(objectKey string, valueKey string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SerializeString0(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactor) SerializeUint(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUint", objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint is a paid mutator transaction binding the contract method 0x129e9002.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeUint(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmTransactor) SerializeUint0(opts *bind.TransactOpts, objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUint0", objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeUint0 is a paid mutator transaction binding the contract method 0xfee9a469.
//
// Solidity: function serializeUint(string objectKey, string valueKey, uint256[] values) returns(string json)
func (_Vm *VmTransactorSession) SerializeUint0(objectKey string, valueKey string, values []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUint0(&_Vm.TransactOpts, objectKey, valueKey, values)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactor) SerializeUintToHex(opts *bind.TransactOpts, objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "serializeUintToHex", objectKey, valueKey, value)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmSession) SerializeUintToHex(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUintToHex(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SerializeUintToHex is a paid mutator transaction binding the contract method 0xae5a2ae8.
//
// Solidity: function serializeUintToHex(string objectKey, string valueKey, uint256 value) returns(string json)
func (_Vm *VmTransactorSession) SerializeUintToHex(objectKey string, valueKey string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SerializeUintToHex(&_Vm.TransactOpts, objectKey, valueKey, value)
}

// SetArbitraryStorage is a paid mutator transaction binding the contract method 0xd3ec2a0b.
//
// Solidity: function setArbitraryStorage(address target, bool overwrite) returns()
func (_Vm *VmTransactor) SetArbitraryStorage(opts *bind.TransactOpts, target common.Address, overwrite bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setArbitraryStorage", target, overwrite)
}

// SetArbitraryStorage is a paid mutator transaction binding the contract method 0xd3ec2a0b.
//
// Solidity: function setArbitraryStorage(address target, bool overwrite) returns()
func (_Vm *VmSession) SetArbitraryStorage(target common.Address, overwrite bool) (*types.Transaction, error) {
	return _Vm.Contract.SetArbitraryStorage(&_Vm.TransactOpts, target, overwrite)
}

// SetArbitraryStorage is a paid mutator transaction binding the contract method 0xd3ec2a0b.
//
// Solidity: function setArbitraryStorage(address target, bool overwrite) returns()
func (_Vm *VmTransactorSession) SetArbitraryStorage(target common.Address, overwrite bool) (*types.Transaction, error) {
	return _Vm.Contract.SetArbitraryStorage(&_Vm.TransactOpts, target, overwrite)
}

// SetArbitraryStorage0 is a paid mutator transaction binding the contract method 0xe1631837.
//
// Solidity: function setArbitraryStorage(address target) returns()
func (_Vm *VmTransactor) SetArbitraryStorage0(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setArbitraryStorage0", target)
}

// SetArbitraryStorage0 is a paid mutator transaction binding the contract method 0xe1631837.
//
// Solidity: function setArbitraryStorage(address target) returns()
func (_Vm *VmSession) SetArbitraryStorage0(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SetArbitraryStorage0(&_Vm.TransactOpts, target)
}

// SetArbitraryStorage0 is a paid mutator transaction binding the contract method 0xe1631837.
//
// Solidity: function setArbitraryStorage(address target) returns()
func (_Vm *VmTransactorSession) SetArbitraryStorage0(target common.Address) (*types.Transaction, error) {
	return _Vm.Contract.SetArbitraryStorage0(&_Vm.TransactOpts, target)
}

// SetBlockhash is a paid mutator transaction binding the contract method 0x5314b54a.
//
// Solidity: function setBlockhash(uint256 blockNumber, bytes32 blockHash) returns()
func (_Vm *VmTransactor) SetBlockhash(opts *bind.TransactOpts, blockNumber *big.Int, blockHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setBlockhash", blockNumber, blockHash)
}

// SetBlockhash is a paid mutator transaction binding the contract method 0x5314b54a.
//
// Solidity: function setBlockhash(uint256 blockNumber, bytes32 blockHash) returns()
func (_Vm *VmSession) SetBlockhash(blockNumber *big.Int, blockHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SetBlockhash(&_Vm.TransactOpts, blockNumber, blockHash)
}

// SetBlockhash is a paid mutator transaction binding the contract method 0x5314b54a.
//
// Solidity: function setBlockhash(uint256 blockNumber, bytes32 blockHash) returns()
func (_Vm *VmTransactorSession) SetBlockhash(blockNumber *big.Int, blockHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SetBlockhash(&_Vm.TransactOpts, blockNumber, blockHash)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmTransactor) SetEnv(opts *bind.TransactOpts, name string, value string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setEnv", name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SetEnv(&_Vm.TransactOpts, name, value)
}

// SetEnv is a paid mutator transaction binding the contract method 0x3d5923ee.
//
// Solidity: function setEnv(string name, string value) returns()
func (_Vm *VmTransactorSession) SetEnv(name string, value string) (*types.Transaction, error) {
	return _Vm.Contract.SetEnv(&_Vm.TransactOpts, name, value)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmTransactor) SetNonce(opts *bind.TransactOpts, account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setNonce", account, newNonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmSession) SetNonce(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonce(&_Vm.TransactOpts, account, newNonce)
}

// SetNonce is a paid mutator transaction binding the contract method 0xf8e18b57.
//
// Solidity: function setNonce(address account, uint64 newNonce) returns()
func (_Vm *VmTransactorSession) SetNonce(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonce(&_Vm.TransactOpts, account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmTransactor) SetNonceUnsafe(opts *bind.TransactOpts, account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "setNonceUnsafe", account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmSession) SetNonceUnsafe(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonceUnsafe(&_Vm.TransactOpts, account, newNonce)
}

// SetNonceUnsafe is a paid mutator transaction binding the contract method 0x9b67b21c.
//
// Solidity: function setNonceUnsafe(address account, uint64 newNonce) returns()
func (_Vm *VmTransactorSession) SetNonceUnsafe(account common.Address, newNonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SetNonceUnsafe(&_Vm.TransactOpts, account, newNonce)
}

// Shuffle is a paid mutator transaction binding the contract method 0x54f1469c.
//
// Solidity: function shuffle(uint256[] array) returns(uint256[])
func (_Vm *VmTransactor) Shuffle(opts *bind.TransactOpts, array []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "shuffle", array)
}

// Shuffle is a paid mutator transaction binding the contract method 0x54f1469c.
//
// Solidity: function shuffle(uint256[] array) returns(uint256[])
func (_Vm *VmSession) Shuffle(array []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Shuffle(&_Vm.TransactOpts, array)
}

// Shuffle is a paid mutator transaction binding the contract method 0x54f1469c.
//
// Solidity: function shuffle(uint256[] array) returns(uint256[])
func (_Vm *VmTransactorSession) Shuffle(array []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Shuffle(&_Vm.TransactOpts, array)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmTransactor) Sign1(opts *bind.TransactOpts, wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "sign1", wallet, digest)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmSession) Sign1(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Sign1(&_Vm.TransactOpts, wallet, digest)
}

// Sign1 is a paid mutator transaction binding the contract method 0xb25c5a25.
//
// Solidity: function sign((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(uint8 v, bytes32 r, bytes32 s)
func (_Vm *VmTransactorSession) Sign1(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Sign1(&_Vm.TransactOpts, wallet, digest)
}

// SignAndAttachDelegation is a paid mutator transaction binding the contract method 0xc7fa7288.
//
// Solidity: function signAndAttachDelegation(address implementation, uint256 privateKey) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactor) SignAndAttachDelegation(opts *bind.TransactOpts, implementation common.Address, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "signAndAttachDelegation", implementation, privateKey)
}

// SignAndAttachDelegation is a paid mutator transaction binding the contract method 0xc7fa7288.
//
// Solidity: function signAndAttachDelegation(address implementation, uint256 privateKey) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmSession) SignAndAttachDelegation(implementation common.Address, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SignAndAttachDelegation(&_Vm.TransactOpts, implementation, privateKey)
}

// SignAndAttachDelegation is a paid mutator transaction binding the contract method 0xc7fa7288.
//
// Solidity: function signAndAttachDelegation(address implementation, uint256 privateKey) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactorSession) SignAndAttachDelegation(implementation common.Address, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SignAndAttachDelegation(&_Vm.TransactOpts, implementation, privateKey)
}

// SignAndAttachDelegation0 is a paid mutator transaction binding the contract method 0xcde3e5be.
//
// Solidity: function signAndAttachDelegation(address implementation, uint256 privateKey, uint64 nonce) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactor) SignAndAttachDelegation0(opts *bind.TransactOpts, implementation common.Address, privateKey *big.Int, nonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "signAndAttachDelegation0", implementation, privateKey, nonce)
}

// SignAndAttachDelegation0 is a paid mutator transaction binding the contract method 0xcde3e5be.
//
// Solidity: function signAndAttachDelegation(address implementation, uint256 privateKey, uint64 nonce) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmSession) SignAndAttachDelegation0(implementation common.Address, privateKey *big.Int, nonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SignAndAttachDelegation0(&_Vm.TransactOpts, implementation, privateKey, nonce)
}

// SignAndAttachDelegation0 is a paid mutator transaction binding the contract method 0xcde3e5be.
//
// Solidity: function signAndAttachDelegation(address implementation, uint256 privateKey, uint64 nonce) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactorSession) SignAndAttachDelegation0(implementation common.Address, privateKey *big.Int, nonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SignAndAttachDelegation0(&_Vm.TransactOpts, implementation, privateKey, nonce)
}

// SignCompact is a paid mutator transaction binding the contract method 0x3d0e292f.
//
// Solidity: function signCompact((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(bytes32 r, bytes32 vs)
func (_Vm *VmTransactor) SignCompact(opts *bind.TransactOpts, wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "signCompact", wallet, digest)
}

// SignCompact is a paid mutator transaction binding the contract method 0x3d0e292f.
//
// Solidity: function signCompact((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(bytes32 r, bytes32 vs)
func (_Vm *VmSession) SignCompact(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SignCompact(&_Vm.TransactOpts, wallet, digest)
}

// SignCompact is a paid mutator transaction binding the contract method 0x3d0e292f.
//
// Solidity: function signCompact((address,uint256,uint256,uint256) wallet, bytes32 digest) returns(bytes32 r, bytes32 vs)
func (_Vm *VmTransactorSession) SignCompact(wallet VmSafeWallet, digest [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.SignCompact(&_Vm.TransactOpts, wallet, digest)
}

// SignDelegation is a paid mutator transaction binding the contract method 0x5b593c7b.
//
// Solidity: function signDelegation(address implementation, uint256 privateKey) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactor) SignDelegation(opts *bind.TransactOpts, implementation common.Address, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "signDelegation", implementation, privateKey)
}

// SignDelegation is a paid mutator transaction binding the contract method 0x5b593c7b.
//
// Solidity: function signDelegation(address implementation, uint256 privateKey) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmSession) SignDelegation(implementation common.Address, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SignDelegation(&_Vm.TransactOpts, implementation, privateKey)
}

// SignDelegation is a paid mutator transaction binding the contract method 0x5b593c7b.
//
// Solidity: function signDelegation(address implementation, uint256 privateKey) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactorSession) SignDelegation(implementation common.Address, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SignDelegation(&_Vm.TransactOpts, implementation, privateKey)
}

// SignDelegation0 is a paid mutator transaction binding the contract method 0xceba2ec3.
//
// Solidity: function signDelegation(address implementation, uint256 privateKey, uint64 nonce) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactor) SignDelegation0(opts *bind.TransactOpts, implementation common.Address, privateKey *big.Int, nonce uint64) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "signDelegation0", implementation, privateKey, nonce)
}

// SignDelegation0 is a paid mutator transaction binding the contract method 0xceba2ec3.
//
// Solidity: function signDelegation(address implementation, uint256 privateKey, uint64 nonce) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmSession) SignDelegation0(implementation common.Address, privateKey *big.Int, nonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SignDelegation0(&_Vm.TransactOpts, implementation, privateKey, nonce)
}

// SignDelegation0 is a paid mutator transaction binding the contract method 0xceba2ec3.
//
// Solidity: function signDelegation(address implementation, uint256 privateKey, uint64 nonce) returns((uint8,bytes32,bytes32,uint64,address) signedDelegation)
func (_Vm *VmTransactorSession) SignDelegation0(implementation common.Address, privateKey *big.Int, nonce uint64) (*types.Transaction, error) {
	return _Vm.Contract.SignDelegation0(&_Vm.TransactOpts, implementation, privateKey, nonce)
}

// Skip is a paid mutator transaction binding the contract method 0xc42a80a7.
//
// Solidity: function skip(bool skipTest, string reason) returns()
func (_Vm *VmTransactor) Skip(opts *bind.TransactOpts, skipTest bool, reason string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "skip", skipTest, reason)
}

// Skip is a paid mutator transaction binding the contract method 0xc42a80a7.
//
// Solidity: function skip(bool skipTest, string reason) returns()
func (_Vm *VmSession) Skip(skipTest bool, reason string) (*types.Transaction, error) {
	return _Vm.Contract.Skip(&_Vm.TransactOpts, skipTest, reason)
}

// Skip is a paid mutator transaction binding the contract method 0xc42a80a7.
//
// Solidity: function skip(bool skipTest, string reason) returns()
func (_Vm *VmTransactorSession) Skip(skipTest bool, reason string) (*types.Transaction, error) {
	return _Vm.Contract.Skip(&_Vm.TransactOpts, skipTest, reason)
}

// Skip0 is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmTransactor) Skip0(opts *bind.TransactOpts, skipTest bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "skip0", skipTest)
}

// Skip0 is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmSession) Skip0(skipTest bool) (*types.Transaction, error) {
	return _Vm.Contract.Skip0(&_Vm.TransactOpts, skipTest)
}

// Skip0 is a paid mutator transaction binding the contract method 0xdd82d13e.
//
// Solidity: function skip(bool skipTest) returns()
func (_Vm *VmTransactorSession) Skip0(skipTest bool) (*types.Transaction, error) {
	return _Vm.Contract.Skip0(&_Vm.TransactOpts, skipTest)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmTransactor) Sleep(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "sleep", duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sleep(&_Vm.TransactOpts, duration)
}

// Sleep is a paid mutator transaction binding the contract method 0xfa9d8713.
//
// Solidity: function sleep(uint256 duration) returns()
func (_Vm *VmTransactorSession) Sleep(duration *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sleep(&_Vm.TransactOpts, duration)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmTransactor) Snapshot(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshot")
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmSession) Snapshot() (*types.Transaction, error) {
	return _Vm.Contract.Snapshot(&_Vm.TransactOpts)
}

// Snapshot is a paid mutator transaction binding the contract method 0x9711715a.
//
// Solidity: function snapshot() returns(uint256 snapshotId)
func (_Vm *VmTransactorSession) Snapshot() (*types.Transaction, error) {
	return _Vm.Contract.Snapshot(&_Vm.TransactOpts)
}

// SnapshotGasLastCall is a paid mutator transaction binding the contract method 0x200c6772.
//
// Solidity: function snapshotGasLastCall(string group, string name) returns(uint256 gasUsed)
func (_Vm *VmTransactor) SnapshotGasLastCall(opts *bind.TransactOpts, group string, name string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshotGasLastCall", group, name)
}

// SnapshotGasLastCall is a paid mutator transaction binding the contract method 0x200c6772.
//
// Solidity: function snapshotGasLastCall(string group, string name) returns(uint256 gasUsed)
func (_Vm *VmSession) SnapshotGasLastCall(group string, name string) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotGasLastCall(&_Vm.TransactOpts, group, name)
}

// SnapshotGasLastCall is a paid mutator transaction binding the contract method 0x200c6772.
//
// Solidity: function snapshotGasLastCall(string group, string name) returns(uint256 gasUsed)
func (_Vm *VmTransactorSession) SnapshotGasLastCall(group string, name string) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotGasLastCall(&_Vm.TransactOpts, group, name)
}

// SnapshotGasLastCall0 is a paid mutator transaction binding the contract method 0xdd9fca12.
//
// Solidity: function snapshotGasLastCall(string name) returns(uint256 gasUsed)
func (_Vm *VmTransactor) SnapshotGasLastCall0(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshotGasLastCall0", name)
}

// SnapshotGasLastCall0 is a paid mutator transaction binding the contract method 0xdd9fca12.
//
// Solidity: function snapshotGasLastCall(string name) returns(uint256 gasUsed)
func (_Vm *VmSession) SnapshotGasLastCall0(name string) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotGasLastCall0(&_Vm.TransactOpts, name)
}

// SnapshotGasLastCall0 is a paid mutator transaction binding the contract method 0xdd9fca12.
//
// Solidity: function snapshotGasLastCall(string name) returns(uint256 gasUsed)
func (_Vm *VmTransactorSession) SnapshotGasLastCall0(name string) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotGasLastCall0(&_Vm.TransactOpts, name)
}

// SnapshotState is a paid mutator transaction binding the contract method 0x9cd23835.
//
// Solidity: function snapshotState() returns(uint256 snapshotId)
func (_Vm *VmTransactor) SnapshotState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshotState")
}

// SnapshotState is a paid mutator transaction binding the contract method 0x9cd23835.
//
// Solidity: function snapshotState() returns(uint256 snapshotId)
func (_Vm *VmSession) SnapshotState() (*types.Transaction, error) {
	return _Vm.Contract.SnapshotState(&_Vm.TransactOpts)
}

// SnapshotState is a paid mutator transaction binding the contract method 0x9cd23835.
//
// Solidity: function snapshotState() returns(uint256 snapshotId)
func (_Vm *VmTransactorSession) SnapshotState() (*types.Transaction, error) {
	return _Vm.Contract.SnapshotState(&_Vm.TransactOpts)
}

// SnapshotValue is a paid mutator transaction binding the contract method 0x51db805a.
//
// Solidity: function snapshotValue(string name, uint256 value) returns()
func (_Vm *VmTransactor) SnapshotValue(opts *bind.TransactOpts, name string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshotValue", name, value)
}

// SnapshotValue is a paid mutator transaction binding the contract method 0x51db805a.
//
// Solidity: function snapshotValue(string name, uint256 value) returns()
func (_Vm *VmSession) SnapshotValue(name string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotValue(&_Vm.TransactOpts, name, value)
}

// SnapshotValue is a paid mutator transaction binding the contract method 0x51db805a.
//
// Solidity: function snapshotValue(string name, uint256 value) returns()
func (_Vm *VmTransactorSession) SnapshotValue(name string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotValue(&_Vm.TransactOpts, name, value)
}

// SnapshotValue0 is a paid mutator transaction binding the contract method 0x6d2b27d8.
//
// Solidity: function snapshotValue(string group, string name, uint256 value) returns()
func (_Vm *VmTransactor) SnapshotValue0(opts *bind.TransactOpts, group string, name string, value *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "snapshotValue0", group, name, value)
}

// SnapshotValue0 is a paid mutator transaction binding the contract method 0x6d2b27d8.
//
// Solidity: function snapshotValue(string group, string name, uint256 value) returns()
func (_Vm *VmSession) SnapshotValue0(group string, name string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotValue0(&_Vm.TransactOpts, group, name, value)
}

// SnapshotValue0 is a paid mutator transaction binding the contract method 0x6d2b27d8.
//
// Solidity: function snapshotValue(string group, string name, uint256 value) returns()
func (_Vm *VmTransactorSession) SnapshotValue0(group string, name string, value *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.SnapshotValue0(&_Vm.TransactOpts, group, name, value)
}

// Sort is a paid mutator transaction binding the contract method 0x9ec8b026.
//
// Solidity: function sort(uint256[] array) returns(uint256[])
func (_Vm *VmTransactor) Sort(opts *bind.TransactOpts, array []*big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "sort", array)
}

// Sort is a paid mutator transaction binding the contract method 0x9ec8b026.
//
// Solidity: function sort(uint256[] array) returns(uint256[])
func (_Vm *VmSession) Sort(array []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sort(&_Vm.TransactOpts, array)
}

// Sort is a paid mutator transaction binding the contract method 0x9ec8b026.
//
// Solidity: function sort(uint256[] array) returns(uint256[])
func (_Vm *VmTransactorSession) Sort(array []*big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Sort(&_Vm.TransactOpts, array)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmTransactor) StartBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast")
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmSession) StartBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast(&_Vm.TransactOpts)
}

// StartBroadcast is a paid mutator transaction binding the contract method 0x7fb5297f.
//
// Solidity: function startBroadcast() returns()
func (_Vm *VmTransactorSession) StartBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast(&_Vm.TransactOpts)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmTransactor) StartBroadcast0(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast0", signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast0(&_Vm.TransactOpts, signer)
}

// StartBroadcast0 is a paid mutator transaction binding the contract method 0x7fec2a8d.
//
// Solidity: function startBroadcast(address signer) returns()
func (_Vm *VmTransactorSession) StartBroadcast0(signer common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast0(&_Vm.TransactOpts, signer)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmTransactor) StartBroadcast1(opts *bind.TransactOpts, privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startBroadcast1", privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast1(&_Vm.TransactOpts, privateKey)
}

// StartBroadcast1 is a paid mutator transaction binding the contract method 0xce817d47.
//
// Solidity: function startBroadcast(uint256 privateKey) returns()
func (_Vm *VmTransactorSession) StartBroadcast1(privateKey *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.StartBroadcast1(&_Vm.TransactOpts, privateKey)
}

// StartDebugTraceRecording is a paid mutator transaction binding the contract method 0x419c8832.
//
// Solidity: function startDebugTraceRecording() returns()
func (_Vm *VmTransactor) StartDebugTraceRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startDebugTraceRecording")
}

// StartDebugTraceRecording is a paid mutator transaction binding the contract method 0x419c8832.
//
// Solidity: function startDebugTraceRecording() returns()
func (_Vm *VmSession) StartDebugTraceRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartDebugTraceRecording(&_Vm.TransactOpts)
}

// StartDebugTraceRecording is a paid mutator transaction binding the contract method 0x419c8832.
//
// Solidity: function startDebugTraceRecording() returns()
func (_Vm *VmTransactorSession) StartDebugTraceRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartDebugTraceRecording(&_Vm.TransactOpts)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmTransactor) StartMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startMappingRecording")
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmSession) StartMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartMappingRecording(&_Vm.TransactOpts)
}

// StartMappingRecording is a paid mutator transaction binding the contract method 0x3e9705c0.
//
// Solidity: function startMappingRecording() returns()
func (_Vm *VmTransactorSession) StartMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartMappingRecording(&_Vm.TransactOpts)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmTransactor) StartPrank(opts *bind.TransactOpts, msgSender common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank", msgSender)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmSession) StartPrank(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank(&_Vm.TransactOpts, msgSender)
}

// StartPrank is a paid mutator transaction binding the contract method 0x06447d56.
//
// Solidity: function startPrank(address msgSender) returns()
func (_Vm *VmTransactorSession) StartPrank(msgSender common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank(&_Vm.TransactOpts, msgSender)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x1cc0b435.
//
// Solidity: function startPrank(address msgSender, bool delegateCall) returns()
func (_Vm *VmTransactor) StartPrank0(opts *bind.TransactOpts, msgSender common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank0", msgSender, delegateCall)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x1cc0b435.
//
// Solidity: function startPrank(address msgSender, bool delegateCall) returns()
func (_Vm *VmSession) StartPrank0(msgSender common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank0(&_Vm.TransactOpts, msgSender, delegateCall)
}

// StartPrank0 is a paid mutator transaction binding the contract method 0x1cc0b435.
//
// Solidity: function startPrank(address msgSender, bool delegateCall) returns()
func (_Vm *VmTransactorSession) StartPrank0(msgSender common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank0(&_Vm.TransactOpts, msgSender, delegateCall)
}

// StartPrank1 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactor) StartPrank1(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank1", msgSender, txOrigin)
}

// StartPrank1 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmSession) StartPrank1(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank1(&_Vm.TransactOpts, msgSender, txOrigin)
}

// StartPrank1 is a paid mutator transaction binding the contract method 0x45b56078.
//
// Solidity: function startPrank(address msgSender, address txOrigin) returns()
func (_Vm *VmTransactorSession) StartPrank1(msgSender common.Address, txOrigin common.Address) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank1(&_Vm.TransactOpts, msgSender, txOrigin)
}

// StartPrank2 is a paid mutator transaction binding the contract method 0x4eb859b5.
//
// Solidity: function startPrank(address msgSender, address txOrigin, bool delegateCall) returns()
func (_Vm *VmTransactor) StartPrank2(opts *bind.TransactOpts, msgSender common.Address, txOrigin common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startPrank2", msgSender, txOrigin, delegateCall)
}

// StartPrank2 is a paid mutator transaction binding the contract method 0x4eb859b5.
//
// Solidity: function startPrank(address msgSender, address txOrigin, bool delegateCall) returns()
func (_Vm *VmSession) StartPrank2(msgSender common.Address, txOrigin common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank2(&_Vm.TransactOpts, msgSender, txOrigin, delegateCall)
}

// StartPrank2 is a paid mutator transaction binding the contract method 0x4eb859b5.
//
// Solidity: function startPrank(address msgSender, address txOrigin, bool delegateCall) returns()
func (_Vm *VmTransactorSession) StartPrank2(msgSender common.Address, txOrigin common.Address, delegateCall bool) (*types.Transaction, error) {
	return _Vm.Contract.StartPrank2(&_Vm.TransactOpts, msgSender, txOrigin, delegateCall)
}

// StartSnapshotGas is a paid mutator transaction binding the contract method 0x3cad9d7b.
//
// Solidity: function startSnapshotGas(string name) returns()
func (_Vm *VmTransactor) StartSnapshotGas(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startSnapshotGas", name)
}

// StartSnapshotGas is a paid mutator transaction binding the contract method 0x3cad9d7b.
//
// Solidity: function startSnapshotGas(string name) returns()
func (_Vm *VmSession) StartSnapshotGas(name string) (*types.Transaction, error) {
	return _Vm.Contract.StartSnapshotGas(&_Vm.TransactOpts, name)
}

// StartSnapshotGas is a paid mutator transaction binding the contract method 0x3cad9d7b.
//
// Solidity: function startSnapshotGas(string name) returns()
func (_Vm *VmTransactorSession) StartSnapshotGas(name string) (*types.Transaction, error) {
	return _Vm.Contract.StartSnapshotGas(&_Vm.TransactOpts, name)
}

// StartSnapshotGas0 is a paid mutator transaction binding the contract method 0x6cd0cc53.
//
// Solidity: function startSnapshotGas(string group, string name) returns()
func (_Vm *VmTransactor) StartSnapshotGas0(opts *bind.TransactOpts, group string, name string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startSnapshotGas0", group, name)
}

// StartSnapshotGas0 is a paid mutator transaction binding the contract method 0x6cd0cc53.
//
// Solidity: function startSnapshotGas(string group, string name) returns()
func (_Vm *VmSession) StartSnapshotGas0(group string, name string) (*types.Transaction, error) {
	return _Vm.Contract.StartSnapshotGas0(&_Vm.TransactOpts, group, name)
}

// StartSnapshotGas0 is a paid mutator transaction binding the contract method 0x6cd0cc53.
//
// Solidity: function startSnapshotGas(string group, string name) returns()
func (_Vm *VmTransactorSession) StartSnapshotGas0(group string, name string) (*types.Transaction, error) {
	return _Vm.Contract.StartSnapshotGas0(&_Vm.TransactOpts, group, name)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmTransactor) StartStateDiffRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "startStateDiffRecording")
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartStateDiffRecording(&_Vm.TransactOpts)
}

// StartStateDiffRecording is a paid mutator transaction binding the contract method 0xcf22e3c9.
//
// Solidity: function startStateDiffRecording() returns()
func (_Vm *VmTransactorSession) StartStateDiffRecording() (*types.Transaction, error) {
	return _Vm.Contract.StartStateDiffRecording(&_Vm.TransactOpts)
}

// StopAndReturnDebugTraceRecording is a paid mutator transaction binding the contract method 0xced398a2.
//
// Solidity: function stopAndReturnDebugTraceRecording() returns((uint256[],bytes,uint8,uint64,bool,address)[] step)
func (_Vm *VmTransactor) StopAndReturnDebugTraceRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopAndReturnDebugTraceRecording")
}

// StopAndReturnDebugTraceRecording is a paid mutator transaction binding the contract method 0xced398a2.
//
// Solidity: function stopAndReturnDebugTraceRecording() returns((uint256[],bytes,uint8,uint64,bool,address)[] step)
func (_Vm *VmSession) StopAndReturnDebugTraceRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnDebugTraceRecording(&_Vm.TransactOpts)
}

// StopAndReturnDebugTraceRecording is a paid mutator transaction binding the contract method 0xced398a2.
//
// Solidity: function stopAndReturnDebugTraceRecording() returns((uint256[],bytes,uint8,uint64,bool,address)[] step)
func (_Vm *VmTransactorSession) StopAndReturnDebugTraceRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnDebugTraceRecording(&_Vm.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_Vm *VmTransactor) StopAndReturnStateDiff(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopAndReturnStateDiff")
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_Vm *VmSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnStateDiff(&_Vm.TransactOpts)
}

// StopAndReturnStateDiff is a paid mutator transaction binding the contract method 0xaa5cf90e.
//
// Solidity: function stopAndReturnStateDiff() returns(((uint256,uint256),uint8,address,address,bool,uint256,uint256,bytes,uint256,bytes,bool,(address,bytes32,bool,bytes32,bytes32,bool)[],uint64)[] accountAccesses)
func (_Vm *VmTransactorSession) StopAndReturnStateDiff() (*types.Transaction, error) {
	return _Vm.Contract.StopAndReturnStateDiff(&_Vm.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmTransactor) StopBroadcast(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopBroadcast")
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmSession) StopBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StopBroadcast(&_Vm.TransactOpts)
}

// StopBroadcast is a paid mutator transaction binding the contract method 0x76eadd36.
//
// Solidity: function stopBroadcast() returns()
func (_Vm *VmTransactorSession) StopBroadcast() (*types.Transaction, error) {
	return _Vm.Contract.StopBroadcast(&_Vm.TransactOpts)
}

// StopExpectSafeMemory is a paid mutator transaction binding the contract method 0x0956441b.
//
// Solidity: function stopExpectSafeMemory() returns()
func (_Vm *VmTransactor) StopExpectSafeMemory(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopExpectSafeMemory")
}

// StopExpectSafeMemory is a paid mutator transaction binding the contract method 0x0956441b.
//
// Solidity: function stopExpectSafeMemory() returns()
func (_Vm *VmSession) StopExpectSafeMemory() (*types.Transaction, error) {
	return _Vm.Contract.StopExpectSafeMemory(&_Vm.TransactOpts)
}

// StopExpectSafeMemory is a paid mutator transaction binding the contract method 0x0956441b.
//
// Solidity: function stopExpectSafeMemory() returns()
func (_Vm *VmTransactorSession) StopExpectSafeMemory() (*types.Transaction, error) {
	return _Vm.Contract.StopExpectSafeMemory(&_Vm.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmTransactor) StopMappingRecording(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopMappingRecording")
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmSession) StopMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopMappingRecording(&_Vm.TransactOpts)
}

// StopMappingRecording is a paid mutator transaction binding the contract method 0x0d4aae9b.
//
// Solidity: function stopMappingRecording() returns()
func (_Vm *VmTransactorSession) StopMappingRecording() (*types.Transaction, error) {
	return _Vm.Contract.StopMappingRecording(&_Vm.TransactOpts)
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmTransactor) StopPrank(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopPrank")
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmSession) StopPrank() (*types.Transaction, error) {
	return _Vm.Contract.StopPrank(&_Vm.TransactOpts)
}

// StopPrank is a paid mutator transaction binding the contract method 0x90c5013b.
//
// Solidity: function stopPrank() returns()
func (_Vm *VmTransactorSession) StopPrank() (*types.Transaction, error) {
	return _Vm.Contract.StopPrank(&_Vm.TransactOpts)
}

// StopSnapshotGas is a paid mutator transaction binding the contract method 0x0c9db707.
//
// Solidity: function stopSnapshotGas(string group, string name) returns(uint256 gasUsed)
func (_Vm *VmTransactor) StopSnapshotGas(opts *bind.TransactOpts, group string, name string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopSnapshotGas", group, name)
}

// StopSnapshotGas is a paid mutator transaction binding the contract method 0x0c9db707.
//
// Solidity: function stopSnapshotGas(string group, string name) returns(uint256 gasUsed)
func (_Vm *VmSession) StopSnapshotGas(group string, name string) (*types.Transaction, error) {
	return _Vm.Contract.StopSnapshotGas(&_Vm.TransactOpts, group, name)
}

// StopSnapshotGas is a paid mutator transaction binding the contract method 0x0c9db707.
//
// Solidity: function stopSnapshotGas(string group, string name) returns(uint256 gasUsed)
func (_Vm *VmTransactorSession) StopSnapshotGas(group string, name string) (*types.Transaction, error) {
	return _Vm.Contract.StopSnapshotGas(&_Vm.TransactOpts, group, name)
}

// StopSnapshotGas0 is a paid mutator transaction binding the contract method 0x773b2805.
//
// Solidity: function stopSnapshotGas(string name) returns(uint256 gasUsed)
func (_Vm *VmTransactor) StopSnapshotGas0(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopSnapshotGas0", name)
}

// StopSnapshotGas0 is a paid mutator transaction binding the contract method 0x773b2805.
//
// Solidity: function stopSnapshotGas(string name) returns(uint256 gasUsed)
func (_Vm *VmSession) StopSnapshotGas0(name string) (*types.Transaction, error) {
	return _Vm.Contract.StopSnapshotGas0(&_Vm.TransactOpts, name)
}

// StopSnapshotGas0 is a paid mutator transaction binding the contract method 0x773b2805.
//
// Solidity: function stopSnapshotGas(string name) returns(uint256 gasUsed)
func (_Vm *VmTransactorSession) StopSnapshotGas0(name string) (*types.Transaction, error) {
	return _Vm.Contract.StopSnapshotGas0(&_Vm.TransactOpts, name)
}

// StopSnapshotGas1 is a paid mutator transaction binding the contract method 0xf6402eda.
//
// Solidity: function stopSnapshotGas() returns(uint256 gasUsed)
func (_Vm *VmTransactor) StopSnapshotGas1(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "stopSnapshotGas1")
}

// StopSnapshotGas1 is a paid mutator transaction binding the contract method 0xf6402eda.
//
// Solidity: function stopSnapshotGas() returns(uint256 gasUsed)
func (_Vm *VmSession) StopSnapshotGas1() (*types.Transaction, error) {
	return _Vm.Contract.StopSnapshotGas1(&_Vm.TransactOpts)
}

// StopSnapshotGas1 is a paid mutator transaction binding the contract method 0xf6402eda.
//
// Solidity: function stopSnapshotGas() returns(uint256 gasUsed)
func (_Vm *VmTransactorSession) StopSnapshotGas1() (*types.Transaction, error) {
	return _Vm.Contract.StopSnapshotGas1(&_Vm.TransactOpts)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmTransactor) Store(opts *bind.TransactOpts, target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "store", target, slot, value)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmSession) Store(target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Store(&_Vm.TransactOpts, target, slot, value)
}

// Store is a paid mutator transaction binding the contract method 0x70ca10bb.
//
// Solidity: function store(address target, bytes32 slot, bytes32 value) returns()
func (_Vm *VmTransactorSession) Store(target common.Address, slot [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Store(&_Vm.TransactOpts, target, slot, value)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactor) Transact(opts *bind.TransactOpts, forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "transact", forkId, txHash)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmSession) Transact(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact(&_Vm.TransactOpts, forkId, txHash)
}

// Transact is a paid mutator transaction binding the contract method 0x4d8abc4b.
//
// Solidity: function transact(uint256 forkId, bytes32 txHash) returns()
func (_Vm *VmTransactorSession) Transact(forkId *big.Int, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact(&_Vm.TransactOpts, forkId, txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmTransactor) Transact0(opts *bind.TransactOpts, txHash [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "transact0", txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmSession) Transact0(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact0(&_Vm.TransactOpts, txHash)
}

// Transact0 is a paid mutator transaction binding the contract method 0xbe646da1.
//
// Solidity: function transact(bytes32 txHash) returns()
func (_Vm *VmTransactorSession) Transact0(txHash [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.Transact0(&_Vm.TransactOpts, txHash)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmTransactor) TryFfi(opts *bind.TransactOpts, commandInput []string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "tryFfi", commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.TryFfi(&_Vm.TransactOpts, commandInput)
}

// TryFfi is a paid mutator transaction binding the contract method 0xf45c1ce7.
//
// Solidity: function tryFfi(string[] commandInput) returns((int32,bytes,bytes) result)
func (_Vm *VmTransactorSession) TryFfi(commandInput []string) (*types.Transaction, error) {
	return _Vm.Contract.TryFfi(&_Vm.TransactOpts, commandInput)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmTransactor) TxGasPrice(opts *bind.TransactOpts, newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "txGasPrice", newGasPrice)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmSession) TxGasPrice(newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.TxGasPrice(&_Vm.TransactOpts, newGasPrice)
}

// TxGasPrice is a paid mutator transaction binding the contract method 0x48f50c0f.
//
// Solidity: function txGasPrice(uint256 newGasPrice) returns()
func (_Vm *VmTransactorSession) TxGasPrice(newGasPrice *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.TxGasPrice(&_Vm.TransactOpts, newGasPrice)
}

// WarmSlot is a paid mutator transaction binding the contract method 0xb23184cf.
//
// Solidity: function warmSlot(address target, bytes32 slot) returns()
func (_Vm *VmTransactor) WarmSlot(opts *bind.TransactOpts, target common.Address, slot [32]byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "warmSlot", target, slot)
}

// WarmSlot is a paid mutator transaction binding the contract method 0xb23184cf.
//
// Solidity: function warmSlot(address target, bytes32 slot) returns()
func (_Vm *VmSession) WarmSlot(target common.Address, slot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.WarmSlot(&_Vm.TransactOpts, target, slot)
}

// WarmSlot is a paid mutator transaction binding the contract method 0xb23184cf.
//
// Solidity: function warmSlot(address target, bytes32 slot) returns()
func (_Vm *VmTransactorSession) WarmSlot(target common.Address, slot [32]byte) (*types.Transaction, error) {
	return _Vm.Contract.WarmSlot(&_Vm.TransactOpts, target, slot)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmTransactor) Warp(opts *bind.TransactOpts, newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "warp", newTimestamp)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmSession) Warp(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Warp(&_Vm.TransactOpts, newTimestamp)
}

// Warp is a paid mutator transaction binding the contract method 0xe5d6bf02.
//
// Solidity: function warp(uint256 newTimestamp) returns()
func (_Vm *VmTransactorSession) Warp(newTimestamp *big.Int) (*types.Transaction, error) {
	return _Vm.Contract.Warp(&_Vm.TransactOpts, newTimestamp)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmTransactor) WriteFile(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeFile", path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteFile(&_Vm.TransactOpts, path, data)
}

// WriteFile is a paid mutator transaction binding the contract method 0x897e0a97.
//
// Solidity: function writeFile(string path, string data) returns()
func (_Vm *VmTransactorSession) WriteFile(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteFile(&_Vm.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmTransactor) WriteFileBinary(opts *bind.TransactOpts, path string, data []byte) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeFileBinary", path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.WriteFileBinary(&_Vm.TransactOpts, path, data)
}

// WriteFileBinary is a paid mutator transaction binding the contract method 0x1f21fc80.
//
// Solidity: function writeFileBinary(string path, bytes data) returns()
func (_Vm *VmTransactorSession) WriteFileBinary(path string, data []byte) (*types.Transaction, error) {
	return _Vm.Contract.WriteFileBinary(&_Vm.TransactOpts, path, data)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmTransactor) WriteJson(opts *bind.TransactOpts, json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeJson", json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteJson is a paid mutator transaction binding the contract method 0x35d6ad46.
//
// Solidity: function writeJson(string json, string path, string valueKey) returns()
func (_Vm *VmTransactorSession) WriteJson(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmTransactor) WriteJson0(opts *bind.TransactOpts, json string, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeJson0", json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson0(&_Vm.TransactOpts, json, path)
}

// WriteJson0 is a paid mutator transaction binding the contract method 0xe23cd19f.
//
// Solidity: function writeJson(string json, string path) returns()
func (_Vm *VmTransactorSession) WriteJson0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteJson0(&_Vm.TransactOpts, json, path)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmTransactor) WriteLine(opts *bind.TransactOpts, path string, data string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeLine", path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteLine(&_Vm.TransactOpts, path, data)
}

// WriteLine is a paid mutator transaction binding the contract method 0x619d897f.
//
// Solidity: function writeLine(string path, string data) returns()
func (_Vm *VmTransactorSession) WriteLine(path string, data string) (*types.Transaction, error) {
	return _Vm.Contract.WriteLine(&_Vm.TransactOpts, path, data)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_Vm *VmTransactor) WriteToml(opts *bind.TransactOpts, json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeToml", json, path, valueKey)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_Vm *VmSession) WriteToml(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteToml is a paid mutator transaction binding the contract method 0x51ac6a33.
//
// Solidity: function writeToml(string json, string path, string valueKey) returns()
func (_Vm *VmTransactorSession) WriteToml(json string, path string, valueKey string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml(&_Vm.TransactOpts, json, path, valueKey)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_Vm *VmTransactor) WriteToml0(opts *bind.TransactOpts, json string, path string) (*types.Transaction, error) {
	return _Vm.contract.Transact(opts, "writeToml0", json, path)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_Vm *VmSession) WriteToml0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml0(&_Vm.TransactOpts, json, path)
}

// WriteToml0 is a paid mutator transaction binding the contract method 0xc0865ba7.
//
// Solidity: function writeToml(string json, string path) returns()
func (_Vm *VmTransactorSession) WriteToml0(json string, path string) (*types.Transaction, error) {
	return _Vm.Contract.WriteToml0(&_Vm.TransactOpts, json, path)
}
