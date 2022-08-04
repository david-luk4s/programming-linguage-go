package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func dup2() {
	counts := make(map[string]int)

	//give names of files
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "bufio")
	} else {
		for _, n := range files {
			f, err := os.Open(n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}

			status := countLines(f, counts, n)

			if status {
				fmt.Printf("existe duplicados em %s\n", n)
			}

			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}

}

func countLines(f *os.File, counts map[string]int, filename string) bool {
	status := false
	input := bufio.NewScanner(f)

	for input.Scan() {
		if input.Text() == ":q" {
			break
		}
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			status = true
		}
	}

	return status
}
