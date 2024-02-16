package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Reflection goes from interface value to reflection object.
	x := 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v.String())
	fmt.Println("type:", v.Type())
	fmt.Println("kind is Float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	var i8 uint8 = 3
	fmt.Println("type:", reflect.TypeOf(i8))
	vi8 := reflect.ValueOf(x)
	fmt.Println("value:", vi8.String())

	// 2. Reflection goes from reflection object to interface value.
	fmt.Println("value:", v.Interface().(float64))
	fmt.Println("value:", v)

	// 3. To modify a reflection object, the value must be settable.
	// v.SetFloat(3.5) // 会 panic
	vp := reflect.ValueOf(&x)
	fmt.Println("CanSet:", vp.CanSet())
	vpelem := vp.Elem()
	fmt.Println("Elem CanSet:", vpelem.CanSet())
	vpelem.SetFloat(3.5)
	fmt.Println("value:", x)

	// Structs
	type T struct {
		A int
		B string
	}
	t := T{1, "hello"}
	vt := reflect.ValueOf(&t)
	s := vt.Elem()
	typeofT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeofT.Field(i).Name, f.Type(), f.Interface())
	}
	s.Field(0).SetInt(2)
	s.Field(1).SetString("world")
	fmt.Println(t)
}
