package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start")

	for i := 0; i < 3; i++ {
		go func(idx int) {
			fmt.Println("Goroutine:", idx, i)
		}(i)
	}
	fmt.Println("End")

	time.Sleep(3 * time.Second)
}
