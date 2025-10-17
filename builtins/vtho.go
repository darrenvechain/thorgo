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

// Raw returns the underlying contract.
func (_VTHO *VTHO) Raw() *contracts.Contract {
	return _VTHO.contract
}

// ==================== View Functions ====================

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_VTHO *VTHO) Allowance(_owner common.Address, _spender common.Address) *VTHOAllowanceCaller {
	return &VTHOAllowanceCaller{caller: _VTHO.contract.Call("allowance", _owner, _spender)}
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_VTHO *VTHO) BalanceOf(_owner common.Address) *VTHOBalanceOfCaller {
	return &VTHOBalanceOfCaller{caller: _VTHO.contract.Call("balanceOf", _owner)}
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_VTHO *VTHO) Decimals() *VTHODecimalsCaller {
	return &VTHODecimalsCaller{caller: _VTHO.contract.Call("decimals")}
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_VTHO *VTHO) Name() *VTHONameCaller {
	return &VTHONameCaller{caller: _VTHO.contract.Call("name")}
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_VTHO *VTHO) Symbol() *VTHOSymbolCaller {
	return &VTHOSymbolCaller{caller: _VTHO.contract.Call("symbol")}
}

// TotalBurned is a free data retrieval call binding the contract method 0xd89135cd.
//
// Solidity: function totalBurned() view returns(uint256)
func (_VTHO *VTHO) TotalBurned() *VTHOTotalBurnedCaller {
	return &VTHOTotalBurnedCaller{caller: _VTHO.contract.Call("totalBurned")}
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_VTHO *VTHO) TotalSupply() *VTHOTotalSupplyCaller {
	return &VTHOTotalSupplyCaller{caller: _VTHO.contract.Call("totalSupply")}
}

// ==================== Transaction Functions ====================

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

// ==================== Event Functions ====================

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_VTHO *VTHO) FilterApproval(criteria []VTHOApprovalCriteria) *VTHOApprovalFilterer {
	filterer := _VTHO.contract.Filter("Approval")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.Owner != nil {
			eventCriteria.Topic1 = *c.Owner
		}
		if c.Spender != nil {
			eventCriteria.Topic2 = *c.Spender
		}
		filterer.Criteria(eventCriteria)
	}

	return &VTHOApprovalFilterer{filterer: filterer, contract: _VTHO.contract}
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_VTHO *VTHO) FilterTransfer(criteria []VTHOTransferCriteria) *VTHOTransferFilterer {
	filterer := _VTHO.contract.Filter("Transfer")

	// Add criteria to the filterer
	for _, c := range criteria {
		eventCriteria := &contracts.EventCriteria{}
		if c.From != nil {
			eventCriteria.Topic1 = *c.From
		}
		if c.To != nil {
			eventCriteria.Topic2 = *c.To
		}
		filterer.Criteria(eventCriteria)
	}

	return &VTHOTransferFilterer{filterer: filterer, contract: _VTHO.contract}
}

// ==================== Event Types and Criteria ====================

// VTHOApproval represents a Approval event raised by the VTHO contract.
type VTHOApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Log     *thorest.EventLog
}

type VTHOApprovalCriteria struct {
	Owner   *common.Address
	Spender *common.Address
}

// VTHOTransfer represents a Transfer event raised by the VTHO contract.
type VTHOTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Log   *thorest.EventLog
}

type VTHOTransferCriteria struct {
	From *common.Address
	To   *common.Address
}

// ==================== Call Result Types ====================

// ==================== Caller Types and Methods ====================

