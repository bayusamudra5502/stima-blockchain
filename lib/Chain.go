package lib

import (
	"encoding/base64"
	"fmt"
)

type Chain struct {
	chain  []*ChainItem
	length int64
}

func NewChain() Chain {
	var ch Chain
	ch.length = 0

	return ch
}

func (c *Chain) AddChain(payload string) {
	var lastHash [32]byte
	lastChain := c.LastChain()

	if lastChain != nil {
		lastHash = lastChain.GetHash()
	} else {
		for i := range lastHash {
			lastHash[i] = 0
		}
	}

	newItem := &ChainItem{Payload: payload, LastHash: lastHash}
	c.chain = append(c.chain, newItem)
	c.length++
}

func (c Chain) LastChain() *ChainItem {
	length := c.length

	if length == 0 {
		return nil
	}

	return c.chain[length-1]
}

func (c Chain) GetChainItem(index int64) *ChainItem {
	if c.length <= int64(index) || index < 0 {
		return nil
	}

	return c.chain[index]
}

func (c Chain) PrintChain() {
	for i, val := range c.chain {
		fmt.Println("Chain Idx : ", i)
		if i == 0 {
			fmt.Println("Last Hash : null")
		} else {
			fmt.Println("Last Hash : ", base64.StdEncoding.EncodeToString(val.LastHash[:]))
		}
		fmt.Println("Payload: ", val.Payload)

		hash := val.GetHash()
		fmt.Println("Hash : ", base64.StdEncoding.EncodeToString(hash[:]))
		fmt.Println()
	}
}
