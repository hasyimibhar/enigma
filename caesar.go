package enigma

type CaesarCipher int

func (n CaesarCipher) Encrypt(from Alphabet) (to Alphabet) {
	if n >= 0 {
		to = Alphabet('a' + ((int(from-'a') + int(n)) % 26))
	} else {
		shift := int(from-'a') + int(n)
		for shift < 0 {
			shift += 26
		}

		to = Alphabet('a' + shift)
	}

	return
}

func (n CaesarCipher) Clone() Cipher {
	return CaesarCipher(int(n))
}

func (n CaesarCipher) Inverse() Cipher {
	return CaesarCipher(int(n) * -1)
}
