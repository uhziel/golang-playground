package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// https://golang.google.cn/pkg/os/exec

func main() {
	cmd := exec.Command("cat", "/etc/shadow")
	err := cmd.Run()
	log.Printf("demo: %v", err)

	cmd = exec.Command("git", "show", "46fd1053bde35f12774c08599e8731e053a607e7:golang/demo-exec/demo-exec.go")
	cmd.Dir = "."
	err = cmd.Run() // output: err is nil
	log.Printf("git show: %v", err)

	cmd = exec.Command("git", "show", "46fd1053bde35f12774c08599e8731e053a607e7:golang/demo-exec/demo-exec.go1")
	err = cmd.Run() // output: err is *os/exec.ExitError
	log.Printf("git show: %v", err)

	//检查下输出 https://pkg.go.dev/os/exec?utm_source=gopls#Cmd.StdoutPipe
	cmd = exec.Command("git", "ls-tree", "--name-only", "46fd1053bde35f12774c08599e8731e053a607e7", "./")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	cmd.Wait()
	/*
	   2021/05/18 17:10:22 demo: exit status 1
	   2021/05/18 17:10:22 git show: <nil>
	   2021/05/18 17:10:37 git show: exit status 128
	*/

	// 演示使用 stdin via https://golang.org/src/os/exec/example_test.go
	cmd = exec.Command("cat", "-")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	io.WriteString(stdin, "test_stdin: hello, zhulei")
	stdin.Close()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
