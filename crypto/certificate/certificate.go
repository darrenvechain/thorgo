package certificate

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/darrenvechain/thorgo/crypto/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Certificate implements the Simple Self-signed Certificate specification.
// See: https://github.com/vechain/VIPs/blob/master/vips/VIP-192.md
type Certificate struct {
	Domain    string         `json:"domain"`
	Payload   Payload        `json:"payload"`
	Purpose   string         `json:"purpose"`
	Signer    string         `json:"signer"`
	Timestamp int64          `json:"timestamp"`
	Signature *hexutil.Bytes `json:"signature,omitempty"`
}

type Payload struct {
	Content string `json:"content"`
	Type    string `json:"type"`
}

// FromBytes decodes a byte array into a Certificate.
func FromBytes(data []byte) (*Certificate, error) {
	cert := &Certificate{}
	if err := json.Unmarshal(data, cert); err != nil {
		return nil, fmt.Errorf("failed to unmarshal certificate: %w", err)
	}
	return cert, nil
}

// Encode encodes the Certificate into a JSON byte array.
func (c *Certificate) Encode() ([]byte, error) {
	cert := *c
	cert.Signature = nil
	return json.Marshal(cert)
}

// SigningHash computes the signing hash of the certificate.
func (c *Certificate) SigningHash() (common.Hash, error) {
	encoded, err := c.Encode()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to encode certificate: %w", err)
	}

	return hash.Blake2b(encoded), nil
}

// Verify checks the signature of the certificate against the signer.
func (c *Certificate) Verify() bool {
	if c.Signature == nil {
		return false
	}
	signingHash, err := c.SigningHash()
	if err != nil {
		return false
	}
	pubkey, err := crypto.SigToPub(signingHash.Bytes(), *c.Signature)
	if err != nil {
		return false
	}
	return crypto.PubkeyToAddress(*pubkey) == common.HexToAddress(c.Signer)
}

// Sign signs the certificate with the given private key and sets the Signature field.
func (c *Certificate) Sign(privateKey *ecdsa.PrivateKey) error {
	signingHash, err := c.SigningHash()
	if err != nil {
		return fmt.Errorf("failed to compute signing hash: %w", err)
	}
	sig, err := crypto.Sign(signingHash.Bytes(), privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign certificate: %w", err)
	}
	signature := hexutil.Bytes(sig)
	c.Signature = &signature
	return nil
}
