package delegatedtx

import (
	"context"
	"log/slog"
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/builtins"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/stretchr/testify/require"
)

func Test_DelegatedTx(t *testing.T) {
	client, cancel := testcontainer.NewSolo()
	t.Cleanup(cancel)

	// Create a delegated transaction manager
	origin := txmanager.FromPK(solo.Keys()[0], client)
	gasPayer := txmanager.FromPK(solo.Keys()[1], client)
	txSender := txmanager.NewDelegated(client, origin, gasPayer)

	// Use the `thorgen` CLI to build your own smart contract wrapper
	vtho, err := builtins.NewVTHO(client)
	require.NoError(t, err)

	// Create a new account to receive the tokens
	recipient, err := txmanager.GeneratePK(client)
	require.NoError(t, err)

	// Call the balanceOf function
	balance, err := vtho.BalanceOf(recipient.Address()).Execute()
	require.NoError(t, err)
	slog.Info("recipient balance before", "balance", balance, "error", err)

	amount := big.NewInt(999)
	receipt, err := vtho.Transfer(recipient.Address(), amount).Receipt(context.Background(), txSender)
	require.NoError(t, err)
	require.False(t, receipt.Reverted)

	newBalance, err := vtho.BalanceOf(recipient.Address()).Execute()
	require.NoError(t, err)
	require.Equal(t, amount, newBalance.Sub(newBalance, balance))
}
