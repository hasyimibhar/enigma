package enigma

import (
	"strings"
)

type Alphabet byte

type Enigma struct {
	Plugboard SubstituteCipher
	Rotors    []*Rotor
}

type RotorConfig struct {
	Name     string
	Position int
}

type Config struct {
	Plugboard string
	Rotors    []*RotorConfig
	Reflector *RotorConfig
}

func NewWithConfig(cfg *Config) (*Enigma, error) {
	plugboard := newPlugboard(cfg.Plugboard)

	rotors := []*Rotor{}
	for _, r := range cfg.Rotors {
		rotor, err := NewRotorWithName(r.Position, r.Name)
		if err != nil {
			return nil, err
		}

		rotors = append(rotors, rotor)
	}

	reflector, err := NewReflectorWithName(cfg.Reflector.Position, cfg.Reflector.Name)
	if err != nil {
		return nil, err
	}

	rotors = append(rotors, reflector)

	return &Enigma{
		Plugboard: plugboard,
		Rotors:    rotors,
	}, nil
}

func (e *Enigma) Transform(from Alphabet) Alphabet {
	to := e.transform().Transform(from)

	for i := 0; i < len(e.Rotors); i++ {
		turnover := e.Rotors[i].Rotate()
		if !turnover {
			break
		}
	}

	return to
}

func (e *Enigma) TransformString(from string) string {
	to := []byte{}
	for _, alph := range []byte(from) {
		to = append(to, byte(e.Transform(Alphabet(alph))))
	}
	return string(to)
}

func (e *Enigma) Clone() *Enigma {
	return &Enigma{
		Plugboard: e.Plugboard.Clone().(SubstituteCipher),
		Rotors:    RotorList(e.Rotors).Clone().(RotorList),
	}
}

func (e *Enigma) transform() Transformation {
	return CombineTransformations(
		e.Plugboard,
		RotorList(e.Rotors),
		RotorList(e.Rotors[:len(e.Rotors)-1]).Inverse(),
		e.Plugboard.Inverse(),
	)
}

func newPlugboard(wiring string) SubstituteCipher {
	wiring = strings.ToLower(wiring)
	pairs := strings.Split(wiring, " ")

	plugboard := NewIdentitySubstituteCipher()

	for _, pair := range pairs {
		plugboard[Alphabet(pair[0])] = Alphabet(pair[1])
		plugboard[Alphabet(pair[1])] = Alphabet(pair[0])
	}

	return plugboard
}
