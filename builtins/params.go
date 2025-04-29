// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package builtins

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"time"

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
	_ = errors.Is
	_ = big.NewInt
	_ = strings.ReplaceAll
	_ = abi.ConvertType
	_ = hexutil.Decode
	_ = context.Background
	_ = tx.NewClause
	_ = blocks.New
	_ = time.Sleep
)

// ParamsMetaData contains all meta data concerning the Params contract.
var ParamsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Set\",\"type\":\"event\"}]",
}

// Params is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Params struct {
	thor     *thorest.Client    // Thor client connection to use
	contract *accounts.Contract // Generic contract wrapper for the low level calls
}

// ParamsTransactor is an auto generated Go binding around an Ethereum, allowing you to transact with the contract.
type ParamsTransactor struct {
	*Params
	contract *accounts.ContractTransactor // Generic contract wrapper for the low level calls
	manager  accounts.TxManager           // TxManager to use
}

// NewParams creates a new instance of Params, bound to a specific deployed contract.
func NewParams(thor *thorest.Client) (*Params, error) {
	parsed, err := ParamsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := accounts.New(thor, common.HexToAddress("0x0000000000000000000000000000506172616d73")).Contract(parsed)
	return &Params{thor: thor, contract: contract}, nil
}

// NewParamsTransactor creates a new instance of ParamsTransactor, bound to a specific deployed contract.
func NewParamsTransactor(thor *thorest.Client, manager accounts.TxManager) (*ParamsTransactor, error) {
	base, err := NewParams(thor)
	if err != nil {
		return nil, err
	}
	return &ParamsTransactor{Params: base, contract: base.contract.Transactor(manager), manager: manager}, nil
}

// Address returns the address of the contract.
func (_Params *Params) Address() common.Address {
	return _Params.contract.Address
}

// Transactor constructs a new transactor for the contract, which allows to send transactions.
func (_Params *Params) Transactor(manager accounts.TxManager) *ParamsTransactor {
	return &ParamsTransactor{Params: _Params, contract: _Params.contract.Transactor(manager), manager: manager}
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *Params) Call(revision thorest.Revision, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.contract.CallAt(revision, method, result, params...)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParamsTransactor *ParamsTransactor) Transact(opts *transactions.Options, vet *big.Int, method string, params ...interface{}) *accounts.Sender {
	return _ParamsTransactor.contract.SendPayable(opts, vet, method, params...)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Params *Params) Executor(revision thorest.Revision) (common.Address, error) {
	var out []interface{}
	err := _Params.Call(revision, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _key) view returns(uint256)
func (_Params *Params) Get(_key [32]byte, revision thorest.Revision) (*big.Int, error) {
	var out []interface{}
	err := _Params.Call(revision, &out, "get", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Set is a paid mutator transaction binding the contract method 0x273f4940.
//
// Solidity: function set(bytes32 _key, uint256 _value) returns()
func (_ParamsTransactor *ParamsTransactor) Set(_key [32]byte, _value *big.Int, opts *transactions.Options) *accounts.Sender {
	return _ParamsTransactor.Transact(opts, big.NewInt(0), "set", _key, _value)
}

// SetAsClause is a transaction clause generator 0x273f4940.
//
// Solidity: function set(bytes32 _key, uint256 _value) returns()
func (_Params *Params) SetAsClause(_key [32]byte, _value *big.Int) (*tx.Clause, error) {
	return _Params.contract.AsClause("set", _key, _value)
}

// ParamsSet represents a Set event raised by the Params contract.
type ParamsSet struct {
	Key   [32]byte
	Value *big.Int
	Log   *thorest.EventLog
}

type ParamsSetCriteria struct {
	Key *[32]byte `abi:"key"`
}

// FilterSet is a free log retrieval operation binding the contract event 0x28e3246f80515f5c1ed987b133ef2f193439b25acba6a5e69f219e896fc9d179.
//
// Solidity: event Set(bytes32 indexed key, uint256 value)
func (_Params *Params) FilterSet(criteria []ParamsSetCriteria, filters *thorest.LogFilters) ([]ParamsSet, error) {
	topicHash := _Params.contract.ABI.Events["Set"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Params.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Key != nil {
			matcher := *c.Key
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
			Address: &_Params.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _Params.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]ParamsSet, len(logs))
	for i, log := range logs {
		event := new(ParamsSet)
		if err := _Params.contract.UnpackLog(event, "Set", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchSet listens for on chain events binding the contract event 0x28e3246f80515f5c1ed987b133ef2f193439b25acba6a5e69f219e896fc9d179.
//
// Solidity: event Set(bytes32 indexed key, uint256 value)
func (_Params *Params) WatchSet(criteria []ParamsSetCriteria, ctx context.Context, bufferSize int64) (chan *ParamsSet, error) {
	topicHash := _Params.contract.ABI.Events["Set"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_Params.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Key != nil {
			matcher := *c.Key
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *ParamsSet, bufferSize)
	blocks := blocks.New(ctx, _Params.thor)
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
						ev := new(ParamsSet)
						if err := _Params.contract.UnpackLog(ev, "Set", log); err != nil {
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
