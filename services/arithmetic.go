package services

import (
	"errors"
	"math"
	"strconv"
)

// PerformArithmeticOperation Function to perform arithmetic operations
func PerformArithmeticOperation(operation string, num1, num2 float64) (string, error) {
	var result float64

	switch operation {
	case "addition":
		result = num1 + num2
	case "subtraction":
		result = num1 - num2
	case "multiplication":
		result = num1 * num2
	case "division":
		if num2 == 0 {
			return "", errors.New("division by zero is not allowed")
		}
		result = num1 / num2
	default:
		return "", errors.New("unsupported operation")
	}

	return strconv.FormatFloat(result, 'f', -1, 64), nil
}

// Sqrt Function to calculate square root
func Sqrt(num float64) (string, error) {
	if num < 0 {
		return "", errors.New("cannot calculate square root of a negative number")
	}
	
	return strconv.FormatFloat(math.Sqrt(num), 'f', -1, 64), nil
}
