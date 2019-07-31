package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

// initilization -> 2.6.2
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func countDiffBitSha256Hash(h1, h2 [32]byte) int {

	sum := 0
	for i := 0; i < 32; i++ {
		sum += int(pc[h1[i]^h2[i]])
	}

	return sum
}

func main() {
	h1 := sha256.Sum256([]byte("aaa"))
	h2 := sha256.Sum256([]byte("bbb"))
	fmt.Printf("diff bit num: %d\n", countDiffBitSha256Hash(h1, h2))
}
