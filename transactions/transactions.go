package transactions

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
)

type Visitor struct {
	client *thorest.Client
	hash   common.Hash
	blocks *blocks.Blocks
}

func New(client *thorest.Client, hash common.Hash) *Visitor {
	return &Visitor{client: client, hash: hash, blocks: blocks.New(client)}
}

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

// Wait for the transaction to be included in a block.
// It will wait for ~6 blocks to be produced.
func (v *Visitor) Wait() (*thorest.TransactionReceipt, error) {
	return v.WaitFor(time.Minute)
}

// WaitFor the transaction to be included in a block.
// It will wait for the given duration.
// If the transaction is not included in a block within the duration, it will return an error.
func (v *Visitor) WaitFor(duration time.Duration) (*thorest.TransactionReceipt, error) {
	receipt, err := v.client.TransactionReceipt(v.hash)
	if err == nil {
		return receipt, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ticker := v.blocks.Ticker()

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("timed out waiting for the tx receipt %s", v.hash.String())
		case <-ticker.C():
			receipt, err = v.client.TransactionReceipt(v.hash)
			if err == nil {
				return receipt, nil
			}
		}
	}
}
