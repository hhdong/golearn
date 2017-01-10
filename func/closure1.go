package main

import (
	"fmt"
	"reflect"
)

func main() {
	fv := func() {
		fmt.Printf("Hello world!\n")
	}
	fv()
	fmt.Printf("%s\n", reflect.TypeOf(fv))
}
