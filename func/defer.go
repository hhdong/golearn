package main

import (
	"fmt"
)

func main() {
	func1()
}

func func1() {
	fmt.Println("In function 1 at the top\n")
	defer func2()
	fmt.Println("In function 1 at the bottom\n")
}

func func2() {
	fmt.Println("In function 2 at the top\n")
}
