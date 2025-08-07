package logs_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/builtins"
	"github.com/darrenvechain/thorgo/contracts"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/logs"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient     *thorest.Client
	vthoContract   *builtins.VTHO
	vthoRaw        *contracts.Contract
	account1       *txmanager.PKManager
	account2       *txmanager.PKManager
	firstBlock     int64
	lastBlock      int64
	firstTimestamp int64
	lastTimestamp  int64
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	vthoContract, _ = builtins.NewVTHO(thorClient)
	abi, _ := builtins.VTHOMetaData.GetAbi()
	vthoRaw = contracts.New(thorClient, vthoContract.Address(), abi)
	account1 = txmanager.FromPK(solo.Keys()[0], thorClient)
	account2 = txmanager.FromPK(solo.Keys()[1], thorClient)

	for i := range int64(5) {
		receipt, err := vthoRaw.
			Send("transfer", account2.Address(), big.NewInt(i*1000)).
			Receipt(context.Background(), account1)
		if err != nil {
			panic(err)
		}

		if i == 0 {
			firstBlock = receipt.Meta.BlockNumber
			firstTimestamp = receipt.Meta.BlockTimestamp
		}
		if i == 4 {
			lastBlock = receipt.Meta.BlockNumber
			lastTimestamp = receipt.Meta.BlockTimestamp
		}
	}

	m.Run()
}

func TestEventFilterer_Range(t *testing.T) {
	// filter by blocks
	events, err := logs.NewEventsFilterer(thorClient).
		RangeUnit("block").
		Range(firstBlock, lastBlock).
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 5)

	// filter by timestamps
	events, err = logs.NewEventsFilterer(thorClient).
		RangeUnit("time").
		Range(firstTimestamp, lastTimestamp-10).
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 4)

	// no range
	events, err = logs.NewEventsFilterer(thorClient).
		Criteria(thorest.EventCriteria{
			Address: &vthoRaw.Address,
		}).
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 5)

	// only from
	events, err = logs.NewEventsFilterer(thorClient).
		From(lastBlock).
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 1)

	// only to
	events, err = logs.NewEventsFilterer(thorClient).
		To(lastBlock - 3).
		Criteria(thorest.EventCriteria{
			Address: &vthoRaw.Address,
		}).
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 2)
}

func TestEventFilterer_Options(t *testing.T) {
	events, err := logs.NewEventsFilterer(thorClient).
		Offset(3).
		Limit(2).
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 2)

	events, err = logs.NewEventsFilterer(thorClient).
		Limit(2).
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 2)
}

func TestEventFilterer_Order(t *testing.T) {
	// ascending order
	events, err := logs.NewEventsFilterer(thorClient).
		From(firstBlock).
		Order("asc").
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 5)
	assert.Equal(t, firstBlock, events[0].Meta.BlockNumber)

	// descending order
	events, err = logs.NewEventsFilterer(thorClient).
		From(firstBlock).
		Order("desc").
		Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 5)
	assert.Equal(t, lastBlock, events[0].Meta.BlockNumber)
}
