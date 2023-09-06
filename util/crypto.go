package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func HmacSha256(data string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateEcdsaKey() (string, string, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}
	privateKeyPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privateKeyPEMBytes := pem.EncodeToMemory(privateKeyPEM)
	privateString := string(privateKeyPEMBytes)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicBase64 := base64.StdEncoding.EncodeToString(publicKeyBytes)
	log.Infof("private key: %s, public key: %s", privateString, publicBase64)
	return privateString, publicBase64, nil
}

func EcdsaSign(privatePemKey, msg string) (string, error) {
	privateKeyPEMBytes := []byte(privatePemKey)
	block, _ := pem.Decode(privateKeyPEMBytes)
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256([]byte(msg))
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	if err != nil {
		return "", err
	}
	signatureBytes := append(r.Bytes(), s.Bytes()...)
	signature := hexutil.Encode(signatureBytes)
	return signature, nil
}
