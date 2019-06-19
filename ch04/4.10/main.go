package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/sh-tatsuno/go-training/ch04/4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	resultStrings := map[string][]string{"under1Month": []string{}, "under1Year": []string{}, "over1Year": []string{}}

	// sort result based on created time of issues
	sort.SliceStable(result.Items,
		func(i, j int) bool {
			return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
		})

	for _, item := range result.Items {
		now := time.Now().UTC() // avoid applying JST
		diffDays := float64(now.Sub(item.CreatedAt).Hours()) / 24
		diffMonth := diffDays / 12
		diffYear := diffMonth / 12

		issue := fmt.Sprintf("%s  #%-5d  %9.9s %.55s %s \n", item.CreatedAt.Format("2006-01-02"), item.Number, item.User.Login, item.Title, item.HTMLURL)
		if diffYear >= 1 {
			resultStrings["over1Year"] = append(resultStrings["over1Year"], issue)
		} else if diffMonth >= 1 {
			resultStrings["under1Year"] = append(resultStrings["under1Year"], issue)
		} else {
			resultStrings["under1Month"] = append(resultStrings["under1Month"], issue)
		}

	}

	fmt.Printf("--------- Issues created under 1 month : total %d issues ---------\n", len(resultStrings["under1Month"]))
	for _, i := range resultStrings["under1Month"] {
		fmt.Print(i)
	}
	fmt.Println()

	fmt.Printf("--------- Issues created under 1 year : total %d issues ---------\n", len(resultStrings["under1Year"]))
	for _, i := range resultStrings["under1Year"] {
		fmt.Print(i)
	}
	fmt.Println()

	fmt.Printf("--------- Issues created over 1 year : total %d issues ---------\n", len(resultStrings["over1Year"]))
	for _, i := range resultStrings["over1Year"] {
		fmt.Print(i)
	}
	fmt.Println()
}
