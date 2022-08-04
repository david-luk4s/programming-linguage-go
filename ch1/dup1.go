package ch1

import (
	"bufio"
	"fmt"
	"os"
)

func dup1() {
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	for input.Scan() {
		if input.Text() == ":q" {
			break
		}
		counts[input.Text()]++
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
