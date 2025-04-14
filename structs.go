package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func trabajarConStructs() {
	p := Person{Name: "John", Age: 30}
	fmt.Println(p)
}
