package main

import (
	"fmt"
)

func splitstring1(str string, i int) (s1, s2 string) {
	orig := []byte(str)
	_s1 := orig[:i]
	_s2 := orig[i:]
	return string(_s1), string(_s2)
}

func reverstring(s string) string {
	orig := []byte(s)
	for i := 0; i < len(orig)/2; i++ {
		tmp := orig[i]
		orig[i] = orig[len(orig)-i-1]
		orig[len(orig)-i-1] = tmp
	}
	return string(orig)
}
func cal(a, b int) int {
	return a * b
}

func mapfunc(f func(int, int) int, data ...int) []int {

	s := make([]int, len(data))
	for k, v := range data {
		s[k] = f(v, 10)
	}
	return s
}
func main() {
	s1, s2 := splitstring1("hello", 2)
	fmt.Println(s1)
	fmt.Println(s2)

	str := "hello world"
	str1 := str[len(str)/2:] + str[:len(str)/2]
	fmt.Println(str1)

	str2 := reverstring("google")
	fmt.Println(str2)

	ret := mapfunc(cal, 1, 2, 3, 4, 5)
	for _, v := range ret {
		fmt.Printf("%d ", v)
	}

}
