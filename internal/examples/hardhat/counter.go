// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package hardhat

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

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
	_ = time.Sleep
	_ = transactions.New
	_ = fmt.Errorf
)

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"Increment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"inc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"by\",\"type\":\"uint256\"}],\"name\":\"incBy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506103cf8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80630c55699c14610043578063371303c01461006157806370119d061461006b575b5f5ffd5b61004b610087565b6040516100589190610187565b60405180910390f35b61006961008c565b005b610085600480360381019061008091906101ce565b6100dc565b005b5f5481565b5f5f81548092919061009d90610226565b91905055507f51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a8160016040516100d291906102af565b60405180910390a1565b5f811161011e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161011590610348565b60405180910390fd5b805f5f82825461012e9190610366565b925050819055507f51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81816040516101649190610187565b60405180910390a150565b5f819050919050565b6101818161016f565b82525050565b5f60208201905061019a5f830184610178565b92915050565b5f5ffd5b6101ad8161016f565b81146101b7575f5ffd5b50565b5f813590506101c8816101a4565b92915050565b5f602082840312156101e3576101e26101a0565b5b5f6101f0848285016101ba565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6102308261016f565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610262576102616101f9565b5b600182019050919050565b5f819050919050565b5f819050919050565b5f61029961029461028f8461026d565b610276565b61016f565b9050919050565b6102a98161027f565b82525050565b5f6020820190506102c25f8301846102a0565b92915050565b5f82825260208201905092915050565b7f696e6342793a20696e6372656d656e742073686f756c6420626520706f7369745f8201527f6976650000000000000000000000000000000000000000000000000000000000602082015250565b5f6103326023836102c8565b915061033d826102d8565b604082019050919050565b5f6020820190508181035f83015261035f81610326565b9050919050565b5f6103708261016f565b915061037b8361016f565b9250828201905080821115610393576103926101f9565b5b9291505056fea26469706673582212208964bf78f883ff62c92c3c09416a09fa5a731da0068f08bc3384c0fdbe22f74f64736f6c634300081c0033",
}

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(ctx context.Context, thor *thorest.Client, sender contracts.TxManager, opts *transactions.Options) (common.Hash, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, nil, err
	}

	bytes, err := hexutil.Decode(CounterMetaData.Bin)
	if err != nil {
		return common.Hash{}, nil, err
	}
	contract, txID, err := contracts.NewDeployer(thor, bytes, parsed).Deploy(ctx, sender, opts)
	if err != nil {
		return common.Hash{}, nil, err
	}
	return txID, &Counter{thor: thor, contract: contract}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Counter struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, thor *thorest.Client) (*Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, address, parsed)
	return &Counter{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Counter *Counter) Address() common.Address {
	return _Counter.contract.Address
}

// Raw returns the underlying contract.
func (_Counter *Counter) Raw() *contracts.Contract {
	return _Counter.contract
}

// ==================== View Functions ====================

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() view returns(uint256)
func (_Counter *Counter) X() *CounterXCaller {
	return &CounterXCaller{caller: _Counter.contract.Call("x")}
}

// ==================== Transaction Functions ====================

// Inc is a paid mutator transaction binding the contract method 0x371303c0.
//
// Solidity: function inc() returns()
func (_Counter *Counter) Inc() *contracts.Sender {
	return contracts.NewSender(_Counter.contract, "inc")
}

// IncBy is a paid mutator transaction binding the contract method 0x70119d06.
//
// Solidity: function incBy(uint256 by) returns()
func (_Counter *Counter) IncBy(by *big.Int) *contracts.Sender {
	return contracts.NewSender(_Counter.contract, "incBy", by)
}

// ==================== Event Functions ====================

// UnpackIncrementLogs unpacks existing logs into typed Increment events.
func (_Counter *Counter) UnpackIncrementLogs(logs []*thorest.EventLog) ([]CounterIncrement, error) {
	events := make([]CounterIncrement, len(logs))
	for i, log := range logs {
		event := CounterIncrement{}
		if err := _Counter.contract.UnpackLog(&event, "Increment", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}
	return events, nil
}

// FilterIncrement is a free log retrieval operation binding the contract event 0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81.
//
// Solidity: event Increment(uint256 by)
func (_Counter *Counter) FilterIncrement() *CounterIncrementFilterer {
	filterer := _Counter.contract.Filter("Increment")

	return &CounterIncrementFilterer{filterer: filterer, contract: _Counter.contract}
}

// ==================== Event IDs ====================

// CounterIncrementEventID is the event ID for Increment
// Solidity: event Increment(uint256 by)
var CounterIncrementEventID = common.HexToHash("0x51af157c2eee40f68107a47a49c32fbbeb0a3c9e5cd37aa56e88e6be92368a81")

// ==================== Event Types and Criteria ====================

// CounterIncrement represents a Increment event raised by the Counter contract.
type CounterIncrement struct {
	By  *big.Int
	Log *thorest.EventLog
}

// ==================== Call Result Types ====================

// ==================== Caller Types and Methods ====================

// CounterXCaller provides typed access to the X method
type CounterXCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x0c55699c.
func (c *CounterXCaller) WithRevision(rev thorest.Revision) *CounterXCaller {
	return &CounterXCaller{caller: c.caller.WithRevision(rev)}
}

// Call executes the raw call to the contract method 0x0c55699c.
func (c *CounterXCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x0c55699c and returns the result.
func (c *CounterXCaller) Execute() (*big.Int, error) {
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

// ==================== Event Filterer Types and Methods ====================

// CounterIncrementFilterer provides typed access to filtering Increment events
type CounterIncrementFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *CounterIncrementFilterer) Unit(unit string) *CounterIncrementFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *CounterIncrementFilterer) IncludeIndexes(include bool) *CounterIncrementFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *CounterIncrementFilterer) Range(from, to int64) *CounterIncrementFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *CounterIncrementFilterer) From(from int64) *CounterIncrementFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *CounterIncrementFilterer) To(to int64) *CounterIncrementFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *CounterIncrementFilterer) Offset(offset int64) *CounterIncrementFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *CounterIncrementFilterer) Limit(limit int64) *CounterIncrementFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *CounterIncrementFilterer) Order(order string) *CounterIncrementFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *CounterIncrementFilterer) Execute() ([]CounterIncrement, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}
	return (&Counter{contract: f.contract}).UnpackIncrementLogs(logs)
}
