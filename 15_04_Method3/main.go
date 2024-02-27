package main

import (
	"fmt"
	"strconv"
)

type Util struct {
}

func (u Util) Atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func (u Util) Atof(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

func main() {
	fmt.Println(Util{}.Atoi("128"))
	fmt.Println(Util{}.Atof("3.14"))
}
