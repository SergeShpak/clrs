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
