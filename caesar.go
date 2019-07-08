package enigma

type CaesarCipher int

func (n CaesarCipher) Transform(from Alphabet) (to Alphabet) {
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

func (n CaesarCipher) Clone() Transformation {
	return CaesarCipher(int(n))
}

func (n CaesarCipher) Inverse() Transformation {
	return CaesarCipher(int(n) * -1)
}
