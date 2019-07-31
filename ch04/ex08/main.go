package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // Unicode文字数
	var utflen [utf8.UTFMax + 1]int // UTF-8長
	invalid := 0                    // 不正なUTF-8文字数

	tags := make(map[string]int) //文字数字などカウント

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++

		// カウント
		if unicode.IsLetter(r) {
			tags["letter"]++
		} else if unicode.IsNumber(r) {
			tags["number"]++
		} else if unicode.IsControl(r) {
			tags["control"]++
		} else if unicode.IsSpace(r) {
			tags["space"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	// 追加
	fmt.Printf("\ntag\tcount\n")
	for t, n := range tags {
		fmt.Printf("%s\t%d\n", t, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
