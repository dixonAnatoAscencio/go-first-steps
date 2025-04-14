package main

import "fmt"

/*
Una firma de función incluye los siguientes componentes:

Nombre de la función: El identificador único para la función.

Parámetros: Los valores que se pasan a la función para su procesamiento. Pueden ser opcionales si la función no los usa.

Tipo de retorno: El tipo de valor que la función devolverá después de su ejecución. Puede ser uno o más valores.
*/

// Firma de la función: suma toma dos enteros como parámetros y devuelve un entero
func suma(a int, b int) int {
	return a + b
}

func sumarDosNumeros() {
	resultado := suma(5, 3)
	fmt.Println("La suma es:", resultado)
}
