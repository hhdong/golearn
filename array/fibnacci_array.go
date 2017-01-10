package main

import (
	"fmt"
)

const MAX = 15

func main() {
	var x int
	x = fib(12)
	fmt.Println(x)
}

func fib(n int) int {
	var fibar [MAX + 1]int
	for i := 1; i <= n; i++ {
		if i == 1 || i == 2 {
			fibar[i] = 1
		} else {
			fibar[i] = fibar[i-1] + fibar[i-2]
		}
	}
	return fibar[n]
}
