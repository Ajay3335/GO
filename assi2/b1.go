package main
import "fmt"
func swap(a*int, b*int){
	temp:=*a
	*a=*b
	*b=temp
}
func main(){
	var x,y int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&x,&y)
	fmt.Println("Before swapping: x =",x,"y =",y)
	swap(&x,&y)
	fmt.Println("After swapping: x =",x,"y =",y)
}