// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package builtins

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/contracts"
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
	_ = errors.Is
	_ = big.NewInt
	_ = strings.ReplaceAll
	_ = abi.ConvertType
	_ = hexutil.Decode
	_ = context.Background
	_ = tx.NewClause
	_ = blocks.New
	_ = time.Sleep
	_ = transactions.New
)

// PrototypeMetaData contains all meta data concerning the Prototype contract.
var PrototypeMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_newMaster\",\"type\":\"address\"}],\"name\":\"setMaster\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"storageFor\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"energy\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"currentSponsor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_credit\",\"type\":\"uint256\"},{\"name\":\"_recoveryRate\",\"type\":\"uint256\"}],\"name\":\"setCreditPlan\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_sponsor\",\"type\":\"address\"}],\"name\":\"selectSponsor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"balance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"sponsor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"creditPlan\",\"outputs\":[{\"name\":\"credit\",\"type\":\"uint256\"},{\"name\":\"recoveryRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addUser\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"hasCode\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"master\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"userCredit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"}],\"name\":\"unsponsor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_self\",\"type\":\"address\"},{\"name\":\"_sponsor\",\"type\":\"address\"}],\"name\":\"isSponsor\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Prototype is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Prototype struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewPrototype creates a new instance of Prototype, bound to a specific deployed contract.
func NewPrototype(thor *thorest.Client) (*Prototype, error) {
	parsed, err := PrototypeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, common.HexToAddress("0x000000000000000000000050726f746f74797065"), parsed)
	return &Prototype{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Prototype *Prototype) Address() common.Address {
	return _Prototype.contract.Address
}

// Balance is a free data retrieval call binding the contract method 0x6d8c859a.
//
// Solidity: function balance(address _self, uint256 _blockNumber) view returns(uint256)
func (_Prototype *Prototype) Balance(_self common.Address, _blockNumber *big.Int) *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_Prototype.contract, "balance", _self, _blockNumber)
}

// PrototypeCreditPlanResult is a free data retrieval call binding the contract method 0x80df45b4.
//
// Solidity: function creditPlan(address _self) view returns(uint256 credit, uint256 recoveryRate)
type PrototypeCreditPlanResult struct {
	Credit       *big.Int
	RecoveryRate *big.Int
}

func (_Prototype *Prototype) CreditPlan(_self common.Address) *contracts.Caller[*PrototypeCreditPlanResult] {
	parser := func(data []interface{}) (*PrototypeCreditPlanResult, error) {
		if len(data) != 2 {
			return nil, errors.New("invalid number of return values")
		}
		out := new(PrototypeCreditPlanResult)

		out.Credit = *abi.ConvertType(data[0], new(*big.Int)).(**big.Int)
		out.RecoveryRate = *abi.ConvertType(data[1], new(*big.Int)).(**big.Int)

		return out, nil
	}

	return contracts.NewCaller[*PrototypeCreditPlanResult](_Prototype.contract, "creditPlan", _self).WithParser(parser)
}

// CurrentSponsor is a free data retrieval call binding the contract method 0x23d8c7db.
//
// Solidity: function currentSponsor(address _self) view returns(address)
func (_Prototype *Prototype) CurrentSponsor(_self common.Address) *contracts.Caller[common.Address] {
	return contracts.NewCaller[common.Address](_Prototype.contract, "currentSponsor", _self)
}

// Energy is a free data retrieval call binding the contract method 0x1e95be45.
//
// Solidity: function energy(address _self, uint256 _blockNumber) view returns(uint256)
func (_Prototype *Prototype) Energy(_self common.Address, _blockNumber *big.Int) *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_Prototype.contract, "energy", _self, _blockNumber)
}

// HasCode is a free data retrieval call binding the contract method 0x9538c4b3.
//
// Solidity: function hasCode(address _self) view returns(bool)
func (_Prototype *Prototype) HasCode(_self common.Address) *contracts.Caller[bool] {
	return contracts.NewCaller[bool](_Prototype.contract, "hasCode", _self)
}

