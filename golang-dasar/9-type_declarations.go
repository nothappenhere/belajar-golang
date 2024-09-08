package main
import "fmt"

func main(){
	//* Membuat alias untuk tipe data, gunakan keyword "type":
	type NoKTP string

	var ktp NoKTP = "123456789"
	fmt.Println(ktp)

	var contohKTP string = "987654321"
	var ktp2 NoKTP = NoKTP(contohKTP)
	fmt.Println(ktp2)
}