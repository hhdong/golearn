package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s/n", typecheck(1))
	fmt.Printf("%s/n", typecheck("2"))
}

func typecheck(values ...interface{}) string {
	for _, value := range values {
		switch value.(type) {
		case int:
			return "int"
		default:
			return "other"
		}
	}
	return "unknown"
}
