package haha

import "fmt"

var printStr string = "Haha"

func Set(str string) {
	printStr = str
}

func Haha() {
	fmt.Println(printStr)
}
