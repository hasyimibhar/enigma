package main

import (
	"fmt"
	"os"

	"github.com/hasyimibhar/enigma/enigma"
)

func main() {
	cipher, err := enigma.NewWithConfig(&enigma.Config{
		Plugboard: "EJ OY IV AQ KW FX MT PS LU BD",
		Rotors: []*enigma.RotorConfig{
			&enigma.RotorConfig{Name: "IV", Position: 14},
			&enigma.RotorConfig{Name: "II", Position: 22},
			&enigma.RotorConfig{Name: "V", Position: 25},
		},
		Reflector: &enigma.RotorConfig{Name: "A"},
	})
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	clone := cipher.Clone()

	plaintext := "helloworld"
	ciphertext := []byte{}

	for _, alph := range plaintext {
		x := byte(cipher.Transform(enigma.Alphabet(alph)))
		ciphertext = append(ciphertext, x)
	}

	fmt.Printf("%s -> %s\n", plaintext, string(ciphertext))

	plaintext = string(ciphertext)
	ciphertext = []byte{}

	for _, alph := range plaintext {
		ciphertext = append(ciphertext, byte(clone.Transform(enigma.Alphabet(alph))))
	}

	fmt.Printf("%s -> %s\n", plaintext, string(ciphertext))
}
