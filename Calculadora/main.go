package main

import (
	"bufio"
	"fmt"
	"operadores/operadores"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Ingresa la operación a realizar: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text() // 5 + 5, 5 - 5, 5 * 5, 5 / 5
	resultado := 0

	if strings.Contains(operacion, "+") {
		resultado = operadores.Sumar(operacion)
	} else if strings.Contains(operacion, "-") {
		resultado = operadores.Restar(operacion)
	} else if strings.Contains(operacion, "*") {
		resultado = operadores.Multiplicar(operacion)
	} else if strings.Contains(operacion, "/") {
		resultado = operadores.Dividir(operacion)
	} else {
		fmt.Println("Tienes que ingresar números enteros separados por un +, -, * o / \nEjemplo: 1+2+3+4+5")
		return
	}

	fmt.Println("El resultado es: " + strconv.Itoa(resultado))

	// Dar enter para terminar el programa
	scanner.Scan()

}
