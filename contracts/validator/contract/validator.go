// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"
	"fmt"

	ethereum "github.com/FRECNET"
	"github.com/FRECNET/accounts/abi"
	"github.com/FRECNET/accounts/abi/bind"
	"github.com/FRECNET/common"
	"github.com/FRECNET/core/types"
	"github.com/FRECNET/event"
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

// FREValidatorABI is the input ABI used to generate the binding from.
const FREValidatorABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"propose\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"unvote\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCandidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasVotedInvalid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getWithdrawCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerToCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getVoters\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getWithdrawBlockNumbers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getLatestKYC\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"candidates\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_invalidCandidate\",\"type\":\"address\"}],\"name\":\"invalidPercent\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"KYCString\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"invalidKYCCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"voterWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"resign\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"getCandidateOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getHashCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"candidateWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"isCandidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minCandidateCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnerCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_invalidCandidate\",\"type\":\"address\"}],\"name\":\"voteInvalidKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kychash\",\"type\":\"string\"}],\"name\":\"uploadKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minVoterCap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_candidates\",\"type\":\"address[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\"},{\"name\":\"_firstOwner\",\"type\":\"address\"},{\"name\":\"_minCandidateCap\",\"type\":\"uint256\"},{\"name\":\"_minVoterCap\",\"type\":\"uint256\"},{\"name\":\"_maxValidatorNumber\",\"type\":\"uint256\"},{\"name\":\"_candidateWithdrawDelay\",\"type\":\"uint256\"},{\"name\":\"_voterWithdrawDelay\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_voter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Unvote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_candidate\",\"type\":\"address\"}],\"name\":\"Resign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_cap\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kycHash\",\"type\":\"string\"}],\"name\":\"UploadedKYC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_masternodeOwner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_masternodes\",\"type\":\"address[]\"}],\"name\":\"InvalidatedNode\",\"type\":\"event\"}]"

