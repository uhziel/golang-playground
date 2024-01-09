package main

import (
	"flag"
	"fmt"
)

// via https://pkg.go.dev/flag

func main() {
	var n = flag.Int("n", 1234, "help message for flag n")
	var kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	var gopherType string
	const (
		defaultGopher = "pocket"
		usage         = "the variety of gopher"
	)
	flag.StringVar(&gopherType, "gopher_type", defaultGopher, usage)
	flag.StringVar(&gopherType, "g", defaultGopher, usage+" (shorthand)")

	flag.Parse()

	fmt.Println("n:", *n)
	fmt.Println("kubecconfig:", *kubeconfig)
	fmt.Println("gopherType:", gopherType)
}
