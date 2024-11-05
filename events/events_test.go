package events

import (
	"testing"

	"github.com/darrenvechain/thorgo/blocks"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient *thorest.Client
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	m.Run()
}

// TestEventByBlockRangeASC fetches events by block range in ascending order
func TestEventByBlockRangeASC(t *testing.T) {
	// Don't apply any criteria, just get all events
	events, err := New(thorClient, []thorest.EventCriteria{}).
		BlockRange(0, 1).
		Asc().
		Apply(0, 100)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

// TestEventsByTimeRangeDESC fetches events by time range in descending order
func TestEventsByTimeRangeDESC(t *testing.T) {
	genesis, err := blocks.New(thorClient).ByNumber(0)
	assert.NoError(t, err)
	best, err := blocks.New(thorClient).Best()
	assert.NoError(t, err)

	events, err := New(thorClient, []thorest.EventCriteria{}).
		TimeRange(genesis.Timestamp, best.Timestamp).
		Desc().
		Apply(0, 100)

	assert.NoError(t, err)
	assert.NotNil(t, events)
}
