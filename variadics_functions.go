/*
Go también soporta funciones variádicas, que son funciones que pueden recibir un número variable de argumentos
de un tipo específico.
*/
package main

import "fmt"

/*
...int indica que la función puede recibir un número variable de enteros como parámetro.

En el cuerpo de la función, numeros se comporta como un slice de enteros ([]int).
*/

// Firma de la función: parámetros variádicos
func sumarNumeros(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

func resultSumaVariadics() {
	resultado := sumarNumeros(1, 2, 3, 4, 5)
	fmt.Println("La suma es:", resultado)
}
