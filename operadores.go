package main

import (
	"fmt"
	"strconv"
)

func usarOperadores() {
	var j int = 10
	var k int = 20

	fmt.Println(j + k)
	fmt.Println(j - k)
	fmt.Println(j * k)
	fmt.Println(j / k)
	fmt.Println(j % k)

	fmt.Println(j == k)
	fmt.Println(j != k)
	fmt.Println(j > k)
	fmt.Println(j < k)
	fmt.Println(j >= k)
	fmt.Println(j <= k)

	// Conversión de strings y números
	strconv.Itoa(42)
	strconv.ParseInt("100", 10, 64)
	strconv.ParseFloat("100.5", 64)
	strconv.ParseBool("true")
}
