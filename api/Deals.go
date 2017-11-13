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
const DealsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetDeals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"},{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_specHash\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_workTime\",\"type\":\"uint256\"}],\"name\":\"OpenDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"GetDealInfo\",\"outputs\":[{\"name\":\"specHach\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"hub\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"workTime\",\"type\":\"uint256\"},{\"name\":\"endTIme\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AcceptDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CloseDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetDealsAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealClosed\",\"type\":\"event\"}]"

// DealsBin is the compiled bytecode used for deploying new contracts.
const DealsBin = `0x60606040526000600155341561001457600080fd5b604051602080610d42833981016040528080519150505b60008054600160a060020a031916600160a060020a0383161790555b505b610cea806100586000396000f3006060604052361561005c5763ffffffff60e060020a600035041663087ca1d881146100615780633b24475f146100d45780635ad5f6ae146101045780635be20de21461016d578063ca6ad1bf14610185578063e45ea8d31461019d575b600080fd5b341561006c57600080fd5b610080600160a060020a03600435166101c2565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100c05780820151818401525b6020016100a7565b505050509050019250505060405180910390f35b34156100df57600080fd5b610102600160a060020a0360043581169060243516604435606435608435610246565b005b341561010f57600080fd5b61011a6004356104f8565b604051978852600160a060020a039687166020890152949095166040808801919091526060870193909352608086019190915260a085015260c084019290925260e0830152610100909101905180910390f35b341561017857600080fd5b61010260043561057d565b005b341561019057600080fd5b61010260043561068b565b005b34156101a857600080fd5b6101b0610c29565b60405190815260200160405180910390f35b6101ca610c61565b6004600083600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561023957602002820191906000526020600020905b815481526020019060010190808311610225575b505050505090505b919050565b83600160a060020a031633600160a060020a031614151561026657600080fd5b60018054810190556101006040519081016040528084815260200185600160a060020a0316815260200186600160a060020a031681526020018381526020016000815260200182815260200160008152602001600060028111156102c657fe5b9052600154600090815260026020526040902081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015160078201805460ff1916600183600281111561038a57fe5b02179055505060008054600160a060020a031691506323b872dd908690309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561040157600080fd5b6102c65a03f1151561041257600080fd5b50505060405180519050151561042757600080fd5b600154600090815260036020526040902054610449908363ffffffff610c3016565b60018054600090815260036020908152604080832094909455600160a060020a038816825260049052919091208054909181016104868382610c73565b916000526020600020900160005b600154909190915055507f873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be9768488585600154604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b5050505050565b600081815260026020819052604082208054600182015482840154600384015460048501546005860154600687015460078801548a998a998a998a998a998a998a9994989397600160a060020a039384169793909216959094939260ff169081111561056057fe5b985098509850985098509850985098505b50919395975091939597565b6000818152600260208190526040909120015433600160a060020a039081169116146105a857600080fd5b60005b60008281526002602081905260409091206007015460ff16908111156105cd57fe5b146105d757600080fd5b600081815260026020526040902060070180546001919060ff191682805b021790555060008181526002602081905260409182902042600482018190556005820154016006820155908101546001909101547f3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f1592600160a060020a039283169290911690849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b50565b60008181526002602052604081206001015433600160a060020a039081169116146106b557600080fd5b60015b60008381526002602081905260409091206007015460ff16908111156106da57fe5b1415610a71576000828152600260205260409020600601544211156107dc576000805483825260026020819052604080842091820154600390920154600160a060020a039384169463a9059cbb94909316929091516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561077657600080fd5b6102c65a03f1151561078757600080fd5b50505060405180519050151561079c57600080fd5b60008281526002602090815260408083206003908101549252909120546107c89163ffffffff610c4a16565b6000838152600360205260409020556109cf565b600082815260026020526040902060058101546003909101548115156107fe57fe5b6000848152600260208190526040808320600481015484549190930154959094044292909203919091029450600160a060020a039283169363a9059cbb9316918591516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561088857600080fd5b6102c65a03f1151561089957600080fd5b5050506040518051905015156108ae57600080fd5b6000828152600360205260409020546108cd908263ffffffff610c4a16565b600083815260036020818152604080842094909455825460029091528383206001810154920154600160a060020a039182169463a9059cbb9493909216929086900391516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561095857600080fd5b6102c65a03f1151561096957600080fd5b50505060405180519050151561097e57600080fd5b60008281526002602090815260408083206003908101549252909120546109ad9183900363ffffffff610c4a16565b6000838152600360209081526040808320939093556002905220426006909101555b6000828152600260208190526040909120600701805460ff19166001835b0217905550600082815260026020819052604091829020908101546001909101547f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89092600160a060020a039283169290911690859051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1610c23565b60005b60008381526002602081905260409091206007015460ff1690811115610a9657fe5b141561005c5760008054838252600260205260408083206001810154600390910154600160a060020a039384169463a9059cbb9490921692909190516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610b1957600080fd5b6102c65a03f11515610b2a57600080fd5b505050604051805190501515610b3f57600080fd5b6000828152600260209081526040808320600390810154925290912054610b6b9163ffffffff610c4a16565b600083815260036020908152604080832093909355600290819052919020600701805460ff19166001836109ed565b0217905550600082815260026020819052604091829020908101546001909101547f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89092600160a060020a039283169290911690859051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1610c23565b600080fd5b5b5b5050565b6001545b90565b600082820183811015610c3f57fe5b8091505b5092915050565b600082821115610c5657fe5b508082035b92915050565b60206040519081016040526000815290565b815481835581811511610c9757600083815260209020610c97918101908301610c9d565b5b505050565b610c2d91905b80821115610cb75760008155600101610ca3565b5090565b905600a165627a7a7230582090496f9a68f83050fd2dee06eea9efa96ab60143e775129c90e664651533388f0029`

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

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, startTime uint256, workTime uint256, endTIme uint256, status uint256)
func (_Deals *DealsCaller) GetDealInfo(opts *bind.CallOpts, dealIndex *big.Int) (struct {
	SpecHach  *big.Int
	Client    common.Address
	Hub       common.Address
	Price     *big.Int
	StartTime *big.Int
	WorkTime  *big.Int
	EndTIme   *big.Int
	Status    *big.Int
}, error) {
	ret := new(struct {
		SpecHach  *big.Int
		Client    common.Address
		Hub       common.Address
		Price     *big.Int
		StartTime *big.Int
		WorkTime  *big.Int
		EndTIme   *big.Int
		Status    *big.Int
	})
	out := ret
	err := _Deals.contract.Call(opts, out, "GetDealInfo", dealIndex)
	return *ret, err
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, startTime uint256, workTime uint256, endTIme uint256, status uint256)
func (_Deals *DealsSession) GetDealInfo(dealIndex *big.Int) (struct {
	SpecHach  *big.Int
	Client    common.Address
	Hub       common.Address
	Price     *big.Int
	StartTime *big.Int
	WorkTime  *big.Int
	EndTIme   *big.Int
	Status    *big.Int
}, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.CallOpts, dealIndex)
}

