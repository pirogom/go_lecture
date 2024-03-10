package main

import "fmt"

type SayWhat interface {
	Say()
	What() string
}

type Pirogom string

func (p Pirogom) Say() {
	fmt.Println(p.What())
}

func (p Pirogom) What() string {
	return string(p)
}

func SaySomething(s SayWhat) {
	s.Say()
}

func main() {
	p := Pirogom("안녕하세요. 모두의프린터 개발자 피로곰입니다.")
	SaySomething(p)
}
