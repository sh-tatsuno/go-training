package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// create issue
// read issue
// update issue
// close issue

//

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string `json:"title"`
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    `json:"body"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

const (
	// ExitCodeOK : exit code
	ExitCodeOK = 0

	// ExitCodeError : error code
	ExitCodeError = 1
)

const IssuesURL = "https://api.github.com/repos/"

func usage() {
	io.WriteString(os.Stderr, usageText)
	flag.PrintDefaults()
}

const usageText = `this is image convert library by go.

In normal usage, you should set -d for directory and -i for input extension.
You also have to set output extension by -o.
You can also set maximum nuber you want to convert by set n.
current available extensions are jpg, jpeg, png, and gif.

Example:
	gophoto -d dir -i .png -o .jpeg -n 10

`

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	var user, password, repository, title, comment string

	// args
	flags := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flags.Usage = usage
	flags.StringVar(&repository, "r", "", "repository")
	flags.StringVar(&user, "u", "", "user")
	flags.StringVar(&password, "p", "", "password")
	flags.StringVar(&title, "t", "", "title")
	flags.StringVar(&comment, "c", "", "comment")
	flags.Parse(args)

	createIssue(user, repository, password, title, comment)

	return ExitCodeOK

}

func createIssue(user, repository, password, title, comment string) error {
	issue := Issue{Title: title, Body: comment}

	url := IssuesURL + repository + "/issues"
	fmt.Println(url)
	input, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer(input),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")
	req.SetBasicAuth(user, password)

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("%v\n", resp)

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create issue failed: %s", resp.Status)
	}

	return nil
}
