package util

import (
	"crypto/rand"
	"math/big"
)

// RandomString generates a string whose length is len.
func RandomString(len int) string {
	strArr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	var str string
	for i := 0; i < len; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len)))
		str += string(strArr[n.Int64()])
	}

	return str
}
