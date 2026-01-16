package main
import(
	"fmt"
	"errors"
)
func f(a int )(int,error){
	if a<0{
		return -1,errors.New("Negative number error")
	}
	return a*2,nil
}
func main(){
	res,err:=f(-5)
	if err!=nil{
		fmt.Println("Error:",err)
	}else{
		fmt.Println("Result:",res)
	}
}	