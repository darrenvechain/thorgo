package thorest

import "github.com/ethereum/go-ethereum/common"

// LogMeta represents the metadata of a log entry.
type LogMeta struct {
	BlockID     common.Hash    `json:"blockID"`
	BlockNumber int64          `json:"blockNumber"`
	BlockTime   int64          `json:"blockTimestamp"`
	TxID        common.Hash    `json:"txID"`
	TxOrigin    common.Address `json:"txOrigin"`
	ClauseIndex int64          `json:"clauseIndex"`
	TxIndex     *int64         `json:"txIndex,omitempty"`  // TxIndex is the index of the transaction in the block
	LogIndex    *int64         `json:"logIndex,omitempty"` // LogIndex is the index of the log in the transaction
}

type LogRange struct {
	Unit *string `json:"unit,omitempty"`
	From *int64  `json:"from,omitempty"`
	To   *int64  `json:"to,omitempty"`
}

type LogOptions struct {
	Offset         *int64 `json:"offset,omitempty"`
	Limit          *int64 `json:"limit,omitempty"`
	IncludeIndexes *bool  `json:"includeIndexes,omitempty"` // IncludeIndexes indicates whether to include transaction and log indexes in the response
}
