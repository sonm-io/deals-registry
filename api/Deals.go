// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// DealsABI is the input ABI used to generate the binding from.
const DealsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"GetDeals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"},{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_specHash\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_workTime\",\"type\":\"uint256\"}],\"name\":\"OpenDeal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"GetDealInfo\",\"outputs\":[{\"name\":\"specHach\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"hub\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"workTime\",\"type\":\"uint256\"},{\"name\":\"endTIme\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AcceptDeal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CloseDeal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetDealsAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DealOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DealAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DealClosed\",\"type\":\"event\"}]"

// DealsBin is the compiled bytecode used for deploying new contracts.
const DealsBin = `0x60606040526000600155341561001457600080fd5b60405160208061151583398101604052808051906020019091905050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050611495806100806000396000f300606060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063087ca1d81461007d5780633b24475f1461010b5780635ad5f6ae1461017e5780635be20de21461023e578063ca6ad1bf14610261578063e45ea8d314610284575b600080fd5b341561008857600080fd5b6100b4600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506102ad565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156100f75780820151818401526020810190506100dc565b505050509050019250505060405180910390f35b341561011657600080fd5b61017c600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190803590602001909190803590602001909190505061034a565b005b341561018957600080fd5b61019f60048080359060200190919050506107ac565b604051808981526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018481526020018381526020018281526020019850505050505050505060405180910390f35b341561024957600080fd5b61025f6004808035906020019091905050610853565b005b341561026c57600080fd5b6102826004808035906020019091905050610a1b565b005b341561028f57600080fd5b6102976113c3565b6040518082815260200191505060405180910390f35b6102b5611404565b600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561033e57602002820191906000526020600020905b81548152602001906001019080831161032a575b50505050509050919050565b8373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561038457600080fd5b6001805401600181905550610100604051908101604052808481526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001600081526020018281526020016000815260200160018152506002600060015481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701559050506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8530856000604051602001526040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b15156105db57600080fd5b6102c65a03f115156105ec57600080fd5b50505060405180519050151561060157600080fd5b61062982600360006001548152602001908152602001600020546113cd90919063ffffffff16565b60036000600154815260200190815260200160002081905550600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060010182816106939190611418565b91600052602060002090016000600154909190915055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff1614151561074857600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060010182816107309190611418565b91600052602060002090016000600154909190915055505b6001548473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be97684860405160405180910390a45050505050565b6000806000806000806000806000600260008b8152602001908152602001600020905080600001548160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168260020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836003015484600401548560050154866006015487600701549850985098509850985098509850985050919395975091939597565b6002600082815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108c357600080fd5b600160026000838152602001908152602001600020600701541415156108e857600080fd5b600280600083815260200190815260200160002060070181905550426002600083815260200190815260200160002060040181905550600260008281526020019081526020016000206005015442016002600083815260200190815260200160002060060181905550806002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f1560405160405180910390a450565b600060028060008481526020019081526020016000206007015414156110b1576002600083815260200190815260200160002060060154421115610bef576000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6002600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660026000868152602001908152602001600020600301546000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1515610b7157600080fd5b6102c65a03f11515610b8257600080fd5b505050604051805190501515610b9757600080fd5b610bd3600260008481526020019081526020016000206003015460036000858152602001908152602001600020546113eb90919063ffffffff16565b6003600084815260200190815260200160002081905550610fc9565b6002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c5f57600080fd5b60026000838152602001908152602001600020600501546002600084815260200190815260200160002060030154811515610c9657fe5b04600260008481526020019081526020016000206004015442030290506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6002600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1515610db557600080fd5b6102c65a03f11515610dc657600080fd5b505050604051805190501515610ddb57600080fd5b610e018160036000858152602001908152602001600020546113eb90919063ffffffff16565b60036000848152602001908152602001600020819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6002600085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836002600087815260200190815260200160002060030154036000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1515610f3257600080fd5b6102c65a03f11515610f4357600080fd5b505050604051805190501515610f5857600080fd5b610f968160026000858152602001908152602001600020600301540360036000858152602001908152602001600020546113eb90919063ffffffff16565b60036000848152602001908152602001600020819055504260026000848152602001908152602001600020600601819055505b60036002600084815260200190815260200160002060070181905550816002600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166002600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89060405160405180910390a46113bf565b6001600260008481526020019081526020016000206007015414156113b9576002600083815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561114057600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6002600085815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660026000868152602001908152602001600020600301546000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b151561125857600080fd5b6102c65a03f1151561126957600080fd5b50505060405180519050151561127e57600080fd5b6112ba600260008481526020019081526020016000206003015460036000858152602001908152602001600020546113eb90919063ffffffff16565b600360008481526020019081526020016000208190555060036002600084815260200190815260200160002060070181905550816002600084815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166002600085815260200190815260200160002060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89060405160405180910390a46113be565b600080fd5b5b5050565b6000600154905090565b60008082840190508381101515156113e157fe5b8091505092915050565b60008282111515156113f957fe5b818303905092915050565b602060405190810160405280600081525090565b81548183558181151161143f5781836000526020600020918201910161143e9190611444565b5b505050565b61146691905b8082111561146257600081600090555060010161144a565b5090565b905600a165627a7a72305820e8ccb6c34aa5702e11d3ed40f711ab8402eff8df03bbdb427971212a8e32fefe0029`

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
	return address, tx, &Deals{DealsCaller: DealsCaller{contract: contract}, DealsTransactor: DealsTransactor{contract: contract}, DealsFilterer: DealsFilterer{contract: contract}}, nil
}

