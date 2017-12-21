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

// SNMTTokenABI is the input ABI used to generate the binding from.
const SNMTTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"mintedAmount\",\"type\":\"uint256\"}],\"name\":\"mintToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"whom\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GiveAway\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// SNMTTokenBin is the compiled bytecode used for deploying new contracts.
const SNMTTokenBin = `0x606060405260408051908101604052600f81527f534f4e4d207465737420746f6b656e00000000000000000000000000000000006020820152600490805161004b9291602001906100e3565b5060408051908101604052600481527f534e4d5400000000000000000000000000000000000000000000000000000000602082015260059080516100939291602001906100e3565b50601260065534156100a457600080fd5b5b5b60038054600160a060020a03191633600160a060020a03161790555b60038054600160a060020a03191633600160a060020a03161790555b610183565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061012457805160ff1916838001178555610151565b82800160010185558215610151579182015b82811115610151578251825591602001919060010190610136565b5b5061015e929150610162565b5090565b61018091905b8082111561015e5760008155600101610168565b5090565b90565b610a17806101926000396000f300606060405236156100c25763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100cf578063095ea7b31461015a57806318160ddd1461019057806323b872dd146101b5578063313ce567146101f157806370a082311461021657806379c65068146102475780638da5cb5b1461027d57806395d89b41146102ac578063a9059cbb14610337578063aa6ca8081461036d578063dd62ed3e14610394578063f2fde38b146103cb575b5b6100cb6103ec565b505b005b34156100da57600080fd5b6100e261049b565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561011f5780820151818401525b602001610106565b50505050905090810190601f16801561014c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561016557600080fd5b61017c600160a060020a0360043516602435610539565b604051901515815260200160405180910390f35b341561019b57600080fd5b6101a36105e0565b60405190815260200160405180910390f35b34156101c057600080fd5b61017c600160a060020a03600435811690602435166044356105e6565b604051901515815260200160405180910390f35b34156101fc57600080fd5b6101a36106fb565b60405190815260200160405180910390f35b341561022157600080fd5b6101a3600160a060020a0360043516610701565b60405190815260200160405180910390f35b341561025257600080fd5b61017c600160a060020a0360043516602435610720565b604051901515815260200160405180910390f35b341561028857600080fd5b6102906107c8565b604051600160a060020a03909116815260200160405180910390f35b34156102b757600080fd5b6100e26107d7565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561011f5780820151818401525b602001610106565b50505050905090810190601f16801561014c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561034257600080fd5b61017c600160a060020a0360043516602435610875565b604051901515815260200160405180910390f35b341561037857600080fd5b61017c6103ec565b604051901515815260200160405180910390f35b341561039f57600080fd5b6101a3600160a060020a0360043581169060243516610935565b60405190815260200160405180910390f35b34156103d657600080fd5b6100cd600160a060020a0360043516610962565b005b600160a060020a033316600090815260016020526040812054683635c9adc5dea0000090610420908263ffffffff6109ba16565b600160a060020a0333166000908152600160205260408120919091555461044d908263ffffffff6109ba16565b507fe08e9d066634006283658128ec91f58d444719d7a07d49f72924da4352ff94ad3382604051600160a060020a03909216825260208201526040908101905180910390a1600191505b5090565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105315780601f1061050657610100808354040283529160200191610531565b820191906000526020600020905b81548152906001019060200180831161051457829003601f168201915b505050505081565b600081158061056b5750600160a060020a03338116600090815260026020908152604080832093871683529290522054155b151561057657600080fd5b600160a060020a03338116600081815260026020908152604080832094881680845294909152908190208590557f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259085905190815260200160405180910390a35060015b92915050565b60005481565b600160a060020a03808416600090815260026020908152604080832033851684528252808320549386168352600190915281205490919061062d908463ffffffff6109ba16565b600160a060020a038086166000908152600160205260408082209390935590871681522054610662908463ffffffff6109d416565b600160a060020a03861660009081526001602052604090205561068b818463ffffffff6109d416565b600160a060020a03808716600081815260026020908152604080832033861684529091529081902093909355908616917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9086905190815260200160405180910390a3600191505b509392505050565b60065481565b600160a060020a0381166000908152600160205260409020545b919050565b60035460009033600160a060020a0390811691161461073e57600080fd5b600160a060020a038316600090815260016020526040902054610767908363ffffffff6109ba16565b5060005461077b908363ffffffff6109ba16565b507fe08e9d066634006283658128ec91f58d444719d7a07d49f72924da4352ff94ad8383604051600160a060020a03909216825260208201526040908101905180910390a15b5b92915050565b600354600160a060020a031681565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105315780601f1061050657610100808354040283529160200191610531565b820191906000526020600020905b81548152906001019060200180831161051457829003601f168201915b505050505081565b600160a060020a03331660009081526001602052604081205461089e908363ffffffff6109d416565b600160a060020a0333811660009081526001602052604080822093909355908516815220546108d3908363ffffffff6109ba16565b600160a060020a0380851660008181526001602052604090819020939093559133909116907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9085905190815260200160405180910390a35060015b92915050565b600160a060020a038083166000908152600260209081526040808320938516835292905220545b92915050565b60035433600160a060020a0390811691161461097d57600080fd5b600160a060020a038116156109b5576003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b5b50565b6000828201838110156109c957fe5b8091505b5092915050565b6000828211156109e057fe5b508082035b929150505600a165627a7a7230582087490a420d63e3c81b42e0375d076ffcd38c7996ae9ca072d965fd088e0f1cf60029`

