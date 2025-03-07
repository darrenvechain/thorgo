package thorest

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type EventLog struct {
	Address *common.Address `json:"address,omitempty"`
	Topics  []common.Hash   `json:"topics"`
	Data    hexutil.Bytes   `json:"data"`
	Meta    *LogMeta        `json:"meta"`
}

type eventFilter struct {
	Range    *filterRange     `json:"range,omitempty"`
	Options  *filterOptions   `json:"options,omitempty"`
	Criteria *[]EventCriteria `json:"criteriaSet,omitempty"`
	Order    *string          `json:"order,omitempty"`
}

type EventCriteria struct {
	Address *common.Address `json:"address,omitempty"`
	Topic0  *common.Hash    `json:"topic0,omitempty"`
	Topic1  *common.Hash    `json:"topic1,omitempty"`
	Topic2  *common.Hash    `json:"topic2,omitempty"`
	Topic3  *common.Hash    `json:"topic3,omitempty"`
	Topic4  *common.Hash    `json:"topic4,omitempty"`
}

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
