package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}
func untrace(s string) { fmt.Println("leaving:", s) }
func un(s string)      { fmt.Println("leaving:", s) }
func a() {
	defer un(trace("a"))
	//trace("a")
	//defer untrace("a")
	fmt.Println("in a")
	b()
}
func b() {
	defer un(trace("a"))
	//trace("b")
	//defer untrace("b")
	fmt.Println("in b")
}
func main() {
	a()
}
