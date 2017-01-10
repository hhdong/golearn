package main

import (
	"fmt"
)

func main() {
	f := fibonacci()
	fmt.Println(f(), f(), f(), f())
}

func fibonacci() func() int {

	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
