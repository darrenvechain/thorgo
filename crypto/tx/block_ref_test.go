// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package tx

import (
	"crypto/rand"
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

var (
	blockID     = common.HexToHash("01354d00bb0fc43ad2ce97ef3124f2e2d6b637184b4085d4cacad54b7ff443c3")
	blockNumber = uint32(20270336)
)

type Tx struct {
	BlockRef BlockRef `json:"blockRef"`
}

func TestBlockRef(t *testing.T) {
	assert.Equal(t, uint32(0), BlockRef{}.Number())
	assert.Equal(t, BlockRef{0, 0, 0, 0xff, 0, 0, 0, 0}, NewBlockRef(0xff))

	var bid common.Hash
	rand.Read(bid[:]) // nolint

	br := NewBlockRefFromID(bid)
	assert.Equal(t, bid[:8], br[:])
}

func TestNewBlockRef(t *testing.T) {
	br := NewBlockRef(blockNumber)
	br2 := NewBlockRefFromID(blockID)

	assert.Equal(t, br.Number(), br2.Number())
}

func TestBlockRef_MarshalJSON(t *testing.T) {
	bytes, err := json.Marshal(&Tx{BlockRef: NewBlockRefFromID(blockID)})
	assert.NoError(t, err)
	assert.Equal(t, `{"blockRef":"0x01354d00bb0fc43a"}`, string(bytes))
}

func TestBlockRef_UnmarshalJSON(t *testing.T) {
	tx := Tx{}
	err := json.Unmarshal([]byte(`{"blockRef":"0x01354d00bb0fc43a"}`), &tx)
	assert.NoError(t, err)
	assert.Equal(t, 20270336, int(tx.BlockRef.Number()))
}
