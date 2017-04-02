package main

import (
	"math"
)

type NewtonFunction func(parameter float64) float64

type NewtonCalculator interface {
	Calculate(function NewtonFunction, derivative NewtonFunction,
		guess float64, precision float64) float64
}

type SimpleNewtonCalculator struct {
}

func (c SimpleNewtonCalculator) Calculate(function NewtonFunction,
	derivative NewtonFunction, guess float64, precision float64) float64 {
	result := guess - function(guess)/derivative(guess)
	if math.Abs(result-guess) < precision {
		return result
	}
	return c.Calculate(function, derivative, guess, precision)
}
