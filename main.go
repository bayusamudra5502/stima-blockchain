package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/bayusamudra5502/stima-blockchain/lib"
)

var CHAIN_LEN_TC = []int64{1_000, 5_000, 10_000, 50_000, 100_000}

const TRY_NUMBER = 5

func main() {
	group := &sync.WaitGroup{}

	fmt.Println("SIZE;ASSERT;BF TIME;DnC Time")

	for _, size := range CHAIN_LEN_TC {
		group.Add(1)

		go (func(size int64) {
			defer group.Done()

			for i := 0; i < TRY_NUMBER; i++ {
				for j := 0; j < 3; j++ {
					ch := lib.BuildRandomChain(int(size))
					ch2, _ := lib.BuildSimilarChain(ch, int64(j))

					startBFTime := time.Now().UnixNano()
					bfResult := lib.FindFirstDiffBF(ch, ch2)

					BFTime := time.Now().UnixNano() - startBFTime

					startDnCTime := time.Now().UnixNano()
					dncResult := lib.FindFirstDiff(ch, ch2)

					DnCTime := time.Now().UnixNano() - startDnCTime

					fmt.Printf("%d;%s;%d;%d\n", size, strconv.FormatBool(dncResult == bfResult), BFTime, DnCTime)
				}
			}
		})(size)
	}

	group.Wait()
}
