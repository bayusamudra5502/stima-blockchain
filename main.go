package main

import (
	"github.com/bayusamudra5502/stima-blockchain/lib"
)

func main() {
	ch := lib.BuildRandomChain(10)
	ch.PrintChain()

	ch2, _ := lib.BuildSimilarChain(ch, 1)
	ch2.PrintChain()
}
