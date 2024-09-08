package helper

import "fmt"

var version = "1.0.0" //* Tidak bisa diakses diluar package
var Application = "golang" //* Bisa diakses dari luar package

//* Tidak bisa diakses diluar package
func sayGoodbye(name string) string{ 
	return "Bye " + name
}

//* Bisa diakses dari luar package
func SayHello(name string) string{
	return "Hello " + name
}

func Contoh(){
	sayGoodbye("Akbar")
	fmt.Println(version)
}