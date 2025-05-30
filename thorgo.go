package thorgo

import (
	"context"

	"github.com/darrenvechain/thorgo/accounts"
	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Thor struct {
	blocks *blocks.Blocks
	client *thorest.Client
}

// New creates a new instance of the Thor client with the given context and URL.
func New(ctx context.Context, url string) *Thor {
	c := thorest.NewClientFromURL(url)
	return &Thor{client: c, blocks: blocks.New(ctx, c)}
}

// NewFromClient creates a new instance of the Thor client with the given context and thorest client.
func NewFromClient(ctx context.Context, c *thorest.Client) *Thor {
	return &Thor{client: c, blocks: blocks.New(ctx, c)}
}

// Blocks provides utility functions to fetch or wait for blocks.
func (t *Thor) Blocks() *blocks.Blocks {
	return t.blocks
}

// Client returns the underlying thorest client.
func (t *Thor) Client() *thorest.Client {
	return t.client
}

// Account can be used to query account information such as balance, code, storage, etc.
// It also provides a way to interact with contracts.
func (t *Thor) Account(address common.Address) *accounts.Visitor {
	return accounts.New(t.client, address)
}

// Transaction provides utility functions to fetch or wait for transactions and their receipts.
func (t *Thor) Transaction(hash common.Hash) *transactions.Visitor {
	return transactions.New(t.client, hash)
}

// Transactor creates a new transaction builder which makes it easier to build, simulate and send transactions.
func (t *Thor) Transactor(clauses []*tx.Clause) *transactions.Transactor {
	return transactions.NewTransactor(t.client, clauses)
}

// Events sets up a query builder to fetch smart contract solidity events.
func (t *Thor) Events(criteria []thorest.EventCriteria, filters *thorest.LogFilters) ([]*thorest.EventLog, error) {
	return t.client.FilterEvents(criteria, filters)
}

// Transfers sets up a query builder to fetch VET transfers.
func (t *Thor) Transfers(criteria []thorest.TransferCriteria, filters *thorest.LogFilters) ([]*thorest.TransferLog, error) {
	return t.client.FilterTransfers(criteria, filters)
}

// Deployer makes it easier to deploy contracts.
func (t *Thor) Deployer(bytecode []byte, abi *abi.ABI) *accounts.Deployer {
	return accounts.NewDeployer(t.client, bytecode, abi)
}