// IsSponsor is a free data retrieval call binding the contract method 0xd87333ac.
//
// Solidity: function isSponsor(address _self, address _sponsor) view returns(bool)
func (_Prototype *Prototype) IsSponsor(_self common.Address, _sponsor common.Address) *contracts.Caller[bool] {
	return contracts.NewCaller[bool](_Prototype.contract, "isSponsor", _self, _sponsor)
}

// IsUser is a free data retrieval call binding the contract method 0x02d43dc8.
//
// Solidity: function isUser(address _self, address _user) view returns(bool)
func (_Prototype *Prototype) IsUser(_self common.Address, _user common.Address) *contracts.Caller[bool] {
	return contracts.NewCaller[bool](_Prototype.contract, "isUser", _self, _user)
}

// Master is a free data retrieval call binding the contract method 0x9ed153c0.
//
// Solidity: function master(address _self) view returns(address)
func (_Prototype *Prototype) Master(_self common.Address) *contracts.Caller[common.Address] {
	return contracts.NewCaller[common.Address](_Prototype.contract, "master", _self)
}

// StorageFor is a free data retrieval call binding the contract method 0x04e7a457.
//
// Solidity: function storageFor(address _self, bytes32 _key) view returns(bytes32)
func (_Prototype *Prototype) StorageFor(_self common.Address, _key [32]byte) *contracts.Caller[[32]byte] {
	return contracts.NewCaller[[32]byte](_Prototype.contract, "storageFor", _self, _key)
}

// UserCredit is a free data retrieval call binding the contract method 0xc9c4fc41.
//
// Solidity: function userCredit(address _self, address _user) view returns(uint256)
func (_Prototype *Prototype) UserCredit(_self common.Address, _user common.Address) *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_Prototype.contract, "userCredit", _self, _user)
}

// AddUser is a paid mutator transaction binding the contract method 0x8ca3b448.
//
// Solidity: function addUser(address _self, address _user) returns()
func (_Prototype *Prototype) AddUser(_self common.Address, _user common.Address) *contracts.Sender {
	return contracts.NewSender(_Prototype.contract, "addUser", _self, _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x22928d6b.
//
// Solidity: function removeUser(address _self, address _user) returns()
func (_Prototype *Prototype) RemoveUser(_self common.Address, _user common.Address) *contracts.Sender {
	return contracts.NewSender(_Prototype.contract, "removeUser", _self, _user)
}

// SelectSponsor is a paid mutator transaction binding the contract method 0x3871a9fb.
//
// Solidity: function selectSponsor(address _self, address _sponsor) returns()
func (_Prototype *Prototype) SelectSponsor(_self common.Address, _sponsor common.Address) *contracts.Sender {
	return contracts.NewSender(_Prototype.contract, "selectSponsor", _self, _sponsor)
}

// SetCreditPlan is a paid mutator transaction binding the contract method 0x3659f8ed.
//
// Solidity: function setCreditPlan(address _self, uint256 _credit, uint256 _recoveryRate) returns()
func (_Prototype *Prototype) SetCreditPlan(_self common.Address, _credit *big.Int, _recoveryRate *big.Int) *contracts.Sender {
	return contracts.NewSender(_Prototype.contract, "setCreditPlan", _self, _credit, _recoveryRate)
}

// SetMaster is a paid mutator transaction binding the contract method 0x01378b58.
//
// Solidity: function setMaster(address _self, address _newMaster) returns()
func (_Prototype *Prototype) SetMaster(_self common.Address, _newMaster common.Address) *contracts.Sender {
	return contracts.NewSender(_Prototype.contract, "setMaster", _self, _newMaster)
}

// Sponsor is a paid mutator transaction binding the contract method 0x766c4f37.
//
// Solidity: function sponsor(address _self) returns()
func (_Prototype *Prototype) Sponsor(_self common.Address) *contracts.Sender {
	return contracts.NewSender(_Prototype.contract, "sponsor", _self)
}

// Unsponsor is a paid mutator transaction binding the contract method 0xcdd2a99f.
//
// Solidity: function unsponsor(address _self) returns()
func (_Prototype *Prototype) Unsponsor(_self common.Address) *contracts.Sender {
	return contracts.NewSender(_Prototype.contract, "unsponsor", _self)
}
