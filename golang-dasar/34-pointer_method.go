package main
import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func main() {
	akbar := Man{"Akbar"}
	akbar.Married()

	fmt.Println(akbar.Name)
}