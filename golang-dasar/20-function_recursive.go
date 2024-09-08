package main
import "fmt"

func factorialLoop(value int)int{
	number := 1
	for i := value; i > 0; i--{
		number *= i
	}
	return number
}

func factorialRecursive(value int) int{
	if value == 1 {
		return 1
	} else{
		return value * factorialRecursive(value-1)
	}
}

func main(){
	// result := 1*2*3*4*5
	// fmt.Println(result)
	fmt.Println(factorialLoop(5))

	fmt.Println(factorialRecursive(5))
}