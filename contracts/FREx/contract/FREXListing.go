// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/FRECNET/accounts/abi"
	"github.com/FRECNET/accounts/abi/bind"
	"github.com/FRECNET/common"
	"github.com/FRECNET/core/types"
)

// FREXListingABI is the input ABI used to generate the binding from.
const FREXListingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"apply\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// FREXListingBin is the compiled bytecode used for deploying new contracts.
const FREXListingBin = `0x608060405234801561001057600080fd5b506102be806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416639d63848a811461005b578063a3ff31b5146100c0578063c6b32f34146100f5575b600080fd5b34801561006757600080fd5b5061007061010b565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100ac578181015183820152602001610094565b505050509050019250505060405180910390f35b3480156100cc57600080fd5b506100e1600160a060020a036004351661016d565b604080519115158252519081900360200190f35b610109600160a060020a036004351661018b565b005b6060600080548060200260200160405190810160405280929190818152602001828054801561016357602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610145575b5050505050905090565b600160a060020a031660009081526001602052604090205460ff1690565b80600160a060020a03811615156101a157600080fd5b600160a060020a03811660009081526001602081905260409091205460ff16151514156101cd57600080fd5b683635c9adc5dea0000034146101e257600080fd5b6040516068903480156108fc02916000818181858888f1935050505015801561020f573d6000803e3d6000fd5b505060008054600180820183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563909101805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039490941693841790556040805160208082018352838252948452919093529190209051815460ff19169015151790555600a165627a7a723058206d2dc0ce827743c25efa82f99e7830ade39d28e17f4d651573f89e0460a6626a0029`

// DeployFREXListing deploys a new Ethereum contract, binding an instance of FREXListing to it.
func DeployFREXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FREXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(FREXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FREXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FREXListing{FREXListingCaller: FREXListingCaller{contract: contract}, FREXListingTransactor: FREXListingTransactor{contract: contract}, FREXListingFilterer: FREXListingFilterer{contract: contract}}, nil
}

// FREXListing is an auto generated Go binding around an Ethereum contract.
type FREXListing struct {
	FREXListingCaller     // Read-only binding to the contract
	FREXListingTransactor // Write-only binding to the contract
	FREXListingFilterer   // Log filterer for contract events
}

// FREXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type FREXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FREXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FREXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FREXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FREXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FREXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FREXListingSession struct {
	Contract     *FREXListing      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FREXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FREXListingCallerSession struct {
	Contract *FREXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FREXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FREXListingTransactorSession struct {
	Contract     *FREXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FREXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type FREXListingRaw struct {
	Contract *FREXListing // Generic contract binding to access the raw methods on
}

// FREXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FREXListingCallerRaw struct {
	Contract *FREXListingCaller // Generic read-only contract binding to access the raw methods on
}

// FREXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FREXListingTransactorRaw struct {
	Contract *FREXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFREXListing creates a new instance of FREXListing, bound to a specific deployed contract.
func NewFREXListing(address common.Address, backend bind.ContractBackend) (*FREXListing, error) {
	contract, err := bindFREXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FREXListing{FREXListingCaller: FREXListingCaller{contract: contract}, FREXListingTransactor: FREXListingTransactor{contract: contract}, FREXListingFilterer: FREXListingFilterer{contract: contract}}, nil
}

// NewFREXListingCaller creates a new read-only instance of FREXListing, bound to a specific deployed contract.
func NewFREXListingCaller(address common.Address, caller bind.ContractCaller) (*FREXListingCaller, error) {
	contract, err := bindFREXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FREXListingCaller{contract: contract}, nil
}

// NewFREXListingTransactor creates a new write-only instance of FREXListing, bound to a specific deployed contract.
func NewFREXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*FREXListingTransactor, error) {
	contract, err := bindFREXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FREXListingTransactor{contract: contract}, nil
}

// NewFREXListingFilterer creates a new log filterer instance of FREXListing, bound to a specific deployed contract.
func NewFREXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*FREXListingFilterer, error) {
	contract, err := bindFREXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FREXListingFilterer{contract: contract}, nil
}

// bindFREXListing binds a generic wrapper to an already deployed contract.
func bindFREXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FREXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FREXListing *FREXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FREXListing.Contract.FREXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FREXListing *FREXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FREXListing.Contract.FREXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FREXListing *FREXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FREXListing.Contract.FREXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FREXListing *FREXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FREXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FREXListing *FREXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FREXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FREXListing *FREXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FREXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_FREXListing *FREXListingCaller) GetTokenStatus(opts *bind.CallOpts, token common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FREXListing.contract.Call(opts, out, "getTokenStatus", token)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_FREXListing *FREXListingSession) GetTokenStatus(token common.Address) (bool, error) {
	return _FREXListing.Contract.GetTokenStatus(&_FREXListing.CallOpts, token)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus(token address) constant returns(bool)
func (_FREXListing *FREXListingCallerSession) GetTokenStatus(token common.Address) (bool, error) {
	return _FREXListing.Contract.GetTokenStatus(&_FREXListing.CallOpts, token)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_FREXListing *FREXListingCaller) Tokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FREXListing.contract.Call(opts, out, "tokens")
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_FREXListing *FREXListingSession) Tokens() ([]common.Address, error) {
	return _FREXListing.Contract.Tokens(&_FREXListing.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() constant returns(address[])
func (_FREXListing *FREXListingCallerSession) Tokens() ([]common.Address, error) {
	return _FREXListing.Contract.Tokens(&_FREXListing.CallOpts)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_FREXListing *FREXListingTransactor) Apply(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _FREXListing.contract.Transact(opts, "apply", token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_FREXListing *FREXListingSession) Apply(token common.Address) (*types.Transaction, error) {
	return _FREXListing.Contract.Apply(&_FREXListing.TransactOpts, token)
}

// Apply is a paid mutator transaction binding the contract method 0xc6b32f34.
//
// Solidity: function apply(token address) returns()
func (_FREXListing *FREXListingTransactorSession) Apply(token common.Address) (*types.Transaction, error) {
	return _FREXListing.Contract.Apply(&_FREXListing.TransactOpts, token)
}
