package thorest

import (
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Account struct {
	Balance *hexutil.Big `json:"balance"`
	Energy  *hexutil.Big `json:"energy"`
	HasCode bool         `json:"hasCode"`
}

type AccountCode struct {
	Code hexutil.Bytes `json:"code"`
}

type AccountStorage struct {
	Value common.Hash `json:"value"`
}

type InspectRequest struct {
	Gas        *uint64         `json:"gas,omitempty"`
	GasPrice   *uint64         `json:"gasPrice,omitempty"`
	Caller     *common.Address `json:"caller,omitempty"`
	ProvedWork *string         `json:"provedWork,omitempty"`
	GasPayer   *common.Address `json:"gasPayer,omitempty"`
	Expiration *uint64         `json:"expiration,omitempty"`
	BlockRef   *hexutil.Big    `json:"blockRef,omitempty"`
	Clauses    []*tx.Clause    `json:"clauses"`
}

type InspectResponse struct {
	Data      hexutil.Bytes `json:"data"`
	Events    []Event       `json:"events"`
	Transfers []Transfer    `json:"transfers"`
	GasUsed   uint64        `json:"gasUsed"`
	Reverted  bool          `json:"reverted"`
	VmError   string        `json:"vmError"`
}
