package main
import "fmt"


func calculateSquaresCubes(num int, squaresChan chan int, cubesChan chan int) {
	str := fmt.Sprintf("%d", num)
	
	for _, digit := range str {
		d := int(digit - '0')
		squaresChan <- d * d
		cubesChan <- d * d * d
	}
	close(squaresChan)
	close(cubesChan)
}

func main() {
	num := 123 
	squaresChan := make(chan int)
	cubesChan := make(chan int)
	
	go calculateSquaresCubes(num, squaresChan, cubesChan)
	
	sumSquares := 0
	sumCubes := 0
	squaresDone := false
	cubesDone := false
	
	for !squaresDone || !cubesDone {
		select {
		case square, ok := <-squaresChan:
			if !ok {
				squaresDone = true
			} else {
				sumSquares += square
			}
		case cube, ok := <-cubesChan:
			if !ok {
				cubesDone = true
			} else {
				sumCubes += cube
			}
		}
	}
	fmt.Printf("Number: %d\n", num)
	fmt.Printf("Sum of squares of digits: %d\n", sumSquares)
	fmt.Printf("Sum of cubes of digits: %d\n", sumCubes)
}

