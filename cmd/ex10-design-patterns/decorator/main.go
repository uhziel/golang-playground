package main

import "fmt"

func main() {
	var pizza Pizza = &VeggieMania{}
	pizza = &TomatoTopping{
		Pizza: pizza,
	}
	pizza = &CheeseTopping{
		Pizza: pizza,
	}
	fmt.Println("pizza+tomato+cheese price:", pizza.GetPrice())
}
