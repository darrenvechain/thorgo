package txmanager

import (
	"crypto/ecdsa"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// PKManager is a struct that manages a private key and provides methods to sign transactions and send clauses.
// Note: It is generally not recommended to use a private key directly in production code.
// Instead, consider using a secure key management solution or hardware wallet.
type PKManager struct {
	key  *ecdsa.PrivateKey
	thor *thorest.Client
}

// FromPK creates a new PKManager instance from an existing private key and a Thor client.
func FromPK(key *ecdsa.PrivateKey, thor *thorest.Client) *PKManager {
	return &PKManager{key: key, thor: thor}
}

// GeneratePK creates a new PKManager instance with a newly generated private key and a Thor client.
func GeneratePK(thor *thorest.Client) (*PKManager, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &PKManager{key: key, thor: thor}, nil
}

// Address returns the address derived from the private key.
func (p *PKManager) Address() (addr common.Address) {
	return crypto.PubkeyToAddress(p.key.PublicKey)
}

// PublicKey returns the public key associated with the private key.
func (p *PKManager) PublicKey() *ecdsa.PublicKey {
	return &p.key.PublicKey
}

// SendClauses sends a transaction with the given clauses and options.
func (p *PKManager) SendClauses(clauses []*tx.Clause, opts *transactions.Options) (*transactions.Visitor, error) {
	tx, err := transactions.NewTransactor(p.thor, clauses).Build(p.Address(), opts)
	if err != nil {
		return nil, err
	}
	signature, err := p.SignTransaction(tx)
	if err != nil {
		return nil, err
	}
	res, err := p.thor.SendTransaction(tx.WithSignature(signature))
	if err != nil {
		return nil, err
	}
	return transactions.New(p.thor, res.ID), nil
}

// SignTransaction signs the transaction with the private key.
// It returns the signature as a byte slice.
func (p *PKManager) SignTransaction(tx *tx.Transaction) ([]byte, error) {
	signature, err := crypto.Sign(tx.SigningHash().Bytes(), p.key)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

// Delegate signs the transaction as the gas payer.
func (p *PKManager) Delegate(tx *tx.Transaction, origin common.Address) ([]byte, error) {
	return crypto.Sign(tx.DelegatorSigningHash(origin).Bytes(), p.key)
}
