package main

import "fmt"

func main() {
	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}

	fmt.Println(pow)
	for i, v := range pow {
		fmt.Println("Index:", i, "Value:", v)
		v = i + 1000
	}
	fmt.Println(pow)
}
