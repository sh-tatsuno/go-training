package popcount

import "testing"

func TestPopCount3(t *testing.T) {
	for i := 0; i < 1024; i++ {
		expected := PopCount(uint64(i))
		actual := PopCount3(uint64(i))
		if actual != expected {
			t.Fatalf("mismatch result in i=%d. actual: %v, expected: %v", i, actual, expected)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount3(uint64(i))
	}
}