// FREValidatorBin is the compiled bytecode used for deploying new contracts.
const FREValidatorBin = `606060405260006009556000600a5534156200001a57600080fd5b60405162003dcc38038062003dcc83398101604052808051820191906020018051820191906020018051906020019091908051906020019091908051906020019091908051906020019091908051906020019091908051906020019091905050600085600b8190555084600c8190555083600d8190555082600e8190555081600f81905550885160098190555060078054806001018281620000bd9190620004f1565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a60008154809291906001019190505550600090505b8851811015620004e25760088054806001018281620001439190620004f1565b916000526020600020900160008b848151811015156200015f57fe5b90602001906020020151909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506060604051908101604052808873ffffffffffffffffffffffffffffffffffffffff1681526020016001151581526020018983815181101515620001ea57fe5b90602001906020020151815250600160008b848151811015156200020a57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160010155905050600260008a83815181101515620002d557fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060010182816200032d9190620004f1565b9160005260206000209001600089909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281620003cf9190620004f1565b916000526020600020900160008b84815181101515620003eb57fe5b90602001906020020151909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600b54600160008b848151811015156200044c57fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550808060010191505062000123565b50505050505050505062000548565b8154818355818115116200051b578183600052602060002091820191016200051a919062000520565b5b505050565b6200054591905b808211156200054157600081600090555060010162000527565b5090565b90565b61387480620005586000396000f300606060405260043610610196576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063012679511461019b578063025e7c27146101c957806302aa9be21461022c57806306a49fce1461026e5780630db02622146102d85780630e3e4fb81461030157806315febd68146103715780632a3640b1146103a85780632d15cc041461042a5780632f9c4bba146104b8578063302b687214610522578063326586521461058e5780633477ee2e14610640578063441a3e70146106a357806358e7525f146106cf5780635b860d271461071c5780635b9cd8cc146107695780636dd7d8ea1461082457806372e44a3814610852578063a9a981a31461089f578063a9ff959e146108c8578063ae6e43f5146108f1578063b642facd1461092a578063c45607df146109a3578063d09f1ab4146109f0578063d161c76714610a19578063d51b9e9314610a42578063d55b7dff14610a93578063ef18374a14610abc578063f2ee3c7d14610ae5578063f5c9512514610b1e578063f8ac9dd514610b4c575b600080fd5b6101c7600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b75565b005b34156101d457600080fd5b6101ea60048080359060200190919050506111fc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561023757600080fd5b61026c600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061123b565b005b341561027957600080fd5b610281611796565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102c45780820151818401526020810190506102a9565b505050509050019250505060405180910390f35b34156102e357600080fd5b6102eb61182a565b6040518082815260200191505060405180910390f35b341561030c57600080fd5b610357600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611830565b604051808215151515815260200191505060405180910390f35b341561037c57600080fd5b610392600480803590602001909190505061185f565b6040518082815260200191505060405180910390f35b34156103b357600080fd5b6103e8600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506118bb565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561043557600080fd5b610461600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611909565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156104a4578082015181840152602081019050610489565b505050509050019250505060405180910390f35b34156104c357600080fd5b6104cb6119dc565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561050e5780820151818401526020810190506104f3565b505050509050019250505060405180910390f35b341561052d57600080fd5b610578600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611a79565b6040518082815260200191505060405180910390f35b341561059957600080fd5b6105c5600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611b03565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106055780820151818401526020810190506105ea565b50505050905090810190601f1680156106325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561064b57600080fd5b6106616004808035906020019091905050611da2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156106ae57600080fd5b6106cd6004808035906020019091908035906020019091905050611de1565b005b34156106da57600080fd5b610706600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061208d565b6040518082815260200191505060405180910390f35b341561072757600080fd5b610753600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506120d9565b6040518082815260200191505060405180910390f35b341561077457600080fd5b6107a9600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506121a1565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156107e95780820151818401526020810190506107ce565b50505050905090810190601f1680156108165780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610850600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061226a565b005b341561085d57600080fd5b610889600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612653565b6040518082815260200191505060405180910390f35b34156108aa57600080fd5b6108b261266b565b6040518082815260200191505060405180910390f35b34156108d357600080fd5b6108db612671565b6040518082815260200191505060405180910390f35b34156108fc57600080fd5b610928600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612677565b005b341561093557600080fd5b610961600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612c36565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156109ae57600080fd5b6109da600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612ca2565b6040518082815260200191505060405180910390f35b34156109fb57600080fd5b610a03612cee565b6040518082815260200191505060405180910390f35b3415610a2457600080fd5b610a2c612cf4565b6040518082815260200191505060405180910390f35b3415610a4d57600080fd5b610a79600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612cfa565b604051808215151515815260200191505060405180910390f35b3415610a9e57600080fd5b610aa6612d53565b6040518082815260200191505060405180910390f35b3415610ac757600080fd5b610acf612d59565b6040518082815260200191505060405180910390f35b3415610af057600080fd5b610b1c600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050612d63565b005b3415610b2957600080fd5b610b4a600480803590602001908201803590602001919091929050506134f1565b005b3415610b5757600080fd5b610b5f6135f0565b6040518082815260200191505060405180910390f35b6000600b543410151515610b8857600080fd5b6000600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050141580610c1c57506000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050115b1515610c2757600080fd5b81600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16151515610c8457600080fd5b610cd934600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101546135f690919063ffffffff16565b915060088054806001018281610cef919061362d565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506060604051908101604052803373ffffffffffffffffffffffffffffffffffffffff16815260200160011515815260200183815250600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff02191690831515021790555060408201518160010155905050610eb834600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546135f690919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610f5160016009546135f690919063ffffffff16565b6009819055506000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905014156110185760078054806001018281610fb6919061362d565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a600081548092919060010191905055505b600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281611069919061362d565b9160005260206000209001600085909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806001018281611109919061362d565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1338434604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b60078181548110151561120b57fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000828280600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101515156112cd57600080fd5b3373ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561140657600b546113f882600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461361490919063ffffffff16565b1015151561140557600080fd5b5b61145b84600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461361490919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555061153384600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461361490919063ffffffff16565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506115cb43600f546135f690919063ffffffff16565b9250611632846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000868152602001908152602001600020546135f690919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000858152602001908152602001600020819055506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010180548060010182816116db9190613659565b9160005260206000209001600085909190915055507faa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2338686604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050505050565b61179e613685565b600880548060200260200160405190810160405280929190818152602001828054801561182057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116117d6575b5050505050905090565b600a5481565b60056020528160005260406000206020528060005260406000206000915091509054906101000a900460ff1681565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000838152602001908152602001600020549050919050565b6006602052816000526040600020818154811015156118d657fe5b90600052602060002090016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611911613685565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156119d057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611986575b50505050509050919050565b6119e4613699565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101805480602002602001604051908101604052809291908181526020018280548015611a6f57602002820191906000526020600020905b815481526020019060010190808311611a5b575b5050505050905090565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b611b0b6136ad565b611b1482612cfa565b15611c655760036000611b2684612c36565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600160036000611b6f86612c36565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905003815481101515611bba57fe5b90600052602060002090018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c595780601f10611c2e57610100808354040283529160200191611c59565b820191906000526020600020905b815481529060010190602001808311611c3c57829003601f168201915b50505050509050611d9d565b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905003815481101515611cf657fe5b90600052602060002090018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611d955780601f10611d6a57610100808354040283529160200191611d95565b820191906000526020600020905b815481529060010190602001808311611d7857829003601f168201915b505050505090505b919050565b600881815481101515611db157fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008282600082111515611df457600080fd5b814310151515611e0357600080fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600084815260200190815260200160002054111515611e6457600080fd5b816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010182815481101515611eb357fe5b906000526020600020900154141515611ecb57600080fd5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008681526020019081526020016000205492506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000868152602001908152602001600020600090556000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010184815481101515611fc457fe5b9060005260206000209001600090553373ffffffffffffffffffffffffffffffffffffffff166108fc849081150290604051600060405180830381858888f19350505050151561201357600080fd5b7ff279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568338685604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001935050505060405180910390a15050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101549050919050565b60008082600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16151561213857600080fd5b61214184612c36565b915061214b612d59565b6064600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540281151561219757fe5b0492505050919050565b6003602052816000526040600020818154811015156121bc57fe5b9060005260206000209001600091509150508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156122625780601f1061223757610100808354040283529160200191612262565b820191906000526020600020905b81548152906001019060200180831161224557829003601f168201915b505050505081565b600c54341015151561227b57600080fd5b80600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff1615156122d757600080fd5b61232c34600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101546135f690919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054141561249b57600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480600101828161244b919061362d565b9160005260206000209001600033909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b61252d34600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546135f690919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc338334604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b60046020528060005260406000206000915090505481565b60095481565b600f5481565b6000806000833373ffffffffffffffffffffffffffffffffffffffff16600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561271957600080fd5b84600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff16151561277557600080fd5b6000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160146101000a81548160ff0219169083151502179055506127e6600160095461361490919063ffffffff16565b600981905550600094505b6008805490508510156128bb578573ffffffffffffffffffffffffffffffffffffffff1660088681548110151561282457fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156128ae5760088581548110151561287b57fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556128bb565b84806001019550506127f1565b600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054935061299284600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001015461361490919063ffffffff16565b600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101819055506000600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550612a7243600e546135f690919063ffffffff16565b9250612ad9846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000868152602001908152602001600020546135f690919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000016000858152602001908152602001600020819055506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018054806001018281612b829190613659565b9160005260206000209001600085909190915055507f4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d33387604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390a1505050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490509050919050565b600d5481565b600e5481565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff169050919050565b600b5481565b6000600a54905090565b600080612d6e613685565b600080600033600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161515612dcf57600080fd5b87600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160149054906101000a900460ff161515612e2b57600080fd5b612e3433612c36565b9750612e3f89612c36565b9650600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151515612ed757600080fd5b6001600560008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506001600460008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282540192505081905550604b612fc4612d59565b6064600460008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020540281151561301057fe5b041015156134e65760016008805490500360405180591061302e5750595b9080825280602002602001820160405250955060009450600093505b600880549050841015613357578673ffffffffffffffffffffffffffffffffffffffff166130b160088681548110151561308057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612c36565b73ffffffffffffffffffffffffffffffffffffffff16141561334a576130e3600160095461361490919063ffffffff16565b6009819055506008848154811015156130f857fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16868680600101975081518110151561313857fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505060088481548110151561318357fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600160006008868154811015156131c457fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a81549060ff021916905560018201600090555050600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006132bb91906136c1565b600660008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061330691906136e2565b600460008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600090555b838060010194505061304a565b600092505b600780549050831015613439578673ffffffffffffffffffffffffffffffffffffffff1660078481548110151561338f57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561342c576007838154811015156133e657fe5b906000526020600020900160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600a6000815480929190600190039190505550613439565b828060010193505061335c565b7fe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e8787604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156134d15780820151818401526020810190506134b6565b50505050905001935050505060405180910390a15b505050505050505050565b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060010182816135429190613703565b91600052602060002090016000848490919290919250919061356592919061372f565b50507f949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a338383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001806020018281038252848482818152602001925080828437820191505094505050505060405180910390a15050565b600c5481565b600080828401905083811015151561360a57fe5b8091505092915050565b600082821115151561362257fe5b818303905092915050565b8154818355818115116136545781836000526020600020918201910161365391906137af565b5b505050565b8154818355818115116136805781836000526020600020918201910161367f91906137af565b5b505050565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b602060405190810160405280600081525090565b50805460008255906000526020600020908101906136df91906137d4565b50565b508054600082559060005260206000209081019061370091906137af565b50565b81548183558181151161372a5781836000526020600020918201910161372991906137d4565b5b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061377057803560ff191683800117855561379e565b8280016001018555821561379e579182015b8281111561379d578235825591602001919060010190613782565b5b5090506137ab91906137af565b5090565b6137d191905b808211156137cd5760008160009055506001016137b5565b5090565b90565b6137fd91905b808211156137f957600081816137f09190613800565b506001016137da565b5090565b90565b50805460018160011615610100020316600290046000825580601f106138265750613845565b601f01602090049060005260206000209081019061384491906137af565b5b505600a165627a7a72305820f5bbb127b52ce86c873faef85cff176563476a5e49a3d88eaa9a06a8f432c9080029`

