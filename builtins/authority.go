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

// AuthorityMetaData contains all meta data concerning the Authority contract.
var AuthorityMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"first\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeMaster\",\"type\":\"address\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeMaster\",\"type\":\"address\"}],\"name\":\"next\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_nodeMaster\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"listed\",\"type\":\"bool\"},{\"name\":\"endorsor\",\"type\":\"address\"},{\"name\":\"identity\",\"type\":\"bytes32\"},{\"name\":\"active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_nodeMaster\",\"type\":\"address\"},{\"name\":\"_endorsor\",\"type\":\"address\"},{\"name\":\"_identity\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"nodeMaster\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"Candidate\",\"type\":\"event\"}]",
}

// Authority is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Authority struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewAuthority creates a new instance of Authority, bound to a specific deployed contract.
func NewAuthority(thor *thorest.Client) (*Authority, error) {
	parsed, err := AuthorityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, common.HexToAddress("0x0000000000000000000000417574686f72697479"), parsed)
	return &Authority{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Authority *Authority) Address() common.Address {
	return _Authority.contract.Address
}

// Raw returns the underlying contract.
func (_Authority *Authority) Raw() *contracts.Contract {
	return _Authority.contract
}

// ==================== View Functions ====================

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Authority *Authority) Executor() *AuthorityExecutorCaller {
	return &AuthorityExecutorCaller{caller: _Authority.contract.Call("executor")}
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(address)
func (_Authority *Authority) First() *AuthorityFirstCaller {
	return &AuthorityFirstCaller{caller: _Authority.contract.Call("first")}
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address _nodeMaster) view returns(bool listed, address endorsor, bytes32 identity, bool active)
func (_Authority *Authority) Get(_nodeMaster common.Address) *AuthorityGetCaller {
	return &AuthorityGetCaller{caller: _Authority.contract.Call("get", _nodeMaster)}
}

// Next is a free data retrieval call binding the contract method 0xab73e316.
//
// Solidity: function next(address _nodeMaster) view returns(address)
func (_Authority *Authority) Next(_nodeMaster common.Address) *AuthorityNextCaller {
	return &AuthorityNextCaller{caller: _Authority.contract.Call("next", _nodeMaster)}
}

// ==================== Transaction Functions ====================

// Add is a paid mutator transaction binding the contract method 0xdc0094b8.
//
// Solidity: function add(address _nodeMaster, address _endorsor, bytes32 _identity) returns()
func (_Authority *Authority) Add(_nodeMaster common.Address, _endorsor common.Address, _identity [32]byte) *contracts.Sender {
	return contracts.NewSender(_Authority.contract, "add", _nodeMaster, _endorsor, _identity)
}

// Revoke is a paid mutator transaction binding the contract method 0x74a8f103.
//
// Solidity: function revoke(address _nodeMaster) returns()
func (_Authority *Authority) Revoke(_nodeMaster common.Address) *contracts.Sender {
	return contracts.NewSender(_Authority.contract, "revoke", _nodeMaster)
}

// ==================== Event Functions ====================

// UnpackCandidateLogs unpacks existing logs into typed Candidate events.
func (_Authority *Authority) UnpackCandidateLogs(logs []*thorest.EventLog) ([]AuthorityCandidate, error) {
	events := make([]AuthorityCandidate, len(logs))
	for i, log := range logs {
		event := AuthorityCandidate{}
		if err := _Authority.contract.UnpackLog(&event, "Candidate", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}
	return events, nil
}

// FilterCandidate is a free log retrieval operation binding the contract event 0xe9e2ad484aeae75ba75479c19d2cbb784b98b2fe4b24dc80a4c8cf142d4c9294.
//
// Solidity: event Candidate(address indexed nodeMaster, bytes32 action)
func (_Authority *Authority) FilterCandidate(criteria []AuthorityCandidateCriteria) *AuthorityCandidateFilterer {
	filterer := _Authority.contract.Filter("Candidate")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.NodeMaster != nil {
			eventCriteria.Topic1 = *c.NodeMaster
		}
		filterer.Criteria(eventCriteria)
	}

	return &AuthorityCandidateFilterer{filterer: filterer, contract: _Authority.contract}
}

// ==================== Event IDs ====================

// AuthorityCandidateEventID is the event ID for Candidate
// Solidity: event Candidate(address indexed nodeMaster, bytes32 action)
var AuthorityCandidateEventID = common.HexToHash("0xe9e2ad484aeae75ba75479c19d2cbb784b98b2fe4b24dc80a4c8cf142d4c9294")

// ==================== Event Types and Criteria ====================

// AuthorityCandidate represents a Candidate event raised by the Authority contract.
type AuthorityCandidate struct {
	NodeMaster common.Address
	Action     [32]byte
	Log        *thorest.EventLog
}

type AuthorityCandidateCriteria struct {
	NodeMaster *common.Address
}

// ==================== Call Result Types ====================

// AuthorityGetResult is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address _nodeMaster) view returns(bool listed, address endorsor, bytes32 identity, bool active)
type AuthorityGetResult struct {
	Listed   bool
	Endorsor common.Address
	Identity [32]byte
	Active   bool
}

// ==================== Caller Types and Methods ====================

// AuthorityExecutorCaller provides typed access to the Executor method
type AuthorityExecutorCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xc34c08e5.
func (c *AuthorityExecutorCaller) WithRevision(rev thorest.Revision) *AuthorityExecutorCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xc34c08e5.
func (c *AuthorityExecutorCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xc34c08e5 and returns the result.
func (c *AuthorityExecutorCaller) Execute() (common.Address, error) {
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

// AuthorityFirstCaller provides typed access to the First method
type AuthorityFirstCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x3df4ddf4.
func (c *AuthorityFirstCaller) WithRevision(rev thorest.Revision) *AuthorityFirstCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x3df4ddf4.
func (c *AuthorityFirstCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x3df4ddf4 and returns the result.
func (c *AuthorityFirstCaller) Execute() (common.Address, error) {
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

// AuthorityGetCaller provides typed access to the Get method
type AuthorityGetCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xc2bc2efc.
func (c *AuthorityGetCaller) WithRevision(rev thorest.Revision) *AuthorityGetCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xc2bc2efc.
func (c *AuthorityGetCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xc2bc2efc and returns the result.
func (c *AuthorityGetCaller) Execute() (*AuthorityGetResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 4 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(AuthorityGetResult)
	out.Listed = *abi.ConvertType(data[0], new(bool)).(*bool)
	out.Endorsor = *abi.ConvertType(data[1], new(common.Address)).(*common.Address)
	out.Identity = *abi.ConvertType(data[2], new([32]byte)).(*[32]byte)
	out.Active = *abi.ConvertType(data[3], new(bool)).(*bool)

	return out, nil
}

// AuthorityNextCaller provides typed access to the Next method
type AuthorityNextCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xab73e316.
func (c *AuthorityNextCaller) WithRevision(rev thorest.Revision) *AuthorityNextCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xab73e316.
func (c *AuthorityNextCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xab73e316 and returns the result.
func (c *AuthorityNextCaller) Execute() (common.Address, error) {
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

// ==================== Event Filterer Types and Methods ====================

// AuthorityCandidateFilterer provides typed access to filtering Candidate events
type AuthorityCandidateFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *AuthorityCandidateFilterer) Unit(unit string) *AuthorityCandidateFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *AuthorityCandidateFilterer) IncludeIndexes(include bool) *AuthorityCandidateFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *AuthorityCandidateFilterer) Range(from, to int64) *AuthorityCandidateFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *AuthorityCandidateFilterer) From(from int64) *AuthorityCandidateFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *AuthorityCandidateFilterer) To(to int64) *AuthorityCandidateFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *AuthorityCandidateFilterer) Offset(offset int64) *AuthorityCandidateFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *AuthorityCandidateFilterer) Limit(limit int64) *AuthorityCandidateFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *AuthorityCandidateFilterer) Order(order string) *AuthorityCandidateFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *AuthorityCandidateFilterer) Execute() ([]AuthorityCandidate, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}
	return (&Authority{contract: f.contract}).UnpackCandidateLogs(logs)
}
