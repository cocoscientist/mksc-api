package helper

import (
	"math/rand"
	"time"
)

const letterSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ApiKeyGenerator(username string) (result string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterSet[r.Int63()%int64(len(letterSet))]
	}
	return string(b)
}
