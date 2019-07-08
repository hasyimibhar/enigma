package enigma

import (
	"fmt"
	"strings"
)

type SubstituteCipher map[Alphabet]Alphabet

func NewIdentitySubstituteCipher() SubstituteCipher {
	alphMap := map[Alphabet]Alphabet{}
	for alph := Alphabet('a'); alph <= Alphabet('z'); alph++ {
		alphMap[alph] = alph
	}
	return alphMap
}

func NewSubstituteCipherWithTable(table string) (SubstituteCipher, error) {
	alphMap := map[Alphabet]Alphabet{}
	seen := map[Alphabet]bool{}

	if len(table) != 26 {
		return SubstituteCipher{}, fmt.Errorf("invalid wiring table: length must be 26")
	}

	table = strings.ToLower(table)

	for i := 0; i < 26; i++ {
		from := Alphabet('a' + i)
		to := Alphabet(table[i])

		if to < 'a' || to > 'z' {
			return SubstituteCipher{}, fmt.Errorf("invalid wiring table: invalid character '%c' at position %d", to, i)
		}

		if seen[to] {
			return SubstituteCipher{}, fmt.Errorf("invalid wiring table: repeated '%c' at position %d", to, i)
		}

		seen[to] = true
		alphMap[from] = to
	}

	return alphMap, nil
}

func (s SubstituteCipher) Transform(from Alphabet) (to Alphabet) {
	to = s[from]
	return
}

func (s SubstituteCipher) Clone() Transformation {
	clone := map[Alphabet]Alphabet{}
	for k, v := range s {
		clone[k] = v
	}
	return SubstituteCipher(clone)
}

func (s SubstituteCipher) Inverse() Transformation {
	inverse := map[Alphabet]Alphabet{}
	for k, v := range s {
		inverse[v] = k
	}
	return SubstituteCipher(inverse)
}
