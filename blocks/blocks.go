package blocks

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/darrenvechain/thorgo/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type subscriber struct {
	sub chan *client.ExpandedBlock
	ctx context.Context
}

type Blocks struct {
	client      *client.Client
	best        atomic.Value
	subscribers sync.Map // Using sync.Map for concurrent access
}

func New(c *client.Client) *Blocks {
	b := &Blocks{client: c}
	go b.poll()
	return b
}

// poll sends the expanded block to all active subscribers.
func (b *Blocks) poll() {
	var previous *client.ExpandedBlock
	var err error
	backoff := 5 * time.Second

	for {
		previous, err = b.Expanded("best")
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

		next, err := b.Expanded("best")
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		if previous.ID != next.ID {
			b.subscribers.Range(func(key, value interface{}) bool {
				sub := value.(subscriber)
				select {
				case <-sub.ctx.Done():
					b.subscribers.Delete(key)
					close(sub.sub)
					return false
				default:
					sub.sub <- next
				}
				return true
			})
			previous = next
		} else {
			// Sleep for a second if the block hasn't changed.
			time.Sleep(1 * time.Second)
			continue
		}
	}
}

// Subscribe adds a new subscriber to the block stream.
// The subscriber will receive the latest block produced.
// The subscriber will be removed when the context is done.
func (b *Blocks) Subscribe(ctx context.Context) <-chan *client.ExpandedBlock {
	sub := make(chan *client.ExpandedBlock)
	id := uuid.New().String()
	s := subscriber{sub: sub, ctx: ctx}
	b.subscribers.Store(id, s)
	return sub
}

// ByID returns the block by the given ID.
func (b *Blocks) ByID(id common.Hash) (*client.Block, error) {
	return b.client.Block(id.Hex())
}

// Best returns the latest block on chain.
func (b *Blocks) Best() (block *client.Block, err error) {
	// Load the best block from the cache.
	if best, ok := b.best.Load().(*client.Block); ok {
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
func (b *Blocks) Finalized() (*client.Block, error) {
	return b.client.Block("finalized")
}

// Justified returns the justified block.
func (b *Blocks) Justified() (*client.Block, error) {
	return b.client.Block("justified")
}

// ByNumber returns the block by the given number.
func (b *Blocks) ByNumber(number uint64) (*client.Block, error) {
	return b.client.Block(fmt.Sprintf("%d", number))
}

// Expanded returns the expanded block information.
// This includes the transactions and receipts.
func (b *Blocks) Expanded(revision string) (*client.ExpandedBlock, error) {
	return b.client.ExpandedBlock(revision)
}

// Ticker waits for the next block to be produced
// Returns the next block
func (b *Blocks) Ticker() (*client.ExpandedBlock, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sub := b.Subscribe(ctx)
	blk := <-sub
	return blk, nil
}
