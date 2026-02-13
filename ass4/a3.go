package main
import "fmt"
type author struct {
	name string
	age int
}

func (a author) show() {
	fmt.Printf("Author Name: %s, Age: %d\n", a.name, a.age)
}
func main() {
	a := author{name: "Ajay Mahajan", age: 22}
	a.show()
}