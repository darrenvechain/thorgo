package contracts

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
)

type ParserFunc[T any] func(data []any) (T, error)

type Caller[T any] struct {
	contract *Contract

	method string
	args   []any

	rev    thorest.Revision
	value  *big.Int
	parser ParserFunc[T]
}

func NewCaller[T any](contract *Contract, method string, args ...any) *Caller[T] {
	return &Caller[T]{
		contract: contract,
		method:   method,
		args:     args,
		rev:      thorest.RevisionBest(),
	}
}

func (c *Caller[T]) WithRevision(rev thorest.Revision) *Caller[T] {
	c.rev = rev
	return c
}

func (c *Caller[T]) WithValue(value *big.Int) *Caller[T] {
	c.value = value
	return c
}

func (c *Caller[T]) WithParser(parser ParserFunc[T]) *Caller[T] {
	c.parser = parser
	return c
}

func (c *Caller[any]) Call() (*thorest.InspectResponse, error) {
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

func (c *Caller[T]) Execute() (T, error) {
	response, err := c.Call()
	var zero T
	if err != nil {
		return zero, err
	}
	if len(response.Data) == 0 {
		return zero, errors.New("no data returned from contract call")
	}
	res, err := c.contract.ABI.Unpack(c.method, response.Data)
	if err != nil {
		return zero, err
	}
	if len(res) == 1 {
		if result, ok := res[0].(T); ok {
			return result, nil
		}
		return zero, fmt.Errorf("unexpected type returned: %T", res[0])
	}
	if c.parser == nil {
		return zero, fmt.Errorf("parser function must be defined for multiple/struct return values, got %d values", len(res))
	}
	return c.parser(res)
}
