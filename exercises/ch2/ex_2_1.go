package ch2

import (
	"bytes"
	"errors"
	"fmt"
)

type BoxedObject struct {
	Payload interface{}
}

func LinearSearch(element int, multiset []int) (int, bool) {
	for pos, el := range multiset {
		if element == el {
			return pos, true
		}
	}
	return -1, false
}

type BinaryNumber struct {
	Arr []bool
}

func InitializeBinaryNumber(representation string) (BinaryNumber, error) {
	numberLen := len(representation)
	arr := make([]bool, numberLen)
	number := BinaryNumber{arr}
	for i, char := range representation {
		if char != '0' && char != '1' {
			errMsg := fmt.Sprintf("Bad character '%c' in number representation on position %d", char, i)
			err := errors.New(errMsg)
			return number, err
		}
		if char == '0' {
			number.Arr[i] = false
			continue
		}
		number.Arr[i] = true
	}
	return number, nil
}

func (number BinaryNumber) ToString() string {
	if len(number.Arr) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	currentIndex := 0
	for number.Arr[currentIndex] != true && currentIndex < len(number.Arr) {
		currentIndex++
	}
	for currentIndex < len(number.Arr) {
		currEl := number.Arr[currentIndex]
		currentIndex++
		if currEl {
			buffer.WriteString("1")
			continue
		}
		buffer.WriteString("0")
	}
	return buffer.String()
}

func Add(first BinaryNumber, second BinaryNumber) BinaryNumber {
	firstLen := len(first.Arr)
	secondLen := len(second.Arr)
	var resultLen int
	if firstLen > secondLen {
		resultLen = firstLen + 1
	}
	if secondLen >= firstLen {
		resultLen = secondLen + 1
	}
	currentFirstIndex := firstLen - 1
	currentSecondIndex := secondLen - 1
	resultBool := make([]bool, resultLen)
	for i := resultLen - 1; i >= 1; i-- {
		var firstDigit bool
		var secondDigit bool
		if currentFirstIndex >= 0 {
			firstDigit = first.Arr[currentFirstIndex]
		}
		if currentSecondIndex >= 0 {
			secondDigit = second.Arr[currentSecondIndex]
		}
		areTwoDigitsTrue := (firstDigit && secondDigit) || (resultBool[i] && (firstDigit || secondDigit))
		if areTwoDigitsTrue {
			resultBool[i-1] = true
		}
		resultBool[i] = (resultBool[i] != firstDigit) != secondDigit
		currentFirstIndex -= 1
		currentSecondIndex -= 1
	}
	var result BinaryNumber
	result.Arr = resultBool
	return result
}
