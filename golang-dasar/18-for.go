package main
import "fmt"

func main(){
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke -", counter)
		counter++
	}

	fmt.Println("")

	for number := 1; number <= 10; number++ {
		fmt.Println("Perulangan ke -", number)
	}

	fmt.Println("")

	//* Perulangan dengan cara manual:
	names := []string{"Muhammad", "Rizky", "Akbar"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	//* Perulangan dengan menggunakan "range":
	for index, value := range names {
		fmt.Println("Index ke -", index, "=", value)
	}
	//* Jika tidak membutuhkan index/value, ganti dengan (_):
	for _, value := range names {
		fmt.Println(value)
	}
}