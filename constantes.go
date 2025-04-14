package main

import "fmt"

const name string = "John Doe"

func trabajarConConstantes() {
	const PI = 3.14
	const radius = 7.0
	area := PI * radius * radius
	fmt.Println("El área es:", area)
	fmt.Println(PI)
	fmt.Println(name)

	const (
		r = 10
		l = 20
		n // igual a l
	)

	fmt.Println(r + l + n)
	fmt.Println(l)
}
