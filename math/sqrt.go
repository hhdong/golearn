package main

import (
	"fmt"
	"math"
)

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}
	return math.Sqrt(f), true
}

func main() {
	t, _ := mySqrt(25.0)
	fmt.Println(t)
}
