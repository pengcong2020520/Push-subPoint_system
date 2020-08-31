// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bcos

import (
	"math/big"
	"strings"

	"github.com/yekai1003/gobcos/accounts/abi"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/common"
	"github.com/yekai1003/gobcos/core/types"
	"github.com/yekai1003/gobcos/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = common.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ControlABI is the input ABI used to generate the binding from.
const ControlABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"jsonData\",\"type\":\"string\"},{\"name\":\"month\",\"type\":\"uint256\"}],\"name\":\"pushLog\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"pass\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"oldPass\",\"type\":\"string\"},{\"name\":\"newPass\",\"type\":\"string\"}],\"name\":\"setPasswd\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"string\"},{\"name\":\"to\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_onwer\",\"type\":\"address\"}],\"name\":\"updateOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ercaddr\",\"type\":\"address\"}],\"name\":\"upgradeErc200\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logaddr\",\"type\":\"address\"}],\"name\":\"upgradeErclog\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"useraddr\",\"type\":\"address\"}],\"name\":\"upgradeUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_onwer\",\"type\":\"address\"},{\"name\":\"sym\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"string\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"itype\",\"type\":\"uint8\"}],\"name\":\"getAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"pass\",\"type\":\"string\"}],\"name\":\"login\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"begin\",\"type\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"queryLog\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userid\",\"type\":\"string\"},{\"name\":\"month\",\"type\":\"uint256\"}],\"name\":\"queryLogByMongh\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Control is an auto generated Go binding around an Ethereum contract.
type Control struct {
	ControlCaller     // Read-only binding to the contract
	ControlTransactor // Write-only binding to the contract
	ControlFilterer   // Log filterer for contract events
}

// ControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControlSession struct {
	Contract     *Control          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControlCallerSession struct {
	Contract *ControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControlTransactorSession struct {
	Contract     *ControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControlRaw struct {
	Contract *Control // Generic contract binding to access the raw methods on
}

// ControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControlCallerRaw struct {
	Contract *ControlCaller // Generic read-only contract binding to access the raw methods on
}

// ControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControlTransactorRaw struct {
	Contract *ControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewControl creates a new instance of Control, bound to a specific deployed contract.
func NewControl(address common.Address, backend bind.ContractBackend) (*Control, error) {
	contract, err := bindControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Control{ControlCaller: ControlCaller{contract: contract}, ControlTransactor: ControlTransactor{contract: contract}, ControlFilterer: ControlFilterer{contract: contract}}, nil
}

// NewControlCaller creates a new read-only instance of Control, bound to a specific deployed contract.
func NewControlCaller(address common.Address, caller bind.ContractCaller) (*ControlCaller, error) {
	contract, err := bindControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControlCaller{contract: contract}, nil
}

// NewControlTransactor creates a new write-only instance of Control, bound to a specific deployed contract.
func NewControlTransactor(address common.Address, transactor bind.ContractTransactor) (*ControlTransactor, error) {
	contract, err := bindControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControlTransactor{contract: contract}, nil
}

// NewControlFilterer creates a new log filterer instance of Control, bound to a specific deployed contract.
func NewControlFilterer(address common.Address, filterer bind.ContractFilterer) (*ControlFilterer, error) {
	contract, err := bindControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControlFilterer{contract: contract}, nil
}

// bindControl binds a generic wrapper to an already deployed contract.
func bindControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control *ControlRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Control.Contract.ControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control *ControlRaw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _Control.Contract.ControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control *ControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _Control.Contract.ControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Control *ControlCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Control.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Control *ControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _Control.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Control *ControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _Control.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string who) constant returns(uint256)
func (_Control *ControlCaller) BalanceOf(opts *bind.CallOpts, who string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Control.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string who) constant returns(uint256)
func (_Control *ControlSession) BalanceOf(who string) (*big.Int, error) {
	return _Control.Contract.BalanceOf(&_Control.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x35ee5f87.
//
// Solidity: function balanceOf(string who) constant returns(uint256)
func (_Control *ControlCallerSession) BalanceOf(who string) (*big.Int, error) {
	return _Control.Contract.BalanceOf(&_Control.CallOpts, who)
}

// GetAddr is a free data retrieval call binding the contract method 0x45cd33a5.
//
// Solidity: function getAddr(uint8 itype) constant returns(address)
func (_Control *ControlCaller) GetAddr(opts *bind.CallOpts, itype uint8) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Control.contract.Call(opts, out, "getAddr", itype)
	return *ret0, err
}

// GetAddr is a free data retrieval call binding the contract method 0x45cd33a5.
//
// Solidity: function getAddr(uint8 itype) constant returns(address)
func (_Control *ControlSession) GetAddr(itype uint8) (common.Address, error) {
	return _Control.Contract.GetAddr(&_Control.CallOpts, itype)
}

// GetAddr is a free data retrieval call binding the contract method 0x45cd33a5.
//
// Solidity: function getAddr(uint8 itype) constant returns(address)
func (_Control *ControlCallerSession) GetAddr(itype uint8) (common.Address, error) {
	return _Control.Contract.GetAddr(&_Control.CallOpts, itype)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string userid, string pass) constant returns(bool)
func (_Control *ControlCaller) Login(opts *bind.CallOpts, userid string, pass string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Control.contract.Call(opts, out, "login", userid, pass)
	return *ret0, err
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string userid, string pass) constant returns(bool)
func (_Control *ControlSession) Login(userid string, pass string) (bool, error) {
	return _Control.Contract.Login(&_Control.CallOpts, userid, pass)
}

// Login is a free data retrieval call binding the contract method 0x58467dbc.
//
// Solidity: function login(string userid, string pass) constant returns(bool)
func (_Control *ControlCallerSession) Login(userid string, pass string) (bool, error) {
	return _Control.Contract.Login(&_Control.CallOpts, userid, pass)
}

// QueryLog is a free data retrieval call binding the contract method 0x07e65cc4.
//
// Solidity: function queryLog(string userid, uint256 begin, uint256 end) constant returns(string[])
func (_Control *ControlCaller) QueryLog(opts *bind.CallOpts, userid string, begin *big.Int, end *big.Int) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Control.contract.Call(opts, out, "queryLog", userid, begin, end)
	return *ret0, err
}

// QueryLog is a free data retrieval call binding the contract method 0x07e65cc4.
//
// Solidity: function queryLog(string userid, uint256 begin, uint256 end) constant returns(string[])
func (_Control *ControlSession) QueryLog(userid string, begin *big.Int, end *big.Int) ([]string, error) {
	return _Control.Contract.QueryLog(&_Control.CallOpts, userid, begin, end)
}

// QueryLog is a free data retrieval call binding the contract method 0x07e65cc4.
//
// Solidity: function queryLog(string userid, uint256 begin, uint256 end) constant returns(string[])
func (_Control *ControlCallerSession) QueryLog(userid string, begin *big.Int, end *big.Int) ([]string, error) {
	return _Control.Contract.QueryLog(&_Control.CallOpts, userid, begin, end)
}

// QueryLogByMongh is a free data retrieval call binding the contract method 0xed8a1c70.
//
// Solidity: function queryLogByMongh(string userid, uint256 month) constant returns(string[])
func (_Control *ControlCaller) QueryLogByMongh(opts *bind.CallOpts, userid string, month *big.Int) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Control.contract.Call(opts, out, "queryLogByMongh", userid, month)
	return *ret0, err
}

