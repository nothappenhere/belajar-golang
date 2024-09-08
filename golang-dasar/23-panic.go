package main
import "fmt"

func endApp(){
	fmt.Println("End App")
}

func runApp(error bool){
	defer endApp()

	if error == true{
		panic("Upss, terdapat ERROR")
	}
}

func main(){
	runApp(true)
	// runApp(false)
}