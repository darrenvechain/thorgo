package accounts

import (
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/common"
)

type TxManager interface {
	SendClauses(clauses []*tx.Clause, opts *transactions.Options) (common.Hash, error)
}

type ContractTransactor struct {
	*Contract
	manager TxManager
}

// Send sends a transaction to the contract with the given method and arguments.
func (c *ContractTransactor) Send(opts *transactions.Options, method string, args ...interface{}) (*transactions.Visitor, error) {
	if opts == nil {
		opts = &transactions.Options{}
	}
	if opts.VET == nil {
		opts.VET = new(big.Int)
	}
	clause, err := c.AsClauseWithVET(opts.VET, method, args...)
	if err != nil {
		return &transactions.Visitor{}, fmt.Errorf("failed to pack method %s: %w", method, err)
	}
	txId, err := c.manager.SendClauses([]*tx.Clause{clause}, opts)
	if err != nil {
		return &transactions.Visitor{}, fmt.Errorf("failed to send transaction: %w", err)
	}
	return transactions.New(c.client, txId), nil
}
