package enigma

import (
	"strings"
	"testing"
)

func TestSubstituteCipher_Identity(t *testing.T) {
	ident := NewIdentitySubstituteCipher()

	for alph := Alphabet('a'); alph <= Alphabet('z'); alph++ {
		x := ident.Transform(alph)
		if x != alph {
			t.Fatalf("expecting '%c' * I => '%c', got '%c' instead", alph, alph, x)
		}
	}
}

func TestSubstituteCipher_CustomWiring(t *testing.T) {
	ident := NewIdentitySubstituteCipher()

	for alph := Alphabet('a'); alph <= Alphabet('z'); alph++ {
		ident[alph] = 'z' - (alph - 'a')
	}

	tests := []struct {
		From Alphabet
		To   Alphabet
	}{
		{'a', 'z'},
		{'b', 'y'},
		{'c', 'x'},
		{'d', 'w'},
		{'e', 'v'},
		{'f', 'u'},
		{'g', 't'},
		{'h', 's'},
		{'i', 'r'},
		{'j', 'q'},
		{'k', 'p'},
		{'l', 'o'},
		{'m', 'n'},
		{'n', 'm'},
		{'o', 'l'},
		{'p', 'k'},
		{'q', 'j'},
		{'r', 'i'},
		{'s', 'h'},
		{'t', 'g'},
		{'u', 'f'},
		{'v', 'e'},
		{'w', 'd'},
		{'x', 'c'},
		{'y', 'b'},
		{'z', 'a'},
	}

	for _, tt := range tests {
		t.Run(string(tt.From), func(t *testing.T) {
			x := ident.Transform(tt.From)
			if x != tt.To {
				t.Fatalf("expecting '%c' * T => '%c', got '%c' instead", tt.From, tt.To, x)
			}
		})
	}
}

func TestSubstituteCipher_KnownRotors(t *testing.T) {
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
			cipher, err := NewSubstituteCipherWithTable(tt.WiringTable)
			if err != nil {
				t.Fatal(err)
			}

			table := strings.ToLower(tt.WiringTable)
			for i := 0; i < 26; i++ {
				from := Alphabet('a' + i)
				expected := Alphabet(table[i])
				actual := cipher.Transform(from)

				if expected != actual {
					t.Fatalf("expecting '%c' * T('%s') => '%c', got '%c' instead", from, tt.Name, expected, actual)
				}
			}
		})
	}
}

func TestSubstituteCipher_Clone(t *testing.T) {
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
			cipher, err := NewSubstituteCipherWithTable(tt.WiringTable)
			if err != nil {
				t.Fatal(err)
			}

			clone := cipher.Clone()

			for i := 0; i < 26; i++ {
				from := Alphabet('a' + i)
				expected := cipher.Transform(from)
				actual := clone.Transform(from)

				if expected != actual {
					t.Fatalf("expecting '%c' * T('%s') => '%c' * T('%s'-clone), got '%c' instead",
						from, tt.Name, expected, tt.Name, actual)
				}
			}
		})
	}
}

func TestSubstituteCipher_InvalidWiringTable(t *testing.T) {
	tests := []struct {
		Name        string
		WiringTable string
		Error       string
	}{
		{"invalid length", "DMTWSILRUYQNKFEJCAZBPGXOH", "invalid wiring table: length must be 26"},
		{"repeated character", "DMTWSILRUYQNKFEJCAZBPGXOHM", "invalid wiring table: repeated 'm' at position 25"},
		{"invalid character", "DMTWSIL!UYQNKFEJCAZBPGXOHV", "invalid wiring table: invalid character '!' at position 7"},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			_, err := NewSubstituteCipherWithTable(tt.WiringTable)
			if err == nil || err.Error() != tt.Error {
				t.Fatalf("expecting error '%s', got '%s' instead", tt.Error, err)
			}
		})
	}
}
