package main

type Shape interface {
	GetType() string
	Accept(visitor Visitor)
}
