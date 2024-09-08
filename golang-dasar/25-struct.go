package main
import "fmt"

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main(){
	//* Bisa seperti ini: 
	var akbar Customer
	akbar.Name = "Akbar"
	akbar.Address = "Bandung"
	akbar.Age = 20

	fmt.Println(akbar)
	// fmt.Println(akbar.Name)
	// fmt.Println(akbar.Address)
	// fmt.Println(akbar.Age)

	fmt.Println("")

	//* Bisa seperti ini: 
	rizky := Customer{
		Name : "Rizky",
		Address : "Bandung",
		Age : 20,
	}
	fmt.Println(rizky)

	fmt.Println("")

	//* Bisa seperti ini: 
	muhammad := Customer{"Muhammad", "Bandung", 20}
	fmt.Println(muhammad)

	muhammad.sayHello("Akbar")
	rizky.sayHello("Akbar")
	akbar.sayHello("Akbar")
}