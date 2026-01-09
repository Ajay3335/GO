package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Choose operation:")
	fmt.Println("1. Addition (+)")
	fmt.Println("2. Subtraction (-)")
	fmt.Println("3. Multiplication (*)")
	fmt.Println("4. Division (/)")
	fmt.Print("Enter choice (1-4 or + - * /): ")
	fmt.Scan(&op)

	switch op {
	case "1", "+":
		fmt.Printf("Result: %v\n", a+b)
	case "2", "-":
		fmt.Printf("Result: %v\n", a-b)
	case "3", "*":
		fmt.Printf("Result: %v\n", a*b)
	case "4", "/":
		if b == 0 {
			fmt.Println("Error: division by zero")
		} else {
			fmt.Printf("Result: %v\n", a/b)
		}
	default:
		fmt.Println("Invalid choice")
	}
}