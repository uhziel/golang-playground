package main

type Nike struct {
}

type NikeShoe struct {
	shoe
}

func (n *Nike) CreateShoe() ShoeInterface {
	return &NikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 10,
		},
	}
}

type NikeShirt struct {
	shirt
}

func (n *Nike) CreateShirt() ShirtInterface {
	return &NikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 21,
		},
	}
}
