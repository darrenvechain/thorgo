package thorest_test

import (
	"testing"

	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
)

var (
	thorClient *thorest.Client
	genesisID  common.Hash
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	genesis, err := thorClient.Block(thorest.RevisionNumber(0))
	if err != nil {
		panic(err)
	}
	genesisID = genesis.ID
	m.Run()
}
