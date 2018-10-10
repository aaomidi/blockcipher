package ecb

import "github.com/aaomidi/blockcipher"

type ECB struct {
	blockcipher.CryptoMethod
}

func (c *ECB) Encrypt(dst, src [][]byte) {
	for i, v := range src {
		encrypted := make([]byte, len(v))

		c.EncryptBlock(encrypted, v)

		dst[i] = encrypted
	}
}

func (c *ECB) Decrypt(dst, src [][]byte) {
	for i, v := range src {
		decrypted := make([]byte, len(v))

		c.DecryptBlock(decrypted, v)

		dst[i] = decrypted
	}
}
