// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler3)
	http.HandleFunc("/count", counter3)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler3 echoes the Path component of the requested URL.
func handler3(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	_, _ = fmt.Fprintf(w, "Host = %q\n", r.Host)
	_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// counter echoes the number of calls so far.
func counter3(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	_, _ = fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
