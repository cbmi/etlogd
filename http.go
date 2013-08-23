package main

import (
	"log"
	"net/http"
)

func ServeHttp(cfg *Config) {
	defer cfg.Mongo.session.Close()

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

		handleMessage(cfg, r.Body)
	})

	log.Printf("Listening on http://%s...\n", cfg.Server.Host)
	handleFatal(http.ListenAndServe(cfg.Server.Host, nil))
}
