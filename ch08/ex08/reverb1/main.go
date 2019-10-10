package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn, done chan struct{}) {
	input := bufio.NewScanner(c)

	// タイマーを張る
	called := make(chan struct{})
	go func(c net.Conn, called <-chan struct{}) {
		for {
			timer := time.After(10 * time.Second)
			select {
			case _, ok := <-called:
				if !ok {
					return
				}
			case <-timer:
				c.Close()
				fmt.Println("connection closed")
				done <- struct{}{}
			}
		}
	}(c, called)

	for input.Scan() {
		called <- struct{}{}

		echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

// clock2と同じもの
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn, done)

		select {
		case <-done:
			return
		}
	}

}
