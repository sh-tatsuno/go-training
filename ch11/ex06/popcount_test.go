package popcount

import (
	"testing"
)

func benchmarkPopCount(b *testing.B, n uint64) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(n))
	}
}

func benchmarkPopCount3(b *testing.B, n uint64) {
	for i := 0; i < b.N; i++ {
		PopCount3(uint64(n))
	}
}

func benchmarkPopCount4(b *testing.B, n uint64) {
	for i := 0; i < b.N; i++ {
		PopCount4(uint64(n))
	}
}

func BenchmarkPopCount_0(b *testing.B)  { benchmarkPopCount(b, 0x0) }
func BenchmarkPopCount3_0(b *testing.B) { benchmarkPopCount3(b, 0x0) }
func BenchmarkPopCount4_0(b *testing.B) { benchmarkPopCount4(b, 0x0) }

func BenchmarkPopCount_F(b *testing.B)  { benchmarkPopCount(b, 0xFFFFFFFF) }
func BenchmarkPopCount3_F(b *testing.B) { benchmarkPopCount3(b, 0xFFFFFFFF) }
func BenchmarkPopCount4_F(b *testing.B) { benchmarkPopCount4(b, 0xFFFFFFFF) }

func BenchmarkPopCount_2F(b *testing.B)  { benchmarkPopCount(b, 0xFFFFFFFFFFFFFFFF) }
func BenchmarkPopCount3_2F(b *testing.B) { benchmarkPopCount3(b, 0xFFFFFFFFFFFFFFFF) }
func BenchmarkPopCount4_2F(b *testing.B) { benchmarkPopCount4(b, 0xFFFFFFFFFFFFFFFF) }
