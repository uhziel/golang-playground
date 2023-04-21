package main

import (
    "fmt"
)

func appendInt(slice []int, elems ...int) []int {
    n := len(slice)
    m := n + len(elems)
    if m > cap(slice) {
        tmp := make([]int, cap(slice) * 2)
        copy(tmp, slice)
        slice = tmp
    }
    slice = slice[0:m]
    copy(slice[n:m], elems)
    return slice
}

func main() {
    s := []int{1, 2, 3}
    s = appendInt(s, 4, 5)
    fmt.Println(s)
}
