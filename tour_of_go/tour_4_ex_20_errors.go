package main

import (
	"fmt"
	"math"
)

// type for custom error handling
type ErrNegativeSqrt float64

// make type ErrNegativeSqrt implement interface 'error'
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt_4(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func tour_4_ex_error() {
	fmt.Println(Sqrt_4(2))   // 1.4142135623730951 <nil>
	fmt.Println(Sqrt_4(-2))  // 0 cannot Sqrt negative number: -2.000000
}