// DeployFREValidator deploys a new Ethereum contract, binding an instance of FREValidator to it.
func DeployFREValidator(auth *bind.TransactOpts, backend bind.ContractBackend, _candidates []common.Address, _caps []*big.Int, _firstOwner common.Address, _minCandidateCap *big.Int, _minVoterCap *big.Int, _maxValidatorNumber *big.Int, _candidateWithdrawDelay *big.Int, _voterWithdrawDelay *big.Int) (common.Address, *types.Transaction, *FREValidator, error) {
	parsed, err := abi.JSON(strings.NewReader(FREValidatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FREValidatorBin), backend, _candidates, _caps, _firstOwner, _minCandidateCap, _minVoterCap, _maxValidatorNumber, _candidateWithdrawDelay, _voterWithdrawDelay)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FREValidator{FREValidatorCaller: FREValidatorCaller{contract: contract}, FREValidatorTransactor: FREValidatorTransactor{contract: contract}, FREValidatorFilterer: FREValidatorFilterer{contract: contract}}, nil
}

// FREValidator is an auto generated Go binding around an Ethereum contract.
type FREValidator struct {
	FREValidatorCaller     // Read-only binding to the contract
	FREValidatorTransactor // Write-only binding to the contract
	FREValidatorFilterer   // Log filterer for contract events
}

// FREValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type FREValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FREValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FREValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FREValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FREValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FREValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FREValidatorSession struct {
	Contract     *FREValidator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FREValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FREValidatorCallerSession struct {
	Contract *FREValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FREValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FREValidatorTransactorSession struct {
	Contract     *FREValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FREValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type FREValidatorRaw struct {
	Contract *FREValidator // Generic contract binding to access the raw methods on
}

// FREValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FREValidatorCallerRaw struct {
	Contract *FREValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// FREValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FREValidatorTransactorRaw struct {
	Contract *FREValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFREValidator creates a new instance of FREValidator, bound to a specific deployed contract.
func NewFREValidator(address common.Address, backend bind.ContractBackend) (*FREValidator, error) {
	contract, err := bindFREValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FREValidator{FREValidatorCaller: FREValidatorCaller{contract: contract}, FREValidatorTransactor: FREValidatorTransactor{contract: contract}, FREValidatorFilterer: FREValidatorFilterer{contract: contract}}, nil
}

// NewFREValidatorCaller creates a new read-only instance of FREValidator, bound to a specific deployed contract.
func NewFREValidatorCaller(address common.Address, caller bind.ContractCaller) (*FREValidatorCaller, error) {
	contract, err := bindFREValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FREValidatorCaller{contract: contract}, nil
}

// NewFREValidatorTransactor creates a new write-only instance of FREValidator, bound to a specific deployed contract.
func NewFREValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*FREValidatorTransactor, error) {
	contract, err := bindFREValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FREValidatorTransactor{contract: contract}, nil
}

// NewFREValidatorFilterer creates a new log filterer instance of FREValidator, bound to a specific deployed contract.
func NewFREValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*FREValidatorFilterer, error) {
	contract, err := bindFREValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FREValidatorFilterer{contract: contract}, nil
}

// bindFREValidator binds a generic wrapper to an already deployed contract.
func bindFREValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FREValidatorABI))
	fmt.Println("bindFREValidator:::",parsed)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FREValidator *FREValidatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FREValidator.Contract.FREValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FREValidator *FREValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FREValidator.Contract.FREValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FREValidator *FREValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FREValidator.Contract.FREValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FREValidator *FREValidatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _FREValidator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FREValidator *FREValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FREValidator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FREValidator *FREValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FREValidator.Contract.contract.Transact(opts, method, params...)
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_FREValidator *FREValidatorCaller) KYCString(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "KYCString", arg0, arg1)
	return *ret0, err
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_FREValidator *FREValidatorSession) KYCString(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _FREValidator.Contract.KYCString(&_FREValidator.CallOpts, arg0, arg1)
}

