package lib

import (
	"crypto/sha256"
	"encoding/base64"
)

type ChainHash [32]byte

type ChainItem struct {
	Payload  string
	LastHash ChainHash
}

func (c ChainItem) GetHash() ChainHash {
	payload := []byte(c.Payload)
	payload = append(payload, c.LastHash[:]...)

	return sha256.Sum256(payload)
}

func (c ChainItem) GetEncode() string {
	payload := []byte(c.Payload)
	payload = append(payload, c.LastHash[:]...)

	hash := c.GetHash()
	payload = append(payload, ':')
	payload = append(payload, hash[:]...)

	return base64.StdEncoding.EncodeToString(payload)
}

func (c ChainItem) GetBase64Hash() string {
	hash := c.GetHash()
	return base64.StdEncoding.EncodeToString(hash[:])
}
