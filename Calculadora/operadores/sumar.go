package operadores

import (
	"fmt"
	"strconv"
	"strings"
)

// Sumar es una función que suma dos enteros y devuelve un entero
func Sumar(operacion string) int {
	valores := strings.Split(operacion, "+") // Lo que hace es separar el string en un array

	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, error := strconv.Atoi(valores[i])

		if error != nil { // Si hay un error, lo imprime y termina el programa
			fmt.Println(error)
			fmt.Println("Tienes que ingresar números enteros separados por un + \nEjemplo: 1+2+3+4+5")
			return 0
		}

		resultado += num
	}

	return resultado
}
