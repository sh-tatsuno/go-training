package editor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"golang.org/x/xerrors"
)

const tmpPath = "tmp"

type EditorInterface interface {
	Use(name string, preset string) (string, error)
}

type Editor struct {
	name string
}

func NewEditor() Editor {
	ed := Editor{}
	return ed
}

func (e Editor) Use(name string, preset string) (string, error) {

	// Make temp editing file
	fPath := getFilePath(tmpPath)
	err := makeFile(fPath, preset)
	if err != nil {
		return "", xerrors.Errorf("failed make edit file. %s\n", err.Error())
	}
	defer deleteFile(fPath)

	// Open text editor
	err = openEditor(name, fPath)
	if err != nil {
		return "", xerrors.Errorf("failed open text editor. %s\n", err.Error())
	}

	// Read edit file
	content, err := ioutil.ReadFile(fPath)
	if err != nil {
		return "", xerrors.Errorf("failed read content. %s\n", err.Error())
	}

	return string(content), nil
}

func getFilePath(about string) string {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		home = os.Getenv("APPDATA")
	}
	fname := filepath.Join(home, tmpPath, fmt.Sprintf("%s_EDITMSG", about))
	return fname
}

func makeFile(fPath string, preset string) (err error) {
	if !isFileExist(fPath) {
		err = ioutil.WriteFile(fPath, []byte(preset), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func isFileExist(fPath string) bool {
	_, err := os.Stat(fPath)
	return err == nil || !os.IsNotExist(err)
}

func deleteFile(fPath string) error {
	return os.Remove(fPath)
}

func openEditor(name string, args ...string) error {
	c := exec.Command(name, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
