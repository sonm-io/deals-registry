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

// MigrationsABI is the input ABI used to generate the binding from.
const MigrationsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"new_address\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"last_completed_migration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"completed\",\"type\":\"uint256\"}],\"name\":\"setCompleted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MigrationsBin is the compiled bytecode used for deploying new contracts.
const MigrationsBin = `0x6060604052341561000f57600080fd5b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506102db8061005e6000396000f300606060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630900f01014610067578063445df0ac146100a05780638da5cb5b146100c9578063fdacd5761461011e575b600080fd5b341561007257600080fd5b61009e600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610141565b005b34156100ab57600080fd5b6100b3610224565b6040518082815260200191505060405180910390f35b34156100d457600080fd5b6100dc61022a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561012957600080fd5b61013f600480803590602001909190505061024f565b005b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415610220578190508073ffffffffffffffffffffffffffffffffffffffff1663fdacd5766001546040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b151561020b57600080fd5b6102c65a03f1151561021c57600080fd5b5050505b5050565b60015481565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614156102ac57806001819055505b505600a165627a7a72305820f356aff50493f383e16a1b32c2e6edf54ed2a3cffa4180d5abfdd6c9234492ce0029`

// DeployMigrations deploys a new Ethereum contract, binding an instance of Migrations to it.
func DeployMigrations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Migrations, error) {
	parsed, err := abi.JSON(strings.NewReader(MigrationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MigrationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}, MigrationsFilterer: MigrationsFilterer{contract: contract}}, nil
}

// Migrations is an auto generated Go binding around an Ethereum contract.
type Migrations struct {
	MigrationsCaller     // Read-only binding to the contract
	MigrationsTransactor // Write-only binding to the contract
	MigrationsFilterer   // Log filterer for contract events
}

// MigrationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MigrationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MigrationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MigrationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MigrationsSession struct {
	Contract     *Migrations       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MigrationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MigrationsCallerSession struct {
	Contract *MigrationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MigrationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MigrationsTransactorSession struct {
	Contract     *MigrationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MigrationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MigrationsRaw struct {
	Contract *Migrations // Generic contract binding to access the raw methods on
}

// MigrationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MigrationsCallerRaw struct {
	Contract *MigrationsCaller // Generic read-only contract binding to access the raw methods on
}

// MigrationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MigrationsTransactorRaw struct {
	Contract *MigrationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMigrations creates a new instance of Migrations, bound to a specific deployed contract.
func NewMigrations(address common.Address, backend bind.ContractBackend) (*Migrations, error) {
	contract, err := bindMigrations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}, MigrationsFilterer: MigrationsFilterer{contract: contract}}, nil
}

// NewMigrationsCaller creates a new read-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsCaller(address common.Address, caller bind.ContractCaller) (*MigrationsCaller, error) {
	contract, err := bindMigrations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsCaller{contract: contract}, nil
}

// NewMigrationsTransactor creates a new write-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsTransactor(address common.Address, transactor bind.ContractTransactor) (*MigrationsTransactor, error) {
	contract, err := bindMigrations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsTransactor{contract: contract}, nil
}

// NewMigrationsFilterer creates a new log filterer instance of Migrations, bound to a specific deployed contract.
func NewMigrationsFilterer(address common.Address, filterer bind.ContractFilterer) (*MigrationsFilterer, error) {
	contract, err := bindMigrations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MigrationsFilterer{contract: contract}, nil
}

// bindMigrations binds a generic wrapper to an already deployed contract.
func bindMigrations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MigrationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.MigrationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transact(opts, method, params...)
}

// Last_completed_migration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() constant returns(uint256)
func (_Migrations *MigrationsCaller) Last_completed_migration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Migrations.contract.Call(opts, out, "last_completed_migration")
	return *ret0, err
}

// Last_completed_migration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() constant returns(uint256)
func (_Migrations *MigrationsSession) Last_completed_migration() (*big.Int, error) {
	return _Migrations.Contract.Last_completed_migration(&_Migrations.CallOpts)
}

// Last_completed_migration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() constant returns(uint256)
func (_Migrations *MigrationsCallerSession) Last_completed_migration() (*big.Int, error) {
	return _Migrations.Contract.Last_completed_migration(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Migrations *MigrationsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Migrations.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Migrations *MigrationsSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Migrations *MigrationsCallerSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(completed uint256) returns()
func (_Migrations *MigrationsTransactor) SetCompleted(opts *bind.TransactOpts, completed *big.Int) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "setCompleted", completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(completed uint256) returns()
func (_Migrations *MigrationsSession) SetCompleted(completed *big.Int) (*types.Transaction, error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(completed uint256) returns()
func (_Migrations *MigrationsTransactorSession) SetCompleted(completed *big.Int) (*types.Transaction, error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(new_address address) returns()
func (_Migrations *MigrationsTransactor) Upgrade(opts *bind.TransactOpts, new_address common.Address) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "upgrade", new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(new_address address) returns()
func (_Migrations *MigrationsSession) Upgrade(new_address common.Address) (*types.Transaction, error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(new_address address) returns()
func (_Migrations *MigrationsTransactorSession) Upgrade(new_address common.Address) (*types.Transaction, error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}
