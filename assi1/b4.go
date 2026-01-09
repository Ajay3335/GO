package main

import "fmt"

func main() {
	var x int = 42
	var p *int = &x
	var pp **int = &p

	fmt.Println("x =", x)
	fmt.Println("p (address of x) =", p)
	fmt.Println("*p (value of x) =", *p)
	fmt.Println("pp (address of p) =", pp)
	fmt.Println("*pp (which is p) =", *pp)
	fmt.Println("**pp (value of x) =", **pp)

	// Modify x through the pointer-to-pointer
	**pp = 100
	fmt.Println("after **pp = 100, x =", x)
}
