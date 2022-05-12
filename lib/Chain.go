package lib

type Chain struct {
	chain  []*ChainItem
	length int64
}

func (c *Chain) AddChain(payload string, signature string) {
	var lastHash [32]byte
	lastChain := c.LastChain()

	if lastChain == nil {
		lastHash = lastChain.GetHash()
	}

	newItem := &ChainItem{payload, signature, lastHash}
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