// QueryLogByMongh is a free data retrieval call binding the contract method 0xed8a1c70.
//
// Solidity: function queryLogByMongh(string userid, uint256 month) constant returns(string[])
func (_Control *ControlSession) QueryLogByMongh(userid string, month *big.Int) ([]string, error) {
	return _Control.Contract.QueryLogByMongh(&_Control.CallOpts, userid, month)
}

// QueryLogByMongh is a free data retrieval call binding the contract method 0xed8a1c70.
//
// Solidity: function queryLogByMongh(string userid, uint256 month) constant returns(string[])
func (_Control *ControlCallerSession) QueryLogByMongh(userid string, month *big.Int) ([]string, error) {
	return _Control.Contract.QueryLogByMongh(&_Control.CallOpts, userid, month)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Control *ControlCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Control.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Control *ControlSession) TotalSupply() (*big.Int, error) {
	return _Control.Contract.TotalSupply(&_Control.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Control *ControlCallerSession) TotalSupply() (*big.Int, error) {
	return _Control.Contract.TotalSupply(&_Control.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns(bool)
func (_Control *ControlTransactor) Mint(opts *bind.TransactOpts, to string, value *big.Int) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns(bool)
func (_Control *ControlSession) Mint(to string, value *big.Int) (*types.RawTransaction, error) {
	return _Control.Contract.Mint(&_Control.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x056b01ce.
//
// Solidity: function mint(string to, uint256 value) returns(bool)
func (_Control *ControlTransactorSession) Mint(to string, value *big.Int) (*types.RawTransaction, error) {
	return _Control.Contract.Mint(&_Control.TransactOpts, to, value)
}

// PushLog is a paid mutator transaction binding the contract method 0x8e34f2d1.
//
// Solidity: function pushLog(string userid, string jsonData, uint256 month) returns()
func (_Control *ControlTransactor) PushLog(opts *bind.TransactOpts, userid string, jsonData string, month *big.Int) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "pushLog", userid, jsonData, month)
}

// PushLog is a paid mutator transaction binding the contract method 0x8e34f2d1.
//
// Solidity: function pushLog(string userid, string jsonData, uint256 month) returns()
func (_Control *ControlSession) PushLog(userid string, jsonData string, month *big.Int) (*types.RawTransaction, error) {
	return _Control.Contract.PushLog(&_Control.TransactOpts, userid, jsonData, month)
}

// PushLog is a paid mutator transaction binding the contract method 0x8e34f2d1.
//
// Solidity: function pushLog(string userid, string jsonData, uint256 month) returns()
func (_Control *ControlTransactorSession) PushLog(userid string, jsonData string, month *big.Int) (*types.RawTransaction, error) {
	return _Control.Contract.PushLog(&_Control.TransactOpts, userid, jsonData, month)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string userid, string pass) returns()
func (_Control *ControlTransactor) Register(opts *bind.TransactOpts, userid string, pass string) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "register", userid, pass)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string userid, string pass) returns()
func (_Control *ControlSession) Register(userid string, pass string) (*types.RawTransaction, error) {
	return _Control.Contract.Register(&_Control.TransactOpts, userid, pass)
}

// Register is a paid mutator transaction binding the contract method 0x3ffbd47f.
//
// Solidity: function register(string userid, string pass) returns()
func (_Control *ControlTransactorSession) Register(userid string, pass string) (*types.RawTransaction, error) {
	return _Control.Contract.Register(&_Control.TransactOpts, userid, pass)
}

// SetPasswd is a paid mutator transaction binding the contract method 0xfa9601c7.
//
// Solidity: function setPasswd(string userid, string oldPass, string newPass) returns()
func (_Control *ControlTransactor) SetPasswd(opts *bind.TransactOpts, userid string, oldPass string, newPass string) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "setPasswd", userid, oldPass, newPass)
}

