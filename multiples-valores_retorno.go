package main

import "fmt"

//Go permite que las funciones devuelvan varios valores,
// lo cual es útil para devolver múltiples resultados de una función.
//  Esto se puede hacer sin necesidad de crear un objeto o estructura adicional.

// Firma de la función: devuelve dos valores, el resultado y un mensaje
func dividir(a int, b int) (int, string) {
	if b == 0 {
		return 0, "Error: División por cero"
	}
	return a / b, "Operación exitosa"
}

func Retorno() {
	resultado, mensaje := dividir(10, 2)
	fmt.Println(resultado, mensaje)
}

func calcularArea(longitud int, ancho int) (area int) {
	area = longitud * ancho
	return
}

func ResultadoArea() {
	area := calcularArea(5, 3)
	fmt.Println("Área del rectángulo:", area)

	/*
		En la función calcularArea, se ha nombrado el valor de retorno como area.
		No es necesario usar return area, ya que Go devolverá automáticamente la variable area
		cuando se alcance el final de la función.
	*/
}

/*
En Go, también puedes nombrar los parámetros y los valores de retorno.
Esto hace que el código sea más legible, ya que puedes ver qué parámetros o valores se están manejando.
*/
