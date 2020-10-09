package utilities

import "math"

// IntegerExponent - raise int a to the power int b and return an int
func IntegerExponent(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
