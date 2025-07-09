package contracts

import (
	"errors"

	"github.com/darrenvechain/thorgo/logs"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// EventCriteria represents the criteria for filtering events with raw matchers
type EventCriteria struct {
	Topic1 any // Raw matcher value, will be converted to hash during Execute
	Topic2 any
	Topic3 any
	Topic4 any
}

// Filterer provides a convenient way to filter and retrieve contract events based on specific criteria.
type Filterer struct {
	*logs.Options
	contract  *Contract
	eventName string
	criteria  []EventCriteria
}

// NewFilterer creates a new Filterer instance for the given contract.
func NewFilterer(contract *Contract, eventName string) *Filterer {
	return &Filterer{
		Options:   &logs.Options{},
		contract:  contract,
		eventName: eventName,
	}
}

// AddCriteria adds event criteria to the filterer.
func (f *Filterer) AddCriteria(criteria EventCriteria) *Filterer {
	if f.criteria == nil {
		f.criteria = make([]EventCriteria, 0)
	}
	f.criteria = append(f.criteria, criteria)
	return f
}

// Execute the raw query to filter events based on the specified criteria.
func (f *Filterer) Execute() ([]*thorest.EventLog, error) {
	event, ok := f.contract.ABI.Events[f.eventName]
	if !ok {
		return nil, errors.New("event not found in contract ABI: " + f.eventName)
	}

	var thorestCriteria []thorest.EventCriteria

	if len(f.criteria) == 0 {
		// If no criteria specified, create a default one for this event
		thorestCriteria = []thorest.EventCriteria{
			{
				Address: &f.contract.Address,
				Topic0:  &event.ID,
			},
		}
	} else {
		// Convert our criteria to thorest criteria, processing matchers
		thorestCriteria = make([]thorest.EventCriteria, len(f.criteria))
		for i, c := range f.criteria {
			tc := thorest.EventCriteria{
				Address: &f.contract.Address,
				Topic0:  &event.ID,
			}
			tc.Address = &f.contract.Address
			tc.Topic0 = &event.ID

			// Process topic matchers
			if c.Topic1 != nil {
				hash, err := f.makeTopicHash(c.Topic1)
				if err != nil {
					return nil, err
				}
				tc.Topic1 = &hash
			}
			if c.Topic2 != nil {
				hash, err := f.makeTopicHash(c.Topic2)
				if err != nil {
					return nil, err
				}
				tc.Topic2 = &hash
			}
			if c.Topic3 != nil {
				hash, err := f.makeTopicHash(c.Topic3)
				if err != nil {
					return nil, err
				}
				tc.Topic3 = &hash
			}
			if c.Topic4 != nil {
				hash, err := f.makeTopicHash(c.Topic4)
				if err != nil {
					return nil, err
				}
				tc.Topic4 = &hash
			}

			thorestCriteria[i] = tc
		}
	}

	rng, opts, order := f.Options.Build()

	filter := &thorest.EventFilter{
		Range:    rng,
		Options:  opts,
		Order:    order,
		Criteria: &thorestCriteria,
	}

	return f.contract.client.FilterEvents(filter)
}

// makeTopicHash converts a matcher value to a topic hash
func (f *Filterer) makeTopicHash(matcher any) (common.Hash, error) {
	topics, err := abi.MakeTopics([]interface{}{matcher})
	if err != nil {
		return common.Hash{}, err
	}
	if len(topics) == 0 || len(topics[0]) == 0 {
		return common.Hash{}, errors.New("failed to generate topic hash")
	}
	return topics[0][0], nil
}
