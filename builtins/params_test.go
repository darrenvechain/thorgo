package builtins

import (
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo"
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

func TestParams_Get(t *testing.T) {
	// init thor
	client, cancel := testcontainer.NewSolo()
	defer cancel()
	thor := thorgo.NewFromClient(client)

	// init contract
	solo1 := txmanager.FromPK(solo.Keys()[0], thor)
	params, err := NewParamsTransactor(thor, solo1)
	assert.NoError(t, err)

	res, err := params.Get(KeyExecutorAddress)
	assert.NoError(t, err)

	// check that solo 1 is the current executor
	assert.Equal(t, solo1.Address(), common.BytesToAddress(res.Bytes()))

	// set solo2 as the current executor
	solo2 := txmanager.FromPK(solo.Keys()[1], thor)
	var key [32]byte = KeyExecutorAddress
	var value = new(big.Int).SetBytes(solo2.Address().Bytes())
	tx, err := params.Set(key, value, &transactions.Options{})
	assert.NoError(t, err)
	_, err = tx.Wait()
	assert.NoError(t, err)

	// check that solo 2 is the updated executor
	res, err = params.Get(KeyExecutorAddress)
	assert.NoError(t, err)
	assert.Equal(t, solo2.Address(), common.BytesToAddress(res.Bytes()))
}
