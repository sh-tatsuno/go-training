package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/sh-tatsuno/go-training/ch04/4.11/gissue/editor"

	"github.com/sh-tatsuno/go-training/ch04/4.11/gissue/entity"
	"golang.org/x/xerrors"
)

type CmderInterface interface {
	List(args []string) error
	Create(args []string) error
	Edit(args []string) error
	Close(args []string) error
	Show(args []string) error
}

type Cmder struct {
	ed       editor.Editor
	user     string
	password string
}

const issuesURL = "https://api.github.com/repos/"

func NewCmder() Cmder {
	ed := editor.NewEditor()
	c := Cmder{
		ed: ed,
	}
	return c
}

func (c *Cmder) init(u string, p string) {
	c.user = u
	c.password = p
}

func (c Cmder) setReq(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")
	req.SetBasicAuth(c.user, c.password)
}

func (c Cmder) read(args []string) (*entity.Issue, error) {
	f := flag.NewFlagSet(os.Args[1], flag.ContinueOnError)
	u := f.String("u", "", "github user name")
	p := f.String("p", "", "github password")
	i := f.Int("i", 0, "issue number")
	f.Parse(args[1:])
	repo := args[0]
	c.init(*u, *p)
	if *i <= 0 {
		return nil, xerrors.New("issue number should be positive number.")
	}

	url := issuesURL + path.Join(repo, "issues", strconv.Itoa(*i))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, xerrors.Errorf("http.NewRequest error: %v", err)
	}
	c.setReq(req)

	client := &http.Client{Timeout: time.Duration(15 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return nil, xerrors.Errorf("client.Do error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s := fmt.Sprintf("response is incorrect. expected: %v, actual: %v", http.StatusOK, resp.StatusCode)
		return nil, xerrors.New(s)
	}

	result := entity.Issue{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
