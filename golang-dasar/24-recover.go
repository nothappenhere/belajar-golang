package main
import "fmt"

func endApp(){
	fmt.Println("End App")
	//* Cara penempatan "recover" yang benar
	massage := recover()
	fmt.Println("Terjadi error, pesan recover:", massage)
}

func runApp(error bool){
	defer endApp()

	if error == true{
		panic("Upss, terdapat ERROR")
	}

	//! cara penempatan "recover" yang salah
	// massage := recover()
	// fmt.Println("Terjadi error, pesan recover:", massage)
}

func main(){
	runApp(true)

	fmt.Println("Test, apakah ini akan ter-eksekusi?")
}