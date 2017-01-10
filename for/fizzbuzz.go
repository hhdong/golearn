package main

import "fmt"

const (
	Fizz     = 3
	FizzBuzz = 15
	Buzz     = 5
)

func main() {

	for i := 1; i <= 100; i++ {
		switch {
		case i%Fizz == 0:
			fmt.Println("Fizz")
		case i%Buzz == 0:
			fmt.Println("Buzz")
		case i%FizzBuzz == 0:
			fmt.Println("FizzBuzz")
		default:
			fmt.Println(i)
		}
	}
	fmt.Println("---------------")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("Fizz")
			} else {
				fmt.Println("FizzBuzz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
