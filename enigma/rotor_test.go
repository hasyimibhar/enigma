package enigma

import (
	"strings"
	"testing"
)

func TestRotor(t *testing.T) {
	// Loop through all possible initial positions
	for ip := 0; ip < 26; ip++ {
		rotor, err := NewRotor(ip, []int{0})
		if err != nil {
			t.Fatal(err)
		}

		// Test 2 full rotations
		for i := 0; i < 26*2; i++ {
			for alph := Alphabet('a'); alph <= Alphabet('z'); alph++ {
				actual := rotor.Transform(alph)
				expected := Alphabet('a' + ((ip + int(alph-'a') + i) % 26))
				if actual != expected {
					t.Fatalf("expecting '%c' * R(%d) => '%c', got '%c' instead", alph, i, actual, expected)
				}
			}

			rotor.Rotate()
		}
	}
}

func TestRotor_WithWiringTable(t *testing.T) {
	tests := []struct {
		Name        string
		WiringTable string
	}{
		{"IC", "DMTWSILRUYQNKFEJCAZBPGXOHV"},
		{"IIC", "HQZGPJTMOBLNCIFDYAWVEUSRKX"},
		{"IIIC", "UQNTLSZFMREHDPXKIBVYGJCWOA"},

		{"I", "JGDQOXUSCAMIFRVTPNEWKBLZYH"},
		{"II", "NTZPSFBOKMWRCJDIVLAEYUXHGQ"},
		{"III", "JVIUBHTCDYAKEQZPOSGXNRMWFL"},
		{"IV", "QYHOGNECVPUZTFDJAXWMKISRBL"},
		{"V", "QWERTZUIOASDFGHJKPYXCVBNML"},

		{"I-K", "PEZUOHXSCVFMTBGLRINQJWAYDK"},
		{"II-K", "ZOUESYDKFWPCIQXHMVBLGNJRAT"},
		{"III-K", "EHRVXGAOBQUSIMZFLYNWKTPDJC"},
		{"UKW-K", "IMETCGFRAYSQBZXWLHKDVUPOJN"},
		{"ETW-K", "QWERTZUIOASDFGHJKPYXCVBNML"},

		{"Enigma I/I", "EKMFLGDQVZNTOWYHXUSPAIBRCJ"},
		{"Enigma I/II", "AJDKSIRUXBLHWTMCQGZNPYFVOE"},
		{"Enigma I/III", "BDFHJLCPRTXVZNYEIWGAKMUSQO"},
		{"M3 Army/IV", "ESOVPZJAYQUIRHXLNFTGKDCMWB"},
		{"M3 Army/V", "VZBRGITYUPSDNHLXAWMJQOFECK"},
		{"M3 & M4 Naval/VI", "JPGVOUMFYQBENHZRDKASXLICTW"},
		{"M3 & M4 Naval/VII", "NZJHGRCXMYSWBOUFAIVLPEKQDT"},
		{"M3 & M4 Naval/VIII", "FKQHTLXOCBJSPDZRAMEWNIUYGV"},

		{"Beta", "LEYJVCNIXWPBQMDRTAKZGFUHOS"},
		{"Gamma", "FSOKANUERHMBTIYCWLQPZXVGJD"},
		{"Reflector A", "EJMZALYXVBWFCRQUONTSPIKHGD"},
		{"Reflector B", "YRUHQSLDPXNGOKMIEBFZCWVJAT"},
		{"Reflector C", "FVPJIAOYEDRZXWGCTKUQSBNMHL"},
		{"Reflector B Thin", "ENKQAUYWJICOPBLMDXZVFTHRGS"},
		{"Reflector C Thin", "RDOBJNTKVEHMLFCWZAXGYIPSUQ"},
		{"Enigma I/ETW", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Loop through all possible initial positions
			for ip := 0; ip < 26; ip++ {
				rotor, err := NewRotorWithWiringTable(ip, []int{0}, tt.WiringTable)
				if err != nil {
					t.Fatal(err)
				}

				wiringTable := strings.ToLower(tt.WiringTable)

				// Test 2 full rotations
				for i := 0; i < 26*2; i++ {
					for alph := Alphabet('a'); alph <= Alphabet('z'); alph++ {
						actual := rotor.Transform(alph)
						expected := Alphabet(wiringTable[(int(alph-'a')+ip+i)%26])
						if actual != expected {
							t.Fatalf("expecting '%c' * R['%s'](%d) => '%c', got '%c' instead", alph, tt.Name, i, actual, expected)
						}
					}

					rotor.Rotate()
				}
			}
		})
	}
}

func TestReflector(t *testing.T) {
	tests := []struct {
		Name        string
		WiringTable string
	}{
		{"Reflector A", "EJMZALYXVBWFCRQUONTSPIKHGD"},
		{"Reflector B", "YRUHQSLDPXNGOKMIEBFZCWVJAT"},
		{"Reflector C", "FVPJIAOYEDRZXWGCTKUQSBNMHL"},
		{"Reflector B Thin", "ENKQAUYWJICOPBLMDXZVFTHRGS"},
		{"Reflector C Thin", "RDOBJNTKVEHMLFCWZAXGYIPSUQ"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			// Loop through all possible initial positions
			for ip := 0; ip < 26; ip++ {
				rotor, err := NewReflectorWithWiringTable(ip, tt.WiringTable, true)
				if err != nil {
					t.Fatal(err)
				}

				wiringTable := strings.ToLower(tt.WiringTable)

				// Test 2 full rotations
				for i := 0; i < 26*2; i++ {
					for alph := Alphabet('a'); alph <= Alphabet('z'); alph++ {
						actual := rotor.Transform(alph)
						expected := Alphabet(wiringTable[(int(alph-'a')+ip+i)%26])
						if actual != expected {
							t.Fatalf("expecting '%c' * R['%s'](%d) => '%c', got '%c' instead", alph, tt.Name, i, actual, expected)
						}
					}

					rotor.Rotate()
				}
			}
		})
	}
}

func TestRotor_Rotate(t *testing.T) {
	rotor, err := NewRotor(0, []int{14, 3, 23})
	if err != nil {
		t.Fatal(err)
	}

	for i := 1; i < 26; i++ {
		turnover := rotor.Rotate()
		if (i == 3 || i == 14 || i == 23) && !turnover {
			t.Fatalf("rotor should turnover at position %d", i)
		} else if !(i == 3 || i == 14 || i == 23) && turnover {
			t.Fatalf("rotor should not turnover at position %d", i)
		}
	}
}
