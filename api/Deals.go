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

// DealsABI is the input ABI used to generate the binding from.
const DealsABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"GetDealAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"GetDealInfo\",\"outputs\":[{\"name\":\"specHach\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"hub\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"AcceptDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"CloseDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_specHash\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"OpenDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealClosed\",\"type\":\"event\"}]"

// DealsBin is the compiled bytecode used for deploying new contracts.
const DealsBin = `0x606060405260008055341561001357600080fd5b5b61059b806100236000396000f300606060405263ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166330e8e25d81146100695780635ad5f6ae1461008e5780635be20de2146100e1578063ca6ad1bf146100f9578063cdb0017514610111575b600080fd5b341561007457600080fd5b61007c610138565b60405190815260200160405180910390f35b341561009957600080fd5b6100a460043561013f565b604051948552600160a060020a039384166020860152919092166040808501919091526060840192909252608083015260a0909101905180910390f35b34156100ec57600080fd5b6100f76004356101a7565b005b341561010457600080fd5b6100f7600435610270565b005b341561011c57600080fd5b6100f7600160a060020a036004351660243560443561033d565b005b6000545b90565b6000818152600160208190526040822080549181015460028083015460038401546004850154879687968796879691959294600160a060020a039182169491909316929160ff9091169081111561019257fe5b955095509550955095505b5091939590929450565b6000818152600160208190526040909120015433600160a060020a039081169116146101d257600080fd5b6000818152600160208190526040909120600401805460ff191682805b021790555060008181526001602081905260409182902060028101549101547f3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f1592600160a060020a039283169290911690849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b50565b60008181526001602052604090206002015433600160a060020a0390811691161461029a57600080fd5b6000818152600160208190526040909120600401805460029260ff1990911690835b021790555060008181526001602081905260409182902060028101549101547f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89092600160a060020a039283169290911690849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b50565b60008054600101908190553390849060a06040519081016040908152868252600160a060020a038085166020840152851690820152606081018590526080810160005b9052600082815260016020526040902081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560608201518160030155608082015160048201805460ff1916600183600281111561042357fe5b02179055505050600160a060020a03861660009081526002602052604080822073__IterableMapping_______________________92637e523645926c01000000000000000000000000620100006bffffffffffffffffffffffff19828b0216040491869151602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600481019390935271ffffffffffffffffffffffffffffffffffff9091166024830152604482015260640160206040518083038186803b15156104fa57600080fd5b6102c65a03f4151561050b57600080fd5b50505060405180519050507f873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be976848838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b5050505050505600a165627a7a7230582014568d5bb54e1397ca6bd45f20f52b6bb6d5e8650bf49488e2f66e777e21c0470029`

