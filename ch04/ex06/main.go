package main

import (
	"unicode"
	"unicode/utf8"
)

func main() {

	// https://golang.org/pkg/unicode/
	// unicode white space
	// '\t', '\n', '\v', '\f', '\r', ' ',

	input := "\nHello\t \rworld \v\f"
	b := []byte(input)
	var result []byte

	for i := 0; i < len(b); i += s {
		r, s := utf8.DecodeRune(b[i:])

		// unicode space判定
		if unicode.IsSpace(r) {
			// 普通のspaceだったら無視
			result = append(result, []byte(" "))
		} else {
			result = append(result, b[i:i+s])
		}
	}

}
