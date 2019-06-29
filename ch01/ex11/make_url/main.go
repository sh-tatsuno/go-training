package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "input file is needed.")
	}
	inputURL := os.Args[1]

	s := "go run main.go"
	for i := 0; i < 100000; i++ {
		s += " " + inputURL
	}

	file, err := os.Create("../fetch/runner.sh")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write(([]byte)(s))
}
