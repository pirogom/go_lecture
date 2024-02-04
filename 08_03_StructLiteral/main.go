package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v  = Vertex{}
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{Y: 2}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v, v1, v2, v3, p)

	v := Vertex{}
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{Y: 2}
	p := &Vertex{1, 2}

	fmt.Println(v, v1, v2, v3, p)
}
