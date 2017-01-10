package main

import "fmt"

var sum, mul, sub int

func main() {
	sum, mul, sub = cal(12, 3)
	fmt.Printf("sum=%d mul=%d sub=%d", sum, mul, sub)

}

func cal(a int, b int) (sum int, mul int, sub int) {
	sum = a + b
	mul = a * b
	sub = a - b
	return
}