// KYCString is a free data retrieval call binding the contract method 0x5b9cd8cc.
//
// Solidity: function KYCString( address,  uint256) constant returns(string)
func (_FREValidator *FREValidatorCallerSession) KYCString(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _FREValidator.Contract.KYCString(&_FREValidator.CallOpts, arg0, arg1)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) CandidateCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "candidateCount")
	return *ret0, err
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_FREValidator *FREValidatorSession) CandidateCount() (*big.Int, error) {
	return _FREValidator.Contract.CandidateCount(&_FREValidator.CallOpts)
}

// CandidateCount is a free data retrieval call binding the contract method 0xa9a981a3.
//
// Solidity: function candidateCount() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) CandidateCount() (*big.Int, error) {
	return _FREValidator.Contract.CandidateCount(&_FREValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) CandidateWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "candidateWithdrawDelay")
	return *ret0, err
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_FREValidator *FREValidatorSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _FREValidator.Contract.CandidateWithdrawDelay(&_FREValidator.CallOpts)
}

// CandidateWithdrawDelay is a free data retrieval call binding the contract method 0xd161c767.
//
// Solidity: function candidateWithdrawDelay() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) CandidateWithdrawDelay() (*big.Int, error) {
	return _FREValidator.Contract.CandidateWithdrawDelay(&_FREValidator.CallOpts)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_FREValidator *FREValidatorCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "candidates", arg0)
	return *ret0, err
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_FREValidator *FREValidatorSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _FREValidator.Contract.Candidates(&_FREValidator.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates( uint256) constant returns(address)
func (_FREValidator *FREValidatorCallerSession) Candidates(arg0 *big.Int) (common.Address, error) {
	return _FREValidator.Contract.Candidates(&_FREValidator.CallOpts, arg0)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_FREValidator *FREValidatorCaller) GetCandidateCap(opts *bind.CallOpts, _candidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getCandidateCap", _candidate)
	return *ret0, err
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_FREValidator *FREValidatorSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _FREValidator.Contract.GetCandidateCap(&_FREValidator.CallOpts, _candidate)
}

// GetCandidateCap is a free data retrieval call binding the contract method 0x58e7525f.
//
// Solidity: function getCandidateCap(_candidate address) constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) GetCandidateCap(_candidate common.Address) (*big.Int, error) {
	return _FREValidator.Contract.GetCandidateCap(&_FREValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_FREValidator *FREValidatorCaller) GetCandidateOwner(opts *bind.CallOpts, _candidate common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getCandidateOwner", _candidate)
	return *ret0, err
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_FREValidator *FREValidatorSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _FREValidator.Contract.GetCandidateOwner(&_FREValidator.CallOpts, _candidate)
}

// GetCandidateOwner is a free data retrieval call binding the contract method 0xb642facd.
//
// Solidity: function getCandidateOwner(_candidate address) constant returns(address)
func (_FREValidator *FREValidatorCallerSession) GetCandidateOwner(_candidate common.Address) (common.Address, error) {
	return _FREValidator.Contract.GetCandidateOwner(&_FREValidator.CallOpts, _candidate)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_FREValidator *FREValidatorCaller) GetCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getCandidates")
	return *ret0, err
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_FREValidator *FREValidatorSession) GetCandidates() ([]common.Address, error) {
	return _FREValidator.Contract.GetCandidates(&_FREValidator.CallOpts)
}

// GetCandidates is a free data retrieval call binding the contract method 0x06a49fce.
//
// Solidity: function getCandidates() constant returns(address[])
func (_FREValidator *FREValidatorCallerSession) GetCandidates() ([]common.Address, error) {
	return _FREValidator.Contract.GetCandidates(&_FREValidator.CallOpts)
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_FREValidator *FREValidatorCaller) GetHashCount(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getHashCount", _address)
	return *ret0, err
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_FREValidator *FREValidatorSession) GetHashCount(_address common.Address) (*big.Int, error) {
	return _FREValidator.Contract.GetHashCount(&_FREValidator.CallOpts, _address)
}

// GetHashCount is a free data retrieval call binding the contract method 0xc45607df.
//
// Solidity: function getHashCount(_address address) constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) GetHashCount(_address common.Address) (*big.Int, error) {
	return _FREValidator.Contract.GetHashCount(&_FREValidator.CallOpts, _address)
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_FREValidator *FREValidatorCaller) GetLatestKYC(opts *bind.CallOpts, _address common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getLatestKYC", _address)
	return *ret0, err
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_FREValidator *FREValidatorSession) GetLatestKYC(_address common.Address) (string, error) {
	return _FREValidator.Contract.GetLatestKYC(&_FREValidator.CallOpts, _address)
}

// GetLatestKYC is a free data retrieval call binding the contract method 0x32658652.
//
// Solidity: function getLatestKYC(_address address) constant returns(string)
func (_FREValidator *FREValidatorCallerSession) GetLatestKYC(_address common.Address) (string, error) {
	return _FREValidator.Contract.GetLatestKYC(&_FREValidator.CallOpts, _address)
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) GetOwnerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getOwnerCount")
	return *ret0, err
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_FREValidator *FREValidatorSession) GetOwnerCount() (*big.Int, error) {
	return _FREValidator.Contract.GetOwnerCount(&_FREValidator.CallOpts)
}

// GetOwnerCount is a free data retrieval call binding the contract method 0xef18374a.
//
// Solidity: function getOwnerCount() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) GetOwnerCount() (*big.Int, error) {
	return _FREValidator.Contract.GetOwnerCount(&_FREValidator.CallOpts)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_FREValidator *FREValidatorCaller) GetVoterCap(opts *bind.CallOpts, _candidate common.Address, _voter common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getVoterCap", _candidate, _voter)
	return *ret0, err
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_FREValidator *FREValidatorSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _FREValidator.Contract.GetVoterCap(&_FREValidator.CallOpts, _candidate, _voter)
}

