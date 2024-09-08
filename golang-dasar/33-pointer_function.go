package main
import "fmt"

type Address struct{
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main(){
	//* Bisa seperti ini:
	// var address *Address = &Address{}
	// changeCountryToIndonesia(address)

	// * Atau seperti ini:
	var address Address = Address{}
	changeCountryToIndonesia(&address)

	fmt.Println(address)
}