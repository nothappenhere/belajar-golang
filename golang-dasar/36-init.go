package main

import (
	"golang-dasar/database"
	_"golang-dasar/internal" //* Blank identifier
	"fmt"
)

func main(){
	fmt.Println(database.GetDB())
}