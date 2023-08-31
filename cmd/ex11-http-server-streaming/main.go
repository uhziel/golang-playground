package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.Int("port", 50001, "listen port")
)

func main() {
	flag.Parse()

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", "hello")
	})

	http.HandleFunc("/stream", func(w http.ResponseWriter, req *http.Request) {
		flusher := w.(http.Flusher)
		for i := 0; i < 5; i++ {
			fmt.Fprintf(w, "%s %d\n", "hello", i)
			flusher.Flush()
			time.Sleep(time.Second)
		}
	})

	log.Println("Starting server at port: ", *port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
