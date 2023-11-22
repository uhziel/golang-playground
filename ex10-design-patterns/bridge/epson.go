package main

import "fmt"

type Epson struct {
}

func (p *Epson) PrintFile() {
	fmt.Println("Print by Epson")
}
