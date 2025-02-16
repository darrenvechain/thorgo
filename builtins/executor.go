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

// ExecutorMetaData contains all meta data concerning the Executor contract.
var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvers\",\"outputs\":[{\"name\":\"identity\",\"type\":\"bytes32\"},{\"name\":\"inPower\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"approverCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approver\",\"type\":\"address\"}],\"name\":\"revokeApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"name\":\"timeProposed\",\"type\":\"uint64\"},{\"name\":\"proposer\",\"type\":\"address\"},{\"name\":\"quorum\",\"type\":\"uint8\"},{\"name\":\"approvalCount\",\"type\":\"uint8\"},{\"name\":\"executed\",\"type\":\"bool\"},{\"name\":\"target\",\"type\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_approver\",\"type\":\"address\"},{\"name\":\"_identity\",\"type\":\"bytes32\"}],\"name\":\"addApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_target\",\"type\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"propose\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"attachVotingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"bytes32\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"detachVotingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalID\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"votingContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"proposalID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"Proposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"Approver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"contractAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"action\",\"type\":\"bytes32\"}],\"name\":\"VotingContract\",\"type\":\"event\"}]",
}

// Executor is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Executor struct {
	thor     *thorest.Client    // Thor client connection to use
	contract *accounts.Contract // Generic contract wrapper for the low level calls
}

// ExecutorTransactor is an auto generated Go binding around an Ethereum, allowing you to transact with the contract.
type ExecutorTransactor struct {
	*Executor
	contract *accounts.ContractTransactor // Generic contract wrapper for the low level calls
	manager  accounts.TxManager           // TxManager to use
}

// NewExecutor creates a new instance of Executor, bound to a specific deployed contract.
func NewExecutor(thor *thorest.Client) (*Executor, error) {
	parsed, err := ExecutorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := accounts.New(thor, common.HexToAddress("0x0000000000000000000000004578656375746f72")).Contract(parsed)
	return &Executor{thor: thor, contract: contract}, nil
}

// NewExecutorTransactor creates a new instance of ExecutorTransactor, bound to a specific deployed contract.
func NewExecutorTransactor(thor *thorest.Client, manager accounts.TxManager) (*ExecutorTransactor, error) {
	base, err := NewExecutor(thor)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransactor{Executor: base, contract: base.contract.Transactor(manager), manager: manager}, nil
}

// Address returns the address of the contract.
func (_Executor *Executor) Address() common.Address {
	return _Executor.contract.Address
}

// Transactor constructs a new transactor for the contract, which allows to send transactions.
func (_Executor *Executor) Transactor(manager accounts.TxManager) *ExecutorTransactor {
	return &ExecutorTransactor{Executor: _Executor, contract: _Executor.contract.Transactor(manager), manager: manager}
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *Executor) Call(revision thorest.Revision, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.contract.CallAt(revision, method, result, params...)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutorTransactor *ExecutorTransactor) Transact(opts *transactions.Options, method string, params ...interface{}) (*transactions.Visitor, error) {
	return _ExecutorTransactor.contract.Send(opts, method, params...)
}

