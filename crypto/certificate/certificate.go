package certificate

import (
	"encoding/json"
	"fmt"

	"github.com/darrenvechain/thorgo/crypto/hash"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

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

func (c *Certificate) Encode() ([]byte, error) {
	cert := *c
	cert.Signature = nil
	return json.Marshal(cert)
}

func (c *Certificate) SigningHash() (common.Hash, error) {
	encoded, err := c.Encode()
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to encode certificate: %w", err)
	}

	return hash.Blake2b(encoded), nil
}

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
