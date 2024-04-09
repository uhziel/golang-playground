package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	// sort.Ints()
	a1 := []int{1, 4, 2, 5}
	sort.Ints(a1)
	fmt.Println("a1:", a1)
	// sort.Floats64()
	a2 := []float64{1.0, 0.8, 1.1, 2, -1.1}
	sort.Float64s(a2)
	fmt.Println("a2:", a2)
	// sort.Strings()
	a3 := []string{"abc", "abb", "ABC"}
	sort.Strings(a3)
	fmt.Println("a3:", a3)

	// sort.Slice() / sort.SliceStable()
	a4 := []struct {
		Name string
		Age  int
	}{
		{
			Name: "lee",
			Age:  33,
		},
		{
			Name: "john",
			Age:  11,
		},
		{
			Name: "tom",
			Age:  30,
		},
		{
			Name: "lucy",
			Age:  30,
		},
	}
	sort.SliceStable(a4, func(i, j int) bool {
		return a4[i].Age < a4[j].Age
	})
	fmt.Println("a4:", a4)

	// sort.Sort()
	a5 := []Person{
		{
			Name: "lee",
			Age:  33,
		},
		{
			Name: "john",
			Age:  11,
		},
		{
			Name: "tom",
			Age:  30,
		},
		{
			Name: "lucy",
			Age:  30,
		},
	}
	sort.Sort(ByAge(a5))
	fmt.Println("a5:", a5)
}
