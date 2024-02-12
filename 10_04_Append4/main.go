package main

import "fmt"

func main() {
	names := [4]string{
		"John", "Paul", "George", "Ringo",
	}
	s := names[:]
	fmt.Println(s, len(s), cap(s))
	s = append(s[0:2], s[3:]...)
	fmt.Println(s, len(s), cap(s))
}
