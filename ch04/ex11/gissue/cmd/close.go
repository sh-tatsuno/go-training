package cmd

import (
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

func (c Cmder) Close(args []string) error {

	 issue, err := c.read(args)
	if err != nil {
		return xerrors.Errorf("c.read, err: %v", err)
	}

	f := flag.NewFlagSet(os.Args[1], flag.ContinueOnError)
	u := f.String("u", "", "github user name")
	p := f.String("p", "", "github password")
	i := f.Int("i", 0, "issue number")
	f.Parse(args[1:])
	repo := args[0]
	c.init(*u, *p)

	issue.State = "closed"
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
		return fmt.Errorf("close issue failed. code: %s, message: %s", resp.Status, resp.Body)
	}

	fmt.Printf("%v", *issue)
	return nil
}
