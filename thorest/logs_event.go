package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// EventLog represents a log entry in the blockchain.
type EventLog struct {
	Address *common.Address `json:"address,omitempty"` // Address is the contract address that generated the log
	Topics  []common.Hash   `json:"topics"`            // Topics are the indexed topics of the log
	Data    hexutil.Bytes   `json:"data"`              // Data is the non-indexed data of the log
	Meta    *LogMeta        `json:"meta"`              // Meta contains metadata about the log
}

type eventFilter struct {
	Range    *filterRange     `json:"range,omitempty"`
	Options  *filterOptions   `json:"options,omitempty"`
	Criteria *[]EventCriteria `json:"criteriaSet,omitempty"`
	Order    *string          `json:"order,omitempty"`
}

// EventCriteria is used to filter events based on their address and topics.
type EventCriteria struct {
	Address *common.Address `json:"address,omitempty"` // Address is the contract address to filter events
	Topic0  *common.Hash    `json:"topic0,omitempty"`  // Topic0 is the first indexed topic to filter events
	Topic1  *common.Hash    `json:"topic1,omitempty"`  // Topic1 is the second indexed topic to filter events
	Topic2  *common.Hash    `json:"topic2,omitempty"`  // Topic2 is the third indexed topic to filter events
	Topic3  *common.Hash    `json:"topic3,omitempty"`  // Topic3 is the fourth indexed topic to filter events
	Topic4  *common.Hash    `json:"topic4,omitempty"`  // Topic4 is the fifth indexed topic to filter events
}

// Matches checks if the event log matches the criteria.
func (e *EventCriteria) Matches(event *EventLog) bool {
	if e.Address != nil && event.Address != nil && e.Address.Cmp(*event.Address) != 0 {
		return false
	}

	matchTopic := func(topic *common.Hash, index int) bool {
		if topic == nil {
			return true // no criteria set, always match
		}
		if len(event.Topics) <= index {
			return false // not enough topics
		}
		return *topic == event.Topics[index] // compare topics
	}

	return matchTopic(e.Topic0, 0) &&
		matchTopic(e.Topic1, 1) &&
		matchTopic(e.Topic2, 2) &&
		matchTopic(e.Topic3, 3) &&
		matchTopic(e.Topic4, 4)
}
