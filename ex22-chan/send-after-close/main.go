package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	time.AfterFunc(3*time.Second, func() {
		close(ch)
	})
	ch <- 1
	fmt.Println("done.")
}

/* output:
panic: send on closed channel

goroutine 1 [running]:
main.main()
        /root/workspace/golang-playground/ex22-chan/send-after-close/main.go:13 +0x7f
exit status 2
*/
