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

// IsDefaultSelector will return true if the revert data is a `Error(string)` or `Panic(uint256)` selector.
func (r *RevertReason) IsDefaultSelector() bool {
	_, err := abi.UnpackRevert(r.res.Output)
	return err == nil
}

// MatchesABI will return true if the revert reason matches the ABI error.
func (r *RevertReason) MatchesABI(abiErr abi.Error) bool {
	return len(r.res.Output) >= 4 && bytes.Equal(r.res.Output[:4], abiErr.ID.Bytes()[0:4])
}

// DecodeInto will decode
func (r *RevertReason) DecodeInto(abiErr abi.Error, value any) error {
	if !r.MatchesABI(abiErr) {
		return errors.New("revert reason does not match ABI error")
	}
	unpacked, err := abiErr.Inputs.Unpack(r.res.Output[4:])
	if err != nil {
		return err
	}
	return abiErr.Inputs.Copy(value, unpacked)
}
