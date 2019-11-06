package word

import (
	"math/rand"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1003))
		if r == rune(0x1002) {
			r = rune(' ')
		}
		if r == rune(0x1001) {
			r = rune(',')
		}
		if r == rune(0x1000) {
			r = rune('.')
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
