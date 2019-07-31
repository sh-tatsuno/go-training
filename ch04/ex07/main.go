package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	size := 0
	for i := 0; i < len(b); i += size {
		_, size = utf8.DecodeRune(b[i:])
		revInside(b[i : i+size])
	}
	revInside(b)
	return b

}

// byteをひっくり返す
func revInside(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func main() {
	b := []byte("こんにちは")
	fmt.Println(string(reverse(b)))
}
