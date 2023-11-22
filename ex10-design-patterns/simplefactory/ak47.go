package main

type AK47 struct {
	gun
}

func NewAK47() Interface {
	return &AK47{
		gun: gun{
			name:  "ak47",
			power: 3,
		},
	}
}