// SetPasswd is a paid mutator transaction binding the contract method 0xfa9601c7.
//
// Solidity: function setPasswd(string userid, string oldPass, string newPass) returns()
func (_Control *ControlSession) SetPasswd(userid string, oldPass string, newPass string) (*types.RawTransaction, error) {
	return _Control.Contract.SetPasswd(&_Control.TransactOpts, userid, oldPass, newPass)
}

// SetPasswd is a paid mutator transaction binding the contract method 0xfa9601c7.
//
// Solidity: function setPasswd(string userid, string oldPass, string newPass) returns()
func (_Control *ControlTransactorSession) SetPasswd(userid string, oldPass string, newPass string) (*types.RawTransaction, error) {
	return _Control.Contract.SetPasswd(&_Control.TransactOpts, userid, oldPass, newPass)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string owner, string to, uint256 value) returns(bool)
func (_Control *ControlTransactor) Transfer(opts *bind.TransactOpts, owner string, to string, value *big.Int) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "transfer", owner, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string owner, string to, uint256 value) returns(bool)
func (_Control *ControlSession) Transfer(owner string, to string, value *big.Int) (*types.RawTransaction, error) {
	return _Control.Contract.Transfer(&_Control.TransactOpts, owner, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0x9b80b050.
//
// Solidity: function transfer(string owner, string to, uint256 value) returns(bool)
func (_Control *ControlTransactorSession) Transfer(owner string, to string, value *big.Int) (*types.RawTransaction, error) {
	return _Control.Contract.Transfer(&_Control.TransactOpts, owner, to, value)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address _onwer) returns()
func (_Control *ControlTransactor) UpdateOwner(opts *bind.TransactOpts, _onwer common.Address) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "updateOwner", _onwer)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address _onwer) returns()
func (_Control *ControlSession) UpdateOwner(_onwer common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpdateOwner(&_Control.TransactOpts, _onwer)
}

