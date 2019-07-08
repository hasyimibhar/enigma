package enigma

// BlockCipher implements the cipher.Block interface.
type BlockCipher struct {
	*Enigma
}

func NewCipher(key *Config) (*BlockCipher, error) {
	enigma, err := NewWithConfig(key)
	if err != nil {
		return nil, err
	}

	return &BlockCipher{enigma}, nil
}

func (c *BlockCipher) BlockSize() int {
	return 8
}

func (c *BlockCipher) Encrypt(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		dst[i] = byte(c.Enigma.Encrypt(Alphabet(src[i])))
	}
}

func (c *BlockCipher) Decrypt(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		dst[i] = byte(c.Enigma.Encrypt(Alphabet(src[i])))
	}
}
