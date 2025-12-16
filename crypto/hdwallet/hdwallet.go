package hdwallet

import (
	"crypto/ecdsa"

	"github.com/btcsuite/btcd/btcutil/hdkeychain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// PathVET is the root path to which custom derivation endpoints are appended.
// As such, the first account will be at m/44'/818'/0'/0, the second
// at m/44'/818'/0'/1, etc.
var PathVET = accounts.DerivationPath{0x80000000 + 44, 0x80000000 + 818, 0x80000000 + 0, 0}
var PathETH = accounts.DefaultRootDerivationPath

// Wallet is the underlying wallet struct.
type Wallet struct {
	masterKey *hdkeychain.ExtendedKey
	//seed      []byte
	path      accounts.DerivationPath
	publicKey *ecdsa.PublicKey
}

// FromSeed generates a wallet from a BIP-39 seed.
func FromSeed(seed []byte) (*Wallet, error) {
	return FromSeedAt(seed, PathVET)
}

// FromSeedAt generates a wallet from a BIP-39 seed and a specific derivation path.
func FromSeedAt(seed []byte, path accounts.DerivationPath) (*Wallet, error) {
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	w := &Wallet{
		masterKey: masterKey,
		path:      path,
	}

	return w.Derive(path)
}

// FromMnemonic generates a wallet using the PathVET derivation path.
func FromMnemonic(mnemonic string) (*Wallet, error) {
	return FromMnemonicAt(mnemonic, PathVET)
}

// FromMnemonicAt generates a wallet from a BIP-39 mnemonic and a specific derivation path.
func FromMnemonicAt(mnemonic string, path accounts.DerivationPath) (*Wallet, error) {
	seed, err := NewSeedFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	w := &Wallet{
		masterKey: masterKey,
		path:      path,
	}

	return w.Derive(path)
}

// Derive returns a new wallet derived from the root seed and master key at the given path.
func (w *Wallet) Derive(path accounts.DerivationPath) (*Wallet, error) {
	key, err := w.derivePath(path)
	if err != nil {
		return nil, err
	}

	publicKey, err := key.ECPubKey()
	if err != nil {
		return nil, err
	}
	publicKeyECDSA := publicKey.ToECDSA()

	return &Wallet{
		masterKey: w.masterKey,
		path:      path,
		publicKey: publicKeyECDSA,
	}, nil
}

// Child returns a child of the current wallet at the given index.
func (w *Wallet) Child(index uint32) (*Wallet, error) {
	path := make(accounts.DerivationPath, len(w.path)+1)
	copy(path, w.path)
	path[len(w.path)] = index
	return w.Derive(path)
}

// PrivateKey returns the ECDSA private key of the account.
func (w *Wallet) PrivateKey() (*ecdsa.PrivateKey, error) {
	key, err := w.derivePath(w.path)
	if err != nil {
		return nil, err
	}

	privateKey, err := key.ECPrivKey()
	if err != nil {
		return nil, err
	}

	return privateKey.ToECDSA(), nil
}

// MustGetPrivateKey returns the ECDSA private key of the account.
// It panics if the private key cannot be retrieved.
func (w *Wallet) MustGetPrivateKey() *ecdsa.PrivateKey {
	key, err := w.PrivateKey()
	if err != nil {
		panic(err)
	}

	return key
}

// PublicKey returns the ECDSA public key of the account.
func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

// Address returns the address of the current master key.
func (w *Wallet) Address() common.Address {
	return crypto.PubkeyToAddress(*w.publicKey)
}

// derivePath derives the key at the given path.
func (w *Wallet) derivePath(path accounts.DerivationPath) (*hdkeychain.ExtendedKey, error) {
	var err error
	key := w.masterKey
	for _, n := range path {
		key, err = key.Derive(n)
		if err != nil {
			return nil, err
		}
	}

	return key, nil
}
