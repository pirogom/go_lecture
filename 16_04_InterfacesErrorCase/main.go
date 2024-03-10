package main

import "fmt"

type Person interface {
	GetInfo() string
}

type Member struct {
	Name     string
	Age      int
	Position string
}

func (m *Member) GetInfo() string {
	return fmt.Sprintf("Name : %s Age : %d Position : %s", m.Name, m.Age, m.Position)
}

func NewMember(name string, age int, pos string) *Person {
	m := Member{name, age, pos}
	return &m
}

func PrintMemberInfo(p *Person) {
	fmt.Println(p.GetInfo())
}

func main() {
	newMember := NewMember("피로곰", 25, "소프트웨어 엔지니어")
	PrintMemberInfo(newMember)
}
