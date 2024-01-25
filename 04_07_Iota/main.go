package main

import "fmt"

const (
	TYPE_NAME = iota
	TYPE_CITY
	TYPE_ADDR
)

const (
	_ = iota
	NUM1
	NUM2
	NUM3
)

const (
	ENUM1 = (iota + iota)
	ENUM2
	ENUM3
)

func main() {
	fmt.Println(TYPE_NAME, TYPE_CITY, TYPE_ADDR)
	fmt.Println(NUM1, NUM2, NUM3)
	fmt.Println(ENUM1, ENUM2, ENUM3)
}
