// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sc

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

// IInfraredVaultUserReward is an auto generated low-level Go binding around an user-defined struct.
type IInfraredVaultUserReward struct {
	Token  common.Address
	Amount *big.Int
}

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxNumberOfRewards\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"RewardPaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"}],\"name\":\"RewardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"RewardStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardsToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newDistributor\",\"type\":\"address\"}],\"name\":\"RewardsDistributorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDuration\",\"type\":\"uint256\"}],\"name\":\"RewardsDurationUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_NUM_REWARD_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllRewardTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getAllRewardsForUser\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIInfraredVault.UserReward[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"getRewardForDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getRewardForUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"infrared\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"lastTimeRewardApplicable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"rewardsDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardsDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodFinish\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardPerTokenStored\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardResidual\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"}],\"name\":\"rewardPerToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardsVault\",\"outputs\":[{\"internalType\":\"contractIRewardVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardsToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardsDuration\",\"type\":\"uint256\"}],\"name\":\"updateRewardsDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userRewardPerTokenPaid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// MAXNUMREWARDTOKENS is a free data retrieval call binding the contract method 0xf65ae959.
//
// Solidity: function MAX_NUM_REWARD_TOKENS() view returns(uint256)
func (_Vault *VaultCaller) MAXNUMREWARDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MAX_NUM_REWARD_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXNUMREWARDTOKENS is a free data retrieval call binding the contract method 0xf65ae959.
//
// Solidity: function MAX_NUM_REWARD_TOKENS() view returns(uint256)
func (_Vault *VaultSession) MAXNUMREWARDTOKENS() (*big.Int, error) {
	return _Vault.Contract.MAXNUMREWARDTOKENS(&_Vault.CallOpts)
}

// MAXNUMREWARDTOKENS is a free data retrieval call binding the contract method 0xf65ae959.
//
// Solidity: function MAX_NUM_REWARD_TOKENS() view returns(uint256)
func (_Vault *VaultCallerSession) MAXNUMREWARDTOKENS() (*big.Int, error) {
	return _Vault.Contract.MAXNUMREWARDTOKENS(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 _balance)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 _balance)
func (_Vault *VaultSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 _balance)
func (_Vault *VaultCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address account, address _rewardsToken) view returns(uint256)
func (_Vault *VaultCaller) Earned(opts *bind.CallOpts, account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "earned", account, _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address account, address _rewardsToken) view returns(uint256)
func (_Vault *VaultSession) Earned(account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.Earned(&_Vault.CallOpts, account, _rewardsToken)
}

// Earned is a free data retrieval call binding the contract method 0x211dc32d.
//
// Solidity: function earned(address account, address _rewardsToken) view returns(uint256)
func (_Vault *VaultCallerSession) Earned(account common.Address, _rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.Earned(&_Vault.CallOpts, account, _rewardsToken)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_Vault *VaultCaller) GetAllRewardTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getAllRewardTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_Vault *VaultSession) GetAllRewardTokens() ([]common.Address, error) {
	return _Vault.Contract.GetAllRewardTokens(&_Vault.CallOpts)
}

// GetAllRewardTokens is a free data retrieval call binding the contract method 0x12edb24c.
//
// Solidity: function getAllRewardTokens() view returns(address[])
func (_Vault *VaultCallerSession) GetAllRewardTokens() ([]common.Address, error) {
	return _Vault.Contract.GetAllRewardTokens(&_Vault.CallOpts)
}

// GetAllRewardsForUser is a free data retrieval call binding the contract method 0x2a170546.
//
// Solidity: function getAllRewardsForUser(address _user) view returns((address,uint256)[])
func (_Vault *VaultCaller) GetAllRewardsForUser(opts *bind.CallOpts, _user common.Address) ([]IInfraredVaultUserReward, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getAllRewardsForUser", _user)

	if err != nil {
		return *new([]IInfraredVaultUserReward), err
	}

	out0 := *abi.ConvertType(out[0], new([]IInfraredVaultUserReward)).(*[]IInfraredVaultUserReward)

	return out0, err

}

