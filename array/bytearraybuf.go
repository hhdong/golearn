package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	fmt.Println(buffer.String())

	var slice1 []byte = new([10]byte)[:]
	for i := 0; i < len(slice1); i++ {
		slice1[i] = (byte)(5 * i)
	}

	slice2 := slice1[:]
	slice3 := slice1[2:5]
	slice4 := append(slice2, slice3)

	for i := 0; i < len(slice4); i++ {
		fmt.Printf("%d ", (int)(slice4[i]))
	}
	s2 := string(slice2)
	fmt.Println(s2)

}

func sliceslice(slice []byte, n int) ([]byte, []byte) {
	return slice[:n], slice[n:]
}
func append(slice, data []byte) []byte {
	var buffer bytes.Buffer
	buffer.Write(slice)
	buffer.Write(data)
	fmt.Println(len(buffer.Bytes()))
	return buffer.Bytes()
}
