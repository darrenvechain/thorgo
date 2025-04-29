package accounts

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
)

type TxManager interface {
	SendClauses(clauses []*tx.Clause, opts *transactions.Options) (*transactions.Visitor, error)
}

type ContractTransactor struct {
	*Contract
	manager TxManager
}

// Send a transaction to the contract with the given method and arguments.
func (c *ContractTransactor) Send(opts *transactions.Options, method string, args ...any) *Sender {
	return c.SendPayable(opts, big.NewInt(0), method, args...)
}

// SendPayable sends a transaction to the contract with the given method and arguments, using the specified VET amount.
func (c *ContractTransactor) SendPayable(opts *transactions.Options, vet *big.Int, method string, args ...any) *Sender {
	return newSender(c, opts, vet, method, args...)
}

type Sender struct {
	contract *ContractTransactor
	opts     *transactions.Options
	method   string
	args     []any
	vet      *big.Int
	sent     atomic.Bool
	visitor  atomic.Pointer[transactions.Visitor]
	mu       sync.Mutex
}

func newSender(contract *ContractTransactor, opts *transactions.Options, vet *big.Int, method string, args ...any) *Sender {
	return &Sender{
		contract: contract,
		opts:     opts,
		method:   method,
		args:     args,
		vet:      vet,
	}
}

func (s *Sender) Contract() *ContractTransactor {
	return s.contract
}

func (s *Sender) Send() (*transactions.Visitor, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.sent.Load() {
		return s.visitor.Load(), nil
	}
	if s.opts == nil {
		s.opts = &transactions.Options{}
	}
	clause, err := s.contract.AsClauseWithVET(s.vet, s.method, s.args...)
	if err != nil {
		return &transactions.Visitor{}, fmt.Errorf("failed to pack method %s: %w", s.method, err)
	}
	res, err := s.contract.manager.SendClauses([]*tx.Clause{clause}, s.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}
	s.visitor.Store(res)
	s.sent.Store(true)
	return res, nil
}

// Receipt waits for the transaction to be mined and returns the receipt.
func (s *Sender) Receipt(ctx context.Context) (*thorest.TransactionReceipt, error) {
	visitor, err := s.Send()
	if err != nil {
		return nil, err
	}
	return visitor.Wait(ctx)
}
