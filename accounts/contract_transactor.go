package accounts

import (
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/transactions"
)

type TxManager interface {
	SendClauses(clauses []*tx.Clause, opts *transactions.Options) (*transactions.Visitor, error)
}

type ContractTransactor struct {
	*Contract
	manager TxManager
}

// Send sends a transaction to the contract with the given method and arguments.
func (c *ContractTransactor) Send(opts *transactions.Options, method string, args ...any) (*transactions.Visitor, error) {
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
	return c.manager.SendClauses([]*tx.Clause{clause}, opts)
}
