package main
import "fmt"
func try(b int)(x,y int){
	x=b+5
	y=b+10
	return
}
func main(){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	a,b:=try(num)
	fmt.Println("Values returned:",a,b)
}