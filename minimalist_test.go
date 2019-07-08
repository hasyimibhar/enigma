package enigma

import (
	"testing"
)

func TestMinimalistEnigma_Deterministic(t *testing.T) {
	minimalist := &MinimalistEnigma{
		P: newPlugboard("EJ OY IV AQ KW FX MT PS LU BD"),
		R: []*MinimalistEnigmaRotor{
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "ESOVPZJAYQUIRHXLNFTGKDCMWB"), Notches: notches("J"), Position: 14},
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "AJDKSIRUXBLHWTMCQGZNPYFVOE"), Notches: notches("E"), Position: 22},
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "VZBRGITYUPSDNHLXAWMJQOFECK"), Notches: notches("Z"), Position: 25},
		},
		U: newSubstituteCipher(t, "EJMZALYXVBWFCRQUONTSPIKHGD"),
	}

	clone := minimalist.Clone()

	plaintext := "loremipsumdolorsitametconsecteturadipiscingelitseddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquautenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatduisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturexcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum"
	ciphertext := minimalist.EncryptString(plaintext)

	actual := clone.EncryptString(plaintext)

	if actual != ciphertext {
		t.Fatal("minimalist enigma machine is not deterministic")
	}
}

func TestMinimalistEnigma_SelfReciprocal(t *testing.T) {
	minimalist := &MinimalistEnigma{
		P: newPlugboard("EJ OY IV AQ KW FX MT PS LU BD"),
		R: []*MinimalistEnigmaRotor{
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "ESOVPZJAYQUIRHXLNFTGKDCMWB"), Notches: notches("J"), Position: 14},
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "AJDKSIRUXBLHWTMCQGZNPYFVOE"), Notches: notches("E"), Position: 22},
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "VZBRGITYUPSDNHLXAWMJQOFECK"), Notches: notches("Z"), Position: 25},
		},
		U: newSubstituteCipher(t, "EJMZALYXVBWFCRQUONTSPIKHGD"),
	}

	clone := minimalist.Clone()

	plaintext := "loremipsumdolorsitametconsecteturadipiscingelitseddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquautenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatduisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturexcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum"
	ciphertext := minimalist.EncryptString(plaintext)

	actual := clone.EncryptString(ciphertext)

	if actual != plaintext {
		t.Fatal("minimalist enigma machine is not self-reciprocal")
	}
}

func TestMinimalistEnigma_Equal(t *testing.T) {
	machine, err := NewWithConfig(&Config{
		Plugboard: "EJ OY IV AQ KW FX MT PS LU BD",
		Rotors: []*RotorConfig{
			&RotorConfig{Name: "IV", Position: 14},
			&RotorConfig{Name: "II", Position: 22},
			&RotorConfig{Name: "V", Position: 25},
		},
		Reflector: &RotorConfig{Name: "A"},
	})
	if err != nil {
		t.Fatal(err)
	}

	minimalist := &MinimalistEnigma{
		P: newPlugboard("EJ OY IV AQ KW FX MT PS LU BD"),
		R: []*MinimalistEnigmaRotor{
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "ESOVPZJAYQUIRHXLNFTGKDCMWB"), Notches: notches("J"), Position: 14},
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "AJDKSIRUXBLHWTMCQGZNPYFVOE"), Notches: notches("E"), Position: 22},
			&MinimalistEnigmaRotor{T: newSubstituteCipher(t, "VZBRGITYUPSDNHLXAWMJQOFECK"), Notches: notches("Z"), Position: 25},
		},
		U: newSubstituteCipher(t, "EJMZALYXVBWFCRQUONTSPIKHGD"),
	}

	plaintext := "loremipsumdolorsitametconsecteturadipiscingelitseddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquautenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatduisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturexcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum"
	machineCiphertext := machine.EncryptString(plaintext)
	minimalistCiphertext := minimalist.EncryptString(plaintext)

	if machineCiphertext != minimalistCiphertext {
		t.Fatal("EM is not equivalent to MEM")
	}
}

func newSubstituteCipher(t *testing.T, table string) SubstituteCipher {
	t.Helper()

	cipher, err := NewSubstituteCipherWithTable(table)
	if err != nil {
		t.Fatal(err)
	}

	return cipher
}
