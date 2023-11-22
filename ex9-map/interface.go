package ex9map

import "fmt"

type Map interface {
	Search(target int) bool
	Add(num int) bool
	Remove(num int) bool
	fmt.Stringer
}
