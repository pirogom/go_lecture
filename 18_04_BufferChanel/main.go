package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1
	fmt.Println("first write")
	ch <- 2
	fmt.Println("second write")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
