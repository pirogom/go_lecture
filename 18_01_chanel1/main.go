package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)

	go func(c chan bool) {
		time.Sleep(3 * time.Second)
		c <- true
	}(c)

	cd := <-c
	fmt.Println(cd)
}
