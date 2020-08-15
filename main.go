package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        handles, ok := r.Header["X-Sandstorm-Preferred-Handle"]
	var handle string
	if !ok {
	    handle = "friend"
	} else {
	    handle = handles[0]
	}
        fmt.Fprintf(w, "Hello %s, you've requested: %s\n", handle, r.URL.Path)
    })

    http.ListenAndServe(":8080", nil)
}
