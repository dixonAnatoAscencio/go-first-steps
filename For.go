package main

import "fmt"

/*
En Go, el for es la única estructura de bucle (no hay while ni do-while).
 Pero es súper flexible y puede comportarse como un for, un while, o un loop infinito.
*/

func TrabajarCicloFor() {
	// 1. For clásico (como en C/Java)
	fmt.Println("For clásico:")
	for i := 0; i < 5; i++ {
		fmt.Println("i:", i)
	}

	// 2. For tipo while
	fmt.Println("\nFor tipo while:")
	j := 0
	for j < 3 {
		fmt.Println("j:", j)
		j++
	}

	// 3. For infinito (con break)
	fmt.Println("\nFor infinito (con break):")
	k := 0
	for {
		if k >= 2 {
			break
		}
		fmt.Println("k:", k)
		k++
	}

	// 4. For con continue
	fmt.Println("\nFor con continue:")
	for l := 0; l < 5; l++ {
		if l%2 == 0 {
			continue // salta los pares
		}
		fmt.Println("l (impares):", l)
	}

	// 5. For range con slice
	fmt.Println("\nFor range con slice:")
	frutas := []string{"manzana", "banana", "pera"}
	for index, fruta := range frutas {
		fmt.Printf("índice %d: %s\n", index, fruta)
	}

	// 6. For range con map
	fmt.Println("\nFor range con map:")
	edades := map[string]int{"Ana": 25, "Luis": 30}
	for nombre, edad := range edades {
		fmt.Printf("%s tiene %d años\n", nombre, edad)
	}

	// 7. For range con string (Unicode)
	fmt.Println("\nFor range con string:")
	texto := "Go ❤️"
	for i, r := range texto {
		fmt.Printf("índice %d: carácter '%c'\n", i, r)
	}

	// Bucle infinito
	for {
		fmt.Println("Este bucle se ejecutará indefinidamente.")

		// Para romper el bucle después de un tiempo
		var respuesta string
		fmt.Println("¿Quieres salir del bucle? (si/no)")
		fmt.Scanln(&respuesta)
		if respuesta == "si" {
			fmt.Println("Saliendo del bucle...")
			break // Rompe el bucle
		}
	}
}
