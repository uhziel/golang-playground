package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.Int("port", 50001, "listen port")
)

func main() {
	flag.Parse()

	http.HandleFunc("/files", func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "main.go")
	})

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
