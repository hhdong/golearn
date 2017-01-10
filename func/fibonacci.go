package main

import (
	"fmt"
	"time"
)

const LIM = 10

var fibs [LIM]uint64

func main() {
	result := 0
	start := time.Now()
	for i := 0; i <= LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is:%d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("cal fib take time:%s\n", delta)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