// GetAllRewardsForUser is a free data retrieval call binding the contract method 0x2a170546.
//
// Solidity: function getAllRewardsForUser(address _user) view returns((address,uint256)[])
func (_Vault *VaultSession) GetAllRewardsForUser(_user common.Address) ([]IInfraredVaultUserReward, error) {
	return _Vault.Contract.GetAllRewardsForUser(&_Vault.CallOpts, _user)
}

// GetAllRewardsForUser is a free data retrieval call binding the contract method 0x2a170546.
//
// Solidity: function getAllRewardsForUser(address _user) view returns((address,uint256)[])
func (_Vault *VaultCallerSession) GetAllRewardsForUser(_user common.Address) ([]IInfraredVaultUserReward, error) {
	return _Vault.Contract.GetAllRewardsForUser(&_Vault.CallOpts, _user)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_Vault *VaultCaller) GetRewardForDuration(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRewardForDuration", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_Vault *VaultSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRewardForDuration(&_Vault.CallOpts, _rewardsToken)
}

// GetRewardForDuration is a free data retrieval call binding the contract method 0xbcd11014.
//
// Solidity: function getRewardForDuration(address _rewardsToken) view returns(uint256)
func (_Vault *VaultCallerSession) GetRewardForDuration(_rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRewardForDuration(&_Vault.CallOpts, _rewardsToken)
}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_Vault *VaultCaller) Infrared(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "infrared")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_Vault *VaultSession) Infrared() (common.Address, error) {
	return _Vault.Contract.Infrared(&_Vault.CallOpts)
}

// Infrared is a free data retrieval call binding the contract method 0x21f96c65.
//
// Solidity: function infrared() view returns(address)
func (_Vault *VaultCallerSession) Infrared() (common.Address, error) {
	return _Vault.Contract.Infrared(&_Vault.CallOpts)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_Vault *VaultCaller) LastTimeRewardApplicable(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "lastTimeRewardApplicable", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_Vault *VaultSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.LastTimeRewardApplicable(&_Vault.CallOpts, _rewardsToken)
}

// LastTimeRewardApplicable is a free data retrieval call binding the contract method 0x638634ee.
//
// Solidity: function lastTimeRewardApplicable(address _rewardsToken) view returns(uint256)
func (_Vault *VaultCallerSession) LastTimeRewardApplicable(_rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.LastTimeRewardApplicable(&_Vault.CallOpts, _rewardsToken)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vault *VaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vault *VaultSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vault *VaultCallerSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardResidual)
func (_Vault *VaultCaller) RewardData(opts *bind.CallOpts, arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardResidual       *big.Int
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "rewardData", arg0)

	outstruct := new(struct {
		RewardsDistributor   common.Address
		RewardsDuration      *big.Int
		PeriodFinish         *big.Int
		RewardRate           *big.Int
		LastUpdateTime       *big.Int
		RewardPerTokenStored *big.Int
		RewardResidual       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RewardsDistributor = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardsDuration = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PeriodFinish = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RewardPerTokenStored = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RewardResidual = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardResidual)
func (_Vault *VaultSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardResidual       *big.Int
}, error) {
	return _Vault.Contract.RewardData(&_Vault.CallOpts, arg0)
}

// RewardData is a free data retrieval call binding the contract method 0x48e5d9f8.
//
// Solidity: function rewardData(address ) view returns(address rewardsDistributor, uint256 rewardsDuration, uint256 periodFinish, uint256 rewardRate, uint256 lastUpdateTime, uint256 rewardPerTokenStored, uint256 rewardResidual)
func (_Vault *VaultCallerSession) RewardData(arg0 common.Address) (struct {
	RewardsDistributor   common.Address
	RewardsDuration      *big.Int
	PeriodFinish         *big.Int
	RewardRate           *big.Int
	LastUpdateTime       *big.Int
	RewardPerTokenStored *big.Int
	RewardResidual       *big.Int
}, error) {
	return _Vault.Contract.RewardData(&_Vault.CallOpts, arg0)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_Vault *VaultCaller) RewardPerToken(opts *bind.CallOpts, _rewardsToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "rewardPerToken", _rewardsToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_Vault *VaultSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.RewardPerToken(&_Vault.CallOpts, _rewardsToken)
}