// UpdateOwner is a paid mutator transaction binding the contract method 0x880cdc31.
//
// Solidity: function updateOwner(address _onwer) returns()
func (_Control *ControlTransactorSession) UpdateOwner(_onwer common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpdateOwner(&_Control.TransactOpts, _onwer)
}

// UpgradeErc200 is a paid mutator transaction binding the contract method 0xb44d8e43.
//
// Solidity: function upgradeErc200(address ercaddr) returns()
func (_Control *ControlTransactor) UpgradeErc200(opts *bind.TransactOpts, ercaddr common.Address) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "upgradeErc200", ercaddr)
}

// UpgradeErc200 is a paid mutator transaction binding the contract method 0xb44d8e43.
//
// Solidity: function upgradeErc200(address ercaddr) returns()
func (_Control *ControlSession) UpgradeErc200(ercaddr common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpgradeErc200(&_Control.TransactOpts, ercaddr)
}

// UpgradeErc200 is a paid mutator transaction binding the contract method 0xb44d8e43.
//
// Solidity: function upgradeErc200(address ercaddr) returns()
func (_Control *ControlTransactorSession) UpgradeErc200(ercaddr common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpgradeErc200(&_Control.TransactOpts, ercaddr)
}

// UpgradeErclog is a paid mutator transaction binding the contract method 0x9a92be03.
//
// Solidity: function upgradeErclog(address logaddr) returns()
func (_Control *ControlTransactor) UpgradeErclog(opts *bind.TransactOpts, logaddr common.Address) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "upgradeErclog", logaddr)
}

// UpgradeErclog is a paid mutator transaction binding the contract method 0x9a92be03.
//
// Solidity: function upgradeErclog(address logaddr) returns()
func (_Control *ControlSession) UpgradeErclog(logaddr common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpgradeErclog(&_Control.TransactOpts, logaddr)
}

// UpgradeErclog is a paid mutator transaction binding the contract method 0x9a92be03.
//
// Solidity: function upgradeErclog(address logaddr) returns()
func (_Control *ControlTransactorSession) UpgradeErclog(logaddr common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpgradeErclog(&_Control.TransactOpts, logaddr)
}

// UpgradeUser is a paid mutator transaction binding the contract method 0x23814798.
//
// Solidity: function upgradeUser(address useraddr) returns()
func (_Control *ControlTransactor) UpgradeUser(opts *bind.TransactOpts, useraddr common.Address) (*types.RawTransaction, error) {
	return _Control.contract.Transact(opts, "upgradeUser", useraddr)
}

// UpgradeUser is a paid mutator transaction binding the contract method 0x23814798.
//
// Solidity: function upgradeUser(address useraddr) returns()
func (_Control *ControlSession) UpgradeUser(useraddr common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpgradeUser(&_Control.TransactOpts, useraddr)
}

// UpgradeUser is a paid mutator transaction binding the contract method 0x23814798.
//
// Solidity: function upgradeUser(address useraddr) returns()
func (_Control *ControlTransactorSession) UpgradeUser(useraddr common.Address) (*types.RawTransaction, error) {
	return _Control.Contract.UpgradeUser(&_Control.TransactOpts, useraddr)
}
