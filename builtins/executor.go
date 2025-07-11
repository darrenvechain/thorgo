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

// ExecutorMetaData contains all meta data concerning the Executor contract.
var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvers\",\"outputs\":[{\"name\":\"identity\",\"type\":\"bytes32\"},{\"name\":\"inPower\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"approverCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approver\",\"type\":\"address\"}],\"name\":\"revokeApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"timeProposed\",\"type\":\"uint64\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"quorum\",\"type\":\"uint8\"},{\"name\":\"approvalCount\",\"type\":\"uint8\"},{\"name\":\"executed\",\"type\":\"bool\"},{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approver\",\"type\":\"address\"},{\"name\":\"_identity\",\"type\":\"bytes32\"}],\"name\":\"addApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"propose\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"attachVotingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"bytes32\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"detachVotingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votingContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"Proposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"Approver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"VotingContract\",\"type\":\"event\"}]",
}

// Executor is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Executor struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewExecutor creates a new instance of Executor, bound to a specific deployed contract.
func NewExecutor(thor *thorest.Client) (*Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, common.HexToAddress("0x0000000000000000000000004578656375746f72"), parsed)
	return &Executor{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_Executor *Executor) Address() common.Address {
	return _Executor.contract.Address
}

// ==================== View Functions ====================

// ApproverCount is a free data retrieval call binding the contract method 0x128e9be6.
//
// Solidity: function approverCount() view returns(uint8)
func (_Executor *Executor) ApproverCount() *ExecutorApproverCountCaller {
	return &ExecutorApproverCountCaller{caller: _Executor.contract.Call("approverCount")}
}

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bytes32 identity, bool inPower)
func (_Executor *Executor) Approvers(arg0 common.Address) *ExecutorApproversCaller {
	return &ExecutorApproversCaller{caller: _Executor.contract.Call("approvers", arg0)}
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint64 timeProposed, address proposer, uint8 quorum, uint8 approvalCount, bool executed, address target, bytes data)
func (_Executor *Executor) Proposals(arg0 [32]byte) *ExecutorProposalsCaller {
	return &ExecutorProposalsCaller{caller: _Executor.contract.Call("proposals", arg0)}
}

// VotingContracts is a free data retrieval call binding the contract method 0xfa06792b.
//
// Solidity: function votingContracts(address ) view returns(bool)
func (_Executor *Executor) VotingContracts(arg0 common.Address) *ExecutorVotingContractsCaller {
	return &ExecutorVotingContractsCaller{caller: _Executor.contract.Call("votingContracts", arg0)}
}

// ==================== Transaction Functions ====================

// AddApprover is a paid mutator transaction binding the contract method 0x3ef0c09e.
//
// Solidity: function addApprover(address _approver, bytes32 _identity) returns()
func (_Executor *Executor) AddApprover(_approver common.Address, _identity [32]byte) *contracts.Sender {
	return contracts.NewSender(_Executor.contract, "addApprover", _approver, _identity)
}

// Approve is a paid mutator transaction binding the contract method 0xa53a1adf.
//
// Solidity: function approve(bytes32 _proposalID) returns()
func (_Executor *Executor) Approve(_proposalID [32]byte) *contracts.Sender {
	return contracts.NewSender(_Executor.contract, "approve", _proposalID)
}

// AttachVotingContract is a paid mutator transaction binding the contract method 0xa1fb668f.
//
// Solidity: function attachVotingContract(address _contract) returns()
func (_Executor *Executor) AttachVotingContract(_contract common.Address) *contracts.Sender {
	return contracts.NewSender(_Executor.contract, "attachVotingContract", _contract)
}

// DetachVotingContract is a paid mutator transaction binding the contract method 0xa83b3bd8.
//
// Solidity: function detachVotingContract(address _contract) returns()
func (_Executor *Executor) DetachVotingContract(_contract common.Address) *contracts.Sender {
	return contracts.NewSender(_Executor.contract, "detachVotingContract", _contract)
}

// Execute is a paid mutator transaction binding the contract method 0xe751f271.
//
// Solidity: function execute(bytes32 _proposalID) returns()
func (_Executor *Executor) Execute(_proposalID [32]byte) *contracts.Sender {
	return contracts.NewSender(_Executor.contract, "execute", _proposalID)
}

// Propose is a paid mutator transaction binding the contract method 0x9d481848.
//
// Solidity: function propose(address _target, bytes _data) returns(bytes32)
func (_Executor *Executor) Propose(_target common.Address, _data []byte) *contracts.Sender {
	return contracts.NewSender(_Executor.contract, "propose", _target, _data)
}

// RevokeApprover is a paid mutator transaction binding the contract method 0x18d13ef7.
//
// Solidity: function revokeApprover(address _approver) returns()
func (_Executor *Executor) RevokeApprover(_approver common.Address) *contracts.Sender {
	return contracts.NewSender(_Executor.contract, "revokeApprover", _approver)
}

