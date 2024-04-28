package main

import "fmt"

func fibonacci(c, done chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	c := make(chan int)
	done := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		done <- 0
	}()
	fibonacci(c, done)
}
