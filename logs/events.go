package logs

import "github.com/darrenvechain/thorgo/thorest"

type EventQuerier struct {
	*Options
	client *thorest.Client
}

func NewEventQuerier(client *thorest.Client) *EventQuerier {
	return &EventQuerier{
		Options: &Options{},
		client:  client,
	}
}

func (eq *EventQuerier) Execute(criteria []thorest.EventCriteria) ([]*thorest.EventLog, error) {
	return eq.client.FilterEvents(&thorest.EventFilter{
		Criteria: &criteria,
		Range:    eq.rnge,
		Options:  eq.opts,
		Order:    eq.order,
	})
}
