package contracts

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
)

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

// WithRevision returns a new Caller with the given revision set.
func (c *Caller) WithRevision(rev thorest.Revision) *Caller {
	return &Caller{
		contract: c.contract,
		method:   c.method,
		args:     c.args,
		rev:      rev,
		value:    c.value,
	}
}

// WithValue returns a new Caller with the given value set.
func (c *Caller) WithValue(value *big.Int) *Caller {
	return &Caller{
		contract: c.contract,
		method:   c.method,
		args:     c.args,
		rev:      c.rev,
		value:    value,
	}
}

func (c *Caller) Clause() (*tx.Clause, error) {
	packed, err := c.contract.ABI.Pack(c.method, c.args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack method: %w", err)
	}
	clause := tx.NewClause(&c.contract.Address).WithData(packed)
	if c.value != nil {
		clause = clause.WithValue(c.value)
	}
	return clause, nil
}

// Call executes the contract call and returns the raw response
func (c *Caller) Call() (*thorest.InspectResponse, error) {
	clause, err := c.Clause()
	if err != nil {
		return nil, err
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
