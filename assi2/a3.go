package main
import "fmt"
func isPalindrome(num int) int{
	rev:=0
	temp:=num
	for num>0{
		digit:=num%10
		rev=rev*10+digit
		num=num/10
	}
	if rev==temp{
		return 1
	}
	return 0
}
func main(){
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if isPalindrome(num)==1{
		fmt.Println(num,"is a palindrome")
	}else{
		fmt.Println(num,"is not a palindrome")
	}
}
