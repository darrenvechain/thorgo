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

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Authority *Authority) Executor() *contracts.Caller[common.Address] {
	return contracts.NewCaller[common.Address](_Authority.contract, "executor")
}

// First is a free data retrieval call binding the contract method 0x3df4ddf4.
//
// Solidity: function first() view returns(address)
func (_Authority *Authority) First() *contracts.Caller[common.Address] {
	return contracts.NewCaller[common.Address](_Authority.contract, "first")
}

// AuthorityGetResult is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address _nodeMaster) view returns(bool listed, address endorsor, bytes32 identity, bool active)
type AuthorityGetResult struct {
	Listed   bool
	Endorsor common.Address
	Identity [32]byte
	Active   bool
}

func (_Authority *Authority) Get(_nodeMaster common.Address) *contracts.Caller[*AuthorityGetResult] {
	parser := func(data []interface{}) (*AuthorityGetResult, error) {
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

	return contracts.NewCaller[*AuthorityGetResult](_Authority.contract, "get", _nodeMaster).WithParser(parser)
}

// Next is a free data retrieval call binding the contract method 0xab73e316.
//
// Solidity: function next(address _nodeMaster) view returns(address)
func (_Authority *Authority) Next(_nodeMaster common.Address) *contracts.Caller[common.Address] {
	return contracts.NewCaller[common.Address](_Authority.contract, "next", _nodeMaster)
}

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

// AuthorityCandidate represents a Candidate event raised by the Authority contract.
type AuthorityCandidate struct {
	NodeMaster common.Address
	Action     [32]byte
	Log        *thorest.EventLog
}

type AuthorityCandidateCriteria struct {
	NodeMaster *common.Address `abi:"nodeMaster"`
}

// FilterCandidate is a free log retrieval operation binding the contract event 0xe9e2ad484aeae75ba75479c19d2cbb784b98b2fe4b24dc80a4c8cf142d4c9294.
//
// Solidity: event Candidate(address indexed nodeMaster, bytes32 action)
func (_Authority *Authority) FilterCandidate(criteria []AuthorityCandidateCriteria, filters *thorest.LogFilters) ([]AuthorityCandidate, error) {
	topicHash := _Authority.contract.ABI.Events["Candidate"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Authority.contract.Address,
			Topic0:  &topicHash,
		}
		if c.NodeMaster != nil {
			matcher := *c.NodeMaster
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
			Address: &_Authority.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Authority.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]AuthorityCandidate, len(logs))
	for i, log := range logs {
		event := new(AuthorityCandidate)
		if err := _Authority.contract.UnpackLog(event, "Candidate", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchCandidate listens for on chain events binding the contract event 0xe9e2ad484aeae75ba75479c19d2cbb784b98b2fe4b24dc80a4c8cf142d4c9294.
//
// Solidity: event Candidate(address indexed nodeMaster, bytes32 action)
func (_Authority *Authority) WatchCandidate(criteria []AuthorityCandidateCriteria, ctx context.Context, bufferSize int64) (chan *AuthorityCandidate, error) {
	topicHash := _Authority.contract.ABI.Events["Candidate"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Authority.contract.Address,
			Topic0:  &topicHash,
		}
		if c.NodeMaster != nil {
			matcher := *c.NodeMaster
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *AuthorityCandidate, bufferSize)
	blocks := blocks.New(ctx, _Authority.thor)
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
						ev := new(AuthorityCandidate)
						if err := _Authority.contract.UnpackLog(ev, "Candidate", log); err != nil {
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
