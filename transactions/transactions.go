package transactions

import (
	"context"
	"errors"
	"fmt"

	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
)

// Visitor is a struct that provides methods to interact with transactions on the VeChainThor blockchain.
type Visitor struct {
	client *thorest.Client
	hash   common.Hash
}

// New creates a new instance of the Visitor struct.
func New(client *thorest.Client, hash common.Hash) *Visitor {
	return &Visitor{client: client, hash: hash}
}

// ID returns the transaction ID.
func (v *Visitor) ID() common.Hash {
	return v.hash
}

// Get fetches the transaction by its hash. This includes the clauses, but not the outputs.
func (v *Visitor) Get() (*thorest.Transaction, error) {
	return v.client.Transaction(v.hash)
}

// Receipt fetches the transaction receipt by its hash. This includes the outputs.
func (v *Visitor) Receipt() (*thorest.TransactionReceipt, error) {
	return v.client.TransactionReceipt(v.hash)
}

// RevertReason fetches the revert reason for the transaction.
func (v *Visitor) RevertReason() (*RevertReason, error) {
	receipt, err := v.client.TransactionReceipt(v.hash)
	if err != nil {
		return nil, err
	}
	if !receipt.Reverted {
		return nil, errors.New("transaction did not revert")
	}
	res, err := v.client.DebugRevertReason(receipt)
	if err != nil {
		return nil, err
	}
	return &RevertReason{res: res}, nil
}

// Raw fetches the raw transaction by its hash.
func (v *Visitor) Raw() (*thorest.RawTransaction, error) {
	return v.client.RawTransaction(v.hash)
}

// Pending includes the transaction in the pending pool when querying for a transaction.
func (v *Visitor) Pending() (*thorest.Transaction, error) {
	return v.client.PendingTransaction(v.hash)
}

// Wait for the transaction to be mined. This function will block until the transaction has been mined or the context is done.
func (v *Visitor) Wait(ctx context.Context) (*thorest.TransactionReceipt, error) {
	receipt, err := v.client.TransactionReceipt(v.hash)
	if err == nil {
		return receipt, nil
	}

	blks := blocks.New(ctx, v.client)
	ticker := blks.Ticker()

	for {
		select {
		case <-ctx.Done():
			_, err := v.client.PendingTransaction(v.hash)
			if err == nil {
				return nil, fmt.Errorf("context cancelled while waiting for transaction %s, it is still pending", v.hash.Hex())
			}
			return nil, fmt.Errorf("context cancelled while waiting for transaction")
		case <-ticker.C():
			receipt, err = v.client.TransactionReceipt(v.hash)
			if err == nil {
				return receipt, nil
			}
		}
	}
}
