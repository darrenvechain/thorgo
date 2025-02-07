package blocks

import (
	"context"
	"testing"
	"time"

	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient *thorest.Client
	blocks     *Blocks
	ctx        = context.Background()
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	blocks = New(ctx, thorClient)
	m.Run()
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
	block, err := blocks.Expanded(thorest.RevisionID(solo.GenesisID()))
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

// TestWaitForNextBlock waits for the next block to be produced
func TestBlocks_Ticker(t *testing.T) {
	ticker := blocks.Ticker()
	timeout := time.NewTimer(12 * time.Second)

	select {
	case <-timeout.C:
		t.Fatal("timed out waiting for the next block")
	case <-ticker.C():
		return
	}
}
