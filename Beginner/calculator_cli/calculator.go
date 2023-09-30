package main

import (
	"fmt"
)

func addition(arr []float64) float64 {
	var sum float64
	for _, num := range arr {
		sum += num
	}
	return sum
}

func subtraction(arr []float64) float64 {
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result -= arr[i]
	}
	return result
}

func multiplication(arr []float64) float64 {
	result := 1.0
	for _, num := range arr {
		result *= num
	}
	return result
}

func division(arr []float64) float64 {
	if len(arr) != 2 || arr[1] == 0 {
		fmt.Println("Error: Invalid division input")
		return 0.0
	}
	return arr[0] / arr[1]
}

func main() {
	fmt.Println("Basic Calculator:")
	fmt.Println()

	for {
		var operation int
		fmt.Print("Select the operation (1. Addition, 2. Subtraction, 3. Multiplication, 4. Division, 5. Exit): ")
		fmt.Scan(&operation)

		if operation == 5 {
			fmt.Println("Exiting the calculator. Goodbye!")
			break
		}

		var num int
		fmt.Print("How many numbers do you want to operate on: ")
		fmt.Scan(&num)

		arr := make([]float64, num)

		for i := 0; i < num; i++ {
			fmt.Printf("Enter number %d: ", i+1)
			fmt.Scan(&arr[i])
		}

		var result float64

		switch operation {
		case 1:
			result = addition(arr)
		case 2:
			result = subtraction(arr)
		case 3:
			result = multiplication(arr)
		case 4:
			result = division(arr)
		default:
			fmt.Println("Invalid operation")
			continue
		}

		fmt.Printf("Result: %.2f\n", result)
	}
}
