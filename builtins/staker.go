// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package builtins

import (
	"context"
	"errors"
	"fmt"
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
	_ = fmt.Errorf
)

// StakerMetaData contains all meta data concerning the Staker contract.
var StakerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"multiplier\",\"type\":\"uint8\"}],\"name\":\"DelegationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"}],\"name\":\"DelegationSignaledExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"DelegationWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"}],\"name\":\"StakeDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"added\",\"type\":\"uint256\"}],\"name\":\"StakeIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"endorser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ValidationQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidationSignaledExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ValidationWithdrawn\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"multiplier\",\"type\":\"uint8\"}],\"name\":\"addDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"}],\"name\":\"addValidation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"decreaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstActive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"first\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstQueued\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"first\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"}],\"name\":\"getDelegation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"multiplier\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isLocked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"}],\"name\":\"getDelegationPeriodDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"startPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"endPeriod\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"stakingPeriod\",\"type\":\"uint32\"}],\"name\":\"getDelegatorsRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"endorser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queuedVET\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"offlineBlock\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidationPeriodDetails\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"period\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"startBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"exitBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"completedPeriods\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidationTotals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedVET\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"queuedVET\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitingVET\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextPeriodWeight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidationsNum\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"activeCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"queuedCount\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"}],\"name\":\"getWithdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawableVET\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"increaseStake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"issued\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"prev\",\"type\":\"address\"}],\"name\":\"next\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nextValidation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queuedStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"queuedVET\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"}],\"name\":\"signalDelegationExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"signalExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalVET\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delegationID\",\"type\":\"uint256\"}],\"name\":\"withdrawDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// Staker is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Staker struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewStaker creates a new instance of Staker, bound to a specific deployed contract.
func NewStaker(thor *thorest.Client) (*Staker, error) {
	parsed, err := StakerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, common.HexToAddress("0x00000000000000000000000000005374616b6572"), parsed)
	return &Staker{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Staker *Staker) Address() common.Address {
	return _Staker.contract.Address
}

// Raw returns the underlying contract.
func (_Staker *Staker) Raw() *contracts.Contract {
	return _Staker.contract
}

// ==================== View Functions ====================

// FirstActive is a free data retrieval call binding the contract method 0xd719835c.
//
// Solidity: function firstActive() view returns(address first)
func (_Staker *Staker) FirstActive() *StakerFirstActiveCaller {
	return &StakerFirstActiveCaller{caller: _Staker.contract.Call("firstActive")}
}

// FirstQueued is a free data retrieval call binding the contract method 0xebe3a069.
//
// Solidity: function firstQueued() view returns(address first)
func (_Staker *Staker) FirstQueued() *StakerFirstQueuedCaller {
	return &StakerFirstQueuedCaller{caller: _Staker.contract.Call("firstQueued")}
}

// GetDelegation is a free data retrieval call binding the contract method 0x0dd35701.
//
// Solidity: function getDelegation(uint256 delegationID) view returns(address validator, uint256 stake, uint8 multiplier, bool isLocked)
func (_Staker *Staker) GetDelegation(delegationID *big.Int) *StakerGetDelegationCaller {
	return &StakerGetDelegationCaller{caller: _Staker.contract.Call("getDelegation", delegationID)}
}

// GetDelegationPeriodDetails is a free data retrieval call binding the contract method 0x0f4b6c66.
//
// Solidity: function getDelegationPeriodDetails(uint256 delegationID) view returns(uint32 startPeriod, uint32 endPeriod)
func (_Staker *Staker) GetDelegationPeriodDetails(delegationID *big.Int) *StakerGetDelegationPeriodDetailsCaller {
	return &StakerGetDelegationPeriodDetailsCaller{caller: _Staker.contract.Call("getDelegationPeriodDetails", delegationID)}
}

// GetDelegatorsRewards is a free data retrieval call binding the contract method 0xfddff039.
//
// Solidity: function getDelegatorsRewards(address validator, uint32 stakingPeriod) view returns(uint256 rewards)
func (_Staker *Staker) GetDelegatorsRewards(validator common.Address, stakingPeriod uint32) *StakerGetDelegatorsRewardsCaller {
	return &StakerGetDelegatorsRewardsCaller{caller: _Staker.contract.Call("getDelegatorsRewards", validator, stakingPeriod)}
}

// GetValidation is a free data retrieval call binding the contract method 0x27cd4de1.
//
// Solidity: function getValidation(address validator) view returns(address endorser, uint256 stake, uint256 weight, uint256 queuedVET, uint8 status, uint32 offlineBlock)
func (_Staker *Staker) GetValidation(validator common.Address) *StakerGetValidationCaller {
	return &StakerGetValidationCaller{caller: _Staker.contract.Call("getValidation", validator)}
}

// GetValidationPeriodDetails is a free data retrieval call binding the contract method 0x1a9e215a.
//
// Solidity: function getValidationPeriodDetails(address validator) view returns(uint32 period, uint32 startBlock, uint32 exitBlock, uint32 completedPeriods)
func (_Staker *Staker) GetValidationPeriodDetails(validator common.Address) *StakerGetValidationPeriodDetailsCaller {
	return &StakerGetValidationPeriodDetailsCaller{caller: _Staker.contract.Call("getValidationPeriodDetails", validator)}
}

