// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package setting

import (
	"strings"

	"github.com/FRECNET/accounts/abi"
	"github.com/FRECNET/accounts/abi/bind"
	"github.com/FRECNET/common"
	"github.com/FRECNET/core/types"
	"github.com/FRECNET/event"
	ethereum "github.com/FRECNET"
)

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146060604052600080fd00a165627a7a72305820b9407d48ebc7efee5c9f08b3b3a957df2939281f5913225e8c1291f069b900490029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SettingABI is the input ABI used to generate the binding from.
const SettingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getTIPIncrease\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TIPIncreaseMasternodes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"uint64\"}],\"name\":\"setEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardPerEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"epochValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"uint64\"}],\"name\":\"setMaxMasterNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"uint64\"}],\"name\":\"setTIPIncrease\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardPerEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxMasterNode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxMasterNode\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"uint64\"}],\"name\":\"setRewardPerEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_maxMasterNode\",\"type\":\"uint64\"},{\"name\":\"_epochValue\",\"type\":\"uint64\"},{\"name\":\"_rewardPerEpoch\",\"type\":\"uint64\"},{\"name\":\"_TIPIncreaseMasternodes\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_modifiedby\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint64\"}],\"name\":\"epoch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_modifiedby\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint64\"}],\"name\":\"reward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_modifiedby\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint64\"}],\"name\":\"masternode\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_modifiedby\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint64\"}],\"name\":\"tipIncreaseMasterNode\",\"type\":\"event\"}]"

// SettingBin is the compiled bytecode used for deploying new contracts.
const SettingBin = `608060405234801561001057600080fd5b50604051608080610a8583398101806040528101908080519060200190929190805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600060146101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555083600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555081600160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555080600160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050505061093d806101486000396000f3006080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806307fe30d6146100ca5780630fe8dad81461010957806311bde4fb146101485780631d4b449a1461017f578063357f170e146101be578063407aed6e146101fd57806344b4f53814610234578063757991a81461026b57806384449a9d146102aa5780638da5cb5b146102e9578063c6322df214610340578063f593484a1461037f578063f74653eb146103be575b600080fd5b3480156100d657600080fd5b506100df6103f5565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b34801561011557600080fd5b5061011e610413565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b34801561015457600080fd5b5061017d600480360381019080803567ffffffffffffffff16906020019092919050505061042d565b005b34801561018b57600080fd5b50610194610533565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b3480156101ca57600080fd5b506101d3610551565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b34801561020957600080fd5b50610232600480360381019080803567ffffffffffffffff16906020019092919050505061056b565b005b34801561024057600080fd5b50610269600480360381019080803567ffffffffffffffff169060200190929190505050610671565b005b34801561027757600080fd5b50610280610777565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b3480156102b657600080fd5b506102bf610794565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b3480156102f557600080fd5b506102fe6107ae565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561034c57600080fd5b506103556107d3565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b34801561038b57600080fd5b506103946107f1565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b3480156103ca57600080fd5b506103f3600480360381019080803567ffffffffffffffff16906020019092919050505061080b565b005b6000600160109054906101000a900467ffffffffffffffff16905090565b600160109054906101000a900467ffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561048857600080fd5b80600060146101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f97e2c7ce57dcaa7fce0e9963c641494722592f51f84555230ffa81a6c3d423fd3382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018267ffffffffffffffff1667ffffffffffffffff1681526020019250505060405180910390a150565b6000600160089054906101000a900467ffffffffffffffff16905090565b600060149054906101000a900467ffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105c657600080fd5b80600160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fdb18383852c9bec12b3a9a5b53acaf6f9e02ba66efc3085bb311b8d440c7f6113382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018267ffffffffffffffff1667ffffffffffffffff1681526020019250505060405180910390a150565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156106cc57600080fd5b80600160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507f3ef0054dcca960f048d0a8cc8bde957d29764bba4ac482f678c90df8f31cc6613382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018267ffffffffffffffff1667ffffffffffffffff1681526020019250505060405180910390a150565b60008060149054906101000a900467ffffffffffffffff16905090565b600160089054906101000a900467ffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160009054906101000a900467ffffffffffffffff16905090565b600160009054906101000a900467ffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561086657600080fd5b80600160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055507fd4d044be35929480ecbaf007df7342a7fe4b676fb1a7c2ff62c5bc55fe6155923382604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018267ffffffffffffffff1667ffffffffffffffff1681526020019250505060405180910390a1505600a165627a7a723058206518f4c57fabfba6d421b195ec73e022dfad1e1fb6c0437f75988a9abec2f8100029`

