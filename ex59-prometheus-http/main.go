package main

import (
	"fmt"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const port = ":3015"

func main() {
	var helloCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_request_count",
		Help: "hello 请求的次数",
	})
	prometheus.MustRegister(helloCounter)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		helloCounter.Inc()
		w.Write([]byte("hello"))
	})
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("listen at", port)
	http.ListenAndServe(port, nil)
}