// GetValidationTotals is a free data retrieval call binding the contract method 0x37d07860.
//
// Solidity: function getValidationTotals(address validator) view returns(uint256 lockedVET, uint256 lockedWeight, uint256 queuedVET, uint256 exitingVET, uint256 nextPeriodWeight)
func (_Staker *Staker) GetValidationTotals(validator common.Address) *StakerGetValidationTotalsCaller {
	return &StakerGetValidationTotalsCaller{caller: _Staker.contract.Call("getValidationTotals", validator)}
}

// GetValidationsNum is a free data retrieval call binding the contract method 0xaf8fef4e.
//
// Solidity: function getValidationsNum() view returns(uint64 activeCount, uint64 queuedCount)
func (_Staker *Staker) GetValidationsNum() *StakerGetValidationsNumCaller {
	return &StakerGetValidationsNumCaller{caller: _Staker.contract.Call("getValidationsNum")}
}

// GetWithdrawable is a free data retrieval call binding the contract method 0x32cc6ae6.
//
// Solidity: function getWithdrawable(address id) view returns(uint256 withdrawableVET)
func (_Staker *Staker) GetWithdrawable(id common.Address) *StakerGetWithdrawableCaller {
	return &StakerGetWithdrawableCaller{caller: _Staker.contract.Call("getWithdrawable", id)}
}

// Issuance is a free data retrieval call binding the contract method 0x863623bb.
//
// Solidity: function issuance() view returns(uint256 issued)
func (_Staker *Staker) Issuance() *StakerIssuanceCaller {
	return &StakerIssuanceCaller{caller: _Staker.contract.Call("issuance")}
}

// Next is a free data retrieval call binding the contract method 0xab73e316.
//
// Solidity: function next(address prev) view returns(address nextValidation)
func (_Staker *Staker) Next(prev common.Address) *StakerNextCaller {
	return &StakerNextCaller{caller: _Staker.contract.Call("next", prev)}
}

// QueuedStake is a free data retrieval call binding the contract method 0xe8e1a8b8.
//
// Solidity: function queuedStake() view returns(uint256 queuedVET)
func (_Staker *Staker) QueuedStake() *StakerQueuedStakeCaller {
	return &StakerQueuedStakeCaller{caller: _Staker.contract.Call("queuedStake")}
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 totalVET, uint256 totalWeight)
func (_Staker *Staker) TotalStake() *StakerTotalStakeCaller {
	return &StakerTotalStakeCaller{caller: _Staker.contract.Call("totalStake")}
}

// ==================== Transaction Functions ====================

// AddDelegation is a paid mutator transaction binding the contract method 0x4bf9842c.
//
// Solidity: function addDelegation(address validator, uint8 multiplier) payable returns(uint256 delegationID)
//
// Setting the value in options is replaced by the vetValue argument.
func (_Staker *Staker) AddDelegation(validator common.Address, multiplier uint8, vetValue *big.Int) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "addDelegation", validator, multiplier).WithVET(vetValue)
}

// AddValidation is a paid mutator transaction binding the contract method 0xc3c4b138.
//
// Solidity: function addValidation(address validator, uint32 period) payable returns()
//
// Setting the value in options is replaced by the vetValue argument.
func (_Staker *Staker) AddValidation(validator common.Address, period uint32, vetValue *big.Int) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "addValidation", validator, period).WithVET(vetValue)
}

// DecreaseStake is a paid mutator transaction binding the contract method 0x1a73ba01.
//
// Solidity: function decreaseStake(address validator, uint256 amount) returns()
func (_Staker *Staker) DecreaseStake(validator common.Address, amount *big.Int) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "decreaseStake", validator, amount)
}

// IncreaseStake is a paid mutator transaction binding the contract method 0x43b0de9a.
//
// Solidity: function increaseStake(address validator) payable returns()
//
// Setting the value in options is replaced by the vetValue argument.
func (_Staker *Staker) IncreaseStake(validator common.Address, vetValue *big.Int) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "increaseStake", validator).WithVET(vetValue)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0xf52564ec.
//
// Solidity: function setBeneficiary(address validator, address beneficiary) returns()
func (_Staker *Staker) SetBeneficiary(validator common.Address, beneficiary common.Address) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "setBeneficiary", validator, beneficiary)
}

// SignalDelegationExit is a paid mutator transaction binding the contract method 0xce88abbd.
//
// Solidity: function signalDelegationExit(uint256 delegationID) returns()
func (_Staker *Staker) SignalDelegationExit(delegationID *big.Int) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "signalDelegationExit", delegationID)
}

// SignalExit is a paid mutator transaction binding the contract method 0xcb652cef.
//
// Solidity: function signalExit(address validator) returns()
func (_Staker *Staker) SignalExit(validator common.Address) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "signalExit", validator)
}

// WithdrawDelegation is a paid mutator transaction binding the contract method 0xdf0e307a.
//
// Solidity: function withdrawDelegation(uint256 delegationID) returns()
func (_Staker *Staker) WithdrawDelegation(delegationID *big.Int) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "withdrawDelegation", delegationID)
}

// WithdrawStake is a paid mutator transaction binding the contract method 0xc23a5cea.
//
// Solidity: function withdrawStake(address validator) returns()
func (_Staker *Staker) WithdrawStake(validator common.Address) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "withdrawStake", validator)
}

// ==================== Event Functions ====================

