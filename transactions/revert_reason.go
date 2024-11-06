package transactions

import (
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type RevertReason struct {
	res *thorest.TxRevertResponse
}

// Data returns the raw revert reason data.
func (r *RevertReason) Data() []byte {
	return r.res.Output
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
	unpacked, err := abiErr.Inputs.Unpack(r.res.Output)
	if err != nil {
		return err
	}
	return abiErr.Inputs.Copy(value, unpacked)
}
