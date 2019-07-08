package main

import (
	"fmt"
	"os"

	"github.com/hasyimibhar/enigma/enigma"
)

func main() {
	machine, err := enigma.NewWithConfig(&enigma.Config{
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

	clone := machine.Clone()

	plaintext := "helloworld"
	ciphertext := machine.TransformString(plaintext)

	fmt.Printf("%s -> %s\n", plaintext, ciphertext)

	plaintext = ciphertext
	ciphertext = clone.TransformString(plaintext)

	fmt.Printf("%s -> %s\n", plaintext, ciphertext)
}
