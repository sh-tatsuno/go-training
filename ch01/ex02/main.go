package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(run(os.Args))
}

func run(args []string) string {
	s := ""
	for i, arg := range args {
		s += strconv.Itoa(i) + " : " + arg + "\n"
	}
	return strings.Trim(s, "\n")
}
