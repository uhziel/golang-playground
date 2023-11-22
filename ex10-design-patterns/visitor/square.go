package main

type Square struct {
}

func (s *Square) GetType() string {
	return "square"
}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitForSquare(s)
}
