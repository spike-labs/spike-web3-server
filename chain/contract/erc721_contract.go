// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// Erc721ContractMetaData contains all meta data concerning the Erc721Contract contract.
var Erc721ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"AdminEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"UpdateUser\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"receivers\",\"type\":\"address[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"tokenURIs\",\"type\":\"string[]\"}],\"name\":\"batchMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"disableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"enableAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"baseURI_\",\"type\":\"string\"}],\"name\":\"setBaseTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setDefaultRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"feeNumerator\",\"type\":\"uint96\"}],\"name\":\"setTokenRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"setTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"expires\",\"type\":\"uint64\"}],\"name\":\"setUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"direct\",\"type\":\"bool\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"userOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Erc721ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc721ContractMetaData.ABI instead.
var Erc721ContractABI = Erc721ContractMetaData.ABI

// Erc721Contract is an auto generated Go binding around an Ethereum contract.
type Erc721Contract struct {
	Erc721ContractCaller     // Read-only binding to the contract
	Erc721ContractTransactor // Write-only binding to the contract
	Erc721ContractFilterer   // Log filterer for contract events
}

// Erc721ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type Erc721ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc721ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc721ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc721ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc721ContractSession struct {
	Contract     *Erc721Contract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc721ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc721ContractCallerSession struct {
	Contract *Erc721ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// Erc721ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc721ContractTransactorSession struct {
	Contract     *Erc721ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// Erc721ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type Erc721ContractRaw struct {
	Contract *Erc721Contract // Generic contract binding to access the raw methods on
}

// Erc721ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc721ContractCallerRaw struct {
	Contract *Erc721ContractCaller // Generic read-only contract binding to access the raw methods on
}

// Erc721ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc721ContractTransactorRaw struct {
	Contract *Erc721ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErc721Contract creates a new instance of Erc721Contract, bound to a specific deployed contract.
func NewErc721Contract(address common.Address, backend bind.ContractBackend) (*Erc721Contract, error) {
	contract, err := bindErc721Contract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc721Contract{Erc721ContractCaller: Erc721ContractCaller{contract: contract}, Erc721ContractTransactor: Erc721ContractTransactor{contract: contract}, Erc721ContractFilterer: Erc721ContractFilterer{contract: contract}}, nil
}

// NewErc721ContractCaller creates a new read-only instance of Erc721Contract, bound to a specific deployed contract.
func NewErc721ContractCaller(address common.Address, caller bind.ContractCaller) (*Erc721ContractCaller, error) {
	contract, err := bindErc721Contract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractCaller{contract: contract}, nil
}

// NewErc721ContractTransactor creates a new write-only instance of Erc721Contract, bound to a specific deployed contract.
func NewErc721ContractTransactor(address common.Address, transactor bind.ContractTransactor) (*Erc721ContractTransactor, error) {
	contract, err := bindErc721Contract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractTransactor{contract: contract}, nil
}

// NewErc721ContractFilterer creates a new log filterer instance of Erc721Contract, bound to a specific deployed contract.
func NewErc721ContractFilterer(address common.Address, filterer bind.ContractFilterer) (*Erc721ContractFilterer, error) {
	contract, err := bindErc721Contract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractFilterer{contract: contract}, nil
}

