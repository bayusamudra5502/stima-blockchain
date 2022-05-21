package test

import (
	"testing"

	"github.com/bayusamudra5502/stima-blockchain/lib"
	"github.com/stretchr/testify/assert"
)

func min(arr []int64) int64 {
	if len(arr) == 0 {
		return -1
	}

	result := arr[0]

	for _, val := range arr {
		if result > val {
			result = val
		}
	}

	return result
}

func TestBF(t *testing.T) {
	size := []int{8, 64, 1024, 16284, 30000}
	diff := []int64{0, 1, 2, 4}

	for _, i := range size {
		ch := lib.BuildRandomChain(i)

		for _, j := range diff {
			chSim, ans := lib.BuildSimilarChain(ch, j)
			minval := min(ans)
			actualMin := lib.FindFirstDiffBF(ch, chSim)

			assert.Equal(t, minval, actualMin)
		}
	}
}

func TestDnC(t *testing.T) {
	size := []int{8, 64, 1024, 16284, 30000}
	diff := []int64{0, 1, 2, 4}

	for _, i := range size {
		ch := lib.BuildRandomChain(i)

		for _, j := range diff {
			chSim, ans := lib.BuildSimilarChain(ch, j)
			minval := min(ans)
			actualMin := lib.FindFirstDiff(ch, chSim)

			assert.Equal(t, minval, actualMin)
		}
	}
}
