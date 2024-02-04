package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{}
	v.X = 1
	v.Y = 1

	v2 := Vertex{2, 2}

	fmt.Println(v)
	fmt.Println(v2)
}
