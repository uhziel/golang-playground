package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.String("port", "6000", "监听的端口")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	addr := fmt.Sprintf(":%s", *port)

	fmt.Printf("listen at %s\n", addr)
	panic(http.ListenAndServe(addr, nil))
}
