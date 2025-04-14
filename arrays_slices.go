package main

import "fmt"

func manejarArraysYSlices() {
	// Arrays
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	arr3 := arr2[1:4]
	fmt.Println(arr3)

	arr4 := arr2[1:4:5]
	fmt.Println(arr4)

	// Slices
	myArray := [4]string{"apple", "banana", "cherry", "date"}
	mySlice := myArray[1:3]
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))

	slice := make([]string, 5)
	slice[0] = "apple"
	slice[1] = "banana"
	fmt.Println(len(slice))
}
