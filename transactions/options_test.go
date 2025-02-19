package transactions

import (
	"testing"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestOptionsBuilder_Build(t *testing.T) {
	b := new(OptionsBuilder).
		Nonce(100).
		GasPayer(common.Address{1}).
		Gas(1_000_000).
		Expiration(100).
		BlockRef(tx.NewBlockRef(100))

	o := b.Build()

	assert.Equal(t, uint64(100), *o.Nonce)
	assert.Equal(t, common.Address{1}, *o.GasPayer)
	assert.Equal(t, uint64(1_000_000), *o.Gas)
	assert.Equal(t, uint32(100), *o.Expiration)
}
