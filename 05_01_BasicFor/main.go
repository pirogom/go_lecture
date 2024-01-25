package main

import "fmt"

func basicFor() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func basicFor2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func basicFor3() {
	for {
		fmt.Println("이것은 무한루프입니다!")
	}
}

func main() {
	basicFor()
	basicFor2()
}
