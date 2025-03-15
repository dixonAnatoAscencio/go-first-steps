package main

import "fmt"

const name string = "John Doe"

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

	fmt.Println(c + d)

	//Short hand declaration
	e := 15
	fmt.Println(e)

	//Multiple variable declaration
	var (
		f = 50
		g = 50
	)
	fmt.Println(f + g)

	const PI = 3.14
	const radius = 7.0
	area := PI * radius * radius
	fmt.Println("el area es: ", area)

	fmt.Println(PI)

	fmt.Println(name)

	// Type conversion
	var h int = 10
	var i int32
	i = int32(h)
	fmt.Println(i)

	//multiple constant declaration
	const (
		r = 10
		l = 20
		k // k will have the same value as l
	)

	fmt.Println(r + l + k)
	fmt.Println(l)

}
