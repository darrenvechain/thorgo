package builtins

import (
	"testing"

	"github.com/darrenvechain/thorgo"
	"github.com/stretchr/testify/assert"
)

var thor = thorgo.New("https://mainnet.vechain.org")

func TestEnergy(t *testing.T) {
	energy, err := NewEnergy(thor)
	assert.NoError(t, err)

	energyBal, err := energy.BalanceOf(energy.Address())
	assert.NoError(t, err)
	assert.NotZero(t, energyBal)
}
