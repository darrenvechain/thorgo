package logs

import "github.com/darrenvechain/thorgo/thorest"

type Options struct {
	rnge  *thorest.LogRange
	opts  *thorest.LogOptions
	order *string
}

// IncludeIndexes sets whether to include indexes in the log filter.
func (o *Options) IncludeIndexes(include bool) *Options {
	if o.opts == nil {
		o.opts = &thorest.LogOptions{}
	}
	o.opts.IncludeIndexes = &include
	return o
}

// RangeUnit sets the range unit for the log filter.
// The unit can be "block" or "time".
func (o *Options) RangeUnit(unit string) *Options {
	if o.rnge == nil {
		o.rnge = &thorest.LogRange{}
	}
	o.rnge.Unit = &unit
	return o
}

// Range sets the range for the log filter.
func (o *Options) Range(from, to int64) *Options {
	if o.rnge == nil {
		o.rnge = &thorest.LogRange{}
	}
	o.rnge.From = &from
	o.rnge.To = &to
	return o
}

// From sets the starting point for the log filter range.
func (o *Options) From(from int64) *Options {
	if o.rnge == nil {
		o.rnge = &thorest.LogRange{}
	}
	o.rnge.From = &from
	return o
}

// To sets the endpoint for the log filter range.
func (o *Options) To(to int64) *Options {
	if o.rnge == nil {
		o.rnge = &thorest.LogRange{}
	}
	o.rnge.To = &to
	return o
}

// Offset sets the offset for the log filter.
func (o *Options) Offset(offset int64) *Options {
	if o.opts == nil {
		o.opts = &thorest.LogOptions{}
	}
	o.opts.Offset = &offset
	return o
}

// Limit sets the limit for the log filter.
func (o *Options) Limit(limit int64) *Options {
	if o.opts == nil {
		o.opts = &thorest.LogOptions{}
	}
	o.opts.Limit = &limit
	return o
}

// Order sets the order of the log filter results.
// The order can be "asc" for ascending or "desc" for descending.
func (o *Options) Order(order string) *Options {
	o.order = &order
	return o
}

// Build finalizes the options and returns the log range, options, and order.
func (o *Options) Build() (*thorest.LogRange, *thorest.LogOptions, *string) {
	return o.rnge, o.opts, o.order
}
