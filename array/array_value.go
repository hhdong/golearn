package main

import (
	"fmt"
)

func printArry(arr [3]int) {
	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\n")
}

func main() {
	var arr [3]int
	for i := 0; i < 3; i++ {
		arr[i] = i
	}
	printArry(arr)
	arr2 := arr
	arr2[2] = 89
	printArry(arr2)
}