// ==================== Event Functions ====================

// FilterApprover is a free log retrieval operation binding the contract event 0x770115cde75e60f17b265d7e0c5e39c57abf243bc316c7e5c2f8d851771da6ac.
//
// Solidity: event Approver(address indexed approver, bytes32 action)
func (_Executor *Executor) FilterApprover(criteria []ExecutorApproverCriteria) *ExecutorApproverFilterer {
	filterer := _Executor.contract.Filter("Approver")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := contracts.EventCriteria{}
		if c.Approver != nil {
			eventCriteria.Topic1 = *c.Approver
		}
		filterer.AddCriteria(eventCriteria)
	}

	return &ExecutorApproverFilterer{filterer: filterer, contract: _Executor.contract}
}

// FilterProposal is a free log retrieval operation binding the contract event 0x7d9bcf5c6cdade398a64a03053a982851ccea20dc827dbc130754b9e78c7c31a.
//
// Solidity: event Proposal(bytes32 indexed proposalID, bytes32 action)
func (_Executor *Executor) FilterProposal(criteria []ExecutorProposalCriteria) *ExecutorProposalFilterer {
	filterer := _Executor.contract.Filter("Proposal")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := contracts.EventCriteria{}
		if c.ProposalID != nil {
			eventCriteria.Topic1 = *c.ProposalID
		}
		filterer.AddCriteria(eventCriteria)
	}

	return &ExecutorProposalFilterer{filterer: filterer, contract: _Executor.contract}
}

// FilterVotingContract is a free log retrieval operation binding the contract event 0xf4cb5443be666f872bc8a75293e99e2204a6573e5eb3d2d485d866f2e13c7ea4.
//
// Solidity: event VotingContract(address indexed contractAddr, bytes32 action)
func (_Executor *Executor) FilterVotingContract(criteria []ExecutorVotingContractCriteria) *ExecutorVotingContractFilterer {
	filterer := _Executor.contract.Filter("VotingContract")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := contracts.EventCriteria{}
		if c.ContractAddr != nil {
			eventCriteria.Topic1 = *c.ContractAddr
		}
		filterer.AddCriteria(eventCriteria)
	}

	return &ExecutorVotingContractFilterer{filterer: filterer, contract: _Executor.contract}
}

// ==================== Event Types and Criteria ====================

// ExecutorApprover represents a Approver event raised by the Executor contract.
type ExecutorApprover struct {
	Approver common.Address
	Action   [32]byte
	Log      *thorest.EventLog
}

type ExecutorApproverCriteria struct {
	Approver *common.Address
}

// ExecutorProposal represents a Proposal event raised by the Executor contract.
type ExecutorProposal struct {
	ProposalID [32]byte
	Action     [32]byte
	Log        *thorest.EventLog
}

type ExecutorProposalCriteria struct {
	ProposalID *[32]byte
}

// ExecutorVotingContract represents a VotingContract event raised by the Executor contract.
type ExecutorVotingContract struct {
	ContractAddr common.Address
	Action       [32]byte
	Log          *thorest.EventLog
}

type ExecutorVotingContractCriteria struct {
	ContractAddr *common.Address
}

// ==================== Call Result Types ====================

// ExecutorApproversResult is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bytes32 identity, bool inPower)
type ExecutorApproversResult struct {
	Identity [32]byte
	InPower  bool
}

// ExecutorProposalsResult is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint64 timeProposed, address proposer, uint8 quorum, uint8 approvalCount, bool executed, address target, bytes data)
type ExecutorProposalsResult struct {
	TimeProposed  uint64
	Proposer      common.Address
	Quorum        uint8
	ApprovalCount uint8
	Executed      bool
	Target        common.Address
	Data          []byte
}

// ==================== Caller Types and Methods ====================

