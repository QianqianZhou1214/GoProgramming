package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, fname := range os.Args[1:] { // loop all file names in cmdLine
		var lc, wc, cc int
		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()
			wc += len(strings.Fields(s)) //Filds is a func takes string and return a slice of string
			cc += len(s)
			lc++
		}

		fmt.Printf("%7d %7d %7d %s\n", lc, wc, cc, fname)
		file.Close()
	}
}
