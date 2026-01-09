package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	if n >= -9 && n <= 9 {
		fmt.Printf("%d is a single-digit number\n", n)
	} else {
		fmt.Printf("%d is not a single-digit number\n", n)
	}
}
