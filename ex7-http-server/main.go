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

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", "hello")
	})

	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("listen at: %s", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
