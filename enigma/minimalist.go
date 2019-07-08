package enigma

type MinimalistEnigmaRotor struct {
	T        SubstituteCipher
	Position int
	Notches  []int
}

type MinimalistEnigma struct {
	P SubstituteCipher
	R []*MinimalistEnigmaRotor
	U SubstituteCipher
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
		tfs = append(tfs, CombineTransformations(
			CaesarCipher(e.R[i].Position),
			e.R[i].T,
			CaesarCipher(e.R[i].Position).Inverse(),
		))
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
