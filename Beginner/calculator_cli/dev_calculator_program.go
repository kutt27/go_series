package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	args := os.Args

	if len(args) != 4 {
		fmt.Println("Usage: calculator <number> <operator> <number>")
		return
	}

	num1, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Println("Invalid number:", args[1])
		return
	}

	operator := args[2]

	num2, err := strconv.ParseFloat(args[3], 64)
	if err != nil {
		fmt.Println("Invalid number:", args[3])
		return
	}

	var result float64
	var errorMessage error

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result, errorMessage = divide(num1, num2)
	default:
		fmt.Println("Invalid operator:", operator)
		return
	}

	if errorMessage != nil {
		fmt.Println("Error:", errorMessage)
		return
	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}