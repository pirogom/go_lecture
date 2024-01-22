package main

import "fmt"

var c int
var python int
var java bool

var spring, django, express bool

func main() {
	fmt.Println(c, python, java)
	fmt.Println(spring, django, express)
	localVar()
}

func localVar() {
	var i int
	var spring bool
	spring = true
	fmt.Println(i, c, python, java, spring)
}
