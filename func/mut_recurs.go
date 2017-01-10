package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is event:is %t\n", 16, event(16))
	fmt.Printf("%d is odd:is %t\n", 17, event(17))
	fmt.Printf("%d is odd:is %t\n", 18, event(18))
}

func event(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return event(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
