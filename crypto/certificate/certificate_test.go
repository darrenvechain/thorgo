package certificate

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var (
	privKey, _  = crypto.HexToECDSA("37174033db12d0976a60e1699007057fb19bcaeffb1092c08c9e7ac5d519ff37")
	signer      = crypto.PubkeyToAddress(privKey.PublicKey)
	signature   = hexutil.Bytes(common.Hex2Bytes("4043e6cd7f21b62474cf3f7337177450ba85e349b878f152f6181e53fe616f0976aed20786fbbc66441d13b3b0c9402e7821b8dcd84a2947db613d0535cb541c00"))
	signingHash = "0xd8ab73da48ec11a58467856de337702a4deb2f5a362185b3fb72774b961a4675"
)

func newCert(signature *hexutil.Bytes) Certificate {
	return Certificate{
		Domain: "localhost",
		Payload: Payload{
			Type:    "text",
			Content: "fyi",
		},
		Purpose:   "identification",
		Signer:    strings.ToLower(signer.String()),
		Timestamp: 1545035330,
		Signature: signature,
	}
}

func TestCertificate_FromBytes(t *testing.T) {
	sig := hexutil.Encode(signature)
	testCases := []struct {
		name    string
		rawJson string
	}{
		{
			name:    "signed certificate",
			rawJson: fmt.Sprintf(`{"domain":"localhost","payload":{"content":"fyi","type":"text"},"purpose":"identification","signer":"%s","timestamp":1545035330,"signature":"%s"}`, strings.ToLower(signer.String()), sig),
		},
		{
			name:    "unsigned certificate",
			rawJson: fmt.Sprintf(`{"domain":"localhost","payload":{"content":"fyi","type":"text"},"purpose":"identification","signer":"%s","timestamp":1545035330}`, strings.ToLower(signer.String())),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cert, err := FromBytes([]byte(tc.rawJson))
			assert.NoError(t, err)
			assert.Equal(t, signer, common.HexToAddress(cert.Signer))
			assert.Equal(t, "localhost", cert.Domain)
			assert.Equal(t, "fyi", cert.Payload.Content)
			assert.Equal(t, "text", cert.Payload.Type)
			assert.Equal(t, "identification", cert.Purpose)
			assert.Equal(t, int64(1545035330), cert.Timestamp)
			if tc.name == "signed certificate" {
				assert.Equal(t, signature, *cert.Signature)
			} else {
				assert.Nil(t, cert.Signature)
			}
		})
	}
}

func TestCertificate_Encode(t *testing.T) {
	cert := newCert(nil)
	encoded, err := cert.Encode()
	assert.NoError(t, err)

	expected := fmt.Sprintf(`{"domain":"localhost","payload":{"content":"fyi","type":"text"},"purpose":"identification","signer":"%s","timestamp":1545035330}`, strings.ToLower(signer.String()))
	assert.Equal(t, expected, string(encoded))
}

func TestCertificate_SigningHash(t *testing.T) {
	cert := newCert(nil)
	hash, err := cert.SigningHash()
	assert.NoError(t, err)
	assert.Equal(t, signingHash, hash.String())
}

func TestCertificate_Verify(t *testing.T) {
	cert := newCert(&signature)
	assert.True(t, cert.Verify())
}

func TestCertificate_Unmarshal(t *testing.T) {
	sig := hexutil.Encode(signature)
	testCases := []struct {
		name     string
		rawJson  string
		verified bool
	}{
		{
			name:     "signed certificate",
			rawJson:  fmt.Sprintf(`{"domain":"localhost","payload":{"content":"fyi","type":"text"},"purpose":"identification","signer":"%s","timestamp":1545035330,"signature":"%s"}`, strings.ToLower(signer.String()), sig),
			verified: true,
		},
		{
			name:     "unsigned certificate",
			rawJson:  fmt.Sprintf(`{"domain":"localhost","payload":{"content":"fyi","type":"text"},"purpose":"identification","signer":"%s","timestamp":1545035330}`, strings.ToLower(signer.String())),
			verified: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cert := Certificate{}
			err := json.Unmarshal([]byte(tc.rawJson), &cert)
			assert.NoError(t, err)
			assert.Equal(t, tc.verified, cert.Verify())
		})
	}
}
