package main

import "fmt"

/*
¿Por qué usar punteros?
Para modificar valores desde funciones (pasar por referencia).

Para evitar copias innecesarias (eficiencia).

Para trabajar con estructuras de datos complejas (como structs).
*/

/*
Sintaxis clave
Expresión	Significado
*int	Un pointer a un entero.
&x	Dirección de memoria de la variable x.
*p	Valor al que apunta el puntero p.


*/

type Persona struct {
	Nombre string
	Edad   int
}

func Pointers() {
	x := 10
	var p *int = &x // p apunta a la dirección de x

	fmt.Println("Valor de x:", x)             // 10
	fmt.Println("Dirección de x:", &x)        // 0xc0000120a8 (por ejemplo)
	fmt.Println("Valor de p:", p)             // dirección de x
	fmt.Println("Valor al que apunta p:", *p) // 10

	*p = 20                             // cambia el valor de x a través del puntero
	fmt.Println("Nuevo valor de x:", x) // 20
}

func cambiarNombre(p *Persona) {
	p.Nombre = "Ana"
}

func Principal() {
	persona := Persona{Nombre: "Luis", Edad: 30}
	cambiarNombre(&persona)
	fmt.Println(persona.Nombre) // Ana
}

/*
¿Qué pasa si un puntero no apunta a nada?
Un puntero sin inicializar vale nil.
Esto significa que no apunta a ninguna dirección de memoria válida.
Esto puede causar un panic si intentas desreferenciarlo.


var p *int
fmt.Println(p) // nil
*/
