package transactions_test

import (
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/crypto/tx"

	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestContractClause(t *testing.T) {
	// account 1
	account1 := solo.Keys()[0]
	account1Addr := crypto.PubkeyToAddress(account1.PublicKey)

	// account 2
	account2 := solo.Keys()[1]
	account2Addr := crypto.PubkeyToAddress(account2.PublicKey)

	// transfer clause
	clause := tx.NewClause(&account2Addr).WithData([]byte{}).WithValue(big.NewInt(1000))
	txbuilder := transactions.NewTransactor(thorClient, []*tx.Clause{clause})

	// simulation
	simulation, err := txbuilder.Simulate(account1Addr, &transactions.Options{})
	assert.NoError(t, err)
	assert.False(t, simulation.Reverted())

	// build
	trx, err := txbuilder.Build(account1Addr, &transactions.Options{})
	assert.NoError(t, err)
	assert.Equal(t, trx.Type(), tx.TypeDynamicFee)

	// sign
	signingHash := trx.SigningHash()
	signature, _ := crypto.Sign(signingHash.Bytes(), account1)
	signedTx := trx.WithSignature(signature)

	// send
	res, err := thorClient.SendTransaction(signedTx)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
