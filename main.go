package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/bayusamudra5502/stima-blockchain/lib"
)

const TRY_NUMBER = 5

func main() {
	group := &sync.WaitGroup{}

	fmt.Println("SIZE;ASSERT;BF TIME;DnC Time")

	for i := 0; i < 3; i++ {
		group.Add(1)

		go (func() {
			defer group.Done()

			for i := 0; i < TRY_NUMBER; i++ {
				size := rand.Int31n(99_000) + 1_000

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
		})()
	}

	group.Wait()
}
