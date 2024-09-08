package main

import (
	"golang-dasar/helper"
	"fmt"
)

func main(){
	result := helper.SayHello("Akbar")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) //* Tidak bisa diakses
	// fmt.Println(helper.sayGoodbye("Akbar")) //* Tidak bisa diakses
	fmt.Println(helper.Contoh())

	
}

