package main

type Adidas struct {
}

type AdidasShoe struct {
	shoe
}

func (a *Adidas) CreateShoe() ShoeInterface {
	return &AdidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 10,
		},
	}
}

type AdidasShirt struct {
	shirt
}

func (a *Adidas) CreateShirt() ShirtInterface {
	return &AdidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 21,
		},
	}
}