// GetVoterCap is a free data retrieval call binding the contract method 0x302b6872.
//
// Solidity: function getVoterCap(_candidate address, _voter address) constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) GetVoterCap(_candidate common.Address, _voter common.Address) (*big.Int, error) {
	return _FREValidator.Contract.GetVoterCap(&_FREValidator.CallOpts, _candidate, _voter)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_FREValidator *FREValidatorCaller) GetVoters(opts *bind.CallOpts, _candidate common.Address) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getVoters", _candidate)
	return *ret0, err
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_FREValidator *FREValidatorSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _FREValidator.Contract.GetVoters(&_FREValidator.CallOpts, _candidate)
}

// GetVoters is a free data retrieval call binding the contract method 0x2d15cc04.
//
// Solidity: function getVoters(_candidate address) constant returns(address[])
func (_FREValidator *FREValidatorCallerSession) GetVoters(_candidate common.Address) ([]common.Address, error) {
	return _FREValidator.Contract.GetVoters(&_FREValidator.CallOpts, _candidate)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_FREValidator *FREValidatorCaller) GetWithdrawBlockNumbers(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getWithdrawBlockNumbers")
	return *ret0, err
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_FREValidator *FREValidatorSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _FREValidator.Contract.GetWithdrawBlockNumbers(&_FREValidator.CallOpts)
}

// GetWithdrawBlockNumbers is a free data retrieval call binding the contract method 0x2f9c4bba.
//
// Solidity: function getWithdrawBlockNumbers() constant returns(uint256[])
func (_FREValidator *FREValidatorCallerSession) GetWithdrawBlockNumbers() ([]*big.Int, error) {
	return _FREValidator.Contract.GetWithdrawBlockNumbers(&_FREValidator.CallOpts)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_FREValidator *FREValidatorCaller) GetWithdrawCap(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "getWithdrawCap", _blockNumber)
	return *ret0, err
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_FREValidator *FREValidatorSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _FREValidator.Contract.GetWithdrawCap(&_FREValidator.CallOpts, _blockNumber)
}

// GetWithdrawCap is a free data retrieval call binding the contract method 0x15febd68.
//
// Solidity: function getWithdrawCap(_blockNumber uint256) constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) GetWithdrawCap(_blockNumber *big.Int) (*big.Int, error) {
	return _FREValidator.Contract.GetWithdrawCap(&_FREValidator.CallOpts, _blockNumber)
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_FREValidator *FREValidatorCaller) HasVotedInvalid(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "hasVotedInvalid", arg0, arg1)
	return *ret0, err
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_FREValidator *FREValidatorSession) HasVotedInvalid(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _FREValidator.Contract.HasVotedInvalid(&_FREValidator.CallOpts, arg0, arg1)
}

// HasVotedInvalid is a free data retrieval call binding the contract method 0x0e3e4fb8.
//
// Solidity: function hasVotedInvalid( address,  address) constant returns(bool)
func (_FREValidator *FREValidatorCallerSession) HasVotedInvalid(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _FREValidator.Contract.HasVotedInvalid(&_FREValidator.CallOpts, arg0, arg1)
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_FREValidator *FREValidatorCaller) InvalidKYCCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "invalidKYCCount", arg0)
	return *ret0, err
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_FREValidator *FREValidatorSession) InvalidKYCCount(arg0 common.Address) (*big.Int, error) {
	return _FREValidator.Contract.InvalidKYCCount(&_FREValidator.CallOpts, arg0)
}

// InvalidKYCCount is a free data retrieval call binding the contract method 0x72e44a38.
//
// Solidity: function invalidKYCCount( address) constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) InvalidKYCCount(arg0 common.Address) (*big.Int, error) {
	return _FREValidator.Contract.InvalidKYCCount(&_FREValidator.CallOpts, arg0)
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_FREValidator *FREValidatorCaller) InvalidPercent(opts *bind.CallOpts, _invalidCandidate common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "invalidPercent", _invalidCandidate)
	return *ret0, err
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_FREValidator *FREValidatorSession) InvalidPercent(_invalidCandidate common.Address) (*big.Int, error) {
	return _FREValidator.Contract.InvalidPercent(&_FREValidator.CallOpts, _invalidCandidate)
}

// InvalidPercent is a free data retrieval call binding the contract method 0x5b860d27.
//
// Solidity: function invalidPercent(_invalidCandidate address) constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) InvalidPercent(_invalidCandidate common.Address) (*big.Int, error) {
	return _FREValidator.Contract.InvalidPercent(&_FREValidator.CallOpts, _invalidCandidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_FREValidator *FREValidatorCaller) IsCandidate(opts *bind.CallOpts, _candidate common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "isCandidate", _candidate)
	return *ret0, err
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_FREValidator *FREValidatorSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _FREValidator.Contract.IsCandidate(&_FREValidator.CallOpts, _candidate)
}

// IsCandidate is a free data retrieval call binding the contract method 0xd51b9e93.
//
// Solidity: function isCandidate(_candidate address) constant returns(bool)
func (_FREValidator *FREValidatorCallerSession) IsCandidate(_candidate common.Address) (bool, error) {
	return _FREValidator.Contract.IsCandidate(&_FREValidator.CallOpts, _candidate)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "maxValidatorNumber")
	return *ret0, err
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_FREValidator *FREValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _FREValidator.Contract.MaxValidatorNumber(&_FREValidator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _FREValidator.Contract.MaxValidatorNumber(&_FREValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) MinCandidateCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "minCandidateCap")
	return *ret0, err
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_FREValidator *FREValidatorSession) MinCandidateCap() (*big.Int, error) {
	return _FREValidator.Contract.MinCandidateCap(&_FREValidator.CallOpts)
}

// MinCandidateCap is a free data retrieval call binding the contract method 0xd55b7dff.
//
// Solidity: function minCandidateCap() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) MinCandidateCap() (*big.Int, error) {
	return _FREValidator.Contract.MinCandidateCap(&_FREValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) MinVoterCap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "minVoterCap")
	return *ret0, err
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_FREValidator *FREValidatorSession) MinVoterCap() (*big.Int, error) {
	return _FREValidator.Contract.MinVoterCap(&_FREValidator.CallOpts)
}

