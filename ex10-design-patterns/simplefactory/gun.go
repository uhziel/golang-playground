package main

type Interface interface {
	SetName(v string)
	GetName() string
	SetPower(v int)
	GetPower() int
}

type gun struct {
	name  string
	power int
}

func (g *gun) SetName(v string) {
	g.name = v
}

func (g *gun) GetName() string {
	return g.name
}

func (g *gun) SetPower(v int) {
	g.power = v
}

func (g *gun) GetPower() int {
	return g.power
}
