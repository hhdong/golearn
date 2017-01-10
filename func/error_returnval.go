package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1)
	}
	ret2, err2 := MySqrt(1)
	if err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2)
	}
}

func MySqrt(x float64) (float64, error) {
	if x < 0 {
		return float64(math.NaN()), errors.New("x < 0")
	}
	return math.Sqrt(x), nil
}
func MySqrt_2(x float64) (ret float64, err error) {
	if x < 0 {
		ret = float64(math.NaN())
		err = errors.New("x<0")
		return
	}
	ret = float64(math.Sqrt(x))
	err = nil
	return
}
