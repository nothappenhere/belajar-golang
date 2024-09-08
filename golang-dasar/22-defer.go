package main
import "fmt"

func logging(){
	fmt.Println("Function ini akan dieksekusi terakhir")
}

func runApp(){
	//* Bisa seperti ini walaupun terdapat logic error akan tetap dieksekusi atau seperti dibawah-nya:
	defer logging()
	fmt.Println("Run application")
	// logging() //! dengan catatan, jika logic diatasnya terdapat eror function ini tidak akan terpanggil
}

func main(){
	runApp()
}