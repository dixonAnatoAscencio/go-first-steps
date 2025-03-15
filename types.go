package main

import "fmt"

func main() {

	fmt.Println("Hello, World!" + " " + "I am learning Go programming language")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println(1 + 13)
	fmt.Println(7.0 / 3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//uint
	var a uint = 10
	fmt.Println(a)

	//complex number
	var b complex64 = 1 + 2i
	fmt.Println(b)

	//Variables declaration and initialization
	var c int
	c = 10
	fmt.Println(c)

	var d int = 20
	fmt.Println(d)

}
