package main

import "fmt"

type Windows struct {
	Printer Printer
}

func (w *Windows) SetPrinter(printer Printer) {
	w.Printer = printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for windows")
	w.Printer.PrintFile()
}
