package main

import "fmt"

// Operations on Channel
// There are two major operations which can be done on a channel

// Send
// ch <- val

// Receive
// val := <- ch

func main() {

	var (
	// ok bool
	)

	var err chan error

	errChan := make(chan error)

	ch := make(chan int, 3)

	// first assignment
	ch <- 5

	// second assignment
	ch <- 6

	// channel closing
	close(ch)

	// third assignment
	// it will product panic: send on closed channel
	// it happens becase channel is closed
	// ch <- 7

	fmt.Println(err)
	fmt.Println(errChan)
	fmt.Println(Sum(ch))

}

func Sum(ch chan int) string {
	sum := 0
	for val := range ch {
		sum += val
	}
	return fmt.Sprintf("Sum: %d\n", sum)
}
