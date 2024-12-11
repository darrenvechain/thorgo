package blocks

import (
	"errors"
	"sync/atomic"
	"time"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
)

type Blocks struct {
	client *thorest.Client
	best   atomic.Value
	signal Signal
}

func New(c *thorest.Client) *Blocks {
	b := &Blocks{client: c}
	go b.poll()
	return b
}

// poll sends the expanded block to all active subscribers.
func (b *Blocks) poll() {
	var previous *thorest.ExpandedBlock
	var err error
	backoff := 5 * time.Second

	for {
		previous, err = b.Expanded(thorest.RevisionBest())
		if err != nil {
			time.Sleep(backoff)
			continue
		}
		break
	}

	for {
		next, err := b.Expanded(thorest.RevisionBest())
		if errors.Is(err, thorest.ErrNotFound) {
			time.Sleep(250 * time.Millisecond)
			continue
		}
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		if next.ID == previous.ID {
			time.Sleep(250 * time.Millisecond)
			continue
		}
		previous = next
		b.best.Store(next)
		b.signal.Broadcast()

		nextBlockTime := time.Unix(previous.Timestamp, 0).Add(6 * time.Second)
		now := time.Now().UTC()
		if now.Before(nextBlockTime) {
			time.Sleep(nextBlockTime.Add(100 * time.Millisecond).Sub(now))
		}
	}
}

// Ticker creates a signal Waiter to receive event that the best block changed.
func (b *Blocks) Ticker() Waiter {
	return b.signal.NewWaiter()
}

// ByID returns the block by the given ID.
func (b *Blocks) ByID(id common.Hash) (*thorest.Block, error) {
	return b.client.Block(thorest.RevisionID(id))
}

// Best returns the latest block on chain.
func (b *Blocks) Best() (block *thorest.ExpandedBlock, err error) {
	// Load the best block from the cache.
	if best, ok := b.best.Load().(*thorest.ExpandedBlock); ok {
		// Convert the timestamp to UTC time.
		bestTime := time.Unix(best.Timestamp, 0).UTC()
		if time.Since(bestTime) < 10*time.Second {
			return best, nil
		}
	}

	block, err = b.client.ExpandedBlock(thorest.RevisionBest())
	if err != nil {
		return nil, err
	}

	b.best.Store(block)
	return block, nil
}

// Finalized returns the finalized block.
func (b *Blocks) Finalized() (*thorest.Block, error) {
	return b.client.Block(thorest.RevisionFinalized())
}

// Justified returns the justified block.
func (b *Blocks) Justified() (*thorest.Block, error) {
	return b.client.Block(thorest.RevisionJustified())
}

// ByNumber returns the block by the given number.
func (b *Blocks) ByNumber(number int64) (*thorest.Block, error) {
	return b.client.Block(thorest.RevisionNumber(number))
}

// Expanded returns the expanded block information.
// This includes the transactions and receipts.
func (b *Blocks) Expanded(revision thorest.Revision) (*thorest.ExpandedBlock, error) {
	return b.client.ExpandedBlock(revision)
}
