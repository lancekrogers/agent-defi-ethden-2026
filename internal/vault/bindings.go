// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// ObeyVaultMetaData contains all meta data concerning the ObeyVault contract.
var ObeyVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"asset_\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"agent_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapRouter_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uniswapFactory_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxSwapSize_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxDailyVolume_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxSlippageBps_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"TWAP_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"agent\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"asset\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToAssets\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"convertToShares\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentDay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dailyVolumeUsed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSwap\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountOutMinimum\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"amountOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"guardian\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"heldTokenAt\",\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"heldTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxDailyVolume\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxDeposit\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxMint\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSlippageBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxSwapSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewDeposit\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewMint\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewRedeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previewWithdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAgent\",\"inputs\":[{\"name\":\"newAgent\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxDailyVolume\",\"inputs\":[{\"name\":\"newMax\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxSwapSize\",\"inputs\":[{\"name\":\"newMax\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractISwapRouter02\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalAssets\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"uniswapFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AgentUpdated\",\"inputs\":[{\"name\":\"oldAgent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newAgent\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxDailyVolumeUpdated\",\"inputs\":[{\"name\":\"newMax\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxSwapSizeUpdated\",\"inputs\":[{\"name\":\"newMax\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SwapExecuted\",\"inputs\":[{\"name\":\"tokenIn\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenApprovalUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DailyVolumeExceeded\",\"inputs\":[{\"name\":\"used\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxDeposit\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxMint\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxRedeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC4626ExceededMaxWithdraw\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"assets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAgent\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyGuardian\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SameToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SlippageTooHigh\",\"inputs\":[{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SwapExceedsMaxSize\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotApproved\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ObeyVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use ObeyVaultMetaData.ABI instead.
var ObeyVaultABI = ObeyVaultMetaData.ABI

// ObeyVault is an auto generated Go binding around an Ethereum contract.
type ObeyVault struct {
	ObeyVaultCaller     // Read-only binding to the contract
	ObeyVaultTransactor // Write-only binding to the contract
	ObeyVaultFilterer   // Log filterer for contract events
}

// ObeyVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type ObeyVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObeyVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ObeyVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObeyVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ObeyVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ObeyVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ObeyVaultSession struct {
	Contract     *ObeyVault        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ObeyVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ObeyVaultCallerSession struct {
	Contract *ObeyVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ObeyVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ObeyVaultTransactorSession struct {
	Contract     *ObeyVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ObeyVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type ObeyVaultRaw struct {
	Contract *ObeyVault // Generic contract binding to access the raw methods on
}

// ObeyVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ObeyVaultCallerRaw struct {
	Contract *ObeyVaultCaller // Generic read-only contract binding to access the raw methods on
}

// ObeyVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ObeyVaultTransactorRaw struct {
	Contract *ObeyVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewObeyVault creates a new instance of ObeyVault, bound to a specific deployed contract.
func NewObeyVault(address common.Address, backend bind.ContractBackend) (*ObeyVault, error) {
	contract, err := bindObeyVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ObeyVault{ObeyVaultCaller: ObeyVaultCaller{contract: contract}, ObeyVaultTransactor: ObeyVaultTransactor{contract: contract}, ObeyVaultFilterer: ObeyVaultFilterer{contract: contract}}, nil
}

// NewObeyVaultCaller creates a new read-only instance of ObeyVault, bound to a specific deployed contract.
func NewObeyVaultCaller(address common.Address, caller bind.ContractCaller) (*ObeyVaultCaller, error) {
	contract, err := bindObeyVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultCaller{contract: contract}, nil
}

// NewObeyVaultTransactor creates a new write-only instance of ObeyVault, bound to a specific deployed contract.
func NewObeyVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*ObeyVaultTransactor, error) {
	contract, err := bindObeyVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultTransactor{contract: contract}, nil
}

// NewObeyVaultFilterer creates a new log filterer instance of ObeyVault, bound to a specific deployed contract.
func NewObeyVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*ObeyVaultFilterer, error) {
	contract, err := bindObeyVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultFilterer{contract: contract}, nil
}

// bindObeyVault binds a generic wrapper to an already deployed contract.
func bindObeyVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ObeyVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ObeyVault *ObeyVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ObeyVault.Contract.ObeyVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ObeyVault *ObeyVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ObeyVault.Contract.ObeyVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ObeyVault *ObeyVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ObeyVault.Contract.ObeyVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ObeyVault *ObeyVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ObeyVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ObeyVault *ObeyVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ObeyVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ObeyVault *ObeyVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ObeyVault.Contract.contract.Transact(opts, method, params...)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_ObeyVault *ObeyVaultCaller) TWAPPERIOD(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "TWAP_PERIOD")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_ObeyVault *ObeyVaultSession) TWAPPERIOD() (uint32, error) {
	return _ObeyVault.Contract.TWAPPERIOD(&_ObeyVault.CallOpts)
}

// TWAPPERIOD is a free data retrieval call binding the contract method 0x7ca25184.
//
// Solidity: function TWAP_PERIOD() view returns(uint32)
func (_ObeyVault *ObeyVaultCallerSession) TWAPPERIOD() (uint32, error) {
	return _ObeyVault.Contract.TWAPPERIOD(&_ObeyVault.CallOpts)
}

