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
const DealsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"GetDealAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"GetDealInfo\",\"outputs\":[{\"name\":\"specHach\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"hub\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"AcceptDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"}],\"name\":\"GetDealByClient\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"IsPriceAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"}],\"name\":\"GetDealByHubAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"CloseDeal\",\"outputs\":[],\"payable\":true,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"},{\"name\":\"_specHash\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"OpenDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealClosed\",\"type\":\"event\"}]"

// DealsBin is the compiled bytecode used for deploying new contracts.
const DealsBin = `0x6060604052600a6000556000600255341561001957600080fd5b604051602080610b40833981016040528080519150505b60018054600160a060020a031916600160a060020a0383161790555b505b610ae38061005d6000396000f300606060405236156100725763ffffffff60e060020a60003504166330e8e25d81146100775780635ad5f6ae1461009c5780635be20de2146100ef57806368d93ff1146101075780638027b2dd1461017a578063a60d4dff146101b0578063ca6ad1bf14610223578063cdb0017514610230575b600080fd5b341561008257600080fd5b61008a610257565b60405190815260200160405180910390f35b34156100a757600080fd5b6100b260043561025e565b604051948552600160a060020a039384166020860152919092166040808501919091526060840192909252608083015260a0909101905180910390f35b34156100fa57600080fd5b6101056004356102c3565b005b341561011257600080fd5b610126600160a060020a03600435166103d0565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101665780820151818401525b60200161014d565b505050509050019250505060405180910390f35b341561018557600080fd5b61019c600160a060020a0360043516602435610454565b604051901515815260200160405180910390f35b34156101bb57600080fd5b610126600160a060020a03600435166104e2565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101665780820151818401525b60200161014d565b505050509050019250505060405180910390f35b610105600435610566565b005b341561023b57600080fd5b610105600160a060020a036004351660243560443561077e565b005b6002545b90565b6000818152600360208190526040822080546001820154600280840154948401546004850154879687968796879691959094600160a060020a0392831694939092169260ff16908111156102ae57fe5b955095509550955095505b5091939590929450565b60008181526003602052604090206002015433600160a060020a039081169116146102ed57600080fd5b600081815260036020819052604082206002810154925491015461032592600160a060020a0316919081151561031f57fe5b04610454565b151561033057600080fd5b600081815260036020526040902060040180546001919060ff191682805b02179055506000818152600360205260409081902060028101546001909101547f3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f1592600160a060020a039283169290911690849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b50565b6103d8610a5a565b6006600083600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561044757602002820191906000526020600020905b815481526020019060010190808311610433575b505050505090505b919050565b60015460009081908390600160a060020a031663dd62ed3e8630856040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b15156104bb57600080fd5b6102c65a03f115156104cc57600080fd5b5050506040518051905003101590505b92915050565b6104ea610a5a565b6005600083600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561044757602002820191906000526020600020905b815481526020019060010190808311610433575b505050505090505b919050565b60008181526003602052604081206004015460ff16905b81600281111561058957fe5b14156105be5760008281526003602052604090206002015433600160a060020a039081169116146105b957600080fd5b6106db565b60015b8160028111156105cd57fe5b14156106db5760015460008381526003602081905260408083206002810154920154600160a060020a039485169463a9059cbb949316929091516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561064e57600080fd5b6102c65a03f1151561065f57600080fd5b50505060405180519050151561067457600080fd5b600082815260036020818152604080842092830154600190930154600160a060020a031684526004909152909120546106b29163ffffffff610a2916565b600083815260036020908152604080832060010154600160a060020a0316835260049091529020555b5b600082815260036020526040902060040180546002919060ff19166001835b02179055506000828152600360205260409081902060028101546001909101547f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89092600160a060020a039283169290911690859051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b5050565b8233600061078c8285610454565b151561079757600080fd5b600154600160a060020a03166323b872dd83308760006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561080357600080fd5b6102c65a03f1151561081457600080fd5b50505060405180519050151561082957600080fd5b600160a060020a038216600090815260046020526040902054610852908563ffffffff610a4016565b600160a060020a03831660009081526004602052604090819020919091556002805460010190819055915060a090519081016040908152868252600160a060020a038085166020840152851690820152606081018590526080810160005b9052600082815260036020526040902081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560608201518160030155608082015160048201805460ff1916600183600281111561095357fe5b02179055505050600160a060020a03821660009081526006602052604090208054600181016109828382610a6c565b916000526020600020900160005b5082905550600160a060020a03831660009081526005602052604090208054600181016109bd8382610a6c565b916000526020600020900160005b50829055507f873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be976848838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b505050505050565b600082821115610a3557fe5b508082035b92915050565b600082820183811015610a4f57fe5b8091505b5092915050565b60206040519081016040526000815290565b815481835581811511610a9057600083815260209020610a90918101908301610a96565b5b505050565b61025b91905b80821115610ab05760008155600101610a9c565b5090565b905600a165627a7a72305820d7991be4b6701d591fd655dbc1be7484f32cd6bf915a9037a6fe2166160a431c0029`

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

// GetDealAmount is a free data retrieval call binding the contract method 0x30e8e25d.
//
// Solidity: function GetDealAmount() constant returns(uint256)
func (_Deals *DealsCaller) GetDealAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDealAmount")
	return *ret0, err
}

