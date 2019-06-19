package main

import (
	"fmt"
	"io"
	"os"
)

// ex) go run main.go src

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "input file is needed.")
	}
	src := os.Args[1]
	dst := os.Stdout

	f, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.Copy(dst, f)
	if err != nil {
		panic(err)
	}
}
