package simplemath

import "math"

func Sqrt(value int) int{
	return int(math.Sqrt(float64(value)))
}