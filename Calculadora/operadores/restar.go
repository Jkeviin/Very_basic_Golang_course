package operadores

import (
	"fmt"
	"strconv"
	"strings"
)

// Restar es una función que resta dos enteros y devuelve un entero
func Restar(operacion string) int {
	valores := strings.Split(operacion, "-") // Lo que hace es separar el string en un array

	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, error := strconv.Atoi(valores[i])

		if error != nil { // Si hay un error, lo imprime y termina el programa
			fmt.Println(error)
			fmt.Println("Tienes que ingresar números enteros separados por un - \nEjemplo: 1-2-3-4-5")
			return 0
		}

		if resultado == 0 {
			resultado = num
			continue
		}

		resultado -= num
	}

	return resultado
}
