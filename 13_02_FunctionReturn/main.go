package main

import "fmt"

func functionReturn() func() string {
	return func() string {
		return "피로곰의 모두의 프린터"
	}
}

func someText() string {
	return "피로곰의 Go언어 속성강의"
}

func functionReturn2() func() string {
	return someText
}

func main() {
	rf := functionReturn()
	fmt.Println(rf())
	rf = functionReturn2()
	fmt.Println(rf())
}
