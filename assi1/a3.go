package main

import "fmt"

func main() {
	a := 10
	b := 20

	fmt.Println("Before Swap:")
	fmt.Println("a =", a, "b =", b)

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("After Swap:")
	fmt.Println("a =", a, "b =", b)
}

