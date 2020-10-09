package utilities

import (
	"testing"
)

type EqualIntSlicesTestData struct {
	a     []int
	b     []int
	equal bool
}

func TestEqualIntSlices(t *testing.T) {
	items := []EqualIntSlicesTestData{
		{[]int{1, 2, 3, 45, 110}, []int{1, 2, 3, 45, 110}, true},
		{[]int{1, 2, 3, 45, 110}, []int{3, 2, 3, 3, 12}, false},
		{[]int{1, 2, 3, 45, 110}, []int{110, 45, 3, 2, 1}, false},
	}

	for _, item := range items {
		if EqualIntSlices(item.a, item.b) != item.equal {
			t.Errorf("EqualIntSlices incorrect. Got %t. Expected %t", !item.equal, item.equal)
		}
	}
}

type SumSlicesTestData struct {
	a []int
	b []int
	c []int
}

func TestSumSlices(t *testing.T) {
	items := []SumSlicesTestData{
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{5, 7, 9}},
		{[]int{-1, 2, -3}, []int{4, -5, -6}, []int{3, -3, -9}},
	}

	for _, item := range items {
		output := SumSlices(item.a, item.b)

		if !EqualIntSlices(output, item.c) {
			t.Errorf("SumSlices incorrect. Got: %d. Expected %d.", output, item.c)
		}
	}
}
