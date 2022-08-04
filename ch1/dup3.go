package ch1

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup3() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, v := range counts {
		fmt.Printf("%d\t%s\n", v, line)
	}

}
