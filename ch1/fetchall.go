package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func fetchall() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go ffetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Printf(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func ffetch(url string, ch chan<- string) {
	start := time.Now()
	file, err := os.OpenFile(url, os.O_CREATE|os.O_RDWR, 0644)
	defer file.Close()

	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}

	//nbytes, err := io.Copy(io.Discard, resp.Body)
	bytes, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s %v", url, err)
		return
	}

	file.Write(bytes)

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, len(bytes), url)
}
