// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ApusMarket

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ApusDataClientConfig is an auto generated low-level Go binding around an user-defined struct.
type ApusDataClientConfig struct {
	Owner            common.Address
	Id               *big.Int
	Url              string
	MinFee           *big.Int
	MaxZkEvmInstance uint8
	CurInstance      uint8
	Stat             uint8
}

// ApusMarketMetaData contains all meta data concerning the ApusMarket contract.
var ApusMarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"clients\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxZkEvmInstance\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"curInstance\",\"type\":\"uint8\"},{\"internalType\":\"enumApusData.ClientStatus\",\"name\":\"stat\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getClientCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"getAvilableClientCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getUserClients\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxZkEvmInstance\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"curInstance\",\"type\":\"uint8\"},{\"internalType\":\"enumApusData.ClientStatus\",\"name\":\"stat\",\"type\":\"uint8\"}],\"internalType\":\"structApusData.ClientConfig[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_as\",\"type\":\"address\"}],\"name\":\"setTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxZkEvmInstance\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"curInstance\",\"type\":\"uint8\"},{\"internalType\":\"enumApusData.ClientStatus\",\"name\":\"stat\",\"type\":\"uint8\"}],\"internalType\":\"structApusData.ClientConfig\",\"name\":\"cf\",\"type\":\"tuple\"}],\"name\":\"joinMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLowestN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxZkEvmInstance\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"curInstance\",\"type\":\"uint8\"},{\"internalType\":\"enumApusData.ClientStatus\",\"name\":\"stat\",\"type\":\"uint8\"}],\"internalType\":\"structApusData.ClientConfig\",\"name\":\"cf\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"clientID\",\"type\":\"uint256\"}],\"name\":\"getProverConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxZkEvmInstance\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"curInstance\",\"type\":\"uint8\"},{\"internalType\":\"enumApusData.ClientStatus\",\"name\":\"stat\",\"type\":\"uint8\"}],\"internalType\":\"structApusData.ClientConfig\",\"name\":\"cf\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"clientID\",\"type\":\"uint256\"}],\"name\":\"dispatchTaskToClient\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"clientID\",\"type\":\"uint256\"}],\"name\":\"releaseTaskToClient\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"clientID\",\"type\":\"uint256\"}],\"name\":\"offlineClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketCapacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// ApusMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use ApusMarketMetaData.ABI instead.
var ApusMarketABI = ApusMarketMetaData.ABI

// ApusMarket is an auto generated Go binding around an Ethereum contract.
type ApusMarket struct {
	ApusMarketCaller     // Read-only binding to the contract
	ApusMarketTransactor // Write-only binding to the contract
	ApusMarketFilterer   // Log filterer for contract events
}

// ApusMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApusMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApusMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApusMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApusMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApusMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApusMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApusMarketSession struct {
	Contract     *ApusMarket       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApusMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApusMarketCallerSession struct {
	Contract *ApusMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ApusMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApusMarketTransactorSession struct {
	Contract     *ApusMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ApusMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApusMarketRaw struct {
	Contract *ApusMarket // Generic contract binding to access the raw methods on
}

// ApusMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApusMarketCallerRaw struct {
	Contract *ApusMarketCaller // Generic read-only contract binding to access the raw methods on
}

// ApusMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApusMarketTransactorRaw struct {
	Contract *ApusMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApusMarket creates a new instance of ApusMarket, bound to a specific deployed contract.
func NewApusMarket(address common.Address, backend bind.ContractBackend) (*ApusMarket, error) {
	contract, err := bindApusMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApusMarket{ApusMarketCaller: ApusMarketCaller{contract: contract}, ApusMarketTransactor: ApusMarketTransactor{contract: contract}, ApusMarketFilterer: ApusMarketFilterer{contract: contract}}, nil
}

// NewApusMarketCaller creates a new read-only instance of ApusMarket, bound to a specific deployed contract.
func NewApusMarketCaller(address common.Address, caller bind.ContractCaller) (*ApusMarketCaller, error) {
	contract, err := bindApusMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApusMarketCaller{contract: contract}, nil
}

// NewApusMarketTransactor creates a new write-only instance of ApusMarket, bound to a specific deployed contract.
func NewApusMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*ApusMarketTransactor, error) {
	contract, err := bindApusMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApusMarketTransactor{contract: contract}, nil
}

