package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sh-tatsuno/go-training/ch02/2.2/tmpconv"
)

func main() {
	args := os.Args[1:]
	var s string
	if len(args) < 1 {
		var sc = bufio.NewScanner(os.Stdin)
		if sc.Scan() {
			s = sc.Text()
		}
		args = strings.Split(s, " ")
	}

	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tmpconv.Fahrenheit(t)
		c := tmpconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tmpconv.FToC(f), c, tmpconv.CToF(c))
	}

}
