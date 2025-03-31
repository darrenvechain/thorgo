package thorest_test

import (
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/internal/datagen"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestClient_SendTransaction(t *testing.T) {
	account1 := solo.Keys()[0]
	account2 := solo.Keys()[1]
	account2Addr := crypto.PubkeyToAddress(account2.PublicKey)

	vetClause := tx.NewClause(&account2Addr).
		WithValue(big.NewInt(1000))

	tag, err := thorClient.ChainTag()
	assert.NoError(t, err)

	txBody := tx.NewBuilder(tx.TypeLegacy).
		Gas(3_000_000).
		GasPriceCoef(255).
		ChainTag(tag).
		Expiration(100000000).
		BlockRef(tx.NewBlockRef(0)).
		Nonce(datagen.RandUint64()).
		Clause(vetClause).
		Build()

	signingHash := txBody.SigningHash()
	signature, err := crypto.Sign(signingHash[:], account1)
	assert.NoError(t, err)

	signedTx := txBody.WithSignature(signature)

	res, err := thorClient.SendTransaction(signedTx)
	assert.NoError(t, err)
	assert.Equal(t, signedTx.ID().String(), res.ID.String())

	tx, err := thorClient.PendingTransaction(signedTx.ID())
	assert.NoError(t, err)
	assert.Equal(t, signedTx.ID().String(), tx.ID.String())
}
