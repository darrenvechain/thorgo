package thorest

import (
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type Block struct {
	Number       int64          `json:"number"`
	ID           common.Hash    `json:"id"`
	Size         int64          `json:"size"`
	ParentID     common.Hash    `json:"parentID"`
	Timestamp    int64          `json:"timestamp"`
	GasLimit     int64          `json:"gasLimit"`
	Beneficiary  common.Address `json:"beneficiary"`
	GasUsed      int64          `json:"gasUsed"`
	TotalScore   int64          `json:"totalScore"`
	TxsRoot      common.Hash    `json:"txsRoot"`
	TxsFeatures  int64          `json:"txsFeatures"`
	StateRoot    common.Hash    `json:"stateRoot"`
	ReceiptsRoot common.Hash    `json:"receiptsRoot"`
	Com          bool           `json:"com"`
	Signer       common.Address `json:"signer"`
	IsTrunk      bool           `json:"isTrunk"`
	IsFinalized  bool           `json:"isFinalized"`
	BaseFee      *hexutil.Big   `json:"baseFee"`
	Transactions []common.Hash  `json:"transactions"`
}

func (b *Block) ChainTag() byte {
	return b.ID[len(b.ID)-1]
}

func (b *Block) BlockRef() tx.BlockRef {
	return tx.NewBlockRefFromID(b.ID)
}

type BlockTransaction struct {
	ID                   common.Hash     `json:"id"`
	Type                 string          `json:"txType,omitempty"`
	ChainTag             byte            `json:"chainTag"`
	BlockRef             tx.BlockRef     `json:"blockRef"`
	Expiration           int64           `json:"expiration"`
	Clauses              []tx.Clause     `json:"clauses"`
	GasPriceCoef         int64           `json:"gasPriceCoef,omitempty"`
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas,omitempty"`
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas,omitempty"`
	Gas                  int64           `json:"gas"`
	Origin               common.Address  `json:"origin"`
	Delegator            *common.Address `json:"delegator,omitempty"`
	Nonce                *hexutil.Big    `json:"nonce"`
	DependsOn            *common.Hash    `json:"dependsOn,omitempty"`
	Size                 int64           `json:"size"`
	GasUsed              int64           `json:"gasUsed"`
	GasPayer             common.Address  `json:"gasPayer"`
	Paid                 *hexutil.Big    `json:"paid"`
	Reward               *hexutil.Big    `json:"reward"`
	Reverted             bool            `json:"reverted"`
	Outputs              []Output        `json:"outputs"`
}

type ExpandedBlock struct {
	Number       int64               `json:"number"`
	ID           common.Hash         `json:"id"`
	Size         int64               `json:"size"`
	ParentID     common.Hash         `json:"parentID"`
	Timestamp    int64               `json:"timestamp"`
	GasLimit     int64               `json:"gasLimit"`
	Beneficiary  common.Address      `json:"beneficiary"`
	GasUsed      int64               `json:"gasUsed"`
	TotalScore   int64               `json:"totalScore"`
	TxsRoot      common.Hash         `json:"txsRoot"`
	TxsFeatures  int64               `json:"txsFeatures"`
	StateRoot    common.Hash         `json:"stateRoot"`
	ReceiptsRoot common.Hash         `json:"receiptsRoot"`
	Com          bool                `json:"com"`
	Signer       common.Address      `json:"signer"`
	IsTrunk      bool                `json:"isTrunk"`
	IsFinalized  bool                `json:"isFinalized"`
	BaseFee      *hexutil.Big        `json:"baseFee"`
	Transactions []*BlockTransaction `json:"transactions"`
}

func (b *ExpandedBlock) Events() []*EventLog {
	events := make([]*EventLog, 0)
	for _, transaction := range b.Transactions {
		for i, output := range transaction.Outputs {
			for _, ev := range output.Events {
				ev := &EventLog{
					Address: ev.Address,
					Topics:  ev.Topics,
					Data:    ev.Data,
					Meta: &LogMeta{
						BlockID:     b.ID,
						BlockNumber: b.Number,
						BlockTime:   b.Timestamp,
						TxID:        transaction.ID,
						TxOrigin:    transaction.Origin,
						ClauseIndex: int64(i),
					},
				}
				events = append(events, ev)
			}
		}
	}
	return events
}

func (b *ExpandedBlock) FilteredEvents(criteriaSet []EventCriteria) []*EventLog {
	relevant := make([]*EventLog, 0)

	for _, ev := range b.Events() {
		matches := len(criteriaSet) == 0
		for _, c := range criteriaSet {
			if c.Matches(ev) {
				matches = true
				break
			}
		}
		if matches {
			relevant = append(relevant, ev)
		}
	}

	return relevant
}
