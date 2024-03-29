# Enigma machine in Go

The [Enigma machine](https://en.wikipedia.org/wiki/Enigma_machine) implemented in Go as [`cipher.Block`](https://golang.org/pkg/crypto/cipher/#Block) for learning purposes. Obviously, don't use this in production.

## Mathematical Analysis

A single encryption can be [expressed](https://en.wikipedia.org/wiki/Enigma_machine#Mathematical_analysis) as:

    E = P * SUM(R[i], n) * U * SUM(R[i], n)^-1 * P^-1
    R[i] = ROT(p[i]) * T * ROT(-p[i])

where:

- `n` is the number of rotors
- `p[i]` is the rotational position of the i<sup>th</sup> rotor (0 <= `p[i]` < 26)
- `ROT(n)` is a [caesar cipher](https://en.wikipedia.org/wiki/Caesar_cipher)

All elements are from the set of [monoalphabetic ciphers](https://en.wikipedia.org/wiki/Substitution_cipher), which defines the group `<G, *>`, where:

- the operator `*` applies 2 ciphers from left to right:
    
        A * B = B(A(X))

- `A * (B * C) = (A * B) * C` (associatve)
- `A * I = A`, where `I` is a cipher that returns the plaintext (identity)
- `A * A^-1 = I`, where `A^-1` is `A` but backwards (inverse)

with an additional constraint:

- `U` must be self-reciprocal, i.e. `U * U = I`

### Self-reciprocal

It can be shown that an Enigma machine is self-reciprocal, where the ciphertext from an Enigma machine `M`, when fed to another Enigma machine with the exact configurations of `M`, will return the plaintext:

	E * E
	= [P * SUM(R[i], n) * U * SUM(R[i], n)^-1 * P^-1] *
	  [P * SUM(R[i], n) * U * SUM(R[i], n)^-1 * P^-1]
	= P * SUM(R[i], n) * U * SUM(R[i], n)^-1 * (P^-1 * P) * SUM(R[i], n) * U * SUM(R[i], n)^-1 * P^-1
	= P * SUM(R[i], n) * U * (SUM(R[i], n)^-1 * SUM(R[i], n)) * U * SUM(R[i], n)^-1 * P^-1
	= P * SUM(R[i], n) * (U * U) * SUM(R[i], n)^-1 * P^-1
	= P * (SUM(R[i], n) * SUM(R[i], n)^-1) * P^-1
	= P * P^-1
	= I

### Implementation

[`MinimalistEnigma`](https://github.com/hasyimibhar/enigma/blob/master/minimalist.go) demonstrates how [`Enigma`](https://github.com/hasyimibhar/enigma/blob/master/enigma.go) can be reimplemented using algebraic notation.
