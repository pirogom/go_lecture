package main

import (
	"fmt"
	"strconv"
)

type MyString string

func (s MyString) Int32() int32 {
	i, _ := strconv.ParseInt(string(s), 10, 32)
	return int32(i)
}

func (s MyString) Int64() int64 {
	i, _ := strconv.ParseInt(string(s), 10, 64)
	return int64(i)
}

func main() {
	s := MyString("256")
	fmt.Printf("Type: %T Value : %v\n", s, s)
	i32 := s.Int32()
	fmt.Printf("Type: %T Value : %v\n", i32, i32)
	i64 := s.Int64()
	fmt.Printf("Type: %T Value : %v\n", i64, i64)
}
