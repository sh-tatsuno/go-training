package main

import (
	"fmt"

	"github.com/sh-tatsuno/go-training/ch02/2.1/tmpconv"
)

func main() {
	c := tmpconv.AbsoluteZeroC
	fmt.Println(c.String())

	f := tmpconv.CToF(c)
	fmt.Println(f.String())

	k := tmpconv.CToK(c)
	fmt.Println(k.String())
}

// -273.15^C
// -459.66999999999996^F
// 0^K
