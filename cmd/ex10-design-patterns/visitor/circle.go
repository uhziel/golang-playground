package main

type Circle struct {
}

func (c *Circle) GetType() string {
	return "circle"
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitForCircle(c)
}
