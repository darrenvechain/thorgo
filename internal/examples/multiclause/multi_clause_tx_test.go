package multiclause

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/darrenvechain/thorgo/builtins"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/internal/datagen"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/stretchr/testify/require"
)

func Test_MultiClause_Transaction(t *testing.T) {
	client, cancel := testcontainer.NewSolo()
	t.Cleanup(cancel)

	origin := txmanager.FromPK(solo.Keys()[0], client)
	vtho, err := builtins.NewVTHO(client)
	require.NoError(t, err)

	// Create 100 clauses: 50 VTHO transfers and 50 VET transfers
	clauses := make([]*tx.Clause, 100)
	for i := range 50 {
		to := datagen.RandAddress()
		amount := big.NewInt(int64(i + 1))
		clauses[i], err = vtho.Transfer(to, amount).Clause() // VTHO transfer
		require.NoError(t, err)
		clauses[i+50] = tx.NewClause(&to).WithValue(amount) // VET transfer
	}

	// Send the multi-clause transaction
	visitor, err := origin.SendClauses(clauses, &transactions.Options{})
	require.NoError(t, err)
	tCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	t.Cleanup(cancel)
	receipt, err := visitor.Wait(tCtx)
	require.NoError(t, err)
	require.False(t, receipt.Reverted)

	// Verify that all events are logged
	vthoEvents, err := vtho.FilterTransfer([]builtins.VTHOTransferCriteria{
		{
			From: &receipt.Meta.TxOrigin,
		},
	}).Execute()
	require.NoError(t, err)
	require.Len(t, vthoEvents, 50)

	// Verify VET transfers
	vetEvents, err := client.FilterTransfers(&thorest.TransferFilter{
		Range: &thorest.LogRange{
			From: &receipt.Meta.BlockNumber,
			To:   &receipt.Meta.BlockNumber,
		},
	})
	require.NoError(t, err)
	require.Len(t, vetEvents, 50)
}