// NewApusMarketFilterer creates a new log filterer instance of ApusMarket, bound to a specific deployed contract.
func NewApusMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*ApusMarketFilterer, error) {
	contract, err := bindApusMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApusMarketFilterer{contract: contract}, nil
}

// bindApusMarket binds a generic wrapper to an already deployed contract.
func bindApusMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApusMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApusMarket *ApusMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApusMarket.Contract.ApusMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApusMarket *ApusMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApusMarket.Contract.ApusMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApusMarket *ApusMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApusMarket.Contract.ApusMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApusMarket *ApusMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApusMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApusMarket *ApusMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApusMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApusMarket *ApusMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApusMarket.Contract.contract.Transact(opts, method, params...)
}

// Clients is a free data retrieval call binding the contract method 0xf88af21d.
//
// Solidity: function clients(uint256 ) view returns(address owner, uint256 id, string url, uint256 minFee, uint8 maxZkEvmInstance, uint8 curInstance, uint8 stat)
func (_ApusMarket *ApusMarketCaller) Clients(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner            common.Address
	Id               *big.Int
	Url              string
	MinFee           *big.Int
	MaxZkEvmInstance uint8
	CurInstance      uint8
	Stat             uint8
}, error) {
	var out []interface{}
	err := _ApusMarket.contract.Call(opts, &out, "clients", arg0)

	outstruct := new(struct {
		Owner            common.Address
		Id               *big.Int
		Url              string
		MinFee           *big.Int
		MaxZkEvmInstance uint8
		CurInstance      uint8
		Stat             uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Id = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Url = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.MinFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MaxZkEvmInstance = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.CurInstance = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Stat = *abi.ConvertType(out[6], new(uint8)).(*uint8)

	return *outstruct, err

}

// Clients is a free data retrieval call binding the contract method 0xf88af21d.
//
// Solidity: function clients(uint256 ) view returns(address owner, uint256 id, string url, uint256 minFee, uint8 maxZkEvmInstance, uint8 curInstance, uint8 stat)
func (_ApusMarket *ApusMarketSession) Clients(arg0 *big.Int) (struct {
	Owner            common.Address
	Id               *big.Int
	Url              string
	MinFee           *big.Int
	MaxZkEvmInstance uint8
	CurInstance      uint8
	Stat             uint8
}, error) {
	return _ApusMarket.Contract.Clients(&_ApusMarket.CallOpts, arg0)
}

// Clients is a free data retrieval call binding the contract method 0xf88af21d.
//
// Solidity: function clients(uint256 ) view returns(address owner, uint256 id, string url, uint256 minFee, uint8 maxZkEvmInstance, uint8 curInstance, uint8 stat)
func (_ApusMarket *ApusMarketCallerSession) Clients(arg0 *big.Int) (struct {
	Owner            common.Address
	Id               *big.Int
	Url              string
	MinFee           *big.Int
	MaxZkEvmInstance uint8
	CurInstance      uint8
	Stat             uint8
}, error) {
	return _ApusMarket.Contract.Clients(&_ApusMarket.CallOpts, arg0)
}

// GetAvilableClientCount is a free data retrieval call binding the contract method 0x1329d299.
//
// Solidity: function getAvilableClientCount() view returns(uint256)
func (_ApusMarket *ApusMarketCaller) GetAvilableClientCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApusMarket.contract.Call(opts, &out, "getAvilableClientCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAvilableClientCount is a free data retrieval call binding the contract method 0x1329d299.
//
// Solidity: function getAvilableClientCount() view returns(uint256)
func (_ApusMarket *ApusMarketSession) GetAvilableClientCount() (*big.Int, error) {
	return _ApusMarket.Contract.GetAvilableClientCount(&_ApusMarket.CallOpts)
}

// GetAvilableClientCount is a free data retrieval call binding the contract method 0x1329d299.
//
// Solidity: function getAvilableClientCount() view returns(uint256)
func (_ApusMarket *ApusMarketCallerSession) GetAvilableClientCount() (*big.Int, error) {
	return _ApusMarket.Contract.GetAvilableClientCount(&_ApusMarket.CallOpts)
}

// GetClientCount is a free data retrieval call binding the contract method 0x5a81018a.
//
// Solidity: function getClientCount() view returns(uint256)
func (_ApusMarket *ApusMarketCaller) GetClientCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApusMarket.contract.Call(opts, &out, "getClientCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClientCount is a free data retrieval call binding the contract method 0x5a81018a.
//
// Solidity: function getClientCount() view returns(uint256)
func (_ApusMarket *ApusMarketSession) GetClientCount() (*big.Int, error) {
	return _ApusMarket.Contract.GetClientCount(&_ApusMarket.CallOpts)
}

// GetClientCount is a free data retrieval call binding the contract method 0x5a81018a.
//
// Solidity: function getClientCount() view returns(uint256)
func (_ApusMarket *ApusMarketCallerSession) GetClientCount() (*big.Int, error) {
	return _ApusMarket.Contract.GetClientCount(&_ApusMarket.CallOpts)
}

// GetLowestN is a free data retrieval call binding the contract method 0xf476b90d.
//
// Solidity: function getLowestN() view returns(address owner, (address,uint256,string,uint256,uint8,uint8,uint8) cf)
func (_ApusMarket *ApusMarketCaller) GetLowestN(opts *bind.CallOpts) (struct {
	Owner common.Address
	Cf    ApusDataClientConfig
}, error) {
	var out []interface{}
	err := _ApusMarket.contract.Call(opts, &out, "getLowestN")

	outstruct := new(struct {
		Owner common.Address
		Cf    ApusDataClientConfig
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Cf = *abi.ConvertType(out[1], new(ApusDataClientConfig)).(*ApusDataClientConfig)

	return *outstruct, err

}

// GetLowestN is a free data retrieval call binding the contract method 0xf476b90d.
//
// Solidity: function getLowestN() view returns(address owner, (address,uint256,string,uint256,uint8,uint8,uint8) cf)
func (_ApusMarket *ApusMarketSession) GetLowestN() (struct {
	Owner common.Address
	Cf    ApusDataClientConfig
}, error) {
	return _ApusMarket.Contract.GetLowestN(&_ApusMarket.CallOpts)
}

// GetLowestN is a free data retrieval call binding the contract method 0xf476b90d.
//
// Solidity: function getLowestN() view returns(address owner, (address,uint256,string,uint256,uint8,uint8,uint8) cf)
func (_ApusMarket *ApusMarketCallerSession) GetLowestN() (struct {
	Owner common.Address
	Cf    ApusDataClientConfig
}, error) {
	return _ApusMarket.Contract.GetLowestN(&_ApusMarket.CallOpts)
}

// GetProverConfig is a free data retrieval call binding the contract method 0xac459559.
//
// Solidity: function getProverConfig(address owner, uint256 clientID) view returns((address,uint256,string,uint256,uint8,uint8,uint8) cf)
func (_ApusMarket *ApusMarketCaller) GetProverConfig(opts *bind.CallOpts, owner common.Address, clientID *big.Int) (ApusDataClientConfig, error) {
	var out []interface{}
	err := _ApusMarket.contract.Call(opts, &out, "getProverConfig", owner, clientID)

	if err != nil {
		return *new(ApusDataClientConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ApusDataClientConfig)).(*ApusDataClientConfig)

	return out0, err

}

// GetProverConfig is a free data retrieval call binding the contract method 0xac459559.
//
// Solidity: function getProverConfig(address owner, uint256 clientID) view returns((address,uint256,string,uint256,uint8,uint8,uint8) cf)
func (_ApusMarket *ApusMarketSession) GetProverConfig(owner common.Address, clientID *big.Int) (ApusDataClientConfig, error) {
	return _ApusMarket.Contract.GetProverConfig(&_ApusMarket.CallOpts, owner, clientID)
}

// GetProverConfig is a free data retrieval call binding the contract method 0xac459559.
//
// Solidity: function getProverConfig(address owner, uint256 clientID) view returns((address,uint256,string,uint256,uint8,uint8,uint8) cf)
func (_ApusMarket *ApusMarketCallerSession) GetProverConfig(owner common.Address, clientID *big.Int) (ApusDataClientConfig, error) {
	return _ApusMarket.Contract.GetProverConfig(&_ApusMarket.CallOpts, owner, clientID)
}

// GetUserClients is a free data retrieval call binding the contract method 0xca4e5925.
//
// Solidity: function getUserClients(address owner) view returns((address,uint256,string,uint256,uint8,uint8,uint8)[])
func (_ApusMarket *ApusMarketCaller) GetUserClients(opts *bind.CallOpts, owner common.Address) ([]ApusDataClientConfig, error) {
	var out []interface{}
	err := _ApusMarket.contract.Call(opts, &out, "getUserClients", owner)

	if err != nil {
		return *new([]ApusDataClientConfig), err
	}

	out0 := *abi.ConvertType(out[0], new([]ApusDataClientConfig)).(*[]ApusDataClientConfig)

	return out0, err

}

// GetUserClients is a free data retrieval call binding the contract method 0xca4e5925.
//
// Solidity: function getUserClients(address owner) view returns((address,uint256,string,uint256,uint8,uint8,uint8)[])
func (_ApusMarket *ApusMarketSession) GetUserClients(owner common.Address) ([]ApusDataClientConfig, error) {
	return _ApusMarket.Contract.GetUserClients(&_ApusMarket.CallOpts, owner)
}

// GetUserClients is a free data retrieval call binding the contract method 0xca4e5925.
//
// Solidity: function getUserClients(address owner) view returns((address,uint256,string,uint256,uint8,uint8,uint8)[])
func (_ApusMarket *ApusMarketCallerSession) GetUserClients(owner common.Address) ([]ApusDataClientConfig, error) {
	return _ApusMarket.Contract.GetUserClients(&_ApusMarket.CallOpts, owner)
}

// MarketCapacity is a free data retrieval call binding the contract method 0x734ae638.
//
// Solidity: function marketCapacity() view returns(uint256)
func (_ApusMarket *ApusMarketCaller) MarketCapacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ApusMarket.contract.Call(opts, &out, "marketCapacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketCapacity is a free data retrieval call binding the contract method 0x734ae638.
//
// Solidity: function marketCapacity() view returns(uint256)
func (_ApusMarket *ApusMarketSession) MarketCapacity() (*big.Int, error) {
	return _ApusMarket.Contract.MarketCapacity(&_ApusMarket.CallOpts)
}

// MarketCapacity is a free data retrieval call binding the contract method 0x734ae638.
//
// Solidity: function marketCapacity() view returns(uint256)
func (_ApusMarket *ApusMarketCallerSession) MarketCapacity() (*big.Int, error) {
	return _ApusMarket.Contract.MarketCapacity(&_ApusMarket.CallOpts)
}

// DispatchTaskToClient is a paid mutator transaction binding the contract method 0x8c7c7b0d.
//
// Solidity: function dispatchTaskToClient(address owner, uint256 clientID) returns(bool success)
func (_ApusMarket *ApusMarketTransactor) DispatchTaskToClient(opts *bind.TransactOpts, owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.contract.Transact(opts, "dispatchTaskToClient", owner, clientID)
}

// DispatchTaskToClient is a paid mutator transaction binding the contract method 0x8c7c7b0d.
//
// Solidity: function dispatchTaskToClient(address owner, uint256 clientID) returns(bool success)
func (_ApusMarket *ApusMarketSession) DispatchTaskToClient(owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.Contract.DispatchTaskToClient(&_ApusMarket.TransactOpts, owner, clientID)
}

// DispatchTaskToClient is a paid mutator transaction binding the contract method 0x8c7c7b0d.
//
// Solidity: function dispatchTaskToClient(address owner, uint256 clientID) returns(bool success)
func (_ApusMarket *ApusMarketTransactorSession) DispatchTaskToClient(owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.Contract.DispatchTaskToClient(&_ApusMarket.TransactOpts, owner, clientID)
}

// JoinMarket is a paid mutator transaction binding the contract method 0x2d688e74.
//
// Solidity: function joinMarket((address,uint256,string,uint256,uint8,uint8,uint8) cf) returns()
func (_ApusMarket *ApusMarketTransactor) JoinMarket(opts *bind.TransactOpts, cf ApusDataClientConfig) (*types.Transaction, error) {
	return _ApusMarket.contract.Transact(opts, "joinMarket", cf)
}

// JoinMarket is a paid mutator transaction binding the contract method 0x2d688e74.
//
// Solidity: function joinMarket((address,uint256,string,uint256,uint8,uint8,uint8) cf) returns()
func (_ApusMarket *ApusMarketSession) JoinMarket(cf ApusDataClientConfig) (*types.Transaction, error) {
	return _ApusMarket.Contract.JoinMarket(&_ApusMarket.TransactOpts, cf)
}

// JoinMarket is a paid mutator transaction binding the contract method 0x2d688e74.
//
// Solidity: function joinMarket((address,uint256,string,uint256,uint8,uint8,uint8) cf) returns()
func (_ApusMarket *ApusMarketTransactorSession) JoinMarket(cf ApusDataClientConfig) (*types.Transaction, error) {
	return _ApusMarket.Contract.JoinMarket(&_ApusMarket.TransactOpts, cf)
}

// OfflineClient is a paid mutator transaction binding the contract method 0xf60836fc.
//
// Solidity: function offlineClient(address owner, uint256 clientID) returns()
func (_ApusMarket *ApusMarketTransactor) OfflineClient(opts *bind.TransactOpts, owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.contract.Transact(opts, "offlineClient", owner, clientID)
}

// OfflineClient is a paid mutator transaction binding the contract method 0xf60836fc.
//
// Solidity: function offlineClient(address owner, uint256 clientID) returns()
func (_ApusMarket *ApusMarketSession) OfflineClient(owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.Contract.OfflineClient(&_ApusMarket.TransactOpts, owner, clientID)
}

// OfflineClient is a paid mutator transaction binding the contract method 0xf60836fc.
//
// Solidity: function offlineClient(address owner, uint256 clientID) returns()
func (_ApusMarket *ApusMarketTransactorSession) OfflineClient(owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.Contract.OfflineClient(&_ApusMarket.TransactOpts, owner, clientID)
}

// ReleaseTaskToClient is a paid mutator transaction binding the contract method 0xbecb4cda.
//
// Solidity: function releaseTaskToClient(address owner, uint256 clientID) returns(bool)
func (_ApusMarket *ApusMarketTransactor) ReleaseTaskToClient(opts *bind.TransactOpts, owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.contract.Transact(opts, "releaseTaskToClient", owner, clientID)
}

// ReleaseTaskToClient is a paid mutator transaction binding the contract method 0xbecb4cda.
//
// Solidity: function releaseTaskToClient(address owner, uint256 clientID) returns(bool)
func (_ApusMarket *ApusMarketSession) ReleaseTaskToClient(owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.Contract.ReleaseTaskToClient(&_ApusMarket.TransactOpts, owner, clientID)
}

// ReleaseTaskToClient is a paid mutator transaction binding the contract method 0xbecb4cda.
//
// Solidity: function releaseTaskToClient(address owner, uint256 clientID) returns(bool)
func (_ApusMarket *ApusMarketTransactorSession) ReleaseTaskToClient(owner common.Address, clientID *big.Int) (*types.Transaction, error) {
	return _ApusMarket.Contract.ReleaseTaskToClient(&_ApusMarket.TransactOpts, owner, clientID)
}

// SetTask is a paid mutator transaction binding the contract method 0xeda01e9c.
//
// Solidity: function setTask(address _as) returns()
func (_ApusMarket *ApusMarketTransactor) SetTask(opts *bind.TransactOpts, _as common.Address) (*types.Transaction, error) {
	return _ApusMarket.contract.Transact(opts, "setTask", _as)
}

// SetTask is a paid mutator transaction binding the contract method 0xeda01e9c.
//
// Solidity: function setTask(address _as) returns()
func (_ApusMarket *ApusMarketSession) SetTask(_as common.Address) (*types.Transaction, error) {
	return _ApusMarket.Contract.SetTask(&_ApusMarket.TransactOpts, _as)
}

// SetTask is a paid mutator transaction binding the contract method 0xeda01e9c.
//
// Solidity: function setTask(address _as) returns()
func (_ApusMarket *ApusMarketTransactorSession) SetTask(_as common.Address) (*types.Transaction, error) {
	return _ApusMarket.Contract.SetTask(&_ApusMarket.TransactOpts, _as)
}
