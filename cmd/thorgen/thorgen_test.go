package main_test

import (
	"context"
	"math/big"
	"testing"
	"time"

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
	client, cancel = testcontainer.NewSolo()
)

func deployErc20(client *thorest.Client, txsender *txmanager.PKManager) (common.Hash, *testcontract.Erc20Transactor, error) {
	return testcontract.DeployErc20(context.Background(), client, txsender, &transactions.Options{}, "TestToken", "TT")
}

func TestMain(m *testing.M) {
	m.Run()
	cancel()
}

func TestDeploy(t *testing.T) {
	txsender := txmanager.FromPK(solo.Keys()[0], client)

	_, _, err := deployErc20(client, txsender)
	assert.NoError(t, err)
}

func TestCall(t *testing.T) {
	txsender := txmanager.FromPK(solo.Keys()[0], client)

	_, erc20, err := deployErc20(client, txsender)
	assert.NoError(t, err)

	supply, err := erc20.TotalSupply()
	assert.NoError(t, err)

	assert.True(t, supply.Cmp(big.NewInt(0)) == 0)
}

func TestTransactor(t *testing.T) {
	txsender := txmanager.FromPK(solo.Keys()[0], client)
	receiver := txmanager.FromPK(solo.Keys()[1], client)

	_, erc20, err := deployErc20(client, txsender)
	assert.NoError(t, err)

	tx, err := erc20.Mint(receiver.Address(), big.NewInt(1000), &transactions.Options{})
	assert.NoError(t, err)
	receipt, err := tx.Wait(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, receipt)

	supply, err := erc20.TotalSupply()
	assert.NoError(t, err)

	assert.True(t, supply.Cmp(big.NewInt(1000)) == 0)
}

func TestFilter(t *testing.T) {
	txsender := txmanager.FromPK(solo.Keys()[0], client)
	receiver := txmanager.FromPK(solo.Keys()[1], client)

	_, erc20, err := deployErc20(client, txsender)
	assert.NoError(t, err)

	tx, err := erc20.Mint(receiver.Address(), big.NewInt(1000), &transactions.Options{})
	assert.NoError(t, err)
	receipt, err := tx.Wait(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, receipt)

	events, err := erc20.FilterTransfer(make([]testcontract.Erc20TransferCriteria, 0), nil)
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.Equal(t, common.Address{}, events[0].From)
	assert.Equal(t, receiver.Address(), events[0].To)
	assert.True(t, events[0].Value.Cmp(big.NewInt(1000)) == 0)
}

func TestWatch(t *testing.T) {
	txsender := txmanager.FromPK(solo.Keys()[0], client)
	receiver := txmanager.FromPK(solo.Keys()[1], client)

	_, erc20, err := deployErc20(client, txsender)
	assert.NoError(t, err)

	timeout, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	eventChan, err := erc20.WatchTransfer(make([]testcontract.Erc20TransferCriteria, 0), timeout, 10)
	assert.NoError(t, err)

	tx, err := erc20.Mint(receiver.Address(), big.NewInt(1000), &transactions.Options{})
	assert.NoError(t, err)
	receipt, err := tx.Wait(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, receipt)

	select {
	case event := <-eventChan:
		assert.Equal(t, common.Address{}, event.From)
		assert.Equal(t, receiver.Address(), event.To)
		assert.True(t, event.Value.Cmp(big.NewInt(1000)) == 0)
	case <-timeout.Done():
		assert.Fail(t, "timeout")
	}
}
