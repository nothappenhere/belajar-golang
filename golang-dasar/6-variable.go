package main
import "fmt"

func main(){
	//TODO: bisa seperti ini:
	//* var name string
	//* var age int

	//TODO: atau seperti ini:
	//* var name string = "Muhammad Rizky Akbar"
	//* var age int = 20

	//TODO: atau seperti ini:
	//* var name = "Muhammad Rizky Akbar"
	//* var age = 20

	//TODO: atau juga seperti ini (HANYA SEKALI DEKLARASI (:=)):
	name := "Muhammad Rizky Akbar"
	age := 20
	fmt.Println("Nama = ", name)
	fmt.Println("Umur = ", age, "Tahun")
	
	fmt.Println("")
	
	//TODO: tidak boleh lagi pakai (:=)
	name = "Rizky Akbar"
	age = 19
	fmt.Println("Nama = ", name)
	fmt.Println("Umur = ", age, "Tahun")

	fmt.Println("")

	//TODO: multiple variable
	var (
		firstName = "Muhammad"
		middleName = "Rizky"
		lastName = "Akbar"
	)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}