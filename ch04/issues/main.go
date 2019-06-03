package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sh-tatsuno/go-training/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s %s %s\n", item.Number, item.User.Login, item.Title, item.HTMLURL, item.CreatedAt)
	}
}