// VTHOAllowanceCaller provides typed access to the Allowance method
type VTHOAllowanceCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xdd62ed3e.
func (c *VTHOAllowanceCaller) WithRevision(rev thorest.Revision) *VTHOAllowanceCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xdd62ed3e.
func (c *VTHOAllowanceCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xdd62ed3e and returns the result.
func (c *VTHOAllowanceCaller) Execute() (*big.Int, error) {
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

// VTHOBalanceOfCaller provides typed access to the BalanceOf method
type VTHOBalanceOfCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x70a08231.
func (c *VTHOBalanceOfCaller) WithRevision(rev thorest.Revision) *VTHOBalanceOfCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x70a08231.
func (c *VTHOBalanceOfCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x70a08231 and returns the result.
func (c *VTHOBalanceOfCaller) Execute() (*big.Int, error) {
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

// VTHODecimalsCaller provides typed access to the Decimals method
type VTHODecimalsCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x313ce567.
func (c *VTHODecimalsCaller) WithRevision(rev thorest.Revision) *VTHODecimalsCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x313ce567.
func (c *VTHODecimalsCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x313ce567 and returns the result.
func (c *VTHODecimalsCaller) Execute() (uint8, error) {
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

// VTHONameCaller provides typed access to the Name method
type VTHONameCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x06fdde03.
func (c *VTHONameCaller) WithRevision(rev thorest.Revision) *VTHONameCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x06fdde03.
func (c *VTHONameCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x06fdde03 and returns the result.
func (c *VTHONameCaller) Execute() (string, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero string
		return zero, err
	}
	if len(data) != 1 {
		var zero string
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(string); ok {
		return result, nil
	}
	var zero string
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// VTHOSymbolCaller provides typed access to the Symbol method
type VTHOSymbolCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x95d89b41.
func (c *VTHOSymbolCaller) WithRevision(rev thorest.Revision) *VTHOSymbolCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x95d89b41.
func (c *VTHOSymbolCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x95d89b41 and returns the result.
func (c *VTHOSymbolCaller) Execute() (string, error) {
	data, err := c.caller.Execute()
	if err != nil {
		var zero string
		return zero, err
	}
	if len(data) != 1 {
		var zero string
		return zero, errors.New("expected single return value")
	}
	if result, ok := data[0].(string); ok {
		return result, nil
	}
	var zero string
	return zero, fmt.Errorf("unexpected type returned: %T", data[0])
}

// VTHOTotalBurnedCaller provides typed access to the TotalBurned method
type VTHOTotalBurnedCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0xd89135cd.
func (c *VTHOTotalBurnedCaller) WithRevision(rev thorest.Revision) *VTHOTotalBurnedCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0xd89135cd.
func (c *VTHOTotalBurnedCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0xd89135cd and returns the result.
func (c *VTHOTotalBurnedCaller) Execute() (*big.Int, error) {
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

// VTHOTotalSupplyCaller provides typed access to the TotalSupply method
type VTHOTotalSupplyCaller struct {
	caller *contracts.Caller
}

// WithRevision sets the revision for the call to the contract method 0x18160ddd.
func (c *VTHOTotalSupplyCaller) WithRevision(rev thorest.Revision) *VTHOTotalSupplyCaller {
	c.caller.WithRevision(rev)
	return c
}

// Call executes the raw call to the contract method 0x18160ddd.
func (c *VTHOTotalSupplyCaller) Call() (*thorest.InspectResponse, error) {
	return c.caller.Call()
}

// Execute executes the contract method 0x18160ddd and returns the result.
func (c *VTHOTotalSupplyCaller) Execute() (*big.Int, error) {
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

// ==================== Event Filterer Types and Methods ====================

// VTHOApprovalFilterer provides typed access to filtering Approval events
type VTHOApprovalFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *VTHOApprovalFilterer) Unit(unit string) *VTHOApprovalFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *VTHOApprovalFilterer) IncludeIndexes(include bool) *VTHOApprovalFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *VTHOApprovalFilterer) Range(from, to int64) *VTHOApprovalFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *VTHOApprovalFilterer) From(from int64) *VTHOApprovalFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *VTHOApprovalFilterer) To(to int64) *VTHOApprovalFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *VTHOApprovalFilterer) Offset(offset int64) *VTHOApprovalFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *VTHOApprovalFilterer) Limit(limit int64) *VTHOApprovalFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *VTHOApprovalFilterer) Order(order string) *VTHOApprovalFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *VTHOApprovalFilterer) Execute() ([]VTHOApproval, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]VTHOApproval, len(logs))
	for i, log := range logs {
		event := VTHOApproval{}
		if err := f.contract.UnpackLog(&event, "Approval", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}

// VTHOTransferFilterer provides typed access to filtering Transfer events
type VTHOTransferFilterer struct {
	filterer *contracts.Filterer
	contract *contracts.Contract
}

// Unit sets the range type for the filterer. It can be `block` or `time`
func (f *VTHOTransferFilterer) Unit(unit string) *VTHOTransferFilterer {
	f.filterer.RangeUnit(unit)
	return f
}

// IncludeIndexes sets whether to include transaction and log indexes in the response.
func (f *VTHOTransferFilterer) IncludeIndexes(include bool) *VTHOTransferFilterer {
	f.filterer.IncludeIndexes(include)
	return f
}

// Range sets the range for the filterer. It can be a block range or a time range.
func (f *VTHOTransferFilterer) Range(from, to int64) *VTHOTransferFilterer {
	f.filterer.Range(from, to)
	return f
}

// From sets the start time or block number for the filterer.
func (f *VTHOTransferFilterer) From(from int64) *VTHOTransferFilterer {
	f.filterer.From(from)
	return f
}

// To sets the end time or block number for the filterer.
func (f *VTHOTransferFilterer) To(to int64) *VTHOTransferFilterer {
	f.filterer.To(to)
	return f
}

// Offset sets the offset for the filterer, allowing you to skip a number of events.
func (f *VTHOTransferFilterer) Offset(offset int64) *VTHOTransferFilterer {
	f.filterer.Offset(offset)
	return f
}

// Limit sets the maximum number of events to return.
func (f *VTHOTransferFilterer) Limit(limit int64) *VTHOTransferFilterer {
	f.filterer.Limit(limit)
	return f
}

// Order sets the order of the events returned by the filterer. It can be `asc` or `desc`.
func (f *VTHOTransferFilterer) Order(order string) *VTHOTransferFilterer {
	f.filterer.Order(order)
	return f
}

// Execute the query and return the events matching the filter criteria.
func (f *VTHOTransferFilterer) Execute() ([]VTHOTransfer, error) {
	logs, err := f.filterer.Execute()
	if err != nil {
		return nil, err
	}

	events := make([]VTHOTransfer, len(logs))
	for i, log := range logs {
		event := VTHOTransfer{}
		if err := f.contract.UnpackLog(&event, "Transfer", log); err != nil {
			return nil, err
		}
		event.Log = log
		events[i] = event
	}

	return events, nil
}
