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
const DealsABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"GetDealAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"GetDealInfo\",\"outputs\":[{\"name\":\"specHach\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"hub\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"AcceptDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"IsPriceAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"CloseDeal\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"},{\"name\":\"_specHash\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"OpenDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealClosed\",\"type\":\"event\"}]"

// DealsBin is the compiled bytecode used for deploying new contracts.
const DealsBin = `0x6060604052600a6000556000600255341561001957600080fd5b60405160208061070b833981016040528080519150505b60018054600160a060020a031916600160a060020a0383161790555b505b6106ae8061005d6000396000f3006060604052361561005c5763ffffffff60e060020a60003504166330e8e25d81146100615780635ad5f6ae146100865780635be20de2146100d95780638027b2dd146100f1578063ca6ad1bf14610127578063cdb0017514610134575b600080fd5b341561006c57600080fd5b61007461015b565b60405190815260200160405180910390f35b341561009157600080fd5b61009c600435610162565b604051948552600160a060020a039384166020860152919092166040808501919091526060840192909252608083015260a0909101905180910390f35b34156100e457600080fd5b6100ef6004356101c7565b005b34156100fc57600080fd5b610113600160a060020a03600435166024356102d9565b604051901515815260200160405180910390f35b6100ef600435610367565b005b341561013f57600080fd5b6100ef600160a060020a0360043516602435604435610525565b005b6002545b90565b6000818152600360208190526040822080546001820154600280840154948401546004850154879687968796879691959094600160a060020a0392831694939092169260ff16908111156101b257fe5b955095509550955095505b5091939590929450565b60008181526003602052604090206002015433600160a060020a039081169116146101f157600080fd5b600081815260036020819052604082206002810154925491015461022e92600160a060020a0316919081151561022357fe5b0462030d40026102d9565b151561023957600080fd5b600081815260036020526040902060040180546001919060ff191682805b02179055506000818152600360205260409081902060028101546001909101547f3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f1592600160a060020a039283169290911690849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b50565b60015460009081908390600160a060020a031663dd62ed3e8630856040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561034057600080fd5b6102c65a03f1151561035157600080fd5b5050506040518051905003101590505b92915050565b60008181526003602052604081206004015460ff16905b81600281111561038a57fe5b14156103bf5760008281526003602052604090206002015433600160a060020a039081169116146103ba57600080fd5b610482565b60015b8160028111156103ce57fe5b1415610482576001805460008481526003602081905260408083209485015460028601549590920154600160a060020a03948516956323b872dd95938416949316929091516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561046657600080fd5b6102c65a03f1151561047757600080fd5b505050604051805150505b5b600082815260036020526040902060040180546002919060ff19166001835b02179055506000828152600360205260409081902060028101546001909101547f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89092600160a060020a039283169290911690859051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b5050565b8233600061053382856102d9565b151561053e57600080fd5b50600280546001019081905560a06040519081016040908152868252600160a060020a038085166020840152851690820152606081018590526080810160005b9052600082815260036020526040902081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560608201518160030155608082015160048201805460ff1916600183600281111561062157fe5b02179055509050507f873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be976848838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b5050505050505600a165627a7a72305820a40eb36c36dd1bedd003a3182a75aee989f5f2affec2c823492810ea11fb56ef0029`

// DeployDeals deploys a new Ethereum contract, binding an instance of Deals to it.
func DeployDeals(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Deals, error) {
	parsed, err := abi.JSON(strings.NewReader(DealsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DealsBin), backend, _token)
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

// IsPriceAllowed is a paid mutator transaction binding the contract method 0x8027b2dd.
//
// Solidity: function IsPriceAllowed(addr address, _price uint256) returns(bool)
func (_Deals *DealsTransactor) IsPriceAllowed(opts *bind.TransactOpts, addr common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "IsPriceAllowed", addr, _price)
}

// IsPriceAllowed is a paid mutator transaction binding the contract method 0x8027b2dd.
//
// Solidity: function IsPriceAllowed(addr address, _price uint256) returns(bool)
func (_Deals *DealsSession) IsPriceAllowed(addr common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.IsPriceAllowed(&_Deals.TransactOpts, addr, _price)
}

// IsPriceAllowed is a paid mutator transaction binding the contract method 0x8027b2dd.
//
// Solidity: function IsPriceAllowed(addr address, _price uint256) returns(bool)
func (_Deals *DealsTransactorSession) IsPriceAllowed(addr common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.IsPriceAllowed(&_Deals.TransactOpts, addr, _price)
}

// OpenDeal is a paid mutator transaction binding the contract method 0xcdb00175.
//
// Solidity: function OpenDeal(_hub address, _specHash uint256, _price uint256) returns()
func (_Deals *DealsTransactor) OpenDeal(opts *bind.TransactOpts, _hub common.Address, _specHash *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "OpenDeal", _hub, _specHash, _price)
}

// OpenDeal is a paid mutator transaction binding the contract method 0xcdb00175.
//
// Solidity: function OpenDeal(_hub address, _specHash uint256, _price uint256) returns()
func (_Deals *DealsSession) OpenDeal(_hub common.Address, _specHash *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _hub, _specHash, _price)
}

// OpenDeal is a paid mutator transaction binding the contract method 0xcdb00175.
//
// Solidity: function OpenDeal(_hub address, _specHash uint256, _price uint256) returns()
func (_Deals *DealsTransactorSession) OpenDeal(_hub common.Address, _specHash *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _hub, _specHash, _price)
}
