package main;

import "fmt";

func main (){
	a:=10;
	b:=3;

	// Suma
	resultado := a + b;
	fmt.Println("Suma:", resultado);

	// Resta
	resultado = a - b;
	fmt.Println("Resta:", resultado);

	// Multiplicación
	resultado = a * b;
	fmt.Println("Multiplicación:", resultado);

	// División
	resultado = a / b;
	fmt.Println("División:", resultado);

	// Modulo
	resultado = a % b;
	fmt.Println("Modulo:", resultado);

	// Incremental
	a++;
	fmt.Println("Incremental:", a);

	// Decremental
	a--;
	fmt.Println("Decremental:", a);

	// Asignación
	a += 10;
	fmt.Println("Asignación:", a);

	// Comparación
	fmt.Println("Comparación:", a == b);

	// Diferencia
	fmt.Println("Diferencia:", a != b);

	// Mayor que
	fmt.Println("Mayor que:", a > b);
}