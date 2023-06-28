package main

import "fmt"

type FactoryInterface interface {
	CreateShoe() ShoeInterface
	CreateShirt() ShirtInterface
}

func NewFactory(factory string) (FactoryInterface, error) {
	switch factory {
	case "nike":
		return &Nike{}, nil
	case "adidas":
		return &Adidas{}, nil
	default:
		return nil, fmt.Errorf("wrong factory type: %s", factory)
	}
}
