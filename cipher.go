package enigma

import (
	"strings"
)

type Cipher interface {
	Encrypt(from Alphabet) Alphabet
	Clone() Cipher
	Inverse() Cipher
}

func CombineCiphers(tfs ...Cipher) Cipher {
	combined := CipherList{}
	for _, tf := range tfs {
		combined = append(combined, tf)
	}
	return combined
}

type CipherList []Cipher

func (tl CipherList) Encrypt(from Alphabet) Alphabet {
	to := from
	for _, tf := range tl {
		to = tf.Encrypt(to)
	}
	return to
}

func (tl CipherList) Clone() Cipher {
	clone := []Cipher{}
	for _, tf := range tl {
		clone = append(clone, tf.Clone())
	}
	return CipherList(clone)
}

func (tl CipherList) Inverse() Cipher {
	inverse := []Cipher{}
	for i := len(tl) - 1; i >= 0; i-- {
		inverse = append(inverse, tl[i].Inverse())
	}
	return CipherList(inverse)
}

func EncryptString(tf Cipher, from string) string {
	to := []byte{}
	for _, alph := range []byte(from) {
		to = append(to, byte(tf.Encrypt(Alphabet(alph))))
	}
	return string(to)
}

func CipherAsString(tf Cipher) string {
	str := []byte{}
	for i := 0; i < 26; i++ {
		str = append(str, byte(tf.Encrypt(Alphabet('a'+byte(i)))))
	}
	return strings.ToUpper(string(str))
}
