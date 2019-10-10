package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sh-tatsuno/go-training/ch08/ex06/links"
)

type urlStruct struct {
	url   string
	depth int
}

// go run main.go --url https://kakaku.com/item/K0001088733/ --depth 3
func main() {
	var (
		baseURL  = flag.String("url", "", "base url")
		maxDepth = flag.Int("depth", 3, "max depth")
	)
	flag.Parse()
	worklist := make(chan []urlStruct)
	unseenLinks := make(chan urlStruct)

	go func() {
		worklist <- createURLStruct([]string{*baseURL}, 0)
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				if link.depth < *maxDepth {
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if link.depth >= *maxDepth {
				return
			}
			if !seen[link.url] {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(u urlStruct) []urlStruct {
	//fmt.Println(u.url)
	fmt.Println(u)
	tokens <- struct{}{}
	list, err := links.Extract(u.url)
	<-tokens

	if err != nil {
		log.Print(err)
	}
	return createURLStruct(list, u.depth)
}

func createURLStruct(list []string, prevDepth int) []urlStruct {
	var u []urlStruct
	for _, l := range list {
		u = append(u, urlStruct{l, prevDepth + 1})
	}
	return u
}
