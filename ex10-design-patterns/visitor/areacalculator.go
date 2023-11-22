package main

import "fmt"

type AreaCalculator struct {
	Area int
}

func (a *AreaCalculator) VisitForSquare(square *Square) {
	fmt.Println("AreaCalculator for Square")
}

func (a *AreaCalculator) VisitForCircle(circle *Circle) {
	fmt.Println("AreaCalculator for Circle")
}

func (a *AreaCalculator) VisitForRectangle(rectangle *Rectangle) {
	fmt.Println("AreaCalculator for Rectangle")
}
