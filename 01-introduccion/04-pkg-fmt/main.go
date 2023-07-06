package main;

import "fmt";

func main (){
	hola := "Hola";
	mundo := "Mundo";

	fmt.Println(hola);
	fmt.Println(mundo);

	nombre := "Kevin";
	edad := 21;

	fmt.Printf("Hola, %s, tienes %d años \n", nombre, edad);  // %s = string, %d = int
	fmt.Printf("Hola, %v, tienes %v años \n", nombre, edad); // %v = variable (string, int, float, bool, etc)

	mensaje := fmt.Sprintf("Hola, %s, tienes %d años \n", nombre, edad); // Se usa Sprintf para guardar el mensaje en una variable

	fmt.Println(mensaje);

	fmt.Printf("Hola, %q \n", mensaje); // %q = string con comillas

	fmt.Printf("Hola, %T \n", mensaje); // %T = tipo de dato

	fmt.Printf("Hola, %% \n"); // %% = imprimir %

	// Pedir datos al usuario
	var nombre2 string;
	var edad2 int;

	fmt.Print("Ingresa tu nombre: ");
	fmt.Scanf("%s \n", &nombre2); // & = referencia de la variable
	fmt.Print("Ingresa tu edad: ");
	fmt.Scanf("%d \n", &edad2);

	fmt.Printf("Hola, %s, tienes %d años \n", nombre2, edad2);
}