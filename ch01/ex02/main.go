package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args {
		s := strconv.Itoa(i) + " : " + arg
		fmt.Println(s)
	}
}
