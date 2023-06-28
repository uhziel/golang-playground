package main

import "fmt"

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}

func NewBuilder(builder string) (Builder, error) {
	switch builder {
	case "igloo":
		return &IglooBuilder{}, nil
	case "normal":
		return &NormalBuilder{}, nil
	default:
		return nil, fmt.Errorf("wrong builder type: %s", builder)
	}
}
