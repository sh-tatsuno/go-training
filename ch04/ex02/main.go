package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {

	s := flag.String("s", "sha256", "sha256, sha384, sha512")
	flag.Parse()

	if (*s != "sha256") && (*s != "sha384") && (*s != "sha512") {
		fmt.Fprintf(os.Stderr, "should use sha256, sha384, sha512. your input is: %s\n", *s)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		d := scanner.Bytes()
		switch *s {
		case "sha256":
			res := sha256.Sum256(d)
			fmt.Printf("%x\n", res)
		case "sha384":
			res := sha512.Sum384(d)
			fmt.Printf("%x\n", res)
		case "sha512":
			res := sha512.Sum512(d)
			fmt.Printf("%x\n", res)
		}
	}
}
