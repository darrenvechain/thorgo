package accounts_test

import (
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/accounts"
	"github.com/darrenvechain/thorgo/builtins"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient   *thorest.Client
	thor         *thorgo.Thor
	vthoContract *builtins.VTHO
	vthoRaw      *accounts.Contract
	account1     *txmanager.PKManager
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	thor = thorgo.NewFromClient(thorClient)
	vthoContract, _ = builtins.NewVTHO(thor)
	abi, _ := builtins.VTHOMetaData.GetAbi()
	vthoRaw = accounts.NewContract(thorClient, vthoContract.Address(), abi)
	account1 = txmanager.FromPK(solo.Keys()[0], thor)
	m.Run()
}

// TestGetAccount fetches a thor solo account and checks if the balance and energy are greater than 0
func TestVisitor_Get(t *testing.T) {
	acc, err := accounts.New(thorClient, account1.Address()).Get()

	assert.NoError(t, err, "Account.httpGet should not return an error")
	assert.NotNil(t, acc, "Account.httpGet should return an account")

	assert.Greater(t, acc.Balance.ToInt().Uint64(), uint64(0))
	assert.Greater(t, acc.Energy.ToInt().Uint64(), uint64(0))
	assert.False(t, acc.HasCode)
}

// TestGetAccountForRevision fetches a thor solo account for the genesis block
// and checks if the balance and energy are greater than 0
func TestVisitor_Get_At(t *testing.T) {
	acc, err := accounts.New(thorClient, account1.Address()).Revision(thorest.RevisionID(solo.GenesisID())).Get()

	assert.NoError(t, err, "Account.httpGet should not return an error")
	assert.NotNil(t, acc, "Account.httpGet should return an account")

	assert.Greater(t, acc.Balance.ToInt().Uint64(), uint64(0))
	assert.Greater(t, acc.Energy.ToInt().Uint64(), uint64(0))
	assert.False(t, acc.HasCode)
}

// TestGetCode fetches the code of the VTHO contract and checks if the code length is greater than 2 (0x)
func TestVisitor_Code(t *testing.T) {
	vtho, err := accounts.New(thorClient, vthoContract.Address()).Code()

	assert.NoError(t, err, "Account.Code should not return an error")
	assert.NotNil(t, vtho, "Account.Code should return a code")
	assert.Greater(t, len(vtho.Code), 2)
}

// TestGetCodeForRevision fetches the code of the VTHO contract for the genesis block
func TestVisitor_Code_At(t *testing.T) {
	vtho, err := accounts.New(thorClient, vthoContract.Address()).
		Revision(thorest.RevisionID(solo.GenesisID())).
		Code()

	assert.NoError(t, err, "Account.Code should not return an error")
	assert.NotNil(t, vtho, "Account.Code should return a code")
	assert.Greater(t, len(vtho.Code), 2)
}

// TestGetStorage fetches a storage position of the VTHO contract and checks if the value is empty
func TestVisitor_Storage(t *testing.T) {
	storage, err := accounts.New(thorClient, vthoContract.Address()).Storage(common.Hash{})

	assert.NoError(t, err, "Account.Storage should not return an error")
	assert.NotNil(t, storage, "Account.Storage should return a storage")
	assert.Equal(t, common.Hash{}, storage.Value)
}

// TestGetStorageForRevision fetches a storage position of the VTHO contract for the genesis block
func TestVisitor_Storage_At(t *testing.T) {
	storage, err := accounts.New(thorClient, vthoContract.Address()).
		Revision(thorest.RevisionID(solo.GenesisID())).
		Storage(common.Hash{})

	assert.NoError(t, err, "Account.Storage should not return an error")
	assert.NotNil(t, storage, "Account.Storage should return a storage")
	assert.Equal(t, common.Hash{}, storage.Value)
}

func TestVisitor_Call(t *testing.T) {
	clause, err := vthoRaw.AsClause("balanceOf", account1.Address())
	assert.NoError(t, err)

	res, err := accounts.New(thorClient, vthoContract.Address()).Call(clause.Data())
	assert.NoError(t, err)
	balance := new(big.Int).SetBytes(res.Data)
	assert.Greater(t, balance.Uint64(), big.NewInt(0).Uint64())
}

func TestVisitor_Call_At(t *testing.T) {
	clause, err := vthoRaw.AsClause("balanceOf", account1.Address())
	assert.NoError(t, err)

	res, err := accounts.New(thorClient, vthoContract.Address()).
		Revision(thorest.RevisionID(solo.GenesisID())).
		Call(clause.Data())
	assert.NoError(t, err)
	balance := new(big.Int).SetBytes(res.Data)
	assert.Greater(t, balance.Uint64(), big.NewInt(0).Uint64())
}
