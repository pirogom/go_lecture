package main

import "fmt"

func deferEx() {
	fmt.Println("deferEx 함수 시작")
	defer fmt.Println("deferEx 함수 종료")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d\n", i)
	}
}

func main() {
	deferEx()
}
