package main

import "fmt"

func basicFunc() {
	fmt.Println("basicFunc:", "Basic Func")
}

func 이거슨한글함수() {
	fmt.Println("이거슨한글함수:", "세종대왕만세!!")
}

func basicParam(x int, y int) {
	fmt.Println("basicParam:", x, y)
}

func basicParam2(x, y int) {
	fmt.Println("basicParam2:", x, y)
}

func basicReturn(x int, y int) int {
	return x + y
}

func basicReturn2(x int, y int) (int, int) {
	return x + 100, y + 100
}

func returnName(x int, y int) (r1 int) {
	r1 = x + y
	return
}

func returnName2(x int, y int) (r1 int, r2 int) {
	r1 = x + 10
	r2 = x + 20
	return
}

func returnName3(x int, y int) (r1, r2 int) {
	r1 = x + 10
	r2 = x + 20
	return
}

func returnEx() (msg1, msg2, msg3, msg4, msg5 string) {
	msg1 = "첫번째 메시지"
	msg2 = "패스!"
	msg3 = "세번째 메시지"
	msg4 = "패스!"
	msg5 = "다섯번째 메시지"
	return
}

func main() {
	basicFunc()

	이거슨한글함수()

	basicParam(1, 2)

	basicParam2(3, 4)

	fmt.Print("basicReturn: ")
	fmt.Println(basicReturn(100, 101))

	fmt.Print("basicReturn2: ")
	fmt.Println(basicReturn2(1000, 2000))

	fmt.Println("returnName:", returnName(1, 2))

	fmt.Print("returnName2: ")
	fmt.Println(returnName2(1, 2))

	fmt.Print("returnName3: ")
	fmt.Println(returnName3(1, 2))

	fmt.Print("returnEx: ")
	fmt.Println(returnEx())

	fmt.Print("returnEx Pass: ")
	msg1, _, msg3, _, msg5 := returnEx()
	fmt.Println(msg1, msg3, msg5)
}
