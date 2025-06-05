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

// VTHOMetaData contains all meta data concerning the VTHO contract.
var VTHOMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// VTHO is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type VTHO struct {
	thor     *thorest.Client     // Thor client connection to use
	contract *contracts.Contract // Generic contract wrapper for the low level calls
}

// NewVTHO creates a new instance of VTHO, bound to a specific deployed contract.
func NewVTHO(thor *thorest.Client) (*VTHO, error) {
	parsed, err := VTHOMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := contracts.New(thor, common.HexToAddress("0x0000000000000000000000000000456e65726779"), parsed)
	return &VTHO{thor: thor, contract: contract}, nil
}

// Address returns the address of the contract.
func (_VTHO *VTHO) Address() common.Address {
	return _VTHO.contract.Address
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_VTHO *VTHO) Allowance(_owner common.Address, _spender common.Address) *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_VTHO.contract, "allowance", _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_VTHO *VTHO) BalanceOf(_owner common.Address) *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_VTHO.contract, "balanceOf", _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_VTHO *VTHO) Decimals() *contracts.Caller[uint8] {
	return contracts.NewCaller[uint8](_VTHO.contract, "decimals")
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_VTHO *VTHO) Name() *contracts.Caller[string] {
	return contracts.NewCaller[string](_VTHO.contract, "name")
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_VTHO *VTHO) Symbol() *contracts.Caller[string] {
	return contracts.NewCaller[string](_VTHO.contract, "symbol")
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_VTHO *VTHO) TotalBurned() *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_VTHO.contract, "totalBurned")
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VTHO *VTHO) TotalSupply() *contracts.Caller[*big.Int] {
	return contracts.NewCaller[*big.Int](_VTHO.contract, "totalSupply")
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_VTHO *VTHO) Approve(_spender common.Address, _value *big.Int) *contracts.Sender {
	return contracts.NewSender(_VTHO.contract, "approve", _spender, _value)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address _from, address _to, uint256 _amount) returns(bool success)
func (_VTHO *VTHO) Move(_from common.Address, _to common.Address, _amount *big.Int) *contracts.Sender {
	return contracts.NewSender(_VTHO.contract, "move", _from, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_VTHO *VTHO) Transfer(_to common.Address, _amount *big.Int) *contracts.Sender {
	return contracts.NewSender(_VTHO.contract, "transfer", _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_VTHO *VTHO) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) *contracts.Sender {
	return contracts.NewSender(_VTHO.contract, "transferFrom", _from, _to, _amount)
}

// VTHOApproval represents a Approval event raised by the VTHO contract.
type VTHOApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Log     *thorest.EventLog
}

type VTHOApprovalCriteria struct {
	Owner   *common.Address `abi:"_owner"`
	Spender *common.Address `abi:"_spender"`
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_VTHO *VTHO) FilterApproval(criteria []VTHOApprovalCriteria, filters *thorest.LogFilters) ([]VTHOApproval, error) {
	topicHash := _VTHO.contract.ABI.Events["Approval"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_VTHO.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Owner != nil {
			matcher := *c.Owner
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.Spender != nil {
			matcher := *c.Spender
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_VTHO.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _VTHO.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]VTHOApproval, len(logs))
	for i, log := range logs {
		event := new(VTHOApproval)
		if err := _VTHO.contract.UnpackLog(event, "Approval", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchApproval listens for on chain events binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_VTHO *VTHO) WatchApproval(criteria []VTHOApprovalCriteria, ctx context.Context, bufferSize int64) (chan *VTHOApproval, error) {
	topicHash := _VTHO.contract.ABI.Events["Approval"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_VTHO.contract.Address,
			Topic0:  &topicHash,
		}
		if c.Owner != nil {
			matcher := *c.Owner
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.Spender != nil {
			matcher := *c.Spender
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *VTHOApproval, bufferSize)
	blocks := blocks.New(ctx, _VTHO.thor)
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
						ev := new(VTHOApproval)
						if err := _VTHO.contract.UnpackLog(ev, "Approval", log); err != nil {
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

// VTHOTransfer represents a Transfer event raised by the VTHO contract.
type VTHOTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Log   *thorest.EventLog
}

type VTHOTransferCriteria struct {
	From *common.Address `abi:"_from"`
	To   *common.Address `abi:"_to"`
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_VTHO *VTHO) FilterTransfer(criteria []VTHOTransferCriteria, filters *thorest.LogFilters) ([]VTHOTransfer, error) {
	topicHash := _VTHO.contract.ABI.Events["Transfer"].ID

	criteriaSet := make([]thorest.EventCriteria, len(criteria))
	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_VTHO.contract.Address,
			Topic0:  &topicHash,
		}
		if c.From != nil {
			matcher := *c.From
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.To != nil {
			matcher := *c.To
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	if len(criteriaSet) == 0 {
		criteriaSet = append(criteriaSet, thorest.EventCriteria{
			Address: &_VTHO.contract.Address,
			Topic0:  &topicHash,
		})
	}

	logs, err := _VTHO.thor.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	events := make([]VTHOTransfer, len(logs))
	for i, log := range logs {
		event := new(VTHOTransfer)
		if err := _VTHO.contract.UnpackLog(event, "Transfer", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = *event
	}

	return events, nil
}

// WatchTransfer listens for on chain events binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_VTHO *VTHO) WatchTransfer(criteria []VTHOTransferCriteria, ctx context.Context, bufferSize int64) (chan *VTHOTransfer, error) {
	topicHash := _VTHO.contract.ABI.Events["Transfer"].ID
	criteriaSet := make([]thorest.EventCriteria, len(criteria))

	for i, c := range criteria {
		crteria := thorest.EventCriteria{
			Address: &_VTHO.contract.Address,
			Topic0:  &topicHash,
		}
		if c.From != nil {
			matcher := *c.From
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic1 = &topics[0][0]
		}
		if c.To != nil {
			matcher := *c.To
			topics, err := abi.MakeTopics([]interface{}{matcher})
			if err != nil {
				return nil, err
			}
			crteria.Topic2 = &topics[0][0]
		}

		criteriaSet[i] = crteria
	}

	eventChan := make(chan *VTHOTransfer, bufferSize)
	blocks := blocks.New(ctx, _VTHO.thor)
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
						ev := new(VTHOTransfer)
						if err := _VTHO.contract.UnpackLog(ev, "Transfer", log); err != nil {
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
