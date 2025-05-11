package txmanager

import (
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Signer is an interface for signing transactions
type Signer interface {
	// SignTransaction signs the transaction and returns the signature
	SignTransaction(tx *tx.Transaction) ([]byte, error)
	// Address returns the origin of the transaction to be signed
	Address() common.Address
}

// Delegator handles the payment of transaction fees
type Delegator interface {
	// Delegate returns the delegator signature for the transaction
	Delegate(tx *tx.Transaction, origin common.Address) ([]byte, error)
}

// DelegateRequest is the request structure for HTTP requests according to VIP-201
// See: https://github.com/vechain/VIPs/blob/master/vips/VIP-201.md
type DelegateRequest struct {
	Origin common.Address `json:"origin"`
	Raw    hexutil.Bytes  `json:"raw"`
}

// DelegateResponse is the response structure for HTTP requests according to VIP-201
// See: https://github.com/vechain/VIPs/blob/master/vips/VIP-201.md#explicit-grant-flow
type DelegateResponse struct {
	Signature hexutil.Bytes `json:"signature"`
}
