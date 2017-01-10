package main

import (
	"fmt"
)

func main() {
	x := min(1, 2, 3, 0)
	fmt.Printf("The minimum is:%d\n", x)

	arr := []int{3, 65, 13, 6}
	x = min(arr...)
	fmt.Printf("The minimum is:%d\n", x)
}
func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

func min_2(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}
