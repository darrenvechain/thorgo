package datagen

import (
	"crypto/rand"
	mathrand "math/rand/v2"

	"github.com/ethereum/go-ethereum/common"
)

func RandBytes(n int) []byte {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return bytes
}

func RandAddress() (addr common.Address) {
	rand.Read(addr[:])
	return
}

func RandomHash() common.Hash {
	var b32 common.Hash

	rand.Read(b32[:])
	return b32
}

func RandInt() int {
	return mathrand.Int() //#nosec G404
}

func RandUint64() uint64 {
	return mathrand.Uint64() //#nosec G404
}

func RandIntN(n int) int {
	return mathrand.N(n) //#nosec G404
}
