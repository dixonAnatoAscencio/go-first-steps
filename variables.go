package main

import "fmt"

func trabajarConVariables() {
	var a uint = 10
	fmt.Println(a)

	var b complex64 = 1 + 2i
	fmt.Println(b)

	var c int
	c = 10
	fmt.Println(c)

	var d int = 20
	fmt.Println(d)
	fmt.Println(c + d)

	e := 15
	fmt.Println(e)

	var (
		f = 50
		g = 50
	)
	fmt.Println(f + g)

	var h int = 10
	var i int32
	i = int32(h)
	fmt.Println(i)

	var myInt int = 65
	var myFloat float64 = float64(myInt)
	var myString string = string(myInt)
	fmt.Println(myFloat)
	fmt.Println(myString)
}
