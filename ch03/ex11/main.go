package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func getRuneAt(s string, i int) rune {
	rs := []rune(s)
	return rs[i]
}

func main() {
	s := os.Args[1]
	pm := ""
	if s[0] == '-' || s[0] == '+' {
		pm = string(getRuneAt(s, 0))
		s = s[1:]
	}
	slice := strings.Split(s, ".")
	decimal := ""
	if len(slice) == 2 {
		decimal = "." + slice[1]
	} else if len(slice) > 2 {
		return
	}
	s = slice[0]
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
	fmt.Printf("answer: %s", pm+buf.String()+decimal)
}
