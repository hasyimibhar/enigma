package main

import (
	"crypto/cipher"
)

type ECB struct {
	block   cipher.Block
	encrypt bool
}

func NewECBEncrypter(block cipher.Block) cipher.BlockMode {
	return ECB{block, true}
}

func NewECBDecrypter(block cipher.Block) cipher.BlockMode {
	return ECB{block, false}
}

func (e ECB) BlockSize() int {
	return e.block.BlockSize()
}

func (e ECB) CryptBlocks(dst, src []byte) {
	if len(dst) < len(src) {
		panic("src is bigger than dst")
	}
	if e.block.BlockSize() > 1 {
		panic("blocksize > 1 is not supported")
	}

	if e.encrypt {
		for i := 0; i < len(src); i++ {
			e.block.Encrypt(dst[i:], src[i:])
		}
	} else {
		for i := 0; i < len(src); i++ {
			e.block.Decrypt(dst[i:], src[i:])
		}
	}
}
