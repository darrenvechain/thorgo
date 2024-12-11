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

// VTHOMetaData contains all meta data concerning the VTHO contract.
var VTHOMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"move\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBurned\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// VTHO is an auto generated Go binding around an Ethereum contract, allowing you to query and create clauses.
type VTHO struct {
	thor     *thorgo.Thor       // Thor connection to use
	contract *accounts.Contract // Generic contract wrapper for the low level calls
}

// VTHOTransactor is an auto generated Go binding around an Ethereum, allowing you to transact with the contract.
type VTHOTransactor struct {
	*VTHO
	contract *accounts.ContractTransactor // Generic contract wrapper for the low level calls
	manager  accounts.TxManager           // TxManager to use
}

// NewVTHO creates a new instance of VTHO, bound to a specific deployed contract.
func NewVTHO(thor *thorgo.Thor) (*VTHO, error) {
	parsed, err := VTHOMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	contract := thor.Account(common.HexToAddress("0x0000000000000000000000000000456e65726779")).Contract(parsed)
	return &VTHO{thor: thor, contract: contract}, nil
}

// NewVTHOTransactor creates a new instance of VTHOTransactor, bound to a specific deployed contract.
func NewVTHOTransactor(thor *thorgo.Thor, manager accounts.TxManager) (*VTHOTransactor, error) {
	base, err := NewVTHO(thor)
	if err != nil {
		return nil, err
	}
	return &VTHOTransactor{VTHO: base, contract: base.contract.Transactor(manager), manager: manager}, nil
}

// Address returns the address of the contract.
func (_VTHO *VTHO) Address() common.Address {
	return _VTHO.contract.Address
}

// Transactor constructs a new transactor for the contract, which allows to send transactions.
func (_VTHO *VTHO) Transactor(manager accounts.TxManager) *VTHOTransactor {
	return &VTHOTransactor{VTHO: _VTHO, contract: _VTHO.contract.Transactor(manager), manager: manager}
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VTHO *VTHO) Call(revision thorest.Revision, result *[]interface{}, method string, params ...interface{}) error {
	return _VTHO.contract.CallAt(revision, method, result, params...)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VTHOTransactor *VTHOTransactor) Transact(opts *transactions.Options, method string, params ...interface{}) (*transactions.Visitor, error) {
	return _VTHOTransactor.contract.Send(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_VTHO *VTHO) Allowance(_owner common.Address, _spender common.Address, revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _VTHO.Call(rev, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_VTHO *VTHO) BalanceOf(_owner common.Address, revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _VTHO.Call(rev, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_VTHO *VTHO) Decimals(revision ...thorest.Revision) (uint8, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _VTHO.Call(rev, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_VTHO *VTHO) Name(revision ...thorest.Revision) (string, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _VTHO.Call(rev, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_VTHO *VTHO) Symbol(revision ...thorest.Revision) (string, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _VTHO.Call(rev, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_VTHO *VTHO) TotalBurned(revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _VTHO.Call(rev, &out, "totalBurned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VTHO *VTHO) TotalSupply(revision ...thorest.Revision) (*big.Int, error) {
	var rev thorest.Revision
	if len(revision) > 0 {
		rev = revision[0]
	} else {
		rev = thorest.RevisionBest()
	}

	var out []interface{}
	err := _VTHO.Call(rev, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_VTHOTransactor *VTHOTransactor) Approve(_spender common.Address, _value *big.Int, opts *transactions.Options) (*transactions.Visitor, error) {
	return _VTHOTransactor.Transact(opts, "approve", _spender, _value)
}

// ApproveAsClause is a transaction clause generator 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_VTHO *VTHO) ApproveAsClause(_spender common.Address, _value *big.Int, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _VTHO.contract.AsClauseWithVET(val, "approve", _spender, _value)
}

// Move is a paid mutator transaction binding the contract method 0xbb35783b.
//
// Solidity: function move(address _from, address _to, uint256 _amount) returns(bool success)
func (_VTHOTransactor *VTHOTransactor) Move(_from common.Address, _to common.Address, _amount *big.Int, opts *transactions.Options) (*transactions.Visitor, error) {
	return _VTHOTransactor.Transact(opts, "move", _from, _to, _amount)
}

// MoveAsClause is a transaction clause generator 0xbb35783b.
//
// Solidity: function move(address _from, address _to, uint256 _amount) returns(bool success)
func (_VTHO *VTHO) MoveAsClause(_from common.Address, _to common.Address, _amount *big.Int, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _VTHO.contract.AsClauseWithVET(val, "move", _from, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_VTHOTransactor *VTHOTransactor) Transfer(_to common.Address, _amount *big.Int, opts *transactions.Options) (*transactions.Visitor, error) {
	return _VTHOTransactor.Transact(opts, "transfer", _to, _amount)
}

// TransferAsClause is a transaction clause generator 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_VTHO *VTHO) TransferAsClause(_to common.Address, _amount *big.Int, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _VTHO.contract.AsClauseWithVET(val, "transfer", _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_VTHOTransactor *VTHOTransactor) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int, opts *transactions.Options) (*transactions.Visitor, error) {
	return _VTHOTransactor.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFromAsClause is a transaction clause generator 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_VTHO *VTHO) TransferFromAsClause(_from common.Address, _to common.Address, _amount *big.Int, vetValue ...*big.Int) (*tx.Clause, error) {
	var val *big.Int
	if len(vetValue) > 0 {
		val = vetValue[0]
	} else {
		val = big.NewInt(0)
	}
	return _VTHO.contract.AsClauseWithVET(val, "transferFrom", _from, _to, _amount)
}

// VTHOApproval represents a Approval event raised by the VTHO contract.
type VTHOApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Log     thorest.EventLog
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

	logs, err := _VTHO.thor.Client.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	inputs := _VTHO.contract.ABI.Events["Approval"].Inputs
	var indexed abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
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
	ticker := _VTHO.thor.Blocks.Ticker()

	go func() {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				block, err := _VTHO.thor.Blocks.Best()
				if err != nil {
					continue
				}
				for _, tx := range block.Transactions {
					for index, outputs := range tx.Outputs {
						for _, event := range outputs.Events {
							for _, c := range criteriaSet {
								if !c.Matches(event) {
									continue
								}
							}

							log := thorest.EventLog{
								Address: &_VTHO.contract.Address,
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

							ev := new(VTHOApproval)
							if err := _VTHO.contract.UnpackLog(ev, "Approval", log); err != nil {
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

// VTHOTransfer represents a Transfer event raised by the VTHO contract.
type VTHOTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Log   thorest.EventLog
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

	logs, err := _VTHO.thor.Client.FilterEvents(criteriaSet, filters)
	if err != nil {
		return nil, err
	}

	inputs := _VTHO.contract.ABI.Events["Transfer"].Inputs
	var indexed abi.Arguments
	for _, arg := range inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
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
	ticker := _VTHO.thor.Blocks.Ticker()

	go func() {
		defer close(eventChan)

		for {
			select {
			case <-ticker.C():
				block, err := _VTHO.thor.Blocks.Best()
				if err != nil {
					continue
				}
				for _, tx := range block.Transactions {
					for index, outputs := range tx.Outputs {
						for _, event := range outputs.Events {
							for _, c := range criteriaSet {
								if !c.Matches(event) {
									continue
								}
							}

							log := thorest.EventLog{
								Address: &_VTHO.contract.Address,
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

							ev := new(VTHOTransfer)
							if err := _VTHO.contract.UnpackLog(ev, "Transfer", log); err != nil {
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
