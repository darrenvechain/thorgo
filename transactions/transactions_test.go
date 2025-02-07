package transactions_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/internal/testcontract"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient *thorest.Client
	account1   *txmanager.PKManager
	account2   *txmanager.PKManager
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	account1 = txmanager.FromPK(solo.Keys()[0], thorClient)
	account2 = txmanager.FromPK(solo.Keys()[1], thorClient)
	m.Run()
}

// TestTransactions demonstrates how to build, sign, send, and wait for a transaction
func TestTransactions(t *testing.T) {
	// build a transaction
	to := account2.Address()
	vetClause := tx.NewClause(&to).WithValue(big.NewInt(1000))
	unsigned, err := transactions.NewTransactor(thorClient, []*tx.Clause{vetClause}).Build(account1.Address(), &transactions.Options{})
	assert.NoError(t, err)

	// sign it
	signature, err := account1.SignTransaction(unsigned)
	assert.NoError(t, err)
	signed := unsigned.WithSignature(signature)

	// send it
	res, err := thorClient.SendTransaction(signed)
	assert.NoError(t, err)

	tx := transactions.New(thorClient, res.ID)

	// fetch the pending transaction
	pending, err := tx.Pending()
	assert.NoError(t, err)
	assert.NotNil(t, pending)

	// wait for the receipt
	receipt, err := tx.Wait(context.Background())
	assert.NoError(t, err)
	assert.False(t, receipt.Reverted)

	// raw tx
	raw, err := tx.Raw()
	assert.NoError(t, err)
	assert.NotNil(t, raw)
}

func TestVisitor_RevertReason(t *testing.T) {
	balance := big.NewInt(1000)
	transferAmount := big.NewInt(1001)

	// setup contracts + funding
	deploymentTxID, erc20, err := testcontract.DeployErc20(context.Background(), thorClient, account1, &transactions.Options{}, "Erc20", "ERC")
	assert.NoError(t, err)
	_, err = transactions.New(thorClient, deploymentTxID).Wait(context.Background())
	assert.NoError(t, err)
	erc20Funding, err := erc20.Mint(account1.Address(), balance, &transactions.Options{})
	assert.NoError(t, err)
	_, err = erc20Funding.Wait(context.Background())
	assert.NoError(t, err)

	// send funds too much erc20 tokens
	transfer, err := erc20.Transfer(account2.Address(), transferAmount, &transactions.Options{})
	assert.NoError(t, err)
	receipt, err := transfer.Wait(context.Background())
	assert.NoError(t, err)
	assert.True(t, receipt.Reverted)

	// get revert reason
	erc20ABI, err := testcontract.Erc20MetaData.GetAbi()
	assert.NoError(t, err)
	reason, err := transfer.RevertReason()
	assert.NoError(t, err)

	type ERC20InsufficientBalance struct {
		Sender  common.Address
		Balance *big.Int
		Needed  *big.Int
	}
	errABI := erc20ABI.Errors["ERC20InsufficientBalance"]

	// decode revert reason
	var decoded = ERC20InsufficientBalance{}
	err = reason.DecodeInto(errABI, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, decoded.Sender, account1.Address())
	assert.Equal(t, decoded.Balance, balance)
	assert.Equal(t, decoded.Needed, transferAmount)
}
