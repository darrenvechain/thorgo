package accounts_test

import (
	"math/big"
	"testing"

	"github.com/darrenvechain/thorgo/client"
	"github.com/darrenvechain/thorgo/events"
	"github.com/darrenvechain/thorgo/txmanager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestContract_Call(t *testing.T) {
	// name
	var name []interface{}
	err := vthoRaw.Call("name", &name)
	assert.NoError(t, err)
	assert.Equal(t, "VeThor", name[0])

	// symbol
	var symbol []interface{}
	err = vthoRaw.Call("symbol", &symbol)
	assert.NoError(t, err)
	assert.Equal(t, "VTHO", symbol[0])

	// decimals
	var decimals []interface{}
	err = vthoRaw.Call("decimals", &decimals)
	assert.NoError(t, err)
	assert.Equal(t, uint8(18), decimals[0])
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
	receiver, err := txmanager.GeneratePK(thor)
	assert.NoError(t, err)

	// transfer clause
	clause, err := vthoRaw.AsClause("transfer", receiver.Address(), big.NewInt(1000))
	assert.NoError(t, err)
	assert.Equal(t, clause.Value(), big.NewInt(0))
	assert.Equal(t, clause.To().Hex(), vthoContract.Address().Hex())
}

func TestContract_Send(t *testing.T) {
	receiver, err := txmanager.GeneratePK(thor)
	assert.NoError(t, err)

	tx, err := vthoRaw.Send(account1, "transfer", receiver.Address(), big.NewInt(1000))
	assert.NoError(t, err)

	receipt, err := tx.Wait()
	assert.NoError(t, err)
	assert.False(t, receipt.Reverted)
}

func TestContract_EventCriteria(t *testing.T) {
	receiver, err := txmanager.GeneratePK(thor)
	assert.NoError(t, err)

	tx, err := vthoRaw.Send(account1, "transfer", receiver.Address(), big.NewInt(1000))
	assert.NoError(t, err)

	receipt, _ := tx.Wait()
	assert.False(t, receipt.Reverted)

	// event criteria - match the newly created receiver
	criteria, err := vthoRaw.EventCriteria("Transfer", nil, receiver.Address())
	assert.NoError(t, err)

	// fetch events
	transfers, err := events.New(thorClient, []client.EventCriteria{criteria}).Apply(0, 100)
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
