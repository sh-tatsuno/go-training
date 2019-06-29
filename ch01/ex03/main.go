package main

import (
	"fmt"
	"os"
	"strings"
)

func echoSlow(args []string) string {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	return s
}

func echoFast(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	fmt.Println(echoSlow(os.Args))
	fmt.Println(echoFast(os.Args))
}
