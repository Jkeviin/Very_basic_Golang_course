package main

import "fmt"

func saludar(nombre string) string {
	return "Hola, " + nombre
}

func sumar(a int, b int) int {
	return a + b
}

func main() {
	nombre := "Kevin"
	edad := 21

	fmt.Println(saludar(nombre))
	fmt.Println(edad)

	resultado := sumar(10, 20)
	fmt.Println(resultado)
}
