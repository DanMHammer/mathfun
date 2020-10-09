package polynomials

import (
	"errors"
	"math"
	"mathfun/utilities"
)

// Polynomial represented as list of coefficients from constant to nth degree
type Polynomial struct {
	Coefficients []float64
}

// Add - Add 2 polynomials by summing equal degree coefficients
func add(a Polynomial, b Polynomial) Polynomial {
	return Polynomial{utilities.SumSlicesFloat(a.Coefficients, b.Coefficients)}
}

// Equal - Check if 2 polynomials have the same coefficients at each degree
func equal(a Polynomial, b Polynomial) bool {
	return utilities.EqualFloatSlices(a.Coefficients, b.Coefficients)
}

// Multiply - Multiply 2 polynomials by foiling
func multiply(a Polynomial, b Polynomial) Polynomial {
	new := make([]float64, len(a.Coefficients)+len(b.Coefficients)-1)

	for idx, x := range a.Coefficients {
		for idy, y := range b.Coefficients {
			new[idx+idy] += x * y
		}
	}
	return Polynomial{new}
}

// Evaluate - evaluate polynomial a when the variable = x
func evaluate(a Polynomial, x float64) float64 {
	result := 0.0

	for i := 0; i < len(a.Coefficients); i++ {
		result += a.Coefficients[i] * math.Pow(x, float64(i))
	}

	return result
}

// interpolate - Return the unique polynomial of degree at most n passing through the given n+1 points
func interpolate(points [][]float64) (Polynomial, error) {
	if len(points) == 0 {
		return Polynomial{}, errors.New("must provide at least 1 point")
	}

	xvalues := make(map[float64]int)

	for _, val := range points {
		_, exist := xvalues[val[0]]

		if exist {
			return Polynomial{}, errors.New("not all x values are distinct")
		}

		xvalues[val[0]] = 1
	}

	terms := Polynomial{}

	for i := 0; i < len(points); i++ {
		terms = add(terms, singleterm(points, i))
	}

	return terms, nil
}

// singleterm - Return one term of interpolated polynomial
func singleterm(points [][]float64, i int) Polynomial {
	term := Polynomial{[]float64{1}}
	xi, yi := points[i][0], points[i][1]

	for j, p := range points {
		if j == i {
			continue
		}

		xj := p[0]

		term = multiply(term, Polynomial{[]float64{-xj / (xi - xj), 1.0 / (xi - xj)}})
	}

	return multiply(term, Polynomial{[]float64{yi}})
}
