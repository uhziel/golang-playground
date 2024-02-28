package main

import "fmt"

func main() {
	var m map[string]string
	v, ok := m["foo"]
	fmt.Printf("%#v %v\n", v, ok)
	m["foo"] = "bar"
}