// ExecutorApproverCountCaller provides typed access to the ApproverCount method
type ExecutorApproverCountCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x128e9be6.
func (c *ExecutorApproverCountCaller) WithRevision(rev thorest.Revision) *ExecutorApproverCountCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x128e9be6.
func (c *ExecutorApproverCountCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x128e9be6 and returns the result.
func (c *ExecutorApproverCountCaller) Execute() (uint8, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero uint8
		return zero, err
	}
	if len(data) != 1 {
		var zero uint8
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(uint8); ok {
		return result, nil
	}
	var zero uint8
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// ExecutorApproversCaller provides typed access to the Approvers method
type ExecutorApproversCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x0a144391.
func (c *ExecutorApproversCaller) WithRevision(rev thorest.Revision) *ExecutorApproversCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x0a144391.
func (c *ExecutorApproversCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x0a144391 and returns the result.
func (c *ExecutorApproversCaller) Execute() (*ExecutorApproversResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 2 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(ExecutorApproversResult)
	out.Identity = *abi.ConvertType(data[0], new([32]byte)).(*[32]byte)
	out.InPower = *abi.ConvertType(data[1], new(bool)).(*bool)

	return out, nil
}

// ExecutorProposalsCaller provides typed access to the Proposals method
type ExecutorProposalsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x32ed5b12.
func (c *ExecutorProposalsCaller) WithRevision(rev thorest.Revision) *ExecutorProposalsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x32ed5b12.
func (c *ExecutorProposalsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x32ed5b12 and returns the result.
func (c *ExecutorProposalsCaller) Execute() (*ExecutorProposalsResult, error) {
	data, err := c.caller.Execute()
	if err != nil {
		return nil, err
	}
	if len(data) != 7 {
		return nil, errors.New("invalid number of return values")
	}
	out := new(ExecutorProposalsResult)
	out.TimeProposed = *abi.ConvertType(data[0], new(uint64)).(*uint64)
	out.Proposer = *abi.ConvertType(data[1], new(common.Address)).(*common.Address)
	out.Quorum = *abi.ConvertType(data[2], new(uint8)).(*uint8)
	out.ApprovalCount = *abi.ConvertType(data[3], new(uint8)).(*uint8)
	out.Executed = *abi.ConvertType(data[4], new(bool)).(*bool)
	out.Target = *abi.ConvertType(data[5], new(common.Address)).(*common.Address)
	out.Data = *abi.ConvertType(data[6], new([]byte)).(*[]byte)

	return out, nil
}

// ExecutorVotingContractsCaller provides typed access to the VotingContracts method
type ExecutorVotingContractsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xfa06792b.
func (c *ExecutorVotingContractsCaller) WithRevision(rev thorest.Revision) *ExecutorVotingContractsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xfa06792b.
func (c *ExecutorVotingContractsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xfa06792b and returns the result.
func (c *ExecutorVotingContractsCaller) Execute() (bool, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero bool
		return zero, err
	}
	if len(data) != 1 {
		var zero bool
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(bool); ok {
		return result, nil
	}
	var zero bool
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// ==================== Event Filterer Types and Methods ====================

// ExecutorApproverFilterer provides typed access to filtering Approver events
type ExecutorApproverFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *ExecutorApproverFilterer) Unit(unit string) *ExecutorApproverFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *ExecutorApproverFilterer) Range(from, to int64) *ExecutorApproverFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *ExecutorApproverFilterer) From(from int64) *ExecutorApproverFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *ExecutorApproverFilterer) To(to int64) *ExecutorApproverFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *ExecutorApproverFilterer) Offset(offset int64) *ExecutorApproverFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *ExecutorApproverFilterer) Limit(limit int64) *ExecutorApproverFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *ExecutorApproverFilterer) Order(order string) *ExecutorApproverFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *ExecutorApproverFilterer) Execute() ([]ExecutorApprover, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]ExecutorApprover, len(logs))
	for i, log := range logs {
		event := new(ExecutorApprover)
		if err := f.contract.UnpackLog(event, "Approver", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// ExecutorProposalFilterer provides typed access to filtering Proposal events
type ExecutorProposalFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *ExecutorProposalFilterer) Unit(unit string) *ExecutorProposalFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *ExecutorProposalFilterer) Range(from, to int64) *ExecutorProposalFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *ExecutorProposalFilterer) From(from int64) *ExecutorProposalFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *ExecutorProposalFilterer) To(to int64) *ExecutorProposalFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *ExecutorProposalFilterer) Offset(offset int64) *ExecutorProposalFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *ExecutorProposalFilterer) Limit(limit int64) *ExecutorProposalFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *ExecutorProposalFilterer) Order(order string) *ExecutorProposalFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *ExecutorProposalFilterer) Execute() ([]ExecutorProposal, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]ExecutorProposal, len(logs))
	for i, log := range logs {
		event := new(ExecutorProposal)
		if err := f.contract.UnpackLog(event, "Proposal", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// ExecutorVotingContractFilterer provides typed access to filtering VotingContract events
type ExecutorVotingContractFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *ExecutorVotingContractFilterer) Unit(unit string) *ExecutorVotingContractFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *ExecutorVotingContractFilterer) Range(from, to int64) *ExecutorVotingContractFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *ExecutorVotingContractFilterer) From(from int64) *ExecutorVotingContractFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *ExecutorVotingContractFilterer) To(to int64) *ExecutorVotingContractFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *ExecutorVotingContractFilterer) Offset(offset int64) *ExecutorVotingContractFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *ExecutorVotingContractFilterer) Limit(limit int64) *ExecutorVotingContractFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *ExecutorVotingContractFilterer) Order(order string) *ExecutorVotingContractFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *ExecutorVotingContractFilterer) Execute() ([]ExecutorVotingContract, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]ExecutorVotingContract, len(logs))
	for i, log := range logs {
		event := new(ExecutorVotingContract)
		if err := f.contract.UnpackLog(event, "VotingContract", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}
