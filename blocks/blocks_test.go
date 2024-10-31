package blocks

import (
	"context"
	"testing"
	"time"

	"github.com/darrenvechain/thorgo/client"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient *client.Client
	blocks     *Blocks
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	blocks = New(thorClient)
	m.Run()
}

// TestGetBestBlock fetches the best block from the network
func TestBlocks_Best(t *testing.T) {
	block, err := blocks.Best()
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

// TestGetBlockByNumber fetches a block by its number
func TestBlocks_ByNumber(t *testing.T) {
	block, err := blocks.ByNumber(0)
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

// TestGetBlockByID fetches a block by its ID
func TestBlocks_ByID(t *testing.T) {
	block, err := blocks.ByID(solo.GenesisID())
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

// TestGetFinalizedBlock fetches the finalized block from the network
func TestBlocks_Finalized(t *testing.T) {
	block, err := blocks.Finalized()
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

// TestGetExpandedBlock fetches a block where all the transactions are expanded
// It accepts a revision, which can be a block ID, block number, "best" or "finalized"
func TestBlocks_Expanded(t *testing.T) {
	block, err := blocks.Expanded(solo.GenesisID().Hex())
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

// TestWaitForNextBlock waits for the next block to be produced
func TestBlocks_Ticker(t *testing.T) {
	block, err := blocks.Ticker()
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

func TestBlocks_Subscribe(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sub := blocks.Subscribe(ctx)
	ticker := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-ticker.C:
			t.Fatal("timed out waiting for block")
		case blk := <-sub:
			assert.NotNil(t, blk)
			return
		}
	}
}

func TestBlocks_Unsubscribe(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// cancel the context before we start the subscription
	cancel()
	sub := blocks.Subscribe(ctx)
	blk, ok := <-sub
	assert.Nil(t, blk)
	assert.False(t, ok)
}
