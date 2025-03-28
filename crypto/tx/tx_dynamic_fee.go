// Copyright (c) 2024 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package tx

import (
	"io"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type DynamicFeeTransaction struct {
	ChainTag             byte
	BlockRef             uint64
	Expiration           uint32
	Clauses              []*Clause
	Gas                  uint64
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	DependsOn            *common.Hash `rlp:"nil"`
	Nonce                uint64
	Reserved             reserved
	Signature            []byte
}

func (t *DynamicFeeTransaction) txType() byte {
	return TypeDynamic
}

func (t *DynamicFeeTransaction) copy() TxData {
	cpy := &DynamicFeeTransaction{
		ChainTag:             t.ChainTag,
		BlockRef:             t.BlockRef,
		Expiration:           t.Expiration,
		Clauses:              make([]*Clause, len(t.Clauses)),
		Gas:                  t.Gas,
		MaxFeePerGas:         new(big.Int),
		MaxPriorityFeePerGas: new(big.Int),
		DependsOn:            t.DependsOn,
		Nonce:                t.Nonce,
		Reserved:             t.Reserved,
		Signature:            t.Signature,
	}
	copy(cpy.Clauses, t.Clauses)
	if t.MaxFeePerGas != nil {
		cpy.MaxFeePerGas.Set(t.MaxFeePerGas)
	}
	if t.MaxPriorityFeePerGas != nil {
		cpy.MaxPriorityFeePerGas.Set(t.MaxPriorityFeePerGas)
	}
	return cpy
}

func (t *DynamicFeeTransaction) chainTag() byte {
	return t.ChainTag
}

func (t *DynamicFeeTransaction) blockRef() uint64 {
	return t.BlockRef
}

func (t *DynamicFeeTransaction) expiration() uint32 {
	return t.Expiration
}

func (t *DynamicFeeTransaction) clauses() []*Clause {
	return t.Clauses
}

func (t *DynamicFeeTransaction) gas() uint64 {
	return t.Gas
}

func (t *DynamicFeeTransaction) gasPriceCoef() uint8 {
	return 0
}

func (t *DynamicFeeTransaction) maxFeePerGas() *big.Int {
	return t.MaxFeePerGas
}

func (t *DynamicFeeTransaction) maxPriorityFeePerGas() *big.Int {
	return t.MaxPriorityFeePerGas
}

func (t *DynamicFeeTransaction) dependsOn() *common.Hash {
	return t.DependsOn
}

func (t *DynamicFeeTransaction) nonce() uint64 {
	return t.Nonce
}

func (t *DynamicFeeTransaction) reserved() reserved {
	return t.Reserved
}

func (t *DynamicFeeTransaction) signature() []byte {
	return t.Signature
}

func (t *DynamicFeeTransaction) setSignature(sig []byte) {
	t.Signature = sig
}

func (t *DynamicFeeTransaction) hashWithoutNonce(origin common.Address) *common.Hash {
	b := hash.Blake2bFn(func(w io.Writer) {
		rlp.Encode(w, []any{
			t.chainTag(),
			t.blockRef(),
			t.expiration(),
			t.clauses(),
			t.maxFeePerGas(),
			t.maxPriorityFeePerGas(),
			t.dependsOn(),
			t.nonce(),
			t.reserved(),
			origin,
		})
	})
	return &b
}

func (t *DynamicFeeTransaction) encode(w io.Writer) error {
	return rlp.Encode(w, []any{
		t.ChainTag,
		t.BlockRef,
		t.Expiration,
		t.Clauses,
		t.Gas,
		t.MaxFeePerGas,
		t.MaxPriorityFeePerGas,
		t.DependsOn,
		t.Nonce,
		&t.Reserved,
	})
}
