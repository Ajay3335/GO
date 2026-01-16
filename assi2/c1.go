package main
import "fmt"
func callbyvalue() {	
	var a int = 10
	var b int = 20
	fmt.Println("Before swapping in callbyvalue a:", a, "b:", b)
	swap(a, b)
	fmt.Println("After swapping in callbyvalue a:", a, "b:", b)
}
func swap(x int, y int) {
	var temp int
	temp = x
	x = y
	y = temp
}
func main() {
	callbyvalue()
}