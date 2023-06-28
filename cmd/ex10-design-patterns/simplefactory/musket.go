package main

type Musket struct {
	gun
}

func NewMusket() Interface {
	return &Musket{
		gun: gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
