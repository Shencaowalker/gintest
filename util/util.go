package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letter = []byte("qwertyuiopasdfghjklzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}
