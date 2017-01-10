package trans

import (
	"fmt"
	"math"
)

var Pia float64

func init() {
	Pia = 4 * math.Atan(1)
	sayHello()
}

func sayGoodBye() {
	fmt.Printf("sayGoodBye")
}
