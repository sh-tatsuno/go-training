package main

import (
	"bufio"
	"fmt"
	"os"
)

// ex)
// go run dup.go testdata/01.txt testdata/02.txt

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, -1, counts)
	} else {
		for i, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup err: %v\n", err)
				continue
			}
			countLines(f, i, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, i int, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if i < 0 {
			counts[input.Text()]++
		} else {
			counts[os.Args[i+1]+"\t"+input.Text()]++
		}
	}
}
