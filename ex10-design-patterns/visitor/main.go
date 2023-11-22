package main

func main() {
	square := &Square{}
	circle := &Circle{}
	rectangle := &Rectangle{}

	areaCalculator := &AreaCalculator{}
	middleCoordinates := &MiddleCoordinates{}

	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	rectangle.Accept(middleCoordinates)
}
