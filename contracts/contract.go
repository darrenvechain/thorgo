package contracts

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Contract is a generic smart contract wrapper for VeChainThor.
type Contract struct {
	client  *thorest.Client
	ABI     *abi.ABI
	Address common.Address
}

// New creates a new instance of the Contract struct.
func New(
	client *thorest.Client,
	address common.Address,
	abi *abi.ABI,
) *Contract {
	return &Contract{client: client, Address: address, ABI: abi}
}

// Call creates a contract caller for the specified method and arguments.
func (c *Contract) Call(method string, args ...any) *Caller {
	return NewCaller(c, method, args...)
}

// Send creates a contract transaction sender for the specified method and arguments.
func (c *Contract) Send(method string, args ...any) *Sender {
	return NewSender(c, method, args...)
}

// Filter creates a contract event filterer for the specified event name.
func (c *Contract) Filter(eventName string) *Filterer {
	return NewFilterer(c, eventName)
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

// UnpackLog unpacks a retrieved log into the provided output structure.
// For example:
//
//		type TransferEvent struct {
//			From   common.Address
//			To     common.Address
//			Value  *big.Int
//		}
//
//	    contract.UnpackLog(&transferEvent, "Transfer", log)
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
