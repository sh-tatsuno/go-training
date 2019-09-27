package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var p int
	flag.IntVar(&p, "port", 8000, "port number")
	flag.Parse()

	address := "localhost:" + strconv.Itoa(p)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 接続切断
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	tz := os.Getenv("TZ")
	loc, err := time.LoadLocation(tz)
	if err != nil {
		log.Print(err)
		loc, _ = time.LoadLocation("UTC")
	}
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(loc).Format("15:04:05\n"))
		if err != nil {
			log.Print(err)
			return // 接続切断
		}
		time.Sleep(1 * time.Second)
	}
}
