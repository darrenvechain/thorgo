package accounts

import (
	"math/big"

	"github.com/darrenvechain/thorgo/crypto/tx"
	"github.com/darrenvechain/thorgo/thorest"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Visitor is a struct that provides methods to interact with the VeChain accounts API.
// It allows you to fetch account information, code, storage, and perform read-only contract calls.
type Visitor struct {
	client   *thorest.Client
	account  common.Address
	revision *thorest.Revision
}

// New creates a new Visitor instance.
func New(c *thorest.Client, account common.Address) *Visitor {
	return &Visitor{client: c, account: account}
}

// Revision sets the optional revision for the API calls.
func (a *Visitor) Revision(revision thorest.Revision) *Visitor {
	a.revision = &revision
	return a
}

// Get fetches the account information for the given address.
func (a *Visitor) Get() (*thorest.Account, error) {
	if a.revision == nil {
		return a.client.Account(a.account)
	}
	return a.client.AccountAt(a.account, *a.revision)
}

// Code fetches the byte code of the contract at the given address.
func (a *Visitor) Code() (*thorest.AccountCode, error) {
	if a.revision == nil {
		return a.client.AccountCode(a.account)
	}

	return a.client.AccountCodeAt(a.account, *a.revision)
}

// Storage fetches the storage value for the given key.
func (a *Visitor) Storage(key common.Hash) (*thorest.AccountStorage, error) {
	if a.revision == nil {
		return a.client.AccountStorage(a.account, key)
	}

	return a.client.AccountStorageAt(a.account, key, *a.revision)
}

// Call executes a read-only contract call.
func (a *Visitor) Call(calldata []byte) (*thorest.InspectResponse, error) {
	clause := tx.NewClause(&a.account).WithData(calldata).WithValue(big.NewInt(0))
	request := thorest.InspectRequest{Clauses: []*tx.Clause{clause}}
	var (
		inspection []thorest.InspectResponse
		err        error
	)

	if a.revision == nil {
		inspection, err = a.client.Inspect(request)
	} else {
		inspection, err = a.client.InspectAt(request, *a.revision)
	}

	if err != nil {
		return nil, err
	}

	return &inspection[0], nil
}

// Contract returns a new Contract instance.
func (a *Visitor) Contract(abi *abi.ABI) *Contract {
	return NewContract(a.client, a.account, abi)
}
