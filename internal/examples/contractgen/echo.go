// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractgen

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

// EchoContractMetaData contains all meta data concerning the EchoContract contract.
var EchoContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"Echoed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"}],\"name\":\"echo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608080604052346100165761018d908161001c8239f35b600080fdfe60806040818152600436101561001457600080fd5b6000803560e01c63f15da7291461002a57600080fd5b3461015457602092836003193601126101505767ffffffffffffffff9160043583811161015057366023820112156101505780600401359084821161014c576024810190602483369201011161014c57868452818785015281818786013782868386010152601f19947f984ec84768d7671494d262d4e7fa6dab3eed5e37433c68c0c5d476b57679312686601f850116958881888101030190a1865194603f01861685019081118582101761013857918088928596979594895281885283880137850101528351948592818452845191828186015281955b8387106101205750508394508582601f949501015201168101030190f35b86810182015189880189015295810195889550610102565b634e487b7160e01b84526041600452602484fd5b8280fd5b5080fd5b80fdfea264697066735822122051657eb27142f8a682469c9903f405e05543d43d71b65608cb7482515444123764736f6c63430008140033",
}

// DeployEchoContract deploys a new Ethereum contract, binding an instance of EchoContract to it.
func DeployEchoContract(ctx context.Context, thor *thorest.Client, sender contracts.TxManager, opts *transactions.Options) (common.Hash, *EchoContract, error) {
	parsed, err := EchoContractMetaData.GetAbi()
	if err != nil {
		return common.Hash{}, nil, err
	}

	bytes, err := hexutil.Decode(EchoContractMetaData.Bin)
	if err != nil {
		return common.Hash{}, nil, err
	}
	contract, txID, err := contracts.NewDeployer(thor, bytes, parsed).Deploy(ctx, sender, opts)
	if err != nil {
		return common.Hash{}, nil, err
	}
	return txID, &EchoContract{thor: thor, contract: contract}, nil
}

// EchoContract is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type EchoContract struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewEchoContract creates a new instance of EchoContract, bound to a specific deployed contract.
func NewEchoContract(address common.Address, thor *thorest.Client) (*EchoContract, error) {
	parsed, err := EchoContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, address, parsed)
	return &EchoContract{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_EchoContract *EchoContract) Address() common.Address {
	return _EchoContract.contract.Address
}

// Raw returns the underlying contract.
func (_EchoContract *EchoContract) Raw() *contracts.Contract {
	return _EchoContract.contract
}

// ==================== View Functions ====================

// ==================== Transaction Functions ====================

// Echo is a paid mutator transaction binding the contract method 0xf15da729.
//
// Solidity: function echo(string message) returns(string)
func (_EchoContract *EchoContract) Echo(message string) *contracts.Sender {
	return contracts.NewSender(_EchoContract.contract, "echo", message)
}

// ==================== Event Functions ====================

// UnpackEchoedLogs unpacks existing logs into typed Echoed events.
func (_EchoContract *EchoContract) UnpackEchoedLogs(logs []*thorest.EventLog) ([]EchoContractEchoed, error) {
	events := make([]EchoContractEchoed, len(logs))
	for i, log := range logs {
		event := EchoContractEchoed{}
		if err := _EchoContract.contract.UnpackLog(&event, "Echoed", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}
	return events, nil
}

// FilterEchoed is a free log retrieval operation binding the contract event 0x984ec84768d7671494d262d4e7fa6dab3eed5e37433c68c0c5d476b576793126.
//
// Solidity: event Echoed(string message)
func (_EchoContract *EchoContract) FilterEchoed() *EchoContractEchoedFilterer {
	filterer := _EchoContract.contract.Filter("Echoed")

	return &EchoContractEchoedFilterer{filterer: filterer, contract: _EchoContract.contract}
}

// ==================== Event IDs ====================

// EchoContractEchoedEventID is the event ID for Echoed
// Solidity: event Echoed(string message)
var EchoContractEchoedEventID = common.HexToHash("0x984ec84768d7671494d262d4e7fa6dab3eed5e37433c68c0c5d476b576793126")

// ==================== Event Types and Criteria ====================

// EchoContractEchoed represents a Echoed event raised by the EchoContract contract.
type EchoContractEchoed struct {
	Message string
	Log     *thorest.EventLog
}

// ==================== Call Result Types ====================

// ==================== Caller Types and Methods ====================

// ==================== Event Filterer Types and Methods ====================

// EchoContractEchoedFilterer provides typed access to filtering Echoed events
type EchoContractEchoedFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *EchoContractEchoedFilterer) Unit(unit string) *EchoContractEchoedFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *EchoContractEchoedFilterer) IncludeIndexes(include bool) *EchoContractEchoedFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *EchoContractEchoedFilterer) Range(from, to int64) *EchoContractEchoedFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *EchoContractEchoedFilterer) From(from int64) *EchoContractEchoedFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *EchoContractEchoedFilterer) To(to int64) *EchoContractEchoedFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *EchoContractEchoedFilterer) Offset(offset int64) *EchoContractEchoedFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *EchoContractEchoedFilterer) Limit(limit int64) *EchoContractEchoedFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *EchoContractEchoedFilterer) Order(order string) *EchoContractEchoedFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *EchoContractEchoedFilterer) Execute() ([]EchoContractEchoed, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}
	return (&EchoContract{contract: f.contract}).UnpackEchoedLogs(logs)
}
