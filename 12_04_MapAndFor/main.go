package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["모프"] = "모두의프린터"
	m["모피"] = "모두의PDF"
	m["모플"] = "모두의플러그"
	m["모리"] = "모두의리모트"

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
