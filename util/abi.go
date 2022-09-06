package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"strings"
)

func GetABI(abiJSON string) abi.ABI {
	wrapABI, _ := abi.JSON(strings.NewReader(abiJSON))
	return wrapABI
}

func EventSignHash(eventTopic string) string {
	eventSignature := []byte(eventTopic)
	hash := crypto.Keccak256Hash(eventSignature)
	return hash.Hex()
}

func GetTxMethodName(method string) []byte {
	methodName := []byte(method)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(methodName)
	return hash.Sum(nil)[:4]
}