// GetDealInfo is a free data retrieval call binding the contract method 0x5ad5f6ae.
//
// Solidity: function GetDealInfo(dealIndex uint256) constant returns(specHach uint256, client address, hub address, price uint256, startTime uint256, workTime uint256, endTIme uint256, status uint256)
func (_Deals *DealsCallerSession) GetDealInfo(dealIndex *big.Int) (struct {
	SpecHach  *big.Int
	Client    common.Address
	Hub       common.Address
	Price     *big.Int
	StartTime *big.Int
	WorkTime  *big.Int
	EndTIme   *big.Int
	Status    *big.Int
}, error) {
	return _Deals.Contract.GetDealInfo(&_Deals.CallOpts, dealIndex)
}

// GetDeals is a free data retrieval call binding the contract method 0x087ca1d8.
//
// Solidity: function GetDeals(addr address) constant returns(uint256[])
func (_Deals *DealsCaller) GetDeals(opts *bind.CallOpts, addr common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDeals", addr)
	return *ret0, err
}

// GetDeals is a free data retrieval call binding the contract method 0x087ca1d8.
//
// Solidity: function GetDeals(addr address) constant returns(uint256[])
func (_Deals *DealsSession) GetDeals(addr common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDeals(&_Deals.CallOpts, addr)
}

