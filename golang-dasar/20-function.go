package main
import "fmt"

//* Function tanpa parameter
func sayHello(){
	fmt.Println("Hello World!")
}

//* Function dengan 2 parameter
func greeting(firstName string, lastName string){
	fmt.Println("Hallo", firstName, lastName)
}

//* Function dengan 2 parameter yang mengembalikan nilai "int"
func sum(a int, b int) int{
	result := a + b
	return result
}

//* Function dengan 2 parameter yang mengembalikan banyak nilai "string, string"
/* 
? 	func getFullName() (string, string) {
?	return "Rizky", "Akbar"
?	}
*/
func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func main(){
	//* Memanggil function dengan tanpa parameter
	sayHello()

	//* Memanggil function dengan 2 parameter
	greeting("Rizky", "Akbar")

	//* Memanggil function dengan 2 parameter, bisa dengan salah satu dari 2 cara berikut:
	fmt.Println(sum(5,5))
	
	result := sum(4, 5)
	fmt.Println(result)

	//* Memanggil function dengan 2 parameter, yang mengembalikan banyak nilai:
	/*
	? firstName, lastName := getFullName()
	? fmt.Println(firstName, lastName)
	* Jika tidak membutuhkan salah satu nilainya, gunakan (_):
	? firstName, _ := getFullName()
	? fmt.Println(firstName)
	*/
	fmt.Println(getFullName("Rizky", "Akbar"))
}