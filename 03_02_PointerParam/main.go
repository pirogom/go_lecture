package main

import "fmt"

func nonPointerParam(a int, b int) {
	a = 101
	b = 102
}

func pointerParam(a *int, b *int) {
	*a = 101
	*b = 102
}

func main() {
	a := 1
	b := 2

	fmt.Println(a, b)
	nonPointerParam(a, b)
	fmt.Println(a, b)
	pointerParam(&a, &b)
	fmt.Println(a, b)
}