// DeploySetting deploys a new Ethereum contract, binding an instance of Setting to it.
func DeploySetting(auth *bind.TransactOpts, backend bind.ContractBackend, _maxMasterNode uint64, _epochValue uint64, _rewardPerEpoch uint64, _TIPIncreaseMasternodes uint64) (common.Address, *types.Transaction, *Setting, error) {
	parsed, err := abi.JSON(strings.NewReader(SettingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SettingBin), backend, _maxMasterNode, _epochValue, _rewardPerEpoch, _TIPIncreaseMasternodes)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Setting{SettingCaller: SettingCaller{contract: contract}, SettingTransactor: SettingTransactor{contract: contract}, SettingFilterer: SettingFilterer{contract: contract}}, nil
}

// Setting is an auto generated Go binding around an Ethereum contract.
type Setting struct {
	SettingCaller     // Read-only binding to the contract
	SettingTransactor // Write-only binding to the contract
	SettingFilterer   // Log filterer for contract events
}

// SettingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SettingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SettingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SettingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SettingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SettingSession struct {
	Contract     *Setting          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SettingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SettingCallerSession struct {
	Contract *SettingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SettingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SettingTransactorSession struct {
	Contract     *SettingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SettingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SettingRaw struct {
	Contract *Setting // Generic contract binding to access the raw methods on
}

// SettingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SettingCallerRaw struct {
	Contract *SettingCaller // Generic read-only contract binding to access the raw methods on
}

// SettingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SettingTransactorRaw struct {
	Contract *SettingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSetting creates a new instance of Setting, bound to a specific deployed contract.
func NewSetting(address common.Address, backend bind.ContractBackend) (*Setting, error) {
	contract, err := bindSetting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Setting{SettingCaller: SettingCaller{contract: contract}, SettingTransactor: SettingTransactor{contract: contract}, SettingFilterer: SettingFilterer{contract: contract}}, nil
}

// NewSettingCaller creates a new read-only instance of Setting, bound to a specific deployed contract.
func NewSettingCaller(address common.Address, caller bind.ContractCaller) (*SettingCaller, error) {
	contract, err := bindSetting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SettingCaller{contract: contract}, nil
}

// NewSettingTransactor creates a new write-only instance of Setting, bound to a specific deployed contract.
func NewSettingTransactor(address common.Address, transactor bind.ContractTransactor) (*SettingTransactor, error) {
	contract, err := bindSetting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SettingTransactor{contract: contract}, nil
}

// NewSettingFilterer creates a new log filterer instance of Setting, bound to a specific deployed contract.
func NewSettingFilterer(address common.Address, filterer bind.ContractFilterer) (*SettingFilterer, error) {
	contract, err := bindSetting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SettingFilterer{contract: contract}, nil
}

// bindSetting binds a generic wrapper to an already deployed contract.
func bindSetting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SettingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Setting *SettingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Setting.Contract.SettingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Setting *SettingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Setting.Contract.SettingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Setting *SettingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Setting.Contract.SettingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Setting *SettingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Setting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Setting *SettingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Setting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Setting *SettingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Setting.Contract.contract.Transact(opts, method, params...)
}

// TIPIncreaseMasternodes is a free data retrieval call binding the contract method 0x0fe8dad8.
//
// Solidity: function TIPIncreaseMasternodes() constant returns(uint64)
func (_Setting *SettingCaller) TIPIncreaseMasternodes(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "TIPIncreaseMasternodes")
	return *ret0, err
}

// TIPIncreaseMasternodes is a free data retrieval call binding the contract method 0x0fe8dad8.
//
// Solidity: function TIPIncreaseMasternodes() constant returns(uint64)
func (_Setting *SettingSession) TIPIncreaseMasternodes() (uint64, error) {
	return _Setting.Contract.TIPIncreaseMasternodes(&_Setting.CallOpts)
}

