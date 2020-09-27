package randstr

import (
	rand "math/rand"
	time "time"
)

var (
	randStr = New(
		rand.NewSource(time.Now().UnixNano()),
		[]rune(LettersLowerStr+LettersUpperStr+LettersNumStr),
	)
)

// Gen generates random string.
func Gen(length uint) string {
	return randStr.Gen(length)
}

// GetSrcLetters gets srcLetters.
func GetSrcLetters() []rune {
	return randStr.srcLetters
}

// SetSrcLetters sets srcLetters.
func SetSrcLetters(letters []rune) error {
	if len(letters) <= 0 {
		return ErrZeroLengthLetters
	}

	randStr.srcLetters = letters
	return nil
}

// GetRandModule gets randModule.
func GetRandModule() *rand.Rand {
	return randStr.GetRandModule()
}

// SetRandModule sets randModule.
func SetRandModule(newRandModule *rand.Rand) {
	randStr.SetRandModule(newRandModule)
}
