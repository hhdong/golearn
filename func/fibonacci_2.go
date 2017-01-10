package main

import (
	"fmt"
)

func main() {
	result := 0
	index := 0
	for i := 0; i <= 10; i++ {
		result, index = fibonacci(i)
		fmt.Printf("fibonacci(%d) is:%d\n", index, result)
	}
}

func fibonacci(n int) (res, index int) {
	index = n
	var n1, n2 int
	if n <= 1 {
		res = 1
	} else {
		n1, _ = fibonacci(n - 1)
		n2, _ = fibonacci(n - 2)
		res = n1 + n2
	}
	return
}
