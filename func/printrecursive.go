package main

import (
	"fmt"
)

func main() {
	printrec(1)
}

func printrec(n int) {
	if n > 10 {
		return
	}
	fmt.Printf("%d\n", n)
	printrec(n + 1)
}
