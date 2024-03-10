package main

import (
	"fmt"
	"math"
)

type MyType interface {
	Get() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Get() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a MyType
	v := Vertex{3, 4}

	a = &v // 구현체 v의 포인터
	a = v  // 오류!

	fmt.Println(a.Get()) // 5
}
