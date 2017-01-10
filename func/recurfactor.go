package main

import (
	"fmt"
)

func main() {
	fmt.Printf("result %d", factor(3))
}
func factor(n int) (res int) {
	if n == 0 {
		res = 1
	} else {
		res = factor(n-1) * n
	}
	return
}
