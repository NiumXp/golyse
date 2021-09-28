package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var done = make(chan struct{})

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/running", runningHandler).Methods("GET")
	r.HandleFunc("/stop", stopHandler).Methods("GET")

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	<-done

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
}

func runningHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func stopHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	close(done)
}
