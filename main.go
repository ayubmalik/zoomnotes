package main

import (
	"fmt"
	"io"
	"os"
	"path"
)

const (
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	src := path.Join(home, "Documents", "Zoom")
	dst := "notes/README.md"

	return SaveZoomNotes(src, dst)
}