// TIPIncreaseMasternodes is a free data retrieval call binding the contract method 0x0fe8dad8.
//
// Solidity: function TIPIncreaseMasternodes() constant returns(uint64)
func (_Setting *SettingCallerSession) TIPIncreaseMasternodes() (uint64, error) {
	return _Setting.Contract.TIPIncreaseMasternodes(&_Setting.CallOpts)
}

// EpochValue is a free data retrieval call binding the contract method 0x357f170e.
//
// Solidity: function epochValue() constant returns(uint64)
func (_Setting *SettingCaller) EpochValue(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "epochValue")
	return *ret0, err
}

// EpochValue is a free data retrieval call binding the contract method 0x357f170e.
//
// Solidity: function epochValue() constant returns(uint64)
func (_Setting *SettingSession) EpochValue() (uint64, error) {
	return _Setting.Contract.EpochValue(&_Setting.CallOpts)
}

// EpochValue is a free data retrieval call binding the contract method 0x357f170e.
//
// Solidity: function epochValue() constant returns(uint64)
func (_Setting *SettingCallerSession) EpochValue() (uint64, error) {
	return _Setting.Contract.EpochValue(&_Setting.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint64)
func (_Setting *SettingCaller) GetEpoch(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "getEpoch")
	return *ret0, err
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint64)
func (_Setting *SettingSession) GetEpoch() (uint64, error) {
	return _Setting.Contract.GetEpoch(&_Setting.CallOpts)
}

// GetEpoch is a free data retrieval call binding the contract method 0x757991a8.
//
// Solidity: function getEpoch() constant returns(uint64)
func (_Setting *SettingCallerSession) GetEpoch() (uint64, error) {
	return _Setting.Contract.GetEpoch(&_Setting.CallOpts)
}

// GetMaxMasterNode is a free data retrieval call binding the contract method 0xc6322df2.
//
// Solidity: function getMaxMasterNode() constant returns(uint64)
func (_Setting *SettingCaller) GetMaxMasterNode(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "getMaxMasterNode")
	return *ret0, err
}

// GetMaxMasterNode is a free data retrieval call binding the contract method 0xc6322df2.
//
// Solidity: function getMaxMasterNode() constant returns(uint64)
func (_Setting *SettingSession) GetMaxMasterNode() (uint64, error) {
	return _Setting.Contract.GetMaxMasterNode(&_Setting.CallOpts)
}

// GetMaxMasterNode is a free data retrieval call binding the contract method 0xc6322df2.
//
// Solidity: function getMaxMasterNode() constant returns(uint64)
func (_Setting *SettingCallerSession) GetMaxMasterNode() (uint64, error) {
	return _Setting.Contract.GetMaxMasterNode(&_Setting.CallOpts)
}

// GetRewardPerEpoch is a free data retrieval call binding the contract method 0x1d4b449a.
//
// Solidity: function getRewardPerEpoch() constant returns(uint64)
func (_Setting *SettingCaller) GetRewardPerEpoch(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "getRewardPerEpoch")
	return *ret0, err
}

// GetRewardPerEpoch is a free data retrieval call binding the contract method 0x1d4b449a.
//
// Solidity: function getRewardPerEpoch() constant returns(uint64)
func (_Setting *SettingSession) GetRewardPerEpoch() (uint64, error) {
	return _Setting.Contract.GetRewardPerEpoch(&_Setting.CallOpts)
}

// GetRewardPerEpoch is a free data retrieval call binding the contract method 0x1d4b449a.
//
// Solidity: function getRewardPerEpoch() constant returns(uint64)
func (_Setting *SettingCallerSession) GetRewardPerEpoch() (uint64, error) {
	return _Setting.Contract.GetRewardPerEpoch(&_Setting.CallOpts)
}

// GetTIPIncrease is a free data retrieval call binding the contract method 0x07fe30d6.
//
// Solidity: function getTIPIncrease() constant returns(uint64)
func (_Setting *SettingCaller) GetTIPIncrease(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "getTIPIncrease")
	return *ret0, err
}

// GetTIPIncrease is a free data retrieval call binding the contract method 0x07fe30d6.
//
// Solidity: function getTIPIncrease() constant returns(uint64)
func (_Setting *SettingSession) GetTIPIncrease() (uint64, error) {
	return _Setting.Contract.GetTIPIncrease(&_Setting.CallOpts)
}

