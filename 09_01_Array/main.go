package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "A World"

	fmt.Println(a)

	var b [2]string = [2]string{"Hello", "B World"}

	fmt.Println(b)

	c := [2]string{}
	c[0] = "Hello"
	c[1] = "C World"

	fmt.Println(c)

	d := [2]string{"Hello", "D World"}

	fmt.Println(d)
}
