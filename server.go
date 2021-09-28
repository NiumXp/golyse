package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/running", runningHandler)
	http.HandleFunc("/stop", stopHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func runningHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	go func() { os.Exit(0) }()
}