// GetDealAmount is a free data retrieval call binding the contract method 0x30e8e25d.
//
// Solidity: function GetDealAmount() constant returns(uint256)
func (_Deals *DealsSession) GetDealAmount() (*big.Int, error) {
	return _Deals.Contract.GetDealAmount(&_Deals.CallOpts)
}

// GetDealAmount is a free data retrieval call binding the contract method 0x30e8e25d.
//
// Solidity: function GetDealAmount() constant returns(uint256)
func (_Deals *DealsCallerSession) GetDealAmount() (*big.Int, error) {
	return _Deals.Contract.GetDealAmount(&_Deals.CallOpts)
}

// GetDealByClient is a free data retrieval call binding the contract method 0x68d93ff1.
//
// Solidity: function GetDealByClient(_client address) constant returns(uint256[])
func (_Deals *DealsCaller) GetDealByClient(opts *bind.CallOpts, _client common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDealByClient", _client)
	return *ret0, err
}

// GetDealByClient is a free data retrieval call binding the contract method 0x68d93ff1.
//
// Solidity: function GetDealByClient(_client address) constant returns(uint256[])
func (_Deals *DealsSession) GetDealByClient(_client common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealByClient(&_Deals.CallOpts, _client)
}

// GetDealByClient is a free data retrieval call binding the contract method 0x68d93ff1.
//
// Solidity: function GetDealByClient(_client address) constant returns(uint256[])
func (_Deals *DealsCallerSession) GetDealByClient(_client common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealByClient(&_Deals.CallOpts, _client)
}

// GetDealByHubAddress is a free data retrieval call binding the contract method 0xa60d4dff.
//
// Solidity: function GetDealByHubAddress(_hub address) constant returns(uint256[])
func (_Deals *DealsCaller) GetDealByHubAddress(opts *bind.CallOpts, _hub common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDealByHubAddress", _hub)
	return *ret0, err
}

// GetDealByHubAddress is a free data retrieval call binding the contract method 0xa60d4dff.
//
// Solidity: function GetDealByHubAddress(_hub address) constant returns(uint256[])
func (_Deals *DealsSession) GetDealByHubAddress(_hub common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealByHubAddress(&_Deals.CallOpts, _hub)
}

// GetDealByHubAddress is a free data retrieval call binding the contract method 0xa60d4dff.
//
// Solidity: function GetDealByHubAddress(_hub address) constant returns(uint256[])
func (_Deals *DealsCallerSession) GetDealByHubAddress(_hub common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealByHubAddress(&_Deals.CallOpts, _hub)
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, status uint256)
func (_Deals *DealsCaller) GetDealInfo(opts *bind.CallOpts, dealIndex *big.Int) (struct {
	SpecHach *big.Int
	Client   common.Address
	Hub      common.Address
	Price    *big.Int
	Status   *big.Int
}, error) {
	ret := new(struct {
		SpecHach *big.Int
		Client   common.Address
		Hub      common.Address
		Price    *big.Int
		Status   *big.Int
	})
	out := ret
	err := _Deals.contract.Call(opts, out, "GetDealInfo", dealIndex)
	return *ret, err
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, status uint256)
func (_Deals *DealsSession) GetDealInfo(dealIndex *big.Int) (struct {
	SpecHach *big.Int
	Client   common.Address
	Hub      common.Address
	Price    *big.Int
	Status   *big.Int
}, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.CallOpts, dealIndex)
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, status uint256)
func (_Deals *DealsCallerSession) GetDealInfo(dealIndex *big.Int) (struct {
	SpecHach *big.Int
	Client   common.Address
	Hub      common.Address
	Price    *big.Int
	Status   *big.Int
}, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.CallOpts, dealIndex)
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
