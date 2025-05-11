package txmanager

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/darrenvechain/thorgo/transactions"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// DelegatedManager is a transaction manager that delegates the payment of transaction fees to a Delegator
type DelegatedManager struct {
	thor     *thorest.Client
	gasPayer Delegator
	origin   Signer
}

// NewDelegated creates a new DelegatedManager instance
func NewDelegated(thor *thorest.Client, origin Signer, gasPayer Delegator) *DelegatedManager {
	return &DelegatedManager{
		thor:     thor,
		origin:   origin,
		gasPayer: gasPayer,
	}
}

// SignTransaction signs the transaction with both the origin and the delegator
// Returns the combined signature of the origin and the delegator
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

// SendClauses sends the transaction with the given clauses and options
func (d *DelegatedManager) SendClauses(clauses []*tx.Clause, opts *transactions.Options) (*transactions.Visitor, error) {
	if opts == nil {
		opts = &transactions.Options{}
	}
	if opts.Delegation == nil || !*opts.Delegation {
		delegated := true
		opts.Delegation = &delegated
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

// URLDelegator is a delegator that uses a remote URL to pay for transaction fees
type URLDelegator struct {
	url string
}

// NewUrlDelegator creates a new URLDelegator instance
func NewUrlDelegator(url string) *URLDelegator {
	return &URLDelegator{url: url}
}

// Delegate requests the delegator's signature for the given transaction
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
