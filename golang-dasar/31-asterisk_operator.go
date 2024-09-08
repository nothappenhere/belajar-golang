package main
import "fmt"

type Address struct{
	City, Province, Country string
}

func main(){
	var address1 Address = Address{"Bandung", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 //* pointer, reference address1

	address2.City = "Karawang"

	fmt.Println(address1)
	fmt.Println(address2)
	
	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}