package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func walk(root string) []string {
	paths := []string{}
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		paths = append(paths, path)
		return nil
	})
	return paths
}

func digester(path string) result {
	data, err := os.ReadFile(path)
	if err != nil {
		return result{
			path: path,
			err:  err,
		}
	}

	return result{
		path: path,
		sum:  md5.Sum(data),
	}
}

func main() {
	root := "."
	paths := walk(root)

	results := make(map[string][md5.Size]byte)
	for _, path := range paths {
		result := digester(path)
		if result.err != nil {
			fmt.Println(result.err)
		}
		results[result.path] = result.sum
	}

	for _, path := range paths {
		fmt.Printf("%x %s\n", results[path], path)
	}
}
