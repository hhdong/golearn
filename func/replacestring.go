package main

import "fmt"

//import "strings"

func main() {
	fmt.Println(replace("hello我是中国asdf"))
}

func replace(s string) string {

	var b []byte
	b = make([]byte, len(s), len(s))
	start := 0
	for start < len(s) {
		if IsAscii((int)(s[start])) == true {
			b[start] = (byte)(s[start])
		} else {
			b[start] = '?'
		}
		start += 1
	}
	return string(b)
}

func IsAscii(c int) bool {
	if c > 255 {
		return false
	}
	return true
}
