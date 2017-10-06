// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// UserABI is the input ABI used to generate the binding from.
const UserABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sum\",\"outputs\":[{\"name\":\"s\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UserBin is the compiled bytecode used for deploying new contracts.
const UserBin = `0x6060604052341561000f57600080fd5b5b6103478061001f6000396000f300606060405263ffffffff60e060020a6000350416631d834a1b811461002f578063853255cc1461005a575b600080fd5b341561003a57600080fd5b61004860043560243561007f565b60405190815260200160405180910390f35b341561006557600080fd5b610048610110565b60405190815260200160405180910390f35b600073__IterableMapping_______________________637e5236456000858560006040516020015260405160e060020a63ffffffff861602815260048101939093526024830191909152604482015260640160206040518083038186803b15156100e957600080fd5b6102c65a03f415156100fa57600080fd5b5050506040518051505060025490505b92915050565b60008060008073__IterableMapping_______________________63a21ab7166000806040516020015260405160e060020a63ffffffff8416028152600481019190915260240160206040518083038186803b151561016e57600080fd5b6102c65a03f4151561017f57600080fd5b50505060405180519350505b73__IterableMapping_______________________63c8fccc6960008560006040516020015260405160e060020a63ffffffff85160281526004810192909252602482015260440160206040518083038186803b15156101ea57600080fd5b6102c65a03f415156101fb57600080fd5b50505060405180519050156103145773__IterableMapping_______________________6375a3e8e860008560006040516040015260405160e060020a63ffffffff851602815260048101929092526024820152604401604080518083038186803b151561026857600080fd5b6102c65a03f4151561027957600080fd5b505050604051805190602001805190509150915080840193505b73__IterableMapping_______________________6388d0443760008560006040516020015260405160e060020a63ffffffff85160281526004810192909252602482015260440160206040518083038186803b15156102f257600080fd5b6102c65a03f4151561030357600080fd5b50505060405180519050925061018b565b5b505050905600a165627a7a72305820668f5c5131f3082ed57cd955e1e4fa03225f31c77119b23c440d6f410bd5cdb40029`

// DeployUser deploys a new Ethereum contract, binding an instance of User to it.
func DeployUser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *User, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}}, nil
}

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// Insert is a paid mutator transaction binding the contract method 0x1d834a1b.
//
// Solidity: function insert(k uint256, v uint256) returns(size uint256)
func (_User *UserTransactor) Insert(opts *bind.TransactOpts, k *big.Int, v *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "insert", k, v)
}

// Insert is a paid mutator transaction binding the contract method 0x1d834a1b.
//
// Solidity: function insert(k uint256, v uint256) returns(size uint256)
func (_User *UserSession) Insert(k *big.Int, v *big.Int) (*types.Transaction, error) {
	return _User.Contract.Insert(&_User.TransactOpts, k, v)
}

// Insert is a paid mutator transaction binding the contract method 0x1d834a1b.
//
// Solidity: function insert(k uint256, v uint256) returns(size uint256)
func (_User *UserTransactorSession) Insert(k *big.Int, v *big.Int) (*types.Transaction, error) {
	return _User.Contract.Insert(&_User.TransactOpts, k, v)
}

// Sum is a paid mutator transaction binding the contract method 0x853255cc.
//
// Solidity: function sum() returns(s uint256)
func (_User *UserTransactor) Sum(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "sum")
}

// Sum is a paid mutator transaction binding the contract method 0x853255cc.
//
// Solidity: function sum() returns(s uint256)
func (_User *UserSession) Sum() (*types.Transaction, error) {
	return _User.Contract.Sum(&_User.TransactOpts)
}

// Sum is a paid mutator transaction binding the contract method 0x853255cc.
//
// Solidity: function sum() returns(s uint256)
func (_User *UserTransactorSession) Sum() (*types.Transaction, error) {
	return _User.Contract.Sum(&_User.TransactOpts)
}
