package main
import "fmt"
func recursiveSum(n int)int{
	if n==0{
		return 0
	}
	return n + recursiveSum(n-1)


}
func main(){
	fmt.Println("Sum of 1 to 5:", recursiveSum(5))
}