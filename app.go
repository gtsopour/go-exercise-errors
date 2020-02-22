package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return x, ErrNegativeSqrt(x)
	}
    z := float64(1)
    for {
      t := z
      z -= (z*z - x) / (2*z)
      //fmt.Printf("z %v, t = %v\n", z, t)
      if math.Abs(t-z) < 1e-6 {
        break
      }
    }
    return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
