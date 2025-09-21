package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fname := range os.Args[1:] { // loop all file names in cmdLine
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if _, err := io.Copy(os.Stdout, file); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		file.Close()
	}
}
