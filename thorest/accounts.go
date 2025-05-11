package thorest

import (
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Account represents the state of an account on the VeChainThor blockchain.
type Account struct {
	Balance *hexutil.Big `json:"balance"` // Balance in wei
	Energy  *hexutil.Big `json:"energy"`  // Energy in wei
	HasCode bool         `json:"hasCode"` // HasCode indicates if the account is a smart contract
}

// AccountCode represents the bytecode of a smart contract account.
type AccountCode struct {
	Code hexutil.Bytes `json:"code"` // Code is the bytecode of the smart contract
}

// AccountStorage represents the storage value of an account storage slot.
type AccountStorage struct {
	Value common.Hash `json:"value"` // Value is the value stored at the specified storage slot
}

// InspectRequest represents a request to inspect a transaction or contract call.
type InspectRequest struct {
	Gas        *uint64         `json:"gas,omitempty"`        // Gas (optional) is used inside the builtin contracts
	GasPrice   *uint64         `json:"gasPrice,omitempty"`   // GasPrice (optional) is used inside the builtin contracts
	Caller     *common.Address `json:"caller,omitempty"`     // Caller (optional) is the calling account for the clauses
	ProvedWork *string         `json:"provedWork,omitempty"` // ProvedWork (optional) is used inside the builtin contracts
	GasPayer   *common.Address `json:"gasPayer,omitempty"`   // GasPayer (optional) is the account that pays for the gas
	Expiration *uint64         `json:"expiration,omitempty"` // Expiration (optional) is used inside the builtin contracts
	BlockRef   *hexutil.Big    `json:"blockRef,omitempty"`   // BlockRef (optional) is used inside the builtin contracts
	Clauses    []*tx.Clause    `json:"clauses"`              // Clauses is the list of clauses to inspect
}

// InspectResponse represents the response from an inspect request.
type InspectResponse struct {
	Data      hexutil.Bytes `json:"data"`      // Data is the output data from the contract call
	Events    []Event       `json:"events"`    // Events represent any smart contract events emitted during the call
	Transfers []Transfer    `json:"transfers"` // Transfers represent any VET transfers made during the call
	GasUsed   uint64        `json:"gasUsed"`   // GasUsed is the amount of gas used during the call
	Reverted  bool          `json:"reverted"`  // Reverted indicates if the call was reverted
	VmError   string        `json:"vmError"`   // VmError is the error message if the EVM experienced an error
}
