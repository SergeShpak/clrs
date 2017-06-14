package ch2

import (
	"testing"
)

func Test_LinearSearch_Sanity(t *testing.T) {
	multiset := []int{3, 7, 4, 2, 1, 12, -2, 0, -15}
	el := 1
	elPos, isFound := LinearSearch(el, multiset)
	expectedPos := 4
	if !isFound || elPos != expectedPos {
		t.Errorf("Expected: %d -- Actual: %d", expectedPos, elPos)
	}
}

func Test_LinearSearch_ElementNotPresent_ReturnFalse(t *testing.T) {
	multiset := []int{3, 7, 4, 2, 1, 12, -2, 0, -15}
	el := 100
	elPos, isFound := LinearSearch(el, multiset)
	if isFound {
		t.Errorf("Expected: not found -- Actual: found, returned position: %d",
			elPos)
	}
}

func Test_BinaryNumber_Initialize_Sanity(t *testing.T) {
	stringRepresentation := "11101"
	result, err := InitializeBinaryNumber(stringRepresentation)
	if err != nil {
		t.Error(err)
	}
	expected := []bool{true, true, true, false, true}
	if len(result.Arr) != len(stringRepresentation) {
		t.Errorf("Expected length: %d -- Actual length: %d",
			len(stringRepresentation), len(result.Arr))
	}
	for i, el := range result.Arr {
		if el != expected[i] {
			t.Errorf("On position %d: Expected element: %t -- Actual element: %t",
				i, el, expected[i])
		}
	}
}

func Test_BinaryNumber_ToString_Sanity(t *testing.T) {
	stringRepresentation := "11101"
	result, err := InitializeBinaryNumber(stringRepresentation)
	if err != nil {
		t.Error(err)
	}
	numberStr := result.ToString()
	if stringRepresentation != numberStr {
		t.Errorf("Expected: %s -- Actual: %s", stringRepresentation, numberStr)
	}
}

func Test_BinaryNumber_ToString_TrimmingLeadingZeroes(t *testing.T) {
	stringRepresentation := "00011"
	result, err := InitializeBinaryNumber(stringRepresentation)
	if err != nil {
		t.Error(err)
	}
	numberStr := result.ToString()
	expectedRepresentation := "11"
	if expectedRepresentation != numberStr {
		t.Errorf("Expected: %s -- Actual: %s", expectedRepresentation, numberStr)
	}
}

func Test_BinaryNumber_ToString_EmptyBinaryResultsInEmptyString(t *testing.T) {
	var number BinaryNumber
	number.Arr = make([]bool, 0)
	if number.ToString() != "" {
		t.Error("Result string is not empty")
	}
}

func Test_BinaryNumber_Add_Sanity(t *testing.T) {
	first, err := InitializeBinaryNumber("101")
	if err != nil {
		t.Error(err)
	}
	second, err := InitializeBinaryNumber("111")
	if err != nil {
		t.Error(err)
	}
	result := Add(first, second)
	expectedStr := "1100"
	resultStr := result.ToString()
	if resultStr != expectedStr {
		t.Errorf("Expected: %s -- Actual: %s", expectedStr, resultStr)
	}
}

func Test_BinaryNumber_Add_SecondNumberShorter(t *testing.T) {
	first, err := InitializeBinaryNumber("1101")
	if err != nil {
		t.Error(err)
	}
	second, err := InitializeBinaryNumber("11")
	if err != nil {
		t.Error(err)
	}
	result := Add(first, second)
	expectedStr := "10000"
	resultStr := result.ToString()
	if resultStr != expectedStr {
		t.Errorf("Expected: %s -- Actual: %s", expectedStr, resultStr)
	}
}

func Test_BinaryNumber_Add_FirstNumberShorter(t *testing.T) {
	first, err := InitializeBinaryNumber("11")
	if err != nil {
		t.Error(err)
	}
	second, err := InitializeBinaryNumber("1101")
	if err != nil {
		t.Error(err)
	}
	result := Add(first, second)
	expectedStr := "10000"
	resultStr := result.ToString()
	if resultStr != expectedStr {
		t.Errorf("Expected: %s -- Actual: %s", expectedStr, resultStr)
	}
}
