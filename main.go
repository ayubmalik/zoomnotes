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

func run(args []string, w io.Writer) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	src := path.Join(home, "Documents", "Zoom")
	dst := "output.md"

	fmt.Fprintf(w, "Saving file %s\n", dst)
	return SaveZoomNotes(src, dst)
}
