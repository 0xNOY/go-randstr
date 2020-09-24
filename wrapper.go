package randstr

import (
	rand "math/rand"
	time "time"
)

var (
	// Letters is letters used generating random string.
	Letters = []rune(StrLettersLower + StrLettersUpper + StrLettersNum)
	randStr = New(
		rand.NewSource(time.Now().UnixNano()),
		Letters,
	)
)

// Gen generates random string.
func Gen(length uint) string {
	syncLetters()
	return randStr.Gen(length)
}

// GetRandModule gets randModule.
func GetRandModule() *rand.Rand {
	return randStr.GetRandModule()
}

// SetRandModule sets randModule.
func SetRandModule(newRandModule *rand.Rand) {
	randStr.SetRandModule(newRandModule)
}

func syncLetters() {
	randStr.Letters = Letters
}
