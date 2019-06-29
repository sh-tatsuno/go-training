package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(run(os.Args))
}

func run(args []string) string {
	return strings.Join(args, " ")
}
