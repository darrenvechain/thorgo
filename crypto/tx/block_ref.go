// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package tx

import (
	"encoding/binary"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// BlockRef is block reference.
type BlockRef [8]byte

// Number extracts block number.
func (b BlockRef) Number() uint32 {
	return binary.BigEndian.Uint32(b[:])
}

// NewBlockRef create block reference with block number.
func NewBlockRef(blockNum uint32) (br BlockRef) {
	binary.BigEndian.PutUint32(br[:], blockNum)
	return
}

// NewBlockRefFromID create block reference from block id.
func NewBlockRefFromID(blockID common.Hash) (br BlockRef) {
	copy(br[:], blockID[:])
	return
}

func (b *BlockRef) UnmarshalJSON(data []byte) error {
	// block ref is returned as a hex string from the API
	encoded := string(data[1 : len(data)-1])
	encoded = strings.TrimPrefix(encoded, "0x")
	decoded := common.Hex2Bytes(encoded)
	copy(b[:], decoded)
	return nil
}

func (b *BlockRef) MarshalJSON() ([]byte, error) {
	hex := hexutil.Encode(b[:])
	return []byte(`"` + hex + `"`), nil
}
