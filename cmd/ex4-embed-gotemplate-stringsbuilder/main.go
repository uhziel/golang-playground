package main

import (
  "fmt"
  "strings"
  "text/template"
  _ "embed"
)

type Person struct {
  Name string
  Foo []string
}

//go:embed hello.txt
var hello string

func main() {
  tmpl, err := template.New("test").Parse(hello)
  if err != nil {
    panic(err)
  }
  lilei := Person{
    Name: "lilei",
    Foo: []string{
      "bar1",
      "bar2",
    },
  }
  b := new(strings.Builder)
  err = tmpl.Execute(b, lilei)
  if err != nil {
    panic(err)
  }
  fmt.Print(b.String())
}
