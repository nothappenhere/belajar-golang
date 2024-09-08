package main

import(
	"fmt"
	"strings"
)

func main(){
	var slice1 []string = []string{
		"akbar",
		"rizky",
		"muhammad",
	}
	for idx, val := range(slice1){
		fmt.Println("index", idx, "=", val)
	}


	fmt.Println(strings.Contains("Muhammad Rizky Akbar", "Rizky"))
	fmt.Println(strings.ToLower("Muhammad Rizky Akbar"))
	fmt.Println(strings.ToUpper("Muhammad Rizky Akbar"))
	fmt.Println(strings.Split("Muhammad Rizky Akbar", " "))
	fmt.Println(strings.Trim("         Muhammad Rizky Akbar         ", " "))
	fmt.Println(strings.ReplaceAll("Muhammad Rizky Akbar", "Rizky", "Akbar"))
}