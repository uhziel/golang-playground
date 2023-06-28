package main

type ShoeInterface interface {
	SetLogo(v string)
	GetLogo() string
	SetSize(v int)
	GetSize() int
}

type shoe struct {
	logo string
	size int
}

func (s *shoe) SetLogo(v string) {
	s.logo = v
}

func (s *shoe) GetLogo() string {
	return s.logo
}

func (s *shoe) SetSize(v int) {
	s.size = v
}

func (s *shoe) GetSize() int {
	return s.size
}
