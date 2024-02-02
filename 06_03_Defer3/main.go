package main

import "fmt"

func main() {
	fmt.Println("시작")
	defer fmt.Println("디퍼1")
	defer fmt.Println("디퍼2")
	defer fmt.Println("디퍼3")
	fmt.Println("종료")
}
