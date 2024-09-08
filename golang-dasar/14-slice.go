package main
import "fmt"

func main(){
	names := [...]string{"Muhammad", "Rizky", "Akbar", "Akbar", "Rizky", "Muhammad"}

	//* Deklarasi manual:
	// var slice1 []string = names[4:6]
	slice1 := names[4:6]
	fmt.Println(slice1)

	//* Deklarasi manual:
	// var slice2 []string = names[:3]
	slice2 := names[:3]
	fmt.Println(slice2)

	//* Deklarasi manual:
	var slice3 []string = names[3:]
	// slice3 := names[3:]
	fmt.Println(slice3)

	//* Deklarasi manual:
	var slice4 []string = names[:]
	// slice4 := names[:] 
	fmt.Println(slice4)

	fmt.Println("")

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	fmt.Println("")

	var newSLice []string = make([]string, 2, 5)
	newSLice[0] = "Rizky"
	newSLice[1] = "Akbar"
	// newSLice[2] = "Akbar" //! Error, harusnya menggunakan append()

	fmt.Println(newSLice)
	fmt.Println(len(newSLice))
	fmt.Println(cap(newSLice))
	
	newSLice2 := append(newSLice, "Akbar")
	fmt.Println(newSLice2)
	fmt.Println(len(newSLice2))
	fmt.Println(cap(newSLice2))
	

	fmt.Println("")

	fromSlcie := days[:]
	toSlice := make([]string, len(fromSlcie), cap(fromSlcie))

	copy(toSlice, fromSlcie)
	fmt.Println(fromSlcie)
	fmt.Println(toSlice)

	fmt.Println("")

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}