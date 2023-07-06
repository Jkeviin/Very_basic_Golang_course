package main

import "fmt"

func main() {

	Carrera1 := new(Carrera)
	Carrera1.nombreCarrera = "Carrera de Go"
	Carrera1.duracion = 12
	Carrera1.nombre = "Curso profesional de Go"
	Carrera1.url = "https://google.com"
	Carrera1.habilidad = []string{"Backend", "APIs", "Testing", "Devops"}

	Carrera1.Inscribirse("Juanito")
}

// curso es una estructura
type Curso struct {
	nombre    string
	url       string
	habilidad []string
}

// inscribirse es un metodo de la estructura curso
func (c Curso) Inscribirse(nombre string) {
	fmt.Printf("La persona %s se ha inscrito al curso %s \n", nombre, c.nombre)
}

// carreras es una estructura que hereda de curso
type Carrera struct {
	nombreCarrera string
	duracion      int
	Curso
}
