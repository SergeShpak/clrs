package main

import (
	"math"
)

const DefaultPrecision = 0.01

func Ex2_2() float64 {
	c := SimpleNewtonCalculator{}
	function := func(n float64) float64 {
		return 8*n*n - 64*n*math.Log10(n)
	}
	derivative := func(n float64) float64 {
		return 16*n - 64*(math.Log(n)-1)/math.Ln10
	}
	guess := float64(4)
	result := calculateEfficiencyResults(function, derivative, c, guess,
		DefaultPrecision)
	return result
}

func Ex2_3() float64 {
	c := SimpleNewtonCalculator{}
	function := func(n float64) float64 {
		return math.Pow(2, n) - 100*n*n
	}
	derivative := func(n float64) float64 {
		return math.Pow(2, n)*math.Ln2 - 200*n
	}
	guess := float64(50)
	result := calculateEfficiencyResults(function, derivative, c, guess,
		DefaultPrecision)
	return processExerciseResult(result)
}

func calculateEfficiencyResults(function NewtonFunction,
	derivative NewtonFunction, c NewtonCalculator, guess float64,
	precision float64) float64 {
	result := c.Calculate(function, derivative, guess, precision)
	return processExerciseResult(result)
}

func processExerciseResult(raw float64) float64 {
	return math.Ceil(raw)
}
