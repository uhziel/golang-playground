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

func walk(done <-chan struct{}, root string) (<-chan Result, <-chan error) {
	ch := make(chan Result)
	errCh := make(chan error)
	go func() {
		defer close(ch)

		var wg sync.WaitGroup

		err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.Type().IsRegular() {
				return nil
			}

			wg.Add(1)

			go func() {
				defer wg.Done()

				data, err := os.ReadFile(path)
				r := Result{path: path}

				if err != nil {
					r.err = err
				} else {
					r.data = md5.Sum(data)
				}

				select {
				case <-done:
				case ch <- r:
				}
			}()

			select {
			case <-done:
				return errors.New("user cancel")
			default:
				return nil
			}
		})

		errCh <- err
		wg.Wait()
	}()

	return ch, errCh
}

func MD5Sum(done <-chan struct{}, root string) (map[string][md5.Size]byte, error) {
	results := make(map[string][md5.Size]byte)

	ch, errCh := walk(done, root)

L:
	for {
		select {
		case <-done:
			break L
		case err := <-errCh:
			if err != nil {
				return results, fmt.Errorf("errCh: %w", err)
			}
		case r, ok := <-ch:
			if !ok {
				break L
			}
			if r.err != nil {
				return results, fmt.Errorf("r.err from ch: %w", r.err)
			}
			results[r.path] = r.data
		}
	}

	return results, nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s <path>\n", os.Args[0])
		os.Exit(1)
	}
	path := os.Args[1]

	done := make(chan struct{})
	time.AfterFunc(time.Millisecond, func() {
		close(done)
	})

	results, err := MD5Sum(done, path)
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
