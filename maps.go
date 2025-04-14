package main

import "fmt"

/*
Un map es una estructura de datos que te permite almacenar pares clave-valor,
similar a un diccionario en Python o un objeto en JavaScript.

Clave (key): Es única y sirve para acceder a su valor.
Valor (value): Es el dato asociado a esa clave.
*/

func trabajarConMaps() {
	// Crear un map
	personas := map[string]int{
		"Ana":  25,
		"Juan": 30,
	}

	// Acceder a un valor
	fmt.Println("Edad de Ana:", personas["Ana"])

	// Agregar un nuevo valor
	personas["Luis"] = 40

	// Verificar existencia
	edad, existe := personas["Pedro"]
	if existe {
		fmt.Println("Edad de Pedro:", edad)
	} else {
		fmt.Println("Pedro no está en el mapa")
	}

	// Eliminar
	delete(personas, "Juan")

	// Recorrer
	for nombre, edad := range personas {
		fmt.Println(nombre, "tiene", edad, "años")
	}
}
