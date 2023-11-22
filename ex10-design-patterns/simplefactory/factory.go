package main

import "fmt"

func New(gun string) (Interface, error) {
	switch gun {
	case "ak47":
		return NewAK47(), nil
	case "musket":
		return NewMusket(), nil
	default:
		return nil, fmt.Errorf("wrong gun type passed")
	}
}
