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
const DealsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_client\",\"type\":\"address\"}],\"name\":\"GetDealsByClient\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"},{\"name\":\"_client\",\"type\":\"address\"},{\"name\":\"_specHash\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_workTime\",\"type\":\"uint256\"}],\"name\":\"OpenDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dealIndex\",\"type\":\"uint256\"}],\"name\":\"GetDealInfo\",\"outputs\":[{\"name\":\"specHach\",\"type\":\"uint256\"},{\"name\":\"client\",\"type\":\"address\"},{\"name\":\"hub\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"startTime\",\"type\":\"uint256\"},{\"name\":\"workTime\",\"type\":\"uint256\"},{\"name\":\"endTIme\",\"type\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AcceptDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CancelDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"IsPriceAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hub\",\"type\":\"address\"}],\"name\":\"GetDealsByHubAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"CloseDeal\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetDealsAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"hub\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"client\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DealClosed\",\"type\":\"event\"}]"

// DealsBin is the compiled bytecode used for deploying new contracts.
const DealsBin = `0x6060604052600a6000556000600255341561001957600080fd5b604051602080610f84833981016040528080519150505b60018054600160a060020a031916600160a060020a0383161790555b505b610f278061005d6000396000f3006060604052361561007d5763ffffffff60e060020a600035041663371c20bf81146100825780633b24475f146100f55780635ad5f6ae146101255780635be20de21461018e57806371994335146101a65780638027b2dd146101be5780638bcb6c9b146101f4578063ca6ad1bf14610267578063e45ea8d31461027f575b600080fd5b341561008d57600080fd5b6100a1600160a060020a03600435166102a4565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100e15780820151818401525b6020016100c8565b505050509050019250505060405180910390f35b341561010057600080fd5b610123600160a060020a0360043581169060243516604435606435608435610328565b005b341561013057600080fd5b61013b600435610605565b604051978852600160a060020a039687166020890152949095166040808801919091526060870193909352608086019190915260a085015260c084019290925260e0830152610100909101905180910390f35b341561019957600080fd5b61012360043561068e565b005b34156101b157600080fd5b610123600435610798565b005b34156101c957600080fd5b6101e0600160a060020a036004351660243561096a565b604051901515815260200160405180910390f35b34156101ff57600080fd5b6100a1600160a060020a03600435166109f8565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156100e15780820151818401525b6020016100c8565b505050509050019250505060405180910390f35b341561027257600080fd5b610123600435610a7c565b005b341561028a57600080fd5b610292610e66565b60405190815260200160405180910390f35b6102ac610e9e565b6006600083600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561031b57602002820191906000526020600020905b815481526020019060010190808311610307575b505050505090505b919050565b600084600160a060020a031633600160a060020a031614151561034a57600080fd5b5060028054600101908190556101006040519081016040528085815260200186600160a060020a0316815260200187600160a060020a031681526020018481526020016000815260200183815260200160008152602001600060028111156103ae57fe5b9052600082815260036020526040902081518155602082015160018201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055604082015160028201805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015160078201805460ff1916600183600281111561046f57fe5b021790555050600154600160a060020a031690506323b872dd86308660006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156104e357600080fd5b6102c65a03f115156104f457600080fd5b50505060405180519050151561050957600080fd5b600081815260046020526040902054610528908463ffffffff610e6d16565b600082815260046020908152604080832093909355600160a060020a0388168252600690522080546001810161055e8382610eb0565b916000526020600020900160005b5082905550600160a060020a03861660009081526005602052604090208054600181016105998382610eb0565b916000526020600020900160005b50829055507f873cb35202fef184c9f8ee23c04e36dc38f3e26fb285224ca574a837be976848868683604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b505050505050565b60008181526003602081905260408220805460018201546002808401549484015460048501546005860154600687015460078801548a998a998a998a998a998a998a9994989397600160a060020a03938416979690931695909493929160ff9091169081111561067157fe5b985098509850985098509850985098505b50919395975091939597565b60008181526003602052604090206002015433600160a060020a039081169116146106b857600080fd5b60005b60008281526003602052604090206007015460ff1660028111156106db57fe5b146106e557600080fd5b600081815260036020526040902060070180546001919060ff191682805b0217905550600081815260036020526040908190204260048201819055600582015401600682015560028101546001909101547f3a38edea6028913403c74ce8433c90eca94f4ca074d318d8cb77be5290ba4f1592600160a060020a039283169290911690849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b50565b60008181526003602052604090206001015433600160a060020a039081169116146107c257600080fd5b60005b60008281526003602052604090206007015460ff1660028111156107e557fe5b146107ef57600080fd5b600180546000838152600360208190526040808320948501549490910154600160a060020a039384169463a9059cbb9416929091516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561086b57600080fd5b6102c65a03f1151561087c57600080fd5b50505060405180519050151561089157600080fd5b60008181526003602081815260408084209092015460049091529120546108bd9163ffffffff610e8716565b600082815260046020908152604080832093909355600390522060070180546002919060ff19166001835b02179055506000818152600360205260409081902060028101546001909101547f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89092600160a060020a039283169290911690849051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b50565b60015460009081908390600160a060020a031663dd62ed3e8630856040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b15156109d157600080fd5b6102c65a03f115156109e257600080fd5b5050506040518051905003101590505b92915050565b610a00610e9e565b6005600083600160a060020a0316600160a060020a0316815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561031b57602002820191906000526020600020905b815481526020019060010190808311610307575b505050505090505b919050565b60008181526003602052604081206001015433600160a060020a03908116911614610aa657600080fd5b60015b60008381526003602052604090206007015460ff166002811115610ac957fe5b14610ad357600080fd5b600082815260036020526040902060060154421115610bcd5760015460008381526003602081905260408083206002810154920154600160a060020a039485169463a9059cbb949316929091516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610b6757600080fd5b6102c65a03f11515610b7857600080fd5b505050604051805190501515610b8d57600080fd5b6000828152600360208181526040808420909201546004909152912054610bb99163ffffffff610e8716565b600083815260046020526040902055610dc4565b60008281526003602081905260409091206005810154910154811515610bef57fe5b6000848152600360205260408082206004810154600154600290920154429190910395909404949094029450600160a060020a039384169363a9059cbb9316918591516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610c7957600080fd5b6102c65a03f11515610c8a57600080fd5b505050604051805190501515610c9f57600080fd5b600082815260046020526040902054610cbe908263ffffffff610e8716565b60008381526004602090815260408083209390935560018054600392839052848420918201549190920154600160a060020a039283169463a9059cbb9493909216929086900391516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b1515610d4d57600080fd5b6102c65a03f11515610d5e57600080fd5b505050604051805190501515610d7357600080fd5b6000828152600360208181526040808420909201546004909152912054610da29183900363ffffffff610e8716565b6000838152600460209081526040808320939093556003905220426006909101555b600082815260036020526040902060070180546002919060ff19166001835b02179055506000828152600360205260409081902060028101546001909101547f72615f99a62a6cc2f8452d5c0c9cbc5683995297e1d988f09bb1471d4eefb89092600160a060020a039283169290911690859051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15b5050565b6002545b90565b600082820183811015610e7c57fe5b8091505b5092915050565b600082821115610e9357fe5b508082035b92915050565b60206040519081016040526000815290565b815481835581811511610ed457600083815260209020610ed4918101908301610eda565b5b505050565b610e6a91905b80821115610ef45760008155600101610ee0565b5090565b905600a165627a7a7230582008c284895f806df7bc12f4f797086816995a962dff48257346531a2faba177b90029`

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

// GetDealsByClient is a free data retrieval call binding the contract method 0x371c20bf.
//
// Solidity: function GetDealsByClient(_client address) constant returns(uint256[])
func (_Deals *DealsCaller) GetDealsByClient(opts *bind.CallOpts, _client common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDealsByClient", _client)
	return *ret0, err
}

// GetDealsByClient is a free data retrieval call binding the contract method 0x371c20bf.
//
// Solidity: function GetDealsByClient(_client address) constant returns(uint256[])
func (_Deals *DealsSession) GetDealsByClient(_client common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealsByClient(&_Deals.CallOpts, _client)
}

// GetDealsByClient is a free data retrieval call binding the contract method 0x371c20bf.
//
// Solidity: function GetDealsByClient(_client address) constant returns(uint256[])
func (_Deals *DealsCallerSession) GetDealsByClient(_client common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealsByClient(&_Deals.CallOpts, _client)
}

// GetDealsByHubAddress is a free data retrieval call binding the contract method 0x8bcb6c9b.
//
// Solidity: function GetDealsByHubAddress(_hub address) constant returns(uint256[])
func (_Deals *DealsCaller) GetDealsByHubAddress(opts *bind.CallOpts, _hub common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Deals.contract.Call(opts, out, "GetDealsByHubAddress", _hub)
	return *ret0, err
}

// GetDealsByHubAddress is a free data retrieval call binding the contract method 0x8bcb6c9b.
//
// Solidity: function GetDealsByHubAddress(_hub address) constant returns(uint256[])
func (_Deals *DealsSession) GetDealsByHubAddress(_hub common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealsByHubAddress(&_Deals.CallOpts, _hub)
}

// GetDealsByHubAddress is a free data retrieval call binding the contract method 0x8bcb6c9b.
//
// Solidity: function GetDealsByHubAddress(_hub address) constant returns(uint256[])
func (_Deals *DealsCallerSession) GetDealsByHubAddress(_hub common.Address) ([]*big.Int, error) {
	return _Deals.Contract.GetDealsByHubAddress(&_Deals.CallOpts, _hub)
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

// CancelDeal is a paid mutator transaction binding the contract method 0x71994335.
//
// Solidity: function CancelDeal(id uint256) returns()
func (_Deals *DealsTransactor) CancelDeal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Deals.contract.Transact(opts, "CancelDeal", id)
}

// CancelDeal is a paid mutator transaction binding the contract method 0x71994335.
//
// Solidity: function CancelDeal(id uint256) returns()
func (_Deals *DealsSession) CancelDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CancelDeal(&_Deals.TransactOpts, id)
}

// CancelDeal is a paid mutator transaction binding the contract method 0x71994335.
//
// Solidity: function CancelDeal(id uint256) returns()
func (_Deals *DealsTransactorSession) CancelDeal(id *big.Int) (*types.Transaction, error) {
	return _Deals.Contract.CancelDeal(&_Deals.TransactOpts, id)
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
