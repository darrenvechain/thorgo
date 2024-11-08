// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package builtins

import (
	"context"
	"errors"
	"math/big"
	"strings"

	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/accounts"
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
)

// ParamsMetaData contains all meta data concerning the Params contract.
var ParamsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Set\",\"type\":\"event\"}]",
}

// Params is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type Params struct {
	thor     *thorgo.Thor       // Thor connection to use
	contract *accounts.Contract // Generic contract wrapper for the low level calls
}

// ParamsTransactor is an auto generated Go binding around an Ethereum, allowing you to transact with the contract.
type ParamsTransactor struct {
	*Params
	manager accounts.TxManager // TxManager to use
}

// NewParams creates a new instance of Params, bound to a specific deployed contract.
func NewParams(thor *thorgo.Thor) (*Params, error) {
	parsed, err := ParamsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := thor.Account(common.HexToAddress("0x0000000000000000000000000000506172616d73")).Contract(parsed)
	if err != nil {
		return nil, err
	}
	return &Params{thor: thor, contract: contract}, nil
}

// NewParamsTransactor creates a new instance of ParamsTransactor, bound to a specific deployed contract.
func NewParamsTransactor(thor *thorgo.Thor, manager accounts.TxManager) (*ParamsTransactor, error) {
	base, err := NewParams(thor)
	if err != nil {
		return nil, err
	}
	return &ParamsTransactor{Params: base, manager: manager}, nil
}

// Address returns the address of the contract.
func (_Params *Params) Address() common.Address {
	return _Params.contract.Address
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *Params) Call(revision thorest.Revision, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.contract.Call(method, result, params...)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParamsTransactor *ParamsTransactor) Transact(vetValue *big.Int, method string, params ...interface{}) (*transactions.Visitor, error) {
	return _ParamsTransactor.contract.SendWithVET(_ParamsTransactor.manager, vetValue, method, params...)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Params *Params) Executor(revision ...thorest.Revision) (common.Address, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Params.Call(rev, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _key) view returns(uint256)
func (_Params *Params) Get(_key [32]byte, revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _Params.Call(rev, &out, "get", _key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Set is a paid mutator transaction binding the contract method 0x273f4940.
//
// Solidity: function set(bytes32 _key, uint256 _value) returns()
func (_ParamsTransactor *ParamsTransactor) Set(_key [32]byte, _value *big.Int, vetValue ...*big.Int) (*transactions.Visitor, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _ParamsTransactor.Transact(val, "set", _key, _value)
}

// SetAsClause is a transaction clause generator 0x273f4940.
//
// Solidity: function set(bytes32 _key, uint256 _value) returns()
func (_Params *Params) SetAsClause(_key [32]byte, _value *big.Int, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _Params.contract.AsClauseWithVET(val, "set", _key, _value)
}

// ParamsSet represents a Set event raised by the Params contract.
type ParamsSet struct {
	Key   [32]byte
	Value *big.Int
	Log   thorest.EventLog
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

	logs, err := _Params.thor.Client.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	inputs := _Params.contract.ABI.Events["Set"].Inputs
	var indexed abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
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
func (_Params *Params) WatchSet(criteria []ParamsSetCriteria, ctx context.Context, bufferSize int) (chan *ParamsSet, error) {
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
	blockSub := _Params.thor.Blocks.Subscribe(ctx, bufferSize)

	go func() {
		defer close(eventChan)

		for {
			select {
			case block := <-blockSub:
				for _, tx := range block.Transactions {
					for index, outputs := range tx.Outputs {
						for _, event := range outputs.Events {
							if event.Address != _Params.contract.Address {
								continue
							}
							if topicHash != event.Topics[0] {
								continue
							}
							for _, c := range criteriaSet {
								if c.Topic1 != nil && *c.Topic1 != event.Topics[1] {
									continue
								}
								if c.Topic2 != nil && *c.Topic2 != event.Topics[2] {
									continue
								}
								if c.Topic3 != nil && *c.Topic3 != event.Topics[3] {
									continue
								}
								if c.Topic4 != nil && *c.Topic4 != event.Topics[4] {
									continue
								}
							}

							log := thorest.EventLog{
								Address: &_Params.contract.Address,
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

							ev := new(ParamsSet)
							if err := _Params.contract.UnpackLog(ev, "Set", log); err != nil {
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
