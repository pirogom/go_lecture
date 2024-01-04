package main

import (
	"fmt"
	"math"
	"math/rand" // math 는 rand 패키지의 경로 일뿐 패키지명은 rand
)

func main() {
	fmt.Println(math.Abs(-100))      // math 패키지의 Abs
	fmt.Println(rand.Intn(100))      // math/rand 패키지의 Intn
	fmt.Println(math.rand.Intn(100)) // 오류!
}
