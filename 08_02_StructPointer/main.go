package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	fmt.Println(v)
	p := &v

	//(*p).X = 100
	p.X = 100
	fmt.Println(v)
}
