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

	log.Println(fmt.Sprintf("listening at :%d", *port))
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", *port), http.FileServer(http.Dir("."))))
}
