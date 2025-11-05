package txbuilding

import (
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/internal/datagen"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/stretchr/testify/require"
)

func Test_Example_TxBuilding(t *testing.T) {
	// setup client & tx manager
	client, cancel := testcontainer.NewSolo()
	t.Cleanup(cancel)
	origin := txmanager.FromPK(solo.Keys()[0], client)
	to := datagen.RandAddress()

	// simple VET transfer
	clause := tx.NewClause(&to).WithValue(big.NewInt(1))

	// use custom options when sending the transaction
	options := new(transactions.OptionsBuilder).
		Gas(50_000).
		Expiration(100).
		BlockRef(tx.NewBlockRef(0)).
		Build()

	// create a transactor with the clause
	transactor := transactions.NewTransactor(client, []*tx.Clause{clause})

	// simulate the transaction before sending
	simulation, err := transactor.Simulate(origin.Address(), options)
	require.NoError(t, err)
	require.False(t, simulation.Reverted())

	// build and send the transaction
	tx, err := transactor.Build(origin.Address(), options)
	require.NoError(t, err)
	signature, err := origin.SignTransaction(tx)
	require.NoError(t, err)
	tx = tx.WithSignature(signature)

	// send the transaction
	result, err := client.SendTransaction(tx)
	require.NoError(t, err)
	require.NotNil(t, result.ID)
}
