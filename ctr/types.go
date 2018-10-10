package ctr

import "github.com/aaomidi/blockcipher"

type CTR struct {
	IVIncrement
	blockcipher.CryptoMethod
}

type IVIncrement interface {
	Increment(iv []byte)
}

func (c *CTR) Apply(dst, src [][]byte, IV []byte) {
	for i, v := range src {
		dst[i] = make([]byte, len(v))

		encrypted := make([]byte, len(IV))

		c.EncryptBlock(encrypted, IV)

		blockcipher.XORBytes(dst[i], v, encrypted)

		c.Increment(IV)
	}
}
