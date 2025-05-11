package thorest

import (
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Block represents a block in the VeChainThor blockchain.
type Block struct {
	Number       int64          `json:"number"`                  // Number is the block number
	ID           common.Hash    `json:"id"`                      // ID is the block ID
	Size         int64          `json:"size"`                    // Size is the size of the block in bytes.
	ParentID     common.Hash    `json:"parentID"`                // ParentID is the ID of the parent block.
	Timestamp    int64          `json:"timestamp"`               // Timestamp is the block timestamp in seconds since the epoch.
	GasLimit     int64          `json:"gasLimit"`                // GasLimit is the maximum amount of gas that can be used in the block.
	Beneficiary  common.Address `json:"beneficiary"`             // Beneficiary is the address that receives the block reward.
	GasUsed      int64          `json:"gasUsed"`                 // GasUsed is the total amount of gas used in the block.
	TotalScore   int64          `json:"totalScore"`              // TotalScore is the accumulated number of validators for each block.
	TxsRoot      common.Hash    `json:"txsRoot"`                 // TxsRoot is the root hash of the transactions in the block.
	TxsFeatures  int64          `json:"txsFeatures"`             // TxsFeatures represents the supported transaction features.
	StateRoot    common.Hash    `json:"stateRoot"`               // StateRoot is the root hash of the state trie after applying the block.
	ReceiptsRoot common.Hash    `json:"receiptsRoot"`            // ReceiptsRoot is the root hash of the receipts trie after applying the block.
	Com          bool           `json:"com"`                     // Com indicates whether the validators should vote on finality.
	Signer       common.Address `json:"signer"`                  // Signer is the address of the block signer.
	IsTrunk      bool           `json:"isTrunk"`                 // IsTrunk indicates whether the block is part of the main chain.
	IsFinalized  bool           `json:"isFinalized"`             // IsFinalized indicates whether the block is finalized.
	BaseFee      *hexutil.Big   `json:"baseFeePerGas,omitempty"` // BaseFee is present after the GALACTICA hardfork
	Transactions []common.Hash  `json:"transactions"`            // Transactions is the list of transaction IDs in the block.
}

// BlockRef returns a BlockRef object for the block.
func (b *Block) BlockRef() tx.BlockRef {
	return tx.NewBlockRefFromID(b.ID)
}

// BlockTransaction represents a transaction in a given block.
type BlockTransaction struct {
	ID                   common.Hash     `json:"id"`                             // ID is the transaction ID
	Type                 uint8           `json:"type,omitempty"`                 // Type is the transaction type
	ChainTag             byte            `json:"chainTag"`                       // ChainTag is the chain tag
	BlockRef             tx.BlockRef     `json:"blockRef"`                       // BlockRef is the block reference
	Expiration           int64           `json:"expiration"`                     // Expiration is the number of blocks on top of BlockRef, after which the transaction is considered expired
	Clauses              []tx.Clause     `json:"clauses"`                        // Clauses is the list of clauses in the transaction
	GasPriceCoef         int64           `json:"gasPriceCoef,omitempty"`         // GasPriceCoef is the coefficient used to calculate the gas price
	MaxFeePerGas         *hexutil.Big    `json:"maxFeePerGas,omitempty"`         // MaxFeePerGas is the maximum fee per gas that the sender is willing to pay
	MaxPriorityFeePerGas *hexutil.Big    `json:"maxPriorityFeePerGas,omitempty"` // MaxPriorityFeePerGas is the maximum priority fee per gas that the sender is willing to pay
	Gas                  int64           `json:"gas"`                            // Gas is the maximum amount of gas that can be used in the transaction
	Origin               common.Address  `json:"origin"`                         // Origin is the address of the sender
	Delegator            *common.Address `json:"delegator,omitempty"`            // Delegator is the address that paid for the transaction if it was delegated
	Nonce                *hexutil.Big    `json:"nonce"`                          // Nonce is the transaction nonce
	DependsOn            *common.Hash    `json:"dependsOn,omitempty"`            // DependsOn is the ID of the transaction that this transaction depends on
	Size                 int64           `json:"size"`                           // Size is the size of the transaction in bytes
	GasUsed              int64           `json:"gasUsed"`                        // GasUsed is the amount of gas used by the transaction
	GasPayer             common.Address  `json:"gasPayer"`                       // GasPayer is the address that pays for the gas
	Paid                 *hexutil.Big    `json:"paid"`                           // Paid is the amount of VTHO paid for the transaction
	Reward               *hexutil.Big    `json:"reward"`                         // Reward is the amount of VTHO given to the block producer
	Reverted             bool            `json:"reverted"`                       // Reverted indicates whether the transaction was reverted
	Outputs              []Output        `json:"outputs"`                        // Outputs is the list of Output for the transaction
}

type ExpandedBlock struct {
	Number       int64               `json:"number"`                  // Number is the block number
	ID           common.Hash         `json:"id"`                      // ID is the block ID
	Size         int64               `json:"size"`                    // Size is the size of the block in bytes.
	ParentID     common.Hash         `json:"parentID"`                // ParentID is the ID of the parent block.
	Timestamp    int64               `json:"timestamp"`               // Timestamp is the block timestamp in seconds since the epoch.
	GasLimit     int64               `json:"gasLimit"`                // GasLimit is the maximum amount of gas that can be used in the block.
	Beneficiary  common.Address      `json:"beneficiary"`             // Beneficiary is the address that receives the block reward.
	GasUsed      int64               `json:"gasUsed"`                 // GasUsed is the total amount of gas used in the block.
	TotalScore   int64               `json:"totalScore"`              // TotalScore is the accumulated number of validators for each block.
	TxsRoot      common.Hash         `json:"txsRoot"`                 // TxsRoot is the root hash of the transactions in the block.
	TxsFeatures  int64               `json:"txsFeatures"`             // TxsFeatures represents the supported transaction features.
	StateRoot    common.Hash         `json:"stateRoot"`               // StateRoot is the root hash of the state trie after applying the block.
	ReceiptsRoot common.Hash         `json:"receiptsRoot"`            // ReceiptsRoot is the root hash of the receipts trie after applying the block.
	Com          bool                `json:"com"`                     // Com indicates whether the validators should vote on finality.
	Signer       common.Address      `json:"signer"`                  // Signer is the address of the block signer.
	IsTrunk      bool                `json:"isTrunk"`                 // IsTrunk indicates whether the block is part of the main chain.
	IsFinalized  bool                `json:"isFinalized"`             // IsFinalized indicates whether the block is finalized.
	BaseFee      *hexutil.Big        `json:"baseFeePerGas,omitempty"` // BaseFee is present after the GALACTICA hardfork
	Transactions []*BlockTransaction `json:"transactions"`            // Transactions is the list of transactions and their receipts in the block.
}

// Events return all smart contract events in the block.
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

// FilteredEvents returns all smart contract events in the block that match the given criteria.
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
