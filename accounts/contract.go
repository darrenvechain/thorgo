package accounts

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Contract is a generic representation of a smart contract.
type Contract struct {
	client  *thorest.Client
	ABI     *abi.ABI
	Address common.Address
}

// NewContract creates a new contract instance.
func NewContract(
	client *thorest.Client,
	address common.Address,
	abi *abi.ABI,
) *Contract {
	return &Contract{client: client, Address: address, ABI: abi}
}

// Call executes a read-only contract call.
func (c *Contract) Call(method string, results *[]any, args ...any) error {
	return c.CallAt(thorest.RevisionNext(), method, results, args...)
}

// CallAt executes a read-only contract call at a specific revision.
func (c *Contract) CallAt(revision thorest.Revision, method string, results *[]any, args ...any) error {
	if results == nil {
		results = new([]any)
	}
	packed, err := c.ABI.Pack(method, args...)
	if err != nil {
		return fmt.Errorf("failed to pack method %s: %w", method, err)
	}
	clause := tx.NewClause(&c.Address).WithData(packed).WithValue(big.NewInt(0))
	request := thorest.InspectRequest{
		Clauses: []*tx.Clause{clause},
	}
	response, err := c.client.InspectAt(request, revision)
	if err != nil {
		return fmt.Errorf("failed to inspect contract: %w", err)
	}
	inspection := response[0]
	if inspection.Reverted {
		return errors.New("contract call reverted")
	}
	if inspection.VmError != "" {
		return errors.New(inspection.VmError)
	}
	if len(*results) == 0 {
		res, err := c.ABI.Unpack(method, inspection.Data)
		*results = res
		return err
	}
	res := *results
	return c.ABI.UnpackIntoInterface(res[0], method, inspection.Data)
}

// DecodeCall decodes the result of a contract call, for example, decoding a clause's 'data'.
// The data must include the method signature.
func (c *Contract) DecodeCall(data []byte, value any) error {
	var method string
	for name, m := range c.ABI.Methods {
		if len(data) >= 4 && bytes.Equal(data[:4], m.ID) {
			method = name
			break
		}
	}

	if method == "" {
		return errors.New("method signature not found")
	}

	data = data[4:]

	err := c.ABI.UnpackIntoInterface(value, method, data)
	if err != nil {
		return fmt.Errorf("failed to unpack method %s: %w", method, err)
	}
	return nil
}

// AsClause returns a transaction clause for the given method and arguments.
func (c *Contract) AsClause(method string, args ...any) (*tx.Clause, error) {
	packed, err := c.ABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack method %s: %w", method, err)
	}
	return tx.NewClause(&c.Address).WithData(packed).WithValue(big.NewInt(0)), nil
}

// AsClauseWithVET returns a transaction clause for the given method, value, and arguments.
func (c *Contract) AsClauseWithVET(vet *big.Int, method string, args ...any) (*tx.Clause, error) {
	packed, err := c.ABI.Pack(method, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack method %s: %w", method, err)
	}
	return tx.NewClause(&c.Address).WithData(packed).WithValue(vet), nil
}

func (c *Contract) Transactor(manager TxManager) *ContractTransactor {
	return &ContractTransactor{Contract: c, manager: manager}
}

// EventCriteria generates criteria to query contract events by name.
// Matchers correspond to event input parameters and must be in the same order as the event's inputs.
// Use nil for any event input you want to ignore.
//
// For example, consider the following event:
//
//	event Transfer(address indexed from, address indexed to, uint256 value);
//
// To filter events based on the 'to' address while ignoring the 'from' address and 'value', you can pass nil for those values:
//
//	to := common.HexToAddress("0x87AA2B76f29583E4A9095DBb6029A9C41994E25B")
//	criteria, err := contract.EventCriteria("Transfer", nil, &to)
//
// Returns an EventCriteria object and any error encountered.
func (c *Contract) EventCriteria(name string, matchers ...any) (thorest.EventCriteria, error) {
	ev, ok := c.ABI.Events[name]
	if !ok {
		return thorest.EventCriteria{}, fmt.Errorf("event %s not found", name)
	}
	criteria := thorest.EventCriteria{
		Address: &c.Address,
		Topic0:  &ev.ID,
	}

	for i := range ev.Inputs {
		if i >= len(matchers) {
			break
		}
		if matchers[i] == nil {
			continue
		}
		if !ev.Inputs[i].Indexed {
			return thorest.EventCriteria{}, errors.New("can't match non-indexed event inputs")
		}
		topics, err := abi.MakeTopics(
			[]any{matchers[i]},
		)
		if err != nil {
			return thorest.EventCriteria{}, err
		}

		switch i + 1 {
		case 1:
			criteria.Topic1 = &topics[0][0]
		case 2:
			criteria.Topic2 = &topics[0][0]
		case 3:
			criteria.Topic3 = &topics[0][0]
		case 4:
			criteria.Topic4 = &topics[0][0]
		}
	}

	return criteria, nil
}

type Event struct {
	Name string
	Args map[string]any
	Log  *thorest.EventLog
}

// DecodeEvents parses logs into a slice of decoded events.
// The logs are typically obtained from a contract's filtered events.
//
// For example, consider the Solidity event:
//
//	event Transfer(address indexed from, address indexed to, uint256 value);
//
// To retrieve and decode "Transfer" events where the 'to' address matches a given value, you would:
//
//	to := common.HexToAddress("0x87AA2B76f29583E4A9095DBb6029A9C41994E25B")
//	logs, _ := client.FilterEvents(contract.EventCriteria("Transfer", nil, &to))
//	events, _ := contract.DecodeEvents(logs)
//
// Once decoded, you can iterate over the events and access their name and arguments:
//
//	for _, event := range events {
//	  fmt.Println(event.Name, event.Args)
//	}
//
// This function returns a slice of decoded event objects and any error encountered.
func (c *Contract) DecodeEvents(logs []*thorest.EventLog) ([]Event, error) {
	var decoded []Event
	for _, log := range logs {
		if len(log.Topics) < 2 {
			continue
		}

		eventABI, err := c.ABI.EventByID(log.Topics[0])
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

// UnpackLog unpacks a retrieved log into the provided output structure.
func (c *Contract) UnpackLog(out any, event string, log *thorest.EventLog) error {
	if len(log.Topics) == 0 {
		return errors.New("anonymous events are not supported")
	}
	if log.Topics[0] != c.ABI.Events[event].ID {
		return errors.New("event signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := c.ABI.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.ABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