// MinVoterCap is a free data retrieval call binding the contract method 0xf8ac9dd5.
//
// Solidity: function minVoterCap() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) MinVoterCap() (*big.Int, error) {
	return _FREValidator.Contract.MinVoterCap(&_FREValidator.CallOpts)
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) OwnerCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "ownerCount")
	return *ret0, err
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_FREValidator *FREValidatorSession) OwnerCount() (*big.Int, error) {
	return _FREValidator.Contract.OwnerCount(&_FREValidator.CallOpts)
}

// OwnerCount is a free data retrieval call binding the contract method 0x0db02622.
//
// Solidity: function ownerCount() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) OwnerCount() (*big.Int, error) {
	return _FREValidator.Contract.OwnerCount(&_FREValidator.CallOpts)
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_FREValidator *FREValidatorCaller) OwnerToCandidate(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "ownerToCandidate", arg0, arg1)
	return *ret0, err
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_FREValidator *FREValidatorSession) OwnerToCandidate(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FREValidator.Contract.OwnerToCandidate(&_FREValidator.CallOpts, arg0, arg1)
}

// OwnerToCandidate is a free data retrieval call binding the contract method 0x2a3640b1.
//
// Solidity: function ownerToCandidate( address,  uint256) constant returns(address)
func (_FREValidator *FREValidatorCallerSession) OwnerToCandidate(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FREValidator.Contract.OwnerToCandidate(&_FREValidator.CallOpts, arg0, arg1)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_FREValidator *FREValidatorCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_FREValidator *FREValidatorSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _FREValidator.Contract.Owners(&_FREValidator.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_FREValidator *FREValidatorCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _FREValidator.Contract.Owners(&_FREValidator.CallOpts, arg0)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_FREValidator *FREValidatorCaller) VoterWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _FREValidator.contract.Call(opts, out, "voterWithdrawDelay")
	return *ret0, err
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_FREValidator *FREValidatorSession) VoterWithdrawDelay() (*big.Int, error) {
	return _FREValidator.Contract.VoterWithdrawDelay(&_FREValidator.CallOpts)
}

