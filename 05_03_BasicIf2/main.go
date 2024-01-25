package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	var v float64
	if v = math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow3(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else if v == lim {
		fmt.Println("v == lim")
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
	fmt.Println(pow2(3, 2, 10), pow2(3, 3, 20))
	fmt.Println(pow3(3, 2, 10), pow3(3, 3, 20))
}
