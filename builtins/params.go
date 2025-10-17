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

// ParamsMetaData contains all meta data concerning the Params contract.
var ParamsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Set\",\"type\":\"event\"}]",
}

// Params is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Params struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewParams creates a new instance of Params, bound to a specific deployed contract.
func NewParams(thor *thorest.Client) (*Params, error) {
	parsed, err := ParamsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, common.HexToAddress("0x0000000000000000000000000000506172616d73"), parsed)
	return &Params{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Params *Params) Address() common.Address {
	return _Params.contract.Address
}

// Raw returns the underlying contract.
func (_Params *Params) Raw() *contracts.Contract {
	return _Params.contract
}

// ==================== View Functions ====================

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Params *Params) Executor() *ParamsExecutorCaller {
	return &ParamsExecutorCaller{caller: _Params.contract.Call("executor")}
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _key) view returns(uint256)
func (_Params *Params) Get(_key [32]byte) *ParamsGetCaller {
	return &ParamsGetCaller{caller: _Params.contract.Call("get", _key)}
}

// ==================== Transaction Functions ====================

// Set is a paid mutator transaction binding the contract method 0x273f4940.
//
// Solidity: function set(bytes32 _key, uint256 _value) returns()
func (_Params *Params) Set(_key [32]byte, _value *big.Int) *contracts.Sender {
	return contracts.NewSender(_Params.contract, "set", _key, _value)
}

// ==================== Event Functions ====================

// FilterSet is a free log retrieval operation binding the contract event 0x28e3246f80515f5c1ed987b133ef2f193439b25acba6a5e69f219e896fc9d179.
//
// Solidity: event Set(bytes32 indexed key, uint256 value)
func (_Params *Params) FilterSet(criteria []ParamsSetCriteria) *ParamsSetFilterer {
	filterer := _Params.contract.Filter("Set")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Key != nil {
			eventCriteria.Topic1 = *c.Key
		}
		filterer.Criteria(eventCriteria)
	}

	return &ParamsSetFilterer{filterer: filterer, contract: _Params.contract}
}

// ==================== Event Types and Criteria ====================

// ParamsSet represents a Set event raised by the Params contract.
type ParamsSet struct {
	Key   [32]byte
	Value *big.Int
	Log   *thorest.EventLog
}

type ParamsSetCriteria struct {
	Key *[32]byte
}

// ==================== Call Result Types ====================

// ==================== Caller Types and Methods ====================

// ParamsExecutorCaller provides typed access to the Executor method
type ParamsExecutorCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xc34c08e5.
func (c *ParamsExecutorCaller) WithRevision(rev thorest.Revision) *ParamsExecutorCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xc34c08e5.
func (c *ParamsExecutorCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xc34c08e5 and returns the result.
func (c *ParamsExecutorCaller) Execute() (common.Address, error) {
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

// ParamsGetCaller provides typed access to the Get method
type ParamsGetCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x8eaa6ac0.
func (c *ParamsGetCaller) WithRevision(rev thorest.Revision) *ParamsGetCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x8eaa6ac0.
func (c *ParamsGetCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x8eaa6ac0 and returns the result.
func (c *ParamsGetCaller) Execute() (*big.Int, error) {
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

// ParamsSetFilterer provides typed access to filtering Set events
type ParamsSetFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *ParamsSetFilterer) Unit(unit string) *ParamsSetFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *ParamsSetFilterer) IncludeIndexes(include bool) *ParamsSetFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *ParamsSetFilterer) Range(from, to int64) *ParamsSetFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *ParamsSetFilterer) From(from int64) *ParamsSetFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *ParamsSetFilterer) To(to int64) *ParamsSetFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *ParamsSetFilterer) Offset(offset int64) *ParamsSetFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *ParamsSetFilterer) Limit(limit int64) *ParamsSetFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *ParamsSetFilterer) Order(order string) *ParamsSetFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *ParamsSetFilterer) Execute() ([]ParamsSet, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]ParamsSet, len(logs))
	for i, log := range logs {
		event := ParamsSet{}
		if err := f.contract.UnpackLog(&event, "Set", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}
