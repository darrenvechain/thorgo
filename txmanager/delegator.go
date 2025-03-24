package txmanager

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// DelegatedManager is a transaction manager that delegates the payment of transaction fees to a Delegator
type DelegatedManager struct {
	thor     *thorest.Client
	gasPayer Delegator
	origin   Signer
}

type Signer interface {
	SignTransaction(tx *tx.Transaction) ([]byte, error)
	Address() common.Address
}

func NewDelegatedManager(thor *thorest.Client, origin Signer, gasPayer Delegator) *DelegatedManager {
	return &DelegatedManager{
		thor:     thor,
		origin:   origin,
		gasPayer: gasPayer,
	}
}

func (d *DelegatedManager) SignTransaction(tx *tx.Transaction) ([]byte, error) {
	signature, err := d.origin.SignTransaction(tx)
	if err != nil {
		return nil, err
	}
	delegatorSig, err := d.gasPayer.Delegate(tx, d.Address())
	if err != nil {
		return nil, err
	}
	signature = append(signature, delegatorSig...)
	return signature, nil
}

func (d *DelegatedManager) SendClauses(clauses []*tx.Clause, opts *transactions.Options) (*transactions.Visitor, error) {
	if opts == nil {
		opts = &transactions.Options{}
	}
	if opts.Delegation == nil || !*opts.Delegation {
		opts.Delegation = new(bool)
		*opts.Delegation = true
	}

	tx, err := transactions.NewTransactor(d.thor, clauses).Build(d.origin.Address(), opts)
	if err != nil {
		return nil, err
	}
	signature, err := d.SignTransaction(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}
	tx = tx.WithSignature(signature)
	res, err := d.thor.SendTransaction(tx)
	if err != nil {
		return nil, fmt.Errorf("failed to send transaction: %w", err)
	}
	return transactions.New(d.thor, res.ID), nil
}

// Address returns the address of the origin manager
func (d *DelegatedManager) Address() common.Address {
	return d.origin.Address()
}

// PKDelegator is a delegator that uses a private key to pay for transaction fees
type PKDelegator struct {
	key *ecdsa.PrivateKey
}

func NewDelegator(key *ecdsa.PrivateKey) *PKDelegator {
	return &PKDelegator{key: key}
}

func (p *PKDelegator) PublicKey() *ecdsa.PublicKey {
	return &p.key.PublicKey
}

func (p *PKDelegator) Address() (addr common.Address) {
	return crypto.PubkeyToAddress(p.key.PublicKey)
}

func (p *PKDelegator) Delegate(tx *tx.Transaction, origin common.Address) ([]byte, error) {
	return crypto.Sign(tx.DelegatorSigningHash(origin).Bytes(), p.key)
}

// URLDelegator is a delegator that uses a remote URL to pay for transaction fees
type URLDelegator struct {
	url string
}

func NewUrlDelegator(url string) *URLDelegator {
	return &URLDelegator{url: url}
}

func (p *URLDelegator) Delegate(tx *tx.Transaction, origin common.Address) ([]byte, error) {
	encoded, err := tx.MarshalBinary()
	if err != nil {
		return nil, err
	}

	req := &DelegateRequest{
		Origin: origin,
		Raw:    hexutil.Bytes(encoded),
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(p.url, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("200 OK expected")
	}

	defer res.Body.Close()
	var response DelegateResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Signature, nil
}
