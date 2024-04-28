package main

import (
	"fmt"
)

func firstWrite(c chan string) {
	c <- "First Write data"
}

func secondWrite(c chan string) {
	c <- "Second Write data"
}

func main() {
	c := make(chan string)

	go firstWrite(c)
	go secondWrite(c)

	firstRead := <-c
	fmt.Println(firstRead)
	secondRead := <-c
	fmt.Println(secondRead)
}
