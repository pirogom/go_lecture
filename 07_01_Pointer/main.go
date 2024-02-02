package main

import "fmt"

func main() {
	i := 100

	// p := &i

	var p *int
	p = &i
	fmt.Println(p, &i)

	fmt.Println(*p)

	*p = *p / 2
	fmt.Println(*p, i)
}
