package main

import (
	"fmt"
	"strings"
	"time"
)

func emptySwitch() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func caseBreak() {
	n := "피로곰's 모두의 프린터"

	switch {
	case strings.Contains(n, "피로곰"):
		fmt.Println("피로곰")
	case strings.Contains(n, "모두의"):
		fmt.Println("모두의")
	case strings.Contains(n, "프린터"):
		fmt.Println("프린터")
	default:
		fmt.Println("없음")
	}
}

func caseContinue() {
	n := "피로곰's 모두의 프린터"

	switch {
	case strings.Contains(n, "피로곰"):
		fmt.Println("피로곰")
		fallthrough
	case strings.Contains(n, "모두의"):
		fmt.Println("모두의")
		fallthrough
	case strings.Contains(n, "프린터"):
		fmt.Println("프린터")
		fallthrough
	default:
		fmt.Println("없음")
	}
}

func main() {
	emptySwitch()
	caseBreak()
	caseContinue()
}
