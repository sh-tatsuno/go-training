package main

import (
	"fmt"
	"os"
)

func main() {
	s1 := string(os.Args[1])
	s2 := string(os.Args[2])
	m := map[string]int{}

	for _, c := range s1 {
		_, ok := m[string(c)]
		if !ok {
			m[string(c)] = 1
		} else {
			m[string(c)]++
		}
	}

	for _, c := range s2 {
		_, ok := m[string(c)]
		if !ok {
			m[string(c)] = -1
		} else {
			m[string(c)]--
		}
	}

	for k := range m {
		if m[k] != 0 {
			fmt.Println("they are not anagram")
			return
		}
	}
	fmt.Println("they are anagram")

	return
}
