package main

import "fmt"

const Pi = 3.14
const Mid = 1 << 99
const Big = 1 << 100

func main() {
	const (
		World        = "World"
		NumConst int = 123
	)

	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day", NumConst)
	fmt.Printf("Type : %T, Value : %v\n", World, World)
	//fmt.Printf("Type : %T, Value : %v\n", Big, Big)
	if Mid < Big {
		fmt.Println("Mid < Big")
	}
}