// Agent is a free data retrieval call binding the contract method 0xf5ff5c76.
//
// Solidity: function agent() view returns(address)
func (_ObeyVault *ObeyVaultCaller) Agent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "agent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Agent is a free data retrieval call binding the contract method 0xf5ff5c76.
//
// Solidity: function agent() view returns(address)
func (_ObeyVault *ObeyVaultSession) Agent() (common.Address, error) {
	return _ObeyVault.Contract.Agent(&_ObeyVault.CallOpts)
}

// Agent is a free data retrieval call binding the contract method 0xf5ff5c76.
//
// Solidity: function agent() view returns(address)
func (_ObeyVault *ObeyVaultCallerSession) Agent() (common.Address, error) {
	return _ObeyVault.Contract.Agent(&_ObeyVault.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.Allowance(&_ObeyVault.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.Allowance(&_ObeyVault.CallOpts, owner, spender)
}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) view returns(bool)
func (_ObeyVault *ObeyVaultCaller) ApprovedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "approvedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) view returns(bool)
func (_ObeyVault *ObeyVaultSession) ApprovedTokens(arg0 common.Address) (bool, error) {
	return _ObeyVault.Contract.ApprovedTokens(&_ObeyVault.CallOpts, arg0)
}

// ApprovedTokens is a free data retrieval call binding the contract method 0x6d1ea3fa.
//
// Solidity: function approvedTokens(address ) view returns(bool)
func (_ObeyVault *ObeyVaultCallerSession) ApprovedTokens(arg0 common.Address) (bool, error) {
	return _ObeyVault.Contract.ApprovedTokens(&_ObeyVault.CallOpts, arg0)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ObeyVault *ObeyVaultCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ObeyVault *ObeyVaultSession) Asset() (common.Address, error) {
	return _ObeyVault.Contract.Asset(&_ObeyVault.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_ObeyVault *ObeyVaultCallerSession) Asset() (common.Address, error) {
	return _ObeyVault.Contract.Asset(&_ObeyVault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.BalanceOf(&_ObeyVault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.BalanceOf(&_ObeyVault.CallOpts, account)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.ConvertToAssets(&_ObeyVault.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.ConvertToAssets(&_ObeyVault.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.ConvertToShares(&_ObeyVault.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.ConvertToShares(&_ObeyVault.CallOpts, assets)
}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) CurrentDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "currentDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) CurrentDay() (*big.Int, error) {
	return _ObeyVault.Contract.CurrentDay(&_ObeyVault.CallOpts)
}

// CurrentDay is a free data retrieval call binding the contract method 0x5c9302c9.
//
// Solidity: function currentDay() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) CurrentDay() (*big.Int, error) {
	return _ObeyVault.Contract.CurrentDay(&_ObeyVault.CallOpts)
}

// DailyVolumeUsed is a free data retrieval call binding the contract method 0xb79308d1.
//
// Solidity: function dailyVolumeUsed() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) DailyVolumeUsed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "dailyVolumeUsed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyVolumeUsed is a free data retrieval call binding the contract method 0xb79308d1.
//
// Solidity: function dailyVolumeUsed() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) DailyVolumeUsed() (*big.Int, error) {
	return _ObeyVault.Contract.DailyVolumeUsed(&_ObeyVault.CallOpts)
}

// DailyVolumeUsed is a free data retrieval call binding the contract method 0xb79308d1.
//
// Solidity: function dailyVolumeUsed() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) DailyVolumeUsed() (*big.Int, error) {
	return _ObeyVault.Contract.DailyVolumeUsed(&_ObeyVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ObeyVault *ObeyVaultCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ObeyVault *ObeyVaultSession) Decimals() (uint8, error) {
	return _ObeyVault.Contract.Decimals(&_ObeyVault.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ObeyVault *ObeyVaultCallerSession) Decimals() (uint8, error) {
	return _ObeyVault.Contract.Decimals(&_ObeyVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_ObeyVault *ObeyVaultCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_ObeyVault *ObeyVaultSession) Guardian() (common.Address, error) {
	return _ObeyVault.Contract.Guardian(&_ObeyVault.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_ObeyVault *ObeyVaultCallerSession) Guardian() (common.Address, error) {
	return _ObeyVault.Contract.Guardian(&_ObeyVault.CallOpts)
}

// HeldTokenAt is a free data retrieval call binding the contract method 0xbfc7dfa3.
//
// Solidity: function heldTokenAt(uint256 index) view returns(address)
func (_ObeyVault *ObeyVaultCaller) HeldTokenAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "heldTokenAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HeldTokenAt is a free data retrieval call binding the contract method 0xbfc7dfa3.
//
// Solidity: function heldTokenAt(uint256 index) view returns(address)
func (_ObeyVault *ObeyVaultSession) HeldTokenAt(index *big.Int) (common.Address, error) {
	return _ObeyVault.Contract.HeldTokenAt(&_ObeyVault.CallOpts, index)
}

// HeldTokenAt is a free data retrieval call binding the contract method 0xbfc7dfa3.
//
// Solidity: function heldTokenAt(uint256 index) view returns(address)
func (_ObeyVault *ObeyVaultCallerSession) HeldTokenAt(index *big.Int) (common.Address, error) {
	return _ObeyVault.Contract.HeldTokenAt(&_ObeyVault.CallOpts, index)
}

// HeldTokenCount is a free data retrieval call binding the contract method 0x7c99dd43.
//
// Solidity: function heldTokenCount() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) HeldTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "heldTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HeldTokenCount is a free data retrieval call binding the contract method 0x7c99dd43.
//
// Solidity: function heldTokenCount() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) HeldTokenCount() (*big.Int, error) {
	return _ObeyVault.Contract.HeldTokenCount(&_ObeyVault.CallOpts)
}

// HeldTokenCount is a free data retrieval call binding the contract method 0x7c99dd43.
//
// Solidity: function heldTokenCount() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) HeldTokenCount() (*big.Int, error) {
	return _ObeyVault.Contract.HeldTokenCount(&_ObeyVault.CallOpts)
}

// MaxDailyVolume is a free data retrieval call binding the contract method 0x07df2f18.
//
// Solidity: function maxDailyVolume() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) MaxDailyVolume(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "maxDailyVolume")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDailyVolume is a free data retrieval call binding the contract method 0x07df2f18.
//
// Solidity: function maxDailyVolume() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) MaxDailyVolume() (*big.Int, error) {
	return _ObeyVault.Contract.MaxDailyVolume(&_ObeyVault.CallOpts)
}

// MaxDailyVolume is a free data retrieval call binding the contract method 0x07df2f18.
//
// Solidity: function maxDailyVolume() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) MaxDailyVolume() (*big.Int, error) {
	return _ObeyVault.Contract.MaxDailyVolume(&_ObeyVault.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxDeposit(&_ObeyVault.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxDeposit(&_ObeyVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxMint(&_ObeyVault.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxMint(&_ObeyVault.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxRedeem(&_ObeyVault.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxRedeem(&_ObeyVault.CallOpts, owner)
}

// MaxSlippageBps is a free data retrieval call binding the contract method 0xc4aa7395.
//
// Solidity: function maxSlippageBps() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) MaxSlippageBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "maxSlippageBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSlippageBps is a free data retrieval call binding the contract method 0xc4aa7395.
//
// Solidity: function maxSlippageBps() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) MaxSlippageBps() (*big.Int, error) {
	return _ObeyVault.Contract.MaxSlippageBps(&_ObeyVault.CallOpts)
}

// MaxSlippageBps is a free data retrieval call binding the contract method 0xc4aa7395.
//
// Solidity: function maxSlippageBps() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) MaxSlippageBps() (*big.Int, error) {
	return _ObeyVault.Contract.MaxSlippageBps(&_ObeyVault.CallOpts)
}

// MaxSwapSize is a free data retrieval call binding the contract method 0x4f28cac2.
//
// Solidity: function maxSwapSize() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) MaxSwapSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "maxSwapSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSwapSize is a free data retrieval call binding the contract method 0x4f28cac2.
//
// Solidity: function maxSwapSize() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) MaxSwapSize() (*big.Int, error) {
	return _ObeyVault.Contract.MaxSwapSize(&_ObeyVault.CallOpts)
}

// MaxSwapSize is a free data retrieval call binding the contract method 0x4f28cac2.
//
// Solidity: function maxSwapSize() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) MaxSwapSize() (*big.Int, error) {
	return _ObeyVault.Contract.MaxSwapSize(&_ObeyVault.CallOpts)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxWithdraw(&_ObeyVault.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _ObeyVault.Contract.MaxWithdraw(&_ObeyVault.CallOpts, owner)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ObeyVault *ObeyVaultCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ObeyVault *ObeyVaultSession) Name() (string, error) {
	return _ObeyVault.Contract.Name(&_ObeyVault.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ObeyVault *ObeyVaultCallerSession) Name() (string, error) {
	return _ObeyVault.Contract.Name(&_ObeyVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ObeyVault *ObeyVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ObeyVault *ObeyVaultSession) Paused() (bool, error) {
	return _ObeyVault.Contract.Paused(&_ObeyVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ObeyVault *ObeyVaultCallerSession) Paused() (bool, error) {
	return _ObeyVault.Contract.Paused(&_ObeyVault.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewDeposit(&_ObeyVault.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewDeposit(&_ObeyVault.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewMint(&_ObeyVault.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewMint(&_ObeyVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewRedeem(&_ObeyVault.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewRedeem(&_ObeyVault.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewWithdraw(&_ObeyVault.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _ObeyVault.Contract.PreviewWithdraw(&_ObeyVault.CallOpts, assets)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_ObeyVault *ObeyVaultCaller) SwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "swapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_ObeyVault *ObeyVaultSession) SwapRouter() (common.Address, error) {
	return _ObeyVault.Contract.SwapRouter(&_ObeyVault.CallOpts)
}

// SwapRouter is a free data retrieval call binding the contract method 0xc31c9c07.
//
// Solidity: function swapRouter() view returns(address)
func (_ObeyVault *ObeyVaultCallerSession) SwapRouter() (common.Address, error) {
	return _ObeyVault.Contract.SwapRouter(&_ObeyVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ObeyVault *ObeyVaultCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ObeyVault *ObeyVaultSession) Symbol() (string, error) {
	return _ObeyVault.Contract.Symbol(&_ObeyVault.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ObeyVault *ObeyVaultCallerSession) Symbol() (string, error) {
	return _ObeyVault.Contract.Symbol(&_ObeyVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) TotalAssets() (*big.Int, error) {
	return _ObeyVault.Contract.TotalAssets(&_ObeyVault.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) TotalAssets() (*big.Int, error) {
	return _ObeyVault.Contract.TotalAssets(&_ObeyVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ObeyVault *ObeyVaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ObeyVault *ObeyVaultSession) TotalSupply() (*big.Int, error) {
	return _ObeyVault.Contract.TotalSupply(&_ObeyVault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ObeyVault *ObeyVaultCallerSession) TotalSupply() (*big.Int, error) {
	return _ObeyVault.Contract.TotalSupply(&_ObeyVault.CallOpts)
}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_ObeyVault *ObeyVaultCaller) UniswapFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ObeyVault.contract.Call(opts, &out, "uniswapFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_ObeyVault *ObeyVaultSession) UniswapFactory() (common.Address, error) {
	return _ObeyVault.Contract.UniswapFactory(&_ObeyVault.CallOpts)
}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_ObeyVault *ObeyVaultCallerSession) UniswapFactory() (common.Address, error) {
	return _ObeyVault.Contract.UniswapFactory(&_ObeyVault.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.Approve(&_ObeyVault.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.Approve(&_ObeyVault.TransactOpts, spender, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ObeyVault *ObeyVaultTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ObeyVault *ObeyVaultSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Deposit(&_ObeyVault.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256)
func (_ObeyVault *ObeyVaultTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Deposit(&_ObeyVault.TransactOpts, assets, receiver)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x232174c0.
//
// Solidity: function executeSwap(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMinimum, bytes reason) returns(uint256 amountOut)
func (_ObeyVault *ObeyVaultTransactor) ExecuteSwap(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOutMinimum *big.Int, reason []byte) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "executeSwap", tokenIn, tokenOut, amountIn, amountOutMinimum, reason)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x232174c0.
//
// Solidity: function executeSwap(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMinimum, bytes reason) returns(uint256 amountOut)
func (_ObeyVault *ObeyVaultSession) ExecuteSwap(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOutMinimum *big.Int, reason []byte) (*types.Transaction, error) {
	return _ObeyVault.Contract.ExecuteSwap(&_ObeyVault.TransactOpts, tokenIn, tokenOut, amountIn, amountOutMinimum, reason)
}

// ExecuteSwap is a paid mutator transaction binding the contract method 0x232174c0.
//
// Solidity: function executeSwap(address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOutMinimum, bytes reason) returns(uint256 amountOut)
func (_ObeyVault *ObeyVaultTransactorSession) ExecuteSwap(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int, amountOutMinimum *big.Int, reason []byte) (*types.Transaction, error) {
	return _ObeyVault.Contract.ExecuteSwap(&_ObeyVault.TransactOpts, tokenIn, tokenOut, amountIn, amountOutMinimum, reason)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ObeyVault *ObeyVaultTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ObeyVault *ObeyVaultSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Mint(&_ObeyVault.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256)
func (_ObeyVault *ObeyVaultTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Mint(&_ObeyVault.TransactOpts, shares, receiver)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ObeyVault *ObeyVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ObeyVault *ObeyVaultSession) Pause() (*types.Transaction, error) {
	return _ObeyVault.Contract.Pause(&_ObeyVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ObeyVault *ObeyVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _ObeyVault.Contract.Pause(&_ObeyVault.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ObeyVault *ObeyVaultTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ObeyVault *ObeyVaultSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Redeem(&_ObeyVault.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256)
func (_ObeyVault *ObeyVaultTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Redeem(&_ObeyVault.TransactOpts, shares, receiver, owner)
}

// SetAgent is a paid mutator transaction binding the contract method 0xbcf685ed.
//
// Solidity: function setAgent(address newAgent) returns()
func (_ObeyVault *ObeyVaultTransactor) SetAgent(opts *bind.TransactOpts, newAgent common.Address) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "setAgent", newAgent)
}

// SetAgent is a paid mutator transaction binding the contract method 0xbcf685ed.
//
// Solidity: function setAgent(address newAgent) returns()
func (_ObeyVault *ObeyVaultSession) SetAgent(newAgent common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetAgent(&_ObeyVault.TransactOpts, newAgent)
}

// SetAgent is a paid mutator transaction binding the contract method 0xbcf685ed.
//
// Solidity: function setAgent(address newAgent) returns()
func (_ObeyVault *ObeyVaultTransactorSession) SetAgent(newAgent common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetAgent(&_ObeyVault.TransactOpts, newAgent)
}

// SetApprovedToken is a paid mutator transaction binding the contract method 0x7e6a3400.
//
// Solidity: function setApprovedToken(address token, bool approved) returns()
func (_ObeyVault *ObeyVaultTransactor) SetApprovedToken(opts *bind.TransactOpts, token common.Address, approved bool) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "setApprovedToken", token, approved)
}

// SetApprovedToken is a paid mutator transaction binding the contract method 0x7e6a3400.
//
// Solidity: function setApprovedToken(address token, bool approved) returns()
func (_ObeyVault *ObeyVaultSession) SetApprovedToken(token common.Address, approved bool) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetApprovedToken(&_ObeyVault.TransactOpts, token, approved)
}

// SetApprovedToken is a paid mutator transaction binding the contract method 0x7e6a3400.
//
// Solidity: function setApprovedToken(address token, bool approved) returns()
func (_ObeyVault *ObeyVaultTransactorSession) SetApprovedToken(token common.Address, approved bool) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetApprovedToken(&_ObeyVault.TransactOpts, token, approved)
}

// SetMaxDailyVolume is a paid mutator transaction binding the contract method 0x5c3f2f4a.
//
// Solidity: function setMaxDailyVolume(uint256 newMax) returns()
func (_ObeyVault *ObeyVaultTransactor) SetMaxDailyVolume(opts *bind.TransactOpts, newMax *big.Int) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "setMaxDailyVolume", newMax)
}

// SetMaxDailyVolume is a paid mutator transaction binding the contract method 0x5c3f2f4a.
//
// Solidity: function setMaxDailyVolume(uint256 newMax) returns()
func (_ObeyVault *ObeyVaultSession) SetMaxDailyVolume(newMax *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetMaxDailyVolume(&_ObeyVault.TransactOpts, newMax)
}

// SetMaxDailyVolume is a paid mutator transaction binding the contract method 0x5c3f2f4a.
//
// Solidity: function setMaxDailyVolume(uint256 newMax) returns()
func (_ObeyVault *ObeyVaultTransactorSession) SetMaxDailyVolume(newMax *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetMaxDailyVolume(&_ObeyVault.TransactOpts, newMax)
}

// SetMaxSwapSize is a paid mutator transaction binding the contract method 0x355a395f.
//
// Solidity: function setMaxSwapSize(uint256 newMax) returns()
func (_ObeyVault *ObeyVaultTransactor) SetMaxSwapSize(opts *bind.TransactOpts, newMax *big.Int) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "setMaxSwapSize", newMax)
}

// SetMaxSwapSize is a paid mutator transaction binding the contract method 0x355a395f.
//
// Solidity: function setMaxSwapSize(uint256 newMax) returns()
func (_ObeyVault *ObeyVaultSession) SetMaxSwapSize(newMax *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetMaxSwapSize(&_ObeyVault.TransactOpts, newMax)
}

// SetMaxSwapSize is a paid mutator transaction binding the contract method 0x355a395f.
//
// Solidity: function setMaxSwapSize(uint256 newMax) returns()
func (_ObeyVault *ObeyVaultTransactorSession) SetMaxSwapSize(newMax *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.SetMaxSwapSize(&_ObeyVault.TransactOpts, newMax)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.Transfer(&_ObeyVault.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.Transfer(&_ObeyVault.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.TransferFrom(&_ObeyVault.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ObeyVault *ObeyVaultTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ObeyVault.Contract.TransferFrom(&_ObeyVault.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ObeyVault *ObeyVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ObeyVault *ObeyVaultSession) Unpause() (*types.Transaction, error) {
	return _ObeyVault.Contract.Unpause(&_ObeyVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ObeyVault *ObeyVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _ObeyVault.Contract.Unpause(&_ObeyVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ObeyVault *ObeyVaultTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ObeyVault.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ObeyVault *ObeyVaultSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Withdraw(&_ObeyVault.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256)
func (_ObeyVault *ObeyVaultTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _ObeyVault.Contract.Withdraw(&_ObeyVault.TransactOpts, assets, receiver, owner)
}

// ObeyVaultAgentUpdatedIterator is returned from FilterAgentUpdated and is used to iterate over the raw logs and unpacked data for AgentUpdated events raised by the ObeyVault contract.
type ObeyVaultAgentUpdatedIterator struct {
	Event *ObeyVaultAgentUpdated // Event containing the contract specifics and raw log

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
func (it *ObeyVaultAgentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultAgentUpdated)
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
		it.Event = new(ObeyVaultAgentUpdated)
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
func (it *ObeyVaultAgentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultAgentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultAgentUpdated represents a AgentUpdated event raised by the ObeyVault contract.
type ObeyVaultAgentUpdated struct {
	OldAgent common.Address
	NewAgent common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAgentUpdated is a free log retrieval operation binding the contract event 0xee7760cb405d3beb1072ae2e716857daa45a5dd52bc1383a21173199d32c7dee.
//
// Solidity: event AgentUpdated(address indexed oldAgent, address indexed newAgent)
func (_ObeyVault *ObeyVaultFilterer) FilterAgentUpdated(opts *bind.FilterOpts, oldAgent []common.Address, newAgent []common.Address) (*ObeyVaultAgentUpdatedIterator, error) {

	var oldAgentRule []interface{}
	for _, oldAgentItem := range oldAgent {
		oldAgentRule = append(oldAgentRule, oldAgentItem)
	}
	var newAgentRule []interface{}
	for _, newAgentItem := range newAgent {
		newAgentRule = append(newAgentRule, newAgentItem)
	}

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "AgentUpdated", oldAgentRule, newAgentRule)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultAgentUpdatedIterator{contract: _ObeyVault.contract, event: "AgentUpdated", logs: logs, sub: sub}, nil
}

// WatchAgentUpdated is a free log subscription operation binding the contract event 0xee7760cb405d3beb1072ae2e716857daa45a5dd52bc1383a21173199d32c7dee.
//
// Solidity: event AgentUpdated(address indexed oldAgent, address indexed newAgent)
func (_ObeyVault *ObeyVaultFilterer) WatchAgentUpdated(opts *bind.WatchOpts, sink chan<- *ObeyVaultAgentUpdated, oldAgent []common.Address, newAgent []common.Address) (event.Subscription, error) {

	var oldAgentRule []interface{}
	for _, oldAgentItem := range oldAgent {
		oldAgentRule = append(oldAgentRule, oldAgentItem)
	}
	var newAgentRule []interface{}
	for _, newAgentItem := range newAgent {
		newAgentRule = append(newAgentRule, newAgentItem)
	}

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "AgentUpdated", oldAgentRule, newAgentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultAgentUpdated)
				if err := _ObeyVault.contract.UnpackLog(event, "AgentUpdated", log); err != nil {
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

// ParseAgentUpdated is a log parse operation binding the contract event 0xee7760cb405d3beb1072ae2e716857daa45a5dd52bc1383a21173199d32c7dee.
//
// Solidity: event AgentUpdated(address indexed oldAgent, address indexed newAgent)
func (_ObeyVault *ObeyVaultFilterer) ParseAgentUpdated(log types.Log) (*ObeyVaultAgentUpdated, error) {
	event := new(ObeyVaultAgentUpdated)
	if err := _ObeyVault.contract.UnpackLog(event, "AgentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ObeyVault contract.
type ObeyVaultApprovalIterator struct {
	Event *ObeyVaultApproval // Event containing the contract specifics and raw log

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
func (it *ObeyVaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultApproval)
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
		it.Event = new(ObeyVaultApproval)
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
func (it *ObeyVaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultApproval represents a Approval event raised by the ObeyVault contract.
type ObeyVaultApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ObeyVault *ObeyVaultFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ObeyVaultApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultApprovalIterator{contract: _ObeyVault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ObeyVault *ObeyVaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ObeyVaultApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultApproval)
				if err := _ObeyVault.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ObeyVault *ObeyVaultFilterer) ParseApproval(log types.Log) (*ObeyVaultApproval, error) {
	event := new(ObeyVaultApproval)
	if err := _ObeyVault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ObeyVault contract.
type ObeyVaultDepositIterator struct {
	Event *ObeyVaultDeposit // Event containing the contract specifics and raw log

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
func (it *ObeyVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultDeposit)
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
		it.Event = new(ObeyVaultDeposit)
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
func (it *ObeyVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultDeposit represents a Deposit event raised by the ObeyVault contract.
type ObeyVaultDeposit struct {
	Sender common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ObeyVault *ObeyVaultFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address, owner []common.Address) (*ObeyVaultDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultDepositIterator{contract: _ObeyVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ObeyVault *ObeyVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ObeyVaultDeposit, sender []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "Deposit", senderRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultDeposit)
				if err := _ObeyVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed sender, address indexed owner, uint256 assets, uint256 shares)
func (_ObeyVault *ObeyVaultFilterer) ParseDeposit(log types.Log) (*ObeyVaultDeposit, error) {
	event := new(ObeyVaultDeposit)
	if err := _ObeyVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultMaxDailyVolumeUpdatedIterator is returned from FilterMaxDailyVolumeUpdated and is used to iterate over the raw logs and unpacked data for MaxDailyVolumeUpdated events raised by the ObeyVault contract.
type ObeyVaultMaxDailyVolumeUpdatedIterator struct {
	Event *ObeyVaultMaxDailyVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *ObeyVaultMaxDailyVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultMaxDailyVolumeUpdated)
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
		it.Event = new(ObeyVaultMaxDailyVolumeUpdated)
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
func (it *ObeyVaultMaxDailyVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultMaxDailyVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultMaxDailyVolumeUpdated represents a MaxDailyVolumeUpdated event raised by the ObeyVault contract.
type ObeyVaultMaxDailyVolumeUpdated struct {
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxDailyVolumeUpdated is a free log retrieval operation binding the contract event 0xa1f53dad4024b4e77813595cc395a33825582866232b16fdd6998f2a2bc9367b.
//
// Solidity: event MaxDailyVolumeUpdated(uint256 newMax)
func (_ObeyVault *ObeyVaultFilterer) FilterMaxDailyVolumeUpdated(opts *bind.FilterOpts) (*ObeyVaultMaxDailyVolumeUpdatedIterator, error) {

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "MaxDailyVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &ObeyVaultMaxDailyVolumeUpdatedIterator{contract: _ObeyVault.contract, event: "MaxDailyVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxDailyVolumeUpdated is a free log subscription operation binding the contract event 0xa1f53dad4024b4e77813595cc395a33825582866232b16fdd6998f2a2bc9367b.
//
// Solidity: event MaxDailyVolumeUpdated(uint256 newMax)
func (_ObeyVault *ObeyVaultFilterer) WatchMaxDailyVolumeUpdated(opts *bind.WatchOpts, sink chan<- *ObeyVaultMaxDailyVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "MaxDailyVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultMaxDailyVolumeUpdated)
				if err := _ObeyVault.contract.UnpackLog(event, "MaxDailyVolumeUpdated", log); err != nil {
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

// ParseMaxDailyVolumeUpdated is a log parse operation binding the contract event 0xa1f53dad4024b4e77813595cc395a33825582866232b16fdd6998f2a2bc9367b.
//
// Solidity: event MaxDailyVolumeUpdated(uint256 newMax)
func (_ObeyVault *ObeyVaultFilterer) ParseMaxDailyVolumeUpdated(log types.Log) (*ObeyVaultMaxDailyVolumeUpdated, error) {
	event := new(ObeyVaultMaxDailyVolumeUpdated)
	if err := _ObeyVault.contract.UnpackLog(event, "MaxDailyVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultMaxSwapSizeUpdatedIterator is returned from FilterMaxSwapSizeUpdated and is used to iterate over the raw logs and unpacked data for MaxSwapSizeUpdated events raised by the ObeyVault contract.
type ObeyVaultMaxSwapSizeUpdatedIterator struct {
	Event *ObeyVaultMaxSwapSizeUpdated // Event containing the contract specifics and raw log

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
func (it *ObeyVaultMaxSwapSizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultMaxSwapSizeUpdated)
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
		it.Event = new(ObeyVaultMaxSwapSizeUpdated)
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
func (it *ObeyVaultMaxSwapSizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultMaxSwapSizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultMaxSwapSizeUpdated represents a MaxSwapSizeUpdated event raised by the ObeyVault contract.
type ObeyVaultMaxSwapSizeUpdated struct {
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxSwapSizeUpdated is a free log retrieval operation binding the contract event 0x60cdec16253805a936970e4fccab47d8699308d353824c8d78944f3913b57284.
//
// Solidity: event MaxSwapSizeUpdated(uint256 newMax)
func (_ObeyVault *ObeyVaultFilterer) FilterMaxSwapSizeUpdated(opts *bind.FilterOpts) (*ObeyVaultMaxSwapSizeUpdatedIterator, error) {

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "MaxSwapSizeUpdated")
	if err != nil {
		return nil, err
	}
	return &ObeyVaultMaxSwapSizeUpdatedIterator{contract: _ObeyVault.contract, event: "MaxSwapSizeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSwapSizeUpdated is a free log subscription operation binding the contract event 0x60cdec16253805a936970e4fccab47d8699308d353824c8d78944f3913b57284.
//
// Solidity: event MaxSwapSizeUpdated(uint256 newMax)
func (_ObeyVault *ObeyVaultFilterer) WatchMaxSwapSizeUpdated(opts *bind.WatchOpts, sink chan<- *ObeyVaultMaxSwapSizeUpdated) (event.Subscription, error) {

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "MaxSwapSizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultMaxSwapSizeUpdated)
				if err := _ObeyVault.contract.UnpackLog(event, "MaxSwapSizeUpdated", log); err != nil {
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

// ParseMaxSwapSizeUpdated is a log parse operation binding the contract event 0x60cdec16253805a936970e4fccab47d8699308d353824c8d78944f3913b57284.
//
// Solidity: event MaxSwapSizeUpdated(uint256 newMax)
func (_ObeyVault *ObeyVaultFilterer) ParseMaxSwapSizeUpdated(log types.Log) (*ObeyVaultMaxSwapSizeUpdated, error) {
	event := new(ObeyVaultMaxSwapSizeUpdated)
	if err := _ObeyVault.contract.UnpackLog(event, "MaxSwapSizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ObeyVault contract.
type ObeyVaultPausedIterator struct {
	Event *ObeyVaultPaused // Event containing the contract specifics and raw log

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
func (it *ObeyVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultPaused)
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
		it.Event = new(ObeyVaultPaused)
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
func (it *ObeyVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultPaused represents a Paused event raised by the ObeyVault contract.
type ObeyVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ObeyVault *ObeyVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*ObeyVaultPausedIterator, error) {

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ObeyVaultPausedIterator{contract: _ObeyVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ObeyVault *ObeyVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ObeyVaultPaused) (event.Subscription, error) {

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultPaused)
				if err := _ObeyVault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ObeyVault *ObeyVaultFilterer) ParsePaused(log types.Log) (*ObeyVaultPaused, error) {
	event := new(ObeyVaultPaused)
	if err := _ObeyVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultSwapExecutedIterator is returned from FilterSwapExecuted and is used to iterate over the raw logs and unpacked data for SwapExecuted events raised by the ObeyVault contract.
type ObeyVaultSwapExecutedIterator struct {
	Event *ObeyVaultSwapExecuted // Event containing the contract specifics and raw log

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
func (it *ObeyVaultSwapExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultSwapExecuted)
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
		it.Event = new(ObeyVaultSwapExecuted)
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
func (it *ObeyVaultSwapExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultSwapExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultSwapExecuted represents a SwapExecuted event raised by the ObeyVault contract.
type ObeyVaultSwapExecuted struct {
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Reason    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwapExecuted is a free log retrieval operation binding the contract event 0xf5be4cde7f6a949be6fa81a42a771e5c6a3e8d9e630a233b6fcf3fb1691ce8ac.
//
// Solidity: event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut, bytes reason)
func (_ObeyVault *ObeyVaultFilterer) FilterSwapExecuted(opts *bind.FilterOpts, tokenIn []common.Address, tokenOut []common.Address) (*ObeyVaultSwapExecutedIterator, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "SwapExecuted", tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultSwapExecutedIterator{contract: _ObeyVault.contract, event: "SwapExecuted", logs: logs, sub: sub}, nil
}

// WatchSwapExecuted is a free log subscription operation binding the contract event 0xf5be4cde7f6a949be6fa81a42a771e5c6a3e8d9e630a233b6fcf3fb1691ce8ac.
//
// Solidity: event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut, bytes reason)
func (_ObeyVault *ObeyVaultFilterer) WatchSwapExecuted(opts *bind.WatchOpts, sink chan<- *ObeyVaultSwapExecuted, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "SwapExecuted", tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultSwapExecuted)
				if err := _ObeyVault.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
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

// ParseSwapExecuted is a log parse operation binding the contract event 0xf5be4cde7f6a949be6fa81a42a771e5c6a3e8d9e630a233b6fcf3fb1691ce8ac.
//
// Solidity: event SwapExecuted(address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut, bytes reason)
func (_ObeyVault *ObeyVaultFilterer) ParseSwapExecuted(log types.Log) (*ObeyVaultSwapExecuted, error) {
	event := new(ObeyVaultSwapExecuted)
	if err := _ObeyVault.contract.UnpackLog(event, "SwapExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultTokenApprovalUpdatedIterator is returned from FilterTokenApprovalUpdated and is used to iterate over the raw logs and unpacked data for TokenApprovalUpdated events raised by the ObeyVault contract.
type ObeyVaultTokenApprovalUpdatedIterator struct {
	Event *ObeyVaultTokenApprovalUpdated // Event containing the contract specifics and raw log

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
func (it *ObeyVaultTokenApprovalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultTokenApprovalUpdated)
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
		it.Event = new(ObeyVaultTokenApprovalUpdated)
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
func (it *ObeyVaultTokenApprovalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultTokenApprovalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultTokenApprovalUpdated represents a TokenApprovalUpdated event raised by the ObeyVault contract.
type ObeyVaultTokenApprovalUpdated struct {
	Token    common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenApprovalUpdated is a free log retrieval operation binding the contract event 0xacddef0bc85f1f88cae59ecb80b161e03a04bb438d9aa0e1f8cd5a34a1c48ca9.
//
// Solidity: event TokenApprovalUpdated(address indexed token, bool approved)
func (_ObeyVault *ObeyVaultFilterer) FilterTokenApprovalUpdated(opts *bind.FilterOpts, token []common.Address) (*ObeyVaultTokenApprovalUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "TokenApprovalUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultTokenApprovalUpdatedIterator{contract: _ObeyVault.contract, event: "TokenApprovalUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenApprovalUpdated is a free log subscription operation binding the contract event 0xacddef0bc85f1f88cae59ecb80b161e03a04bb438d9aa0e1f8cd5a34a1c48ca9.
//
// Solidity: event TokenApprovalUpdated(address indexed token, bool approved)
func (_ObeyVault *ObeyVaultFilterer) WatchTokenApprovalUpdated(opts *bind.WatchOpts, sink chan<- *ObeyVaultTokenApprovalUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "TokenApprovalUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultTokenApprovalUpdated)
				if err := _ObeyVault.contract.UnpackLog(event, "TokenApprovalUpdated", log); err != nil {
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

// ParseTokenApprovalUpdated is a log parse operation binding the contract event 0xacddef0bc85f1f88cae59ecb80b161e03a04bb438d9aa0e1f8cd5a34a1c48ca9.
//
// Solidity: event TokenApprovalUpdated(address indexed token, bool approved)
func (_ObeyVault *ObeyVaultFilterer) ParseTokenApprovalUpdated(log types.Log) (*ObeyVaultTokenApprovalUpdated, error) {
	event := new(ObeyVaultTokenApprovalUpdated)
	if err := _ObeyVault.contract.UnpackLog(event, "TokenApprovalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ObeyVault contract.
type ObeyVaultTransferIterator struct {
	Event *ObeyVaultTransfer // Event containing the contract specifics and raw log

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
func (it *ObeyVaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultTransfer)
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
		it.Event = new(ObeyVaultTransfer)
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
func (it *ObeyVaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultTransfer represents a Transfer event raised by the ObeyVault contract.
type ObeyVaultTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ObeyVault *ObeyVaultFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ObeyVaultTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultTransferIterator{contract: _ObeyVault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ObeyVault *ObeyVaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ObeyVaultTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultTransfer)
				if err := _ObeyVault.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ObeyVault *ObeyVaultFilterer) ParseTransfer(log types.Log) (*ObeyVaultTransfer, error) {
	event := new(ObeyVaultTransfer)
	if err := _ObeyVault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ObeyVault contract.
type ObeyVaultUnpausedIterator struct {
	Event *ObeyVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *ObeyVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultUnpaused)
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
		it.Event = new(ObeyVaultUnpaused)
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
func (it *ObeyVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultUnpaused represents a Unpaused event raised by the ObeyVault contract.
type ObeyVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ObeyVault *ObeyVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ObeyVaultUnpausedIterator, error) {

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ObeyVaultUnpausedIterator{contract: _ObeyVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ObeyVault *ObeyVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ObeyVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultUnpaused)
				if err := _ObeyVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ObeyVault *ObeyVaultFilterer) ParseUnpaused(log types.Log) (*ObeyVaultUnpaused, error) {
	event := new(ObeyVaultUnpaused)
	if err := _ObeyVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ObeyVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ObeyVault contract.
type ObeyVaultWithdrawIterator struct {
	Event *ObeyVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *ObeyVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ObeyVaultWithdraw)
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
		it.Event = new(ObeyVaultWithdraw)
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
func (it *ObeyVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ObeyVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ObeyVaultWithdraw represents a Withdraw event raised by the ObeyVault contract.
type ObeyVaultWithdraw struct {
	Sender   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ObeyVault *ObeyVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address, owner []common.Address) (*ObeyVaultWithdrawIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ObeyVault.contract.FilterLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ObeyVaultWithdrawIterator{contract: _ObeyVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ObeyVault *ObeyVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ObeyVaultWithdraw, sender []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ObeyVault.contract.WatchLogs(opts, "Withdraw", senderRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ObeyVaultWithdraw)
				if err := _ObeyVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed sender, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_ObeyVault *ObeyVaultFilterer) ParseWithdraw(log types.Log) (*ObeyVaultWithdraw, error) {
	event := new(ObeyVaultWithdraw)
	if err := _ObeyVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
