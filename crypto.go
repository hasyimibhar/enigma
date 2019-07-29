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
	return 1
}

func (c *BlockCipher) Encrypt(dst, src []byte) {
	dst[0] = byte(c.Enigma.Encrypt(Alphabet(src[0])))
}

func (c *BlockCipher) Decrypt(dst, src []byte) {
	dst[0] = byte(c.Enigma.Encrypt(Alphabet(src[0])))
}
