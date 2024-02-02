package main

import (
	"fmt"
	"os"
)

func deferFile() {
	f, err := os.Open("./test.txt")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	b := make([]byte, 128)
	f.Read(b)

	fmt.Println(string(b))
}

func main() {
	deferFile()
}
