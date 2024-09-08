package main
import "fmt"

func main(){
	name := "Akbar"

	switch name {
	case "Muhammad":
		fmt.Println("hallo muhammad.")
	case "Rizky":
		fmt.Println("hallo rizky.")
	case "Akbar":
		fmt.Println("hallo akbar.")
	default:
		fmt.Println("hi, nama kamu siapa?.")
	}

	switch length := len(name); length <= 5 {
	case true :
		fmt.Println("Namanya sudah pas")		
	case false :
		fmt.Println("Namanya kepanjangan!")
	}

	length := len(name)
	switch {
	case length <= 5 :
		fmt.Println("Namanya sudah pas nih")		
	case length > 5 :
		fmt.Println("Namanya panjang banget!")
	}
}