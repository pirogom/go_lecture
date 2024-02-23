package main

import "fmt"

type sumFuncType func(a int, b int, c int) int
type divFuncType func(a int, b int) int
type returnFuncType func() string

func sum(a int, b int, c int) int {
	return a + b + c
}

func div(a int, b int) int {
	return a / b
}

func functionParam(sf sumFuncType, df divFuncType) returnFuncType {
	return func() string {
		return fmt.Sprintf("Sum : %d Div : %d", sf(1, 2, 3), df(10, 2))
	}
}

func main() {
	rf := functionParam(sum, div)
	fmt.Println(rf())
}
