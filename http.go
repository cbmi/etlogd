package main

import (
	"fmt"
	"net/http"
)

func ServeHttp(host string, port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Currently only write only server..
		if r.Method != "POST" {
			w.WriteHeader(405)
			return
		}

		// Ensure the header exists and the media type is supported
		elem, ok := r.Header["Content-Type"]

		if !ok || elem[0] != "application/json" {
			w.WriteHeader(415)
			return
		}

		// Ignore empty payloads
		if r.ContentLength == 0 {
			w.WriteHeader(422)
			return
		}

		handleMessage(r.Body)
	})

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Listening on http://%s...\n", addr)
	handleError(http.ListenAndServe(addr, nil))
}
