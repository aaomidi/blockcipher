package cbc

import (
	"github.com/aaomidi/blockcipher"
)

type CBC struct {
	blockcipher.CryptoMethod
}

func (c *CBC) Encrypt(dst, src [][]byte, IV []byte) {
	lastBlock := IV
	for i, v := range src {
		dst[i] = make([]byte, len(v))

		blockcipher.XORBytes(dst[i], v, lastBlock)
		c.EncryptBlock(dst[i], dst[i])

		lastBlock = dst[i]
	}
}

func (c *CBC) Decrypt(dst, src [][]byte, IV []byte) {
	lastBlock := IV
	for i, v := range src {
		dst[i] = make([]byte, len(v))

		c.DecryptBlock(dst[i], v)

		blockcipher.XORBytes(dst[i], dst[i], lastBlock)
		lastBlock = v
	}
}
