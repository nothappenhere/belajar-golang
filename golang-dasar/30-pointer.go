package main
import "fmt"

type Address struct{
	City, Province, Country string
}

func main(){
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 Address = address1 //* Copy value address1
	var address3 *Address = &address1 //* pointer, reference address1

	address2.City = "Karawang"
	address3.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}