package main

import "fmt"

type Mac struct {
	Printer Printer
}

func (m *Mac) SetPrinter(printer Printer) {
	m.Printer = printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
	m.Printer.PrintFile()
}
