package main

import (
    "sync"
    "fmt"
)

func main() {
    var onceA, onceB sync.Once
    var funcA func()
    funcB := func() { onceA.Do(funcA) }
    funcA = func() { onceB.Do(funcB) }
    onceA.Do(funcA)

    fmt.Printf("normal exit")
}
