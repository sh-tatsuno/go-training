package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/xerrors"
)

type GoJSON struct {
	Dir        string
	ImportPath string
	Name       string
	Target     string
	Stale      bool
	Root       string
	GoFiles    []string
	Imports    []string
	Deps       []string
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return xerrors.New("have to use more args.")
	}
	newArgs := []string{"list", "-json"}
	for _, arg := range args {
		newArgs = append(newArgs, arg+"...")
	}

	out, err := exec.Command("go", newArgs...).Output()
	if err != nil {
		fmt.Printf("Command Exec Error. err: %v\n", err)
	}

	scanner := bufio.NewScanner(bytes.NewBuffer(out))
	jsonTxt := ""
	var deps []string
	var goJSON GoJSON

	for scanner.Scan() {
		t := scanner.Text()
		jsonTxt += t
		if string(t[0]) == "}" {
			err = json.Unmarshal([]byte(jsonTxt), &goJSON)
			if err != nil {
				fmt.Printf("Cannot unmarshal. err: %v\n", err)
			}
			deps = append(deps, goJSON.Deps...)
			jsonTxt = ""
		}

	}

	// 重複除外
	m := make(map[string]struct{})
	newList := make([]string, 0)
	for _, element := range deps {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; !ok {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}
	fmt.Printf("%v", newList)
	return nil
}
