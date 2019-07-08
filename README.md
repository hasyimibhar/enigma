# Enigma Machine in Go

The Enigma Machine implemented in Go for learning purposes.

## Mathematical Analysis

A single encryption can be expressed with the following expression:

    E(X) = P * SUM(R[i], n) * U * SUM(R[i], n)^-1 * P^-1
    R[i] = ROT(p[i]) * T * ROT(-p[i])


where:

- `X` is a string of alphabets (a-z)
- the operator `*` concatenates 2 transformations from left to right:
    
        A = B * C
        A(X) = C(B(X))

- `n` is the number of rotors
- `p[i]` is the rotational position of the i<sup>th</sup> rotor (0 <= `p[i]` < 26)
- `ROT(n)` is a [caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher)
- `P`, `T`, `U` are [monoalphabetic ciphers](https://en.wikipedia.org/wiki/Substitution_cipher)
- `U` must have the following property: `U * U^-1 = I`

[`MinimalistEnigma`](https://github.com/hasyimibhar/enigma/blob/master/enigma/minimalist.go) demonstrates how [`Enigma`](https://github.com/hasyimibhar/enigma/blob/master/enigma/enigma.go) can be reimplemented using algebraic notation.

## Sample usage:

```go
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

	plaintext := "helloworld"
	ciphertext := machine.TransformString(plaintext)
	fmt.Printf("%s -> %s\n", plaintext, ciphertext)
}
```

Output:
```
helloworld -> nsdskqdugk
```
