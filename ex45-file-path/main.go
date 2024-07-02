package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

/*
https://gobyexample.com/file-paths
https://pkg.go.dev/path/filepath?utm_source=gopls#section-documentation
*/

const FILE1 = "/home/zhulei/workspace/xxx-playground/golang/file-path/go.mod"
const FILE2 = "/home/zhulei/workspace/xxx-playground/golang/file-path/foobar"
const FILE3 = "go.mod"
const DIR1 = "/home/zhulei/workspace/xxx-playground/golang/file-path"
const DIR2 = "/home/zhulei/workspace/xxx-playground/golang/file-path/"

func printFilePath(PATH string) {
	fmt.Println("==")
	fmt.Println("PATH:", PATH)
	fmt.Printf("Dir: %s, Base: %s, Ext: %s$\n", filepath.Dir(PATH), filepath.Base(PATH), filepath.Ext(PATH))
	dir, file := filepath.Split(PATH) // Split() 主要是把路径字符串切分为两部分。和上面的 Dir() 有细微的差别，具体请看程序输出
	fmt.Printf("Split (Dir: %s, Base: %s)\n", dir, file)
	fmt.Printf("Join .. : %s\n", filepath.Join(PATH, ".."))
}

func getCheckedFiles(path, dir string) []string {
	var checkedFiles []string
	if !strings.HasPrefix(path, dir) {
		return checkedFiles
	}
	dir = filepath.Join(dir, ".")
	for checkedFile := path; dir != checkedFile; checkedFile = filepath.Dir(checkedFile) {
		checkedFiles = append(checkedFiles, checkedFile+".meta")
	}
	return checkedFiles
}

func main() {
	printFilePath(FILE1)
	printFilePath(FILE2)
	printFilePath(FILE3)
	printFilePath(DIR1)
	printFilePath(DIR2)
	for _, checkedFile := range getCheckedFiles(FILE1, "/home/zhulei/workspace/xxx-playground") {
		fmt.Println(checkedFile)
	}
}
