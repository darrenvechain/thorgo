package accounts

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/darrenvechain/thorgo/client"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Deployer struct {
	client   *client.Client
	bytecode []byte
	abi      *abi.ABI
	value    *big.Int
}

func NewDeployer(client *client.Client, bytecode []byte, abi *abi.ABI) *Deployer {
	return &Deployer{client: client, bytecode: bytecode, abi: abi, value: big.NewInt(0)}
}

func (d *Deployer) WithValue(value *big.Int) *Deployer {
	d.value = value
	return d
}

func (d *Deployer) Deploy(sender TxManager, args ...interface{}) (*Contract, common.Hash, error) {
	clause, err := d.AsClause(args...)
	txID := common.Hash{}
	if err != nil {
		return nil, txID, fmt.Errorf("failed to pack contract arguments: %w", err)
	}
	txID, err = sender.SendClauses([]*tx.Clause{clause})
	if err != nil {
		return nil, txID, fmt.Errorf("failed to send contract deployment transaction: %w", err)
	}
	receipt, err := transactions.New(d.client, txID).Wait()
	if err != nil {
		return nil, txID, fmt.Errorf("failed to wait for contract deployment: %w", err)
	}
	if receipt.Reverted {
		return nil, txID, errors.New("contract deployment reverted")
	}

	address := common.HexToAddress(receipt.Outputs[0].ContractAddress)

	return NewContract(d.client, address, d.abi), txID, nil
}

// AsClause returns the contract deployment clause.
func (d *Deployer) AsClause(args ...interface{}) (*tx.Clause, error) {
	contractArgs, err := d.abi.Pack("", args...)
	if err != nil {
		return nil, fmt.Errorf("failed to pack contract arguments: %w", err)
	}
	bytecode := append(d.bytecode, contractArgs...)
	clause := tx.NewClause(nil).WithData(bytecode).WithValue(d.value)
	return clause, nil
}
