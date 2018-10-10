package blockcipher

type CryptoMethod interface {
	EncryptBlock(dst, src []byte)
	DecryptBlock(dst, src []byte)
}
