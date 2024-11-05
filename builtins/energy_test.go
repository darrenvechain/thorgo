package builtins

import (
	"testing"

	"github.com/darrenvechain/thorgo"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var thor = thorgo.New("https://mainnet.vechain.org")

func TestEnergy_Name(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	name, err := energy.Name()
	assert.NoError(t, err)
	assert.Equal(t, "VeThor", name)
}

func TestEnergy_Symbol(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	symbol, err := energy.Symbol()
	assert.NoError(t, err)
	assert.Equal(t, "VTHO", symbol)
}

func TestEnergy_Decimals(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	decimals, err := energy.Decimals()
	assert.NoError(t, err)
	assert.Equal(t, uint8(18), decimals)
}

func TestEnergy_TotalSupply(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	totalSupply, err := energy.TotalSupply()
	assert.NoError(t, err)
	assert.NotZero(t, totalSupply)
}

func TestEnergy_BalanceOf(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	energyBal, err := energy.BalanceOf(energy.Address())
	assert.NoError(t, err)
	assert.NotZero(t, energyBal)
}

func TestEnergy_FilterTransfer(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	maxEvents := int64(10)
	events, err := energy.FilterTransfer(make([]EnergyTransferCriteria, 0), &thorest.FilterOptions{Limit: &maxEvents}, nil)
	assert.NoError(t, err)
	assert.Equal(t, maxEvents, int64(len(events)))
}

func TestEnergy_FilterApproval(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	maxEvents := int64(10)
	events, err := energy.FilterApproval(make([]EnergyApprovalCriteria, 0), &thorest.FilterOptions{Limit: &maxEvents}, nil)
	assert.NoError(t, err)
	assert.Equal(t, maxEvents, int64(len(events)))
}

func TestEnergy_FilterTransfer_WithCriteria(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	to := common.HexToAddress("0xbF89016e670595AAa225eEfa0a84B7FB17b8dAC8")

	maxEvents := int64(10)
	events, err := energy.FilterTransfer([]EnergyTransferCriteria{{To: &to}}, &thorest.FilterOptions{Limit: &maxEvents}, nil)
	assert.NoError(t, err)
	assert.NotZero(t, len(events))
}
