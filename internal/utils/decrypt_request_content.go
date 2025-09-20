package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/totoro-paradise/goapp/internal/data"
)

// DecryptRequestContent takes a base64 encoded RSA payload encrypted with the
// project's public key and returns the decoded JSON payload as a generic map.
// It mirrors the behaviour of the Node implementation but fixes the bug where
// the ciphertext was double stringified before the decryption step.
func DecryptRequestContent(req string) (map[string]any, error) {
	block, _ := pem.Decode([]byte(data.PrivateKey))
	if block == nil {
		return nil, errors.New("rsa: invalid private key pem block")
	}

	privateKeyAny, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("rsa: parse private key: %w", err)
	}

	privateKey, ok := privateKeyAny.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("rsa: unexpected private key type")
	}

	cipherBytes, err := base64.StdEncoding.DecodeString(req)
	if err != nil {
		return nil, fmt.Errorf("rsa: decode base64 payload: %w", err)
	}

	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherBytes)
	if err != nil {
		return nil, fmt.Errorf("rsa: decrypt payload: %w", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(decrypted, &payload); err != nil {
		return nil, fmt.Errorf("rsa: parse decrypted payload: %w", err)
	}

	return payload, nil
}
