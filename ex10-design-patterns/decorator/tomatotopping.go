package main

type TomatoTopping struct {
	Pizza Pizza
}

func (t *TomatoTopping) GetPrice() int {
	return t.Pizza.GetPrice() + 2
}
