package blocks

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Blocks provides utility functions to fetch or wait for blocks.
type Blocks struct {
	ctx     context.Context
	client  *thorest.Client
	best    atomic.Value
	subs    event.SubscriptionScope
	feed    event.Feed
	done    chan struct{}
	mu      sync.Mutex
	polling bool
}

// New creates a new Blocks instance and starts polling for the best block.
func New(ctx context.Context, c *thorest.Client) *Blocks {
	return &Blocks{client: c, done: make(chan struct{}), ctx: ctx}
}

// Close stops the block polling and closes all active subscriptions.
func (b *Blocks) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.polling {
		return
	}

	close(b.done)
	b.subs.Close()
	b.polling = false
}

// poll sends the expanded block to all active subscribers when a new best block is detected.
func (b *Blocks) poll(ctx context.Context) {
	var previous *thorest.ExpandedBlock

	for {
		select {
		case <-b.done:
			return
		case <-ctx.Done():
			return
		default:
			next, err := b.Expanded(thorest.RevisionBest())
			if errors.Is(err, thorest.ErrNotFound) {
				time.Sleep(250 * time.Millisecond)
				continue
			}
			if err != nil {
				time.Sleep(2 * time.Second)
				continue
			}
			if previous != nil && next.ID == previous.ID {
				time.Sleep(250 * time.Millisecond)
				continue
			}
			previous = next
			b.best.Store(next)
			b.feed.Send(next)

			nextBlockTime := time.Unix(previous.Timestamp, 0).Add(6 * time.Second)
			now := time.Now().UTC()
			if now.Before(nextBlockTime) {
				time.Sleep(nextBlockTime.Add(100 * time.Millisecond).Sub(now))
			}
		}
	}
}

// Subscribe to new blocks. When the best block ID changes, the new expanded block will be sent to the given channel.
// If the polling is not already started, it will be started.
func (b *Blocks) Subscribe(blockChan chan *thorest.ExpandedBlock) event.Subscription {
	b.mu.Lock()
	defer b.mu.Unlock()

	if !b.polling {
		b.polling = true
		b.done = make(chan struct{})
		b.subs = event.SubscriptionScope{}
		b.feed = event.Feed{}
		go b.poll(b.ctx)
	}

	return b.subs.Track(b.feed.Subscribe(blockChan))
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
		if time.Since(bestTime) < 6*time.Second {
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
