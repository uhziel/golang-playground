package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

func MD5Sum(root string) (map[string][md5.Size]byte, error) {
	results := make(map[string][md5.Size]byte)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil { // 这个判断应该放在最前，防止出现 root 为 `...` 出现 d 为 nil 的情况
			return err
		}

		if !d.Type().IsRegular() { // 防止 path 是链接但链接到 dir 的情况，os.ReadFile() 会 follow ln 导致因为是目录而失败
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		s := md5.Sum(data)

		results[path] = s

		return nil
	})

	if err != nil {
		return results, err

	}
	return results, nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]

	results, err := MD5Sum(path)
	if err != nil {
		panic(err)
	}

	paths := make([]string, 0, len(results))
	for k := range results {
		paths = append(paths, k)
	}
	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x %s\n", results[path], path)
	}
}
