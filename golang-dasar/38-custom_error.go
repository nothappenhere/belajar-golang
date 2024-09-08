package main


import (
	"fmt"
)

type validationError struct{
	Massage string
}

func (v *validationError) Error() string{
	return v.Massage
}


type notFoundError struct{
	Massage string
}

func (n *notFoundError) Error() string{
	return n.Massage
}


func SaveData(id string, data any) error{
	if id == ""{
		return &validationError{Massage: "Validation Error!"}
	} else if id != "Akbar" {
		return &notFoundError{Massage: "Not Found Error"}
	}
	return nil
}

func main(){
	err := SaveData("4", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("NOt found error:", notFoundErr.Error())
		// } else{
		// 	fmt.Println("unknown error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("NOt found error:", finalError.Error())
		default:
			fmt.Println("unknown error:", finalError.Error())
		}
	}else{
		fmt.Println("Sukses")
	}
}