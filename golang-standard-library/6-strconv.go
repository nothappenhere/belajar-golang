package main

import(
	"fmt"
	"strconv"
)

func main(){
	result, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println("Error", err.Error())
	}else{
		fmt.Println("Hasilnya =", result)
	}
}