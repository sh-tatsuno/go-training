package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	var p int
	flag.IntVar(&p, "port", 8000, "port number")
	flag.Parse()
	address := "localhost:" + strconv.Itoa(p)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
