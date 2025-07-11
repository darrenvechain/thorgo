package txmanager_test

import (
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/contracts"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/stretchr/testify/assert"
)

var (
	client *thorest.Client
)

func TestMain(m *testing.M) {
	var cancel func()
	client, cancel = testcontainer.NewSolo()
	defer cancel()
	m.Run()
}

var (
	// PKManager should implement accounts.TxManager
	_ contracts.TxManager = &txmanager.PKManager{}
)

// TestPKSigner demonstrates ease the ease of sending a transaction using a private key signer
func TestPKSigner(t *testing.T) {
	signer, err := txmanager.GeneratePK(client)
	assert.NoError(t, err)

	to, err := txmanager.GeneratePK(client)
	assert.NoError(t, err)
	toAddr := to.Address()
	vetClause := tx.NewClause(&toAddr).WithValue(big.NewInt(1000))

	tx := tx.NewBuilder(tx.TypeLegacy).
		GasPriceCoef(1).
		Gas(100000).
		Clause(vetClause).
		ChainTag(10).
		BlockRef(tx.NewBlockRef(100)).
		Build()

	signature, err := signer.SignTransaction(tx)
	assert.NoError(t, err)
	signedTx := tx.WithSignature(signature)
	origin, err := signedTx.Origin()
	assert.NoError(t, err)
	assert.Equal(t, signer.Address(), origin)
}
