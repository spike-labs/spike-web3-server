package util

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
)

func getABI(abiJSON string) abi.ABI {
	wrapABI, _ := abi.JSON(strings.NewReader(abiJSON))
	return wrapABI
}
