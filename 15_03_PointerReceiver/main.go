package main

import "fmt"

type MyString string

func (s MyString) Set(v string) {
	fmt.Printf("Value : %v Address : %p\n", v, &s)
	s = MyString(v)
}

func (s *MyString) SetPointer(v string) {
	fmt.Printf("Value : %v Address : %p\n", v, s)
	*s = MyString(v)
}

func main() {
	s := MyString("피로곰")
	fmt.Printf("value : %v address : %p\n", s, &s)

	s.Set("모두의프린터")
	fmt.Println(s)
	s.SetPointer("모두의프린터")
	fmt.Println(s)
}
