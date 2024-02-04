package main

import "fmt"

func main() {
	names := [4]string{"피로곰", "모두의", "프린터", "만세"}

	fmt.Println(names)

	a := names[0:2]
	fmt.Println("A Slice:", a)

	b := names[1:3]
	fmt.Println("B Slice:", b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
}
