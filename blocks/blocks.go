package blocks

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type subscriber struct {
	channel chan *thorest.ExpandedBlock
	ctx     context.Context
}

type Blocks struct {
	client      *thorest.Client
	best        atomic.Value
	subscribers sync.Map // Using sync.Map for concurrent access
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
		nextBlockTime := time.Unix(previous.Timestamp, 0).Add(10 * time.Second)
		now := time.Now().UTC()
		if now.Before(nextBlockTime) {
			time.Sleep(nextBlockTime.Add(100 * time.Millisecond).Sub(now))
		}
		next, err := b.Expanded(thorest.RevisionBest())
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		if next.ID == previous.ID {
			time.Sleep(1 * time.Second)
			continue
		}

		b.subscribers.Range(func(key, value interface{}) bool {
			sub := value.(subscriber)
			go func(sub subscriber) {
				select {
				case <-sub.ctx.Done():
					b.subscribers.Delete(key)
					close(sub.channel)
				default:
					select {
					case sub.channel <- next:
					case <-time.After(1 * time.Second):
						return
					}
				}
			}(sub)
			return true
		})
		previous = next
	}
}

// Subscribe adds a new subscriber to the block stream.
// The subscriber will receive the latest block produced.
// The subscriber will be removed when the context is done.
func (b *Blocks) Subscribe(ctx context.Context, bufferSize int) <-chan *thorest.ExpandedBlock {
	sub := make(chan *thorest.ExpandedBlock, bufferSize)
	s := subscriber{channel: sub, ctx: ctx}
	b.subscribers.Store(uuid.New(), s)
	return sub
}

// ByID returns the block by the given ID.
func (b *Blocks) ByID(id common.Hash) (*thorest.Block, error) {
	return b.client.Block(id.Hex())
}

// Best returns the latest block on chain.
func (b *Blocks) Best() (block *thorest.Block, err error) {
	// Load the best block from the cache.
	if best, ok := b.best.Load().(*thorest.Block); ok {
		// Convert the timestamp to UTC time.
		bestTime := time.Unix(best.Timestamp, 0).UTC()
		if time.Since(bestTime) < 10*time.Second {
			return best, nil
		}
	}

	block, err = b.client.Block("best")
	if err != nil {
		return nil, err
	}

	b.best.Store(block)
	return block, nil
}

// Finalized returns the finalized block.
func (b *Blocks) Finalized() (*thorest.Block, error) {
	return b.client.Block("finalized")
}

// Justified returns the justified block.
func (b *Blocks) Justified() (*thorest.Block, error) {
	return b.client.Block("justified")
}

// ByNumber returns the block by the given number.
func (b *Blocks) ByNumber(number uint64) (*thorest.Block, error) {
	return b.client.Block(fmt.Sprintf("%d", number))
}

// Expanded returns the expanded block information.
// This includes the transactions and receipts.
func (b *Blocks) Expanded(revision thorest.Revision) (*thorest.ExpandedBlock, error) {
	return b.client.ExpandedBlock(revision)
}

// Ticker waits for the next block to be produced
// Returns the next block
func (b *Blocks) Ticker() (*thorest.ExpandedBlock, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sub := b.Subscribe(ctx, 1)
	blk := <-sub
	return blk, nil
}
