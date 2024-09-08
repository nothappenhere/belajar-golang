package main
import "fmt"

func main(){
	name := "Rizky Akbar"

	if (name == "Rizky") {
		fmt.Println("hello rizky")
	} else if (name == "Akbar") {
		fmt.Println("hello akbar")		
	} else {
		fmt.Println("Nama kamu siapa?")
	}

	if length := len(name); length <= 5 {
		fmt.Println("Namanya sudah pas")		
	} else {
		fmt.Println("Namanya kepanjangan!")
	}
}