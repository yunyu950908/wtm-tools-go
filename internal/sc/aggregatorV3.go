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

// AggregatorV3MetaData contains all meta data concerning the AggregatorV3 contract.
var AggregatorV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AggregatorV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorV3MetaData.ABI instead.
var AggregatorV3ABI = AggregatorV3MetaData.ABI

// AggregatorV3 is an auto generated Go binding around an Ethereum contract.
type AggregatorV3 struct {
	AggregatorV3Caller     // Read-only binding to the contract
	AggregatorV3Transactor // Write-only binding to the contract
	AggregatorV3Filterer   // Log filterer for contract events
}

// AggregatorV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorV3Session struct {
	Contract     *AggregatorV3     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatorV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorV3CallerSession struct {
	Contract *AggregatorV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AggregatorV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorV3TransactorSession struct {
	Contract     *AggregatorV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AggregatorV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorV3Raw struct {
	Contract *AggregatorV3 // Generic contract binding to access the raw methods on
}

// AggregatorV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorV3CallerRaw struct {
	Contract *AggregatorV3Caller // Generic read-only contract binding to access the raw methods on
}

// AggregatorV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorV3TransactorRaw struct {
	Contract *AggregatorV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregatorV3 creates a new instance of AggregatorV3, bound to a specific deployed contract.
func NewAggregatorV3(address common.Address, backend bind.ContractBackend) (*AggregatorV3, error) {
	contract, err := bindAggregatorV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3{AggregatorV3Caller: AggregatorV3Caller{contract: contract}, AggregatorV3Transactor: AggregatorV3Transactor{contract: contract}, AggregatorV3Filterer: AggregatorV3Filterer{contract: contract}}, nil
}

// NewAggregatorV3Caller creates a new read-only instance of AggregatorV3, bound to a specific deployed contract.
func NewAggregatorV3Caller(address common.Address, caller bind.ContractCaller) (*AggregatorV3Caller, error) {
	contract, err := bindAggregatorV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3Caller{contract: contract}, nil
}

// NewAggregatorV3Transactor creates a new write-only instance of AggregatorV3, bound to a specific deployed contract.
func NewAggregatorV3Transactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorV3Transactor, error) {
	contract, err := bindAggregatorV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3Transactor{contract: contract}, nil
}

// NewAggregatorV3Filterer creates a new log filterer instance of AggregatorV3, bound to a specific deployed contract.
func NewAggregatorV3Filterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorV3Filterer, error) {
	contract, err := bindAggregatorV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorV3Filterer{contract: contract}, nil
}

// bindAggregatorV3 binds a generic wrapper to an already deployed contract.
func bindAggregatorV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregatorV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorV3 *AggregatorV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV3.Contract.AggregatorV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorV3 *AggregatorV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV3.Contract.AggregatorV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorV3 *AggregatorV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV3.Contract.AggregatorV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregatorV3 *AggregatorV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregatorV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregatorV3 *AggregatorV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregatorV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregatorV3 *AggregatorV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregatorV3.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AggregatorV3 *AggregatorV3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AggregatorV3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AggregatorV3 *AggregatorV3Session) Decimals() (uint8, error) {
	return _AggregatorV3.Contract.Decimals(&_AggregatorV3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AggregatorV3 *AggregatorV3CallerSession) Decimals() (uint8, error) {
	return _AggregatorV3.Contract.Decimals(&_AggregatorV3.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3 *AggregatorV3Caller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _AggregatorV3.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3 *AggregatorV3Session) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3.Contract.LatestRoundData(&_AggregatorV3.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_AggregatorV3 *AggregatorV3CallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _AggregatorV3.Contract.LatestRoundData(&_AggregatorV3.CallOpts)
}
