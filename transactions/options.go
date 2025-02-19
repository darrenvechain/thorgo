package transactions

import (
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
)

type Options struct {
	Nonce                *uint64
	GasPayer             *common.Address
	Delegation           *bool
	Gas                  *uint64
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	Expiration           *uint32
	BlockRef             *tx.BlockRef
	DependsOn            *common.Hash
	ChainTag             *byte
	VET                  *big.Int
}

type OptionsBuilder struct {
	options Options
}

// Nonce sets the nonce for the transaction.
func (b *OptionsBuilder) Nonce(nonce uint64) *OptionsBuilder {
	b.options.Nonce = &nonce
	return b
}

// GasPayer enables the delegation feature and sets the gas payer when simulating the transaction.
func (b *OptionsBuilder) GasPayer(payer common.Address) *OptionsBuilder {
	b.options.GasPayer = &payer
	return b
}

// Delegated enables the delegation feature. Use `GasPayer` for a more accurate simulation.
func (b *OptionsBuilder) Delegated() *OptionsBuilder {
	delegation := true
	b.options.Delegation = &delegation
	return b
}

// Gas sets the gas provision for the transaction. If not set, it will be estimated.
func (b *OptionsBuilder) Gas(gas uint64) *OptionsBuilder {
	b.options.Gas = &gas
	return b
}

// MaxFeePerGas sets the maximum fee per gas for the transaction.
func (b *OptionsBuilder) MaxFeePerGas(fee *big.Int) *OptionsBuilder {
	b.options.MaxFeePerGas = fee
	return b
}

// MaxPriorityFeePerGas sets the maximum priority fee per gas for the transaction.
func (b *OptionsBuilder) MaxPriorityFeePerGas(fee *big.Int) *OptionsBuilder {
	b.options.MaxPriorityFeePerGas = fee
	return b
}

// Expiration sets the expiration block count. Defaults to 30 blocks (5 minutes) if not set.
func (b *OptionsBuilder) Expiration(exp uint32) *OptionsBuilder {
	b.options.Expiration = &exp
	return b
}

// BlockRef sets the block reference. Defaults to the "best" block reference if not set.
func (b *OptionsBuilder) BlockRef(br tx.BlockRef) *OptionsBuilder {
	b.options.BlockRef = &br
	return b
}

// DependsOn sets the transaction that this transaction depends on.
func (b *OptionsBuilder) DependsOn(txID common.Hash) *OptionsBuilder {
	b.options.DependsOn = &txID
	return b
}

// ChainTag sets the chain tag for the transaction. Defaults to the chain tag of the genesis block if not set.
func (b *OptionsBuilder) ChainTag(tag byte) *OptionsBuilder {
	b.options.ChainTag = &tag
	return b
}

// VET is an additional option when sending a single clause transaction.
func (b *OptionsBuilder) VET(vet *big.Int) *OptionsBuilder {
	b.options.VET = vet
	return b
}

func (b *OptionsBuilder) Build() *Options {
	return &b.options
}
