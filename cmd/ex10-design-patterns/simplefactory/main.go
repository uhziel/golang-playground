package main

import "fmt"

func main() {
	ak47, _ := New("ak47")
	printDetail(ak47)
	musket, _ := New("musket")
	printDetail(musket)
}

func printDetail(gun Interface) {
	fmt.Println("Name:", gun.GetName())
	fmt.Println("Power:", gun.GetPower())
}
