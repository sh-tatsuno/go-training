package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	s := os.Args[1]
	var buf bytes.Buffer
	mod := (len(s) + 1) % 3
	for i := 0; i < len(s); i++ {
		mod++
		buf.WriteByte(s[i])
		if mod == 3 && i != (len(s)-1) {
			buf.WriteByte(',')
			mod = 0
		}
	}
	fmt.Printf("answer: %s", buf.String())
}
