package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/sh-tatsuno/go-training/ch04/4.11/gissue/entity"
	"golang.org/x/xerrors"
)

func (c Cmder) List(args []string) error {

	f := flag.NewFlagSet(os.Args[1], flag.ContinueOnError)
	u := f.String("u", "", "github user name")
	p := f.String("p", "", "github password")
	f.Parse(args[1:])
	repo := args[0]
	c.init(*u, *p)

	url := issuesURL + path.Join(repo, "issues")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return xerrors.Errorf("http.NewRequest error: %v", err)
	}
	c.setReq(req)

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return xerrors.Errorf("client.Do error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s := fmt.Sprintf("response is incorrect. expected: %v, actual: %v", http.StatusOK, resp.StatusCode)
		return xerrors.New(s)
	}

	result := []entity.Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	fmt.Printf("%v\n", result)

	return nil
}
