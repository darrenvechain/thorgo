// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testcontract

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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unlockBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b506107b78061001f6000396000f3fe6080604052600436106100345760003560e01c80633ccfd60b14610039578063a694fc3a14610050578063fc7e286d1461006c575b600080fd5b34801561004557600080fd5b5061004e6100aa565b005b61006a60048036038101906100659190610478565b61029f565b005b34801561007857600080fd5b50610093600480360381019061008e9190610503565b610419565b6040516100a192919061053f565b60405180910390f35b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060400160405290816000820154815260200160018201548152505090504381600001511115610152576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610149906105c5565b60405180910390fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001018190555060003373ffffffffffffffffffffffffffffffffffffffff1682602001516040516101c390610616565b60006040518083038185875af1925050503d8060008114610200576040519150601f19603f3d011682016040523d82523d6000602084013e610205565b606091505b5050905080610249576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024090610677565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b6583602001516040516102939190610697565b60405180910390a25050565b600034116102e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102d9906106fe565b60405180910390fd5b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020604051806040016040529081600082015481526020016001820154815250509050814361034f919061074d565b816000018181525050348160200181815161036a919061074d565b91508181525050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000155602082015181600101559050503373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3460405161040d9190610697565b60405180910390a25050565b60006020528060005260406000206000915090508060000154908060010154905082565b600080fd5b6000819050919050565b61045581610442565b811461046057600080fd5b50565b6000813590506104728161044c565b92915050565b60006020828403121561048e5761048d61043d565b5b600061049c84828501610463565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006104d0826104a5565b9050919050565b6104e0816104c5565b81146104eb57600080fd5b50565b6000813590506104fd816104d7565b92915050565b6000602082840312156105195761051861043d565b5b6000610527848285016104ee565b91505092915050565b61053981610442565b82525050565b60006040820190506105546000830185610530565b6105616020830184610530565b9392505050565b600082825260208201905092915050565b7f596f752063616e27742077697468647261772079657400000000000000000000600082015250565b60006105af601683610568565b91506105ba82610579565b602082019050919050565b600060208201905081810360008301526105de816105a2565b9050919050565b600081905092915050565b50565b60006106006000836105e5565b915061060b826105f0565b600082019050919050565b6000610621826105f3565b9150819050919050565b7f5472616e73666572206661696c65642e00000000000000000000000000000000600082015250565b6000610661601083610568565b915061066c8261062b565b602082019050919050565b6000602082019050818103600083015261069081610654565b9050919050565b60006020820190506106ac6000830184610530565b92915050565b7f76616c756520697320656d707479000000000000000000000000000000000000600082015250565b60006106e8600e83610568565b91506106f3826106b2565b602082019050919050565b60006020820190508181036000830152610717816106db565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061075882610442565b915061076383610442565b925082820190508082111561077b5761077a61071e565b5b9291505056fea264697066735822122069fc707adb3355686afe9a3a4fd2c45422434b176050cbb0ed7bb841309a2d5964736f6c634300081c0033",
}

// DeployStaker deploys a new Ethereum contract, binding an instance of Staker to it.
func DeployStaker(ctx context.Context, thor *thorest.Client, sender contracts.TxManager, opts *transactions.Options) (common.Hash, *Staker, error) {
	parsed, err := StakerMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, nil, err
	}

	bytes, err := hexutil.Decode(StakerMetaData.Bin)
	if err != nil {
		return common.Hash{}, nil, err
	}
	contract, txID, err := contracts.NewDeployer(thor, bytes, parsed).Deploy(ctx, sender, opts)
	if err != nil {
		return common.Hash{}, nil, err
	}
	return txID, &Staker{thor: thor, contract: contract}, nil
}

// Staker is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Staker struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewStaker creates a new instance of Staker, bound to a specific deployed contract.
func NewStaker(address common.Address, thor *thorest.Client) (*Staker, error) {
	parsed, err := StakerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, address, parsed)
	return &Staker{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Staker *Staker) Address() common.Address {
	return _Staker.contract.Address
}

// ==================== View Functions ====================

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 unlockBlock, uint256 amount)
func (_Staker *Staker) Deposits(arg0 common.Address) *StakerDepositsCaller {
	return &StakerDepositsCaller{caller: _Staker.contract.Call("deposits", arg0)}
}

