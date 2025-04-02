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
func (t *Transactor) Simulate(caller common.Address, options *Options) (*Simulation, error) {
	request := thorest.InspectRequest{
		Clauses: t.clauses,
		Caller:  &caller,
	}

	if options.GasPayer != nil {
		request.GasPayer = options.GasPayer
	}

	response, err := t.client.Inspect(request)
	if err != nil {
		return nil, err
	}

	lastResult := response[len(response)-1]

	var consumedGas uint64
	for _, res := range response {
		consumedGas += res.GasUsed
	}

	intrinsicGas, err := tx.IntrinsicGas(t.clauses...)
	if err != nil {
		return nil, err
	}

	if intrinsicGas > math.MaxInt64 {
		return nil, fmt.Errorf("intrinsic gas exceeds maximum int64")
	}

	return &Simulation{
		consumedGas:  consumedGas,
		vmError:      lastResult.VmError,
		reverted:     lastResult.Reverted,
		outputs:      response,
		intrinsicGas: intrinsicGas,
	}, nil
}

// Build constructs the transaction, applying defaults where necessary.
func (t *Transactor) Build(caller common.Address, options *Options) (*tx.Transaction, error) {
	if options == nil {
		options = &Options{}
	}
	best, err := t.client.BestBlock()
	if err != nil {
		return nil, err
	}

	txType := tx.TypeLegacy
	if best.BaseFee != nil {
		txType = tx.TypeDynamicFee
	}

	builder := tx.NewBuilder(txType)

	for _, clause := range t.clauses {
		builder.Clause(clause)
	}

	if options.Nonce != nil {
		builder.Nonce(*options.Nonce)
	} else {
		randomNonce, err := wire.RandomUint64()
		if err != nil {
			return nil, err
		}
		builder.Nonce(randomNonce)
	}

	if options.GasPayer != nil || (options.Delegation != nil && *options.Delegation) {
		builder.Features(tx.DelegationFeature)
	}

	if options.Gas != nil {
		builder.Gas(*options.Gas)
	} else {
		simulation, err := t.Simulate(caller, options)
		if err != nil {
			return nil, err
		}
		builder.Gas(simulation.TotalGas())
	}

	switch txType {
	case tx.TypeLegacy:
		if options.GasPriceCoef != nil {
			builder.GasPriceCoef(*options.GasPriceCoef)
		}
	case tx.TypeDynamicFee:
		if options.MaxPriorityFeePerGas != nil {
			builder.MaxPriorityFeePerGas(options.MaxPriorityFeePerGas)
		}
		if options.MaxFeePerGas != nil {
			builder.MaxFeePerGas(options.MaxFeePerGas)
		} else {
			fees, err := t.client.FeesHistory(thorest.RevisionNext(), 1, []float64{})
			if err != nil {
				return nil, err
			}
			if len(fees.BaseFeePerGas) < 1 {
				return nil, fmt.Errorf("missing base fees from fees history")
			}
			suggestion, err := t.client.FeesPriority()
			if err != nil {
				return nil, err
			}
			maxFee := fees.BaseFeePerGas[0].ToInt()
			maxFee = maxFee.Add(maxFee, suggestion.MaxPriorityFeePerGas.ToInt())
			builder.MaxFeePerGas(maxFee)
		}
	}

	if options.Expiration != nil {
		builder.Expiration(*options.Expiration)
	} else {
		builder.Expiration(30)
	}

	if options.BlockRef != nil {
		builder.BlockRef(*options.BlockRef)
	} else {
		builder.BlockRef(best.BlockRef())
	}

	if options.DependsOn != nil {
		builder.DependsOn(options.DependsOn)
	}

	if options.ChainTag != nil {
		builder.ChainTag(*options.ChainTag)
	} else {
		chainTag, err := t.client.ChainTag()
		if err != nil {
			return nil, err
		}
		builder.ChainTag(chainTag)
	}

	return builder.Build(), nil
}
