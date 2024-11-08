package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type TransferLog struct {
	Sender    common.Address `json:"sender"`
	Recipient common.Address `json:"recipient"`
	Amount    hexutil.Big    `json:"amount"`
	Meta      LogMeta        `json:"meta"`
}

type TransferCriteria struct {
	TxOrigin  *common.Address `json:"txOrigin,omitempty"`
	Sender    *common.Address `json:"sender,omitempty"`
	Recipient *common.Address `json:"recipient,omitempty"`
}

type transferFilter struct {
	Range    *filterRange        `json:"range,omitempty"`
	Options  *filterOptions      `json:"options,omitempty"`
	Criteria *[]TransferCriteria `json:"criteriaSet,omitempty"`
	Order    *string             `json:"order,omitempty"`
}
