package transfers

import (
	"testing"

	blocks2 "github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/client"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient *client.Client
	blocks     *blocks2.Blocks
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	blocks = blocks2.New(thorClient)
	m.Run()
}

// TestTransfersByBlockRangeASC fetches transfers by block range in ascending order
func TestTransfersByBlockRangeASC(t *testing.T) {
	// Don't apply any criteria, just get all events
	events, err := New(thorClient, []client.TransferCriteria{}).
		BlockRange(0, 1).
		Asc().
		Apply(0, 100)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

// TestTransfersByTimeRangeDESC fetches transfers by time range in descending order
func TestTransfersByTimeRangeDESC(t *testing.T) {
	genesis, err := blocks.ByNumber(0)
	assert.NoError(t, err)
	best, err := blocks.Best()
	assert.NoError(t, err)

	events, err := New(thorClient, []client.TransferCriteria{}).
		TimeRange(genesis.Timestamp, best.Timestamp).
		Desc().
		Apply(0, 100)

	assert.NoError(t, err)
	assert.NotNil(t, events)
}
