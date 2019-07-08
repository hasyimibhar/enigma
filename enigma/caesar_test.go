package enigma

import (
	"strconv"
	"testing"
)

func TestCaesarCipher(t *testing.T) {
	tests := []struct {
		N                  int
		Plaintext          string
		ExpectedCiphertext string
	}{
		{0, "helloworld", "helloworld"},
		{13, "helloworld", "uryybjbeyq"},
		{-1, "helloworld", "gdkknvnqkc"},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.N), func(t *testing.T) {
			cipher := CaesarCipher(tt.N)
			ciphertext := TransformString(cipher, tt.Plaintext)

			if ciphertext != tt.ExpectedCiphertext {
				t.Fatalf("expecting '%s' * ROT%d => '%s', got '%s' instead",
					tt.Plaintext, tt.N, tt.ExpectedCiphertext, string(ciphertext))
			}
		})
	}
}

func TestCaesarCipher_Identity(t *testing.T) {
	ident := CaesarCipher(0)
	ciphertext := TransformString(ident, "helloworld")
	if ciphertext != "helloworld" {
		t.Fatal("caesar cipher does not have an identity element")
	}
}

func TestCaesarCipher_Inverse(t *testing.T) {
	cipher := CaesarCipher(17)
	inverse := cipher.Inverse()
	ciphertext := TransformString(CombineTransformations(cipher, inverse), "helloworld")

	if ciphertext != "helloworld" {
		t.Fatal("caesar cipher does not have an identity element")
	}
}
