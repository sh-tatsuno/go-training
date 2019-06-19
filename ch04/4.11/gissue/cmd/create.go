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
	"time"

	"github.com/sh-tatsuno/go-training/ch04/4.11/gissue/entity"
)

func (c Cmder) Create(args []string) error {

	f := flag.NewFlagSet(os.Args[1], flag.ContinueOnError)
	u := f.String("u", "", "github user name")
	p := f.String("p", "", "github password")
	e := f.String("e", "vim", "github password")
	f.Parse(args[1:])
	repo := args[0]
	c.init(*u, *p)

	// title
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	title := stdin.Text()
	if err := stdin.Err(); err != nil {
		return err
	}

	// open & use editor
	content, err := c.ed.Use(*e, "")
	if err != nil {
		return err
	}

	issue := entity.Issue{
		Title: title,
		Body:  content,
	}

	input, err := json.Marshal(issue)
	if err != nil {
		return err
	}

	url := issuesURL + path.Join(repo, "issues")
	req, err := http.NewRequest(
		"POST",
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

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("create issue failed: %s", resp.Status)
	}

	// save
	return nil
}
