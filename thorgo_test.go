package thorgo

import (
	"context"
	"testing"

	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var thor *Thor
var ctx = context.Background()

func TestMain(t *testing.M) {
	clt, cancel := testcontainer.NewSolo()
	defer cancel()
	thor = NewFromClient(ctx, clt)
	t.Run()
}

func TestFromClient(t *testing.T) {
	thor := NewFromClient(ctx, thor.Client())
	assert.NotNil(t, thor)
	tag, err := thor.Client().ChainTag()
	assert.NoError(t, err)
	assert.Equal(t, solo.ChainTag(), tag)
}

func TestBlock(t *testing.T) {
	block, err := thor.Blocks().ByNumber(0)
	assert.NoError(t, err)
	assert.NotNil(t, block)
}

func TestGetAccount(t *testing.T) {
	soloAccount := common.HexToAddress("0xf077b491b355E64048cE21E3A6Fc4751eEeA77fa")
	acc, err := thor.Account(soloAccount).Get()
	assert.NoError(t, err, "Account.httpGet should not return an error")
	assert.NotNil(t, acc, "Account.httpGet should return an account")

	assert.Greater(t, acc.Balance.ToInt().Uint64(), uint64(0))
	assert.Greater(t, acc.Energy.ToInt().Uint64(), uint64(0))
	assert.False(t, acc.HasCode)
}

func TestTransfers(t *testing.T) {
	criteria := make([]thorest.TransferCriteria, 0)

	transfers, err := thor.Transfers().Execute(criteria)

	assert.NoError(t, err)
	assert.NotNil(t, transfers)
}

func TestEvents(t *testing.T) {
	criteria := make([]thorest.EventCriteria, 0)

	events, err := thor.Events().Execute(criteria)

	assert.NoError(t, err)
	assert.NotNil(t, events)
}
