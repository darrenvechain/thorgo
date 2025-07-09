package contracts

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
)

type ParserFunc[T any] func(data []any) (T, error)

type Caller struct {
	contract *Contract
	method   string
	args     []any
	rev      thorest.Revision
	value    *big.Int
}

func NewCaller(contract *Contract, method string, args ...any) *Caller {
	return &Caller{
		contract: contract,
		method:   method,
		args:     args,
		rev:      thorest.RevisionBest(),
	}
}

func (c *Caller) WithRevision(rev thorest.Revision) *Caller {
	c.rev = rev
	return c
}

func (c *Caller) WithValue(value *big.Int) *Caller {
	c.value = value
	return c
}

// Call executes the contract call and returns the raw response
func (c *Caller) Call() (*thorest.InspectResponse, error) {
	packed, err := c.contract.ABI.Pack(c.method, c.args...)
	if err != nil {
		return nil, errors.New("failed to pack method: " + err.Error())
	}
	clause := tx.NewClause(&c.contract.Address).WithData(packed)
	if c.value != nil {
		clause = clause.WithValue(c.value)
	}
	request := thorest.InspectRequest{
		Clauses: []*tx.Clause{clause},
	}
	response, err := c.contract.client.InspectAt(request, c.rev)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect contract: %w", err)
	}
	if len(response) != 1 {
		return nil, errors.New("unexpected number of responses from inspect")
	}
	output := response[0]
	if output.Reverted {
		return &output, errors.New("contract call reverted")
	}
	return &output, nil
}

// UnpackResult unpacks the response data using the ABI
func (c *Caller) UnpackResult(response *thorest.InspectResponse) ([]any, error) {
	if len(response.Data) == 0 {
		return nil, errors.New("no data returned from contract call")
	}
	return c.contract.ABI.Unpack(c.method, response.Data)
}

// Execute calls the contract and returns the unpacked result
func (c *Caller) Execute() ([]any, error) {
	response, err := c.Call()
	if err != nil {
		return nil, err
	}
	return c.UnpackResult(response)
}