// GetDeals is a free data retrieval call binding the contract method 0x087ca1d8.
//
// Solidity: function GetDeals(addr address) constant returns(uint256[])
func (_Deals *DealsCallerSession) GetDeals(addr common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDeals(&_Deals.CallOpts, addr)
}

// GetDealsAmount is a free data retrieval call binding the contract method 0xe45ea8d3.
//
// Solidity: function GetDealsAmount() constant returns(uint256)
func (_Deals *DealsCaller) GetDealsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDealsAmount")
	return *ret0, err
}

// GetDealsAmount is a free data retrieval call binding the contract method 0xe45ea8d3.
//
// Solidity: function GetDealsAmount() constant returns(uint256)
func (_Deals *DealsSession) GetDealsAmount() (*big.Int, error) {
	return _Deals.Contract.GetDealsAmount(&_Deals.CallOpts)
}

// GetDealsAmount is a free data retrieval call binding the contract method 0xe45ea8d3.
//
// Solidity: function GetDealsAmount() constant returns(uint256)
func (_Deals *DealsCallerSession) GetDealsAmount() (*big.Int, error) {
	return _Deals.Contract.GetDealsAmount(&_Deals.CallOpts)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(id uint256) returns()
func (_Deals *DealsTransactor) AcceptDeal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "AcceptDeal", id)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(id uint256) returns()
func (_Deals *DealsSession) AcceptDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.AcceptDeal(&_Deals.TransactOpts, id)
}

// AcceptDeal is a paid mutator transaction binding the contract method 0x5be20de2.
//
// Solidity: function AcceptDeal(id uint256) returns()
func (_Deals *DealsTransactorSession) AcceptDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.AcceptDeal(&_Deals.TransactOpts, id)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(id uint256) returns()
func (_Deals *DealsTransactor) CloseDeal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "CloseDeal", id)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(id uint256) returns()
func (_Deals *DealsSession) CloseDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CloseDeal(&_Deals.TransactOpts, id)
}

// CloseDeal is a paid mutator transaction binding the contract method 0xca6ad1bf.
//
// Solidity: function CloseDeal(id uint256) returns()
func (_Deals *DealsTransactorSession) CloseDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CloseDeal(&_Deals.TransactOpts, id)
}

// OpenDeal is a paid mutator transaction binding the contract method 0x3b24475f.
//
// Solidity: function OpenDeal(_hub address, _client address, _specHash uint256, _price uint256, _workTime uint256) returns()
func (_Deals *DealsTransactor) OpenDeal(opts *bind.TransactOpts, _hub common.Address, _client common.Address, _specHash *big.Int, _price *big.Int, _workTime *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "OpenDeal", _hub, _client, _specHash, _price, _workTime)
}

// OpenDeal is a paid mutator transaction binding the contract method 0x3b24475f.
//
// Solidity: function OpenDeal(_hub address, _client address, _specHash uint256, _price uint256, _workTime uint256) returns()
func (_Deals *DealsSession) OpenDeal(_hub common.Address, _client common.Address, _specHash *big.Int, _price *big.Int, _workTime *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _hub, _client, _specHash, _price, _workTime)
}

// OpenDeal is a paid mutator transaction binding the contract method 0x3b24475f.
//
// Solidity: function OpenDeal(_hub address, _client address, _specHash uint256, _price uint256, _workTime uint256) returns()
func (_Deals *DealsTransactorSession) OpenDeal(_hub common.Address, _client common.Address, _specHash *big.Int, _price *big.Int, _workTime *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.OpenDeal(&_Deals.TransactOpts, _hub, _client, _specHash, _price, _workTime)
}