// Deals is an auto generated Go binding around an Ethereum contract.
type Deals struct {
	DealsCaller     // Read-only binding to the contract
	DealsTransactor // Write-only binding to the contract
	DealsFilterer   // Log filterer for contract events
}

// DealsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DealsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DealsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DealsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DealsFilterer struct {
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
	contract, err := bindDeals(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deals{DealsCaller: DealsCaller{contract: contract}, DealsTransactor: DealsTransactor{contract: contract}, DealsFilterer: DealsFilterer{contract: contract}}, nil
}

// NewDealsCaller creates a new read-only instance of Deals, bound to a specific deployed contract.
func NewDealsCaller(address common.Address, caller bind.ContractCaller) (*DealsCaller, error) {
	contract, err := bindDeals(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DealsCaller{contract: contract}, nil
}

// NewDealsTransactor creates a new write-only instance of Deals, bound to a specific deployed contract.
func NewDealsTransactor(address common.Address, transactor bind.ContractTransactor) (*DealsTransactor, error) {
	contract, err := bindDeals(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DealsTransactor{contract: contract}, nil
}

// NewDealsFilterer creates a new log filterer instance of Deals, bound to a specific deployed contract.
func NewDealsFilterer(address common.Address, filterer bind.ContractFilterer) (*DealsFilterer, error) {
	contract, err := bindDeals(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DealsFilterer{contract: contract}, nil
}

// bindDeals binds a generic wrapper to an already deployed contract.
func bindDeals(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DealsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// DealsDealAcceptedIterator is returned from FilterDealAccepted and is used to iterate over the raw logs and unpacked data for DealAccepted events raised by the Deals contract.
type DealsDealAcceptedIterator struct {
	Event *DealsDealAccepted // Event containing the contract specifics and raw log

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
func (it *DealsDealAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DealsDealAccepted)
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
		it.Event = new(DealsDealAccepted)
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
func (it *DealsDealAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DealsDealAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DealsDealAccepted represents a DealAccepted event raised by the Deals contract.
type DealsDealAccepted struct {
	Hub    common.Address
	Client common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealAccepted is a free log retrieval operation binding the contract event 0x3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f15.
//
// Solidity: event DealAccepted(hub indexed address, client indexed address, id indexed uint256)
func (_Deals *DealsFilterer) FilterDealAccepted(opts *bind.FilterOpts, hub []common.Address, client []common.Address, id []*big.Int) (*DealsDealAcceptedIterator, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Deals.contract.FilterLogs(opts, "DealAccepted", hubRule, clientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &DealsDealAcceptedIterator{contract: _Deals.contract, event: "DealAccepted", logs: logs, sub: sub}, nil
}

// WatchDealAccepted is a free log subscription operation binding the contract event 0x3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f15.
//
// Solidity: event DealAccepted(hub indexed address, client indexed address, id indexed uint256)
func (_Deals *DealsFilterer) WatchDealAccepted(opts *bind.WatchOpts, sink chan<- *DealsDealAccepted, hub []common.Address, client []common.Address, id []*big.Int) (event.Subscription, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Deals.contract.WatchLogs(opts, "DealAccepted", hubRule, clientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DealsDealAccepted)
				if err := _Deals.contract.UnpackLog(event, "DealAccepted", log); err != nil {
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

// DealsDealClosedIterator is returned from FilterDealClosed and is used to iterate over the raw logs and unpacked data for DealClosed events raised by the Deals contract.
type DealsDealClosedIterator struct {
	Event *DealsDealClosed // Event containing the contract specifics and raw log

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
func (it *DealsDealClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DealsDealClosed)
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
		it.Event = new(DealsDealClosed)
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
func (it *DealsDealClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DealsDealClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DealsDealClosed represents a DealClosed event raised by the Deals contract.
type DealsDealClosed struct {
	Hub    common.Address
	Client common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealClosed is a free log retrieval operation binding the contract event 0x72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb890.
//
// Solidity: event DealClosed(hub indexed address, client indexed address, id indexed uint256)
func (_Deals *DealsFilterer) FilterDealClosed(opts *bind.FilterOpts, hub []common.Address, client []common.Address, id []*big.Int) (*DealsDealClosedIterator, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Deals.contract.FilterLogs(opts, "DealClosed", hubRule, clientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &DealsDealClosedIterator{contract: _Deals.contract, event: "DealClosed", logs: logs, sub: sub}, nil
}

// WatchDealClosed is a free log subscription operation binding the contract event 0x72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb890.
//
// Solidity: event DealClosed(hub indexed address, client indexed address, id indexed uint256)
func (_Deals *DealsFilterer) WatchDealClosed(opts *bind.WatchOpts, sink chan<- *DealsDealClosed, hub []common.Address, client []common.Address, id []*big.Int) (event.Subscription, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Deals.contract.WatchLogs(opts, "DealClosed", hubRule, clientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DealsDealClosed)
				if err := _Deals.contract.UnpackLog(event, "DealClosed", log); err != nil {
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

// DealsDealOpenedIterator is returned from FilterDealOpened and is used to iterate over the raw logs and unpacked data for DealOpened events raised by the Deals contract.
type DealsDealOpenedIterator struct {
	Event *DealsDealOpened // Event containing the contract specifics and raw log

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
func (it *DealsDealOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DealsDealOpened)
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
		it.Event = new(DealsDealOpened)
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
func (it *DealsDealOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DealsDealOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DealsDealOpened represents a DealOpened event raised by the Deals contract.
type DealsDealOpened struct {
	Hub    common.Address
	Client common.Address
	Id     *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDealOpened is a free log retrieval operation binding the contract event 0x873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be976848.
//
// Solidity: event DealOpened(hub indexed address, client indexed address, id indexed uint256)
func (_Deals *DealsFilterer) FilterDealOpened(opts *bind.FilterOpts, hub []common.Address, client []common.Address, id []*big.Int) (*DealsDealOpenedIterator, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Deals.contract.FilterLogs(opts, "DealOpened", hubRule, clientRule, idRule)
	if err != nil {
		return nil, err
	}
	return &DealsDealOpenedIterator{contract: _Deals.contract, event: "DealOpened", logs: logs, sub: sub}, nil
}

// WatchDealOpened is a free log subscription operation binding the contract event 0x873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be976848.
//
// Solidity: event DealOpened(hub indexed address, client indexed address, id indexed uint256)
func (_Deals *DealsFilterer) WatchDealOpened(opts *bind.WatchOpts, sink chan<- *DealsDealOpened, hub []common.Address, client []common.Address, id []*big.Int) (event.Subscription, error) {

	var hubRule []interface{}
	for _, hubItem := range hub {
		hubRule = append(hubRule, hubItem)
	}
	var clientRule []interface{}
	for _, clientItem := range client {
		clientRule = append(clientRule, clientItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Deals.contract.WatchLogs(opts, "DealOpened", hubRule, clientRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DealsDealOpened)
				if err := _Deals.contract.UnpackLog(event, "DealOpened", log); err != nil {
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
