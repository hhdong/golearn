package main

import "fmt"

func main() {
	k := 6
	switch k {
	case 4:
		fmt.Printf("was <=4\n")
	case 5:
		fmt.Printf("was <=5\n")
	case 6:
		fmt.Printf("was <=6\n")
	default:
		fmt.Println("default case")
	}
}
