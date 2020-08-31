// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PingABI is the input ABI used to generate the binding from.
const PingABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getMsg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_msg\",\"type\":\"string\"}],\"name\":\"setMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PingFuncSigs maps the 4-byte function signature to its string representation.
var PingFuncSigs = map[string]string{
	"b5fdeb23": "getMsg()",
	"c4784fd4": "setMsg(string)",
}

// PingBin is the compiled bytecode used for deploying new contracts.
var PingBin = "0x60c060405260056080819052643837b7339960d91b60a0908152610026916001919061007c565b5034801561003357600080fd5b506040516103f33803806103f38339818101604052602081101561005657600080fd5b5051600080546001600160a01b0319166001600160a01b0390921691909117905561010f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100bd57805160ff19168380011785556100ea565b828001600101855582156100ea579182015b828111156100ea5782518255916020019190600101906100cf565b506100f69291506100fa565b5090565b5b808211156100f657600081556001016100fb565b6102d58061011e6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063b5fdeb231461003b578063c4784fd4146100b8575b600080fd5b610043610160565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561007d578181015183820152602001610065565b50505050905090810190601f1680156100aa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61015e600480360360208110156100ce57600080fd5b8101906020810181356401000000008111156100e957600080fd5b8201836020820111156100fb57600080fd5b8035906020019184600183028401116401000000008311171561011d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506101f5945050505050565b005b60018054604080516020601f600260001961010087891615020190951694909404938401819004810282018101909252828152606093909290918301828280156101eb5780601f106101c0576101008083540402835291602001916101eb565b820191906000526020600020905b8154815290600101906020018083116101ce57829003601f168201915b5050505050905090565b805161020890600190602084019061020c565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061024d57805160ff191683800117855561027a565b8280016001018555821561027a579182015b8281111561027a57825182559160200191906001019061025f565b5061028692915061028a565b5090565b5b80821115610286576000815560010161028b56fea264697066735822122032624bcde668c67da520ac17f278163327d7e2a721ba1b29f391dda9918a428564736f6c63430007000033"

// DeployPing deploys a new Ethereum contract, binding an instance of Ping to it.
func DeployPing(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *Ping, error) {
	parsed, err := abi.JSON(strings.NewReader(PingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PingBin), backend, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ping{PingCaller: PingCaller{contract: contract}, PingTransactor: PingTransactor{contract: contract}, PingFilterer: PingFilterer{contract: contract}}, nil
}

// Ping is an auto generated Go binding around an Ethereum contract.
type Ping struct {
	PingCaller     // Read-only binding to the contract
	PingTransactor // Write-only binding to the contract
	PingFilterer   // Log filterer for contract events
}

// PingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PingSession struct {
	Contract     *Ping             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PingCallerSession struct {
	Contract *PingCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PingTransactorSession struct {
	Contract     *PingTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PingRaw struct {
	Contract *Ping // Generic contract binding to access the raw methods on
}

// PingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PingCallerRaw struct {
	Contract *PingCaller // Generic read-only contract binding to access the raw methods on
}

// PingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PingTransactorRaw struct {
	Contract *PingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPing creates a new instance of Ping, bound to a specific deployed contract.
func NewPing(address common.Address, backend bind.ContractBackend) (*Ping, error) {
	contract, err := bindPing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ping{PingCaller: PingCaller{contract: contract}, PingTransactor: PingTransactor{contract: contract}, PingFilterer: PingFilterer{contract: contract}}, nil
}

// NewPingCaller creates a new read-only instance of Ping, bound to a specific deployed contract.
func NewPingCaller(address common.Address, caller bind.ContractCaller) (*PingCaller, error) {
	contract, err := bindPing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingCaller{contract: contract}, nil
}

// NewPingTransactor creates a new write-only instance of Ping, bound to a specific deployed contract.
func NewPingTransactor(address common.Address, transactor bind.ContractTransactor) (*PingTransactor, error) {
	contract, err := bindPing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingTransactor{contract: contract}, nil
}

// NewPingFilterer creates a new log filterer instance of Ping, bound to a specific deployed contract.
func NewPingFilterer(address common.Address, filterer bind.ContractFilterer) (*PingFilterer, error) {
	contract, err := bindPing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingFilterer{contract: contract}, nil
}

// bindPing binds a generic wrapper to an already deployed contract.
func bindPing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ping *PingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ping.Contract.PingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ping *PingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ping.Contract.PingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ping *PingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ping.Contract.PingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ping *PingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ping *PingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ping *PingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ping.Contract.contract.Transact(opts, method, params...)
}

// GetMsg is a free data retrieval call binding the contract method 0xb5fdeb23.
//
// Solidity: function getMsg() view returns(string)
func (_Ping *PingCaller) GetMsg(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ping.contract.Call(opts, out, "getMsg")
	return *ret0, err
}

// GetMsg is a free data retrieval call binding the contract method 0xb5fdeb23.
//
// Solidity: function getMsg() view returns(string)
func (_Ping *PingSession) GetMsg() (string, error) {
	return _Ping.Contract.GetMsg(&_Ping.CallOpts)
}

// GetMsg is a free data retrieval call binding the contract method 0xb5fdeb23.
//
// Solidity: function getMsg() view returns(string)
func (_Ping *PingCallerSession) GetMsg() (string, error) {
	return _Ping.Contract.GetMsg(&_Ping.CallOpts)
}

// SetMsg is a paid mutator transaction binding the contract method 0xc4784fd4.
//
// Solidity: function setMsg(string _msg) returns()
func (_Ping *PingTransactor) SetMsg(opts *bind.TransactOpts, _msg string) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "setMsg", _msg)
}

// SetMsg is a paid mutator transaction binding the contract method 0xc4784fd4.
//
// Solidity: function setMsg(string _msg) returns()
func (_Ping *PingSession) SetMsg(_msg string) (*types.Transaction, error) {
	return _Ping.Contract.SetMsg(&_Ping.TransactOpts, _msg)
}

// SetMsg is a paid mutator transaction binding the contract method 0xc4784fd4.
//
// Solidity: function setMsg(string _msg) returns()
func (_Ping *PingTransactorSession) SetMsg(_msg string) (*types.Transaction, error) {
	return _Ping.Contract.SetMsg(&_Ping.TransactOpts, _msg)
}