// DeploySNMTToken deploys a new Ethereum contract, binding an instance of SNMTToken to it.
func DeploySNMTToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SNMTToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SNMTTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SNMTTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SNMTToken{SNMTTokenCaller: SNMTTokenCaller{contract: contract}, SNMTTokenTransactor: SNMTTokenTransactor{contract: contract}}, nil
}

// SNMTToken is an auto generated Go binding around an Ethereum contract.
type SNMTToken struct {
	SNMTTokenCaller     // Read-only binding to the contract
	SNMTTokenTransactor // Write-only binding to the contract
}

// SNMTTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SNMTTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SNMTTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SNMTTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SNMTTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SNMTTokenSession struct {
	Contract     *SNMTToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SNMTTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SNMTTokenCallerSession struct {
	Contract *SNMTTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SNMTTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SNMTTokenTransactorSession struct {
	Contract     *SNMTTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SNMTTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SNMTTokenRaw struct {
	Contract *SNMTToken // Generic contract binding to access the raw methods on
}

// SNMTTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SNMTTokenCallerRaw struct {
	Contract *SNMTTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SNMTTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SNMTTokenTransactorRaw struct {
	Contract *SNMTTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSNMTToken creates a new instance of SNMTToken, bound to a specific deployed contract.
func NewSNMTToken(address common.Address, backend bind.ContractBackend) (*SNMTToken, error) {
	contract, err := bindSNMTToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SNMTToken{SNMTTokenCaller: SNMTTokenCaller{contract: contract}, SNMTTokenTransactor: SNMTTokenTransactor{contract: contract}}, nil
}

// NewSNMTTokenCaller creates a new read-only instance of SNMTToken, bound to a specific deployed contract.
func NewSNMTTokenCaller(address common.Address, caller bind.ContractCaller) (*SNMTTokenCaller, error) {
	contract, err := bindSNMTToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SNMTTokenCaller{contract: contract}, nil
}

// NewSNMTTokenTransactor creates a new write-only instance of SNMTToken, bound to a specific deployed contract.
func NewSNMTTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SNMTTokenTransactor, error) {
	contract, err := bindSNMTToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SNMTTokenTransactor{contract: contract}, nil
}

// bindSNMTToken binds a generic wrapper to an already deployed contract.
func bindSNMTToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SNMTTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SNMTToken *SNMTTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SNMTToken.Contract.SNMTTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SNMTToken *SNMTTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SNMTToken.Contract.SNMTTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SNMTToken *SNMTTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SNMTToken.Contract.SNMTTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SNMTToken *SNMTTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SNMTToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SNMTToken *SNMTTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SNMTToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SNMTToken *SNMTTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SNMTToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SNMTToken *SNMTTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNMTToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SNMTToken *SNMTTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SNMTToken.Contract.Allowance(&_SNMTToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_SNMTToken *SNMTTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _SNMTToken.Contract.Allowance(&_SNMTToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SNMTToken *SNMTTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNMTToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SNMTToken *SNMTTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SNMTToken.Contract.BalanceOf(&_SNMTToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_SNMTToken *SNMTTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _SNMTToken.Contract.BalanceOf(&_SNMTToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SNMTToken *SNMTTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNMTToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SNMTToken *SNMTTokenSession) Decimals() (*big.Int, error) {
	return _SNMTToken.Contract.Decimals(&_SNMTToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SNMTToken *SNMTTokenCallerSession) Decimals() (*big.Int, error) {
	return _SNMTToken.Contract.Decimals(&_SNMTToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SNMTToken *SNMTTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SNMTToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SNMTToken *SNMTTokenSession) Name() (string, error) {
	return _SNMTToken.Contract.Name(&_SNMTToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SNMTToken *SNMTTokenCallerSession) Name() (string, error) {
	return _SNMTToken.Contract.Name(&_SNMTToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SNMTToken *SNMTTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SNMTToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SNMTToken *SNMTTokenSession) Owner() (common.Address, error) {
	return _SNMTToken.Contract.Owner(&_SNMTToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SNMTToken *SNMTTokenCallerSession) Owner() (common.Address, error) {
	return _SNMTToken.Contract.Owner(&_SNMTToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SNMTToken *SNMTTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SNMTToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SNMTToken *SNMTTokenSession) Symbol() (string, error) {
	return _SNMTToken.Contract.Symbol(&_SNMTToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_SNMTToken *SNMTTokenCallerSession) Symbol() (string, error) {
	return _SNMTToken.Contract.Symbol(&_SNMTToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SNMTToken *SNMTTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SNMTToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SNMTToken *SNMTTokenSession) TotalSupply() (*big.Int, error) {
	return _SNMTToken.Contract.TotalSupply(&_SNMTToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SNMTToken *SNMTTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SNMTToken.Contract.TotalSupply(&_SNMTToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.Approve(&_SNMTToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.Approve(&_SNMTToken.TransactOpts, _spender, _value)
}

// GetTokens is a paid mutator transaction binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() returns(bool)
func (_SNMTToken *SNMTTokenTransactor) GetTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SNMTToken.contract.Transact(opts, "getTokens")
}

// GetTokens is a paid mutator transaction binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() returns(bool)
func (_SNMTToken *SNMTTokenSession) GetTokens() (*types.Transaction, error) {
	return _SNMTToken.Contract.GetTokens(&_SNMTToken.TransactOpts)
}

// GetTokens is a paid mutator transaction binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() returns(bool)
func (_SNMTToken *SNMTTokenTransactorSession) GetTokens() (*types.Transaction, error) {
	return _SNMTToken.Contract.GetTokens(&_SNMTToken.TransactOpts)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactor) MintToken(opts *bind.TransactOpts, target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _SNMTToken.contract.Transact(opts, "mintToken", target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns(bool)
func (_SNMTToken *SNMTTokenSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.MintToken(&_SNMTToken.TransactOpts, target, mintedAmount)
}

// MintToken is a paid mutator transaction binding the contract method 0x79c65068.
//
// Solidity: function mintToken(target address, mintedAmount uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactorSession) MintToken(target common.Address, mintedAmount *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.MintToken(&_SNMTToken.TransactOpts, target, mintedAmount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.Transfer(&_SNMTToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.Transfer(&_SNMTToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.TransferFrom(&_SNMTToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_SNMTToken *SNMTTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _SNMTToken.Contract.TransferFrom(&_SNMTToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SNMTToken *SNMTTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SNMTToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SNMTToken *SNMTTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SNMTToken.Contract.TransferOwnership(&_SNMTToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_SNMTToken *SNMTTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SNMTToken.Contract.TransferOwnership(&_SNMTToken.TransactOpts, newOwner)
}
