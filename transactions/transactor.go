package transactions

import (
	"fmt"
	"math"

	"github.com/btcsuite/btcd/wire"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
)

// Transactor is a transaction builder that can be used to simulate, build and send transactions.
type Transactor struct {
	client  *thorest.Client
	clauses []*tx.Clause
}

func NewTransactor(client *thorest.Client, clauses []*tx.Clause) *Transactor {
	return &Transactor{
		client:  client,
		clauses: clauses,
	}
}

// Simulate estimates the gas usage and checks for errors or reversion in the transaction.
func (t *Transactor) Simulate(caller common.Address, options *Options) (Simulation, error) {
	request := thorest.InspectRequest{
		Clauses: t.clauses,
		Caller:  &caller,
	}

	if options != nil && options.GasPayer != nil {
		request.GasPayer = options.GasPayer
	}

	response, err := t.client.Inspect(request)
	if err != nil {
		return Simulation{}, err
	}

	lastResult := response[len(response)-1]

	var consumedGas uint64
	for _, res := range response {
		consumedGas += res.GasUsed
	}

	intrinsicGas, err := tx.IntrinsicGas(t.clauses...)
	if err != nil {
		return Simulation{}, err
	}

	if intrinsicGas > math.MaxInt64 {
		return Simulation{}, fmt.Errorf("intrinsic gas exceeds maximum int64")
	}

	return Simulation{
		consumedGas:  consumedGas,
		vmError:      lastResult.VmError,
		reverted:     lastResult.Reverted,
		outputs:      response,
		intrinsicGas: intrinsicGas,
	}, nil
}

// Build constructs the transaction, applying defaults where necessary.
func (t *Transactor) Build(caller common.Address, options *Options) (*tx.Transaction, error) {
	builder := new(tx.Builder)

	for _, clause := range t.clauses {
		builder.Clause(clause)
	}

	if options != nil && options.Nonce != nil {
		builder.Nonce(*options.Nonce)
	} else {
		randomNonce, err := wire.RandomUint64()
		if err != nil {
			return nil, err
		}
		builder.Nonce(randomNonce)
	}

	if options != nil && (options.GasPayer != nil || (options.Delegation != nil && *options.Delegation)) {
		builder.Features(tx.DelegationFeature)
	}

	if options != nil && options.Gas != nil {
		builder.Gas(*options.Gas)
	} else {
		simulation, err := t.Simulate(caller, options)
		if err != nil {
			return nil, err
		}
		builder.Gas(simulation.TotalGas())
	}

	if options != nil && options.GasPriceCoef != nil {
		builder.GasPriceCoef(*options.GasPriceCoef)
	}

	if options != nil && options.Expiration != nil {
		builder.Expiration(*options.Expiration)
	} else {
		builder.Expiration(30)
	}

	if options != nil && options.BlockRef != nil {
		builder.BlockRef(*options.BlockRef)
	} else {
		best, err := t.client.BestBlock()
		if err != nil {
			return nil, err
		}
		builder.BlockRef(best.BlockRef())
	}

	if options != nil && options.DependsOn != nil {
		builder.DependsOn(options.DependsOn)
	}

	if options != nil && options.ChainTag != nil {
		builder.ChainTag(*options.ChainTag)
	} else {
		genesis, err := t.client.GenesisBlock()
		if err != nil {
			return nil, err
		}
		builder.ChainTag(genesis.ChainTag())
	}

	return builder.Build(), nil
}

type Signer interface {
	SignTransaction(tx *tx.Transaction) ([]byte, error)
	Address() common.Address
}

// Send will submit the transaction to the network.
func (t *Transactor) Send(signer Signer, options *Options) (*Visitor, error) {
	tx, err := t.Build(signer.Address(), options)
	if err != nil {
		return nil, fmt.Errorf("failed to build transaction: %w", err)
	}

	signature, err := signer.SignTransaction(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}
	tx = tx.WithSignature(signature)

	res, err := t.client.SendTransaction(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}

	return New(t.client, res.ID), nil
}
