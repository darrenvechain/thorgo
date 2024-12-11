package thorest_test

import (
	"testing"

	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestClient_Block(t *testing.T) {
	block, err := thorClient.Block(thorest.RevisionBest())
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

func TestClient_BestBlock(t *testing.T) {
	block, err := thorClient.BestBlock()
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

func TestClient_GenesisBlock(t *testing.T) {
	block, err := thorClient.GenesisBlock()
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

func TestClient_ExpandedBlock(t *testing.T) {
	block, err := thorClient.ExpandedBlock(thorest.RevisionNumber(0))
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

func TestClient_ExpandedBlockWithTxs(t *testing.T) {
	c := thorest.NewClientFromURL("https://mainnet.vechain.org")

	blk, err := c.ExpandedBlock(thorest.RevisionID(common.HexToHash("0x0125fb07988ff3c36b261b5f7227688c1c0473c4873825ac299bc256ea991b0f")))
	assert.NoError(t, err)

	assert.NotNil(t, blk)
	assert.NotNil(t, blk.Transactions)
	assert.Greater(t, len(blk.Transactions), 0)
}

func TestClient_ChainTag(t *testing.T) {
	chainTag, err := thorClient.ChainTag()
	assert.NoError(t, err)
	assert.Equal(t, solo.ChainTag(), chainTag)
}

func TestClient_BlockRef(t *testing.T) {
	genesis, err := thorClient.Block(thorest.RevisionNumber(0))
	assert.NoError(t, err)
	assert.NotNil(t, genesis)
	assert.Equal(t, genesis.BlockRef().Number(), uint32(0))
}
