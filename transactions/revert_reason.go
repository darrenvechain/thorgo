package transactions

import (
	"bytes"
	"errors"

	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type RevertReason struct {
	res *thorest.TxRevertResponse
}

// Response returns the raw revert reason response.
func (r *RevertReason) Response() *thorest.TxRevertResponse {
	return r.res
}

// Decode will decode the revert reason if it is `Error(string)` or `Panic(uint256)`
func (r *RevertReason) Decode() (string, error) {
	return abi.UnpackRevert(r.res.Output)
}

// IsKnownSelector will return true if the revert reason is a known selector.
func (r *RevertReason) IsKnownSelector() bool {
	_, err := abi.UnpackRevert(r.res.Output)
	return err == nil
}

// DecodeInto will decode
func (r *RevertReason) DecodeInto(abiErr abi.Error, value interface{}) error {
	if len(r.res.Output) < 4 || !bytes.Equal(r.res.Output[:4], abiErr.ID.Bytes()[0:4]) {
		return errors.New("selector does not match")
	}

	unpacked, err := abiErr.Inputs.Unpack(r.res.Output[4:])
	if err != nil {
		return err
	}
	return abiErr.Inputs.Copy(value, unpacked)
}
