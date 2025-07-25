package calculator

import (
	"errors"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divide by 0 is not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	return math.Sqrt(a), nil
}
