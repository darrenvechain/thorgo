package thorest

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestEventCriteria_Matches(t *testing.T) {
	hexToHash := func(s string) *common.Hash {
		hash := common.HexToHash(s)
		return &hash
	}

	tests := []struct {
		name     string
		criteria EventCriteria
		event    *EventLog
		expected bool
	}{
		{
			name:     "Match on Address only",
			criteria: EventCriteria{Address: &common.Address{0x1}},
			event:    &EventLog{Address: &common.Address{0x1}, Topics: nil},
			expected: true,
		},
		{
			name:     "Address mismatch",
			criteria: EventCriteria{Address: &common.Address{0x1}},
			event:    &EventLog{Address: &common.Address{0x2}, Topics: nil},
			expected: false,
		},
		{
			name: "Match on Topic0 only",
			criteria: EventCriteria{
				Topic0: hexToHash("0xabc123"),
			},
			event: &EventLog{
				Topics: []common.Hash{
					common.HexToHash("0xabc123"),
				},
			},
			expected: true,
		},
		{
			name: "Topic0 mismatch",
			criteria: EventCriteria{
				Topic0: hexToHash("0xabc123"),
			},
			event: &EventLog{
				Topics: []common.Hash{
					common.HexToHash("0xdef456"),
				},
			},
			expected: false,
		},
		{
			name: "Topic index out of bounds",
			criteria: EventCriteria{
				Topic1: hexToHash("0xabc123"),
			},
			event: &EventLog{
				Topics: []common.Hash{
					common.HexToHash("0xabc123"),
				},
			},
			expected: false,
		},
		{
			name: "Match on multiple topics",
			criteria: EventCriteria{
				Topic0: hexToHash("0xabc123"),
				Topic1: hexToHash("0xdef456"),
			},
			event: &EventLog{
				Topics: []common.Hash{
					common.HexToHash("0xabc123"),
					common.HexToHash("0xdef456"),
				},
			},
			expected: true,
		},
		{
			name: "Partial match on multiple topics",
			criteria: EventCriteria{
				Topic0: hexToHash("0xabc123"),
				Topic1: hexToHash("0xdef456"),
				Topic2: hexToHash("0xfed789"),
			},
			event: &EventLog{
				Topics: []common.Hash{
					common.HexToHash("0xabc123"),
					common.HexToHash("0xdef456"),
				},
			},
			expected: false,
		},
		{
			name: "No criteria set (matches anything)",
			criteria: EventCriteria{
				Address: nil,
				Topic0:  nil,
				Topic1:  nil,
			},
			event: &EventLog{
				Address: &common.Address{0x1},
				Topics: []common.Hash{
					common.HexToHash("0xabc123"),
				},
			},
			expected: true,
		},
		{
			name: "Empty topics in event with non-nil criteria",
			criteria: EventCriteria{
				Topic0: hexToHash("0xabc123"),
			},
			event: &EventLog{
				Topics: []common.Hash{},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.criteria.Matches(tt.event)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
