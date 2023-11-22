package main

import "fmt"

func main() {
	nike, _ := NewFactory("nike")
	adidas, _ := NewFactory("adidas")

	nikeShoe := nike.CreateShoe()
	nikeShirt := nike.CreateShirt()

	adidasShoe := adidas.CreateShoe()
	adidasShirt := adidas.CreateShirt()

	printShoeDetail(nikeShoe)
	printShirtDetail(nikeShirt)

	printShoeDetail(adidasShoe)
	printShirtDetail(adidasShirt)
}

func printShoeDetail(shoe ShoeInterface) {
	fmt.Println("shoe logo:", shoe.GetLogo())
	fmt.Println("shoe size:", shoe.GetSize())
}

func printShirtDetail(shirt ShirtInterface) {
	fmt.Println("shirt logo:", shirt.GetLogo())
	fmt.Println("shirt size:", shirt.GetSize())
}
