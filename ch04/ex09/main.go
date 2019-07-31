package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for {
		if !input.Scan() {
			break
		}
		words[input.Text()]++
	}
	fmt.Println("word\tcount")
	for word, n := range words {
		fmt.Printf("%s\t%d\n", word, n)
	}
}
