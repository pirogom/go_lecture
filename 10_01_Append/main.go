package main

import "fmt"

func main() {
	v := []int{1, 2, 3}
	fmt.Println(v) // [1 2 3]

	v = append(v, 4)
	fmt.Println(v) // [1 2 3 4]

	v = append(v, 5, 6)
	fmt.Println(v) // [1 2 3 4 5 6]

	x := []int{7, 8}
	v = append(v, x...)
	fmt.Println(v) // [1 2 3 4 5 6 7 8]

	v = append(v, []int{9, 10}...)
	fmt.Println(v) // [1 2 3 4 5 6 7 8 9 10]
}
