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

// ParamsExecutorCaller provides typed access to the Executor method
type ParamsExecutorCaller struct {
	caller *contracts.Caller
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Params *Params) Executor() *ParamsExecutorCaller {
	return &ParamsExecutorCaller{
		caller: _Params.contract.Call("executor"),
	}
}

func (c *ParamsExecutorCaller) WithRevision(rev thorest.Revision) *ParamsExecutorCaller {
	c.caller.WithRevision(rev)
	return c
}

func (c *ParamsExecutorCaller) WithValue(value *big.Int) *ParamsExecutorCaller {
	c.caller.WithValue(value)
	return c
}

func (c *ParamsExecutorCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

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

// Get is a free data retrieval call binding the contract method 0x8eaa6ac0.
//
// Solidity: function get(bytes32 _key) view returns(uint256)
func (_Params *Params) Get(_key [32]byte) *ParamsGetCaller {
	return &ParamsGetCaller{
		caller: _Params.contract.Call("get", _key),
	}
}

func (c *ParamsGetCaller) WithRevision(rev thorest.Revision) *ParamsGetCaller {
	c.caller.WithRevision(rev)
	return c
}

func (c *ParamsGetCaller) WithValue(value *big.Int) *ParamsGetCaller {
	c.caller.WithValue(value)
	return c
}

func (c *ParamsGetCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

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

// Set is a paid mutator transaction binding the contract method 0x273f4940.
//
// Solidity: function set(bytes32 _key, uint256 _value) returns()
func (_Params *Params) Set(_key [32]byte, _value *big.Int) *contracts.Sender {
	return contracts.NewSender(_Params.contract, "set", _key, _value)
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

// ParamsSetFilterer provides typed access to filtering Set events
type ParamsSetFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// FilterSet is a free log retrieval operation binding the contract event 0x28e3246f80515f5c1ed987b133ef2f193439b25acba6a5e69f219e896fc9d179.
//
// Solidity: event Set(bytes32 indexed key, uint256 value)
func (_Params *Params) FilterSet(criteria []ParamsSetCriteria) *ParamsSetFilterer {
	filterer := _Params.contract.Filter("Set")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := contracts.EventCriteria{}
		if c.Key != nil {
			eventCriteria.Topic1 = *c.Key
		}
		filterer.AddCriteria(eventCriteria)
	}

	return &ParamsSetFilterer{filterer: filterer, contract: _Params.contract}
}

func (f *ParamsSetFilterer) Range(from, to int64) *ParamsSetFilterer {
	f.filterer.Range(from, to)
	return f
}

func (f *ParamsSetFilterer) From(from int64) *ParamsSetFilterer {
	f.filterer.From(from)
	return f
}

func (f *ParamsSetFilterer) To(to int64) *ParamsSetFilterer {
	f.filterer.To(to)
	return f
}

func (f *ParamsSetFilterer) Offset(offset int64) *ParamsSetFilterer {
	f.filterer.Offset(offset)
	return f
}

func (f *ParamsSetFilterer) Limit(limit int64) *ParamsSetFilterer {
	f.filterer.Limit(limit)
	return f
}

func (f *ParamsSetFilterer) Order(order string) *ParamsSetFilterer {
	f.filterer.Order(order)
	return f
}

func (f *ParamsSetFilterer) Execute() ([]ParamsSet, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]ParamsSet, len(logs))
	for i, log := range logs {
		event := new(ParamsSet)
		if err := f.contract.UnpackLog(event, "Set", log); err != nil {
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
