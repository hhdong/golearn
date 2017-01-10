package main

import (
	"fmt"
)

func main() {
	//var slice1 []int = make([]int, 10)
	var slice1 []int = new([10]int)[:]
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at%d is %d\n", i, slice1[i])
	}
	fmt.Printf("the length of slice1 is %d\n", len(slice1))
	fmt.Printf("the capcity of slice1 is %d\n", cap(slice1))

	s := make([]byte, 5)
	for i := 0; i < len(s); i++ {
		s[i] = (byte)(i)
	}
	s = s[2:4]
	fmt.Printf("the length of s is %d\n", len(s))
	fmt.Printf("the capcity of s is %d\n", cap(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("s at%d is %d\n", i, s[i])
	}
}
