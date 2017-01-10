package main

import (
	"fmt"
)

func main() {
	week := map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 7: "Sunday"}
	for k, v := range week {
		fmt.Printf("%d : %s\n", k, v)
	}
	value, isPresent := week[1]
	if isPresent {
		fmt.Print("yes \n")
	} else {
		fmt.Print("no\n")
	}
	fmt.Print(value)
}
