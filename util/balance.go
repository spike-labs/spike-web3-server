package util

import "math/big"

func ParseBalance(balance *big.Int) string {
	decimalBalance := ToDecimal(balance, 18)
	return decimalBalance.String()
}
