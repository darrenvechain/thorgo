package thorest_test

import (
	"strings"
	"testing"

	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestClient_Account(t *testing.T) {
	acc, err := thorClient.Account(common.HexToAddress("0xd1d37b8913563fC25BC5bB2E669eB3dBC6b87762"))
	assert.NoError(t, err)

	assert.Zero(t, acc.Balance.ToInt().Int64())
	assert.Zero(t, acc.Energy.ToInt().Int64())
	assert.False(t, acc.HasCode)
}

func TestClient_AccountAt(t *testing.T) {
	acc, err := thorClient.AccountAt(
		common.HexToAddress("0xd1d37b8913563fC25BC5bB2E669eB3dBC6b87762"),
		thorest.RevisionID(solo.GenesisID()),
	)

	assert.NoError(t, err)
	assert.Equal(t, int64(0), acc.Balance.ToInt().Int64())
	assert.Equal(t, int64(0), acc.Energy.ToInt().Int64())
	assert.False(t, acc.HasCode)
}

func TestClient_AccountCode(t *testing.T) {
	res, err := thorClient.AccountCode(common.HexToAddress("0x0000000000000000000000000000456E65726779"))
	assert.NoError(t, err)
	assert.Greater(t, len(res.Code), 2)
}

func TestClient_AccountCodeAt(t *testing.T) {
	res, err := thorClient.AccountCodeAt(
		common.HexToAddress("0x0000000000000000000000000000456E65726779"),
		thorest.RevisionID(solo.GenesisID()),
	)
	assert.NoError(t, err)
	assert.Greater(t, len(res.Code), 2)
}

func TestClient_AccountStorage(t *testing.T) {
	res, err := thorClient.AccountStorage(
		common.HexToAddress("0x0000000000000000000000000000456E65726779"),
		common.HexToHash(strings.Repeat("0", 64)),
	)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Value), 32)
}

func TestClient_AccountStorageAt(t *testing.T) {
	res, err := thorClient.AccountStorageAt(
		common.HexToAddress("0x0000000000000000000000000000456E65726779"),
		common.HexToHash(strings.Repeat("0", 64)),
		thorest.RevisionID(solo.GenesisID()),
	)

	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(res.Value), 32)
}
