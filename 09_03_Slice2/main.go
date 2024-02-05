package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	f := s[:]
	fmt.Println(f)

	v1 := s[1:4]
	fmt.Println(v1)

	v2 := s[:2]
	fmt.Println(v2)

	v3 := s[1:]
	fmt.Println(v3)
}
