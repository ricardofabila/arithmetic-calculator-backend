package services_test

import (
	"testing"

	"github.com/ricardofabila/arithmetic-calculator-backend/services"
)

// TestPerformArithmeticOperationAddition tests the addition operation
func TestPerformArithmeticOperationAddition(t *testing.T) {
	operation := "addition"
	num1 := 5.0
	num2 := 3.0
	expected := "8"

	result, err := services.PerformArithmeticOperation(operation, num1, num2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

// TestPerformArithmeticOperationSubtraction tests the subtraction operation
func TestPerformArithmeticOperationSubtraction(t *testing.T) {
	operation := "subtraction"
	num1 := 5.0
	num2 := 3.0
	expected := "2"

	result, err := services.PerformArithmeticOperation(operation, num1, num2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

// TestPerformArithmeticOperationMultiplication tests the multiplication operation
func TestPerformArithmeticOperationMultiplication(t *testing.T) {
	operation := "multiplication"
	num1 := 5.0
	num2 := 3.0
	expected := "15"

	result, err := services.PerformArithmeticOperation(operation, num1, num2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

// TestPerformArithmeticOperationDivision tests the division operation
func TestPerformArithmeticOperationDivision(t *testing.T) {
	operation := "division"
	num1 := 6.0
	num2 := 3.0
	expected := "2"

	result, err := services.PerformArithmeticOperation(operation, num1, num2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

// TestPerformArithmeticOperationDivisionByZero tests division by zero
func TestPerformArithmeticOperationDivisionByZero(t *testing.T) {
	operation := "division"
	num1 := 5.0
	num2 := 0.0

	_, err := services.PerformArithmeticOperation(operation, num1, num2)
	if err == nil {
		t.Errorf("expected an error for division by zero, but got nil")
	}

	expectedErrMsg := "division by zero is not allowed"
	if err != nil && err.Error() != expectedErrMsg {
		t.Errorf("expected error message %s, but got %s", expectedErrMsg, err.Error())
	}
}

// TestPerformArithmeticOperationUnsupportedOperation tests an unsupported operation
func TestPerformArithmeticOperationUnsupportedOperation(t *testing.T) {
	operation := "modulus"
	num1 := 5.0
	num2 := 3.0

	_, err := services.PerformArithmeticOperation(operation, num1, num2)
	if err == nil {
		t.Errorf("expected an error for unsupported operation, but got nil")
	}

	expectedErrMsg := "unsupported operation"
	if err != nil && err.Error() != expectedErrMsg {
		t.Errorf("expected error message %s, but got %s", expectedErrMsg, err.Error())
	}
}

// TestSqrtPositiveNumber tests the Sqrt function with positive numbers
func TestSqrtPositiveNumber(t *testing.T) {
	num := 9.0
	expected := "3"

	result, err := services.Sqrt(num)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

// TestSqrtZero tests the Sqrt function with zero
func TestSqrtZero(t *testing.T) {
	num := 0.0
	expected := "0"

	result, err := services.Sqrt(num)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("expected %s, but got %s", expected, result)
	}
}

// TestSqrtNegativeNumber tests the Sqrt function with a negative number
func TestSqrtNegativeNumber(t *testing.T) {
	num := -9.0

	_, err := services.Sqrt(num)
	if err == nil {
		t.Errorf("expected an error for negative input, but got nil")
	}
}
