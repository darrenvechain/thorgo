package thorest

import "github.com/ethereum/go-ethereum/common"

var unitBlock = "block"

type LogMeta struct {
	BlockID     common.Hash    `json:"blockID"`
	BlockNumber int64          `json:"blockNumber"`
	BlockTime   int64          `json:"blockTimestamp"`
	TxID        common.Hash    `json:"txID"`
	TxOrigin    common.Address `json:"txOrigin"`
	ClauseIndex int64          `json:"clauseIndex"`
}

type filterRange struct {
	Unit *string `json:"unit,omitempty"`
	From *int64  `json:"from,omitempty"`
	To   *int64  `json:"to,omitempty"`
}

type filterOptions struct {
	Offset *int64 `json:"offset,omitempty"`
	Limit  *int64 `json:"limit,omitempty"`
}

// LogFilters is used to filter transfer and event logs.
// Example:
//
//	filters := new(thorest.LogFilters).BlockRange(0, 1000).Limit(10).Offset(0).Order("desc")
type LogFilters struct {
	filterRange *filterRange
	options     *filterOptions
	order       *string
}

// Order sets the order of the logs, either "asc" or "desc".
func (lf *LogFilters) Order(order string) *LogFilters {
	lf.order = &order
	return lf
}

// Options sets the offset and limit of the logs.
func (lf *LogFilters) Options(offset, limit int64) *LogFilters {
	lf.options = &filterOptions{
		Offset: &offset,
		Limit:  &limit,
	}
	return lf
}

// Limit sets the limit of the number of logs to fetch.
func (lf *LogFilters) Limit(limit int64) *LogFilters {
	if lf.options == nil {
		lf.options = &filterOptions{}
	}
	lf.options.Limit = &limit
	return lf
}

// Offset sets the offset of the logs to fetch.
func (lf *LogFilters) Offset(offset int64) *LogFilters {
	if lf.options == nil {
		lf.options = &filterOptions{}
	}
	lf.options.Offset = &offset
	return lf
}

// Range sets the range of logs to filter, either from a block number or a timestamp.
func (lf *LogFilters) Range(from, to int64, unit string) *LogFilters {
	lf.filterRange = &filterRange{
		From: &from,
		To:   &to,
		Unit: &unit,
	}
	return lf
}

// BlockRange sets the range of blocks to filter logs from
func (lf *LogFilters) BlockRange(from, to int64) *LogFilters {
	lf.Range(from, to, unitBlock)
	return lf
}

// ToRange sets the maximum block number to filter logs from
func (lf *LogFilters) ToRange(to int64) *LogFilters {
	lf.filterRange = &filterRange{
		To:   &to,
		Unit: &unitBlock,
	}
	return lf
}

// FromRange sets the minimum block number to filter logs from
func (lf *LogFilters) FromRange(from int64) *LogFilters {
	lf.filterRange = &filterRange{
		From: &from,
		Unit: &unitBlock,
	}
	return lf
}
