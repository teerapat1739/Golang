package main

import (
	"fmt"
	"net/http"
)


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}

		w.Header().Set("Content-type","application/json")
		fmt.Fprintf(w, `{"message":"welcome to the homepage!"}`)
	})

	http.ListenAndServe(":8080", mux)
}
