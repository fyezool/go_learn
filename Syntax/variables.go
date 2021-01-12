package main

import (
	"fmt"
)

/*
3 ways to declare variables in golang

1. using var keyword
var name string = "Using var keyword"

2. using shorthand syntax
data := 42

3. using var() method for multiple var
 var (
	 a = 1
	 b = 2
	 c = 3
 )
*/

var (
	a = 1
	b = 1
	c = 1
)

func main() {
	data := "Box this lap"            // shorthand style
	fmt.Println(a, b, c)              //Print value of var a,b,c
	fmt.Printf("%T %T %T\n", a, b, c) //print type of abc var
	fmt.Println(data)                 //print value of data
	fmt.Printf("%s %T\n", data, data) //print var value with data type %s %T
}
