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

type Mop struct {
	Owner   string
	Product string
}

func (m Mop) Say() {
	fmt.Println(m.What())
}

func (m Mop) What() string {
	return fmt.Sprintf("안녕하세요. %s 개발자 %s입니다.", m.Product, m.Owner)
}

func SaySomething(s SayWhat) {
	s.Say()
}

func main() {
	p := Pirogom("안녕하세요. 모두의프린터 개발자 피로곰입니다.")
	SaySomething(p)

	m := Mop{"피로곰", "모두의프린터"}
	SaySomething(m)
}
