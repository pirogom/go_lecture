package main

import "fmt"

func main() {
	v := []int{2, 3, 5, 7, 11, 13}
	s := v[1:4]

	fmt.Println("array:", v)
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, 100, 101)
	fmt.Println("array:", v)
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))

	s[0] = 200
	s[1] = 201
	fmt.Println("array:", v)
	fmt.Println("slice:", s, "len:", len(s), "cap:", cap(s))
}