// FilterBeneficiarySet is a free log retrieval operation binding the contract event 0x2906d223dc4163733bb374af8641c7e9ae256e2bae53c90e0c9a2be2e611ae44.
//
// Solidity: event BeneficiarySet(address indexed validator, address beneficiary)
func (_Staker *Staker) FilterBeneficiarySet(criteria []StakerBeneficiarySetCriteria) *StakerBeneficiarySetFilterer {
	filterer := _Staker.contract.Filter("BeneficiarySet")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Validator != nil {
			eventCriteria.Topic1 = *c.Validator
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerBeneficiarySetFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterDelegationAdded is a free log retrieval operation binding the contract event 0xe2eec2bad00ad70d7deab3014f1f737b6a5f8f1c948ef017b70ec34025fb4be5.
//
// Solidity: event DelegationAdded(address indexed validator, uint256 indexed delegationID, uint256 stake, uint8 multiplier)
func (_Staker *Staker) FilterDelegationAdded(criteria []StakerDelegationAddedCriteria) *StakerDelegationAddedFilterer {
	filterer := _Staker.contract.Filter("DelegationAdded")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Validator != nil {
			eventCriteria.Topic1 = *c.Validator
		}
		if c.DelegationID != nil {
			eventCriteria.Topic2 = c.DelegationID
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerDelegationAddedFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterDelegationSignaledExit is a free log retrieval operation binding the contract event 0xd5188446f0faca180c03ed7a73869f965c4e99eb4730b54c02544229ad45feb3.
//
// Solidity: event DelegationSignaledExit(uint256 indexed delegationID)
func (_Staker *Staker) FilterDelegationSignaledExit(criteria []StakerDelegationSignaledExitCriteria) *StakerDelegationSignaledExitFilterer {
	filterer := _Staker.contract.Filter("DelegationSignaledExit")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.DelegationID != nil {
			eventCriteria.Topic1 = c.DelegationID
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerDelegationSignaledExitFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterDelegationWithdrawn is a free log retrieval operation binding the contract event 0x0841064fa7f404b1ce2629504bda6d18463bb267aafc387ef9146f5bd0376dfc.
//
// Solidity: event DelegationWithdrawn(uint256 indexed delegationID, uint256 stake)
func (_Staker *Staker) FilterDelegationWithdrawn(criteria []StakerDelegationWithdrawnCriteria) *StakerDelegationWithdrawnFilterer {
	filterer := _Staker.contract.Filter("DelegationWithdrawn")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.DelegationID != nil {
			eventCriteria.Topic1 = c.DelegationID
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerDelegationWithdrawnFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterStakeDecreased is a free log retrieval operation binding the contract event 0x700865370ffb2a65a2b0242e6a64b21ac907ed5ecd46c9cffc729c177b2b1c69.
//
// Solidity: event StakeDecreased(address indexed validator, uint256 removed)
func (_Staker *Staker) FilterStakeDecreased(criteria []StakerStakeDecreasedCriteria) *StakerStakeDecreasedFilterer {
	filterer := _Staker.contract.Filter("StakeDecreased")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Validator != nil {
			eventCriteria.Topic1 = *c.Validator
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerStakeDecreasedFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterStakeIncreased is a free log retrieval operation binding the contract event 0x8b0ed825817a2e696c9a931715af4609fc60e1701f09c89ee7645130e937eb2d.
//
// Solidity: event StakeIncreased(address indexed validator, uint256 added)
func (_Staker *Staker) FilterStakeIncreased(criteria []StakerStakeIncreasedCriteria) *StakerStakeIncreasedFilterer {
	filterer := _Staker.contract.Filter("StakeIncreased")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Validator != nil {
			eventCriteria.Topic1 = *c.Validator
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerStakeIncreasedFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterValidationQueued is a free log retrieval operation binding the contract event 0x24082cc07dcc5c94fa94ab9cf2415cc4d8879a961b3d08f086d413fcde8d058c.
//
// Solidity: event ValidationQueued(address indexed validator, address indexed endorser, uint32 period, uint256 stake)
func (_Staker *Staker) FilterValidationQueued(criteria []StakerValidationQueuedCriteria) *StakerValidationQueuedFilterer {
	filterer := _Staker.contract.Filter("ValidationQueued")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Validator != nil {
			eventCriteria.Topic1 = *c.Validator
		}
		if c.Endorser != nil {
			eventCriteria.Topic2 = *c.Endorser
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerValidationQueuedFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterValidationSignaledExit is a free log retrieval operation binding the contract event 0xc42bb85f7e889d54497752aca8a1b93fb0a75d4664f4e463024c69ff6df56b4f.
//
// Solidity: event ValidationSignaledExit(address indexed validator)
func (_Staker *Staker) FilterValidationSignaledExit(criteria []StakerValidationSignaledExitCriteria) *StakerValidationSignaledExitFilterer {
	filterer := _Staker.contract.Filter("ValidationSignaledExit")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Validator != nil {
			eventCriteria.Topic1 = *c.Validator
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerValidationSignaledExitFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterValidationWithdrawn is a free log retrieval operation binding the contract event 0x8e0a6cffaa8510e8eac358ca73120eb97d99b0eadf08e4d965b717c71b6334ff.
//
// Solidity: event ValidationWithdrawn(address indexed validator, uint256 stake)
func (_Staker *Staker) FilterValidationWithdrawn(criteria []StakerValidationWithdrawnCriteria) *StakerValidationWithdrawnFilterer {
	filterer := _Staker.contract.Filter("ValidationWithdrawn")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Validator != nil {
			eventCriteria.Topic1 = *c.Validator
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerValidationWithdrawnFilterer{filterer: filterer, contract: _Staker.contract}
}

// ==================== Event Types and Criteria ====================

// StakerBeneficiarySet represents a BeneficiarySet event raised by the Staker contract.
type StakerBeneficiarySet struct {
	Validator   common.Address
	Beneficiary common.Address
	Log         *thorest.EventLog
}

type StakerBeneficiarySetCriteria struct {
	Validator *common.Address
}

// StakerDelegationAdded represents a DelegationAdded event raised by the Staker contract.
type StakerDelegationAdded struct {
	Validator    common.Address
	DelegationID *big.Int
	Stake        *big.Int
	Multiplier   uint8
	Log          *thorest.EventLog
}

type StakerDelegationAddedCriteria struct {
	Validator    *common.Address
	DelegationID *big.Int
}

// StakerDelegationSignaledExit represents a DelegationSignaledExit event raised by the Staker contract.
type StakerDelegationSignaledExit struct {
	DelegationID *big.Int
	Log          *thorest.EventLog
}

type StakerDelegationSignaledExitCriteria struct {
	DelegationID *big.Int
}

// StakerDelegationWithdrawn represents a DelegationWithdrawn event raised by the Staker contract.
type StakerDelegationWithdrawn struct {
	DelegationID *big.Int
	Stake        *big.Int
	Log          *thorest.EventLog
}

type StakerDelegationWithdrawnCriteria struct {
	DelegationID *big.Int
}

// StakerStakeDecreased represents a StakeDecreased event raised by the Staker contract.
type StakerStakeDecreased struct {
	Validator common.Address
	Removed   *big.Int
	Log       *thorest.EventLog
}

type StakerStakeDecreasedCriteria struct {
	Validator *common.Address
}

// StakerStakeIncreased represents a StakeIncreased event raised by the Staker contract.
type StakerStakeIncreased struct {
	Validator common.Address
	Added     *big.Int
	Log       *thorest.EventLog
}

type StakerStakeIncreasedCriteria struct {
	Validator *common.Address
}

// StakerValidationQueued represents a ValidationQueued event raised by the Staker contract.
type StakerValidationQueued struct {
	Validator common.Address
	Endorser  common.Address
	Period    uint32
	Stake     *big.Int
	Log       *thorest.EventLog
}

type StakerValidationQueuedCriteria struct {
	Validator *common.Address
	Endorser  *common.Address
}

// StakerValidationSignaledExit represents a ValidationSignaledExit event raised by the Staker contract.
type StakerValidationSignaledExit struct {
	Validator common.Address
	Log       *thorest.EventLog
}

type StakerValidationSignaledExitCriteria struct {
	Validator *common.Address
}

// StakerValidationWithdrawn represents a ValidationWithdrawn event raised by the Staker contract.
type StakerValidationWithdrawn struct {
	Validator common.Address
	Stake     *big.Int
	Log       *thorest.EventLog
}

type StakerValidationWithdrawnCriteria struct {
	Validator *common.Address
}

// ==================== Call Result Types ====================

// StakerGetDelegationResult is a free data retrieval call binding the contract method 0x0dd35701.
//
// Solidity: function getDelegation(uint256 delegationID) view returns(address validator, uint256 stake, uint8 multiplier, bool isLocked)
type StakerGetDelegationResult struct {
	Validator  common.Address
	Stake      *big.Int
	Multiplier uint8
	IsLocked   bool
}

// StakerGetDelegationPeriodDetailsResult is a free data retrieval call binding the contract method 0x0f4b6c66.
//
// Solidity: function getDelegationPeriodDetails(uint256 delegationID) view returns(uint32 startPeriod, uint32 endPeriod)
type StakerGetDelegationPeriodDetailsResult struct {
	StartPeriod uint32
	EndPeriod   uint32
}

// StakerGetValidationResult is a free data retrieval call binding the contract method 0x27cd4de1.
//
// Solidity: function getValidation(address validator) view returns(address endorser, uint256 stake, uint256 weight, uint256 queuedVET, uint8 status, uint32 offlineBlock)
type StakerGetValidationResult struct {
	Endorser     common.Address
	Stake        *big.Int
	Weight       *big.Int
	QueuedVET    *big.Int
	Status       uint8
	OfflineBlock uint32
}

// StakerGetValidationPeriodDetailsResult is a free data retrieval call binding the contract method 0x1a9e215a.
//
// Solidity: function getValidationPeriodDetails(address validator) view returns(uint32 period, uint32 startBlock, uint32 exitBlock, uint32 completedPeriods)
type StakerGetValidationPeriodDetailsResult struct {
	Period           uint32
	StartBlock       uint32
	ExitBlock        uint32
	CompletedPeriods uint32
}

// StakerGetValidationTotalsResult is a free data retrieval call binding the contract method 0x37d07860.
//
// Solidity: function getValidationTotals(address validator) view returns(uint256 lockedVET, uint256 lockedWeight, uint256 queuedVET, uint256 exitingVET, uint256 nextPeriodWeight)
type StakerGetValidationTotalsResult struct {
	LockedVET        *big.Int
	LockedWeight     *big.Int
	QueuedVET        *big.Int
	ExitingVET       *big.Int
	NextPeriodWeight *big.Int
}

// StakerGetValidationsNumResult is a free data retrieval call binding the contract method 0xaf8fef4e.
//
// Solidity: function getValidationsNum() view returns(uint64 activeCount, uint64 queuedCount)
type StakerGetValidationsNumResult struct {
	ActiveCount uint64
	QueuedCount uint64
}

// StakerTotalStakeResult is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256 totalVET, uint256 totalWeight)
type StakerTotalStakeResult struct {
	TotalVET    *big.Int
	TotalWeight *big.Int
}

// ==================== Caller Types and Methods ====================

// StakerFirstActiveCaller provides typed access to the FirstActive method
type StakerFirstActiveCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xd719835c.
func (c *StakerFirstActiveCaller) WithRevision(rev thorest.Revision) *StakerFirstActiveCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xd719835c.
func (c *StakerFirstActiveCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xd719835c and returns the result.
func (c *StakerFirstActiveCaller) Execute() (common.Address, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero common.Address
		return zero, err
	}
	if len(data) != 1 {
		var zero common.Address
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(common.Address); ok {
		return result, nil
	}
	var zero common.Address
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// StakerFirstQueuedCaller provides typed access to the FirstQueued method
type StakerFirstQueuedCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xebe3a069.
func (c *StakerFirstQueuedCaller) WithRevision(rev thorest.Revision) *StakerFirstQueuedCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xebe3a069.
func (c *StakerFirstQueuedCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xebe3a069 and returns the result.
func (c *StakerFirstQueuedCaller) Execute() (common.Address, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero common.Address
		return zero, err
	}
	if len(data) != 1 {
		var zero common.Address
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(common.Address); ok {
		return result, nil
	}
	var zero common.Address
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// StakerGetDelegationCaller provides typed access to the GetDelegation method
type StakerGetDelegationCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x0dd35701.
func (c *StakerGetDelegationCaller) WithRevision(rev thorest.Revision) *StakerGetDelegationCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x0dd35701.
func (c *StakerGetDelegationCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x0dd35701 and returns the result.
func (c *StakerGetDelegationCaller) Execute() (*StakerGetDelegationResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 4 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerGetDelegationResult)
	out.Validator = *abi.ConvertType(data[0], new(common.Address)).(*common.Address)
	out.Stake = *abi.ConvertType(data[1], new(*big.Int)).(**big.Int)
	out.Multiplier = *abi.ConvertType(data[2], new(uint8)).(*uint8)
	out.IsLocked = *abi.ConvertType(data[3], new(bool)).(*bool)

	return out, nil
}

// StakerGetDelegationPeriodDetailsCaller provides typed access to the GetDelegationPeriodDetails method
type StakerGetDelegationPeriodDetailsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x0f4b6c66.
func (c *StakerGetDelegationPeriodDetailsCaller) WithRevision(rev thorest.Revision) *StakerGetDelegationPeriodDetailsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x0f4b6c66.
func (c *StakerGetDelegationPeriodDetailsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x0f4b6c66 and returns the result.
func (c *StakerGetDelegationPeriodDetailsCaller) Execute() (*StakerGetDelegationPeriodDetailsResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerGetDelegationPeriodDetailsResult)
	out.StartPeriod = *abi.ConvertType(data[0], new(uint32)).(*uint32)
	out.EndPeriod = *abi.ConvertType(data[1], new(uint32)).(*uint32)

	return out, nil
}

// StakerGetDelegatorsRewardsCaller provides typed access to the GetDelegatorsRewards method
type StakerGetDelegatorsRewardsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xfddff039.
func (c *StakerGetDelegatorsRewardsCaller) WithRevision(rev thorest.Revision) *StakerGetDelegatorsRewardsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xfddff039.
func (c *StakerGetDelegatorsRewardsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xfddff039 and returns the result.
func (c *StakerGetDelegatorsRewardsCaller) Execute() (*big.Int, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero *big.Int
		return zero, err
	}
	if len(data) != 1 {
		var zero *big.Int
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(*big.Int); ok {
		return result, nil
	}
	var zero *big.Int
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// StakerGetValidationCaller provides typed access to the GetValidation method
type StakerGetValidationCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x27cd4de1.
func (c *StakerGetValidationCaller) WithRevision(rev thorest.Revision) *StakerGetValidationCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x27cd4de1.
func (c *StakerGetValidationCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x27cd4de1 and returns the result.
func (c *StakerGetValidationCaller) Execute() (*StakerGetValidationResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 6 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerGetValidationResult)
	out.Endorser = *abi.ConvertType(data[0], new(common.Address)).(*common.Address)
	out.Stake = *abi.ConvertType(data[1], new(*big.Int)).(**big.Int)
	out.Weight = *abi.ConvertType(data[2], new(*big.Int)).(**big.Int)
	out.QueuedVET = *abi.ConvertType(data[3], new(*big.Int)).(**big.Int)
	out.Status = *abi.ConvertType(data[4], new(uint8)).(*uint8)
	out.OfflineBlock = *abi.ConvertType(data[5], new(uint32)).(*uint32)

	return out, nil
}

// StakerGetValidationPeriodDetailsCaller provides typed access to the GetValidationPeriodDetails method
type StakerGetValidationPeriodDetailsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x1a9e215a.
func (c *StakerGetValidationPeriodDetailsCaller) WithRevision(rev thorest.Revision) *StakerGetValidationPeriodDetailsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x1a9e215a.
func (c *StakerGetValidationPeriodDetailsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x1a9e215a and returns the result.
func (c *StakerGetValidationPeriodDetailsCaller) Execute() (*StakerGetValidationPeriodDetailsResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 4 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerGetValidationPeriodDetailsResult)
	out.Period = *abi.ConvertType(data[0], new(uint32)).(*uint32)
	out.StartBlock = *abi.ConvertType(data[1], new(uint32)).(*uint32)
	out.ExitBlock = *abi.ConvertType(data[2], new(uint32)).(*uint32)
	out.CompletedPeriods = *abi.ConvertType(data[3], new(uint32)).(*uint32)

	return out, nil
}

// StakerGetValidationTotalsCaller provides typed access to the GetValidationTotals method
type StakerGetValidationTotalsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x37d07860.
func (c *StakerGetValidationTotalsCaller) WithRevision(rev thorest.Revision) *StakerGetValidationTotalsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x37d07860.
func (c *StakerGetValidationTotalsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x37d07860 and returns the result.
func (c *StakerGetValidationTotalsCaller) Execute() (*StakerGetValidationTotalsResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 5 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerGetValidationTotalsResult)
	out.LockedVET = *abi.ConvertType(data[0], new(*big.Int)).(**big.Int)
	out.LockedWeight = *abi.ConvertType(data[1], new(*big.Int)).(**big.Int)
	out.QueuedVET = *abi.ConvertType(data[2], new(*big.Int)).(**big.Int)
	out.ExitingVET = *abi.ConvertType(data[3], new(*big.Int)).(**big.Int)
	out.NextPeriodWeight = *abi.ConvertType(data[4], new(*big.Int)).(**big.Int)

	return out, nil
}

// StakerGetValidationsNumCaller provides typed access to the GetValidationsNum method
type StakerGetValidationsNumCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xaf8fef4e.
func (c *StakerGetValidationsNumCaller) WithRevision(rev thorest.Revision) *StakerGetValidationsNumCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xaf8fef4e.
func (c *StakerGetValidationsNumCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xaf8fef4e and returns the result.
func (c *StakerGetValidationsNumCaller) Execute() (*StakerGetValidationsNumResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerGetValidationsNumResult)
	out.ActiveCount = *abi.ConvertType(data[0], new(uint64)).(*uint64)
	out.QueuedCount = *abi.ConvertType(data[1], new(uint64)).(*uint64)

	return out, nil
}

// StakerGetWithdrawableCaller provides typed access to the GetWithdrawable method
type StakerGetWithdrawableCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x32cc6ae6.
func (c *StakerGetWithdrawableCaller) WithRevision(rev thorest.Revision) *StakerGetWithdrawableCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x32cc6ae6.
func (c *StakerGetWithdrawableCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x32cc6ae6 and returns the result.
func (c *StakerGetWithdrawableCaller) Execute() (*big.Int, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero *big.Int
		return zero, err
	}
	if len(data) != 1 {
		var zero *big.Int
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(*big.Int); ok {
		return result, nil
	}
	var zero *big.Int
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// StakerIssuanceCaller provides typed access to the Issuance method
type StakerIssuanceCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x863623bb.
func (c *StakerIssuanceCaller) WithRevision(rev thorest.Revision) *StakerIssuanceCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x863623bb.
func (c *StakerIssuanceCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x863623bb and returns the result.
func (c *StakerIssuanceCaller) Execute() (*big.Int, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero *big.Int
		return zero, err
	}
	if len(data) != 1 {
		var zero *big.Int
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(*big.Int); ok {
		return result, nil
	}
	var zero *big.Int
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// StakerNextCaller provides typed access to the Next method
type StakerNextCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xab73e316.
func (c *StakerNextCaller) WithRevision(rev thorest.Revision) *StakerNextCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xab73e316.
func (c *StakerNextCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xab73e316 and returns the result.
func (c *StakerNextCaller) Execute() (common.Address, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero common.Address
		return zero, err
	}
	if len(data) != 1 {
		var zero common.Address
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(common.Address); ok {
		return result, nil
	}
	var zero common.Address
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// StakerQueuedStakeCaller provides typed access to the QueuedStake method
type StakerQueuedStakeCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xe8e1a8b8.
func (c *StakerQueuedStakeCaller) WithRevision(rev thorest.Revision) *StakerQueuedStakeCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xe8e1a8b8.
func (c *StakerQueuedStakeCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xe8e1a8b8 and returns the result.
func (c *StakerQueuedStakeCaller) Execute() (*big.Int, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero *big.Int
		return zero, err
	}
	if len(data) != 1 {
		var zero *big.Int
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(*big.Int); ok {
		return result, nil
	}
	var zero *big.Int
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// StakerTotalStakeCaller provides typed access to the TotalStake method
type StakerTotalStakeCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x8b0e9f3f.
func (c *StakerTotalStakeCaller) WithRevision(rev thorest.Revision) *StakerTotalStakeCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x8b0e9f3f.
func (c *StakerTotalStakeCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x8b0e9f3f and returns the result.
func (c *StakerTotalStakeCaller) Execute() (*StakerTotalStakeResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerTotalStakeResult)
	out.TotalVET = *abi.ConvertType(data[0], new(*big.Int)).(**big.Int)
	out.TotalWeight = *abi.ConvertType(data[1], new(*big.Int)).(**big.Int)

	return out, nil
}

// ==================== Event Filterer Types and Methods ====================

// StakerBeneficiarySetFilterer provides typed access to filtering BeneficiarySet events
type StakerBeneficiarySetFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerBeneficiarySetFilterer) Unit(unit string) *StakerBeneficiarySetFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerBeneficiarySetFilterer) IncludeIndexes(include bool) *StakerBeneficiarySetFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerBeneficiarySetFilterer) Range(from, to int64) *StakerBeneficiarySetFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerBeneficiarySetFilterer) From(from int64) *StakerBeneficiarySetFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerBeneficiarySetFilterer) To(to int64) *StakerBeneficiarySetFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerBeneficiarySetFilterer) Offset(offset int64) *StakerBeneficiarySetFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerBeneficiarySetFilterer) Limit(limit int64) *StakerBeneficiarySetFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerBeneficiarySetFilterer) Order(order string) *StakerBeneficiarySetFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerBeneficiarySetFilterer) Execute() ([]StakerBeneficiarySet, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerBeneficiarySet, len(logs))
	for i, log := range logs {
		event := StakerBeneficiarySet{}
		if err := f.contract.UnpackLog(&event, "BeneficiarySet", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerDelegationAddedFilterer provides typed access to filtering DelegationAdded events
type StakerDelegationAddedFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerDelegationAddedFilterer) Unit(unit string) *StakerDelegationAddedFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerDelegationAddedFilterer) IncludeIndexes(include bool) *StakerDelegationAddedFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerDelegationAddedFilterer) Range(from, to int64) *StakerDelegationAddedFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerDelegationAddedFilterer) From(from int64) *StakerDelegationAddedFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerDelegationAddedFilterer) To(to int64) *StakerDelegationAddedFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerDelegationAddedFilterer) Offset(offset int64) *StakerDelegationAddedFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerDelegationAddedFilterer) Limit(limit int64) *StakerDelegationAddedFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerDelegationAddedFilterer) Order(order string) *StakerDelegationAddedFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerDelegationAddedFilterer) Execute() ([]StakerDelegationAdded, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerDelegationAdded, len(logs))
	for i, log := range logs {
		event := StakerDelegationAdded{}
		if err := f.contract.UnpackLog(&event, "DelegationAdded", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerDelegationSignaledExitFilterer provides typed access to filtering DelegationSignaledExit events
type StakerDelegationSignaledExitFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerDelegationSignaledExitFilterer) Unit(unit string) *StakerDelegationSignaledExitFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerDelegationSignaledExitFilterer) IncludeIndexes(include bool) *StakerDelegationSignaledExitFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerDelegationSignaledExitFilterer) Range(from, to int64) *StakerDelegationSignaledExitFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerDelegationSignaledExitFilterer) From(from int64) *StakerDelegationSignaledExitFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerDelegationSignaledExitFilterer) To(to int64) *StakerDelegationSignaledExitFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerDelegationSignaledExitFilterer) Offset(offset int64) *StakerDelegationSignaledExitFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerDelegationSignaledExitFilterer) Limit(limit int64) *StakerDelegationSignaledExitFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerDelegationSignaledExitFilterer) Order(order string) *StakerDelegationSignaledExitFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerDelegationSignaledExitFilterer) Execute() ([]StakerDelegationSignaledExit, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerDelegationSignaledExit, len(logs))
	for i, log := range logs {
		event := StakerDelegationSignaledExit{}
		if err := f.contract.UnpackLog(&event, "DelegationSignaledExit", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerDelegationWithdrawnFilterer provides typed access to filtering DelegationWithdrawn events
type StakerDelegationWithdrawnFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerDelegationWithdrawnFilterer) Unit(unit string) *StakerDelegationWithdrawnFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerDelegationWithdrawnFilterer) IncludeIndexes(include bool) *StakerDelegationWithdrawnFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerDelegationWithdrawnFilterer) Range(from, to int64) *StakerDelegationWithdrawnFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerDelegationWithdrawnFilterer) From(from int64) *StakerDelegationWithdrawnFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerDelegationWithdrawnFilterer) To(to int64) *StakerDelegationWithdrawnFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerDelegationWithdrawnFilterer) Offset(offset int64) *StakerDelegationWithdrawnFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerDelegationWithdrawnFilterer) Limit(limit int64) *StakerDelegationWithdrawnFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerDelegationWithdrawnFilterer) Order(order string) *StakerDelegationWithdrawnFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerDelegationWithdrawnFilterer) Execute() ([]StakerDelegationWithdrawn, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerDelegationWithdrawn, len(logs))
	for i, log := range logs {
		event := StakerDelegationWithdrawn{}
		if err := f.contract.UnpackLog(&event, "DelegationWithdrawn", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerStakeDecreasedFilterer provides typed access to filtering StakeDecreased events
type StakerStakeDecreasedFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerStakeDecreasedFilterer) Unit(unit string) *StakerStakeDecreasedFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerStakeDecreasedFilterer) IncludeIndexes(include bool) *StakerStakeDecreasedFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerStakeDecreasedFilterer) Range(from, to int64) *StakerStakeDecreasedFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerStakeDecreasedFilterer) From(from int64) *StakerStakeDecreasedFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerStakeDecreasedFilterer) To(to int64) *StakerStakeDecreasedFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerStakeDecreasedFilterer) Offset(offset int64) *StakerStakeDecreasedFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerStakeDecreasedFilterer) Limit(limit int64) *StakerStakeDecreasedFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerStakeDecreasedFilterer) Order(order string) *StakerStakeDecreasedFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerStakeDecreasedFilterer) Execute() ([]StakerStakeDecreased, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerStakeDecreased, len(logs))
	for i, log := range logs {
		event := StakerStakeDecreased{}
		if err := f.contract.UnpackLog(&event, "StakeDecreased", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerStakeIncreasedFilterer provides typed access to filtering StakeIncreased events
type StakerStakeIncreasedFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerStakeIncreasedFilterer) Unit(unit string) *StakerStakeIncreasedFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerStakeIncreasedFilterer) IncludeIndexes(include bool) *StakerStakeIncreasedFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerStakeIncreasedFilterer) Range(from, to int64) *StakerStakeIncreasedFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerStakeIncreasedFilterer) From(from int64) *StakerStakeIncreasedFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerStakeIncreasedFilterer) To(to int64) *StakerStakeIncreasedFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerStakeIncreasedFilterer) Offset(offset int64) *StakerStakeIncreasedFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerStakeIncreasedFilterer) Limit(limit int64) *StakerStakeIncreasedFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerStakeIncreasedFilterer) Order(order string) *StakerStakeIncreasedFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerStakeIncreasedFilterer) Execute() ([]StakerStakeIncreased, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerStakeIncreased, len(logs))
	for i, log := range logs {
		event := StakerStakeIncreased{}
		if err := f.contract.UnpackLog(&event, "StakeIncreased", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerValidationQueuedFilterer provides typed access to filtering ValidationQueued events
type StakerValidationQueuedFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerValidationQueuedFilterer) Unit(unit string) *StakerValidationQueuedFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerValidationQueuedFilterer) IncludeIndexes(include bool) *StakerValidationQueuedFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerValidationQueuedFilterer) Range(from, to int64) *StakerValidationQueuedFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerValidationQueuedFilterer) From(from int64) *StakerValidationQueuedFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerValidationQueuedFilterer) To(to int64) *StakerValidationQueuedFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerValidationQueuedFilterer) Offset(offset int64) *StakerValidationQueuedFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerValidationQueuedFilterer) Limit(limit int64) *StakerValidationQueuedFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerValidationQueuedFilterer) Order(order string) *StakerValidationQueuedFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerValidationQueuedFilterer) Execute() ([]StakerValidationQueued, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerValidationQueued, len(logs))
	for i, log := range logs {
		event := StakerValidationQueued{}
		if err := f.contract.UnpackLog(&event, "ValidationQueued", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerValidationSignaledExitFilterer provides typed access to filtering ValidationSignaledExit events
type StakerValidationSignaledExitFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerValidationSignaledExitFilterer) Unit(unit string) *StakerValidationSignaledExitFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerValidationSignaledExitFilterer) IncludeIndexes(include bool) *StakerValidationSignaledExitFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerValidationSignaledExitFilterer) Range(from, to int64) *StakerValidationSignaledExitFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerValidationSignaledExitFilterer) From(from int64) *StakerValidationSignaledExitFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerValidationSignaledExitFilterer) To(to int64) *StakerValidationSignaledExitFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerValidationSignaledExitFilterer) Offset(offset int64) *StakerValidationSignaledExitFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerValidationSignaledExitFilterer) Limit(limit int64) *StakerValidationSignaledExitFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerValidationSignaledExitFilterer) Order(order string) *StakerValidationSignaledExitFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerValidationSignaledExitFilterer) Execute() ([]StakerValidationSignaledExit, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerValidationSignaledExit, len(logs))
	for i, log := range logs {
		event := StakerValidationSignaledExit{}
		if err := f.contract.UnpackLog(&event, "ValidationSignaledExit", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// StakerValidationWithdrawnFilterer provides typed access to filtering ValidationWithdrawn events
type StakerValidationWithdrawnFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerValidationWithdrawnFilterer) Unit(unit string) *StakerValidationWithdrawnFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerValidationWithdrawnFilterer) IncludeIndexes(include bool) *StakerValidationWithdrawnFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerValidationWithdrawnFilterer) Range(from, to int64) *StakerValidationWithdrawnFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerValidationWithdrawnFilterer) From(from int64) *StakerValidationWithdrawnFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerValidationWithdrawnFilterer) To(to int64) *StakerValidationWithdrawnFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerValidationWithdrawnFilterer) Offset(offset int64) *StakerValidationWithdrawnFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerValidationWithdrawnFilterer) Limit(limit int64) *StakerValidationWithdrawnFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerValidationWithdrawnFilterer) Order(order string) *StakerValidationWithdrawnFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerValidationWithdrawnFilterer) Execute() ([]StakerValidationWithdrawn, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerValidationWithdrawn, len(logs))
	for i, log := range logs {
		event := StakerValidationWithdrawn{}
		if err := f.contract.UnpackLog(&event, "ValidationWithdrawn", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}
