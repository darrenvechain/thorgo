package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// FeesHistory represents historical fee data, including the gas-used ratio per block
type FeesHistory struct {
	OldestBlock   common.Hash      `json:"oldestBlock"`      // OldestBlock is the hash of the oldest block in the history
	BaseFeePerGas []*hexutil.Big   `json:"baseFeePerGas"`    // BaseFeePerGas is the base fee per gas for each block in the history
	GasUsedRatios []float64        `json:"gasUsedRatios"`    // GasUsedRatios is the gas-used ratio for each block in the history
	Reward        [][]*hexutil.Big `json:"reward,omitempty"` // Reward is the reward for each block in the history
}

// FeesPriority is an estimation of the priority fee for a transaction to be included in a block
type FeesPriority struct {
	MaxPriorityFeePerGas *hexutil.Big `json:"maxPriorityFeePerGas"` // MaxPriorityFeePerGas is the maximum priority fee per gas
}
