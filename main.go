package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte(`{"status": "ok"}`))
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")), mux)
	if err != nil {
		log.Fatal(err)
	}
}
