package main

type Computer interface {
	SetPrinter(printer Printer)
	Print()
}
