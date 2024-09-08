package main
import "fmt"

func main(){
	var nilai32 int32 = 32768
	fmt.Println("Nilai int32 = ", nilai32)

	var nilai64 = int64(nilai32)
	fmt.Println("Nilai int64 = ", nilai64)
	
	//* Terjadi overflow, sehingga nilai-nya akan menjadi minus:
	nilai16 := int16(nilai32)
	fmt.Println("Nilai int16 = ", nilai16)

	fmt.Println("")

	var name string = "Muhammad Rizky Akbar"
	fmt.Println("Nama lengkap = ", name)

	//* Nilai-nya akan menjadi BYTE:
	var e uint8 = name[0]
	fmt.Println("Index ke-0 BYTE = ", e)

	/**
	 * * Sehingga perlu dilakukan konversi tipe data menjadi STRING,
	 * * Bisa dengan salah satu dari 2 cara berikut:
	 * TODO eString := string(e)
	 * TODO eString = string(name[0])
	 */
	eString := string(e)
	 //? eString = string(name[0])
	fmt.Println("Index ke-0 STRING = ", eString)
}