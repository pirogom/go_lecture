package main

import (
	"fmt"
	"runtime"
)

func basicSwitch(v string) {
	switch v {
	case "피로곰":
		fmt.Println("모두의프린터")
	case "모두의프린터":
		fmt.Println("피로곰")
	default:
		fmt.Println("피로곰's 모두의 프린터")
	}
}

func basicSwitch2() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MacOS X")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("(Other OS) -", os)
	}
}

func basicSwitch3() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin", "linux", "windows":
		fmt.Println("Famous OS -", os)
	default:
		fmt.Println("Other OS -", os)
	}
}

func main() {
	basicSwitch("피로곰")
	basicSwitch("모두의프린터")
	basicSwitch("아몰랑")

	basicSwitch2()
	basicSwitch3()
}
