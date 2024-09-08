package main
import "fmt"

func main(){
	var plus1 = 10
	var plus2 = 10
	plus3 := plus1 + plus2
	fmt.Println(plus3)
	//* Augmeted Assignments 
	plus3 += 10
	fmt.Println(plus3)

	fmt.Println("")

	var minus1 = 15
	var minus2 = 10
	minus3 := minus1 - minus2
	fmt.Println(minus3)
	//* Augmeted Assignments
	minus3 -= 10
	fmt.Println(minus3)

	fmt.Println("")

	var mul1 = 20
	var mul2 = 10
	mul3 := mul1 * mul2
	fmt.Println(mul3)
	//* Augmeted Assignments
	mul3 *= 10
	fmt.Println(mul3)

	fmt.Println("")

	var div1 = 25
	var div2 = 10
	div3 := div1 / div2
	fmt.Println(div3)
	//* Augmeted Assignments
	div3 /= 10
	fmt.Println(div3)

	fmt.Println("")

	var mod1 = 10
	var mod2 = 4
	mod3 := mod1 % mod2
	fmt.Println(mod3)
	//* Augmeted Assignments
	mod3 %= 10
	fmt.Println(mod3)

	fmt.Println("")

	var i = 1
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)
}