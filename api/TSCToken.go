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

// TSCTokenABI is the input ABI used to generate the binding from.
const TSCTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"mintAndAllow\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TSCTokenBin is the compiled bytecode used for deploying new contracts.
const TSCTokenBin = `0x606060405260408051908101604052600e81527f545343207465737420746f6b656e0000000000000000000000000000000000006020820152600490805161004b929160200190610109565b5060408051908101604052600381527f747363000000000000000000000000000000000000000000000000000000000060208201526005908051610093929160200190610109565b50601260065534156100a457600080fd5b5b5b60038054600160a060020a03191633600160a060020a03161790555b6b8f76d1b6eaf63dc370000000600081815560038054600160a060020a03191633600160a060020a039081169190911791829055168152600160205260409020555b6101a9565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061014a57805160ff1916838001178555610177565b82800160010185558215610177579182015b8281111561017757825182559160200191906001019061015c565b5b50610184929150610188565b5090565b6101a691905b80821115610184576000815560010161018e565b5090565b90565b610a85806101b86000396000f300606060405236156100c25763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100d5578063095ea7b31461016057806318160ddd1461019657806323b872dd146101bb578063313ce567146101f757806340c10f191461021c57806353a7a1f61461025257806370a082311461028f5780638da5cb5b146102c057806395d89b41146102ef578063a9059cbb1461037a578063dd62ed3e146103b0578063f2fde38b146103e7575b5b6100d1333461271002610408565b505b005b34156100e057600080fd5b6100e86104a3565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101255780820151818401525b60200161010c565b50505050905090810190601f1680156101525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016b57600080fd5b610182600160a060020a0360043516602435610541565b604051901515815260200160405180910390f35b34156101a157600080fd5b6101a96105e8565b60405190815260200160405180910390f35b34156101c657600080fd5b610182600160a060020a03600435811690602435166044356105ee565b604051901515815260200160405180910390f35b341561020257600080fd5b6101a9610703565b60405190815260200160405180910390f35b341561022757600080fd5b610182600160a060020a0360043516602435610408565b604051901515815260200160405180910390f35b341561025d57600080fd5b610182600160a060020a036004358116906024359060443516610709565b604051901515815260200160405180910390f35b341561029a57600080fd5b6101a9600160a060020a0360043516610817565b60405190815260200160405180910390f35b34156102cb57600080fd5b6102d3610836565b604051600160a060020a03909116815260200160405180910390f35b34156102fa57600080fd5b6100e8610845565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101255780820151818401525b60200161010c565b50505050905090810190601f1680156101525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561038557600080fd5b610182600160a060020a03600435166024356108e3565b604051901515815260200160405180910390f35b34156103bb57600080fd5b6101a9600160a060020a03600435811690602435166109a3565b60405190815260200160405180910390f35b34156103f257600080fd5b6100d3600160a060020a03600435166109d0565b005b6000805461041c908363ffffffff610a2816565b6000908155600160a060020a038416815260016020526040902054610447908363ffffffff610a2816565b600160a060020a0384166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859084905190815260200160405180910390a25060015b92915050565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105395780601f1061050e57610100808354040283529160200191610539565b820191906000526020600020905b81548152906001019060200180831161051c57829003601f168201915b505050505081565b60008115806105735750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561057e57600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a038084166000908152600260209081526040808320338516845282528083205493861683526001909152812054909190610635908463ffffffff610a2816565b600160a060020a03808616600090815260016020526040808220939093559087168152205461066a908463ffffffff610a4216565b600160a060020a038616600090815260016020526040902055610693818463ffffffff610a4216565b600160a060020a03808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3600191505b509392505050565b60065481565b6000805461071d908463ffffffff610a2816565b6000908155600160a060020a038516815260016020526040902054610748908463ffffffff610a2816565b600160a060020a0385166000818152600160205260409081902092909255907f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859085905190815260200160405180910390a2600160a060020a0380851660008181526001602081815260408084205433871680865260028452828620978a16808752978452828620829055959094529190527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591905190815260200160405180910390a35060015b9392505050565b600160a060020a0381166000908152600160205260409020545b919050565b600354600160a060020a031681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105395780601f1061050e57610100808354040283529160200191610539565b820191906000526020600020905b81548152906001019060200180831161051c57829003601f168201915b505050505081565b600160a060020a03331660009081526001602052604081205461090c908363ffffffff610a4216565b600160a060020a033381166000908152600160205260408082209390935590851681522054610941908363ffffffff610a2816565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060015b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b92915050565b60035433600160a060020a039081169116146109eb57600080fd5b600160a060020a03811615610a23576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b5b50565b600082820183811015610a3757fe5b8091505b5092915050565b600082821115610a4e57fe5b508082035b929150505600a165627a7a72305820a8f1eefd4408fa20f7871bb9e2f20447a34de1800d5fb13fa4568af081ba6ae00029`

