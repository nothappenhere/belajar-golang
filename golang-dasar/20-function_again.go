package main
import "fmt"

//* named return value
func getFullName() (firstName, middleName, lastName string){
	firstName = "Muhammad"
	middleName = "Rizky"
	lastName = "Akbar"

	return firstName, middleName, lastName
}

//* variadic function
func sumAll(numbers ...int)int{
	total := 0
	for _, number := range numbers{
		total += number
	}
	return total
}

//* Function as value
func getGoodBye(name string)string{
	return "Bye " + name
}

//* Function as parameter
func sayHelloWithFilter(name string, filter func(string)string){
	filtered := filter(name)
	fmt.Println("Helo", filtered)
}
/* bisa menggunakan function diatas, atas dibawah berikut ini:
?	type Filter func(string)string
?	func sayHelloWithFilter(name string, filter Filter){
?		filtered := filter(name)
?		fmt.Println("Helo", filtered)
?	}
*/
func spamFilter(name string)string{
	if name == "Anjing"{
		return "..."
	} else {
		return name
	}
}


func main(){
	//* named return value
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)

	//* variadic function
	total := sumAll(10,20,30,40,50,60,70,80,90)
	fmt.Println(total)

	//* Function as value
	goodbye := getGoodBye
	fmt.Println(goodbye("Akbar"))

	//* Function as parameter
	sayHelloWithFilter("Akbar", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}