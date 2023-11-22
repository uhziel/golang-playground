package main

type ShirtInterface interface {
	SetLogo(v string)
	GetLogo() string
	SetSize(v int)
	GetSize() int
}

type shirt struct {
	logo string
	size int
}

func (s *shirt) SetLogo(v string) {
	s.logo = v
}

func (s *shirt) GetLogo() string {
	return s.logo
}

func (s *shirt) SetSize(v int) {
	s.size = v
}

func (s *shirt) GetSize() int {
	return s.size
}
