package test

import (
	"math/rand"
	"testing"

	"github.com/bayusamudra5502/stima-blockchain/lib"
)

func BenchmarkBF(b *testing.B) {
	ch := lib.BuildRandomChain(b.N)
	diff := int64(rand.Float64() / 4 * float64(b.N))

	chF, _ := lib.BuildSimilarChain(ch, diff)

	lib.FindFirstDiffBF(ch, chF)
}

func BenchmarkDnC(b *testing.B) {
	ch := lib.BuildRandomChain(b.N)
	diff := int64(rand.Float64() / 4 * float64(b.N))

	chF, _ := lib.BuildSimilarChain(ch, diff)

	lib.FindFirstDiffBF(ch, chF)
}
