package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s, i)
	}
}

func main() {
	go say("Hello, go world!")
	fmt.Println("Start!")
	time.Sleep(6 * time.Second)
}
