package ch2

import (
	"testing"
)

func Test_BinarySearch_Sanity(t *testing.T) {
	multiset := []int{3, 7, 4, 2, 1, 12, -2, 0, -15}
	el := 0
	expectedInd := 2
	actualInd := BinarySearch(multiset, el)
	if actualInd != expectedInd {
		t.Errorf("Expected: %d -- Actual: %d", expectedInd, actualInd)
	}
}

func Test_BinarySearch_MissingElement(t *testing.T) {
	multiset := []int{3, 7, 4, 2, 1, 12, -2, 0, -15}
	el := 42
	expectedInd := -1
	actualInd := BinarySearch(multiset, el)
	if actualInd != expectedInd {
		t.Errorf("Expected: %d -- Actual: %d", expectedInd, actualInd)
	}
}

func Test_BinarySearch_NilSlice(t *testing.T) {
	var multiset []int
	multiset = nil
	el := 0
	expectedInd := -1
	actualInd := BinarySearch(multiset, el)
	if actualInd != expectedInd {
		t.Errorf("Expected: %d -- Actual: %d", expectedInd, actualInd)
	}
}

func Test_BinarySearch_SliceOfZeroLength(t *testing.T) {
	multiset := make([]int, 0)
	el := 0
	expectedInd := -1
	actualInd := BinarySearch(multiset, el)
	if actualInd != expectedInd {
		t.Errorf("Expected: %d -- Actual: %d", expectedInd, actualInd)
	}
}

var twoElementsSumTests = []struct {
	Stock          []int
	Sum            int
	ExpectedResult bool
	ExpectedFirst  int
	ExpectedSecond int
}{
	{[]int{3, 7, 4, 2, 1, 12, -2, 0, -15}, 19, true, 7, 12},
	{[]int{3, 7, 4, 2, 1, 12, -2, 0, -15}, -17, true, -15, -2},
	{nil, 5, false, -1, -1},
	{[]int{3}, 3, false, -1, -1},
}

func Test_TwoElementsSum(t *testing.T) {
	for _, test := range twoElementsSumTests {
		actualResult, actualFirst, actualSecond := TwoElementsSum(test.Stock,
			test.Sum)
		if test.ExpectedResult != actualResult || test.ExpectedFirst != actualFirst || test.ExpectedSecond != actualSecond {
			t.Errorf("TwoElementsSum(%v, %d) => %t, %d, %d, want %t, %d, %d",
				test.Stock, test.Sum, actualResult, actualFirst, actualSecond,
				test.ExpectedResult, test.ExpectedFirst, test.ExpectedSecond)
		}
	}
}
