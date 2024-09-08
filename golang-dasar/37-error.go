package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0{
		return 0, errors.New("Tidak dapat dibagi dengan 0!")
	} else{
		return nilai / pembagi, nil
	}
}

func main(){
	hasil, err := pembagian(100, 0)
	if err == nil{
		fmt.Println("Hasilnya adalah", hasil)
	}else{
		fmt.Println("Error", err.Error())
	}
}