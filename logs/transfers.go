package logs

import "github.com/darrenvechain/thorgo/thorest"

type TransfersFilterer struct {
	opts     *Options
	criteria []thorest.TransferCriteria
	client   *thorest.Client
}

func NewTransfersFilterer(client *thorest.Client) *TransfersFilterer {
	return &TransfersFilterer{
		opts:   &Options{},
		client: client,
	}
}

func (t *TransfersFilterer) Criteria(criteria ...thorest.TransferCriteria) *TransfersFilterer {
	if t.criteria == nil {
		t.criteria = make([]thorest.TransferCriteria, 0)
	}
	t.criteria = append(t.criteria, criteria...)
	return t
}

func (t *TransfersFilterer) Execute() ([]*thorest.TransferLog, error) {
	return t.client.FilterTransfers(&thorest.TransferFilter{
		Criteria: &t.criteria,
		Range:    t.opts.rnge,
		Options:  t.opts.opts,
		Order:    t.opts.order,
	})
}

// RangeUnit sets the range unit for the log filter.
// The unit can be "block" or "time".
func (t *TransfersFilterer) RangeUnit(unit string) *TransfersFilterer {
	t.opts.RangeUnit(unit)
	return t
}

// Range sets the range for the log filter.
func (t *TransfersFilterer) Range(from, to int64) *TransfersFilterer {
	t.opts.Range(from, to)
	return t
}

// From sets the starting point for the log filter range.
func (t *TransfersFilterer) From(from int64) *TransfersFilterer {
	t.opts.From(from)
	return t
}

// To sets the endpoint for the log filter range.
func (t *TransfersFilterer) To(to int64) *TransfersFilterer {
	t.opts.To(to)
	return t
}

// Offset sets the offset for the log filter.
func (t *TransfersFilterer) Offset(offset int64) *TransfersFilterer {
	t.opts.Offset(offset)
	return t
}

// Limit sets the limit for the log filter.
func (t *TransfersFilterer) Limit(limit int64) *TransfersFilterer {
	t.opts.Limit(limit)
	return t
}

// Order sets the order of the log filter results.
// The order can be "asc" for ascending or "desc" for descending.
func (t *TransfersFilterer) Order(order string) *TransfersFilterer {
	t.opts.Order(order)
	return t
}
