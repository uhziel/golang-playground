package main

type CheeseTopping struct {
	Pizza Pizza
}

func (c *CheeseTopping) GetPrice() int {
	return c.Pizza.GetPrice() + 3
}
