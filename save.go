package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// SaveZoomNotes all Zoom meeting notes from src directory to dst file.
func SaveZoomNotes(src, dst string) error {
	noteFiles, err := findMeetingNotes(src)
	if err != nil {
		return errors.Wrap(err, "Save")
	}

	w, err := os.Create(dst)
	must(err, "cannot create file: "+dst)
	defer w.Close()

	for _, f := range noteFiles {
		md, err := formatMarkdown(f)
		if err != nil {
			return err
		}
		w.WriteString(md)

	}
	return nil
}

func findMeetingNotes(dir string) (paths []string, err error) {
	log.Println("zoom dir", dir)
	paths = make([]string, 0)
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && strings.Contains(info.Name(), "meeting_saved_chat") {
			paths = append(paths, path)
		}
		return nil
	})

	return paths, err
}

func formatMarkdown(file string) (string, error) {
	info, _ := os.Stat(file)
	var s strings.Builder

	s.WriteString(fmt.Sprintf("# %s\n\n", info.ModTime().Format("02/01/2006")))

	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// remove prefix e.g. in line "18:01:23    From Roger Mellie : Hello world"
	// will only keep "Hello world"
	regex := regexp.MustCompile("^\\d\\d:\\d\\d:\\d\\d.*From.*:")
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n') // 0x0A separator = newline
		if err == io.EOF {
			break
		}

		if err != nil {
			return "", nil
		}

		// strip fixed width date
		line = regex.ReplaceAllString(line, "")
		line = strings.TrimSpace(line)
		s.WriteString(line + "\n\n")
	}
	s.WriteString("\n")
	return s.String(), nil
}

func must(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, ":")
	}
}
