package main
import "fmt"

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name){
		fmt.Println("You're blocked", name)
	} else{
		fmt.Println("Welcome", name)
	}
}

func main(){
	blacklist := func(name string)bool{
		return name == "anjing"
	}

	registerUser("akbar", blacklist)
	registerUser("anjing", func(name string)bool{
		return name == "anjing"
	})
}