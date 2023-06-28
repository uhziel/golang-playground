package main

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) ChangeBuilder(builder Builder) {
	d.builder = builder
}

func (d *Director) BuildHouse() House {
	d.builder.SetWindowType()
	d.builder.SetDoorType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