// DeployDeals deploys a new Ethereum contract, binding an instance of Deals to it.
func DeployDeals(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Deals, error) {
	parsed, err := abi.JSON(strings.NewReader(DealsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DealsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Deals{DealsCaller: DealsCaller{contract: contract}, DealsTransactor: DealsTransactor{contract: contract}}, nil
}

// Deals is an auto generated Go binding around an Ethereum contract.
type Deals struct {
	DealsCaller     // Read-only binding to the contract
	DealsTransactor // Write-only binding to the contract
}

// DealsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DealsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DealsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DealsSession struct {
	Contract     *Deals            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DealsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DealsCallerSession struct {
	Contract *DealsCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DealsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DealsTransactorSession struct {
	Contract     *DealsTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DealsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DealsRaw struct {
	Contract *Deals // Generic contract binding to access the raw methods on
}

// DealsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DealsCallerRaw struct {
	Contract *DealsCaller // Generic read-only contract binding to access the raw methods on
}

// DealsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DealsTransactorRaw struct {
	Contract *DealsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeals creates a new instance of Deals, bound to a specific deployed contract.
func NewDeals(address common.Address, backend bind.ContractBackend) (*Deals, error) {
	contract, err := bindDeals(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deals{DealsCaller: DealsCaller{contract: contract}, DealsTransactor: DealsTransactor{contract: contract}}, nil
}

// NewDealsCaller creates a new read-only instance of Deals, bound to a specific deployed contract.
func NewDealsCaller(address common.Address, caller bind.ContractCaller) (*DealsCaller, error) {
	contract, err := bindDeals(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &DealsCaller{contract: contract}, nil
}

// NewDealsTransactor creates a new write-only instance of Deals, bound to a specific deployed contract.
func NewDealsTransactor(address common.Address, transactor bind.ContractTransactor) (*DealsTransactor, error) {
	contract, err := bindDeals(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &DealsTransactor{contract: contract}, nil
}

// bindDeals binds a generic wrapper to an already deployed contract.
func bindDeals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DealsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deals *DealsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deals.Contract.DealsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deals *DealsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deals.Contract.DealsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deals *DealsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deals.Contract.DealsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deals *DealsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deals.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deals *DealsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deals.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deals *DealsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deals.Contract.contract.Transact(opts, method, params...)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(dealIndex uint256) returns()
func (_Deals *DealsTransactor) AcceptDeal(opts *bind.TransactOpts, dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "AcceptDeal", dealIndex)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(dealIndex uint256) returns()
func (_Deals *DealsSession) AcceptDeal(dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.AcceptDeal(&_Deals.TransactOpts, dealIndex)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(dealIndex uint256) returns()
func (_Deals *DealsTransactorSession) AcceptDeal(dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.AcceptDeal(&_Deals.TransactOpts, dealIndex)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(dealIndex uint256) returns()
func (_Deals *DealsTransactor) CloseDeal(opts *bind.TransactOpts, dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "CloseDeal", dealIndex)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(dealIndex uint256) returns()
func (_Deals *DealsSession) CloseDeal(dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CloseDeal(&_Deals.TransactOpts, dealIndex)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(dealIndex uint256) returns()
func (_Deals *DealsTransactorSession) CloseDeal(dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CloseDeal(&_Deals.TransactOpts, dealIndex)
}

// GetDealAmount is a paid mutator transaction binding the contract method 0x30e8e25d.
//
// Solidity: function GetDealAmount() returns(uint256)
func (_Deals *DealsTransactor) GetDealAmount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "GetDealAmount")
}

// GetDealAmount is a paid mutator transaction binding the contract method 0x30e8e25d.
//
// Solidity: function GetDealAmount() returns(uint256)
func (_Deals *DealsSession) GetDealAmount() (*types.Transaction, error) {
	return _Deals.Contract.GetDealAmount(&_Deals.TransactOpts)
}

// GetDealAmount is a paid mutator transaction binding the contract method 0x30e8e25d.
//
// Solidity: function GetDealAmount() returns(uint256)
func (_Deals *DealsTransactorSession) GetDealAmount() (*types.Transaction, error) {
	return _Deals.Contract.GetDealAmount(&_Deals.TransactOpts)
}

// GetDealInfo is a paid mutator transaction binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) returns(specHach uint256, client address, hub address, price uint256, status uint256)
func (_Deals *DealsTransactor) GetDealInfo(opts *bind.TransactOpts, dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "GetDealInfo", dealIndex)
}

// GetDealInfo is a paid mutator transaction binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) returns(specHach uint256, client address, hub address, price uint256, status uint256)
func (_Deals *DealsSession) GetDealInfo(dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.TransactOpts, dealIndex)
}

// GetDealInfo is a paid mutator transaction binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) returns(specHach uint256, client address, hub address, price uint256, status uint256)
func (_Deals *DealsTransactorSession) GetDealInfo(dealIndex *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.TransactOpts, dealIndex)
}

// OpenDeal is a paid mutator transaction binding the contract method 0xcdb00175.
//
// Solidity: function OpenDeal(_client address, _specHash uint256, _price uint256) returns()
func (_Deals *DealsTransactor) OpenDeal(opts *bind.TransactOpts, _client common.Address, _specHash *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "OpenDeal", _client, _specHash, _price)
}

// OpenDeal is a paid mutator transaction binding the contract method 0xcdb00175.
//
// Solidity: function OpenDeal(_client address, _specHash uint256, _price uint256) returns()
func (_Deals *DealsSession) OpenDeal(_client common.Address, _specHash *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _client, _specHash, _price)
}

// OpenDeal is a paid mutator transaction binding the contract method 0xcdb00175.
//
// Solidity: function OpenDeal(_client address, _specHash uint256, _price uint256) returns()
func (_Deals *DealsTransactorSession) OpenDeal(_client common.Address, _specHash *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _client, _specHash, _price)
}
