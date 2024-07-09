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
	"time"
)

type Result struct {
	path string
	data [md5.Size]byte
	err  error
}

var (
	errUserCancel = errors.New("user cancel")
)

func walkDir(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	ch := make(chan string)
	errCh := make(chan error, 1)
	go func() {
		defer close(ch)
		defer close(errCh)

		errCh <- filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.Type().IsRegular() {
				return nil
			}

			select {
			case <-done:
				return errUserCancel
			case ch <- path:
			}

			return nil
		})
	}()

	return ch, errCh
}

func md5Sum(done <-chan struct{}, in <-chan string) <-chan Result {
	ch := make(chan Result)
	go func() {
		defer close(ch)

		for path := range in {
			r := Result{
				path: path,
			}
			content, err := os.ReadFile(path)
			if err != nil {
				r.err = err
			} else {
				r.data = md5.Sum(content)
			}

			select {
			case <-done:
				return
			case ch <- r:
			}
		}
	}()

	return ch
}

func fanIn(done <-chan struct{}, sources ...<-chan Result) <-chan Result {
	ch := make(chan Result)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(sources))

		for _, s := range sources {
			go func(s <-chan Result) {
				defer wg.Done()
				for r := range s {
					select {
					case <-done:
						return
					case ch <- r:
					}
				}
			}(s)
		}

		go func() {
			wg.Wait()
			close(ch)
		}()
	}()

	return ch
}

const (
	workerNum = 3
)

func MD5Sum(done <-chan struct{}, root string) (map[string][md5.Size]byte, error) {
	ch, errCh := walkDir(done, root)

	var resultChs []<-chan Result
	for i := 0; i < workerNum; i++ {
		resultChs = append(resultChs, md5Sum(done, ch))
	}

	m := make(map[string][md5.Size]byte)
	for r := range fanIn(done, resultChs...) {
		if r.err != nil {
			return m, r.err
		}

		m[r.path] = r.data
	}

	if err := <-errCh; err != nil {
		return m, err
	}

	return m, nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]

	done := make(chan struct{})
	time.AfterFunc(30*time.Millisecond, func() {
		close(done)
	})

	results, err := MD5Sum(done, path)
	if err != nil {
		//fmt.Println(results)
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

	panic("show me the stack")
}
