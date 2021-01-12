package main

import (
	"fmt"
)

var a int = 10
var b int = 2

func main() {
	d := a + b
	e := a - b
	f := a / b
	g := a * b
	h := a % b

	fmt.Printf("The addition operation of %d and %d is %d\n", a, b, d)
	fmt.Printf("The substraction operation of %d and %d is %d\n", a, b, e)
	fmt.Printf("The division operation of %d and %d is %d\n", a, b, f)
	fmt.Printf("The multiplication operation of %d and %d is %d\n", a, b, g)
	fmt.Printf("The modulo operation of %d and %d is %d\n", a, b, h)

}
