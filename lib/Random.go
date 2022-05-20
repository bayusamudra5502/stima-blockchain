package lib

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func BuildRandomChain(length int) Chain {
	ch := NewChain()

	for i := 0; i < length; i++ {
		len := rand.Int63n(1000) + 1
		ch.AddChain(RandomString(len))
	}

	return ch
}

func BuildSimilarChain(seed Chain, diff int) Chain {
	ch := NewChain()
	var i int64

	for i = 0; i < seed.length; i++ {
		p := rand.Float64()

		if p > 0.65 && diff > 0 {
			len := rand.Int63n(1000) + 1

			ch.AddChain(RandomString(len))
			diff--
		} else {
			ch.AddChain(ch.GetChainItem(i).Payload)
		}
	}

	return ch
}

func RandomString(length int64) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
