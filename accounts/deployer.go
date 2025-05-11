package accounts

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Deployer is a struct that provides methods to deploy smart contracts on VeChainThor.
type Deployer struct {
	client   *thorest.Client
	bytecode []byte
	abi      *abi.ABI
	value    *big.Int
}

// NewDeployer creates a new instance of the Deployer struct.
func NewDeployer(client *thorest.Client, bytecode []byte, abi *abi.ABI) *Deployer {
	return &Deployer{client: client, bytecode: bytecode, abi: abi, value: big.NewInt(0)}
}

// Deploy a smart contract to the VeChainThor blockchain.
func (d *Deployer) Deploy(ctx context.Context, sender TxManager, opts *transactions.Options, args ...any) (*Contract, common.Hash, error) {
	if opts == nil {
		opts = &transactions.Options{}
	}
	clause, err := d.AsClause(args...)
	txID := common.Hash{}
	if err != nil {
		return nil, txID, fmt.Errorf("failed to pack contract arguments: %w", err)
	}
	trx, err := sender.SendClauses([]*tx.Clause{clause}, opts)
	if err != nil {
		return nil, txID, fmt.Errorf("failed to send contract deployment transaction: %w", err)
	}
	receipt, err := trx.Wait(ctx)
	if err != nil {
		return nil, txID, fmt.Errorf("failed to wait for contract deployment: %w", err)
	}
	if receipt.Reverted {
		return nil, txID, errors.New("contract deployment reverted")
	}

	address := receipt.Outputs[0].ContractAddress

	return NewContract(d.client, *address, d.abi), trx.ID(), nil
}

// AsClause returns the contract deployment clause.
func (d *Deployer) AsClause(args ...any) (*tx.Clause, error) {
	contractArgs, err := d.abi.Pack("", args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack contract arguments: %w", err)
	}
	bytecode := append(d.bytecode, contractArgs...)
	clause := tx.NewClause(nil).WithData(bytecode).WithValue(d.value)
	return clause, nil
}
