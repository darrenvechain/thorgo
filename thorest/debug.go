package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type debugTraceClause struct {
	Target string      `json:"target"`
	Name   string      `json:"name"`
	Config interface{} `json:"config"`
}

type TxRevertResponse struct {
	From    common.Address `json:"from"`
	Gas     hexutil.Big    `json:"gas"`
	GasUsed hexutil.Big    `json:"gasUsed"`
	To      common.Address `json:"to"`
	Input   hexutil.Bytes  `json:"input"`
	Output  hexutil.Bytes  `json:"output"`
	Error   string         `json:"error"`
	Value   hexutil.Big    `json:"value"`
	Type    string         `json:"type"`
}
