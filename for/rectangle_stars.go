package main

import "fmt"

const (
	w = 10
	h = 15
)

func main() {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i == 0 || i == (h-1) {
				fmt.Print("*")
			} else {
				if j == 0 || j == w-1 {
					fmt.Print("*")
				} else {

					fmt.Print(" ")
				}
			}
		}
		fmt.Println("")
	}
}