// bindErc721Contract binds a generic wrapper to an already deployed contract.
func bindErc721Contract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc721ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Contract *Erc721ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721Contract.Contract.Erc721ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Contract *Erc721ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Erc721ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Contract *Erc721ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Erc721ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc721Contract *Erc721ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc721Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc721Contract *Erc721ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc721Contract *Erc721ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc721Contract.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Erc721Contract *Erc721ContractCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Erc721Contract *Erc721ContractSession) Admins(arg0 common.Address) (bool, error) {
	return _Erc721Contract.Contract.Admins(&_Erc721Contract.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_Erc721Contract *Erc721ContractCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _Erc721Contract.Contract.Admins(&_Erc721Contract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721Contract *Erc721ContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721Contract *Erc721ContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721Contract.Contract.BalanceOf(&_Erc721Contract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Erc721Contract *Erc721ContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc721Contract.Contract.BalanceOf(&_Erc721Contract.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721Contract.Contract.GetApproved(&_Erc721Contract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Erc721Contract.Contract.GetApproved(&_Erc721Contract.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721Contract *Erc721ContractCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721Contract *Erc721ContractSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721Contract.Contract.IsApprovedForAll(&_Erc721Contract.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Erc721Contract *Erc721ContractCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Erc721Contract.Contract.IsApprovedForAll(&_Erc721Contract.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721Contract *Erc721ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721Contract *Erc721ContractSession) Name() (string, error) {
	return _Erc721Contract.Contract.Name(&_Erc721Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc721Contract *Erc721ContractCallerSession) Name() (string, error) {
	return _Erc721Contract.Contract.Name(&_Erc721Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721Contract *Erc721ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721Contract *Erc721ContractSession) Owner() (common.Address, error) {
	return _Erc721Contract.Contract.Owner(&_Erc721Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Erc721Contract *Erc721ContractCallerSession) Owner() (common.Address, error) {
	return _Erc721Contract.Contract.Owner(&_Erc721Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721Contract.Contract.OwnerOf(&_Erc721Contract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721Contract.Contract.OwnerOf(&_Erc721Contract.CallOpts, tokenId)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Erc721Contract *Erc721ContractCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Erc721Contract *Erc721ContractSession) PendingOwner() (common.Address, error) {
	return _Erc721Contract.Contract.PendingOwner(&_Erc721Contract.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Erc721Contract *Erc721ContractCallerSession) PendingOwner() (common.Address, error) {
	return _Erc721Contract.Contract.PendingOwner(&_Erc721Contract.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Erc721Contract *Erc721ContractCaller) RoyaltyInfo(opts *bind.CallOpts, _tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "royaltyInfo", _tokenId, _salePrice)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Erc721Contract *Erc721ContractSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Erc721Contract.Contract.RoyaltyInfo(&_Erc721Contract.CallOpts, _tokenId, _salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 _tokenId, uint256 _salePrice) view returns(address, uint256)
func (_Erc721Contract *Erc721ContractCallerSession) RoyaltyInfo(_tokenId *big.Int, _salePrice *big.Int) (common.Address, *big.Int, error) {
	return _Erc721Contract.Contract.RoyaltyInfo(&_Erc721Contract.CallOpts, _tokenId, _salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721Contract *Erc721ContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721Contract *Erc721ContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721Contract.Contract.SupportsInterface(&_Erc721Contract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Erc721Contract *Erc721ContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Erc721Contract.Contract.SupportsInterface(&_Erc721Contract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721Contract *Erc721ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721Contract *Erc721ContractSession) Symbol() (string, error) {
	return _Erc721Contract.Contract.Symbol(&_Erc721Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc721Contract *Erc721ContractCallerSession) Symbol() (string, error) {
	return _Erc721Contract.Contract.Symbol(&_Erc721Contract.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721Contract *Erc721ContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721Contract *Erc721ContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721Contract.Contract.TokenURI(&_Erc721Contract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Erc721Contract *Erc721ContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Erc721Contract.Contract.TokenURI(&_Erc721Contract.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_Erc721Contract *Erc721ContractCaller) UserExpires(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "userExpires", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_Erc721Contract *Erc721ContractSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _Erc721Contract.Contract.UserExpires(&_Erc721Contract.CallOpts, tokenId)
}

// UserExpires is a free data retrieval call binding the contract method 0x8fc88c48.
//
// Solidity: function userExpires(uint256 tokenId) view returns(uint256)
func (_Erc721Contract *Erc721ContractCallerSession) UserExpires(tokenId *big.Int) (*big.Int, error) {
	return _Erc721Contract.Contract.UserExpires(&_Erc721Contract.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractCaller) UserOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Erc721Contract.contract.Call(opts, &out, "userOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721Contract.Contract.UserOf(&_Erc721Contract.CallOpts, tokenId)
}

// UserOf is a free data retrieval call binding the contract method 0xc2f1f14a.
//
// Solidity: function userOf(uint256 tokenId) view returns(address)
func (_Erc721Contract *Erc721ContractCallerSession) UserOf(tokenId *big.Int) (common.Address, error) {
	return _Erc721Contract.Contract.UserOf(&_Erc721Contract.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Approve(&_Erc721Contract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Approve(&_Erc721Contract.TransactOpts, to, tokenId)
}

// BatchMint is a paid mutator transaction binding the contract method 0x5a86c41a.
//
// Solidity: function batchMint(uint256[] tokenIds, address[] receivers) returns()
func (_Erc721Contract *Erc721ContractTransactor) BatchMint(opts *bind.TransactOpts, tokenIds []*big.Int, receivers []common.Address) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "batchMint", tokenIds, receivers)
}

// BatchMint is a paid mutator transaction binding the contract method 0x5a86c41a.
//
// Solidity: function batchMint(uint256[] tokenIds, address[] receivers) returns()
func (_Erc721Contract *Erc721ContractSession) BatchMint(tokenIds []*big.Int, receivers []common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.BatchMint(&_Erc721Contract.TransactOpts, tokenIds, receivers)
}

// BatchMint is a paid mutator transaction binding the contract method 0x5a86c41a.
//
// Solidity: function batchMint(uint256[] tokenIds, address[] receivers) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) BatchMint(tokenIds []*big.Int, receivers []common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.BatchMint(&_Erc721Contract.TransactOpts, tokenIds, receivers)
}

// BatchMint0 is a paid mutator transaction binding the contract method 0xbbd64d44.
//
// Solidity: function batchMint(uint256[] tokenIds, address to, string[] tokenURIs) returns()
func (_Erc721Contract *Erc721ContractTransactor) BatchMint0(opts *bind.TransactOpts, tokenIds []*big.Int, to common.Address, tokenURIs []string) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "batchMint0", tokenIds, to, tokenURIs)
}

// BatchMint0 is a paid mutator transaction binding the contract method 0xbbd64d44.
//
// Solidity: function batchMint(uint256[] tokenIds, address to, string[] tokenURIs) returns()
func (_Erc721Contract *Erc721ContractSession) BatchMint0(tokenIds []*big.Int, to common.Address, tokenURIs []string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.BatchMint0(&_Erc721Contract.TransactOpts, tokenIds, to, tokenURIs)
}

// BatchMint0 is a paid mutator transaction binding the contract method 0xbbd64d44.
//
// Solidity: function batchMint(uint256[] tokenIds, address to, string[] tokenURIs) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) BatchMint0(tokenIds []*big.Int, to common.Address, tokenURIs []string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.BatchMint0(&_Erc721Contract.TransactOpts, tokenIds, to, tokenURIs)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Erc721Contract *Erc721ContractTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Erc721Contract *Erc721ContractSession) ClaimOwnership() (*types.Transaction, error) {
	return _Erc721Contract.Contract.ClaimOwnership(&_Erc721Contract.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Erc721Contract *Erc721ContractTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Erc721Contract.Contract.ClaimOwnership(&_Erc721Contract.TransactOpts)
}

// DisableAdmin is a paid mutator transaction binding the contract method 0x751e9a9c.
//
// Solidity: function disableAdmin(address _addr) returns()
func (_Erc721Contract *Erc721ContractTransactor) DisableAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "disableAdmin", _addr)
}

// DisableAdmin is a paid mutator transaction binding the contract method 0x751e9a9c.
//
// Solidity: function disableAdmin(address _addr) returns()
func (_Erc721Contract *Erc721ContractSession) DisableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.DisableAdmin(&_Erc721Contract.TransactOpts, _addr)
}

// DisableAdmin is a paid mutator transaction binding the contract method 0x751e9a9c.
//
// Solidity: function disableAdmin(address _addr) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) DisableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.DisableAdmin(&_Erc721Contract.TransactOpts, _addr)
}

// EnableAdmin is a paid mutator transaction binding the contract method 0xbea532ff.
//
// Solidity: function enableAdmin(address _addr) returns()
func (_Erc721Contract *Erc721ContractTransactor) EnableAdmin(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "enableAdmin", _addr)
}

// EnableAdmin is a paid mutator transaction binding the contract method 0xbea532ff.
//
// Solidity: function enableAdmin(address _addr) returns()
func (_Erc721Contract *Erc721ContractSession) EnableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.EnableAdmin(&_Erc721Contract.TransactOpts, _addr)
}

// EnableAdmin is a paid mutator transaction binding the contract method 0xbea532ff.
//
// Solidity: function enableAdmin(address _addr) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) EnableAdmin(_addr common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.EnableAdmin(&_Erc721Contract.TransactOpts, _addr)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_Erc721Contract *Erc721ContractTransactor) Mint(opts *bind.TransactOpts, tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "mint", tokenId, to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_Erc721Contract *Erc721ContractSession) Mint(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Mint(&_Erc721Contract.TransactOpts, tokenId, to)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 tokenId, address to) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) Mint(tokenId *big.Int, to common.Address) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Mint(&_Erc721Contract.TransactOpts, tokenId, to)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe67e402c.
//
// Solidity: function mint(uint256 tokenId, address to, string _tokenURI) returns()
func (_Erc721Contract *Erc721ContractTransactor) Mint0(opts *bind.TransactOpts, tokenId *big.Int, to common.Address, _tokenURI string) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "mint0", tokenId, to, _tokenURI)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe67e402c.
//
// Solidity: function mint(uint256 tokenId, address to, string _tokenURI) returns()
func (_Erc721Contract *Erc721ContractSession) Mint0(tokenId *big.Int, to common.Address, _tokenURI string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Mint0(&_Erc721Contract.TransactOpts, tokenId, to, _tokenURI)
}

// Mint0 is a paid mutator transaction binding the contract method 0xe67e402c.
//
// Solidity: function mint(uint256 tokenId, address to, string _tokenURI) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) Mint0(tokenId *big.Int, to common.Address, _tokenURI string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.Mint0(&_Erc721Contract.TransactOpts, tokenId, to, _tokenURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721Contract *Erc721ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721Contract *Erc721ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc721Contract.Contract.RenounceOwnership(&_Erc721Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Erc721Contract *Erc721ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Erc721Contract.Contract.RenounceOwnership(&_Erc721Contract.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SafeTransferFrom(&_Erc721Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SafeTransferFrom(&_Erc721Contract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc721Contract *Erc721ContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc721Contract *Erc721ContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SafeTransferFrom0(&_Erc721Contract.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SafeTransferFrom0(&_Erc721Contract.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721Contract *Erc721ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721Contract *Erc721ContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetApprovalForAll(&_Erc721Contract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetApprovalForAll(&_Erc721Contract.TransactOpts, operator, approved)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseURI_) returns()
func (_Erc721Contract *Erc721ContractTransactor) SetBaseTokenURI(opts *bind.TransactOpts, baseURI_ string) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "setBaseTokenURI", baseURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseURI_) returns()
func (_Erc721Contract *Erc721ContractSession) SetBaseTokenURI(baseURI_ string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetBaseTokenURI(&_Erc721Contract.TransactOpts, baseURI_)
}

// SetBaseTokenURI is a paid mutator transaction binding the contract method 0x30176e13.
//
// Solidity: function setBaseTokenURI(string baseURI_) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SetBaseTokenURI(baseURI_ string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetBaseTokenURI(&_Erc721Contract.TransactOpts, baseURI_)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Erc721Contract *Erc721ContractTransactor) SetDefaultRoyalty(opts *bind.TransactOpts, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "setDefaultRoyalty", receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Erc721Contract *Erc721ContractSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetDefaultRoyalty(&_Erc721Contract.TransactOpts, receiver, feeNumerator)
}

// SetDefaultRoyalty is a paid mutator transaction binding the contract method 0x04634d8d.
//
// Solidity: function setDefaultRoyalty(address receiver, uint96 feeNumerator) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SetDefaultRoyalty(receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetDefaultRoyalty(&_Erc721Contract.TransactOpts, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Erc721Contract *Erc721ContractTransactor) SetTokenRoyalty(opts *bind.TransactOpts, tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "setTokenRoyalty", tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Erc721Contract *Erc721ContractSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetTokenRoyalty(&_Erc721Contract.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenRoyalty is a paid mutator transaction binding the contract method 0x5944c753.
//
// Solidity: function setTokenRoyalty(uint256 tokenId, address receiver, uint96 feeNumerator) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SetTokenRoyalty(tokenId *big.Int, receiver common.Address, feeNumerator *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetTokenRoyalty(&_Erc721Contract.TransactOpts, tokenId, receiver, feeNumerator)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Erc721Contract *Erc721ContractTransactor) SetTokenURI(opts *bind.TransactOpts, tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "setTokenURI", tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Erc721Contract *Erc721ContractSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetTokenURI(&_Erc721Contract.TransactOpts, tokenId, _tokenURI)
}

// SetTokenURI is a paid mutator transaction binding the contract method 0x162094c4.
//
// Solidity: function setTokenURI(uint256 tokenId, string _tokenURI) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SetTokenURI(tokenId *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetTokenURI(&_Erc721Contract.TransactOpts, tokenId, _tokenURI)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_Erc721Contract *Erc721ContractTransactor) SetUser(opts *bind.TransactOpts, tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "setUser", tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_Erc721Contract *Erc721ContractSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetUser(&_Erc721Contract.TransactOpts, tokenId, user, expires)
}

// SetUser is a paid mutator transaction binding the contract method 0xe030565e.
//
// Solidity: function setUser(uint256 tokenId, address user, uint64 expires) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) SetUser(tokenId *big.Int, user common.Address, expires uint64) (*types.Transaction, error) {
	return _Erc721Contract.Contract.SetUser(&_Erc721Contract.TransactOpts, tokenId, user, expires)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.TransferFrom(&_Erc721Contract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Erc721Contract.Contract.TransferFrom(&_Erc721Contract.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address newOwner, bool direct) returns()
func (_Erc721Contract *Erc721ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address, direct bool) (*types.Transaction, error) {
	return _Erc721Contract.contract.Transact(opts, "transferOwnership", newOwner, direct)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address newOwner, bool direct) returns()
func (_Erc721Contract *Erc721ContractSession) TransferOwnership(newOwner common.Address, direct bool) (*types.Transaction, error) {
	return _Erc721Contract.Contract.TransferOwnership(&_Erc721Contract.TransactOpts, newOwner, direct)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xb242e534.
//
// Solidity: function transferOwnership(address newOwner, bool direct) returns()
func (_Erc721Contract *Erc721ContractTransactorSession) TransferOwnership(newOwner common.Address, direct bool) (*types.Transaction, error) {
	return _Erc721Contract.Contract.TransferOwnership(&_Erc721Contract.TransactOpts, newOwner, direct)
}

// Erc721ContractAdminEnabledIterator is returned from FilterAdminEnabled and is used to iterate over the raw logs and unpacked data for AdminEnabled events raised by the Erc721Contract contract.
type Erc721ContractAdminEnabledIterator struct {
	Event *Erc721ContractAdminEnabled // Event containing the contract specifics and raw log

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
func (it *Erc721ContractAdminEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContractAdminEnabled)
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
		it.Event = new(Erc721ContractAdminEnabled)
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
func (it *Erc721ContractAdminEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContractAdminEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContractAdminEnabled represents a AdminEnabled event raised by the Erc721Contract contract.
type Erc721ContractAdminEnabled struct {
	Admin   common.Address
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAdminEnabled is a free log retrieval operation binding the contract event 0xb3714ef62726d54fcac8919290b50b999590b382adefc7f78eddf14aafb9ba5e.
//
// Solidity: event AdminEnabled(address admin, bool enabled)
func (_Erc721Contract *Erc721ContractFilterer) FilterAdminEnabled(opts *bind.FilterOpts) (*Erc721ContractAdminEnabledIterator, error) {

	logs, sub, err := _Erc721Contract.contract.FilterLogs(opts, "AdminEnabled")
	if err != nil {
		return nil, err
	}
	return &Erc721ContractAdminEnabledIterator{contract: _Erc721Contract.contract, event: "AdminEnabled", logs: logs, sub: sub}, nil
}

// WatchAdminEnabled is a free log subscription operation binding the contract event 0xb3714ef62726d54fcac8919290b50b999590b382adefc7f78eddf14aafb9ba5e.
//
// Solidity: event AdminEnabled(address admin, bool enabled)
func (_Erc721Contract *Erc721ContractFilterer) WatchAdminEnabled(opts *bind.WatchOpts, sink chan<- *Erc721ContractAdminEnabled) (event.Subscription, error) {

	logs, sub, err := _Erc721Contract.contract.WatchLogs(opts, "AdminEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContractAdminEnabled)
				if err := _Erc721Contract.contract.UnpackLog(event, "AdminEnabled", log); err != nil {
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

// ParseAdminEnabled is a log parse operation binding the contract event 0xb3714ef62726d54fcac8919290b50b999590b382adefc7f78eddf14aafb9ba5e.
//
// Solidity: event AdminEnabled(address admin, bool enabled)
func (_Erc721Contract *Erc721ContractFilterer) ParseAdminEnabled(log types.Log) (*Erc721ContractAdminEnabled, error) {
	event := new(Erc721ContractAdminEnabled)
	if err := _Erc721Contract.contract.UnpackLog(event, "AdminEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc721Contract contract.
type Erc721ContractApprovalIterator struct {
	Event *Erc721ContractApproval // Event containing the contract specifics and raw log

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
func (it *Erc721ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContractApproval)
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
		it.Event = new(Erc721ContractApproval)
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
func (it *Erc721ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContractApproval represents a Approval event raised by the Erc721Contract contract.
type Erc721ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721Contract *Erc721ContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*Erc721ContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721Contract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractApprovalIterator{contract: _Erc721Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721Contract *Erc721ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc721ContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721Contract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContractApproval)
				if err := _Erc721Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Erc721Contract *Erc721ContractFilterer) ParseApproval(log types.Log) (*Erc721ContractApproval, error) {
	event := new(Erc721ContractApproval)
	if err := _Erc721Contract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Erc721Contract contract.
type Erc721ContractApprovalForAllIterator struct {
	Event *Erc721ContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *Erc721ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContractApprovalForAll)
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
		it.Event = new(Erc721ContractApprovalForAll)
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
func (it *Erc721ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContractApprovalForAll represents a ApprovalForAll event raised by the Erc721Contract contract.
type Erc721ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721Contract *Erc721ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*Erc721ContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721Contract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractApprovalForAllIterator{contract: _Erc721Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721Contract *Erc721ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *Erc721ContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Erc721Contract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContractApprovalForAll)
				if err := _Erc721Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Erc721Contract *Erc721ContractFilterer) ParseApprovalForAll(log types.Log) (*Erc721ContractApprovalForAll, error) {
	event := new(Erc721ContractApprovalForAll)
	if err := _Erc721Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Erc721Contract contract.
type Erc721ContractOwnershipTransferredIterator struct {
	Event *Erc721ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Erc721ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContractOwnershipTransferred)
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
		it.Event = new(Erc721ContractOwnershipTransferred)
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
func (it *Erc721ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Erc721Contract contract.
type Erc721ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721Contract *Erc721ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Erc721ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc721Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractOwnershipTransferredIterator{contract: _Erc721Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721Contract *Erc721ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Erc721ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Erc721Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContractOwnershipTransferred)
				if err := _Erc721Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Erc721Contract *Erc721ContractFilterer) ParseOwnershipTransferred(log types.Log) (*Erc721ContractOwnershipTransferred, error) {
	event := new(Erc721ContractOwnershipTransferred)
	if err := _Erc721Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc721Contract contract.
type Erc721ContractTransferIterator struct {
	Event *Erc721ContractTransfer // Event containing the contract specifics and raw log

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
func (it *Erc721ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContractTransfer)
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
		it.Event = new(Erc721ContractTransfer)
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
func (it *Erc721ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContractTransfer represents a Transfer event raised by the Erc721Contract contract.
type Erc721ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721Contract *Erc721ContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*Erc721ContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721Contract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractTransferIterator{contract: _Erc721Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721Contract *Erc721ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc721ContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Erc721Contract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContractTransfer)
				if err := _Erc721Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Erc721Contract *Erc721ContractFilterer) ParseTransfer(log types.Log) (*Erc721ContractTransfer, error) {
	event := new(Erc721ContractTransfer)
	if err := _Erc721Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc721ContractUpdateUserIterator is returned from FilterUpdateUser and is used to iterate over the raw logs and unpacked data for UpdateUser events raised by the Erc721Contract contract.
type Erc721ContractUpdateUserIterator struct {
	Event *Erc721ContractUpdateUser // Event containing the contract specifics and raw log

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
func (it *Erc721ContractUpdateUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc721ContractUpdateUser)
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
		it.Event = new(Erc721ContractUpdateUser)
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
func (it *Erc721ContractUpdateUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc721ContractUpdateUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc721ContractUpdateUser represents a UpdateUser event raised by the Erc721Contract contract.
type Erc721ContractUpdateUser struct {
	TokenId *big.Int
	User    common.Address
	Expires uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateUser is a free log retrieval operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_Erc721Contract *Erc721ContractFilterer) FilterUpdateUser(opts *bind.FilterOpts, tokenId []*big.Int, user []common.Address) (*Erc721ContractUpdateUserIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Erc721Contract.contract.FilterLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return &Erc721ContractUpdateUserIterator{contract: _Erc721Contract.contract, event: "UpdateUser", logs: logs, sub: sub}, nil
}

// WatchUpdateUser is a free log subscription operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_Erc721Contract *Erc721ContractFilterer) WatchUpdateUser(opts *bind.WatchOpts, sink chan<- *Erc721ContractUpdateUser, tokenId []*big.Int, user []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Erc721Contract.contract.WatchLogs(opts, "UpdateUser", tokenIdRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc721ContractUpdateUser)
				if err := _Erc721Contract.contract.UnpackLog(event, "UpdateUser", log); err != nil {
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

// ParseUpdateUser is a log parse operation binding the contract event 0x4e06b4e7000e659094299b3533b47b6aa8ad048e95e872d23d1f4ee55af89cfe.
//
// Solidity: event UpdateUser(uint256 indexed tokenId, address indexed user, uint64 expires)
func (_Erc721Contract *Erc721ContractFilterer) ParseUpdateUser(log types.Log) (*Erc721ContractUpdateUser, error) {
	event := new(Erc721ContractUpdateUser)
	if err := _Erc721Contract.contract.UnpackLog(event, "UpdateUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
