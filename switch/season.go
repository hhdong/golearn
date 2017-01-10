package main

import "fmt"

func getseason(month int) string {

	switch month {
	case 12, 1, 2:
		return "Winter"
	case 3, 4, 5:
		return "Spring"
	case 6, 7, 8:
		return "Summer"
	case 9, 10, 11:
		return "Autumn"
	}
	return "error "
}

func main() {
	fmt.Printf(getseason(3))
}
