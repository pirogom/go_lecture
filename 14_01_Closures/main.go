package main

import "fmt"

func printNumber() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	plus, minus := printNumber(), printNumber()
	for i := 0; i < 10; i++ {
		fmt.Println(i, plus(i), minus(-i))
	}
}
