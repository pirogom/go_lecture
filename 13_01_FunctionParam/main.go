package main

import "fmt"

func sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func functionParam(fn func(args ...int) int) int {
	return fn(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}

func main() {
	fmt.Println(functionParam(sum))
}
