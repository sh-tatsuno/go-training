package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"golang.org/x/xerrors"
)

func (c Cmder) Edit(args []string) error {

	issue, err := c.read(args)
	if err != nil {
		return xerrors.Errorf("c.read, err: %v", err)
	}

	f := flag.NewFlagSet(os.Args[1], flag.ContinueOnError)
	u := f.String("u", "", "github user name")
	p := f.String("p", "", "github password")
	e := f.String("e", "vim", "github password")
	i := f.Int("i", 0, "issue number")
	f.Parse(args[1:])
	repo := args[0]
	c.init(*u, *p)

	// title -> 初期値をどう入れるか分からなかった
	fmt.Println("title: ")
	stdin := bufio.NewScanner(os.Stdin)

	stdin.Scan()
	title := stdin.Text()
	if err := stdin.Err(); err != nil {
		return err
	}

	// open & use editor
	content, err := c.ed.Use(*e, issue.Body)
	if err != nil {
		return err
	}

	issue.Title = title
	issue.Body = content

	input, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	url := issuesURL + path.Join(repo, "issues", strconv.Itoa(*i))
	req, err := http.NewRequest(
		"PATCH",
		url,
		bytes.NewBuffer(input),
	)
	if err != nil {
		return err
	}
	c.setReq(req)

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return xerrors.Errorf("edit issue failed: %s\n", resp.Status)
	}

	fmt.Printf("%v", *issue)

	return nil
}
