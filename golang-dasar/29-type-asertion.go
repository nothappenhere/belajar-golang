package main
import "fmt"

func random() interface{} {
	return true
}


func main(){
	var result any = random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) //! PANIC
	// fmt.Println(resultInt)
	
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown", value)
	}
}