package main

import (
	"fmt"
)

func printArry(arr []int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\n")
}

func main() {
	var arr [15]int
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	printArry(arr[:])
}
