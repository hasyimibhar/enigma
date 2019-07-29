package main

import (
	"fmt"

	"github.com/hasyimibhar/enigma"
)

func main() {
	sharedKey := &enigma.Config{
		Plugboard: "EJ OY IV AQ KW FX MT PS LU BD",
		Rotors: []*enigma.RotorConfig{
			&enigma.RotorConfig{Name: "IV", Position: 14},
			&enigma.RotorConfig{Name: "II", Position: 22},
			&enigma.RotorConfig{Name: "V", Position: 25},
		},
		Reflector: &enigma.RotorConfig{Name: "A"},
	}

	aliceBlock, _ := enigma.NewCipher(sharedKey)
	bobBlock, _ := enigma.NewCipher(sharedKey)

	plaintext := []byte("loremipsumdolorsitamet")
	ciphertext := make([]byte, len(plaintext))

	alice := NewECBEncrypter(aliceBlock)
	bob := NewECBDecrypter(bobBlock)

	alice.CryptBlocks(ciphertext, plaintext)

	fmt.Printf("[alice] %s -> %s\n", string(plaintext), (ciphertext))

	bob.CryptBlocks(plaintext, ciphertext)
	fmt.Printf("[bob] %s -> %s\n", string(ciphertext), (plaintext))
}
