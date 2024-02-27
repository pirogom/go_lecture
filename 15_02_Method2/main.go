package main

import (
	"fmt"
)

type Version struct {
	Major   int16
	Minor   int16
	AppName string
}

func (v Version) Version() string {
	return fmt.Sprintf("%d.%d", v.Major, v.Minor)
}

func (v Version) Title() string {
	return fmt.Sprintf("%s(%s)", v.AppName, v.Version())
}

func main() {
	v := Version{3, 16, "모두의 프린터"}
	fmt.Println(v.Title())
}
