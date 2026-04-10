package util

import (
	"crypto/rand"
	"math/big"
)

// RandomNumber generates a number randomly from min to max
func RandomNumber(min int64, max int64) int64 {
	rangeSize := big.NewInt(max - min + 1)

	n, _ := rand.Int(rand.Reader, rangeSize)
	return n.Int64() + min
}
