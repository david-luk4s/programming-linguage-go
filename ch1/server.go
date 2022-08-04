package ch1

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mux sync.Mutex

func server() {
	http.HandleFunc("/home", handler)
	http.HandleFunc("/count", Count)
	log.Panic(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	count++
	mux.Unlock()
	fmt.Fprintf(w, "router: %s\n", r.URL.Path)
	fmt.Fprintf(w, "agent: %s\n", r.UserAgent())
	fmt.Fprintf(w, "addr: %s\n", r.RemoteAddr)
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s:%s\n", k, v)
	}
}

func Count(w http.ResponseWriter, r *http.Request) {
	mux.Lock()
	fmt.Fprintf(w, "count: %d", count)
	mux.Unlock()
}
