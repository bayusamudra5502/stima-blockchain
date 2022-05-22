package lib

import (
	"math/rand"
	"time"
)

const MAX_BYTES = 2048

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func BuildRandomChain(length int) Chain {
	ch := NewChain()

	for i := 0; i < length; i++ {
		len := rand.Int63n(MAX_BYTES-1) + 1
		ch.AddChain(RandomString(len))
	}

	return ch
}

func BuildSimilarChain(seed Chain, diff int64) (Chain, []int64) {
	ch := NewChain()
	var idx []int64
	var i int64

	for i = 0; i < diff; i++ {
		isFound := false
		newIdx := rand.Int63n(seed.length)

		for !isFound {
			newIdx = rand.Int63n(seed.length)
			isFound = true
			for _, val := range idx {
				if val == newIdx {
					isFound = false
				}
			}
		}

		idx = append(idx, newIdx)
	}

	for i = 0; i < seed.length; i++ {
		isRandom := false

		for _, val := range idx {
			if val == i {
				isRandom = true
				break
			}
		}

		if isRandom {
			len := rand.Int63n(MAX_BYTES-1) + 1

			ch.AddChain(RandomString(len))
			diff--
		} else {
			ch.AddChain(seed.GetChainItem(i).Payload)
		}
	}

	return ch, idx
}

func RandomString(length int64) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
