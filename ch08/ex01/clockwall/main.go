// WIP

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	args := os.Args[1:]

	var conns []io.Reader
	for _, address := range args {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal(err)
		}
		conns = append(conns, conn)
		defer conn.Close()
	}

	go mustCopy(os.Stdout, conns)
}

func mustCopy(dst io.Writer, src []io.Reader) {
	for _, s := range src {
		if _, err := io.Copy(dst, s); err != nil {
			log.Fatal(err)
		}
	}
}
