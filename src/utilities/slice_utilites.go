package utilities

// SumSlices - find the sum of 2 slices of ints by adding all of the terms
func SumSlices(a []int, b []int) []int {
	values := []int{}

	for idx := 0; idx < len(a) && idx < len(b); idx++ {
		values = append(values, a[idx]+b[idx])
	}

	if len(a) > len(b) {
		values = append(values, a[len(b):]...)
	} else if len(b) > len(a) {
		values = append(values, b[len(a):]...)
	}

	return values
}

// EqualIntSlices - check if 2 slices of ints are the same
func EqualIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// SumSlicesFloat - find the sum of 2 slices of ints by adding all of the terms
func SumSlicesFloat(a []float64, b []float64) []float64 {
	values := []float64{}

	for idx := 0; idx < len(a) && idx < len(b); idx++ {
		values = append(values, a[idx]+b[idx])
	}

	if len(a) > len(b) {
		values = append(values, a[len(b):]...)
	} else if len(b) > len(a) {
		values = append(values, b[len(a):]...)
	}

	return values
}

// EqualFloatSlices - check if 2 slices of float64 are the same
func EqualFloatSlices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
