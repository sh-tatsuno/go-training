package main

import (
	"fmt"

	"github.com/sh-tatsuno/go-training/ch02/2.5/popcount"
)

func main() {
	val := uint64(10)
	fmt.Println(popcount.PopCount(val))
	fmt.Println(popcount.PopCount4(val))
}
