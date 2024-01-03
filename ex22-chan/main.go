package main

func main() {
	ch := make(chan int)

	closerecvch(ch)
}

func closerecvch(ch <-chan int) {
	close(ch)
}
