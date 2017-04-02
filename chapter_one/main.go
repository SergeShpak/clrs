package main

import (
	"fmt"
)

type NewtonCalculateFunc func(float64) float64

func main() {
	result := Ex2_2()
	fmt.Printf("Exercise 1.2-2 result: %f\n", result)
	result = Ex2_3()
	fmt.Printf("Exercise 1.2-3 result: %f\n", result)
}
