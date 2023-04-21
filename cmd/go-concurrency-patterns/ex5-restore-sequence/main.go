package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Msg  string
	Wait chan<- bool
}

func boring(name string) <-chan Message {
	ch := make(chan Message)
	chWait := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- Message{
				Msg:  fmt.Sprintf("%s say: %d", name, i),
				Wait: chWait,
			}
			time.Sleep(time.Duration((rand.Intn(5) + 1)) * time.Second)

			<-chWait
		}
		close(ch)
	}()
	return ch
}

func fanIn(chs ...<-chan Message) <-chan Message {
	chAllInOne := make(chan Message)
	for _, ch := range chs {
		go func(ch <-chan Message) {
			for {
				if v, ok := <-ch; ok {
					chAllInOne <- v
				}
			}
		}(ch)
	}
	return chAllInOne
}

func main() {
	ch := fanIn(boring("Lee"), boring("John"))

	fmt.Println("I'm listening.")

	for i := 0; i < 2; i++ {
		msg1 := <-ch
		fmt.Println(msg1.Msg)
		msg2 := <-ch
		fmt.Println(msg2.Msg)

		msg1.Wait <- true
		msg2.Wait <- true
	}

	fmt.Println("I'm leaving. You are too boring.")
}
