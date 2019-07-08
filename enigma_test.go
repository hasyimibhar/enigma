package enigma

import (
	"testing"
)

func TestEnigma_Deterministic(t *testing.T) {
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

	clone := machine.Clone()

	plaintext := "loremipsumdolorsitametconsecteturadipiscingelitseddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquautenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatduisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturexcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum"
	ciphertext := machine.EncryptString(plaintext)

	actual := clone.EncryptString(plaintext)
	if actual != ciphertext {
		t.Fatal("enigma machine is not deterministic")
	}
}

func TestEnigma_SelfReciprocal(t *testing.T) {
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

	clone := machine.Clone()

	plaintext := "loremipsumdolorsitametconsecteturadipiscingelitseddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquautenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatduisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturexcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum"
	ciphertext := machine.EncryptString(plaintext)

	actual := clone.EncryptString(ciphertext)
	if actual != plaintext {
		t.Fatal("enigma machine is not self-reciprocal")
	}
}
