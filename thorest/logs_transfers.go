package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// TransferLog represents a log entry for a transfer event in the blockchain.
type TransferLog struct {
	Sender    common.Address `json:"sender"`    // Sender is where the VET is sent from
	Recipient common.Address `json:"recipient"` // Recipient is the destination of the transaction
	Amount    *hexutil.Big   `json:"amount"`    // Amount is the amount of VET transferred
	Meta      *LogMeta       `json:"meta"`      // Meta contains metadata about the transfer log
}

type TransferCriteria struct {
	TxOrigin  *common.Address `json:"txOrigin,omitempty"`  // TxOrigin is the origin address of the transaction
	Sender    *common.Address `json:"sender,omitempty"`    // Sender is the address of the sender
	Recipient *common.Address `json:"recipient,omitempty"` // Recipient is the address of the recipient
}

type transferFilter struct {
	Range    *filterRange        `json:"range,omitempty"`
	Options  *filterOptions      `json:"options,omitempty"`
	Criteria *[]TransferCriteria `json:"criteriaSet,omitempty"`
	Order    *string             `json:"order,omitempty"`
}
