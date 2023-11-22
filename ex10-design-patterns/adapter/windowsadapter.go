package main

import "fmt"

type WindowsAdapter struct {
	Windows *Windows
}

func (a *WindowsAdapter) Insert2LightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	a.Windows.Insert2USBPort()
}
