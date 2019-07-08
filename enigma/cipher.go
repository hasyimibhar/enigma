package enigma

// Cipher implements the cipher.Block interface.
type Cipher struct {
	*Enigma
}

func NewCipher(key *Config) (*Cipher, error) {
	enigma, err := NewWithConfig(key)
	if err != nil {
		return nil, err
	}

	return &Cipher{enigma}, nil
}

func (c *Cipher) BlockSize() int {
	return 8
}

func (c *Cipher) Encrypt(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		dst[i] = byte(c.Enigma.Transform(Alphabet(src[i])))
	}
}

func (c *Cipher) Decrypt(dst, src []byte) {
	for i := 0; i < len(src); i++ {
		dst[i] = byte(c.Enigma.Transform(Alphabet(src[i])))
	}
}
