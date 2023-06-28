package main

import "fmt"

type HP struct {
}

func (p *HP) PrintFile() {
	fmt.Println("Print by HP")
}
