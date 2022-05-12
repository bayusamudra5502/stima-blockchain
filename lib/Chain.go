package lib

type Chain struct {
	chain []*ChainItem
}

func (c *Chain) AddChain(payload string, signature string) {
	var lastHash [32]byte
	lastChain := c.LastChain()

	if lastChain == nil {
		lastHash = lastChain.GetHash()
	}

	newItem := &ChainItem{payload, signature, lastHash}
	c.chain = append(c.chain, newItem)
}

func (c Chain) LastChain() *ChainItem {
	length := len(c.chain)

	if length == 0 {
		return nil
	}

	return c.chain[length-1]
}

func (c Chain) GetChainItem(index int) *ChainItem {
	return c.chain[index]
}
