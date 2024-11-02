package events

import (
	"github.com/darrenvechain/thorgo/api"
)

var (
	descending = "desc"
	ascending  = "asc"
	time       = "time"
	block      = "block"
)

type Filter struct {
	client  *api.Client
	request *api.EventFilter
}

func New(c *api.Client, criteria []api.EventCriteria) *Filter {
	return &Filter{client: c, request: &api.EventFilter{
		Criteria: &criteria,
	}}
}

// Desc sets the order of the events to descending. Default is ascending.
func (f *Filter) Desc() *Filter {
	f.request.Order = &descending
	return f
}

// Asc sets the order of the events to ascending. This is the default.
func (f *Filter) Asc() *Filter {
	f.request.Order = &ascending
	return f
}

// BlockRange sets the range of blocks to filter events.
func (f *Filter) BlockRange(from int64, to int64) *Filter {
	f.request.Range = &api.FilterRange{
		From: &from,
		To:   &to,
		Unit: &block,
	}
	return f
}

// TimeRange sets the range of time to filter events.
func (f *Filter) TimeRange(from int64, to int64) *Filter {
	f.request.Range = &api.FilterRange{
		From: &from,
		To:   &to,
		Unit: &time,
	}
	return f
}

// Apply executes the filter and returns the events.
func (f *Filter) Apply(offset int64, limit int64) ([]api.EventLog, error) {
	f.request.Options = &api.FilterOptions{
		Offset: &offset,
		Limit:  &limit,
	}

	return f.client.FilterEvents(f.request)
}
