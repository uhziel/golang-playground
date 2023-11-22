package main

func main() {
	mac := &Mac{}
	windows := &Windows{}

	espon := &Epson{}
	hp := &HP{}

	mac.SetPrinter(espon)
	mac.Print()
	mac.SetPrinter(hp)
	mac.Print()

	windows.SetPrinter(espon)
	windows.Print()
	windows.SetPrinter(hp)
	windows.Print()
}
