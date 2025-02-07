// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package builtins

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/darrenvechain/thorgo/accounts"
	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = bind.Bind
	_ = common.Big1
	_ = abi.ConvertType
	_ = hexutil.MustDecode
	_ = context.Background
	_ = tx.NewClause
	_ = blocks.New
)

// PrototypeMetaData contains all meta data concerning the Prototype contract.
var PrototypeMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_newMaster\",\"type\":\"address\"}],\"name\":\"setMaster\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"storageFor\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"energy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"currentSponsor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_credit\",\"type\":\"uint256\"},{\"name\":\"_recoveryRate\",\"type\":\"uint256\"}],\"name\":\"setCreditPlan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_sponsor\",\"type\":\"address\"}],\"name\":\"selectSponsor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"sponsor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"creditPlan\",\"outputs\":[{\"name\":\"credit\",\"type\":\"uint256\"},{\"name\":\"recoveryRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"hasCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"master\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"unsponsor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_sponsor\",\"type\":\"address\"}],\"name\":\"isSponsor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Prototype is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Prototype struct {
	thor     *thorest.Client    // Thor client connection to use
	contract *accounts.Contract // Generic contract wrapper for the low level calls
}

// PrototypeTransactor is an auto generated Go binding around an Ethereum, allowing you to transact with the contract.
type PrototypeTransactor struct {
	*Prototype
	contract *accounts.ContractTransactor // Generic contract wrapper for the low level calls
	manager  accounts.TxManager           // TxManager to use
}

// NewPrototype creates a new instance of Prototype, bound to a specific deployed contract.
func NewPrototype(thor *thorest.Client) (*Prototype, error) {
	parsed, err := PrototypeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := accounts.New(thor, common.HexToAddress("0x000000000000000000000050726f746f74797065")).Contract(parsed)
	return &Prototype{thor: thor, contract: contract}, nil
}

// NewPrototypeTransactor creates a new instance of PrototypeTransactor, bound to a specific deployed contract.
func NewPrototypeTransactor(thor *thorest.Client, manager accounts.TxManager) (*PrototypeTransactor, error) {
	base, err := NewPrototype(thor)
	if err != nil {
		return nil, err
	}
	return &PrototypeTransactor{Prototype: base, contract: base.contract.Transactor(manager), manager: manager}, nil
}

// Address returns the address of the contract.
func (_Prototype *Prototype) Address() common.Address {
	return _Prototype.contract.Address
}

