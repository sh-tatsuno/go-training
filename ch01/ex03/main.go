package main

import (
	"fmt"
	"os"
	"strings"
)

func echoSlow() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echoFast() {
	fmt.Println(strings.Join(os.Args, " "))
}

func main() {
	echoSlow()

	echoFast()
}