// GetTIPIncrease is a free data retrieval call binding the contract method 0x07fe30d6.
//
// Solidity: function getTIPIncrease() constant returns(uint64)
func (_Setting *SettingCallerSession) GetTIPIncrease() (uint64, error) {
	return _Setting.Contract.GetTIPIncrease(&_Setting.CallOpts)
}

// MaxMasterNode is a free data retrieval call binding the contract method 0xf593484a.
//
// Solidity: function maxMasterNode() constant returns(uint64)
func (_Setting *SettingCaller) MaxMasterNode(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "maxMasterNode")
	return *ret0, err
}

// MaxMasterNode is a free data retrieval call binding the contract method 0xf593484a.
//
// Solidity: function maxMasterNode() constant returns(uint64)
func (_Setting *SettingSession) MaxMasterNode() (uint64, error) {
	return _Setting.Contract.MaxMasterNode(&_Setting.CallOpts)
}

// MaxMasterNode is a free data retrieval call binding the contract method 0xf593484a.
//
// Solidity: function maxMasterNode() constant returns(uint64)
func (_Setting *SettingCallerSession) MaxMasterNode() (uint64, error) {
	return _Setting.Contract.MaxMasterNode(&_Setting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Setting *SettingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Setting *SettingSession) Owner() (common.Address, error) {
	return _Setting.Contract.Owner(&_Setting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Setting *SettingCallerSession) Owner() (common.Address, error) {
	return _Setting.Contract.Owner(&_Setting.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() constant returns(uint64)
func (_Setting *SettingCaller) RewardPerEpoch(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Setting.contract.Call(opts, out, "rewardPerEpoch")
	return *ret0, err
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() constant returns(uint64)
func (_Setting *SettingSession) RewardPerEpoch() (uint64, error) {
	return _Setting.Contract.RewardPerEpoch(&_Setting.CallOpts)
}

// RewardPerEpoch is a free data retrieval call binding the contract method 0x84449a9d.
//
// Solidity: function rewardPerEpoch() constant returns(uint64)
func (_Setting *SettingCallerSession) RewardPerEpoch() (uint64, error) {
	return _Setting.Contract.RewardPerEpoch(&_Setting.CallOpts)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x11bde4fb.
//
// Solidity: function setEpoch(_val uint64) returns()
func (_Setting *SettingTransactor) SetEpoch(opts *bind.TransactOpts, _val uint64) (*types.Transaction, error) {
	return _Setting.contract.Transact(opts, "setEpoch", _val)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x11bde4fb.
//
// Solidity: function setEpoch(_val uint64) returns()
func (_Setting *SettingSession) SetEpoch(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetEpoch(&_Setting.TransactOpts, _val)
}

// SetEpoch is a paid mutator transaction binding the contract method 0x11bde4fb.
//
// Solidity: function setEpoch(_val uint64) returns()
func (_Setting *SettingTransactorSession) SetEpoch(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetEpoch(&_Setting.TransactOpts, _val)
}

// SetMaxMasterNode is a paid mutator transaction binding the contract method 0x407aed6e.
//
// Solidity: function setMaxMasterNode(_val uint64) returns()
func (_Setting *SettingTransactor) SetMaxMasterNode(opts *bind.TransactOpts, _val uint64) (*types.Transaction, error) {
	return _Setting.contract.Transact(opts, "setMaxMasterNode", _val)
}

// SetMaxMasterNode is a paid mutator transaction binding the contract method 0x407aed6e.
//
// Solidity: function setMaxMasterNode(_val uint64) returns()
func (_Setting *SettingSession) SetMaxMasterNode(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetMaxMasterNode(&_Setting.TransactOpts, _val)
}

// SetMaxMasterNode is a paid mutator transaction binding the contract method 0x407aed6e.
//
// Solidity: function setMaxMasterNode(_val uint64) returns()
func (_Setting *SettingTransactorSession) SetMaxMasterNode(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetMaxMasterNode(&_Setting.TransactOpts, _val)
}

// SetRewardPerEpoch is a paid mutator transaction binding the contract method 0xf74653eb.
//
// Solidity: function setRewardPerEpoch(_val uint64) returns()
func (_Setting *SettingTransactor) SetRewardPerEpoch(opts *bind.TransactOpts, _val uint64) (*types.Transaction, error) {
	return _Setting.contract.Transact(opts, "setRewardPerEpoch", _val)
}

// SetRewardPerEpoch is a paid mutator transaction binding the contract method 0xf74653eb.
//
// Solidity: function setRewardPerEpoch(_val uint64) returns()
func (_Setting *SettingSession) SetRewardPerEpoch(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetRewardPerEpoch(&_Setting.TransactOpts, _val)
}

// SetRewardPerEpoch is a paid mutator transaction binding the contract method 0xf74653eb.
//
// Solidity: function setRewardPerEpoch(_val uint64) returns()
func (_Setting *SettingTransactorSession) SetRewardPerEpoch(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetRewardPerEpoch(&_Setting.TransactOpts, _val)
}

// SetTIPIncrease is a paid mutator transaction binding the contract method 0x44b4f538.
//
// Solidity: function setTIPIncrease(_val uint64) returns()
func (_Setting *SettingTransactor) SetTIPIncrease(opts *bind.TransactOpts, _val uint64) (*types.Transaction, error) {
	return _Setting.contract.Transact(opts, "setTIPIncrease", _val)
}

// SetTIPIncrease is a paid mutator transaction binding the contract method 0x44b4f538.
//
// Solidity: function setTIPIncrease(_val uint64) returns()
func (_Setting *SettingSession) SetTIPIncrease(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetTIPIncrease(&_Setting.TransactOpts, _val)
}

// SetTIPIncrease is a paid mutator transaction binding the contract method 0x44b4f538.
//
// Solidity: function setTIPIncrease(_val uint64) returns()
func (_Setting *SettingTransactorSession) SetTIPIncrease(_val uint64) (*types.Transaction, error) {
	return _Setting.Contract.SetTIPIncrease(&_Setting.TransactOpts, _val)
}

// SettingEpochIterator is returned from FilterEpoch and is used to iterate over the raw logs and unpacked data for Epoch events raised by the Setting contract.
type SettingEpochIterator struct {
	Event *SettingEpoch // Event containing the contract specifics and raw log

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
func (it *SettingEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettingEpoch)
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
		it.Event = new(SettingEpoch)
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
func (it *SettingEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettingEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettingEpoch represents a Epoch event raised by the Setting contract.
type SettingEpoch struct {
	Modifiedby common.Address
	Value      uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEpoch is a free log retrieval operation binding the contract event 0x97e2c7ce57dcaa7fce0e9963c641494722592f51f84555230ffa81a6c3d423fd.
//
// Solidity: event epoch(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) FilterEpoch(opts *bind.FilterOpts) (*SettingEpochIterator, error) {

	logs, sub, err := _Setting.contract.FilterLogs(opts, "epoch")
	if err != nil {
		return nil, err
	}
	return &SettingEpochIterator{contract: _Setting.contract, event: "epoch", logs: logs, sub: sub}, nil
}

// WatchEpoch is a free log subscription operation binding the contract event 0x97e2c7ce57dcaa7fce0e9963c641494722592f51f84555230ffa81a6c3d423fd.
//
// Solidity: event epoch(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) WatchEpoch(opts *bind.WatchOpts, sink chan<- *SettingEpoch) (event.Subscription, error) {

	logs, sub, err := _Setting.contract.WatchLogs(opts, "epoch")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettingEpoch)
				if err := _Setting.contract.UnpackLog(event, "epoch", log); err != nil {
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

// SettingMasternodeIterator is returned from FilterMasternode and is used to iterate over the raw logs and unpacked data for Masternode events raised by the Setting contract.
type SettingMasternodeIterator struct {
	Event *SettingMasternode // Event containing the contract specifics and raw log

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
func (it *SettingMasternodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettingMasternode)
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
		it.Event = new(SettingMasternode)
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
func (it *SettingMasternodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettingMasternodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettingMasternode represents a Masternode event raised by the Setting contract.
type SettingMasternode struct {
	Modifiedby common.Address
	Value      uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMasternode is a free log retrieval operation binding the contract event 0xdb18383852c9bec12b3a9a5b53acaf6f9e02ba66efc3085bb311b8d440c7f611.
//
// Solidity: event masternode(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) FilterMasternode(opts *bind.FilterOpts) (*SettingMasternodeIterator, error) {

	logs, sub, err := _Setting.contract.FilterLogs(opts, "masternode")
	if err != nil {
		return nil, err
	}
	return &SettingMasternodeIterator{contract: _Setting.contract, event: "masternode", logs: logs, sub: sub}, nil
}

// WatchMasternode is a free log subscription operation binding the contract event 0xdb18383852c9bec12b3a9a5b53acaf6f9e02ba66efc3085bb311b8d440c7f611.
//
// Solidity: event masternode(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) WatchMasternode(opts *bind.WatchOpts, sink chan<- *SettingMasternode) (event.Subscription, error) {

	logs, sub, err := _Setting.contract.WatchLogs(opts, "masternode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettingMasternode)
				if err := _Setting.contract.UnpackLog(event, "masternode", log); err != nil {
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

// SettingRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the Setting contract.
type SettingRewardIterator struct {
	Event *SettingReward // Event containing the contract specifics and raw log

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
func (it *SettingRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettingReward)
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
		it.Event = new(SettingReward)
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
func (it *SettingRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettingRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettingReward represents a Reward event raised by the Setting contract.
type SettingReward struct {
	Modifiedby common.Address
	Value      uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0xd4d044be35929480ecbaf007df7342a7fe4b676fb1a7c2ff62c5bc55fe615592.
//
// Solidity: event reward(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) FilterReward(opts *bind.FilterOpts) (*SettingRewardIterator, error) {

	logs, sub, err := _Setting.contract.FilterLogs(opts, "reward")
	if err != nil {
		return nil, err
	}
	return &SettingRewardIterator{contract: _Setting.contract, event: "reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0xd4d044be35929480ecbaf007df7342a7fe4b676fb1a7c2ff62c5bc55fe615592.
//
// Solidity: event reward(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *SettingReward) (event.Subscription, error) {

	logs, sub, err := _Setting.contract.WatchLogs(opts, "reward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettingReward)
				if err := _Setting.contract.UnpackLog(event, "reward", log); err != nil {
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

// SettingTipIncreaseMasterNodeIterator is returned from FilterTipIncreaseMasterNode and is used to iterate over the raw logs and unpacked data for TipIncreaseMasterNode events raised by the Setting contract.
type SettingTipIncreaseMasterNodeIterator struct {
	Event *SettingTipIncreaseMasterNode // Event containing the contract specifics and raw log

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
func (it *SettingTipIncreaseMasterNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SettingTipIncreaseMasterNode)
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
		it.Event = new(SettingTipIncreaseMasterNode)
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
func (it *SettingTipIncreaseMasterNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SettingTipIncreaseMasterNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SettingTipIncreaseMasterNode represents a TipIncreaseMasterNode event raised by the Setting contract.
type SettingTipIncreaseMasterNode struct {
	Modifiedby common.Address
	Value      uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTipIncreaseMasterNode is a free log retrieval operation binding the contract event 0x3ef0054dcca960f048d0a8cc8bde957d29764bba4ac482f678c90df8f31cc661.
//
// Solidity: event tipIncreaseMasterNode(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) FilterTipIncreaseMasterNode(opts *bind.FilterOpts) (*SettingTipIncreaseMasterNodeIterator, error) {

	logs, sub, err := _Setting.contract.FilterLogs(opts, "tipIncreaseMasterNode")
	if err != nil {
		return nil, err
	}
	return &SettingTipIncreaseMasterNodeIterator{contract: _Setting.contract, event: "tipIncreaseMasterNode", logs: logs, sub: sub}, nil
}

// WatchTipIncreaseMasterNode is a free log subscription operation binding the contract event 0x3ef0054dcca960f048d0a8cc8bde957d29764bba4ac482f678c90df8f31cc661.
//
// Solidity: event tipIncreaseMasterNode(_modifiedby address, _value uint64)
func (_Setting *SettingFilterer) WatchTipIncreaseMasterNode(opts *bind.WatchOpts, sink chan<- *SettingTipIncreaseMasterNode) (event.Subscription, error) {

	logs, sub, err := _Setting.contract.WatchLogs(opts, "tipIncreaseMasterNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SettingTipIncreaseMasterNode)
				if err := _Setting.contract.UnpackLog(event, "tipIncreaseMasterNode", log); err != nil {
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
