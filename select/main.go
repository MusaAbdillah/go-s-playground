package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		case <-time.After(time.Second * 1):
			fmt.Println("Timeout")
			break
		}
	}

}

func goOne(ch chan string) {
	// time.Sleep(1 * time.Second)
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	// time.Sleep(1 * time.Second)
	ch <- "From goTwo goroutine"
}
