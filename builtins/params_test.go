package builtins

import (
	"context"
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/thorest"

	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	KeyExecutorAddress = common.BytesToHash([]byte("executor"))
)

func TestParams_Set(t *testing.T) {
	// init thor
	client, cancel := testcontainer.NewSolo()
	defer cancel()

	// init contract
	solo1 := txmanager.FromPK(solo.Keys()[0], client)
	params, err := NewParamsTransactor(client, solo1)
	assert.NoError(t, err)

	res, err := params.Get(KeyExecutorAddress, thorest.RevisionBest())
	assert.NoError(t, err)

	// check that solo 1 is the current executor
	assert.Equal(t, solo1.Address(), common.BytesToAddress(res.Bytes()))

	// set solo2 as the current executor
	solo2 := txmanager.FromPK(solo.Keys()[1], client)
	// TODO: Not sure why the value is a big.Int here
	value := new(big.Int).SetBytes(solo2.Address().Bytes())
	tx, err := params.Set(KeyExecutorAddress, value, &transactions.Options{})
	assert.NoError(t, err)
	_, err = tx.Wait(context.Background())
	assert.NoError(t, err)

	// check that solo 2 is the updated executor
	res, err = params.Get(KeyExecutorAddress, thorest.RevisionBest())
	assert.NoError(t, err)
	assert.Equal(t, solo2.Address(), common.BytesToAddress(res.Bytes()))
}

func TestParams_MBP(t *testing.T) {
	client := thorest.NewClientFromURL("https://testnet.vechain.org")
	params, err := NewParams(client)
	assert.NoError(t, err)

	mpbKey := common.BytesToHash([]byte("max-block-proposers"))

	res, err := params.Get(mpbKey, thorest.RevisionBest())
	assert.NoError(t, err)
	assert.Equal(t, int64(14), res.Int64())
}
