package txmanager

import (
	"crypto/ecdsa"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type PKManager struct {
	key  *ecdsa.PrivateKey
	thor *thorest.Client
}

func FromPK(key *ecdsa.PrivateKey, thor *thorest.Client) *PKManager {
	return &PKManager{key: key, thor: thor}
}

func GeneratePK(thor *thorest.Client) (*PKManager, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	return &PKManager{key: key, thor: thor}, nil
}

func (p *PKManager) Address() (addr common.Address) {
	return crypto.PubkeyToAddress(p.key.PublicKey)
}

func (p *PKManager) PublicKey() *ecdsa.PublicKey {
	return &p.key.PublicKey
}

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

func (p *PKManager) SignTransaction(tx *tx.Transaction) ([]byte, error) {
	signature, err := crypto.Sign(tx.SigningHash().Bytes(), p.key)
	if err != nil {
		return nil, err
	}

	return signature, nil
}
