package thorgo

import (
	"testing"

	"github.com/darrenvechain/thorgo/client"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var thor *Thor

func TestMain(t *testing.M) {
	clt, cancel := testcontainer.NewSolo()
	defer cancel()
	thor = NewFromClient(clt)
	t.Run()
}

func TestBadURL(t *testing.T) {
	_, err := NewFromURL("http://localhost:80")
	assert.Error(t, err)
}

func TestFromClient(t *testing.T) {
	thor := NewFromClient(thor.Client)
	assert.NotNil(t, thor)
	assert.Equal(t, solo.ChainTag(), thor.Client.ChainTag())
}

func TestBlock(t *testing.T) {
	block, err := thor.Blocks.ByNumber(0)
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
	// account 1
	account1 := solo.Keys()[0]
	account1Addr := crypto.PubkeyToAddress(account1.PublicKey)

	criteria := make([]client.TransferCriteria, 0)
	criteria = append(criteria, client.TransferCriteria{
		Sender: &account1Addr,
	})

	events, err := thor.Transfers(criteria).
		BlockRange(0, 10000).
		Asc().
		Apply(0, 100)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}

func TestEvents(t *testing.T) {
	criteria := make([]client.EventCriteria, 0)

	events, err := thor.Events(criteria).
		BlockRange(0, 10000).
		Asc().
		Apply(0, 100)
	assert.NoError(t, err)
	assert.NotNil(t, events)
}
