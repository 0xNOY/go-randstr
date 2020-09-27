package randstr

import (
	errors "errors"
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

var (
	// ErrZeroLengthLetters is error.
	ErrZeroLengthLetters = errors.New("srcLetters must not be zero in length")
)

// RandStr is random string generater.
type RandStr struct {
	// srcLetters is source of generated random string.
	srcLetters []rune
	randModule *rand.Rand
}

// New returns new RandStr.
func New(src rand.Source, letters []rune) *RandStr {
	return &RandStr{
		randModule: rand.New(src),
		srcLetters: letters,
	}
}

// Gen generates random string.
func (rs *RandStr) Gen(length uint) string {
	l := len(rs.srcLetters)
	r := make([]rune, length)
	for i := uint(0); i < length; i++ {
		r[i] = rs.srcLetters[rs.randModule.Intn(l)]
	}
	return string(r)
}

// GetSrcLetters gets srcLetters.
func (rs *RandStr) GetSrcLetters() []rune {
	return rs.srcLetters
}

// SetSrcLetters sets srcLetters.
func (rs *RandStr) SetSrcLetters(letters []rune) error {
	if len(letters) <= 0 {
		return ErrZeroLengthLetters
	}

	rs.srcLetters = letters
	return nil
}

// GetRandModule gets randModule.
func (rs *RandStr) GetRandModule() *rand.Rand {
	return rs.randModule
}

// SetRandModule sets randModule.
func (rs *RandStr) SetRandModule(newRandModule *rand.Rand) {
	rs.randModule = newRandModule
}