// DeployTSCToken deploys a new Ethereum contract, binding an instance of TSCToken to it.
func DeployTSCToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TSCToken, error) {
	parsed, err := abi.JSON(strings.NewReader(TSCTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TSCTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TSCToken{TSCTokenCaller: TSCTokenCaller{contract: contract}, TSCTokenTransactor: TSCTokenTransactor{contract: contract}}, nil
}

// TSCToken is an auto generated Go binding around an Ethereum contract.
type TSCToken struct {
	TSCTokenCaller     // Read-only binding to the contract
	TSCTokenTransactor // Write-only binding to the contract
}

// TSCTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TSCTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TSCTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TSCTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TSCTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TSCTokenSession struct {
	Contract     *TSCToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TSCTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TSCTokenCallerSession struct {
	Contract *TSCTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TSCTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TSCTokenTransactorSession struct {
	Contract     *TSCTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TSCTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TSCTokenRaw struct {
	Contract *TSCToken // Generic contract binding to access the raw methods on
}

// TSCTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TSCTokenCallerRaw struct {
	Contract *TSCTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TSCTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TSCTokenTransactorRaw struct {
	Contract *TSCTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTSCToken creates a new instance of TSCToken, bound to a specific deployed contract.
func NewTSCToken(address common.Address, backend bind.ContractBackend) (*TSCToken, error) {
	contract, err := bindTSCToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TSCToken{TSCTokenCaller: TSCTokenCaller{contract: contract}, TSCTokenTransactor: TSCTokenTransactor{contract: contract}}, nil
}

// NewTSCTokenCaller creates a new read-only instance of TSCToken, bound to a specific deployed contract.
func NewTSCTokenCaller(address common.Address, caller bind.ContractCaller) (*TSCTokenCaller, error) {
	contract, err := bindTSCToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TSCTokenCaller{contract: contract}, nil
}

// NewTSCTokenTransactor creates a new write-only instance of TSCToken, bound to a specific deployed contract.
func NewTSCTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TSCTokenTransactor, error) {
	contract, err := bindTSCToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TSCTokenTransactor{contract: contract}, nil
}

// bindTSCToken binds a generic wrapper to an already deployed contract.
func bindTSCToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TSCTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TSCToken *TSCTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TSCToken.Contract.TSCTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TSCToken *TSCTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TSCToken.Contract.TSCTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TSCToken *TSCTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TSCToken.Contract.TSCTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TSCToken *TSCTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TSCToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TSCToken *TSCTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TSCToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TSCToken *TSCTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TSCToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TSCToken *TSCTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TSCToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TSCToken *TSCTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TSCToken.Contract.Allowance(&_TSCToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_TSCToken *TSCTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _TSCToken.Contract.Allowance(&_TSCToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TSCToken *TSCTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TSCToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TSCToken *TSCTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TSCToken.Contract.BalanceOf(&_TSCToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_TSCToken *TSCTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _TSCToken.Contract.BalanceOf(&_TSCToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TSCToken *TSCTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TSCToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TSCToken *TSCTokenSession) Decimals() (*big.Int, error) {
	return _TSCToken.Contract.Decimals(&_TSCToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_TSCToken *TSCTokenCallerSession) Decimals() (*big.Int, error) {
	return _TSCToken.Contract.Decimals(&_TSCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TSCToken *TSCTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TSCToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TSCToken *TSCTokenSession) Name() (string, error) {
	return _TSCToken.Contract.Name(&_TSCToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TSCToken *TSCTokenCallerSession) Name() (string, error) {
	return _TSCToken.Contract.Name(&_TSCToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TSCToken *TSCTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TSCToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TSCToken *TSCTokenSession) Owner() (common.Address, error) {
	return _TSCToken.Contract.Owner(&_TSCToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TSCToken *TSCTokenCallerSession) Owner() (common.Address, error) {
	return _TSCToken.Contract.Owner(&_TSCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TSCToken *TSCTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TSCToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TSCToken *TSCTokenSession) Symbol() (string, error) {
	return _TSCToken.Contract.Symbol(&_TSCToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TSCToken *TSCTokenCallerSession) Symbol() (string, error) {
	return _TSCToken.Contract.Symbol(&_TSCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TSCToken *TSCTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TSCToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TSCToken *TSCTokenSession) TotalSupply() (*big.Int, error) {
	return _TSCToken.Contract.TotalSupply(&_TSCToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TSCToken *TSCTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TSCToken.Contract.TotalSupply(&_TSCToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.Approve(&_TSCToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.Approve(&_TSCToken.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_TSCToken *TSCTokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TSCToken.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_TSCToken *TSCTokenSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.Mint(&_TSCToken.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_TSCToken *TSCTokenTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.Mint(&_TSCToken.TransactOpts, _to, _amount)
}

// MintAndAllow is a paid mutator transaction binding the contract method 0x53a7a1f6.
//
// Solidity: function mintAndAllow(_to address, _amount uint256, _spender address) returns(bool)
func (_TSCToken *TSCTokenTransactor) MintAndAllow(opts *bind.TransactOpts, _to common.Address, _amount *big.Int, _spender common.Address) (*types.Transaction, error) {
	return _TSCToken.contract.Transact(opts, "mintAndAllow", _to, _amount, _spender)
}

// MintAndAllow is a paid mutator transaction binding the contract method 0x53a7a1f6.
//
// Solidity: function mintAndAllow(_to address, _amount uint256, _spender address) returns(bool)
func (_TSCToken *TSCTokenSession) MintAndAllow(_to common.Address, _amount *big.Int, _spender common.Address) (*types.Transaction, error) {
	return _TSCToken.Contract.MintAndAllow(&_TSCToken.TransactOpts, _to, _amount, _spender)
}

// MintAndAllow is a paid mutator transaction binding the contract method 0x53a7a1f6.
//
// Solidity: function mintAndAllow(_to address, _amount uint256, _spender address) returns(bool)
func (_TSCToken *TSCTokenTransactorSession) MintAndAllow(_to common.Address, _amount *big.Int, _spender common.Address) (*types.Transaction, error) {
	return _TSCToken.Contract.MintAndAllow(&_TSCToken.TransactOpts, _to, _amount, _spender)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.Transfer(&_TSCToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.Transfer(&_TSCToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.TransferFrom(&_TSCToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_TSCToken *TSCTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _TSCToken.Contract.TransferFrom(&_TSCToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TSCToken *TSCTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TSCToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TSCToken *TSCTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TSCToken.Contract.TransferOwnership(&_TSCToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TSCToken *TSCTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TSCToken.Contract.TransferOwnership(&_TSCToken.TransactOpts, newOwner)
}
