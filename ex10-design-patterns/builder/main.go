package main

import "fmt"

func main() {
	iglooBuilder := &IglooBuilder{}
	normalBuilder := &NormalBuilder{}

	director := NewDirector(iglooBuilder)
	iglooHouse := director.BuildHouse()
	printDetail(iglooHouse)

	director.ChangeBuilder(normalBuilder)
	normalHouse := director.BuildHouse()
	printDetail(normalHouse)
}

func printDetail(house House) {
	fmt.Println("Window Type:", house.WindowType)
	fmt.Println("Door Type:", house.DoorType)
	fmt.Println("Floor:", house.Floor)
}
