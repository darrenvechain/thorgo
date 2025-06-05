package contracts_test

import (
	"context"
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/builtins"
	"github.com/darrenvechain/thorgo/contracts"
	"github.com/darrenvechain/thorgo/internal/testcontainer"
	"github.com/darrenvechain/thorgo/solo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	thorClient   *thorest.Client
	vthoContract *builtins.VTHO
	vthoRaw      *contracts.Contract
	account1     *txmanager.PKManager
)

func TestMain(m *testing.M) {
	var cancel func()
	thorClient, cancel = testcontainer.NewSolo()
	defer cancel()
	vthoContract, _ = builtins.NewVTHO(thorClient)
	abi, _ := builtins.VTHOMetaData.GetAbi()
	vthoRaw = contracts.New(thorClient, vthoContract.Address(), abi)
	account1 = txmanager.FromPK(solo.Keys()[0], thorClient)
	m.Run()
}

func TestContract_Call(t *testing.T) {
	// name
	name, err := vthoRaw.Call("name").Execute()
	assert.NoError(t, err)
	assert.Equal(t, "VeThor", name)

	// symbol
	symbol, err := vthoRaw.Call("symbol").Execute()
	assert.NoError(t, err)
	assert.Equal(t, "VTHO", symbol)

	// decimals
	decimals, err := vthoRaw.Call("decimals").Execute()
	assert.NoError(t, err)
	assert.Equal(t, uint8(18), decimals)
}

func TestContract_DecodeCall(t *testing.T) {
	packed, err := vthoRaw.ABI.Pack("balanceOf", account1.Address())
	assert.NoError(t, err)

	balance := new(big.Int)
	err = vthoRaw.DecodeCall(packed, &balance)
	assert.NoError(t, err)
	assert.Greater(t, balance.Uint64(), uint64(0))
}

func TestContract_AsClause(t *testing.T) {
	receiver, err := txmanager.GeneratePK(thorClient)
	assert.NoError(t, err)

	// transfer clause
	clause, err := vthoRaw.AsClause("transfer", receiver.Address(), big.NewInt(1000))
	assert.NoError(t, err)
	assert.Equal(t, clause.Value(), big.NewInt(0))
	assert.Equal(t, clause.To().Hex(), vthoContract.Address().Hex())
}

func TestContract_Send(t *testing.T) {
	receiver, err := txmanager.GeneratePK(thorClient)
	assert.NoError(t, err)

	receipt, err := vthoRaw.
		Send("transfer", receiver.Address(), big.NewInt(1000)).
		Receipt(context.Background(), account1)
	assert.NoError(t, err)
	assert.False(t, receipt.Reverted)
}

func TestContract_EventCriteria(t *testing.T) {
	receiver, err := txmanager.GeneratePK(thorClient)
	assert.NoError(t, err)

	receipt, err := vthoRaw.Send("transfer", receiver.Address(), big.NewInt(1000)).
		Receipt(context.Background(), account1)
	assert.NoError(t, err)
	assert.False(t, receipt.Reverted)

	// event criteria - match the newly created receiver
	criteria, err := vthoRaw.EventCriteria("Transfer", nil, receiver.Address())
	assert.NoError(t, err)

	// fetch events
	transfers, err := thorClient.FilterEvents([]thorest.EventCriteria{criteria}, &thorest.LogFilters{})
	assert.NoError(t, err)

	// decode events
	decodedEvs, err := vthoRaw.DecodeEvents(transfers)
	assert.NoError(t, err)

	ev := decodedEvs[0]
	assert.Equal(t, "Transfer", ev.Name)
	assert.NotNil(t, ev.Args["_from"])
	assert.NotNil(t, ev.Args["_to"])
	assert.NotNil(t, ev.Args["_value"])
	assert.IsType(t, common.Address{}, ev.Args["_from"])
	assert.IsType(t, common.Address{}, ev.Args["_to"])
	assert.IsType(t, &big.Int{}, ev.Args["_value"])
}

func TestContract_UnpackLog(t *testing.T) {
	receiver, err := txmanager.GeneratePK(thorClient)
	assert.NoError(t, err)

	receipt, err := vthoRaw.Send("transfer", receiver.Address(), big.NewInt(1000)).
		Receipt(context.Background(), account1)
	assert.NoError(t, err)
	assert.False(t, receipt.Reverted)

	// event criteria - match the newly created receiver
	criteria, err := vthoRaw.EventCriteria("Transfer", nil, receiver.Address())
	assert.NoError(t, err)

	// fetch events
	transfers, err := thorClient.FilterEvents([]thorest.EventCriteria{criteria}, nil)
	assert.NoError(t, err)
	assert.Greater(t, len(transfers), 0)

	// unpack log
	type outStruct struct {
		From  common.Address
		To    common.Address
		Value *big.Int
	}
	res := outStruct{}
	err = vthoRaw.UnpackLog(&res, "Transfer", transfers[0])
	assert.NoError(t, err)
	assert.Equal(t, account1.Address(), res.From)
	assert.Equal(t, receiver.Address().Hex(), res.To.Hex())
	assert.Equal(t, big.NewInt(1000), res.Value)
}
