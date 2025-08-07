package contracts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/contracts"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterer_Limit(t *testing.T) {
	receiver, err := txmanager.GeneratePK(thorClient)
	assert.NoError(t, err)

	_, err = vthoRaw.
		Send("transfer", receiver.Address(), big.NewInt(1000)).
		Receipt(context.Background(), account1)
	require.NoError(t, err)

	events, err := contracts.NewFilterer(vthoRaw, "Transfer").
		Limit(1).
		Criteria(&contracts.EventCriteria{
			Topic2: receiver.Address(),
		}).
		Execute()
	require.NoError(t, err)
	require.Len(t, events, 1)
}
