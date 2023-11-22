package main

type Rectangle struct {
}

func (c *Rectangle) GetType() string {
	return "rectangle"
}

func (c *Rectangle) Accept(visitor Visitor) {
	visitor.VisitForRectangle(c)
}
