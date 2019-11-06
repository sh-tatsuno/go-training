package word

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)

		// check if isPalindrome
		if !IsPalindrome(p) {
			t.Errorf("Is Palindrome(%q) = false", p)
		}
	}
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)

		// check if not isNotPalindrome
		if !IsNotPalindrome(p) {
			t.Errorf("Is Not Palindrome(%q) = true", p)
		}
	}
}
