package main

import "fmt"

func main() {
	file1 := &File{Name: "file1"}
	file2 := &File{Name: "file2"}
	file3 := &File{Name: "file3"}
	folder1 := &Folder{Name: "folder1", Children: []Inode{
		file2,
		file3,
	}}
	folder2 := &Folder{Name: "folder2", Children: []Inode{
		file1,
		folder1,
	}}

	fmt.Println("the tree of folder2:")
	fmt.Print(folder2)

	fmt.Println("the tree of folder2 copy:")
	fmt.Print(folder2.Clone())
}
