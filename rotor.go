package enigma

import (
	"fmt"
	"strings"
)

type Rotor struct {
	wiring    SubstituteCipher
	position  CaesarCipher
	notches   []int
	rotatable bool
	inversed  bool
}

func NewRotor(position int, notches []int) (*Rotor, error) {
	return NewRotorWithWiringTable(position, notches, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
}

func NewReflectorWithWiringTable(position int, wiringTable string, rotatable bool) (*Rotor, error) {
	rotor, err := NewRotorWithWiringTable(position, []int{}, wiringTable)
	if err != nil {
		return nil, err
	}

	rotor.rotatable = rotatable

	for alph := Alphabet('a'); alph <= Alphabet('z'); alph++ {
		image := rotor.wiring.Transform(alph)
		if alph != rotor.wiring.Transform(image) {
			return nil, fmt.Errorf("rotor: invalid reflector wiring table: '%c' * T != '%c' * T", alph, image)
		}
	}

	return rotor, nil
}

func NewRotorWithWiringTable(position int, notches []int, wiringTable string) (*Rotor, error) {
	if !withinRange(position, 0, 25) {
		return nil, fmt.Errorf("position must be in range 0 <= position <= 25")
	}

	seen := map[int]bool{}
	for _, notch := range notches {
		if !withinRange(notch, 0, 25) {
			return nil, fmt.Errorf("notch must be in range 0 <= position <= 25")
		}
		if seen[notch] {
			return nil, fmt.Errorf("notch must be unique")
		}

		seen[notch] = true
	}

	substitute, err := NewSubstituteCipherWithTable(wiringTable)
	if err != nil {
		return nil, fmt.Errorf("rotor: %s", err)
	}

	return &Rotor{
		wiring:    substitute,
		position:  CaesarCipher(position),
		notches:   notches,
		rotatable: true,
	}, nil
}

func NewRotorWithName(position int, name string) (rotor *Rotor, err error) {
	switch name {
	case "I":
		rotor, err = NewRotorWithWiringTable(position, notches("Q"), "EKMFLGDQVZNTOWYHXUSPAIBRCJ")
	case "II":
		rotor, err = NewRotorWithWiringTable(position, notches("E"), "AJDKSIRUXBLHWTMCQGZNPYFVOE")
	case "III":
		rotor, err = NewRotorWithWiringTable(position, notches("V"), "BDFHJLCPRTXVZNYEIWGAKMUSQO")
	case "IV":
		rotor, err = NewRotorWithWiringTable(position, notches("J"), "ESOVPZJAYQUIRHXLNFTGKDCMWB")
	case "V":
		rotor, err = NewRotorWithWiringTable(position, notches("Z"), "VZBRGITYUPSDNHLXAWMJQOFECK")
	case "VI":
		rotor, err = NewRotorWithWiringTable(position, notches("ZM"), "JPGVOUMFYQBENHZRDKASXLICTW")
	case "VII":
		rotor, err = NewRotorWithWiringTable(position, notches("ZM"), "NZJHGRCXMYSWBOUFAIVLPEKQDT")
	case "VIII":
		rotor, err = NewRotorWithWiringTable(position, notches("ZM"), "FKQHTLXOCBJSPDZRAMEWNIUYGV")
	default:
		err = fmt.Errorf("invalid rotor name: %s", name)
	}

	return
}

func NewReflectorWithName(position int, name string) (rotor *Rotor, err error) {
	switch name {
	case "A":
		rotor, err = NewReflectorWithWiringTable(position, "EJMZALYXVBWFCRQUONTSPIKHGD", false)
	case "B":
		rotor, err = NewReflectorWithWiringTable(position, "YRUHQSLDPXNGOKMIEBFZCWVJAT", false)
	case "C":
		rotor, err = NewReflectorWithWiringTable(position, "FVPJIAOYEDRZXWGCTKUQSBNMHL", false)
	case "B-Thin":
		rotor, err = NewReflectorWithWiringTable(position, "ENKQAUYWJICOPBLMDXZVFTHRGS", false)
	case "C-Thin":
		rotor, err = NewReflectorWithWiringTable(position, "RDOBJNTKVEHMLFCWZAXGYIPSUQ", false)
	default:
		err = fmt.Errorf("invalid reflector name: %s", name)
	}

	return
}

// Rotate rotates the rotor. It returns true if
// it reaches its kick position.
func (r *Rotor) Rotate() bool {
	if !r.rotatable {
		return false
	}

	r.position = CaesarCipher((int(r.position) + 1) % 26)

	for _, notch := range r.notches {
		if int(r.position) == notch {
			return true
		}
	}

	return false
}

func (r *Rotor) Transform(from Alphabet) Alphabet {
	tf := CombineTransformations(r.position, r.wiring, r.position.Inverse())
	if r.inversed {
		tf = tf.Inverse()
	}

	return tf.Transform(from)
}

func (r *Rotor) Inverse() Transformation {
	return &Rotor{
		wiring:    r.wiring.Clone().(SubstituteCipher),
		position:  r.position.Clone().(CaesarCipher),
		notches:   intList(r.notches).Clone(),
		rotatable: r.rotatable,
		inversed:  !r.inversed,
	}
}

func (r *Rotor) Clone() Transformation {
	return &Rotor{
		wiring:    r.wiring.Clone().(SubstituteCipher),
		position:  r.position.Clone().(CaesarCipher),
		notches:   intList(r.notches).Clone(),
		rotatable: r.rotatable,
		inversed:  r.inversed,
	}
}

func withinRange(x, min, max int) bool {
	return x >= min && x <= max
}

type RotorList []*Rotor

func (l RotorList) Transform(from Alphabet) Alphabet {
	to := from
	for _, tf := range l {
		to = tf.Transform(to)
	}
	return to
}

func (l RotorList) Clone() Transformation {
	rotors := []*Rotor{}
	for _, r := range l {
		rotors = append(rotors, r.Clone().(*Rotor))
	}
	return RotorList(rotors)
}

func (l RotorList) Inverse() Transformation {
	rotors := []*Rotor{}
	for i := len(l) - 1; i >= 0; i-- {
		rotors = append(rotors, l[i].Inverse().(*Rotor))
	}

	return RotorList(rotors)
}

func notches(s string) []int {
	s = strings.ToLower(s)
	notches := []int{}

	for _, alph := range []byte(s) {
		notches = append(notches, int(alph-'a'))
	}

	return notches
}

type intList []int

func (l intList) Clone() []int {
	clone := []int{}
	for _, i := range l {
		clone = append(clone, i)
	}
	return clone
}
