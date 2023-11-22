package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

func walk(done <-chan bool, root string) (<-chan string, <-chan error) {
	chPaths := make(chan string)
	chErr := make(chan error, 1)

	go func() {
		defer close(chPaths)

		chErr <- filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case chPaths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()

	return chPaths, chErr
}

func digester(done <-chan bool, paths <-chan string, results chan<- result) {
	for {
		select {
		case path, ok := <-paths:
			if !ok {
				return
			}
			data, err := os.ReadFile(path)
			if err != nil {
				results <- result{
					path: path,
					err:  err,
				}
				continue
			}

			results <- result{
				path: path,
				sum:  md5.Sum(data),
			}
		case <-done:
			return
		}
	}
}

func md5all(root string) (map[string][md5.Size]byte, error) {
	done := make(chan bool)
	defer close(done)

	ans := make(map[string][md5.Size]byte)
	paths, errc := walk(done, root)

	results := make(chan result)
	var wg sync.WaitGroup
	workerNum := 10
	wg.Add(workerNum)
	for i := 0; i < workerNum; i++ {
		go func() {
			digester(done, paths, results)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		if result.err != nil {
			return nil, result.err
		}

		ans[result.path] = result.sum
	}

	err := <-errc

	return ans, err
}

func main() {
	root := "."
	results, err := md5all(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	paths := []string{}
	for k := range results {
		paths = append(paths, k)
	}
	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x %s\n", results[path], path)
	}
}
