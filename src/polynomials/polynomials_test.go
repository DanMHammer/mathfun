package polynomials

import (
	"math"
	"testing"
)

type EqualPolynomialTestData struct {
	a     Polynomial
	b     Polynomial
	equal bool
}

func TestEqualPolynomials(t *testing.T) {
	items := []EqualPolynomialTestData{
		{Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{1, 2, 3, 4}}, true},
		{Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{-1, -2, -3, -4}}, false},
		{Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{1, 2, 3}}, false},
		{Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{2, 4, 5, 6}}, false},
		{Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{4, 3, 2, 1}}, false},
	}

	for _, item := range items {
		if equal(item.a, item.b) != item.equal {
			t.Errorf("Equal Polynomials incorrect. Got %t. Expected %t", !item.equal, !item.equal)
		}
	}
}

type SumPolynomialTestData struct {
	a Polynomial
	b Polynomial
	c Polynomial
}

func TestAddPolynomials(t *testing.T) {
	items := []SumPolynomialTestData{
		{Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{5, -4, 3, -2, 1, 4}}, Polynomial{[]float64{6, -2, 6, 2, 1, 4}}},
		{Polynomial{[]float64{5, -4, 3, -2, 1, 4}}, Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{6, -2, 6, 2, 1, 4}}},
	}

	for _, item := range items {
		sum := add(item.a, item.b)
		if !equal(sum, item.c) {
			t.Errorf("Add Polynomials incorrect. Got %f. Expected %f", sum.Coefficients, item.c.Coefficients)
		}
	}
}

type EvaluatePolynomialTestData struct {
	a     Polynomial
	x     float64
	value float64
}

func TestEvaluatePolynomial(t *testing.T) {
	items := []EvaluatePolynomialTestData{
		{Polynomial{[]float64{1, 2, 3, 4}}, 2, 49},
		{Polynomial{[]float64{5, -4, 3, -2, 1, 4}}, -3, -793},
	}

	for _, item := range items {
		value := evaluate(item.a, item.x)
		if value != item.value {
			t.Errorf("Add Polynomials incorrect. Got %f. Expected %f", value, item.value)
		}
	}
}

type MultiplyPolynomialTestData struct {
	a Polynomial
	b Polynomial
	c Polynomial
}

func TestMultiplyPolynomial(t *testing.T) {
	items := []MultiplyPolynomialTestData{
		{Polynomial{[]float64{1, 2, 3, 4}}, Polynomial{[]float64{5, -4, 3, -2, 1, 4}}, Polynomial{[]float64{5, 6, 10, 12, -10, 12, 3, 16, 16}}},
		{Polynomial{[]float64{1, 2, 3}}, Polynomial{[]float64{-1, 2}}, Polynomial{[]float64{-1, 0, 1, 6}}},
	}

	for _, item := range items {
		value := multiply(item.a, item.b)
		if !equal(value, item.c) {
			t.Errorf("Multiply Polynomials incorrect. Got %f. Expected %f", value.Coefficients, item.c.Coefficients)
		}
	}
}

type InterPolatePolynomialTestData struct {
	points   [][]float64
	expected Polynomial
}

func TestInterpolatePolynomial(t *testing.T) {
	items := []InterPolatePolynomialTestData{
		// {[][]float64{
		// 	{1, 1},
		// }, Polynomial{[]float64{1}}},
		// {[][]float64{
		// 	{1, 2},
		// }, Polynomial{[]float64{2}}},
		// {[][]float64{
		// 	{1, 1},
		// 	{2, 0},
		// }, Polynomial{[]float64{2, -1}}},
		{[][]float64{
			{1, 1},
			{2, 4},
			{7, 9},
		}, Polynomial{[]float64{-8.0 / 3.0, 4, -1.0 / 3.0}}},
	}

	for _, item := range items {
		value, err := interpolate(item.points)

		if err != nil {
			t.Errorf("Interpolation error")
		}

		// Fuzzy equals to the hundredths place to account for repeating decimals
		for idx, x := range item.expected.Coefficients {
			if math.Round(x*100)/100 != math.Round(value.Coefficients[idx]*100)/100 {
				t.Errorf("Interpolation incorrect. Got %f. Expected %f", value.Coefficients, item.expected.Coefficients)
			}
		}
	}
}
