package main

import "fmt"

func one(xPtr*int){
	*xPtr = 15
}
func main() {
	var x int
	one(&x)
	fmt.Println(x)
}