// ApproverCount is a free data retrieval call binding the contract method 0x128e9be6.
//
// Solidity: function approverCount() view returns(uint8)
func (_Executor *Executor) ApproverCount(revision ...thorest.Revision) (uint8, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Executor.Call(rev, &out, "approverCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bytes32 identity, bool inPower)
func (_Executor *Executor) Approvers(arg0 common.Address, revision ...thorest.Revision) (struct {
	Identity [32]byte
	InPower  bool
}, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Executor.Call(rev, &out, "approvers", arg0)

	outstruct := new(struct {
		Identity [32]byte
		InPower  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Identity = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.InPower = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(uint64 timeProposed, address proposer, uint8 quorum, uint8 approvalCount, bool executed, address target, bytes data)
func (_Executor *Executor) Proposals(arg0 [32]byte, revision ...thorest.Revision) (struct {
	TimeProposed  uint64
	Proposer      common.Address
	Quorum        uint8
	ApprovalCount uint8
	Executed      bool
	Target        common.Address
	Data          []byte
}, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Executor.Call(rev, &out, "proposals", arg0)

	outstruct := new(struct {
		TimeProposed  uint64
		Proposer      common.Address
		Quorum        uint8
		ApprovalCount uint8
		Executed      bool
		Target        common.Address
		Data          []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TimeProposed = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Proposer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Quorum = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ApprovalCount = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Executed = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Target = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Data = *abi.ConvertType(out[6], new([]byte)).(*[]byte)

	return *outstruct, err

}

// VotingContracts is a free data retrieval call binding the contract method 0xfa06792b.
//
// Solidity: function votingContracts(address ) view returns(bool)
func (_Executor *Executor) VotingContracts(arg0 common.Address, revision ...thorest.Revision) (bool, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Executor.Call(rev, &out, "votingContracts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// AddApprover is a paid mutator transaction binding the contract method 0x3ef0c09e.
//
// Solidity: function addApprover(address _approver, bytes32 _identity) returns()
func (_ExecutorTransactor *ExecutorTransactor) AddApprover(_approver common.Address, _identity [32]byte, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ExecutorTransactor.Transact(opts, "addApprover", _approver, _identity)
}

// AddApproverAsClause is a transaction clause generator 0x3ef0c09e.
//
// Solidity: function addApprover(address _approver, bytes32 _identity) returns()
func (_Executor *Executor) AddApproverAsClause(_approver common.Address, _identity [32]byte, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Executor.contract.AsClauseWithVET(val, "addApprover", _approver, _identity)
}

// Approve is a paid mutator transaction binding the contract method 0xa53a1adf.
//
// Solidity: function approve(bytes32 _proposalID) returns()
func (_ExecutorTransactor *ExecutorTransactor) Approve(_proposalID [32]byte, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ExecutorTransactor.Transact(opts, "approve", _proposalID)
}

// ApproveAsClause is a transaction clause generator 0xa53a1adf.
//
// Solidity: function approve(bytes32 _proposalID) returns()
func (_Executor *Executor) ApproveAsClause(_proposalID [32]byte, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Executor.contract.AsClauseWithVET(val, "approve", _proposalID)
}

// AttachVotingContract is a paid mutator transaction binding the contract method 0xa1fb668f.
//
// Solidity: function attachVotingContract(address _contract) returns()
func (_ExecutorTransactor *ExecutorTransactor) AttachVotingContract(_contract common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ExecutorTransactor.Transact(opts, "attachVotingContract", _contract)
}

// AttachVotingContractAsClause is a transaction clause generator 0xa1fb668f.
//
// Solidity: function attachVotingContract(address _contract) returns()
func (_Executor *Executor) AttachVotingContractAsClause(_contract common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Executor.contract.AsClauseWithVET(val, "attachVotingContract", _contract)
}

// DetachVotingContract is a paid mutator transaction binding the contract method 0xa83b3bd8.
//
// Solidity: function detachVotingContract(address _contract) returns()
func (_ExecutorTransactor *ExecutorTransactor) DetachVotingContract(_contract common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ExecutorTransactor.Transact(opts, "detachVotingContract", _contract)
}

// DetachVotingContractAsClause is a transaction clause generator 0xa83b3bd8.
//
// Solidity: function detachVotingContract(address _contract) returns()
func (_Executor *Executor) DetachVotingContractAsClause(_contract common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Executor.contract.AsClauseWithVET(val, "detachVotingContract", _contract)
}

// Execute is a paid mutator transaction binding the contract method 0xe751f271.
//
// Solidity: function execute(bytes32 _proposalID) returns()
func (_ExecutorTransactor *ExecutorTransactor) Execute(_proposalID [32]byte, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ExecutorTransactor.Transact(opts, "execute", _proposalID)
}

// ExecuteAsClause is a transaction clause generator 0xe751f271.
//
// Solidity: function execute(bytes32 _proposalID) returns()
func (_Executor *Executor) ExecuteAsClause(_proposalID [32]byte, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Executor.contract.AsClauseWithVET(val, "execute", _proposalID)
}

// Propose is a paid mutator transaction binding the contract method 0x9d481848.
//
// Solidity: function propose(address _target, bytes _data) returns(bytes32)
func (_ExecutorTransactor *ExecutorTransactor) Propose(_target common.Address, _data []byte, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ExecutorTransactor.Transact(opts, "propose", _target, _data)
}

// ProposeAsClause is a transaction clause generator 0x9d481848.
//
// Solidity: function propose(address _target, bytes _data) returns(bytes32)
func (_Executor *Executor) ProposeAsClause(_target common.Address, _data []byte, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Executor.contract.AsClauseWithVET(val, "propose", _target, _data)
}

// RevokeApprover is a paid mutator transaction binding the contract method 0x18d13ef7.
//
// Solidity: function revokeApprover(address _approver) returns()
func (_ExecutorTransactor *ExecutorTransactor) RevokeApprover(_approver common.Address, opts *transactions.Options) (*transactions.Visitor, error) {
	return _ExecutorTransactor.Transact(opts, "revokeApprover", _approver)
}

// RevokeApproverAsClause is a transaction clause generator 0x18d13ef7.
//
// Solidity: function revokeApprover(address _approver) returns()
func (_Executor *Executor) RevokeApproverAsClause(_approver common.Address, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Executor.contract.AsClauseWithVET(val, "revokeApprover", _approver)
}

// ExecutorApprover represents a Approver event raised by the Executor contract.
type ExecutorApprover struct {
	Approver common.Address
	Action   [32]byte
	Log      thorest.EventLog
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

	inputs := _Executor.contract.ABI.Events["Approver"].Inputs
	var indexed abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
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

	go func() {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				block, err := blocks.Best()
				if err != nil {
					continue
				}
				for _, tx := range block.Transactions {
					for index, outputs := range tx.Outputs {
						for _, event := range outputs.Events {
							matches := false
							for _, c := range criteriaSet {
								if c.Matches(event) {
									matches = true
									break
								}
							}
							if !matches {
								continue
							}

							log := thorest.EventLog{
								Address: &_Executor.contract.Address,
								Topics:  event.Topics,
								Data:    event.Data,
								Meta: thorest.LogMeta{
									BlockID:     block.ID,
									BlockNumber: block.Number,
									BlockTime:   block.Timestamp,
									TxID:        tx.ID,
									TxOrigin:    tx.Origin,
									ClauseIndex: int64(index),
								},
							}

							ev := new(ExecutorApprover)
							if err := _Executor.contract.UnpackLog(ev, "Approver", log); err != nil {
								continue
							}
							ev.Log = log
							eventChan <- ev
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventChan, nil
}

// ExecutorProposal represents a Proposal event raised by the Executor contract.
type ExecutorProposal struct {
	ProposalID [32]byte
	Action     [32]byte
	Log        thorest.EventLog
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

	inputs := _Executor.contract.ABI.Events["Proposal"].Inputs
	var indexed abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
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

	go func() {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				block, err := blocks.Best()
				if err != nil {
					continue
				}
				for _, tx := range block.Transactions {
					for index, outputs := range tx.Outputs {
						for _, event := range outputs.Events {
							matches := false
							for _, c := range criteriaSet {
								if c.Matches(event) {
									matches = true
									break
								}
							}
							if !matches {
								continue
							}

							log := thorest.EventLog{
								Address: &_Executor.contract.Address,
								Topics:  event.Topics,
								Data:    event.Data,
								Meta: thorest.LogMeta{
									BlockID:     block.ID,
									BlockNumber: block.Number,
									BlockTime:   block.Timestamp,
									TxID:        tx.ID,
									TxOrigin:    tx.Origin,
									ClauseIndex: int64(index),
								},
							}

							ev := new(ExecutorProposal)
							if err := _Executor.contract.UnpackLog(ev, "Proposal", log); err != nil {
								continue
							}
							ev.Log = log
							eventChan <- ev
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventChan, nil
}

// ExecutorVotingContract represents a VotingContract event raised by the Executor contract.
type ExecutorVotingContract struct {
	ContractAddr common.Address
	Action       [32]byte
	Log          thorest.EventLog
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

	inputs := _Executor.contract.ABI.Events["VotingContract"].Inputs
	var indexed abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
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

	go func() {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				block, err := blocks.Best()
				if err != nil {
					continue
				}
				for _, tx := range block.Transactions {
					for index, outputs := range tx.Outputs {
						for _, event := range outputs.Events {
							matches := false
							for _, c := range criteriaSet {
								if c.Matches(event) {
									matches = true
									break
								}
							}
							if !matches {
								continue
							}

							log := thorest.EventLog{
								Address: &_Executor.contract.Address,
								Topics:  event.Topics,
								Data:    event.Data,
								Meta: thorest.LogMeta{
									BlockID:     block.ID,
									BlockNumber: block.Number,
									BlockTime:   block.Timestamp,
									TxID:        tx.ID,
									TxOrigin:    tx.Origin,
									ClauseIndex: int64(index),
								},
							}

							ev := new(ExecutorVotingContract)
							if err := _Executor.contract.UnpackLog(ev, "VotingContract", log); err != nil {
								continue
							}
							ev.Log = log
							eventChan <- ev
						}
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return eventChan, nil
}
