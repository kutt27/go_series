package main

import "fmt"

func addition(arr1[] float32, num int) float32{
	var sum float32
    for i := 0; i < num; i++ {
        sum += arr1[i]
    }
    return sum
}


func substraction(arr1[] float32, num int) float32{
	var total float32
    for i := 0; i < num; i++ {
        total -= arr1[i]
    }
    return total
}

func multiplication(arr1[] float32, num int) float32{
	var total float32
    for i := 0; i < num; i++ {
        total *= arr1[i]
    }
    return total
}

func division(a float32, b float32) float32{
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func main() {
	var operation int
	arr1 := new([10]float32)

	fmt.Println("Basic Calculator: ")
	fmt.Println()

	fmt.Print("Select the operation you want to print: ")
	fmt.Println("\n1. Addition \n2. Subtraction \n3. Multiplication \n4. Division")
	fmt.Scan(&operation)

	var num int

	switch operation {
		case 1:
			fmt.Println("How many elements you want to add: ")
			fmt.Scan(&num)
			for i := 0; i < num; i++ {
				for i := 0; i < num; i++ {
					fmt.Printf("Enter the %d number: ", i+1)
					fmt.Scanf("%f", &arr1[i])
					fmt.Scanln() // Consume newline character
				}
				
			}
			fmt.Println(addition(arr1[:num], num))
		case 2:
			fmt.Println("How many elements you want to add: ")
			fmt.Scan(&num)
			for i := 0; i < num; i++ {
				fmt.Printf("Enter the %d number: ", i)
				fmt.Scanf("%f", &arr1[i])
			}
			fmt.Println(substraction(arr1[:num], num))
		case 3:
			fmt.Println("How many elements you want to add: ")
			fmt.Scan(&num)
			for i := 0; i < num; i++ {
				fmt.Printf("Enter the %d number: ", i)
				fmt.Scanf("%f", &arr1[i])
			}
			fmt.Println(multiplication(arr1[:num], num))
		case 4:
			var a float32 = 0
			var b float32 = 0
			fmt.Println("Which number you want to divide: ")
			fmt.Scanf("%f", &a)
			fmt.Println("By which number: ")
			fmt.Scanf("%f", &b)
			fmt.Print(division(a, b))
		default:
			fmt.Println("Invalid operation")
	}
}