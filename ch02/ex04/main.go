package main

import (
	"fmt"

	"github.com/sh-tatsuno/go-training/ch02/2.4/popcount"
)

func main() {
	val := uint64(15)
	fmt.Println(popcount.PopCount(val))
	fmt.Println(popcount.PopCount3(val))
}
