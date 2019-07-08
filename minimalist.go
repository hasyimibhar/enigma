package enigma

// MinimalistEnigma is an implementation of the Enigma Machine using
// algebraic notation. The encryption of a single alphabet can be
// expressed as:
//
//     E = P * SUM(R[i], n) * U * SUM(R[i], n)^-1 * P^-1
//
// where:
//
//     R = ROT(p) * T * ROT(-p)
//     and
//     P, T, U are monoalphabetic ciphers.
//
type MinimalistEnigma struct {
	P SubstituteCipher
	R []*MinimalistEnigmaRotor
	U SubstituteCipher
}

type MinimalistEnigmaRotor struct {
	T        SubstituteCipher
	Position int
	Notches  []int
}

func (e *MinimalistEnigma) Transform(from Alphabet) Alphabet {
	// E = P * R * U * R^-1 * P^-1
	E := CombineTransformations(
		e.P, e.rotors(), e.U, e.rotors().Inverse(), e.P.Inverse(),
	)

	to := E.Transform(from)

	// Rotate the rotors
	for i := 0; i < len(e.R); i++ {
		e.R[i].Position = (e.R[i].Position + 1) % 26

		var turnover bool
		for _, n := range e.R[i].Notches {
			if e.R[i].Position == n {
				turnover = true
				break
			}
		}

		if !turnover {
			break
		}
	}

	return to
}

func (e *MinimalistEnigma) TransformString(from string) string {
	to := []byte{}
	for _, alph := range []byte(from) {
		to = append(to, byte(e.Transform(Alphabet(alph))))
	}
	return string(to)
}

func (e *MinimalistEnigma) Clone() *MinimalistEnigma {
	return &MinimalistEnigma{
		P: e.P.Clone().(SubstituteCipher),
		R: minimalistEnigmaRotorList(e.R).Clone(),
		U: e.U.Clone().(SubstituteCipher),
	}
}

func (e *MinimalistEnigma) rotors() Transformation {
	tfs := []Transformation{}

	for i := 0; i < len(e.R); i++ {
		// R = p^n * T * p^n^-1
		R := CombineTransformations(
			CaesarCipher(e.R[i].Position),
			e.R[i].T,
			CaesarCipher(e.R[i].Position).Inverse(),
		)

		tfs = append(tfs, R)
	}

	return CombineTransformations(tfs...)
}

func (r *MinimalistEnigmaRotor) Clone() *MinimalistEnigmaRotor {
	return &MinimalistEnigmaRotor{
		T:        r.T.Clone().(SubstituteCipher),
		Position: r.Position,
		Notches:  intList(r.Notches).Clone(),
	}
}

type minimalistEnigmaRotorList []*MinimalistEnigmaRotor

func (rl minimalistEnigmaRotorList) Clone() []*MinimalistEnigmaRotor {
	clone := []*MinimalistEnigmaRotor{}
	for _, r := range rl {
		clone = append(clone, r.Clone())
	}
	return clone
}
