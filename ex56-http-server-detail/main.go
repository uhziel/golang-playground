package main

import (
	"log"
	"net/http"
)

const addr = "0.0.0.0:4567"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Println("listen at", addr)
	log.Fatalln(server.ListenAndServe())
}
