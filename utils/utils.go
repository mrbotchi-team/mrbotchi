package utils

import "crypto/rand"

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); nil != err {
		return nil, err
	}
	return b, nil
}
