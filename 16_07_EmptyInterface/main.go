package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = 3.14
	describe(i)

	v := 256
	i = &v
	describe(i)
}
