package main

import (
	"fmt"
	"strconv"
)

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
		n // k will have the same value as l
	)

	fmt.Println(r + l + n)
	fmt.Println(l)

	var myInt int = 65
	var myFloat float64 = float64(myInt)
	var myString string = string(myInt)
	fmt.Println(myFloat)
	fmt.Println(myString)

	//Operadores aritméticos
	var j int = 10
	var k int = 20
	fmt.Println(j + k)
	fmt.Println(j - k)
	fmt.Println(j * k)
	fmt.Println(j / k)
	fmt.Println(j % k)

	strconv.Itoa(42)
	strconv.ParseInt("100", 10, 64)
	strconv.ParseFloat("100.5", 64)
	strconv.ParseBool("true")

	//Operadores de comparación
	fmt.Println(j == k)
	fmt.Println(j != k)
	fmt.Println(j > k)
	fmt.Println(j < k)
	fmt.Println(j >= k)
	fmt.Println(j <= k)

	//Arrays
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
	//Array initialization
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	//Array length
	fmt.Println(len(arr2))
	//Array slice
	arr3 := arr2[1:4]
	fmt.Println(arr3)
	//Array slice with step
	arr4 := arr2[1:4:5]
	fmt.Println(arr4)
	//Array slice with step and length

	myArray := [4]string{"apple", "banana", "cherry", "date"}
	mySlice := myArray[1:3]
	fmt.Println(mySlice)      // Output: [banana cherry]
	fmt.Println(len(mySlice)) // Output: 2

	// := Short variable declaration, sirve para declarar e incializar una variable en una sola linea
	// y Go infiere automáticamente el tipo de esa variable según el valor que le des.

	//declaracion tradicional
	/*
		var x int = 10  Declarás x, indicás su tipo (int) y le das un valor (10).

		Forma con := (declaración corta):
		x := 10  Declarás x, Go infiere que es un int y le das un valor (10).

		Reglas:
		1. Solo funciona dentro de funciones (no se puede usar fuera del main() o de una función).
		2. Al menos una variable del lado izquierdo debe ser nueva (si no, da error).
		3. El tipo se infiere automáticamente por el compilador según el valor a la derecha.

	*/

	/*
		Slices es una vista dinamica sobre una secuencia de elementos en un array, no tienen un tamaño fijo
		internamente apuntan a un array, pero no tienen un tamaño fijo, por lo que pueden crecer y decrecer

		tienen 3 atributos:
		1. Pointer: referencia al primer elemento del array subyacente
		2. Length: longitud del slice
		3. Capacity: capacidad del slice, que es la longitud del array subyacente
		4. El slice es un puntero a un array, por lo que si modificas el slice, modificas el array
	*/

	var slice []string

	slice = make([]string, 5)
	slice[0] = "apple"
	slice[1] = "banana"
	fmt.Println(len(slice))

	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}
	//fmt.Println(p.Name)
	//fmt.Println(p.Age)
	//p.Name = "Dixon"
	//p.Age = 30
	fmt.Println(p)

}
