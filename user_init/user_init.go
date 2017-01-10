package main

import (
	"fmt"
	"golearn/trans"
)

var twoPi = 2 * trans.Pi
var _pi = trans.Pia

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi)
	fmt.Printf("2*Pi = %g\n", _pi)
}
