package logs

import "github.com/darrenvechain/thorgo/thorest"

type EventFilterer struct {
	opts     *Options
	criteria []thorest.EventCriteria
	client   *thorest.Client
}

func NewEventsFilterer(client *thorest.Client) *EventFilterer {
	return &EventFilterer{
		opts:   &Options{},
		client: client,
	}
}

func (e *EventFilterer) Criteria(criteria thorest.EventCriteria) *EventFilterer {
	if e.criteria == nil {
		e.criteria = make([]thorest.EventCriteria, 0)
	}
	e.criteria = append(e.criteria, criteria)
	return e
}

func (e *EventFilterer) Execute() ([]*thorest.EventLog, error) {
	return e.client.FilterEvents(&thorest.EventFilter{
		Criteria: &e.criteria,
		Range:    e.opts.rnge,
		Options:  e.opts.opts,
		Order:    e.opts.order,
	})
}

// RangeUnit sets the range unit for the log filter.
// The unit can be "block" or "time".
func (e *EventFilterer) RangeUnit(unit string) *EventFilterer {
	e.opts.RangeUnit(unit)
	return e
}

// Range sets the range for the log filter.
func (e *EventFilterer) Range(from, to int64) *EventFilterer {
	e.opts.Range(from, to)
	return e
}

// From sets the starting point for the log filter range.
func (e *EventFilterer) From(from int64) *EventFilterer {
	e.opts.From(from)
	return e
}

// To sets the endpoint for the log filter range.
func (e *EventFilterer) To(to int64) *EventFilterer {
	e.opts.To(to)
	return e
}

func (e *EventFilterer) IncludeIndexes(include bool) *EventFilterer {
	if e.opts.opts == nil {
		e.opts.opts = &thorest.LogOptions{}
	}
	e.opts.opts.IncludeIndexes = &include
	return e
}

// Offset sets the offset for the log filter.
func (e *EventFilterer) Offset(offset int64) *EventFilterer {
	e.opts.Offset(offset)
	return e
}

// Limit sets the limit for the log filter.
func (e *EventFilterer) Limit(limit int64) *EventFilterer {
	e.opts.Limit(limit)
	return e
}

// Order sets the order of the log filter results.
// The order can be "asc" for ascending or "desc" for descending.
func (e *EventFilterer) Order(order string) *EventFilterer {
	e.opts.Order(order)
	return e
}
