package transactions

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/btcsuite/btcd/wire"
	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/common"
)

const GasBuffer = uint64(20_000)

// Transactor is a transaction builder that can be used to simulate and build and send transactions.
type Transactor struct {
	client  *thorest.Client
	clauses []*tx.Clause
}

// NewTransactor creates a new Transactor instance with the given client and clauses.
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
	if len(response) == 0 {
		return nil, errors.New("empty response from inspection")
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
// It sets the transaction type based on the presence of base fees and other options.
// If option values are not provided, it uses defaults or estimates them.
// - Nonce: If not provided, a random nonce is generated.
// - Gas: If not provided, it estimates the gas using simulation.
// - Expiration: If not provided, defaults to 30.
// - BlockRef: If not provided, uses the best block reference.
// - ChainTag: If not provided, retrieves the chain tag from the client.
// - DependsOn: If not provided, no dependency is set.
// - GasPayer: If provided, enables the delegation feature.
// - Delegation: If true, enables the delegation feature.
// - MaxFeePerGas: If not provided, retrieves the max fee from the client.
// - MaxPriorityFeePerGas: If not provided, an additional 5% of the bast fee is added.
// - GasPriceCoef: If not provided, defaults to 0 for legacy transactions.
// - GasBuffer: If not provided, defaults to 20,000 to ensure sufficient gas is available for the transaction.
func (t *Transactor) Build(caller common.Address, options *Options) (*tx.Transaction, error) {
	if options == nil {
		options = &Options{}
	}
	best, err := t.client.BestBlock()
	if err != nil {
		return nil, err
	}

	txType := tx.TypeLegacy  // default to legacy tx
	if best.BaseFee != nil { // use dynamic fee tx if base fee is present
		txType = tx.TypeDynamicFee
	}
	if options.GasPriceCoef != nil { // force legacy tx if GasPriceCoef is set
		txType = tx.TypeLegacy
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

	var gas uint64
	if options.Gas != nil {
		gas = *options.Gas
	} else {
		simulation, err := t.Simulate(caller, options)
		if err != nil {
			return nil, err
		}
		gas = simulation.TotalGas()
	}
	if options.GasBuffer != nil {
		gas += *options.GasBuffer
	} else {
		gas += GasBuffer
	}
	builder.Gas(gas)

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
			maxFee := new(big.Int).Mul(best.BaseFee.ToInt(), big.NewInt(105))
			maxFee = maxFee.Div(maxFee, big.NewInt(100))
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
