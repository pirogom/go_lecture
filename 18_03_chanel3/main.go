package main

import (
	"fmt"
	"time"
)

func first(c chan string) {
	c <- "First Write Data"
	r := <-c
	fmt.Println("First Read:", r)
	fmt.Println("First End")
}

func second(c chan string) {
	r := <-c
	fmt.Println("Second Read:", r)
	c <- "Second Write Data"
	fmt.Println("Second End")
}

func main() {
	c := make(chan string)

	go first(c)
	go second(c)

	time.Sleep(3 * time.Second)
}