// VoterWithdrawDelay is a free data retrieval call binding the contract method 0xa9ff959e.
//
// Solidity: function voterWithdrawDelay() constant returns(uint256)
func (_FREValidator *FREValidatorCallerSession) VoterWithdrawDelay() (*big.Int, error) {
	return _FREValidator.Contract.VoterWithdrawDelay(&_FREValidator.CallOpts)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_FREValidator *FREValidatorTransactor) Propose(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.contract.Transact(opts, "propose", _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_FREValidator *FREValidatorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.Propose(&_FREValidator.TransactOpts, _candidate)
}

// Propose is a paid mutator transaction binding the contract method 0x01267951.
//
// Solidity: function propose(_candidate address) returns()
func (_FREValidator *FREValidatorTransactorSession) Propose(_candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.Propose(&_FREValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_FREValidator *FREValidatorTransactor) Resign(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.contract.Transact(opts, "resign", _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_FREValidator *FREValidatorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.Resign(&_FREValidator.TransactOpts, _candidate)
}

// Resign is a paid mutator transaction binding the contract method 0xae6e43f5.
//
// Solidity: function resign(_candidate address) returns()
func (_FREValidator *FREValidatorTransactorSession) Resign(_candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.Resign(&_FREValidator.TransactOpts, _candidate)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_FREValidator *FREValidatorTransactor) Unvote(opts *bind.TransactOpts, _candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _FREValidator.contract.Transact(opts, "unvote", _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_FREValidator *FREValidatorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _FREValidator.Contract.Unvote(&_FREValidator.TransactOpts, _candidate, _cap)
}

// Unvote is a paid mutator transaction binding the contract method 0x02aa9be2.
//
// Solidity: function unvote(_candidate address, _cap uint256) returns()
func (_FREValidator *FREValidatorTransactorSession) Unvote(_candidate common.Address, _cap *big.Int) (*types.Transaction, error) {
	return _FREValidator.Contract.Unvote(&_FREValidator.TransactOpts, _candidate, _cap)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_FREValidator *FREValidatorTransactor) UploadKYC(opts *bind.TransactOpts, kychash string) (*types.Transaction, error) {
	return _FREValidator.contract.Transact(opts, "uploadKYC", kychash)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_FREValidator *FREValidatorSession) UploadKYC(kychash string) (*types.Transaction, error) {
	return _FREValidator.Contract.UploadKYC(&_FREValidator.TransactOpts, kychash)
}

// UploadKYC is a paid mutator transaction binding the contract method 0xf5c95125.
//
// Solidity: function uploadKYC(kychash string) returns()
func (_FREValidator *FREValidatorTransactorSession) UploadKYC(kychash string) (*types.Transaction, error) {
	return _FREValidator.Contract.UploadKYC(&_FREValidator.TransactOpts, kychash)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_FREValidator *FREValidatorTransactor) Vote(opts *bind.TransactOpts, _candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.contract.Transact(opts, "vote", _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_FREValidator *FREValidatorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.Vote(&_FREValidator.TransactOpts, _candidate)
}

// Vote is a paid mutator transaction binding the contract method 0x6dd7d8ea.
//
// Solidity: function vote(_candidate address) returns()
func (_FREValidator *FREValidatorTransactorSession) Vote(_candidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.Vote(&_FREValidator.TransactOpts, _candidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_FREValidator *FREValidatorTransactor) VoteInvalidKYC(opts *bind.TransactOpts, _invalidCandidate common.Address) (*types.Transaction, error) {
	return _FREValidator.contract.Transact(opts, "voteInvalidKYC", _invalidCandidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_FREValidator *FREValidatorSession) VoteInvalidKYC(_invalidCandidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.VoteInvalidKYC(&_FREValidator.TransactOpts, _invalidCandidate)
}

// VoteInvalidKYC is a paid mutator transaction binding the contract method 0xf2ee3c7d.
//
// Solidity: function voteInvalidKYC(_invalidCandidate address) returns()
func (_FREValidator *FREValidatorTransactorSession) VoteInvalidKYC(_invalidCandidate common.Address) (*types.Transaction, error) {
	return _FREValidator.Contract.VoteInvalidKYC(&_FREValidator.TransactOpts, _invalidCandidate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_FREValidator *FREValidatorTransactor) Withdraw(opts *bind.TransactOpts, _blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _FREValidator.contract.Transact(opts, "withdraw", _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_FREValidator *FREValidatorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _FREValidator.Contract.Withdraw(&_FREValidator.TransactOpts, _blockNumber, _index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(_blockNumber uint256, _index uint256) returns()
func (_FREValidator *FREValidatorTransactorSession) Withdraw(_blockNumber *big.Int, _index *big.Int) (*types.Transaction, error) {
	return _FREValidator.Contract.Withdraw(&_FREValidator.TransactOpts, _blockNumber, _index)
}

// FREValidatorInvalidatedNodeIterator is returned from FilterInvalidatedNode and is used to iterate over the raw logs and unpacked data for InvalidatedNode events raised by the FREValidator contract.
type FREValidatorInvalidatedNodeIterator struct {
	Event *FREValidatorInvalidatedNode // Event containing the contract specifics and raw log

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
func (it *FREValidatorInvalidatedNodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FREValidatorInvalidatedNode)
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
		it.Event = new(FREValidatorInvalidatedNode)
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
func (it *FREValidatorInvalidatedNodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FREValidatorInvalidatedNodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FREValidatorInvalidatedNode represents a InvalidatedNode event raised by the FREValidator contract.
type FREValidatorInvalidatedNode struct {
	MasternodeOwner common.Address
	Masternodes     []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterInvalidatedNode is a free log retrieval operation binding the contract event 0xe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e.
//
// Solidity: event InvalidatedNode(_masternodeOwner address, _masternodes address[])
func (_FREValidator *FREValidatorFilterer) FilterInvalidatedNode(opts *bind.FilterOpts) (*FREValidatorInvalidatedNodeIterator, error) {

	logs, sub, err := _FREValidator.contract.FilterLogs(opts, "InvalidatedNode")
	if err != nil {
		return nil, err
	}
	return &FREValidatorInvalidatedNodeIterator{contract: _FREValidator.contract, event: "InvalidatedNode", logs: logs, sub: sub}, nil
}

// WatchInvalidatedNode is a free log subscription operation binding the contract event 0xe18d61a5bf4aa2ab40afc88aa9039d27ae17ff4ec1c65f5f414df6f02ce8b35e.
//
// Solidity: event InvalidatedNode(_masternodeOwner address, _masternodes address[])
func (_FREValidator *FREValidatorFilterer) WatchInvalidatedNode(opts *bind.WatchOpts, sink chan<- *FREValidatorInvalidatedNode) (event.Subscription, error) {

	logs, sub, err := _FREValidator.contract.WatchLogs(opts, "InvalidatedNode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FREValidatorInvalidatedNode)
				if err := _FREValidator.contract.UnpackLog(event, "InvalidatedNode", log); err != nil {
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

// FREValidatorProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the FREValidator contract.
type FREValidatorProposeIterator struct {
	Event *FREValidatorPropose // Event containing the contract specifics and raw log

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
func (it *FREValidatorProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FREValidatorPropose)
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
		it.Event = new(FREValidatorPropose)
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
func (it *FREValidatorProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FREValidatorProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FREValidatorPropose represents a Propose event raised by the FREValidator contract.
type FREValidatorPropose struct {
	Owner     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_FREValidator *FREValidatorFilterer) FilterPropose(opts *bind.FilterOpts) (*FREValidatorProposeIterator, error) {

	logs, sub, err := _FREValidator.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &FREValidatorProposeIterator{contract: _FREValidator.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x7635f1d87b47fba9f2b09e56eb4be75cca030e0cb179c1602ac9261d39a8f5c1.
//
// Solidity: event Propose(_owner address, _candidate address, _cap uint256)
func (_FREValidator *FREValidatorFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *FREValidatorPropose) (event.Subscription, error) {

	logs, sub, err := _FREValidator.contract.WatchLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FREValidatorPropose)
				if err := _FREValidator.contract.UnpackLog(event, "Propose", log); err != nil {
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

// FREValidatorResignIterator is returned from FilterResign and is used to iterate over the raw logs and unpacked data for Resign events raised by the FREValidator contract.
type FREValidatorResignIterator struct {
	Event *FREValidatorResign // Event containing the contract specifics and raw log

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
func (it *FREValidatorResignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FREValidatorResign)
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
		it.Event = new(FREValidatorResign)
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
func (it *FREValidatorResignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FREValidatorResignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FREValidatorResign represents a Resign event raised by the FREValidator contract.
type FREValidatorResign struct {
	Owner     common.Address
	Candidate common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResign is a free log retrieval operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_FREValidator *FREValidatorFilterer) FilterResign(opts *bind.FilterOpts) (*FREValidatorResignIterator, error) {

	logs, sub, err := _FREValidator.contract.FilterLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return &FREValidatorResignIterator{contract: _FREValidator.contract, event: "Resign", logs: logs, sub: sub}, nil
}

// WatchResign is a free log subscription operation binding the contract event 0x4edf3e325d0063213a39f9085522994a1c44bea5f39e7d63ef61260a1e58c6d3.
//
// Solidity: event Resign(_owner address, _candidate address)
func (_FREValidator *FREValidatorFilterer) WatchResign(opts *bind.WatchOpts, sink chan<- *FREValidatorResign) (event.Subscription, error) {

	logs, sub, err := _FREValidator.contract.WatchLogs(opts, "Resign")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FREValidatorResign)
				if err := _FREValidator.contract.UnpackLog(event, "Resign", log); err != nil {
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

// FREValidatorUnvoteIterator is returned from FilterUnvote and is used to iterate over the raw logs and unpacked data for Unvote events raised by the FREValidator contract.
type FREValidatorUnvoteIterator struct {
	Event *FREValidatorUnvote // Event containing the contract specifics and raw log

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
func (it *FREValidatorUnvoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FREValidatorUnvote)
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
		it.Event = new(FREValidatorUnvote)
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
func (it *FREValidatorUnvoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FREValidatorUnvoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FREValidatorUnvote represents a Unvote event raised by the FREValidator contract.
type FREValidatorUnvote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnvote is a free log retrieval operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_FREValidator *FREValidatorFilterer) FilterUnvote(opts *bind.FilterOpts) (*FREValidatorUnvoteIterator, error) {

	logs, sub, err := _FREValidator.contract.FilterLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return &FREValidatorUnvoteIterator{contract: _FREValidator.contract, event: "Unvote", logs: logs, sub: sub}, nil
}

// WatchUnvote is a free log subscription operation binding the contract event 0xaa0e554f781c3c3b2be110a0557f260f11af9a8aa2c64bc1e7a31dbb21e32fa2.
//
// Solidity: event Unvote(_voter address, _candidate address, _cap uint256)
func (_FREValidator *FREValidatorFilterer) WatchUnvote(opts *bind.WatchOpts, sink chan<- *FREValidatorUnvote) (event.Subscription, error) {

	logs, sub, err := _FREValidator.contract.WatchLogs(opts, "Unvote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FREValidatorUnvote)
				if err := _FREValidator.contract.UnpackLog(event, "Unvote", log); err != nil {
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

// FREValidatorUploadedKYCIterator is returned from FilterUploadedKYC and is used to iterate over the raw logs and unpacked data for UploadedKYC events raised by the FREValidator contract.
type FREValidatorUploadedKYCIterator struct {
	Event *FREValidatorUploadedKYC // Event containing the contract specifics and raw log

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
func (it *FREValidatorUploadedKYCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FREValidatorUploadedKYC)
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
		it.Event = new(FREValidatorUploadedKYC)
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
func (it *FREValidatorUploadedKYCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FREValidatorUploadedKYCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FREValidatorUploadedKYC represents a UploadedKYC event raised by the FREValidator contract.
type FREValidatorUploadedKYC struct {
	Owner   common.Address
	KycHash string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUploadedKYC is a free log retrieval operation binding the contract event 0x949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a.
//
// Solidity: event UploadedKYC(_owner address, kycHash string)
func (_FREValidator *FREValidatorFilterer) FilterUploadedKYC(opts *bind.FilterOpts) (*FREValidatorUploadedKYCIterator, error) {

	logs, sub, err := _FREValidator.contract.FilterLogs(opts, "UploadedKYC")
	if err != nil {
		return nil, err
	}
	return &FREValidatorUploadedKYCIterator{contract: _FREValidator.contract, event: "UploadedKYC", logs: logs, sub: sub}, nil
}

// WatchUploadedKYC is a free log subscription operation binding the contract event 0x949360d814b28a3b393a68909efe1fee120ee09cac30f360a0f80ab5415a611a.
//
// Solidity: event UploadedKYC(_owner address, kycHash string)
func (_FREValidator *FREValidatorFilterer) WatchUploadedKYC(opts *bind.WatchOpts, sink chan<- *FREValidatorUploadedKYC) (event.Subscription, error) {

	logs, sub, err := _FREValidator.contract.WatchLogs(opts, "UploadedKYC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FREValidatorUploadedKYC)
				if err := _FREValidator.contract.UnpackLog(event, "UploadedKYC", log); err != nil {
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

// FREValidatorVoteIterator is returned from FilterVote and is used to iterate over the raw logs and unpacked data for Vote events raised by the FREValidator contract.
type FREValidatorVoteIterator struct {
	Event *FREValidatorVote // Event containing the contract specifics and raw log

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
func (it *FREValidatorVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FREValidatorVote)
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
		it.Event = new(FREValidatorVote)
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
func (it *FREValidatorVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FREValidatorVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FREValidatorVote represents a Vote event raised by the FREValidator contract.
type FREValidatorVote struct {
	Voter     common.Address
	Candidate common.Address
	Cap       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVote is a free log retrieval operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_FREValidator *FREValidatorFilterer) FilterVote(opts *bind.FilterOpts) (*FREValidatorVoteIterator, error) {

	logs, sub, err := _FREValidator.contract.FilterLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return &FREValidatorVoteIterator{contract: _FREValidator.contract, event: "Vote", logs: logs, sub: sub}, nil
}

// WatchVote is a free log subscription operation binding the contract event 0x66a9138482c99e9baf08860110ef332cc0c23b4a199a53593d8db0fc8f96fbfc.
//
// Solidity: event Vote(_voter address, _candidate address, _cap uint256)
func (_FREValidator *FREValidatorFilterer) WatchVote(opts *bind.WatchOpts, sink chan<- *FREValidatorVote) (event.Subscription, error) {

	logs, sub, err := _FREValidator.contract.WatchLogs(opts, "Vote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FREValidatorVote)
				if err := _FREValidator.contract.UnpackLog(event, "Vote", log); err != nil {
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

// FREValidatorWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the FREValidator contract.
type FREValidatorWithdrawIterator struct {
	Event *FREValidatorWithdraw // Event containing the contract specifics and raw log

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
func (it *FREValidatorWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FREValidatorWithdraw)
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
		it.Event = new(FREValidatorWithdraw)
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
func (it *FREValidatorWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FREValidatorWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FREValidatorWithdraw represents a Withdraw event raised by the FREValidator contract.
type FREValidatorWithdraw struct {
	Owner       common.Address
	BlockNumber *big.Int
	Cap         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_FREValidator *FREValidatorFilterer) FilterWithdraw(opts *bind.FilterOpts) (*FREValidatorWithdrawIterator, error) {

	logs, sub, err := _FREValidator.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &FREValidatorWithdrawIterator{contract: _FREValidator.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(_owner address, _blockNumber uint256, _cap uint256)
func (_FREValidator *FREValidatorFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FREValidatorWithdraw) (event.Subscription, error) {

	logs, sub, err := _FREValidator.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FREValidatorWithdraw)
				if err := _FREValidator.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
