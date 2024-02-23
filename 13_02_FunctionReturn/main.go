package main

import "fmt"

func functionReturn() func() string {
	return func() string {
		return "피로곰의 모두의 프린터"
	}
}

func main() {
	rf := functionReturn()
	fmt.Println(rf())
}
