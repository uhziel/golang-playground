package main

import (
	"fmt"
	"net/http"
)

const addr = ":5671"

func main() {
	http.HandleFunc("POST /download", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	fmt.Println("listen at", addr)
	http.ListenAndServe(addr, nil)
}
