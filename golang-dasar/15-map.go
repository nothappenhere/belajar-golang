package main
import "fmt"

func main(){
	// var person map[string]string = map[string]string{}
	// person["name"] = "Akbar"
	// person["address"] = "Bandung"

	person := map[string]string{
		"name" : "Akbar",
		"address" : "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["default-value"])

	book := make(map[string]string)
	book["title"] = "Buku GOLANG"
	book["author"] = "Rizky Akbar"
	book["wrong"] = "Upss!"
	fmt.Println(book)

	// delete(book, "wrong")
	fmt.Println(len(book))
}