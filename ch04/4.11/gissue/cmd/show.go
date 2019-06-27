package cmd

import (
	"fmt"

	"golang.org/x/xerrors"
)

func (c Cmder) Show(args []string) error {
	issue, err := c.read(args)
	if err != nil {
		return xerrors.Errorf("c.read, err: %v", err)
	}

	fmt.Printf("%v", *issue)
	return nil
}
