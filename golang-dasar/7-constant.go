package main
import "fmt"

func main(){
	const firstName string = "Muhammad"
	const lastName = "Rizky Akbar" 

	//! akan terjadi error:
	//* firstName = "Muhammad Rizky"
	//* lastName = "Akbar"

	//TODO: multiple constant
	const (
		name string = "Akbar"
		age = 20
	)
	fmt.Println("Nama = ", name)
	fmt.Println("Umur = ", age, "Tahun")
}