// Transactor constructs a new transactor for the contract, which allows to send transactions.
func (_Prototype *Prototype) Transactor(manager accounts.TxManager) *PrototypeTransactor {
	return &PrototypeTransactor{Prototype: _Prototype, contract: _Prototype.contract.Transactor(manager), manager: manager}
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Prototype *Prototype) Call(revision thorest.Revision, result *[]interface{}, method string, params ...interface{}) error {
	return _Prototype.contract.CallAt(revision, method, result, params...)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrototypeTransactor *PrototypeTransactor) Transact(opts *transactions.Options, method string, params ...interface{}) (*transactions.Visitor, error) {
	return _PrototypeTransactor.contract.Send(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x6d8c859a.
//
// Solidity: function balance(address _self, uint256 _blockNumber) view returns(uint256)
func (_Prototype *Prototype) Balance(_self common.Address, _blockNumber *big.Int, revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "balance", _self, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// CreditPlan is a free data retrieval call binding the contract method 0x80df45b4.
//
// Solidity: function creditPlan(address _self) view returns(uint256 credit, uint256 recoveryRate)
func (_Prototype *Prototype) CreditPlan(_self common.Address, revision ...thorest.Revision) (struct {
	Credit       *big.Int
	RecoveryRate *big.Int
}, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "creditPlan", _self)

	outstruct := new(struct {
		Credit       *big.Int
		RecoveryRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Credit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RecoveryRate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentSponsor is a free data retrieval call binding the contract method 0x23d8c7db.
//
// Solidity: function currentSponsor(address _self) view returns(address)
func (_Prototype *Prototype) CurrentSponsor(_self common.Address, revision ...thorest.Revision) (common.Address, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "currentSponsor", _self)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Energy is a free data retrieval call binding the contract method 0x1e95be45.
//
// Solidity: function energy(address _self, uint256 _blockNumber) view returns(uint256)
func (_Prototype *Prototype) Energy(_self common.Address, _blockNumber *big.Int, revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "energy", _self, _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// HasCode is a free data retrieval call binding the contract method 0x9538c4b3.
//
// Solidity: function hasCode(address _self) view returns(bool)
func (_Prototype *Prototype) HasCode(_self common.Address, revision ...thorest.Revision) (bool, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "hasCode", _self)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// IsSponsor is a free data retrieval call binding the contract method 0xd87333ac.
//
// Solidity: function isSponsor(address _self, address _sponsor) view returns(bool)
func (_Prototype *Prototype) IsSponsor(_self common.Address, _sponsor common.Address, revision ...thorest.Revision) (bool, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "isSponsor", _self, _sponsor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// IsUser is a free data retrieval call binding the contract method 0x02d43dc8.
//
// Solidity: function isUser(address _self, address _user) view returns(bool)
func (_Prototype *Prototype) IsUser(_self common.Address, _user common.Address, revision ...thorest.Revision) (bool, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "isUser", _self, _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// Master is a free data retrieval call binding the contract method 0x9ed153c0.
//
// Solidity: function master(address _self) view returns(address)
func (_Prototype *Prototype) Master(_self common.Address, revision ...thorest.Revision) (common.Address, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "master", _self)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// StorageFor is a free data retrieval call binding the contract method 0x04e7a457.
//
// Solidity: function storageFor(address _self, bytes32 _key) view returns(bytes32)
func (_Prototype *Prototype) StorageFor(_self common.Address, _key [32]byte, revision ...thorest.Revision) ([32]byte, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "storageFor", _self, _key)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err
}

// UserCredit is a free data retrieval call binding the contract method 0xc9c4fc41.
//
// Solidity: function userCredit(address _self, address _user) view returns(uint256)
func (_Prototype *Prototype) UserCredit(_self common.Address, _user common.Address, revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Prototype.Call(rev, &out, "userCredit", _self, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// AddUser is a paid mutator transaction binding the contract method 0x8ca3b448.
//
// Solidity: function addUser(address _self, address _user) returns()
func (_PrototypeTransactor *PrototypeTransactor) AddUser(_self common.Address, _user common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _PrototypeTransactor.Transact(opts, "addUser", _self, _user)
}

// AddUserAsClause is a transaction clause generator 0x8ca3b448.
//
// Solidity: function addUser(address _self, address _user) returns()
func (_Prototype *Prototype) AddUserAsClause(_self common.Address, _user common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Prototype.contract.AsClauseWithVET(val, "addUser", _self, _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x22928d6b.
//
// Solidity: function removeUser(address _self, address _user) returns()
func (_PrototypeTransactor *PrototypeTransactor) RemoveUser(_self common.Address, _user common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _PrototypeTransactor.Transact(opts, "removeUser", _self, _user)
}

// RemoveUserAsClause is a transaction clause generator 0x22928d6b.
//
// Solidity: function removeUser(address _self, address _user) returns()
func (_Prototype *Prototype) RemoveUserAsClause(_self common.Address, _user common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Prototype.contract.AsClauseWithVET(val, "removeUser", _self, _user)
}

// SelectSponsor is a paid mutator transaction binding the contract method 0x3871a9fb.
//
// Solidity: function selectSponsor(address _self, address _sponsor) returns()
func (_PrototypeTransactor *PrototypeTransactor) SelectSponsor(_self common.Address, _sponsor common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _PrototypeTransactor.Transact(opts, "selectSponsor", _self, _sponsor)
}

// SelectSponsorAsClause is a transaction clause generator 0x3871a9fb.
//
// Solidity: function selectSponsor(address _self, address _sponsor) returns()
func (_Prototype *Prototype) SelectSponsorAsClause(_self common.Address, _sponsor common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Prototype.contract.AsClauseWithVET(val, "selectSponsor", _self, _sponsor)
}

// SetCreditPlan is a paid mutator transaction binding the contract method 0x3659f8ed.
//
// Solidity: function setCreditPlan(address _self, uint256 _credit, uint256 _recoveryRate) returns()
func (_PrototypeTransactor *PrototypeTransactor) SetCreditPlan(_self common.Address, _credit *big.Int, _recoveryRate *big.Int, opts *transactions.Options) (*transactions.Visitor, error) {
	return _PrototypeTransactor.Transact(opts, "setCreditPlan", _self, _credit, _recoveryRate)
}

// SetCreditPlanAsClause is a transaction clause generator 0x3659f8ed.
//
// Solidity: function setCreditPlan(address _self, uint256 _credit, uint256 _recoveryRate) returns()
func (_Prototype *Prototype) SetCreditPlanAsClause(_self common.Address, _credit *big.Int, _recoveryRate *big.Int, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Prototype.contract.AsClauseWithVET(val, "setCreditPlan", _self, _credit, _recoveryRate)
}

// SetMaster is a paid mutator transaction binding the contract method 0x01378b58.
//
// Solidity: function setMaster(address _self, address _newMaster) returns()
func (_PrototypeTransactor *PrototypeTransactor) SetMaster(_self common.Address, _newMaster common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _PrototypeTransactor.Transact(opts, "setMaster", _self, _newMaster)
}

// SetMasterAsClause is a transaction clause generator 0x01378b58.
//
// Solidity: function setMaster(address _self, address _newMaster) returns()
func (_Prototype *Prototype) SetMasterAsClause(_self common.Address, _newMaster common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Prototype.contract.AsClauseWithVET(val, "setMaster", _self, _newMaster)
}

// Sponsor is a paid mutator transaction binding the contract method 0x766c4f37.
//
// Solidity: function sponsor(address _self) returns()
func (_PrototypeTransactor *PrototypeTransactor) Sponsor(_self common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _PrototypeTransactor.Transact(opts, "sponsor", _self)
}

// SponsorAsClause is a transaction clause generator 0x766c4f37.
//
// Solidity: function sponsor(address _self) returns()
func (_Prototype *Prototype) SponsorAsClause(_self common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Prototype.contract.AsClauseWithVET(val, "sponsor", _self)
}

// Unsponsor is a paid mutator transaction binding the contract method 0xcdd2a99f.
//
// Solidity: function unsponsor(address _self) returns()
func (_PrototypeTransactor *PrototypeTransactor) Unsponsor(_self common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _PrototypeTransactor.Transact(opts, "unsponsor", _self)
}

// UnsponsorAsClause is a transaction clause generator 0xcdd2a99f.
//
// Solidity: function unsponsor(address _self) returns()
func (_Prototype *Prototype) UnsponsorAsClause(_self common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Prototype.contract.AsClauseWithVET(val, "unsponsor", _self)
}
