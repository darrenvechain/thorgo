package txmanager

import (
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Delegator handles the payment of transaction fees
type Delegator interface {
	Delegate(tx *tx.Transaction, origin common.Address) ([]byte, error)
}

type DelegateRequest struct {
	Origin common.Address `json:"origin"`
	Raw    hexutil.Bytes  `json:"raw"`
}

type DelegateResponse struct {
	Signature hexutil.Bytes `json:"signature"`
}
