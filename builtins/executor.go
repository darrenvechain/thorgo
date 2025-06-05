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

// ApproverCount is a free data retrieval call binding the contract method 0x128e9be6.
//
// Solidity: function approverCount() view returns(uint8)
func (_Executor *Executor) ApproverCount() *contracts.Caller[uint8] {
	return contracts.NewCaller[uint8](_Executor.contract, "approverCount")
}

// ExecutorApproversResult is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bytes32 identity, bool inPower)
type ExecutorApproversResult struct {
	Identity [32]byte
	InPower  bool
}

func (_Executor *Executor) Approvers(arg0 common.Address) *contracts.Caller[*ExecutorApproversResult] {
	parser := func(data []interface{}) (*ExecutorApproversResult, error) {
		if len(data) != 2 {
			return nil, errors.New("invalid number of return values")
		}
		out := new(ExecutorApproversResult)

		out.Identity = *abi.ConvertType(data[0], new([32]byte)).(*[32]byte)
		out.InPower = *abi.ConvertType(data[1], new(bool)).(*bool)

		return out, nil
	}

	return contracts.NewCaller[*ExecutorApproversResult](_Executor.contract, "approvers", arg0).WithParser(parser)
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

func (_Executor *Executor) Proposals(arg0 [32]byte) *contracts.Caller[*ExecutorProposalsResult] {
	parser := func(data []interface{}) (*ExecutorProposalsResult, error) {
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

	return contracts.NewCaller[*ExecutorProposalsResult](_Executor.contract, "proposals", arg0).WithParser(parser)
}

// VotingContracts is a free data retrieval call binding the contract method 0xfa06792b.
//
// Solidity: function votingContracts(address ) view returns(bool)
func (_Executor *Executor) VotingContracts(arg0 common.Address) *contracts.Caller[bool] {
	return contracts.NewCaller[bool](_Executor.contract, "votingContracts", arg0)
}

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

// ExecutorApprover represents a Approver event raised by the Executor contract.
type ExecutorApprover struct {
	Approver common.Address
	Action   [32]byte
	Log      *thorest.EventLog
}

type ExecutorApproverCriteria struct {
	Approver *common.Address `abi:"approver"`
}

// FilterApprover is a free log retrieval operation binding the contract event 0x770115cde75e60f17b265d7e0c5e39c57abf243bc316c7e5c2f8d851771da6ac.
//
// Solidity: event Approver(address indexed approver, bytes32 action)
func (_Executor *Executor) FilterApprover(criteria []ExecutorApproverCriteria, filters *thorest.LogFilters) ([]ExecutorApprover, error) {
	topicHash := _Executor.contract.ABI.Events["Approver"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Approver != nil {
			matcher := *c.Approver
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Executor.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]ExecutorApprover, len(logs))
	for i, log := range logs {
		event := new(ExecutorApprover)
		if err := _Executor.contract.UnpackLog(event, "Approver", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchApprover listens for on chain events binding the contract event 0x770115cde75e60f17b265d7e0c5e39c57abf243bc316c7e5c2f8d851771da6ac.
//
// Solidity: event Approver(address indexed approver, bytes32 action)
func (_Executor *Executor) WatchApprover(criteria []ExecutorApproverCriteria, ctx context.Context, bufferSize int64) (chan *ExecutorApprover, error) {
	topicHash := _Executor.contract.ABI.Events["Approver"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Approver != nil {
			matcher := *c.Approver
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *ExecutorApprover, bufferSize)
	blocks := blocks.New(ctx, _Executor.thor)
	ticker := blocks.Ticker()
	best, err := blocks.Best()
	if err != nil {
		return nil, err
	}

	go func(current int64) {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				for { // loop until the current block is not found
					block, err := blocks.Expanded(thorest.RevisionNumber(current))
					if errors.Is(thorest.ErrNotFound, err) {
						break
					}
					if err != nil {
						time.Sleep(250 * time.Millisecond)
						continue
					}
					current++

					for _, log := range block.FilteredEvents(criteriaSet) {
						ev := new(ExecutorApprover)
						if err := _Executor.contract.UnpackLog(ev, "Approver", log); err != nil {
							continue
						}
						ev.Log = log
						eventChan <- ev
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}(best.Number + 1)

	return eventChan, nil
}

// ExecutorProposal represents a Proposal event raised by the Executor contract.
type ExecutorProposal struct {
	ProposalID [32]byte
	Action     [32]byte
	Log        *thorest.EventLog
}

type ExecutorProposalCriteria struct {
	ProposalID *[32]byte `abi:"proposalID"`
}

// FilterProposal is a free log retrieval operation binding the contract event 0x7d9bcf5c6cdade398a64a03053a982851ccea20dc827dbc130754b9e78c7c31a.
//
// Solidity: event Proposal(bytes32 indexed proposalID, bytes32 action)
func (_Executor *Executor) FilterProposal(criteria []ExecutorProposalCriteria, filters *thorest.LogFilters) ([]ExecutorProposal, error) {
	topicHash := _Executor.contract.ABI.Events["Proposal"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		}
		if c.ProposalID != nil {
			matcher := *c.ProposalID
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Executor.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]ExecutorProposal, len(logs))
	for i, log := range logs {
		event := new(ExecutorProposal)
		if err := _Executor.contract.UnpackLog(event, "Proposal", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchProposal listens for on chain events binding the contract event 0x7d9bcf5c6cdade398a64a03053a982851ccea20dc827dbc130754b9e78c7c31a.
//
// Solidity: event Proposal(bytes32 indexed proposalID, bytes32 action)
func (_Executor *Executor) WatchProposal(criteria []ExecutorProposalCriteria, ctx context.Context, bufferSize int64) (chan *ExecutorProposal, error) {
	topicHash := _Executor.contract.ABI.Events["Proposal"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		}
		if c.ProposalID != nil {
			matcher := *c.ProposalID
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *ExecutorProposal, bufferSize)
	blocks := blocks.New(ctx, _Executor.thor)
	ticker := blocks.Ticker()
	best, err := blocks.Best()
	if err != nil {
		return nil, err
	}

	go func(current int64) {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				for { // loop until the current block is not found
					block, err := blocks.Expanded(thorest.RevisionNumber(current))
					if errors.Is(thorest.ErrNotFound, err) {
						break
					}
					if err != nil {
						time.Sleep(250 * time.Millisecond)
						continue
					}
					current++

					for _, log := range block.FilteredEvents(criteriaSet) {
						ev := new(ExecutorProposal)
						if err := _Executor.contract.UnpackLog(ev, "Proposal", log); err != nil {
							continue
						}
						ev.Log = log
						eventChan <- ev
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}(best.Number + 1)

	return eventChan, nil
}

// ExecutorVotingContract represents a VotingContract event raised by the Executor contract.
type ExecutorVotingContract struct {
	ContractAddr common.Address
	Action       [32]byte
	Log          *thorest.EventLog
}

type ExecutorVotingContractCriteria struct {
	ContractAddr *common.Address `abi:"contractAddr"`
}

// FilterVotingContract is a free log retrieval operation binding the contract event 0xf4cb5443be666f872bc8a75293e99e2204a6573e5eb3d2d485d866f2e13c7ea4.
//
// Solidity: event VotingContract(address indexed contractAddr, bytes32 action)
func (_Executor *Executor) FilterVotingContract(criteria []ExecutorVotingContractCriteria, filters *thorest.LogFilters) ([]ExecutorVotingContract, error) {
	topicHash := _Executor.contract.ABI.Events["VotingContract"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		}
		if c.ContractAddr != nil {
			matcher := *c.ContractAddr
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Executor.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]ExecutorVotingContract, len(logs))
	for i, log := range logs {
		event := new(ExecutorVotingContract)
		if err := _Executor.contract.UnpackLog(event, "VotingContract", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchVotingContract listens for on chain events binding the contract event 0xf4cb5443be666f872bc8a75293e99e2204a6573e5eb3d2d485d866f2e13c7ea4.
//
// Solidity: event VotingContract(address indexed contractAddr, bytes32 action)
func (_Executor *Executor) WatchVotingContract(criteria []ExecutorVotingContractCriteria, ctx context.Context, bufferSize int64) (chan *ExecutorVotingContract, error) {
	topicHash := _Executor.contract.ABI.Events["VotingContract"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Executor.contract.Address,
			Topic0:  &topicHash,
		}
		if c.ContractAddr != nil {
			matcher := *c.ContractAddr
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *ExecutorVotingContract, bufferSize)
	blocks := blocks.New(ctx, _Executor.thor)
	ticker := blocks.Ticker()
	best, err := blocks.Best()
	if err != nil {
		return nil, err
	}

	go func(current int64) {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				for { // loop until the current block is not found
					block, err := blocks.Expanded(thorest.RevisionNumber(current))
					if errors.Is(thorest.ErrNotFound, err) {
						break
					}
					if err != nil {
						time.Sleep(250 * time.Millisecond)
						continue
					}
					current++

					for _, log := range block.FilteredEvents(criteriaSet) {
						ev := new(ExecutorVotingContract)
						if err := _Executor.contract.UnpackLog(ev, "VotingContract", log); err != nil {
							continue
						}
						ev.Log = log
						eventChan <- ev
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}(best.Number + 1)

	return eventChan, nil
}
