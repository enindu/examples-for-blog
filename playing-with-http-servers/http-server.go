package main

import (
	"fmt"
	"net/http"
)

func HTTPServer() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// Point 1
	})

	err := http.ListenAndServe("127.0.0.1:5000", nil)
	if err != nil {
		panic(err)
	}
}

func HTTPServerFinal() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		// Point 1
		if r.Method != http.MethodGet {
			http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Point 2
		fmt.Fprintf(rw, "Hello world!")
	})

	err := http.ListenAndServe("127.0.0.1:5000", nil)
	if err != nil {
		panic(err)
	}
}
