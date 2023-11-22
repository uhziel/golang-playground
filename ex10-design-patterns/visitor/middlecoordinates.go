package main

import "fmt"

type MiddleCoordinates struct {
	X, Y int
}

func (m *MiddleCoordinates) VisitForSquare(square *Square) {
	fmt.Println("MiddleCoordinates for Square")
}

func (m *MiddleCoordinates) VisitForCircle(circle *Circle) {
	fmt.Println("MiddleCoordinates for Circle")
}

func (m *MiddleCoordinates) VisitForRectangle(rectangle *Rectangle) {
	fmt.Println("MiddleCoordinates for Rectangle")
}
