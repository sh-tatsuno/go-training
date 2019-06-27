package main

import (
	"fmt"
	"io"
	"os"

	"github.com/sh-tatsuno/go-training/ch04/4.11/gissue/cmd"
	"golang.org/x/xerrors"
)

func usage() {
	io.WriteString(os.Stderr, UsageText)
}

const UsageText = `this is a handling github isuue tool.
command is below
- list
- create 
- edit
- delete
- show
Further information could be shown in document.
`

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

func main() {
	cmder := cmd.NewCmder()
	os.Exit(run(os.Args[1:], cmder))
}

func run(args []string, cmder cmd.CmderInterface) int {

	// output usage
	if len(args) < 2 {
		usage()
		return ExitCodeError
	}

	// subcommand
	var err error
	switch args[0] {
	case "list":
		err = cmder.List(args[1:])
	case "create":
		err = cmder.Create(args[1:])
	case "show":
		err = cmder.Show(args[1:])
	case "edit":
		err = cmder.Edit(args[1:])
	case "close":
		err = cmder.Close(args[1:])
	default:
		err = xerrors.New("undefined command.")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "err: %v", err)
		return ExitCodeError
	}

	return ExitCodeOK
}
