package transfers

import (
	"errors"

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
	request *api.TransferFilter
}

func New(c *api.Client, criteria []api.TransferCriteria) *Filter {
	return &Filter{client: c, request: &api.TransferFilter{
		Criteria: &criteria,
	}}
}

// Desc sets the order of the transfers to descending. Default is ascending.
func (f *Filter) Desc() *Filter {
	f.request.Order = &descending
	return f
}

// Asc sets the order of the transfers to ascending. This is the default.
func (f *Filter) Asc() *Filter {
	f.request.Order = &ascending
	return f
}

// BlockRange sets the block range for the transfer filter.
func (f *Filter) BlockRange(from int64, to int64) *Filter {
	f.request.Range = &api.FilterRange{
		From: &from,
		To:   &to,
		Unit: &block,
	}
	return f
}

// TimeRange sets the time range for the transfer filter.
func (f *Filter) TimeRange(from int64, to int64) *Filter {
	f.request.Range = &api.FilterRange{
		From: &from,
		To:   &to,
		Unit: &time,
	}
	return f
}

// Apply sends the transfer filter to the node and returns the results.
func (f *Filter) Apply(offset int64, limit int64) ([]api.TransferLog, error) {
	if limit > 256 {
		return nil, errors.New("limit must be less than or equal to 256")
	}
	f.request.Options = &api.FilterOptions{
		Offset: &offset,
		Limit:  &limit,
	}

	return f.client.FilterTransfers(f.request)
}
