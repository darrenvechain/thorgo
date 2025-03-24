// Copyright (c) 2024 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package tx

import (
	"bytes"
	"io"
	"sync"

	"github.com/darrenvechain/thorgo/crypto/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

// deriveBufferPool holds temporary encoder buffers for DeriveSha and TX encoding.
var encodeBufferPool = sync.Pool{
	New: func() any { return new(bytes.Buffer) },
}

func rlpHash(x any) common.Hash {
	return hash.Blake2bFn(func(w io.Writer) {
		rlp.Encode(w, x)
	})
}

// prefixedRlpHash writes the prefix into the hasher before rlp-encoding the
// given interface. It's used for typed transactions.
func prefixedRlpHash(prefix byte, x any) common.Hash {
	return hash.Blake2bFn(func(w io.Writer) {
		w.Write([]byte{prefix})
		rlp.Encode(w, x)
	})
}
