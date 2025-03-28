package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type FeesHistory struct {
	OldestBlock   common.Hash    `json:"oldestBlock"`
	BaseFeePerGas []*hexutil.Big `json:"baseFeePerGas"`
	GasUsedRatios []float64      `json:"gasUsedRatios"`
}

type FeesPriority struct {
	MaxPriorityFeePerGas *hexutil.Big `json:"maxPriorityFeePerGas"`
}
