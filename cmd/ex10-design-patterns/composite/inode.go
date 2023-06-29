package main

import "fmt"

type Inode interface {
	fmt.Stringer
	Clone() Inode
}
