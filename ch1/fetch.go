package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "> fetch:\t%v\n", err)
			os.Exit(1)
		}

		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "> fetch reading:\t%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\n> status code %d\n", resp.StatusCode)
		//fmt.Printf("%d", point)
	}
}
