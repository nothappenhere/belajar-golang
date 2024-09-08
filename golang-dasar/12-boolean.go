package main
import "fmt"

func main()  {
	nilaiAkhir := 90
	absensi := 85

	lulusUjianAkhir := nilaiAkhir > 90
	lulusAbsensi := absensi > 80

	lulus := lulusUjianAkhir && lulusAbsensi
	fmt.Println(lulus)
}