package main

import "fmt"

func main(){
	fmt.Println("Hello World!")

	firstName := "Rizky"
	lastName := "Akbar"
	
	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}