package logs

import "github.com/darrenvechain/thorgo/thorest"

type EventFilterer struct {
	*Options
	criteria []thorest.EventCriteria
	client   *thorest.Client
}

func NewEventsFilterer(client *thorest.Client) *EventFilterer {
	return &EventFilterer{
		Options: &Options{},
		client:  client,
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
		Range:    e.rnge,
		Options:  e.opts,
		Order:    e.order,
	})
}
