// arrays
package main

import "fmt"

func main() {
	var nombres [3]string
	nombres[0] = "Kevin"
	nombres[1] = "Alejandro"
	nombres[2] = "García"

	fmt.Println(nombres)
	fmt.Println(nombres[0])

	edades := [3]int{21, 22, 23}
	fmt.Println(edades)

	// Recorrer un array
	for i := 0; i < len(edades); i++ {
		fmt.Println(edades[i])
	}

	// Recorrer un array con range
	for i, v := range edades {
		fmt.Println(i, v)
	}

	// Recorrer un array con range sin usar el indice
	for _, v := range edades {
		fmt.Println(v)
	}

	// Sumar todos los elementos de un array
	suma := 0
	for _, v := range edades {
		suma += v
	}
	fmt.Println(suma)

	colores := [...]string{"rojo", "azul", "verde"}
	fmt.Println(colores, len(colores))
}