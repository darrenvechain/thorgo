package builtins

import (
	"testing"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var thor = thorest.NewClientFromURL("https://mainnet.vechain.org")
var eventLimit = int64(10)

func TestEnergy_Name(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	name, err := energy.Name().Execute()
	assert.NoError(t, err)
	assert.Equal(t, "VeThor", name)
}

func TestEnergy_Symbol(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	symbol, err := energy.Symbol().Execute()
	assert.NoError(t, err)
	assert.Equal(t, "VTHO", symbol)
}

func TestEnergy_Decimals(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	decimals, err := energy.Decimals().Execute()
	assert.NoError(t, err)
	assert.Equal(t, uint8(18), decimals)
}

func TestEnergy_TotalSupply(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	totalSupply, err := energy.TotalSupply().Execute()
	assert.NoError(t, err)
	assert.NotZero(t, totalSupply)
}

func TestEnergy_BalanceOf(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	energyBal, err := energy.BalanceOf(energy.Address()).Execute()
	assert.NoError(t, err)
	assert.NotZero(t, energyBal)
}

func TestEnergy_FilterTransfer(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	events, err := energy.FilterTransfer(make([]VTHOTransferCriteria, 0)).Limit(eventLimit).Execute()
	assert.NoError(t, err)
	assert.Equal(t, eventLimit, int64(len(events)))
}

func TestEnergy_FilterApproval(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	events, err := energy.FilterApproval(make([]VTHOApprovalCriteria, 0)).Limit(eventLimit).Execute()
	assert.NoError(t, err)
	assert.Equal(t, eventLimit, int64(len(events)))
}

func TestEnergy_FilterTransfer_WithCriteria(t *testing.T) {
	energy, err := NewVTHO(thor)
	assert.NoError(t, err)

	to := common.HexToAddress("0xbF89016e670595AAa225eEfa0a84B7FB17b8dAC8")

	events, err := energy.FilterTransfer([]VTHOTransferCriteria{{To: &to}}).Limit(eventLimit).Execute()
	assert.NoError(t, err)
	assert.NotZero(t, len(events))
}
