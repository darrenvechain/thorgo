package contractgen

import (
	"context"
	"testing"
	"time"

	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/stretchr/testify/require"
)

func TestEcho(t *testing.T) {
	// setup test container and tx manager
	client, cancel := testcontainer.NewSolo()
	t.Cleanup(cancel)
	manager := txmanager.FromPK(solo.Keys()[0], client)

	// setup tx receipt context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	t.Cleanup(cancel)

	// deploy the contract
	_, contract, err := DeployEchoContract(ctx, client, manager, &transactions.Options{})
	require.NoError(t, err)

	// interact with the contract
	visitor, err := contract.Echo("Hello, World!").Send(manager)
	require.NoError(t, err)
	receipt, err := visitor.Wait(ctx)
	require.NoError(t, err)
	require.False(t, receipt.Reverted)

	// filter events
	events, err := contract.FilterEchoed().Execute()
	require.NoError(t, err)
	require.Len(t, events, 1)
	require.Equal(t, "Hello, World!", events[0].Message)
}
