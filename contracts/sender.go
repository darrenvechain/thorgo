package contracts

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type TxManager interface {
	SendClauses(clauses []*tx.Clause, opts *transactions.Options) (*transactions.Visitor, error)
}

type Sender struct {
	contract *Contract
	method   string
	args     []any
	vet      *big.Int
	opts     *transactions.Options
	mu       sync.Mutex
	visitor  atomic.Pointer[transactions.Visitor]
}

func NewSender(contract *Contract, method string, args ...any) *Sender {
	return &Sender{
		contract: contract,
		method:   method,
		args:     args,
		vet:      big.NewInt(0),
	}
}

// WithVET sets the VET amount to be sent with the transaction.
func (s *Sender) WithVET(vet *big.Int) *Sender {
	s.vet = vet
	return s
}

// Simulate the transaction without sending it.
func (s *Sender) Simulate(caller *common.Address) (*thorest.InspectResponse, error) {
	clause, err := s.Clause()
	if err != nil {
		return nil, fmt.Errorf("failed to pack method %s: %w", s.method, err)
	}
	request := thorest.InspectRequest{
		Clauses: []*tx.Clause{clause},
		Caller:  caller,
	}
	response, err := s.contract.client.Inspect(request)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect contract: %w", err)
	}
	if len(response) == 0 {
		return nil, fmt.Errorf("no response from inspection")
	}
	output := &response[0]
	if output.Reverted {
		reason, err := abi.UnpackRevert(output.Data)
		if err != nil {
			return output, fmt.Errorf("failed to unpack revert reason: %w", err)
		}
		if reason == "" {
			return output, fmt.Errorf("failed to unpack revert reason")
		}
		return output, fmt.Errorf("reverted: %s", reason)
	}
	if output.VmError != "" {
		return output, fmt.Errorf("VM error: %s", output.VmError)
	}
	return output, nil
}

// Clause returns a transaction clause that can be used to send a transaction to the contract.
func (s *Sender) Clause() (*tx.Clause, error) {
	return s.contract.Call(s.method, s.args...).WithValue(s.vet).Clause()
}

// WithOptions sets the options for the transaction.
func (s *Sender) WithOptions(opts *transactions.Options) *Sender {
	s.opts = opts
	return s
}

// Send the single clause transaction to the contract with the given method and arguments.
// If the transaction has already been sent, it will return the existing transaction visitor.
func (s *Sender) Send(manager TxManager) (*transactions.Visitor, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	visitor := s.visitor.Load()
	if visitor != nil {
		return visitor, nil
	}
	if s.opts == nil {
		s.opts = &transactions.Options{}
	}
	clause, err := s.Clause()
	if err != nil {
		return nil, err
	}
	res, err := manager.SendClauses([]*tx.Clause{clause}, s.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}
	s.visitor.Store(res)
	return res, nil
}

// Receipt sends the transaction, waits for it to be mined and returns the receipt.
// Later calls to this method will reuse the existing transaction if it has already been sent.
func (s *Sender) Receipt(ctx context.Context, manager TxManager) (*thorest.TransactionReceipt, error) {
	visitor, err := s.Send(manager)
	if err != nil {
		return nil, err
	}
	return visitor.Wait(ctx)
}
