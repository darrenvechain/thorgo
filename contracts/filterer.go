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
	opts      *logs.Options
	contract  *Contract
	eventName string
	criteria  []*EventCriteria
}

// NewFilterer creates a new Filterer instance for the given contract.
func NewFilterer(contract *Contract, eventName string) *Filterer {
	return &Filterer{
		opts:      &logs.Options{},
		contract:  contract,
		eventName: eventName,
	}
}

// Criteria adds event criteria to the filterer.
func (f *Filterer) Criteria(criteria *EventCriteria) *Filterer {
	if f.criteria == nil {
		f.criteria = make([]*EventCriteria, 0)
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

	rng, opts, order := f.opts.Build()

	filter := &thorest.EventFilter{
		Range:    rng,
		Options:  opts,
		Order:    order,
		Criteria: &thorestCriteria,
	}

	return f.contract.client.FilterEvents(filter)
}

// Event represents a decoded event from a contract log.
type Event struct {
	Name string
	Args map[string]any
	Log  *thorest.EventLog
}

// ExecuteAndDecode executes the filter and decodes the events into the provided value type.
func (f *Filterer) ExecuteAndDecode() ([]Event, error) {
	logs, err := f.Execute()
	if err != nil {
		return nil, err
	}

	var decoded []Event
	for _, log := range logs {
		eventABI, err := f.contract.ABI.EventByID(log.Topics[0])
		if err != nil {
			continue
		}

		var indexed abi.Arguments
		for _, arg := range eventABI.Inputs {
			if arg.Indexed {
				indexed = append(indexed, arg)
			}
		}

		values := make(map[string]any)
		err = abi.ParseTopicsIntoMap(values, indexed, log.Topics[1:])
		if err != nil {
			return nil, err
		}

		err = eventABI.Inputs.UnpackIntoMap(values, log.Data)
		if err != nil {
			return nil, err
		}

		decoded = append(decoded, Event{
			Name: eventABI.Name,
			Args: values,
			Log:  log,
		})
	}
	return decoded, nil
}

func (f *Filterer) IncludeIndexes(include bool) *Filterer {
	f.opts.IncludeIndexes(include)
	return f
}

func (f *Filterer) RangeUnit(unit string) *Filterer {
	f.opts.RangeUnit(unit)
	return f
}

func (f *Filterer) Range(from, to int64) *Filterer {
	f.opts.Range(from, to)
	return f
}

func (f *Filterer) From(from int64) *Filterer {
	f.opts.From(from)
	return f
}

func (f *Filterer) To(to int64) *Filterer {
	f.opts.To(to)
	return f
}

func (f *Filterer) Offset(offset int64) *Filterer {
	f.opts.Offset(offset)
	return f
}

func (f *Filterer) Limit(limit int64) *Filterer {
	f.opts.Limit(limit)
	return f
}

func (f *Filterer) Order(order string) *Filterer {
	f.opts.Order(order)
	return f
}

// makeTopicHash converts a matcher value to a topic hash
func (f *Filterer) makeTopicHash(matcher any) (common.Hash, error) {
	topics, err := abi.MakeTopics([]any{matcher})
	if err != nil {
		return common.Hash{}, err
	}
	if len(topics) == 0 || len(topics[0]) == 0 {
		return common.Hash{}, errors.New("failed to generate topic hash")
	}
	return topics[0][0], nil
}
