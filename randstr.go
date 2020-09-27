package randstr

import (
	rand "math/rand"
)

const (
	// LettersNumStr is numbers from 0 to 9.
	LettersNumStr = "0123456789"
	// LettersUpperStr is upper case from A to Z.
	LettersUpperStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// LettersLowerStr is lower case from a to z.
	LettersLowerStr = "abcdefghijklmnopqrstuvwxyz"
)

// RandStr is random string generater.
type RandStr struct {
	// Letters is source of generated random string.
	Letters    []rune
	randModule *rand.Rand
}

// New returns new RandStr.
func New(src rand.Source, letters []rune) *RandStr {
	return &RandStr{
		randModule: rand.New(src),
		Letters:    letters,
	}
}

// Gen generates random string.
func (rs *RandStr) Gen(length uint) string {
	l := len(rs.Letters)
	r := make([]rune, length)
	for i := uint(0); i < length; i++ {
		r[i] = rs.Letters[rs.randModule.Intn(l)]
	}
	return string(r)
}

// GetRandModule gets randModule.
func (rs *RandStr) GetRandModule() *rand.Rand {
	return rs.randModule
}

// SetRandModule sets randModule.
func (rs *RandStr) SetRandModule(newRandModule *rand.Rand) {
	rs.randModule = newRandModule
}
