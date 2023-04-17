package main

import "fmt"

func addr(s string) {
    fmt.Println("addr(s). &s: ", &s)
}

func main() {
    s := "foo"
    fmt.Println("main(). &s: ", &s)
    addr(s)
}
