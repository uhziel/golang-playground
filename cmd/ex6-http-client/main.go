package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	url = flag.String("url", "http://ip.3322.org", "view the page")
)

func main() {
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