// ==================== Transaction Functions ====================

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 blocks) payable returns()
//
// Setting the value in options is replaced by the vetValue argument.
func (_Staker *Staker) Stake(blocks *big.Int, vetValue *big.Int) *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "stake", blocks).WithVET(vetValue)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Staker *Staker) Withdraw() *contracts.Sender {
	return contracts.NewSender(_Staker.contract, "withdraw")
}

// ==================== Event Functions ====================

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed _from, uint256 _value)
func (_Staker *Staker) FilterDeposit(criteria []StakerDepositCriteria) *StakerDepositFilterer {
	filterer := _Staker.contract.Filter("Deposit")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.From != nil {
			eventCriteria.Topic1 = *c.From
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerDepositFilterer{filterer: filterer, contract: _Staker.contract}
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed _to, uint256 _value)
func (_Staker *Staker) FilterWithdrawal(criteria []StakerWithdrawalCriteria) *StakerWithdrawalFilterer {
	filterer := _Staker.contract.Filter("Withdrawal")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.To != nil {
			eventCriteria.Topic1 = *c.To
		}
		filterer.Criteria(eventCriteria)
	}

	return &StakerWithdrawalFilterer{filterer: filterer, contract: _Staker.contract}
}

// ==================== Event Types and Criteria ====================

// StakerDeposit represents a Deposit event raised by the Staker contract.
type StakerDeposit struct {
	From  common.Address
	Value *big.Int
	Log   *thorest.EventLog
}

type StakerDepositCriteria struct {
	From *common.Address
}

// StakerWithdrawal represents a Withdrawal event raised by the Staker contract.
type StakerWithdrawal struct {
	To    common.Address
	Value *big.Int
	Log   *thorest.EventLog
}

type StakerWithdrawalCriteria struct {
	To *common.Address
}

// ==================== Call Result Types ====================

// StakerDepositsResult is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(uint256 unlockBlock, uint256 amount)
type StakerDepositsResult struct {
	UnlockBlock *big.Int
	Amount      *big.Int
}

// ==================== Caller Types and Methods ====================

// StakerDepositsCaller provides typed access to the Deposits method
type StakerDepositsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xfc7e286d.
func (c *StakerDepositsCaller) WithRevision(rev thorest.Revision) *StakerDepositsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xfc7e286d.
func (c *StakerDepositsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xfc7e286d and returns the result.
func (c *StakerDepositsCaller) Execute() (*StakerDepositsResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(StakerDepositsResult)
	out.UnlockBlock = *abi.ConvertType(data[0], new(*big.Int)).(**big.Int)
	out.Amount = *abi.ConvertType(data[1], new(*big.Int)).(**big.Int)

	return out, nil
}

// ==================== Event Filterer Types and Methods ====================

// StakerDepositFilterer provides typed access to filtering Deposit events
type StakerDepositFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerDepositFilterer) Unit(unit string) *StakerDepositFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerDepositFilterer) IncludeIndexes(include bool) *StakerDepositFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerDepositFilterer) Range(from, to int64) *StakerDepositFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerDepositFilterer) From(from int64) *StakerDepositFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerDepositFilterer) To(to int64) *StakerDepositFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerDepositFilterer) Offset(offset int64) *StakerDepositFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerDepositFilterer) Limit(limit int64) *StakerDepositFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerDepositFilterer) Order(order string) *StakerDepositFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerDepositFilterer) Execute() ([]StakerDeposit, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerDeposit, len(logs))
	for i, log := range logs {
		event := new(StakerDeposit)
		if err := f.contract.UnpackLog(event, "Deposit", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// StakerWithdrawalFilterer provides typed access to filtering Withdrawal events
type StakerWithdrawalFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *StakerWithdrawalFilterer) Unit(unit string) *StakerWithdrawalFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *StakerWithdrawalFilterer) IncludeIndexes(include bool) *StakerWithdrawalFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *StakerWithdrawalFilterer) Range(from, to int64) *StakerWithdrawalFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *StakerWithdrawalFilterer) From(from int64) *StakerWithdrawalFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *StakerWithdrawalFilterer) To(to int64) *StakerWithdrawalFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *StakerWithdrawalFilterer) Offset(offset int64) *StakerWithdrawalFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *StakerWithdrawalFilterer) Limit(limit int64) *StakerWithdrawalFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *StakerWithdrawalFilterer) Order(order string) *StakerWithdrawalFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *StakerWithdrawalFilterer) Execute() ([]StakerWithdrawal, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]StakerWithdrawal, len(logs))
	for i, log := range logs {
		event := new(StakerWithdrawal)
		if err := f.contract.UnpackLog(event, "Withdrawal", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}
