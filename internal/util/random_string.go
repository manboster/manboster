package util

import (
	"crypto/rand"
	"math/big"
)

// RandomString generates a string whose length is len.
func RandomString(length int) string {
	strArr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	var str string
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(strArr))))
		str += string(strArr[n.Int64()])
	}

	return str
}
