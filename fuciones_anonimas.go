package main

import "fmt"

/*
Las funciones anónimas en Go son funciones que no tienen un nombre explícito.
A menudo se utilizan cuando se necesita una función para un propósito específico
y no es necesario darle un nombre o reutilizarla en otras partes del código.

Las funciones anónimas en Go son funciones de primera clase,
lo que significa que pueden ser tratadas como cualquier otro valor: pueden ser asignadas a variables,
pasadas como argumentos a otras funciones, y retornadas desde otras funciones.
*/

//func(parameters) (returnType) {
//    // Cuerpo de la función
//}

func Anonima() {
	// Función anónima asignada a una variable
	saludar := func(nombre string) {
		fmt.Println("Hola", nombre)
	}

	// Llamando a la función anónima
	saludar("Juan")
}

/*
La función anónima se define y se asigna a la variable saludar.

Luego, puedes llamar a esa función usando saludar("Juan"), como si fuera una función normal.
*/