// RewardPerToken is a free data retrieval call binding the contract method 0xf1229777.
//
// Solidity: function rewardPerToken(address _rewardsToken) view returns(uint256)
func (_Vault *VaultCallerSession) RewardPerToken(_rewardsToken common.Address) (*big.Int, error) {
	return _Vault.Contract.RewardPerToken(&_Vault.CallOpts, _rewardsToken)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_Vault *VaultCaller) RewardTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "rewardTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_Vault *VaultSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.RewardTokens(&_Vault.CallOpts, arg0)
}

// RewardTokens is a free data retrieval call binding the contract method 0x7bb7bed1.
//
// Solidity: function rewardTokens(uint256 ) view returns(address)
func (_Vault *VaultCallerSession) RewardTokens(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.RewardTokens(&_Vault.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_Vault *VaultCaller) Rewards(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "rewards", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_Vault *VaultSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.Rewards(&_Vault.CallOpts, arg0, arg1)
}

// Rewards is a free data retrieval call binding the contract method 0xe70b9e27.
//
// Solidity: function rewards(address , address ) view returns(uint256)
func (_Vault *VaultCallerSession) Rewards(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.Rewards(&_Vault.CallOpts, arg0, arg1)
}

// RewardsVault is a free data retrieval call binding the contract method 0x5579ed01.
//
// Solidity: function rewardsVault() view returns(address)
func (_Vault *VaultCaller) RewardsVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "rewardsVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsVault is a free data retrieval call binding the contract method 0x5579ed01.
//
// Solidity: function rewardsVault() view returns(address)
func (_Vault *VaultSession) RewardsVault() (common.Address, error) {
	return _Vault.Contract.RewardsVault(&_Vault.CallOpts)
}

// RewardsVault is a free data retrieval call binding the contract method 0x5579ed01.
//
// Solidity: function rewardsVault() view returns(address)
func (_Vault *VaultCallerSession) RewardsVault() (common.Address, error) {
	return _Vault.Contract.RewardsVault(&_Vault.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Vault *VaultCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Vault *VaultSession) StakingToken() (common.Address, error) {
	return _Vault.Contract.StakingToken(&_Vault.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Vault *VaultCallerSession) StakingToken() (common.Address, error) {
	return _Vault.Contract.StakingToken(&_Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vault *VaultCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vault *VaultSession) TotalSupply() (*big.Int, error) {
	return _Vault.Contract.TotalSupply(&_Vault.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vault *VaultCallerSession) TotalSupply() (*big.Int, error) {
	return _Vault.Contract.TotalSupply(&_Vault.CallOpts)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_Vault *VaultCaller) UserRewardPerTokenPaid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "userRewardPerTokenPaid", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_Vault *VaultSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.UserRewardPerTokenPaid(&_Vault.CallOpts, arg0, arg1)
}

// UserRewardPerTokenPaid is a free data retrieval call binding the contract method 0x7035ab98.
//
// Solidity: function userRewardPerTokenPaid(address , address ) view returns(uint256)
func (_Vault *VaultCallerSession) UserRewardPerTokenPaid(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.UserRewardPerTokenPaid(&_Vault.CallOpts, arg0, arg1)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Vault *VaultTransactor) AddReward(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "addReward", _rewardsToken, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Vault *VaultSession) AddReward(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AddReward(&_Vault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// AddReward is a paid mutator transaction binding the contract method 0x9feb8f50.
//
// Solidity: function addReward(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Vault *VaultTransactorSession) AddReward(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AddReward(&_Vault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Vault *VaultTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Vault *VaultSession) Exit() (*types.Transaction, error) {
	return _Vault.Contract.Exit(&_Vault.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Vault *VaultTransactorSession) Exit() (*types.Transaction, error) {
	return _Vault.Contract.Exit(&_Vault.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Vault *VaultTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Vault *VaultSession) GetReward() (*types.Transaction, error) {
	return _Vault.Contract.GetReward(&_Vault.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Vault *VaultTransactorSession) GetReward() (*types.Transaction, error) {
	return _Vault.Contract.GetReward(&_Vault.TransactOpts)
}

// GetRewardForUser is a paid mutator transaction binding the contract method 0xef790a82.
//
// Solidity: function getRewardForUser(address _user) returns()
func (_Vault *VaultTransactor) GetRewardForUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "getRewardForUser", _user)
}

// GetRewardForUser is a paid mutator transaction binding the contract method 0xef790a82.
//
// Solidity: function getRewardForUser(address _user) returns()
func (_Vault *VaultSession) GetRewardForUser(_user common.Address) (*types.Transaction, error) {
	return _Vault.Contract.GetRewardForUser(&_Vault.TransactOpts, _user)
}

// GetRewardForUser is a paid mutator transaction binding the contract method 0xef790a82.
//
// Solidity: function getRewardForUser(address _user) returns()
func (_Vault *VaultTransactorSession) GetRewardForUser(_user common.Address) (*types.Transaction, error) {
	return _Vault.Contract.GetRewardForUser(&_Vault.TransactOpts, _user)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardToken, uint256 _reward) returns()
func (_Vault *VaultTransactor) NotifyRewardAmount(opts *bind.TransactOpts, _rewardToken common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "notifyRewardAmount", _rewardToken, _reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardToken, uint256 _reward) returns()
func (_Vault *VaultSession) NotifyRewardAmount(_rewardToken common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.NotifyRewardAmount(&_Vault.TransactOpts, _rewardToken, _reward)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0xb66503cf.
//
// Solidity: function notifyRewardAmount(address _rewardToken, uint256 _reward) returns()
func (_Vault *VaultTransactorSession) NotifyRewardAmount(_rewardToken common.Address, _reward *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.NotifyRewardAmount(&_Vault.TransactOpts, _rewardToken, _reward)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Vault *VaultTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Vault *VaultSession) PauseStaking() (*types.Transaction, error) {
	return _Vault.Contract.PauseStaking(&_Vault.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Vault *VaultTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _Vault.Contract.PauseStaking(&_Vault.TransactOpts)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) RecoverERC20(opts *bind.TransactOpts, _to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "recoverERC20", _to, _token, _amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_Vault *VaultSession) RecoverERC20(_to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RecoverERC20(&_Vault.TransactOpts, _to, _token, _amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x1171bda9.
//
// Solidity: function recoverERC20(address _to, address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) RecoverERC20(_to common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RecoverERC20(&_Vault.TransactOpts, _to, _token, _amount)
}

// RemoveReward is a paid mutator transaction binding the contract method 0xa4d5e67c.
//
// Solidity: function removeReward(address _rewardsToken) returns()
func (_Vault *VaultTransactor) RemoveReward(opts *bind.TransactOpts, _rewardsToken common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "removeReward", _rewardsToken)
}

// RemoveReward is a paid mutator transaction binding the contract method 0xa4d5e67c.
//
// Solidity: function removeReward(address _rewardsToken) returns()
func (_Vault *VaultSession) RemoveReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveReward(&_Vault.TransactOpts, _rewardsToken)
}

// RemoveReward is a paid mutator transaction binding the contract method 0xa4d5e67c.
//
// Solidity: function removeReward(address _rewardsToken) returns()
func (_Vault *VaultTransactorSession) RemoveReward(_rewardsToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveReward(&_Vault.TransactOpts, _rewardsToken)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Vault *VaultTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Vault *VaultSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Stake(&_Vault.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Vault *VaultTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Stake(&_Vault.TransactOpts, amount)
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0x93f4bcde.
//
// Solidity: function unpauseStaking() returns()
func (_Vault *VaultTransactor) UnpauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "unpauseStaking")
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0x93f4bcde.
//
// Solidity: function unpauseStaking() returns()
func (_Vault *VaultSession) UnpauseStaking() (*types.Transaction, error) {
	return _Vault.Contract.UnpauseStaking(&_Vault.TransactOpts)
}

// UnpauseStaking is a paid mutator transaction binding the contract method 0x93f4bcde.
//
// Solidity: function unpauseStaking() returns()
func (_Vault *VaultTransactorSession) UnpauseStaking() (*types.Transaction, error) {
	return _Vault.Contract.UnpauseStaking(&_Vault.TransactOpts)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xc004ac61.
//
// Solidity: function updateRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Vault *VaultTransactor) UpdateRewardsDuration(opts *bind.TransactOpts, _rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateRewardsDuration", _rewardsToken, _rewardsDuration)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xc004ac61.
//
// Solidity: function updateRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Vault *VaultSession) UpdateRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateRewardsDuration(&_Vault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// UpdateRewardsDuration is a paid mutator transaction binding the contract method 0xc004ac61.
//
// Solidity: function updateRewardsDuration(address _rewardsToken, uint256 _rewardsDuration) returns()
func (_Vault *VaultTransactorSession) UpdateRewardsDuration(_rewardsToken common.Address, _rewardsDuration *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateRewardsDuration(&_Vault.TransactOpts, _rewardsToken, _rewardsDuration)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Vault *VaultSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Vault *VaultTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, amount)
}

// VaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vault contract.
type VaultPausedIterator struct {
	Event *VaultPaused // Event containing the contract specifics and raw log

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
func (it *VaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPaused)
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
		it.Event = new(VaultPaused)
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
func (it *VaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPaused represents a Paused event raised by the Vault contract.
type VaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vault *VaultFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultPausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultPausedIterator{contract: _Vault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vault *VaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultPaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPaused)
				if err := _Vault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Vault *VaultFilterer) ParsePaused(log types.Log) (*VaultPaused, error) {
	event := new(VaultPaused)
	if err := _Vault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the Vault contract.
type VaultRecoveredIterator struct {
	Event *VaultRecovered // Event containing the contract specifics and raw log

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
func (it *VaultRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRecovered)
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
		it.Event = new(VaultRecovered)
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
func (it *VaultRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRecovered represents a Recovered event raised by the Vault contract.
type VaultRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterRecovered(opts *bind.FilterOpts) (*VaultRecoveredIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return &VaultRecoveredIterator{contract: _Vault.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *VaultRecovered) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Recovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRecovered)
				if err := _Vault.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseRecovered(log types.Log) (*VaultRecovered, error) {
	event := new(VaultRecovered)
	if err := _Vault.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Vault contract.
type VaultRewardAddedIterator struct {
	Event *VaultRewardAdded // Event containing the contract specifics and raw log

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
func (it *VaultRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRewardAdded)
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
		it.Event = new(VaultRewardAdded)
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
func (it *VaultRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRewardAdded represents a RewardAdded event raised by the Vault contract.
type VaultRewardAdded struct {
	RewardsToken common.Address
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardsToken, uint256 reward)
func (_Vault *VaultFilterer) FilterRewardAdded(opts *bind.FilterOpts, rewardsToken []common.Address) (*VaultRewardAddedIterator, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RewardAdded", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultRewardAddedIterator{contract: _Vault.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardsToken, uint256 reward)
func (_Vault *VaultFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *VaultRewardAdded, rewardsToken []common.Address) (event.Subscription, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RewardAdded", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRewardAdded)
				if err := _Vault.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xac24935fd910bc682b5ccb1a07b718cadf8cf2f6d1404c4f3ddc3662dae40e29.
//
// Solidity: event RewardAdded(address indexed rewardsToken, uint256 reward)
func (_Vault *VaultFilterer) ParseRewardAdded(log types.Log) (*VaultRewardAdded, error) {
	event := new(VaultRewardAdded)
	if err := _Vault.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRewardPaidIterator is returned from FilterRewardPaid and is used to iterate over the raw logs and unpacked data for RewardPaid events raised by the Vault contract.
type VaultRewardPaidIterator struct {
	Event *VaultRewardPaid // Event containing the contract specifics and raw log

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
func (it *VaultRewardPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRewardPaid)
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
		it.Event = new(VaultRewardPaid)
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
func (it *VaultRewardPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRewardPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRewardPaid represents a RewardPaid event raised by the Vault contract.
type VaultRewardPaid struct {
	User         common.Address
	RewardsToken common.Address
	Reward       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardPaid is a free log retrieval operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardsToken, uint256 reward)
func (_Vault *VaultFilterer) FilterRewardPaid(opts *bind.FilterOpts, user []common.Address, rewardsToken []common.Address) (*VaultRewardPaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RewardPaid", userRule, rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultRewardPaidIterator{contract: _Vault.contract, event: "RewardPaid", logs: logs, sub: sub}, nil
}

// WatchRewardPaid is a free log subscription operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardsToken, uint256 reward)
func (_Vault *VaultFilterer) WatchRewardPaid(opts *bind.WatchOpts, sink chan<- *VaultRewardPaid, user []common.Address, rewardsToken []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RewardPaid", userRule, rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRewardPaid)
				if err := _Vault.contract.UnpackLog(event, "RewardPaid", log); err != nil {
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

// ParseRewardPaid is a log parse operation binding the contract event 0x540798df468d7b23d11f156fdb954cb19ad414d150722a7b6d55ba369dea792e.
//
// Solidity: event RewardPaid(address indexed user, address indexed rewardsToken, uint256 reward)
func (_Vault *VaultFilterer) ParseRewardPaid(log types.Log) (*VaultRewardPaid, error) {
	event := new(VaultRewardPaid)
	if err := _Vault.contract.UnpackLog(event, "RewardPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRewardRemovedIterator is returned from FilterRewardRemoved and is used to iterate over the raw logs and unpacked data for RewardRemoved events raised by the Vault contract.
type VaultRewardRemovedIterator struct {
	Event *VaultRewardRemoved // Event containing the contract specifics and raw log

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
func (it *VaultRewardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRewardRemoved)
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
		it.Event = new(VaultRewardRemoved)
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
func (it *VaultRewardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRewardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRewardRemoved represents a RewardRemoved event raised by the Vault contract.
type VaultRewardRemoved struct {
	RewardsToken common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRewardRemoved is a free log retrieval operation binding the contract event 0x755c47ac85b75fe2251607db5a480aac818b88bb535814bf1e3c4784ae4f6baa.
//
// Solidity: event RewardRemoved(address indexed rewardsToken)
func (_Vault *VaultFilterer) FilterRewardRemoved(opts *bind.FilterOpts, rewardsToken []common.Address) (*VaultRewardRemovedIterator, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RewardRemoved", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultRewardRemovedIterator{contract: _Vault.contract, event: "RewardRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardRemoved is a free log subscription operation binding the contract event 0x755c47ac85b75fe2251607db5a480aac818b88bb535814bf1e3c4784ae4f6baa.
//
// Solidity: event RewardRemoved(address indexed rewardsToken)
func (_Vault *VaultFilterer) WatchRewardRemoved(opts *bind.WatchOpts, sink chan<- *VaultRewardRemoved, rewardsToken []common.Address) (event.Subscription, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RewardRemoved", rewardsTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRewardRemoved)
				if err := _Vault.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
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

// ParseRewardRemoved is a log parse operation binding the contract event 0x755c47ac85b75fe2251607db5a480aac818b88bb535814bf1e3c4784ae4f6baa.
//
// Solidity: event RewardRemoved(address indexed rewardsToken)
func (_Vault *VaultFilterer) ParseRewardRemoved(log types.Log) (*VaultRewardRemoved, error) {
	event := new(VaultRewardRemoved)
	if err := _Vault.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRewardStoredIterator is returned from FilterRewardStored and is used to iterate over the raw logs and unpacked data for RewardStored events raised by the Vault contract.
type VaultRewardStoredIterator struct {
	Event *VaultRewardStored // Event containing the contract specifics and raw log

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
func (it *VaultRewardStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRewardStored)
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
		it.Event = new(VaultRewardStored)
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
func (it *VaultRewardStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRewardStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRewardStored represents a RewardStored event raised by the Vault contract.
type VaultRewardStored struct {
	RewardsToken    common.Address
	RewardsDuration *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRewardStored is a free log retrieval operation binding the contract event 0x7589b0732052d2ded19f37e278ed2ae0d7d2e93b21d3931b73c5200a13671653.
//
// Solidity: event RewardStored(address rewardsToken, uint256 rewardsDuration)
func (_Vault *VaultFilterer) FilterRewardStored(opts *bind.FilterOpts) (*VaultRewardStoredIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RewardStored")
	if err != nil {
		return nil, err
	}
	return &VaultRewardStoredIterator{contract: _Vault.contract, event: "RewardStored", logs: logs, sub: sub}, nil
}

// WatchRewardStored is a free log subscription operation binding the contract event 0x7589b0732052d2ded19f37e278ed2ae0d7d2e93b21d3931b73c5200a13671653.
//
// Solidity: event RewardStored(address rewardsToken, uint256 rewardsDuration)
func (_Vault *VaultFilterer) WatchRewardStored(opts *bind.WatchOpts, sink chan<- *VaultRewardStored) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RewardStored")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRewardStored)
				if err := _Vault.contract.UnpackLog(event, "RewardStored", log); err != nil {
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

// ParseRewardStored is a log parse operation binding the contract event 0x7589b0732052d2ded19f37e278ed2ae0d7d2e93b21d3931b73c5200a13671653.
//
// Solidity: event RewardStored(address rewardsToken, uint256 rewardsDuration)
func (_Vault *VaultFilterer) ParseRewardStored(log types.Log) (*VaultRewardStored, error) {
	event := new(VaultRewardStored)
	if err := _Vault.contract.UnpackLog(event, "RewardStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRewardsDistributorUpdatedIterator is returned from FilterRewardsDistributorUpdated and is used to iterate over the raw logs and unpacked data for RewardsDistributorUpdated events raised by the Vault contract.
type VaultRewardsDistributorUpdatedIterator struct {
	Event *VaultRewardsDistributorUpdated // Event containing the contract specifics and raw log

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
func (it *VaultRewardsDistributorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRewardsDistributorUpdated)
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
		it.Event = new(VaultRewardsDistributorUpdated)
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
func (it *VaultRewardsDistributorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRewardsDistributorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRewardsDistributorUpdated represents a RewardsDistributorUpdated event raised by the Vault contract.
type VaultRewardsDistributorUpdated struct {
	RewardsToken   common.Address
	NewDistributor common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRewardsDistributorUpdated is a free log retrieval operation binding the contract event 0x55bf2bef00411dcc98731a296d0f25718b09b68c42f85f0346efc0ba0b4009e4.
//
// Solidity: event RewardsDistributorUpdated(address indexed rewardsToken, address indexed newDistributor)
func (_Vault *VaultFilterer) FilterRewardsDistributorUpdated(opts *bind.FilterOpts, rewardsToken []common.Address, newDistributor []common.Address) (*VaultRewardsDistributorUpdatedIterator, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}
	var newDistributorRule []interface{}
	for _, newDistributorItem := range newDistributor {
		newDistributorRule = append(newDistributorRule, newDistributorItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RewardsDistributorUpdated", rewardsTokenRule, newDistributorRule)
	if err != nil {
		return nil, err
	}
	return &VaultRewardsDistributorUpdatedIterator{contract: _Vault.contract, event: "RewardsDistributorUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDistributorUpdated is a free log subscription operation binding the contract event 0x55bf2bef00411dcc98731a296d0f25718b09b68c42f85f0346efc0ba0b4009e4.
//
// Solidity: event RewardsDistributorUpdated(address indexed rewardsToken, address indexed newDistributor)
func (_Vault *VaultFilterer) WatchRewardsDistributorUpdated(opts *bind.WatchOpts, sink chan<- *VaultRewardsDistributorUpdated, rewardsToken []common.Address, newDistributor []common.Address) (event.Subscription, error) {

	var rewardsTokenRule []interface{}
	for _, rewardsTokenItem := range rewardsToken {
		rewardsTokenRule = append(rewardsTokenRule, rewardsTokenItem)
	}
	var newDistributorRule []interface{}
	for _, newDistributorItem := range newDistributor {
		newDistributorRule = append(newDistributorRule, newDistributorItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RewardsDistributorUpdated", rewardsTokenRule, newDistributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRewardsDistributorUpdated)
				if err := _Vault.contract.UnpackLog(event, "RewardsDistributorUpdated", log); err != nil {
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

// ParseRewardsDistributorUpdated is a log parse operation binding the contract event 0x55bf2bef00411dcc98731a296d0f25718b09b68c42f85f0346efc0ba0b4009e4.
//
// Solidity: event RewardsDistributorUpdated(address indexed rewardsToken, address indexed newDistributor)
func (_Vault *VaultFilterer) ParseRewardsDistributorUpdated(log types.Log) (*VaultRewardsDistributorUpdated, error) {
	event := new(VaultRewardsDistributorUpdated)
	if err := _Vault.contract.UnpackLog(event, "RewardsDistributorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRewardsDurationUpdatedIterator is returned from FilterRewardsDurationUpdated and is used to iterate over the raw logs and unpacked data for RewardsDurationUpdated events raised by the Vault contract.
type VaultRewardsDurationUpdatedIterator struct {
	Event *VaultRewardsDurationUpdated // Event containing the contract specifics and raw log

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
func (it *VaultRewardsDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRewardsDurationUpdated)
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
		it.Event = new(VaultRewardsDurationUpdated)
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
func (it *VaultRewardsDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRewardsDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRewardsDurationUpdated represents a RewardsDurationUpdated event raised by the Vault contract.
type VaultRewardsDurationUpdated struct {
	Token       common.Address
	NewDuration *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardsDurationUpdated is a free log retrieval operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_Vault *VaultFilterer) FilterRewardsDurationUpdated(opts *bind.FilterOpts) (*VaultRewardsDurationUpdatedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &VaultRewardsDurationUpdatedIterator{contract: _Vault.contract, event: "RewardsDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardsDurationUpdated is a free log subscription operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_Vault *VaultFilterer) WatchRewardsDurationUpdated(opts *bind.WatchOpts, sink chan<- *VaultRewardsDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "RewardsDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRewardsDurationUpdated)
				if err := _Vault.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
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

// ParseRewardsDurationUpdated is a log parse operation binding the contract event 0xad2f86b01ed93b4b3a150d448c61a4f5d8d38075d3c0c64cc0a26fd6e1f49545.
//
// Solidity: event RewardsDurationUpdated(address token, uint256 newDuration)
func (_Vault *VaultFilterer) ParseRewardsDurationUpdated(log types.Log) (*VaultRewardsDurationUpdated, error) {
	event := new(VaultRewardsDurationUpdated)
	if err := _Vault.contract.UnpackLog(event, "RewardsDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Vault contract.
type VaultStakedIterator struct {
	Event *VaultStaked // Event containing the contract specifics and raw log

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
func (it *VaultStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultStaked)
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
		it.Event = new(VaultStaked)
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
func (it *VaultStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultStaked represents a Staked event raised by the Vault contract.
type VaultStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*VaultStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultStakedIterator{contract: _Vault.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *VaultStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultStaked)
				if err := _Vault.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) ParseStaked(log types.Log) (*VaultStaked, error) {
	event := new(VaultStaked)
	if err := _Vault.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Vault contract.
type VaultUnpausedIterator struct {
	Event *VaultUnpaused // Event containing the contract specifics and raw log

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
func (it *VaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUnpaused)
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
		it.Event = new(VaultUnpaused)
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
func (it *VaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUnpaused represents a Unpaused event raised by the Vault contract.
type VaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vault *VaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultUnpausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultUnpausedIterator{contract: _Vault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vault *VaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUnpaused)
				if err := _Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Vault *VaultFilterer) ParseUnpaused(log types.Log) (*VaultUnpaused, error) {
	event := new(VaultUnpaused)
	if err := _Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Vault contract.
type VaultWithdrawnIterator struct {
	Event *VaultWithdrawn // Event containing the contract specifics and raw log

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
func (it *VaultWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWithdrawn)
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
		it.Event = new(VaultWithdrawn)
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
func (it *VaultWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWithdrawn represents a Withdrawn event raised by the Vault contract.
type VaultWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*VaultWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawnIterator{contract: _Vault.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *VaultWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWithdrawn)
				if err := _Vault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Vault *VaultFilterer) ParseWithdrawn(log types.Log) (*VaultWithdrawn, error) {
	event := new(VaultWithdrawn)
	if err := _Vault.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
