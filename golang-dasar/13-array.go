package main
import "fmt"

func main(){
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Rizky"
	names[2] = "Akbar"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println("")

	var values = [3]int{
		90,
		80,
		// 70, //* Wajib diakhiri coma (,)
		// 70 //! Akan error, karena tidak diakhiri coma (,)
	}
	// var values = [3]int{10,10,10} //* Tidak wajib diakhiri coma (,)

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println("")

	var scores = [...]int{
		50,
		40,
		60,
		80,
		90,
		12,
		100,
	}
	fmt.Println(len(scores))
	fmt.Println(scores[1])
	scores[1] = 15
	fmt.Println(scores[1])
}